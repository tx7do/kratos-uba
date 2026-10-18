// Package ai 是 AI 模型客户端工厂：把 sys_ai_providers 表行提炼成可调用的
// OpenAI 兼容客户端（云 = OpenAI 兼容 API，如 DeepSeek/通义；本地 = Ollama）。
//
// 这里只做"表行 → 客户端"的翻译与超时语义修正，不含任何业务逻辑；
// provider 的归属校验、密钥解密在 service 层完成后再进来。
package ai

import (
	"fmt"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

// ModelType 模型部署形态，取值与 proto ai.service.v1.AiProvider.ModelType 一致
// （1=LOCAL，2=CLOUD）。
type ModelType int32

const (
	ModelTypeLocal ModelType = 1
	ModelTypeCloud ModelType = 2
)

// ClientConfig 从 AiProvider 表行提炼的连接配置。
type ClientConfig struct {
	ModelType      int32
	ModelName      string
	BaseUrl        string
	Organization   string
	LocalHost      string
	LocalPort      int32
	TimeoutSeconds int32
}

// NewClient 构造 OpenAI 兼容客户端。
//
// 云端模型必须传解密后的 apiKey（空 key 构造出的客户端调用必 401）；
// 本地 Ollama 无鉴权，key 传空即可（占位 "none"，Ollama 忽略 Authorization）。
//
// 显式注入无超时的 http.Client 覆盖默认 30s Timeout——超时挂在 http.Client 上
// 对 CreateChatCompletionStream 意味着"整条流 30 秒必须收完"，长回复会被腰斩。
// 单轮流式对话的时长由调用方以 context deadline 控制；非流式调用（NL→SQL、
// 结论生成、embedding）走各自 context 超时。
func NewClient(cfg *ClientConfig, decryptedApiKey string) (*openai.Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("ai client config is nil")
	}

	// 本地 Ollama 无鉴权，空 key 占位 "none"（Ollama 忽略 Authorization；
	// go-openai 的 token 只能在 DefaultConfig 传入，无导出字段可后置覆盖）。
	authKey := decryptedApiKey
	if authKey == "" {
		authKey = "none"
	}
	cc := openai.DefaultConfig(authKey)

	switch ModelType(cfg.ModelType) {
	case ModelTypeCloud:
		if cfg.BaseUrl != "" {
			cc.BaseURL = cfg.BaseUrl
		}
		cc.OrgID = cfg.Organization
	case ModelTypeLocal:
		host := cfg.LocalHost
		if host == "" {
			host = "localhost"
		}
		port := cfg.LocalPort
		if port == 0 {
			port = 11434
		}
		cc.BaseURL = fmt.Sprintf("http://%s:%d/v1", host, port)
	default:
		return nil, fmt.Errorf("unsupported ai model type: %d", cfg.ModelType)
	}

	cc.HTTPClient = &http.Client{}

	return openai.NewClientWithConfig(cc), nil
}
