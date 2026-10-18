package ent

import (
	"context"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/tx7do/go-crud/viewer"
	"go.opentelemetry.io/otel/trace"

	authenticationV1 "go-wind-uba/api/gen/go/authentication/service/v1"

	appViewer "go-wind-uba/pkg/entgo/viewer"
	"go-wind-uba/pkg/metadata"
)

// Server 设置 Ent Viewer 到 Context 中的中间件
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			md, err := metadata.FromServerContext(ctx)
			if err != nil {
				reqType := "<nil>"
				if req != nil {
					reqType = reflect.TypeOf(req).String()
				}
				t, _ := transport.FromServerContext(ctx)
				kind, operation, endpoint := "", "", ""
				if t != nil {
					kind = string(t.Kind())
					operation = t.Operation()
					endpoint = t.Endpoint()
				}
				log.Errorf("ent middleware: failed to get metadata from context: %v; req_type=%s; transport=%s; operation=%s; endpoint=%s",
					err, reqType, kind, operation, endpoint)
			}
			if md == nil {
				// 内部服务间调用（collector 应用鉴权/后台任务回调）不带操作者元数据：
				// 以系统上下文执行，绕过租户隐私层——core gRPC 端口仅内网可达，
				// 与 asynq_server 的 StartAllTask 同一信任模型。
				ctx = appViewer.NewSystemViewerContext(ctx)
				return handler(ctx, req)
			}

			var traceID string
			spanContext := trace.SpanContextFromContext(ctx)
			if spanContext.HasTraceID() {
				traceID = spanContext.TraceID().String()
			}

			ctx = viewer.WithContext(ctx, metaDataToUserViewerContext(md, traceID))
			// admin BFF 转发的 gRPC 调用可能缺少操作者元数据（md 存在但内容为空），
			// 此时 metaDataToUserViewerContext 返回 nil → 隐私层拒绝。
			// 回落系统上下文（core gRPC 端口仅内网可达）。
			if _, ok := viewer.FromContext(ctx); !ok {
				ctx = appViewer.NewSystemViewerContext(ctx)
			}

			return handler(ctx, req)
		}
	}
}

func metaDataToUserViewerContext(md *authenticationV1.OperatorMetadata, traceID string) viewer.Context {
	if md == nil {
		return nil
	}

	userViewer := appViewer.NewUserViewer(
		md.GetUserId(),
		md.GetTenantId(),
		md.GetOrgUnitId(),
		traceID,
		md.GetDataScope(),
	)
	return userViewer
}

// InjectViewerForStream 把已认证请求的 OperatorMetadata 解析为 UserViewer 注入 ctx，
// 匿名/异常时回落 SystemViewer。供 kratos server-streaming RPC 使用——
// unary 中间件（Server）不覆盖流式 handler，流式入口需自行调用本函数。
func InjectViewerForStream(ctx context.Context) context.Context {
	var traceID string
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		traceID = spanContext.TraceID().String()
	}

	md, mdErr := metadata.FromServerContext(ctx)
	if mdErr != nil || md == nil {
		return appViewer.NewSystemViewerContext(ctx)
	}
	vc := metaDataToUserViewerContext(md, traceID)
	if vc == nil {
		return appViewer.NewSystemViewerContext(ctx)
	}
	return viewer.WithContext(ctx, vc)
}
