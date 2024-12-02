package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-uba/api/gen/go/report/service/v1"
)

type ReportService struct {
	v1.UnimplementedReportServiceServer
	log *log.Helper
}

func NewReportService(logger log.Logger) *ReportService {
	l := log.NewHelper(log.With(logger, "module", "report/service/report-service"))
	return &ReportService{
		log: l,
	}
}

func (s *ReportService) ListReport(ctx context.Context, req *pagination.PagingRequest) (*v1.ListReportResponse, error) {
	return nil, nil
}

func (s *ReportService) GetReport(ctx context.Context, req *v1.GetReportRequest) (*v1.Report, error) {
	return nil, nil
}

func (s *ReportService) CreateReport(ctx context.Context, req *v1.CreateReportRequest) (*v1.Report, error) {
	return nil, nil
}

func (s *ReportService) UpdateReport(ctx context.Context, req *v1.UpdateReportRequest) (*v1.Report, error) {
	return nil, nil
}

func (s *ReportService) DeleteReport(ctx context.Context, req *v1.DeleteReportRequest) (*emptypb.Empty, error) {
	return nil, nil
}
