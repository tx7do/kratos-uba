package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/kafka"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"

	"kratos-uba/app/logger/service/internal/service"

	"kratos-uba/pkg/topic"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.SaverService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(cfg.Server.Kafka.Endpoints),
		kafka.WithCodec(cfg.Server.Kafka.Codec),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.SaverService) {
	_ = kafka.RegisterSubscriber(srv, ctx,
		topic.UserReportData, topic.LoggerSaverQueue, false,
		svc.SaveUserReport,
	)

	_ = kafka.RegisterSubscriber(srv, ctx,
		topic.EventReportData, topic.LoggerSaverQueue, false,
		svc.SaveEventReport,
	)
}
