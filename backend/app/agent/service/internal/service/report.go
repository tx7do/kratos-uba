package service

import (
	"context"
	util "kratos-bi/pkg/util/time"
	"kratos-bi/pkg/util/trans"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"

	v1 "kratos-bi/gen/api/go/agent/service/v1"
	reportV1 "kratos-bi/gen/api/go/report/service/v1"

	"kratos-bi/pkg/topic"
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

func (s *ReportService) PostReport(_ context.Context, req *v1.PostReportRequest) (*v1.PostReportResponse, error) {

	_ = s.kafkaBroker.Publish(topic.EventReportData, reportV1.RealTimeWarehousingData{
		EventName:  &req.EventName,
		ReportData: &req.Content,
		CreateTime: util.UnixMilliToStringPtr(trans.Int64(time.Now().UnixMilli())),
	})
	return &v1.PostReportResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
