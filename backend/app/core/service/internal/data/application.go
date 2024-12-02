package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"entgo.io/ent/dialect/sql"

	entgo "github.com/tx7do/go-utils/entgo/query"
	timeUtil "github.com/tx7do/go-utils/timeutil"

	"kratos-uba/app/core/service/internal/data/ent"
	"kratos-uba/app/core/service/internal/data/ent/application"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-uba/api/gen/go/user/service/v1"
)

type ApplicationRepo struct {
	data *Data
	log  *log.Helper
}

func NewApplicationRepo(data *Data, logger log.Logger) *ApplicationRepo {
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
		OwnerId:    in.OwnerID,
		Status:     in.Status,
		KeepMonth:  in.KeepMonth,
		CreateTime: timeUtil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeUtil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeUtil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *ApplicationRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Application.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}
	return builder.Count(ctx)
}

func (r *ApplicationRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListApplicationResponse, error) {
	builder := r.data.db.Client().Application.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), application.FieldCreateTime,
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

	items := make([]*v1.Application, 0, len(results))
	for _, item := range results {
		items = append(items, r.convertEntToProto(item))
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListApplicationResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *ApplicationRepo) Get(ctx context.Context, req *v1.GetApplicationRequest) (*v1.Application, error) {
	resp, err := r.data.db.Client().Application.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(resp), err
}

func (r *ApplicationRepo) Create(ctx context.Context, req *v1.CreateApplicationRequest) (*v1.Application, error) {
	resp, err := r.data.db.Client().Application.Create().
		SetNillableName(req.App.Name).
		SetNillableAppID(req.App.AppId).
		SetNillableAppKey(req.App.AppKey).
		SetNillableStatus(req.App.Status).
		SetNillableCreatorID(req.App.CreatorId).
		SetNillableOwnerID(req.App.OwnerId).
		SetNillableRemark(req.App.Remark).
		SetNillableKeepMonth(req.App.KeepMonth).
		SetCreateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(resp), err
}

func (r *ApplicationRepo) Update(ctx context.Context, req *v1.UpdateApplicationRequest) (*v1.Application, error) {
	builder := r.data.db.Client().Application.UpdateOneID(req.Id).
		SetNillableName(req.App.Name).
		SetNillableAppID(req.App.AppId).
		SetNillableAppKey(req.App.AppKey).
		SetNillableStatus(req.App.Status).
		SetNillableOwnerID(req.App.OwnerId).
		SetNillableRemark(req.App.Remark).
		SetNillableKeepMonth(req.App.KeepMonth).
		SetUpdateTime(time.Now())

	resp, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(resp), err
}

func (r *ApplicationRepo) Delete(ctx context.Context, req *v1.DeleteApplicationRequest) (bool, error) {
	err := r.data.db.Client().Application.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
