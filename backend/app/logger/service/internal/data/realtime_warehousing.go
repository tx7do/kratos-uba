package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-uba/gen/api/go/report/service/v1"
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

func (r *RealtimeWarehousingRepo) Create(req *v1.RealTimeWarehousingData) error {
	query := "INSERT INTO realtime_warehousing (event_name, report_data) VALUES (?, ?)"
	_, err := r.data.db.ExecContext(context.Background(), query, req.GetEventName(), req.GetReportData())
	if err != nil {
		r.log.Error(err)
		return err
	}
	return nil
}
