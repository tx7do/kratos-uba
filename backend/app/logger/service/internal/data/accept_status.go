package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-uba/api/gen/go/report/service/v1"
)

type AcceptStatusRepo struct {
	data *Data
	log  *log.Helper
}

func NewAcceptStatusRepo(data *Data, logger log.Logger) *AcceptStatusRepo {
	l := log.NewHelper(log.With(logger, "module", "accept-status/repo/logger-service"))
	return &AcceptStatusRepo{
		data: data,
		log:  l,
	}
}

func (r *AcceptStatusRepo) Create(req *v1.AcceptStatusReportData) error {
	query := "INSERT INTO acceptance_status (data_name, report_type, report_data, status, error_reason, error_handling, part_event, part_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.data.db.ExecContext(context.Background(), query,
		req.GetDataName(), req.GetReportType(), req.GetReportData(), req.GetStatus(), req.GetErrorReason(), req.GetErrorHandling(), req.GetPartEvent(), req.GetPartDate())
	if err != nil {
		r.log.Error(err)
		return err
	}

	return nil
}
