package service

import (
	"context"
	"kratos-bi/app/logger/service/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"

	v1 "kratos-bi/gen/api/go/report/service/v1"
)

type SaverService struct {
	log          *log.Helper
	realtimeRepo *data.RealtimeWarehousingRepo
}

func NewSaverService(
	logger log.Logger,
	realtimeRepo *data.RealtimeWarehousingRepo,
) *SaverService {
	l := log.NewHelper(log.With(logger, "module", "saver/service/logger-service"))
	return &SaverService{
		log:          l,
		realtimeRepo: realtimeRepo,
	}
}

func (s *SaverService) SaveUserReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.UserReport) error {
	return nil
}

func (s *SaverService) SaveEventReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.RealTimeWarehousingData) error {
	return nil
}
