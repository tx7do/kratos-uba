package service

import (
	"context"
	"net"
	"net/http"
	"strings"

	kratosTransport "github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mileusna/useragent"
	"github.com/tx7do/go-utils/geoip/geolite"
	"github.com/tx7do/go-utils/trans"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

// eventGeoClient 复用审计日志的 GeoLite 客户端做 IP 地理兜底。
// 查询库缺失/未配置时 NewClient 仍返回可用实例（Query 返回错误），不影响主链路。
var eventGeoClient, _ = geolite.NewClient()

const (
	headerXForwardedFor = "X-Forwarded-For"
	headerXRealIP       = "X-Real-IP"
	headerUserAgent     = "User-Agent"
	headerReferer       = "Referer"
)

// enrichBehaviorFromRequest 服务端补全：真实 IP、UA 解析（os/browser/设备类型）、
// IP 地理兜底。SDK 只报了 UA/referer 字符串，解析与定位是 collector 的职责 ——
// 客户端不可信，国家/城市不能依赖它自报。
func enrichBehaviorFromRequest(ctx context.Context, behaviorEvent *ubaV1.BehaviorEvent, ci *ubaV1.ClientInfo) {
	req := requestFromContext(ctx)

	// ── 1. 真实 IP：优先 XFF 链首 → X-Real-IP → RemoteAddr ──
	ip := behaviorEvent.GetIp()
	if ip == "" && req != nil {
		ip = realClientIP(req)
		if ip != "" {
			behaviorEvent.Ip = trans.Ptr(trimAndLimit(ip, 64))
		}
	}

	// ── 2. UA 解析：os / browser / 设备类型 ──
	uaStr := ci.GetUserAgent()
	if uaStr == "" && req != nil {
		uaStr = req.Header.Get(headerUserAgent)
	}
	if uaStr != "" {
		ua := useragent.Parse(uaStr)

		if behaviorEvent.Os == nil || behaviorEvent.GetOs() == "" {
			if ua.OS != "" {
				behaviorEvent.Os = trans.Ptr(trimAndLimit(ua.OS, 32))
			}
		}

		// browser / 设备类型没有独立列，进 context 供分析筛选；
		// 只补缺，不覆盖 SDK 显式上报的同名属性。
		if behaviorEvent.Context == nil {
			behaviorEvent.Context = map[string]string{}
		}
		if _, ok := behaviorEvent.Context["browser"]; !ok && ua.Name != "" {
			browser := trimAndLimit(ua.Name+" "+ua.Version, 64)
			behaviorEvent.Context["browser"] = browser
		}
		if _, ok := behaviorEvent.Context["device_type"]; !ok {
			behaviorEvent.Context["device_type"] = deviceType(ua)
		}
	}

	// ── 3. referer 兜底：SDK 未带时从请求头取 ──
	if behaviorEvent.Referer == nil && req != nil {
		if ref := req.Header.Get(headerReferer); ref != "" {
			behaviorEvent.Referer = trans.Ptr(trimAndLimit(ref, 256))
		}
	}

	// ── 4. IP 地理兜底：客户端没报 country/city 时服务端查库 ──
	needGeo := behaviorEvent.Country == nil || behaviorEvent.GetCountry() == "" ||
		behaviorEvent.IpCity == nil || behaviorEvent.GetIpCity() == ""
	if needGeo && eventGeoClient != nil && isPublicIP(ip) {
		if res, err := eventGeoClient.Query(ip); err == nil {
			if (behaviorEvent.Country == nil || behaviorEvent.GetCountry() == "") && res.Country != "" {
				behaviorEvent.Country = trans.Ptr(trimAndLimit(res.Country, 64))
			}
			if (behaviorEvent.IpCity == nil || behaviorEvent.GetIpCity() == "") && res.City != "" {
				behaviorEvent.IpCity = trans.Ptr(trimAndLimit(res.City, 64))
			}
		}
	}
}

// requestFromContext 从 kratos 传输上下文取原始 HTTP 请求；非 HTTP 传输返回 nil。
func requestFromContext(ctx context.Context) *http.Request {
	if tr, ok := kratosTransport.FromServerContext(ctx); ok {
		if htr, ok := tr.(*khttp.Transport); ok {
			return htr.Request()
		}
	}
	return nil
}

// realClientIP 提取客户端真实 IP：XFF 链首（最外层代理应覆写该头防伪造）→
// X-Real-IP → RemoteAddr。与审计日志中间件的口径保持一致。
func realClientIP(request *http.Request) string {
	if request == nil {
		return ""
	}

	if xff := request.Header.Get(headerXForwardedFor); xff != "" {
		for _, ip := range strings.Split(xff, ",") {
			ip = strings.TrimSpace(ip)
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
	}

	if realIP := strings.TrimSpace(request.Header.Get(headerXRealIP)); realIP != "" {
		if net.ParseIP(realIP) != nil {
			return realIP
		}
	}

	host := request.RemoteAddr
	if host == "" {
		return ""
	}
	if trimmed, _, err := net.SplitHostPort(host); err == nil {
		host = trimmed
	}
	return host
}

// isPublicIP 判断是否值得做地理查询：回环/私网/解析失败一律跳过。
func isPublicIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return !ip.IsLoopback() && !ip.IsPrivate() && !ip.IsLinkLocalUnicast() && !ip.IsUnspecified()
}

// deviceType UA 归一化的设备类型。
func deviceType(ua useragent.UserAgent) string {
	switch {
	case ua.Bot:
		return "bot"
	case ua.Mobile:
		return "mobile"
	case ua.Tablet:
		return "tablet"
	case ua.Desktop:
		return "desktop"
	default:
		return "unknown"
	}
}
