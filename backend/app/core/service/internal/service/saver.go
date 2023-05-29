package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"

	v1 "kratos-bi/gen/api/go/report/service/v1"
)

type SaverService struct {
	log *log.Helper
}

func NewSaverService(logger log.Logger) *ApplicationService {
	l := log.NewHelper(log.With(logger, "module", "saver/service/core-service"))
	return &ApplicationService{
		log: l,
	}
}

func (s *SaverService) SaveUserReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.UserReport) error {
	return nil
}

func (s *SaverService) SaveEventReport(ctx context.Context, topic string, headers broker.Headers, msg *v1.EventReport) error {
	return nil
}
