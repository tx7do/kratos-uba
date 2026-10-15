package service

import (
	"context"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

type SqlQueryService struct {
	adminV1.SqlQueryServiceHTTPServer

	log *log.Helper

	sqlQueryServiceClient ubaV1.SqlQueryServiceClient
}

func NewSqlQueryService(
	ctx *bootstrap.Context,
	sqlQueryServiceClient ubaV1.SqlQueryServiceClient,
) *SqlQueryService {
	return &SqlQueryService{
		log:                   ctx.NewLoggerHelper("sql-query/service/admin-service"),
		sqlQueryServiceClient: sqlQueryServiceClient,
	}
}

// ExecuteQuery 转发只读查询到 core。
func (s *SqlQueryService) ExecuteQuery(ctx context.Context, req *ubaV1.ExecuteSqlQueryRequest) (*ubaV1.ExecuteSqlQueryResponse, error) {
	if req == nil || req.GetQuery() == "" {
		return nil, ubaV1.ErrorBadRequest("query is required")
	}

	return s.sqlQueryServiceClient.ExecuteQuery(ctx, req)
}
