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

	"kratos-bi/pkg/util/entgo"
	paging "kratos-bi/pkg/util/pagination"
	util "kratos-bi/pkg/util/time"
)

var _ biz.ApplicationRepo = (*ApplicationRepo)(nil)

type ApplicationRepo struct {
	data *Data
	log  *log.Helper
}

func NewApplicationRepo(data *Data, logger log.Logger) biz.ApplicationRepo {
	l := log.NewHelper(log.With(logger, "module", "app/repo/core-service"))
	return &ApplicationRepo{
		data: data,
		log:  l,
	}
}

func (r *ApplicationRepo) convertEntToProto(in *ent.Application) *v1.Application {
	if in == nil {
		return nil
	}
	return &v1.Application{
		Id:         in.ID,
		Name:       in.Name,
		AppId:      in.AppID,
		AppKey:     in.AppKey,
		Remark:     in.Remark,
		CreatorId:  in.CreatorID,
		CreateTime: util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime: util.UnixMilliToStringPtr(in.UpdateTime),
		DeleteTime: util.UnixMilliToStringPtr(in.DeleteTime),
	}
}

func (r *ApplicationRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListApplicationResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder1 := r.data.db.Application.Query()

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

	builder2 := r.data.db.Application.Query()
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

	items := make([]*v1.Application, 0, len(users))
	for _, u := range users {
		item := r.convertEntToProto(u)
		items = append(items, item)
	}

	return &v1.ListApplicationResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *ApplicationRepo) Get(ctx context.Context, req *v1.GetApplicationRequest) (*v1.Application, error) {
	po, err := r.data.db.Application.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *ApplicationRepo) Create(ctx context.Context, req *v1.CreateApplicationRequest) (*v1.Application, error) {
	po, err := r.data.db.Application.Create().
		SetNillableName(req.App.Name).
		SetNillableAppID(req.App.AppId).
		SetNillableAppKey(req.App.AppKey).
		SetNillableStatus(req.App.Status).
		SetNillableCreatorID(req.App.CreatorId).
		SetNillableRemark(req.App.Remark).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *ApplicationRepo) Update(ctx context.Context, req *v1.UpdateApplicationRequest) (*v1.Application, error) {
	builder := r.data.db.Application.UpdateOneID(req.Id).
		SetNillableName(req.App.Name).
		SetNillableAppID(req.App.AppId).
		SetNillableAppKey(req.App.AppKey).
		SetNillableStatus(req.App.Status).
		SetNillableRemark(req.App.Remark).
		SetUpdateTime(time.Now().UnixMilli())

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *ApplicationRepo) Delete(ctx context.Context, req *v1.DeleteApplicationRequest) (bool, error) {
	err := r.data.db.Application.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
