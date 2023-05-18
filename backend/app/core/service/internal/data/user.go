package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-bi/app/core/service/internal/biz"
	"kratos-bi/app/core/service/internal/data/ent"
	"kratos-bi/app/core/service/internal/data/ent/user"

	"kratos-bi/gen/api/go/common/pagination"
	v1 "kratos-bi/gen/api/go/user/service/v1"

	"kratos-bi/pkg/util/crypto"
	"kratos-bi/pkg/util/entgo"
	paging "kratos-bi/pkg/util/pagination"
	util "kratos-bi/pkg/util/time"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
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
		CreateTime:  util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(in.UpdateTime),
		DeleteTime:  util.UnixMilliToStringPtr(in.DeleteTime),
	}
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.User.Query()

	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder1 = builder1.Where(cond)
		}
	}
	if len(orderCond) != 0 {
		for _, cond := range orderCond {
			builder1 = builder1.Order(cond)
		}
	} else {
		builder1.Order(ent.Desc(user.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder1.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	users, err := builder1.All(ctx)
	if err != nil {
		return nil, err
	}

	builder2 := r.data.db.User.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder2 = builder2.Where(cond)
		}
	}
	count, err := builder2.
		Select(user.FieldID).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.User, 0, len(users))
	for _, u := range users {
		item := r.convertEntToProto(u)
		items = append(items, item)
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	po, err := r.data.db.User.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	po, err := r.data.db.User.Create().
		SetNillableUserName(req.User.UserName).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableDescription(req.User.Description).
		SetNillableRoleID(req.User.RoleId).
		SetNillableAuthority((*user.Authority)(req.User.Authority)).
		SetPassword(ph).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	ph, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	builder := r.data.db.User.UpdateOneID(req.Id).
		SetNillableRealName(req.User.RealName).
		SetNillableEmail(req.User.Email).
		SetNillableDescription(req.User.Description).
		SetNillableRoleID(req.User.RoleId).
		SetNillableAuthority((*user.Authority)(req.User.Authority)).
		SetPassword(ph).
		SetUpdateTime(time.Now().UnixMilli())

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *UserRepo) Delete(ctx context.Context, req *v1.DeleteUserRequest) (bool, error) {
	err := r.data.db.User.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	ret, err := r.data.db.User.
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