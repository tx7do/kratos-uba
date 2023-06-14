package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-bi/app/logger/service/internal/data/ent"
	"time"

	v1 "kratos-bi/gen/api/go/report/service/v1"

	util "kratos-bi/pkg/util/time"
)

type RealtimeWarehousingRepo struct {
	data *Data
	log  *log.Helper
}

func NewRealtimeWarehousingRepo(data *Data, logger log.Logger) *RealtimeWarehousingRepo {
	l := log.NewHelper(log.With(logger, "module", "realtime-warehousing/repo/logger-service"))
	return &RealtimeWarehousingRepo{
		data: data,
		log:  l,
	}
}

func (r *RealtimeWarehousingRepo) convertEntToProto(in *ent.RealtimeWarehousing) *v1.RealTimeWarehousingData {
	if in == nil {
		return nil
	}
	return &v1.RealTimeWarehousingData{
		Id:         &in.ID,
		EventName:  in.EventName,
		ReportData: in.ReportData,
		CreateTime: util.TimeToTimeString(in.CreateTime),
	}
}

func (r *RealtimeWarehousingRepo) Create(ctx context.Context, req *v1.RealTimeWarehousingData) (*v1.RealTimeWarehousingData, error) {
	resp, err := r.data.db.RealtimeWarehousing.Create().
		SetID(req.GetId()).
		SetNillableEventName(req.EventName).
		SetNillableReportData(req.ReportData).
		SetCreateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(resp), err
}
