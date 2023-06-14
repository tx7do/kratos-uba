package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-bi/app/logger/service/internal/data/ent"
	v1 "kratos-bi/gen/api/go/report/service/v1"
	util "kratos-bi/pkg/util/time"
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

func (r *AcceptStatusRepo) convertEntToProto(in *ent.AcceptanceStatus) *v1.AcceptStatusReportData {
	if in == nil {
		return nil
	}
	return &v1.AcceptStatusReportData{
		Id:            &in.ID,
		DataName:      in.DataName,
		ReportType:    in.ReportType,
		ReportData:    in.ReportData,
		ErrorReason:   in.ErrorReason,
		ErrorHandling: in.ErrorHandling,
		Status:        in.Status,
		PartDate:      util.TimeToTimeString(in.PartDate),
	}
}

func (r *AcceptStatusRepo) Create(ctx context.Context, req *v1.AcceptStatusReportData) (*v1.AcceptStatusReportData, error) {
	resp, err := r.data.db.AcceptanceStatus.Create().
		SetID(req.GetId()).
		SetNillableDataName(req.DataName).
		SetNillableReportData(req.ReportData).
		SetNillableErrorReason(req.ErrorReason).
		SetNillableErrorHandling(req.ErrorHandling).
		SetNillableStatus(req.Status).
		SetNillablePartDate(util.StringDateToTime(req.PartDate)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(resp), err
}
