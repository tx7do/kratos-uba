package service

import (
	"github.com/hibiken/asynq"

	notificationV1 "go-wind-uba/api/gen/go/notification/service/v1"
)

// TaskEnqueuer 异步派发能力（与站内信广播同模式）：asynq 未配置时为 nil，
// 全部通知走同步投递。
type TaskEnqueuer interface {
	NewTask(typeName string, msg any, opts ...asynq.Option) error
}

// defaultNotificationRules 初始事件→渠道路由（空表才播，删行不复活）：
// 站内信走台账、监控告警走站内信、风险告警走 Webhook。
var defaultNotificationRules = []struct {
	EventType notificationV1.EventType
	Channel   notificationV1.Channel
	IsAsync   bool
}{
	{
		EventType: notificationV1.EventType_INTERNAL_MESSAGE,
		Channel:   notificationV1.Channel_INTERNAL,
		IsAsync:   false,
	},
	{
		EventType: notificationV1.EventType_MONITOR_ALERT,
		Channel:   notificationV1.Channel_INTERNAL,
		IsAsync:   false,
	},
	{
		EventType: notificationV1.EventType_RISK_ALERT,
		Channel:   notificationV1.Channel_WEBHOOK,
		IsAsync:   true,
	},
	{
		EventType: notificationV1.EventType_REPORT_GENERATED,
		Channel:   notificationV1.Channel_INTERNAL,
		IsAsync:   false,
	},
}
