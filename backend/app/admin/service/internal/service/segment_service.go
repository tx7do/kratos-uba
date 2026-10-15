package service

import (
	"context"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

type SegmentService struct {
	adminV1.SegmentServiceHTTPServer

	log *log.Helper

	segmentServiceClient ubaV1.SegmentServiceClient
}

func NewSegmentService(
	ctx *bootstrap.Context,
	segmentServiceClient ubaV1.SegmentServiceClient,
) *SegmentService {
	return &SegmentService{
		log:                  ctx.NewLoggerHelper("segment/service/admin-service"),
		segmentServiceClient: segmentServiceClient,
	}
}

func (s *SegmentService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListSegmentResponse, error) {
	return s.segmentServiceClient.List(ctx, req)
}

func (s *SegmentService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountSegmentResponse, error) {
	return s.segmentServiceClient.Count(ctx, req)
}

func (s *SegmentService) Get(ctx context.Context, req *ubaV1.GetSegmentRequest) (*ubaV1.Segment, error) {
	return s.segmentServiceClient.Get(ctx, req)
}

func (s *SegmentService) Create(ctx context.Context, req *ubaV1.CreateSegmentRequest) (*ubaV1.Segment, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.segmentServiceClient.Create(ctx, req)
}

func (s *SegmentService) Update(ctx context.Context, req *ubaV1.UpdateSegmentRequest) (*ubaV1.Segment, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.segmentServiceClient.Update(ctx, req)
}

func (s *SegmentService) Delete(ctx context.Context, req *ubaV1.DeleteSegmentRequest) (*emptypb.Empty, error) {
	return s.segmentServiceClient.Delete(ctx, req)
}

func (s *SegmentService) Refresh(ctx context.Context, req *ubaV1.RefreshSegmentRequest) (*ubaV1.Segment, error) {
	return s.segmentServiceClient.Refresh(ctx, req)
}

func (s *SegmentService) ListUsers(ctx context.Context, req *ubaV1.ListSegmentUsersRequest) (*ubaV1.ListSegmentUsersResponse, error) {
	return s.segmentServiceClient.ListUsers(ctx, req)
}
