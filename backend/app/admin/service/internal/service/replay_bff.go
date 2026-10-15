package service

import (
	"context"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/tx7do/kratos-bootstrap/bootstrap"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

type ReplayService struct {
	adminV1.ReplayServiceHTTPServer

	log *log.Helper

	replayServiceClient ubaV1.ReplayServiceClient
}

func NewReplayService(
	ctx *bootstrap.Context,
	replayServiceClient ubaV1.ReplayServiceClient,
) *ReplayService {
	return &ReplayService{
		log:                 ctx.NewLoggerHelper("replay/service/admin-service"),
		replayServiceClient: replayServiceClient,
	}
}

func (s *ReplayService) ListSessions(ctx context.Context, req *ubaV1.ListReplaySessionsRequest) (*ubaV1.ListReplaySessionsResponse, error) {
	return s.replayServiceClient.ListSessions(ctx, req)
}

func (s *ReplayService) GetSessionEvents(ctx context.Context, req *ubaV1.GetReplaySessionRequest) (*ubaV1.GetReplaySessionResponse, error) {
	return s.replayServiceClient.GetSessionEvents(ctx, req)
}

func (s *ReplayService) DeleteSession(ctx context.Context, req *ubaV1.DeleteReplaySessionRequest) (*emptypb.Empty, error) {
	return s.replayServiceClient.DeleteSession(ctx, req)
}
