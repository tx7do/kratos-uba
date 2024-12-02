package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	timeUtil "github.com/tx7do/go-utils/timeutil"

	"kratos-uba/app/core/service/internal/data/ent"
	"kratos-uba/app/core/service/internal/data/ent/user"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-uba/api/gen/go/user/service/v1"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/core-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertEntToProto(in *ent.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:          in.ID,
		UserName:    in.UserName,
		RealName:    in.RealName,
		Email:       in.Email,
		Description: in.Description,
		Password:    in.Password,
		RoleId:      in.RoleID,
		Authority:   (*string)(in.Authority),
		CreateTime:  timeUtil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:  timeUtil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:  timeUtil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}
	return builder.Count(ctx)
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		if item != nil {
			item.Password = nil
		}
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	ret, err := r.data.db.Client().User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	ret, err := r.data.db.Client().User.Create().
		SetNillableUserName(req.User.UserName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableDescription(req.User.Description).
		SetNillableRoleID(req.User.RoleId).
		SetNillableAuthority((*user.Authority)(req.User.Authority)).
		SetPassword(cryptoPassword).
		SetCreateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.Client().User.UpdateOneID(req.Id).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableDescription(req.User.Description).
		SetNillableRoleID(req.User.RoleId).
		SetNillableAuthority((*user.Authority)(req.User.Authority)).
		SetPassword(cryptoPassword).
		SetUpdateTime(time.Now())

	ret, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	u := r.convertEntToProto(ret)
	if u != nil {
		u.Password = nil
	}

	return u, err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	ret, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UserNameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		return &v1.VerifyPasswordResponse{
			Result: v1.VerifyPasswordResult_ACCOUNT_NOT_EXISTS,
		}, v1.ErrorUserNotFound("用户未找到")
	}

	bMatched := crypto.CheckPasswordHash(req.GetPassword(), *ret.Password)
	if !bMatched {
		return &v1.VerifyPasswordResponse{
			Result: v1.VerifyPasswordResult_WRONG_PASSWORD,
		}, v1.ErrorUserNotFound("密码错误")
	}

	return &v1.VerifyPasswordResponse{
		Result: v1.VerifyPasswordResult_SUCCESS,
	}, nil
}
