package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"

	util "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	v1 "kratos-uba/api/gen/go/agent/service/v1"
	reportV1 "kratos-uba/api/gen/go/report/service/v1"

	"kratos-uba/pkg/topic"
)

type ReportService struct {
	v1.ReportServiceHTTPServer

	kafkaBroker broker.Broker
	log         *log.Helper
}

func NewReportService(logger log.Logger, kafkaBroker broker.Broker) *ReportService {
	l := log.NewHelper(log.With(logger, "module", "report/service/agent-service"))
	return &ReportService{
		log:         l,
		kafkaBroker: kafkaBroker,
	}
}

func (s *ReportService) PostReport(ctx context.Context, req *v1.PostReportRequest) (*v1.PostReportResponse, error) {

	_ = s.kafkaBroker.Publish(ctx, topic.EventReportData, reportV1.RealTimeWarehousingData{
		EventName:  &req.EventName,
		ReportData: &req.Content,
		CreateTime: util.UnixMilliToStringPtr(trans.Ptr(time.Now().UnixMilli())),
	})
	return &v1.PostReportResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
