// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             (unknown)
// source: agent/service/v1/i_application.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-bi/gen/api/go/common/pagination"
	v1 "kratos-bi/gen/api/go/user/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationApplicationServiceCreateApplication = "/agent.service.v1.ApplicationService/CreateApplication"
const OperationApplicationServiceDeleteApplication = "/agent.service.v1.ApplicationService/DeleteApplication"
const OperationApplicationServiceGetApplication = "/agent.service.v1.ApplicationService/GetApplication"
const OperationApplicationServiceListApplication = "/agent.service.v1.ApplicationService/ListApplication"
const OperationApplicationServiceUpdateApplication = "/agent.service.v1.ApplicationService/UpdateApplication"

type ApplicationServiceHTTPServer interface {
	// CreateApplication 创建应用
	CreateApplication(context.Context, *v1.CreateApplicationRequest) (*v1.Application, error)
	// DeleteApplication 删除应用
	DeleteApplication(context.Context, *v1.DeleteApplicationRequest) (*emptypb.Empty, error)
	// GetApplication 获取应用数据
	GetApplication(context.Context, *v1.GetApplicationRequest) (*v1.Application, error)
	// ListApplication 获取应用列表
	ListApplication(context.Context, *pagination.PagingRequest) (*v1.ListApplicationResponse, error)
	// UpdateApplication 更新应用
	UpdateApplication(context.Context, *v1.UpdateApplicationRequest) (*v1.Application, error)
}

func RegisterApplicationServiceHTTPServer(s *http.Server, srv ApplicationServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/agent/v1/apps", _ApplicationService_ListApplication0_HTTP_Handler(srv))
	r.GET("/agent/v1/apps/{id}", _ApplicationService_GetApplication0_HTTP_Handler(srv))
	r.POST("/agent/v1/apps", _ApplicationService_CreateApplication0_HTTP_Handler(srv))
	r.PUT("/agent/v1/apps/{id}", _ApplicationService_UpdateApplication0_HTTP_Handler(srv))
	r.DELETE("/agent/v1/apps/{id}", _ApplicationService_DeleteApplication0_HTTP_Handler(srv))
}

func _ApplicationService_ListApplication0_HTTP_Handler(srv ApplicationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplicationServiceListApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListApplication(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListApplicationResponse)
		return ctx.Result(200, reply)
	}
}

func _ApplicationService_GetApplication0_HTTP_Handler(srv ApplicationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetApplicationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplicationServiceGetApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApplication(ctx, req.(*v1.GetApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Application)
		return ctx.Result(200, reply)
	}
}

func _ApplicationService_CreateApplication0_HTTP_Handler(srv ApplicationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateApplicationRequest
		if err := ctx.Bind(&in.App); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplicationServiceCreateApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateApplication(ctx, req.(*v1.CreateApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Application)
		return ctx.Result(200, reply)
	}
}

func _ApplicationService_UpdateApplication0_HTTP_Handler(srv ApplicationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateApplicationRequest
		if err := ctx.Bind(&in.App); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplicationServiceUpdateApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApplication(ctx, req.(*v1.UpdateApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Application)
		return ctx.Result(200, reply)
	}
}

func _ApplicationService_DeleteApplication0_HTTP_Handler(srv ApplicationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteApplicationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplicationServiceDeleteApplication)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteApplication(ctx, req.(*v1.DeleteApplicationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type ApplicationServiceHTTPClient interface {
	CreateApplication(ctx context.Context, req *v1.CreateApplicationRequest, opts ...http.CallOption) (rsp *v1.Application, err error)
	DeleteApplication(ctx context.Context, req *v1.DeleteApplicationRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetApplication(ctx context.Context, req *v1.GetApplicationRequest, opts ...http.CallOption) (rsp *v1.Application, err error)
	ListApplication(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListApplicationResponse, err error)
	UpdateApplication(ctx context.Context, req *v1.UpdateApplicationRequest, opts ...http.CallOption) (rsp *v1.Application, err error)
}

type ApplicationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewApplicationServiceHTTPClient(client *http.Client) ApplicationServiceHTTPClient {
	return &ApplicationServiceHTTPClientImpl{client}
}

func (c *ApplicationServiceHTTPClientImpl) CreateApplication(ctx context.Context, in *v1.CreateApplicationRequest, opts ...http.CallOption) (*v1.Application, error) {
	var out v1.Application
	pattern := "/agent/v1/apps"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApplicationServiceCreateApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.App, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApplicationServiceHTTPClientImpl) DeleteApplication(ctx context.Context, in *v1.DeleteApplicationRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/agent/v1/apps/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApplicationServiceDeleteApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApplicationServiceHTTPClientImpl) GetApplication(ctx context.Context, in *v1.GetApplicationRequest, opts ...http.CallOption) (*v1.Application, error) {
	var out v1.Application
	pattern := "/agent/v1/apps/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApplicationServiceGetApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApplicationServiceHTTPClientImpl) ListApplication(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListApplicationResponse, error) {
	var out v1.ListApplicationResponse
	pattern := "/agent/v1/apps"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApplicationServiceListApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApplicationServiceHTTPClientImpl) UpdateApplication(ctx context.Context, in *v1.UpdateApplicationRequest, opts ...http.CallOption) (*v1.Application, error) {
	var out v1.Application
	pattern := "/agent/v1/apps/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApplicationServiceUpdateApplication))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.App, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
