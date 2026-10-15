// Package netutil 提供绑定 kratos transport 的网络辅助函数。
//
// 通用核心（SSRF 防护、可信代理链的客户端真实 IP 解析）已下沉到
// github.com/tx7do/go-utils/netutil，本包仅保留依赖 kratos transport
// 上下文的便捷入口，内部委托给通用实现。
package netutil

import (
	"context"
	"net/http"

	gonetutil "github.com/tx7do/go-utils/netutil"

	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// ClientIPFromContext 从 kratos context 中提取客户端真实 IP。
// service 层通过此函数即可取得 IP，无需直接依赖中间件包。
// 若上下文非 HTTP 传输或取不到 request，返回空串。
func ClientIPFromContext(ctx context.Context) string {
	req := requestFromContext(ctx)
	if req == nil {
		return ""
	}
	return gonetutil.ClientIPFromRequest(req)
}

// HeaderFromContext 从 kratos context 中提取 HTTP 请求头。
// 用于 service 层读取通过 header 传递的参数（如验证码 id/value）。
// 若上下文非 HTTP 传输或取不到 request，返回 nil。
func HeaderFromContext(ctx context.Context) http.Header {
	req := requestFromContext(ctx)
	if req == nil {
		return nil
	}
	return req.Header
}

// CookieFromContext 从 kratos context 中提取具名 HTTP cookie 值。
// 用于 refresh token 接口从 HttpOnly cookie 中读取刷新令牌。
// 若上下文非 HTTP 传输、取不到 request 或无此 cookie，返回空串。
func CookieFromContext(ctx context.Context, name string) string {
	req := requestFromContext(ctx)
	if req == nil {
		return ""
	}
	c, err := req.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

// requestFromContext 从 kratos server context 中取出 *http.Request。
func requestFromContext(ctx context.Context) *http.Request {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil
	}
	htr, ok := tr.(khttp.Transporter)
	if !ok {
		return nil
	}
	return htr.Request()
}
