package server

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"github.com/tx7do/kratos-transport/transport/asynq"

	"go-wind-uba/app/core/service/internal/service"

	appViewer "go-wind-uba/pkg/entgo/viewer"
	"go-wind-uba/pkg/task"
)

// NewAsynqServer creates a new asynq server.
func NewAsynqServer(
	ctx *bootstrap.Context,
	taskService *service.TaskService,
	notificationService *service.NotificationService,
	campaignService *service.CampaignService,
	replayService *service.ReplayService,
) *asynq.Server {
	cfg := ctx.GetConfig()

	if cfg == nil || cfg.Server == nil || cfg.Server.Asynq == nil {
		return nil
	}

	srv := asynq.NewServer(
		asynq.WithCodec(cfg.Server.Asynq.GetCodec()),
		asynq.WithRedisURI(cfg.Server.Asynq.GetUri()),
		asynq.WithLocation(cfg.Server.Asynq.GetLocation()),
		asynq.WithGracefullyShutdown(cfg.Server.Asynq.GetEnableGracefullyShutdown()),
		asynq.WithShutdownTimeout(cfg.Server.Asynq.GetShutdownTimeout().AsDuration()),
	)

	taskService.RegisterTaskScheduler(srv)

	var err error

	// 通知异步派发：入队能力 + handler（台账状态把关幂等，超时经 ctx 掐断 SMTP 拨号）
	notificationService.RegisterTaskEnqueuer(taskService)
	if err := asynq.RegisterSubscriberWithCtx(srv, task.NotificationDispatchTaskType, notificationService.AsyncNotificationDispatch); err != nil {
		log.Error(err)
	}

	// 通知台账超时清扫 handler（周期调度见下方 NewPeriodicTask）
	if err := asynq.RegisterSubscriber(srv, task.NotificationDeliverySweepTaskType, notificationService.AsyncDeliverySweep); err != nil {
		log.Error(err)
	}

	// 营销触达群发 handler（asynq ProcessAt 定时任务或 Send RPC 入队的即时任务）
	if err := asynq.RegisterSubscriber(srv, task.CampaignSendTaskType, campaignService.AsyncCampaignSend); err != nil {
		log.Error(err)
	}

	// 会话回放保留期清扫（每天 03:30）+ 周期调度
	if err := asynq.RegisterSubscriber(srv, task.ReplayRetentionSweepTaskType, replayService.AsyncReplayRetentionSweep); err != nil {
		log.Error(err)
	}
	if _, err = srv.NewPeriodicTask(task.ReplayRetentionSweepCronSpec, task.ReplayRetentionSweepTaskType, &task.ReplayRetentionSweepTaskData{}); err != nil {
		log.Error(err)
	}

	// 注册任务
	if err = asynq.RegisterSubscriber(srv, task.BackupTaskType, taskService.AsyncBackup); err != nil {
		log.Error(err)
	}

	// 通知台账清扫的系统级周期调度（不写 sys_tasks 表；RestartAllTask 会移除再重建全部周期任务，
	// 故这里只挂 handler、调度挂在 StartAllTask 之后无法自愈——与上游不同，uba 的 RestartAllTask
	// 走 TaskService，需保证重启后恢复：把调度挂到 srv 上并随 StartAllTask 重放）。
	if _, err = srv.NewPeriodicTask(task.NotificationDeliverySweepCronSpec, task.NotificationDeliverySweepTaskType, &task.NotificationDeliverySweepTaskData{}); err != nil {
		log.Error(err)
	}

	// 启动所有的任务
	_, _ = taskService.StartAllTask(appViewer.NewSystemViewerContext(ctx.Context()), nil)

	return srv
}
