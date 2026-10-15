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

type CampaignService struct {
	adminV1.CampaignServiceHTTPServer

	log *log.Helper

	campaignServiceClient ubaV1.CampaignServiceClient
}

func NewCampaignService(
	ctx *bootstrap.Context,
	campaignServiceClient ubaV1.CampaignServiceClient,
) *CampaignService {
	return &CampaignService{
		log:                   ctx.NewLoggerHelper("campaign/service/admin-service"),
		campaignServiceClient: campaignServiceClient,
	}
}

func (s *CampaignService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListCampaignResponse, error) {
	return s.campaignServiceClient.List(ctx, req)
}

func (s *CampaignService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountCampaignResponse, error) {
	return s.campaignServiceClient.Count(ctx, req)
}

func (s *CampaignService) Get(ctx context.Context, req *ubaV1.GetCampaignRequest) (*ubaV1.Campaign, error) {
	return s.campaignServiceClient.Get(ctx, req)
}

func (s *CampaignService) Create(ctx context.Context, req *ubaV1.CreateCampaignRequest) (*ubaV1.Campaign, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorCampaignBadRequest("invalid parameter")
	}

	return s.campaignServiceClient.Create(ctx, req)
}

func (s *CampaignService) Update(ctx context.Context, req *ubaV1.UpdateCampaignRequest) (*ubaV1.Campaign, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorCampaignBadRequest("invalid parameter")
	}

	return s.campaignServiceClient.Update(ctx, req)
}

func (s *CampaignService) Delete(ctx context.Context, req *ubaV1.DeleteCampaignRequest) (*emptypb.Empty, error) {
	return s.campaignServiceClient.Delete(ctx, req)
}

func (s *CampaignService) Send(ctx context.Context, req *ubaV1.SendCampaignRequest) (*ubaV1.Campaign, error) {
	return s.campaignServiceClient.Send(ctx, req)
}

func (s *CampaignService) Cancel(ctx context.Context, req *ubaV1.CancelCampaignRequest) (*ubaV1.Campaign, error) {
	return s.campaignServiceClient.Cancel(ctx, req)
}
