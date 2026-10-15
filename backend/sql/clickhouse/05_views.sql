-- ============================================================
-- UBA 系统 - 物化视图设计
-- 执行顺序：5
--
-- 注意：本文件【不再】创建 mv_sessions_fact / mv_users_dim。
--   物化视图按插入块增量聚合，跨块的同一会话/用户会被拆成多行或以后块覆盖前块，
--   first_active_date 等累计值也只有块内局部真值。sessions_fact / users_dim 改由
--   `uba-ingest ch-etl` 按 Doris 06_etl §1/§2 同款口径做全历史回算（幂等），
--   新鲜度与 Doris 侧一致（按天批量）。存量部署重跑本脚本时，开头的 DROP 会
--   移除旧 MV，已写入的局部数据由回算覆盖（ReplacingMergeTree 按 updated_at 取新）。
-- ============================================================

-- 迁移：移除旧的两条派生 MV（历史局部数据由 ch-etl 回算覆盖）
DROP VIEW IF EXISTS gw_uba.mv_sessions_fact;
DROP VIEW IF EXISTS gw_uba.mv_users_dim;

-- ============================================================
-- 1. 物化视图 - 路径聚合
-- ============================================================
CREATE MATERIALIZED VIEW IF NOT EXISTS gw_uba.mv_path_features
    TO gw_uba.path_features
AS
SELECT concat(session_id, '_', toString(min(event_ts)))                                     AS path_id,
       tenant_id,
       user_id,
       session_id,

       hex(MD5(arrayStringConcat(groupArray(event_name), '->')))                            AS path_hash,
       if(length(groupArray(event_name)) > 0, arrayElement(groupArray(event_name), 1), '')  AS first_event,
       if(length(groupArray(event_name)) > 0, arrayElement(groupArray(event_name), -1), '') AS last_event,
       length(groupArray(event_name))                                                       AS path_length,
       arraySlice(groupArray(event_name), 1, 3)                                             AS first_3_events,
       arraySlice(groupArray(event_name), -3)                                               AS last_3_events,

       maxIf(1, event_name = 'purchase_success')                                            AS is_converted,
       'purchase_success'                                                                   AS conversion_event,
       maxIf(event_time, event_name = 'purchase_success')                                   AS conversion_time,

       min(event_time)                                                                      AS start_time,
       max(event_time)                                                                      AS end_time,
       max(event_ts) - min(event_ts)                                                        AS total_duration_ms,
       count()                                                                              AS step_count

FROM gw_uba.events_fact
WHERE session_id != ''
GROUP BY tenant_id, user_id, session_id;

-- ============================================================
-- 2/3. （已移除）会话聚合 / 用户画像聚合 MV
--   交给 `uba-ingest ch-etl` 全历史回算，见文件头说明。
-- ============================================================


-- ============================================================
-- 4. 物化视图：风险事件 → 用户风险画像
-- ============================================================
CREATE MATERIALIZED VIEW IF NOT EXISTS gw_uba.mv_risk_to_user_risk_profile
    TO gw_uba.user_risk_profile
AS
SELECT tenant_id,
       user_id,

       max(risk_score)                AS risk_score,
       argMax(risk_level, occur_time) AS risk_level,
       groupUniqArray(risk_type)      AS risk_tags,
       max(occur_time)                AS last_risk_time,

       now()                          AS updated_at

FROM gw_uba.risk_events
WHERE user_id != 0
GROUP BY tenant_id, user_id;
