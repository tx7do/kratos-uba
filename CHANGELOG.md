# 更新日志

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/) 规范，
版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

> 历史条目基于 Git 提交记录回溯整理，仅记录对使用者有影响的功能性变更，
> 日常重构、代码风格、CI 调整等不在逐一列举。

## [1.1.0] - 2026-10-15

### 新增

- 增长闭环全链路（分析 → 存档 → 触达 → 行动）：
  - 持久化用户分群：圈选定义保存、定时/手动刷新、成员下钻；漏斗与圈选页一键存为分群。
  - 自定义看板：分析模型组件自由组页（趋势/活跃/分布/漏斗/留存），参数快照存档。
  - 指标阈值告警：事件量/活跃/金额阈值监控，站内信 + Webhook 双通道，静默期防打扰。
  - 定时报告：周期执行分析查询，摘要与 CSV 经站内信/Webhook 定时送达。
- 功能开关 / A-B 实验：多变量 + 千分比权重 + 放量比例，FNV-1a 一致性哈希稳定分流，
  SDK `loadFlags` / `getFlag` / `isFeatureEnabled` 消费，首次取值自动上报 `$feature_flag`
  曝光事件（会话去重）；实验结果按变体聚合曝光与转化。
- 通知域（参考 go-wind-admin 同名域移植）：渠道花名册（SMTP / SMS / Webhook）+
  事件路由规则 + 消息模板（`{{var}}` 渲染）+ 投递台账（幂等锚 / 脱敏 / 超期清扫）+
  用户通知偏好（静音时段 + 分类退订）；collector 新增 flags / visual-events / replays
  公开端点（LookupApp 只读旁路鉴权）。
- 营销触达：分群 × 文案 × 站内信群发，立即/定时发送（asynq ProcessAt），发送进度与统计。
- 可视化埋点：页面 + 元素 XPath → 事件名规则管理，SDK 自动匹配上报，`startVisualPicker`
  调试选点模式。
- 会话回放：rrweb 录制分块上传，会话聚合查询，管理端 rrweb-player 播放抽屉，
  `REPLAY_RETENTION_DAYS` 保留期自动清扫。
- 自助 SQL 查询：只读白名单 + 行数上限 + 30s 超时的 OLAP 查询控制台，结果可导出 CSV。
- 风险事件推送：core 订阅 Kafka 风险事件流（消费组 uba_alert_sender），高风险自动触发 Webhook。
- ID-Mapping 引擎：定期回扫 events_fact 登录用户↔设备共现对，自动构建跨端身份关联。
- 动态标签执行器：规则表达式周期计算人群，自动写入用户标签。
- 采集端补强：
  - SDK 全埋点（自动 page_view / SPA 路由 / 页面停留 / 离线补发 / 采样 / 隐私合规）
  - 服务端 UA 解析（os/browser/设备类型）+ IP 地理兜底（GeoLite）
  - SDK 可视化埋点消费（loadVisualEvents / startVisualPicker）
  - SDK 会话回放录制（startReplayRecording，rrweb 动态导入）

### 变更

- 后台菜单重排：按用户旅程分组（核心分析 → 增长闭环 → 用户洞察 → 数据配置 → 风险通知 → 系统管理），
  百位分组 + 十位留空隙便于插入。
- core gRPC ent 中间件：无操作者元数据的内部服务间调用回落系统上下文（内网信任模型）。
- core gRPC 新增通知域五服务 + 导入服务 + 可视化埋点 + 会话回放 + 自助 SQL 注册。
- README 三语同步增长闭环与通知域描述。
- Makefile api 目标 buf 模板引用改为 `.yaml` 后缀，修复新版 buf 下的构建。

### 修复

- 修复通知域 be_custom.conf 遗留 broken_storage_path 标记导致的 Doris BE 崩溃循环
  （574 次重启）：清除自动持久化的 broken 标记后恢复，数据零丢失。
- 修复 flag 实验结果查询的纳秒/毫秒混用（`int64(30*24*time.Hour)` 误作毫秒传给
  `time.UnixMilli`，产生 year=-80111 非法时间）。
- 修复通知域平台管理员守卫位置：从 core 侧迁至 admin BFF 边缘（JWT 在 BFF 验证后
  才有 claims），平台管理员判定改走角色码 `platform:admin`。

## [1.0.0] - 2026-10-09

### 新增

- 基于多租户的 RBAC 体系：用户、租户、角色、权限、菜单、组织、职位管理。
- 25 种分析模型：通用行为分析 10 / 用户深度洞察 9 / 游戏专项 6，全部走 Doris OLAP 查询。
- Web SDK（TypeScript）：代码埋点 + 点击自动采集（坐标 / XPath / 视口宽度），批量上报 + 重试 + sendBeacon 兜底。
- Collector 埋点采集服务：应用级鉴权（app_id / app_secret）、事件校验、租户覆盖、Kafka 转发。
- Core 核心服务：分析引擎（25 个模型的 Doris OLAP 查询）、事件管理、用户管理、系统管理。
- Admin BFF：管理后台接口聚合，REST API 面向前端。
- uba-ingest：Doris 入仓运维 CLI（建表 · Routine Load · 每日 ETL · 健康度）。
- 多级物化视图：事件日聚合、用户留存、地理分布、渠道版本、热门路径等。
- 审计日志五种（登录 / API / 操作 / 数据访问 / 权限评估）。
- 三语 README（中文 / English / 日本語）。
- Docker Compose 部署支持（全栈中间件 + 后端服务一键部署）。
