package server

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-bi/app/logger/service/internal/service"
	"kratos-bi/gen/api/go/common/conf"
	v1 "kratos-bi/gen/api/go/report/service/v1"

	"kratos-bi/pkg/topic"
)

func UserReportCreator() broker.Any  { return &v1.AcceptStatusReportData{} }
func EventReportCreator() broker.Any { return &v1.RealTimeWarehousingData{} }

type UserReportHandler func(_ context.Context, topic string, headers broker.Headers, msg *v1.AcceptStatusReportData) error
type EventReportHandler func(_ context.Context, topic string, headers broker.Headers, msg *v1.RealTimeWarehousingData) error

func RegisterUserReportHandler(fnc UserReportHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *v1.AcceptStatusReportData:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func RegisterEventReportHandler(fnc EventReportHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *v1.RealTimeWarehousingData:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

// NewKafkaServer create a kafka server.
func NewKafkaServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.SaverService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(cfg.Server.Kafka.Addrs),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.SaverService) {
	_ = srv.RegisterSubscriber(ctx,
		topic.UserReportData, topic.LoggerSaverQueue, false,
		RegisterUserReportHandler(svc.SaveUserReport),
		UserReportCreator,
	)

	_ = srv.RegisterSubscriber(ctx,
		topic.EventReportData, topic.LoggerSaverQueue, false,
		RegisterEventReportHandler(svc.SaveEventReport),
		EventReportCreator,
	)
}
