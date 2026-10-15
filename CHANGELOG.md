# 更新日志

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/) 规范，
版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

> 历史条目基于 Git 提交记录回溯整理，仅记录对使用者有影响的功能性变更，
> 日常重构、代码风格、CI 调整等不在逐一列举。

## [1.0.0] - 2026-10-15

首个公开发布版本（含 2026-10-09 的初始平台基线）。本版本补全了从"分析引擎"到
"增长闭环平台"的全部功能差距，覆盖**分析 → 存档 → 触达 → 行动**完整链路，
共 17 个提交、634 个文件变更、新增 20 张后端表、22 个前端页面目录、
8 个后台引擎、5 个 SDK 新模块；后台一级菜单按用户旅程收拢为 7 组（最深三级）；
ClickHouse 引擎的读数正确性（事件去重 / FINAL 读 / 派生表按天回算）已与 Doris 对齐，
支持二选一部署。

### 新增

#### 增长闭环

- **持久化用户分群**：圈选定义保存（`uba_segments` / `uba_segment_users`），
  支持 include / exclude 条件与时间窗口，定时刷新（`refresh_interval_minutes`）
  与手动刷新（`POST /segments/{id}/refresh`）；漏斗与圈选页一键「保存为分群」；
  成员下钻分页查询；`refreshOnCreate` 创建即刷。分群定义以 JSON 快照
  持久化（`SegmentationDefinition` proto ↔ `map[string]any` 双向序列化），
  刷新走 Doris `Segmentation` OLAP 查询（`UNIQUE KEY` 幂等覆盖）。
- **自定义看板**：`uba_dashboards` 表 + Dashboard 契约，组件 =
  分析模型标识 + 请求参数 JSON 快照 + 栅格宽度；前端看板视图页按
  `a-row/a-col` 网格渲染，支持 event-trend 折线 / active-users 双线 /
  group-by 饼图 / funnel 漏斗 / retention 热力图五种渲染器。
- **指标阈值告警**：`uba_alert_rules` 表 + `AlertEvaluator` 后台评估器
  （每分钟 tick + Redis `SET NX` 分布式锁），支持事件量 / 活跃用户 /
  金额求和 / 金额均值四种聚合方式与 GT / GTE / LT / LTE 四种比较算子，
  命中后走站内信（发给规则创建者）与 Webhook 双通道；静默期
  （`silence_minutes`）防重复通知。
- **定时报告**：`uba_report_subscriptions` 表 + `ReportEngine` 执行器
  （每分钟 tick + Redis 锁，按 `interval_minutes` 到期判定），支持
  event-trend / active-users / group-by / funnel / retention 五种模型；
  结果摘要 + CSV（RFC4180 转义、上限 2000 行）经站内信 / Webhook 推送；
  参数缺时间范围时按执行间隔回看。
- **营销触达**：`uba_campaigns` 表 + `CampaignService`，绑定分群
  （`segment_id`）× 文案（`title` / `content` / `template_code`）× 站内信渠道，
  状态机 DRAFT → SENDING → SENT / PARTIAL / FAILED → CANCELLED；
  群发按分页遍历分群成员逐用户经通知域 SendDirect 投递（台账/脱敏/幂等继承
  通知域）；asynq ProcessAt 延迟任务实现定时发送，取消靠状态机跳过。
- **告警中心管理页**：阈值规则 CRUD（聚合方式 / 比较算子 / 阈值 /
  窗口 / 静默期 / 通知渠道多选），状态与最后触发时间展示。

#### 用户洞察

- **用户 360**：行为序列（BehaviorSequence）+ 用户标签（UserTag）+
  身份识别（ID-Mapping）三页签个体档案页，按用户 ID 聚合查询。
- **会话回放**：`uba_replay_chunks` 表（session_id + chunk_seq 唯一幂等锚）
  + ReplayService（UploadChunk / ListSessions / GetSessionEvents /
  DeleteSession）；collector 新增 `POST /uba/v1/replays` 上传端点
  （LookupApp 只读旁路鉴权）；SDK `startReplayRecording()` 基于 rrweb
  动态导入（未安装优雅降级），分块缓冲（100 事件/块 + 10s 定时）上传；
  管理端 rrweb-player 播放抽屉（动态导入 + CSS，会话全部分块按序拼接喂入）；
  `REPLAY_RETENTION_DAYS` 环境变量控制保留期（缺省 30 天），asynq
  每天 03:30 自动清扫超期分块。

#### 功能开关 / A-B 实验

- `uba_feature_flags` 表：租户内唯一 key、多变量（千分比权重）、
  放量百分比（0~100）、可选目标转化事件、起止时间。
- **一致性哈希分配引擎**：FNV-1a 哈希 `flagKey|distinctID`，同一身份
  永远得到同一变体；登录用户优先、设备 ID 兜底；rollout 与变体权重
  按千分比归一切分，权重全 0 时均分。
- collector 新增 `POST /uba/v1/flags` 端点：复用 report 链路的应用鉴权，
  转发 core `EvaluateFlags`；flag 服务不可用时返回空结果降级。
- SDK `loadFlags()` / `getFlag()` / `isFeatureEnabled()`；首次 `getFlag`
  自动上报 `$feature_flag` 曝光事件（会话内每 flag 去重一次）；
  `identify` / `resetUser` 使取值缓存失效。
- 实验结果：Doris CTE 查询以曝光事件圈定各变体登录用户集，LEFT JOIN
  目标事件算去重转化；管理端「查看实验结果」抽屉展示曝光数、转化数、
  转化率进度条；变体行编辑器（key / 显示名 / 千分比权重 / 增删行 /
  唯一性与权重校验）。

#### 通知域（参考 go-wind-admin 同名域移植）

- `sys_notification_channels`（渠道花名册）：SMTP（TlsMode + 密码加密存储 +
  列表脱敏） / SMS（阿里云 / 腾讯云签名） / Webhook（五种 SignStyle 签名风格
  + SSRF 防线 + 载荷模板）；渠道增删改查、启停、测试发送。
- `sys_notification_rules`（事件路由规则）：一事件一渠道路由 + 是否异步 +
  启停；空表自动播 4 条默认规则（站内信走台账、监控告警走站内信、
  风险告警走 Webhook 异步、报告走站内信）；试投端点按规则当场走完整链路。
- `sys_notification_templates`（消息模板）：`{{var}}` 占位符渲染
  （未识别占位符报错、未闭合保留原文）；按 code 引用，发送方以
  `template_code` 覆写标题 / 正文；试渲染端点供管理页预览。
- `sys_notification_deliveries`（投递台账）：一条通知 × 一个渠道 × 一个
  收件人 = 一行；状态机 SENDING → SENT / FAILED / SKIPPED；幂等锚
  `(request_id, channel)` 唯一约束在 DB 层挡重复；target 脱敏（邮箱留
  首字符 + 域名，其余留末 4 位）；超期 SENDING 由清扫任务（每天 03:30）
  结算为 FAILED。
- `notification_preferences`（用户通知偏好）：静音时段（跨零点合法）+
  分类退订（仅约束全员广播）；只抑制 SSE 实时推送，不抑制落库。
- 发送引擎 `SendDirect` 全流程：模板渲染 → 规则路由（请求显式渠道可覆盖）
  → 落 SENDING 台账 → 渠道 Sender → 同步 / asynq 异步（幂等门 =
  状态非 SENDING 即跳过，SkipRetry 防验证码滞留归档队列）→ 结论回写；
  target 脱敏；INTERNAL 渠道适配器（全仓唯一渠道实现反向依赖服务层的边）
  复用 InternalMessageService 收件行落库 + SSE 推送。
- collector 新增 `POST /uba/v1/flags` 端点：复用 report 链路的应用鉴权
  （app_id / app_secret 常量时间比较 + Redis 缓存），转发 core
  `EvaluateFlags`；flag 服务不可用时返回空结果降级。

#### 自助 SQL 查询

- core 新增 `SqlQueryService.ExecuteQuery`，admin BFF 转发 REST
  `POST /admin/v1/sql-query/execute`。
- 安全管控：仅允许单条 `SELECT / SHOW / DESCRIBE / EXPLAIN`（首关键字
  白名单），拒绝分号多语句与 `--` / `/*` / `#` 注释片段，行数上限
  （默认 200 / 硬顶 1000，多读一行判定 `truncated`），服务端 30 秒超时。
- 结果集用 `google.protobuf.Value` 强类型承载；SQL 控制台页等宽编辑器
  + Ctrl/Cmd+Enter 执行 + 行数上限选择 + 动态列表格 + 截断/耗时标签
  + CSV 导出。

#### 风险事件推送与后台引擎

- **风险事件推送**：core 订阅 Kafka 风险事件流（消费组 `uba_alert_sender`），
  high 以上级别按 `risk.<level>` 事件类型推给匹配的网络钩子；
  `data.kafka` 未配置时静默降级。
- **ID-Mapping 引擎**：每 30 分钟回扫 `events_fact` 里 user_id > 0 且带
  device_id 的事件，把共现对以 `link_source=algorithm` upsert 进 PG 的
  `id_mapping`（按唯一键冲突更新 last_seen），并同步一份到 Doris
  （UNIQUE KEY 幂等）。
- **动态标签执行器**：`TagDefinition.rule.expression` 作为 `users_dim`
  上的布尔谓词周期执行（`refresh_interval_seconds` 到期判定 + SQL 黑名单
  校验），命中人群全量替换进 user_tags 并同步 Doris；TagDefinition 新增
  `last_refreshed_at`。
- **风险事件 Webhook 推送**：core 订阅 Kafka 风险事件流（消费组
  `uba_alert_sender`），high 以上级别按 `risk.<level>` 事件类型推给
  匹配的网络钩子。

#### 采集端补强

- **SDK 全埋点**：自动 `page_view`（包装 `history.pushState/replaceState`
  + `popstate` / `hashchange` SPA 路由监听）、页面停留（切页带 `durationMs`、
  `pagehide` 补发 `page_dwell`，≥1 秒才报避免噪声）。
- **SDK 离线补发**：未上报队列镜像到 localStorage（节流写盘、上限可配），
  上报失败 / beacon 失败 / 页面崩溃不再静默丢事件——下次打开自动补发。
- **SDK 采样与合规**：`sampleRate` 会话级确定性采样（按 sessionId 哈希判定，
  同一会话全采或全不采保证漏斗口径一致）；`requireConsent` + `setConsent`
  授权门控；`respectDoNotTrack` 尊重浏览器 DNT。
- **collector 服务端补全**：真实 IP 提取（XFF 链首 → X-Real-IP → RemoteAddr，
  与审计日志同口径）；UA 解析（`mileusna/useragent`，os 列 + context 属性
  browser / device_type）；IP 地理兜底（GeoLite 客户端，回环 / 私网跳过）。
- **collector 白名单**：新端点 `ListVisualEvents` / `UploadReplayChunk` /
  `PostFlags` 补入 auth 白名单（各自的 handler 内已做应用级鉴权或降级）。

#### 历史数据批量导入

- `uba_import_jobs` 表：events_json 缓存 / 统计三列（总 / 成功 / 失败）/
  错误摘要（前 20 条）/ 状态机（PENDING → RUNNING → COMPLETED / PARTIAL /
  FAILED / CANCELLED）。
- 执行：创建后后台 goroutine 解析 events_json → 分批（500/批）写入 Doris
  `events_fact`（走 EventsFactRepo.Create，UNIQUE KEY 幂等去重），完成后
  清空原始 JSON 释放存储；支持取消（状态机 + 执行中检查）。
- 管理端导入任务列表页（状态徽标 / 进度 / 耗时 / 取消 / 删除），JSON 文件
  上传 + 格式校验；新菜单「历史数据导入」。

#### 分析看板补全

- 看板渲染器扩展：在原有 event-trend / active-users / group-by 基础上补
  funnel（漏斗图）与 retention（同期群留存热力图，按 offset 展开百分比）。
- 分析结果 CSV 导出：`csv-export.ts` 工具（RFC4180 转义 + UTF-8 BOM），
  接入事件趋势 / 漏斗 / 留存 / 用户分群四个核心分析页。

#### 通知域（参考 go-wind-admin 同名域移植）

- 契约：`notification/service/v1`（SendDirect / ListDelivery / GetDelivery +
  RuleService / TemplateService / PreferenceService）+
  `notification_channel/service/v1`（ChannelService 含 SendTestEmail /
  SendTest）+ admin BFF REST 壳五件。
- 发送引擎 `SendDirect` 全流程：模板渲染 → 规则路由（请求显式渠道可覆盖）
  → 落 SENDING 台账 → 渠道 Sender → 同步 / asynq 异步（幂等门 = 状态非
  SENDING 即跳过，SkipRetry 防验证码滞留归档队列）→ 结论回写。
- 渠道实现：SMTP（go-utils/mailer，TlsMode + 密码加密存储 + 列表脱敏）、
  SMS（阿里云 / 腾讯云签名，`go-utils/notify`）、Webhook（go-utils/notify
  五种 SignStyle + SSRF 防线 + 载荷模板）、INTERNAL 缝（复用
  InternalMessageService 收件行落库 + SSE 推送——全仓唯一渠道实现反向
  依赖服务层的边）。
- 用户通知偏好：静音时段（跨零点合法）+ 分类退订（仅约束全员广播）；
  只抑制 SSE 实时推送，不抑制落库。
- admin BFF 新增五个转发服务（channels / rules / templates / deliveries /
  preferences）+ REST 注册。
- 前端四页：渠道花名册（按类型显隐字段、密钥只写不读 `has_*`、测试发送）、
  事件路由规则（含试投）、消息模板（试渲染预览）、投递台账（只读 / 脱敏）；
  新菜单「通知管理」（平台管理员）。
- EventType 补 `RISK_ALERT` / `REPORT_GENERATED` / `CAMPAIGN`（uba 自有事件）。

### 变更

- **后台菜单重排**：按用户旅程分组（核心分析 100 → 增长闭环 200-250 →
  用户洞察 300-310 → 数据配置 400-440 → 风险与通知 500-510 →
  系统管理 900-960），百位分组 + 十位留空隙便于插入；此前按加功能时间
  递增（100→600 / 1004→1014 / 2000+），缺乏信息架构。
- **core gRPC ent 中间件**：内部服务间调用（collector→core / admin BFF
  转发）gRPC 元数据可能不含操作者信息——此时回落系统上下文
  （`NewSystemViewerContext`），与 asynq StartAllTask 同一信任模型；
  此前无元数据时直接放行导致租户隐私层 401。
- **core gRPC 新增注册**：通知域五服务（Notification / Rule / Template /
  Channel / Preference）+ 导入服务 + 可视化埋点 + 会话回放 + 自助 SQL。
- **Makefile api 目标**：buf 模板引用改为 `buf.go.tag.gen.yaml`——旧版 buf
  自动补 `.yaml` 后缀，新版把它当内联 JSON 解析直接报错；修复后
  `make api` 在新版 buf 下恢复可用。
- **proto 枚举前缀去重**：ReportChannel 枚举值加 `REPORT_CHANNEL_` 前缀，
  避免与 AlertChannel 在同一 proto 包内重名冲突。
- README 三语同步增长闭环九项 + 通知域四项描述；SDK README 补全埋点 /
  离线补发 / 采样合规 / flags / 可视化埋点 / 回放 API 文档与新模块目录；
  AGENTS.md 分析模型口径修正（28+ → 25）并补 views 新页面目录。

### 修复

- 修复 flag 实验结果查询的纳秒/毫秒混用：`int64(30*24*time.Hour)` 把
  纳秒值当毫秒传给 `time.UnixMilli`，产生 `year=-80111` 非法时间导致
  Doris 查询失败；改为显式毫秒计算 `int64(30*24*60*60*1000)`。
- 修复通知域平台管理员守卫位置：从 core 侧（无 JWT 中间件 → claims 永远
  为空 → 401）迁至 admin BFF 边缘；平台管理员判定改走角色码
  `platform:admin`（JWT `roc` claim），而非不存在的 `is_platform_admin`
  字段（proto 有此字段但全仓无代码填充）。
- 修复 collector 白名单漏收新端点：`ListVisualEvents` / `UploadReplayChunk`
  / `PostFlags` 未加入 auth 白名单，auth 中间件 401 拦截导致 SDK 拿不到
  数据；补录后 flags / visual-events / replays 三个端点正常可达。
- 修复 Doris BE broken disk 崩溃循环（574 次重启）：BE 自动持久化
  `broken_storage_path` 到 `be_custom.conf`（Oct 14 非正常关机时一次
  DataDir 加载瞬时失败触发），此后每次启动读回该标记拒绝启动。清除
  be_custom.conf 中的 broken_storage_path 行后恢复，数据零丢失。
- 修复 admin→core gRPC 内部调用被系统代理劫持：环境变量
  `HTTP_PROXY=http://127.0.0.1:7890` 导致 Go gRPC 拨号走代理——代理不认识
  etcd 服务发现的内部地址直接 502。启动服务时设 `NO_PROXY=localhost,...`
  排除宿主机内网 IP 即可；`pkg/middleware/ent/ent.go` 同时回落系统上下文
  兜底无操作者元数据的内部调用。

### 初始平台基线（2026-10-09）

- 基于多租户的 RBAC 体系：用户、租户、角色、权限、菜单、组织、职位管理。
- 25 种分析模型：通用行为分析 10（事件趋势 / 漏斗 / 留存 / 维度对比 /
  路径桑基 / 点击热力图 / 归因 / 分布 / 新老对比 / 行为序列）、
  用户深度洞察 9（生命周期 / 流失回流 / 间隔时间 / 矩阵象限 / LTV /
  会话分析 / 异常检测 / 在线统计 / 付费营收）、
  游戏专项 6（关卡分析 / 付费分层 / 滚服留存 / 经济系统 / PCU 在线 /
  实时大屏）。
- Web SDK（TypeScript）：代码埋点 + 点击自动采集（坐标 / XPath / 视口宽度），
  批量上报 + 指数退避重试 + sendBeacon 卸载兜底。
- Collector 埋点采集服务：应用级鉴权（app_id / app_secret SHA-256 常量时间
  比较 + Redis 正 / 负缓存）、事件校验（必填字段 + oneof payload）、
  租户覆盖（服务端权威值覆盖 SDK 上报值，杜绝跨租户伪造）、Kafka 转发
  （uba_events_raw / uba_risk_events）。
- Core 核心服务：分析引擎（25 个模型的 Doris OLAP 查询，含维度拆分 /
  属性过滤 / 多事件多序列）、事件管理、用户管理、系统管理。
- Admin BFF：管理后台接口聚合，REST API 面向前端，etcd 服务发现 +
  gRPC 负载均衡。
- uba-ingest：Doris 入仓运维 CLI（`apply` 建表 + Routine Load、
  `status` 健康度、`etl` 每日回算、`render` 模板渲染），SQL 以
  `go:embed` 编进二进制。
- 多级物化视图：事件日/时聚合、用户留存、地理分布、渠道版本、热门路径等
  9 个 MV + 12 个视图。
- 五类审计日志（登录 / API / 操作 / 数据访问 / 权限评估）生产者全链路。
- 实时大屏（轮询刷新：KPI + 事件流 + 风险流 + 平台分布）。
- 风控规则引擎（可视化配置多维度条件组合）+ 风险事件管理（自动检测 +
  人工审核处理）+ Webhook 告警。
- 多语言（中文 / English / 日本語）+ 暗色主题 + Docker Compose 部署。
