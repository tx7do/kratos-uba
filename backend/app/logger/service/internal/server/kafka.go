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

func UserReportCreator() broker.Any  { return &v1.UserReport{} }
func EventReportCreator() broker.Any { return &v1.EventReport{} }

type UserReportHandler func(_ context.Context, topic string, headers broker.Headers, msg *v1.UserReport) error
type EventReportHandler func(_ context.Context, topic string, headers broker.Headers, msg *v1.EventReport) error

func RegisterUserReportHandler(fnc UserReportHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *v1.UserReport:
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
		case *v1.EventReport:
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

func registerKafkaSubscribers(_ context.Context, srv *kafka.Server, svc *service.SaverService) {
	_, _ = srv.Subscribe(topic.EventReportData,
		RegisterUserReportHandler(svc.SaveUserReport),
		UserReportCreator,
		broker.WithQueueName(topic.LoggerSaverQueue),
	)

	_, _ = srv.Subscribe(topic.UserReportData,
		RegisterEventReportHandler(svc.SaveEventReport),
		EventReportCreator,
		broker.WithQueueName(topic.LoggerSaverQueue),
	)
}
