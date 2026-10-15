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

type ImportService struct {
	adminV1.ImportServiceHTTPServer

	log *log.Helper

	importServiceClient ubaV1.ImportServiceClient
}

func NewImportService(
	ctx *bootstrap.Context,
	importServiceClient ubaV1.ImportServiceClient,
) *ImportService {
	return &ImportService{
		log:                 ctx.NewLoggerHelper("import/service/admin-service"),
		importServiceClient: importServiceClient,
	}
}

func (s *ImportService) Create(ctx context.Context, req *ubaV1.CreateImportJobRequest) (*ubaV1.ImportJob, error) {
	if req == nil || req.GetEventsJson() == "" {
		return nil, ubaV1.ErrorImportBadRequest("events_json is required")
	}
	return s.importServiceClient.Create(ctx, req)
}

func (s *ImportService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListImportJobResponse, error) {
	return s.importServiceClient.List(ctx, req)
}

func (s *ImportService) Get(ctx context.Context, req *ubaV1.GetImportJobRequest) (*ubaV1.ImportJob, error) {
	return s.importServiceClient.Get(ctx, req)
}

func (s *ImportService) Cancel(ctx context.Context, req *ubaV1.CancelImportJobRequest) (*ubaV1.ImportJob, error) {
	return s.importServiceClient.Cancel(ctx, req)
}

func (s *ImportService) Delete(ctx context.Context, req *ubaV1.DeleteImportJobRequest) (*emptypb.Empty, error) {
	return s.importServiceClient.Delete(ctx, req)
}
