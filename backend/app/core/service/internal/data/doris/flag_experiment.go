package doris

import (
	"context"
	"fmt"
	"time"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

// FlagVariantStat 单变体的曝光/转化统计。
type FlagVariantStat struct {
	Variant        string
	ExposedUsers   uint64
	ConvertedUsers uint64
}

// FlagExperimentResult 聚合 A/B 实验结果：以 $feature_flag 曝光事件圈定各变体的
// 登录用户集，统计每个变体里触发过目标事件的去重用户数。
// appFilter 过滤 events_fact.tenant_id（0 表示不过滤）。
func (r *AnalyticsRepo) FlagExperimentResult(
	ctx context.Context,
	flagKey, goalEvent string,
	appFilter uint32,
	startMs, endMs int64,
) ([]FlagVariantStat, error) {
	goal := goalEvent
	if goal == "" {
		goal = "$feature_flag" // 无目标事件时只出曝光（goal 等于曝光事件 → converted == exposed）
	}

	tenantCond := ""
	args := []any{flagKey, time.UnixMilli(startMs), time.UnixMilli(endMs)}
	if appFilter != 0 {
		tenantCond = "AND tenant_id = ?"
		args = append(args, appFilter)
	}

	goalTenantCond := ""
	if appFilter != 0 {
		goalTenantCond = "AND g.tenant_id = ?"
		args = append(args, appFilter)
	}
	args = append(args, goal, time.UnixMilli(startMs), time.UnixMilli(endMs))

	q := fmt.Sprintf(`
WITH exposed AS (
  SELECT user_id, MAX(context['variant']) AS variant
  FROM events_fact
  WHERE event_name = '$feature_flag'
    AND context['flag_key'] = ?
    AND user_id > 0
    AND event_time >= ? AND event_time < ?
    %s
  GROUP BY user_id
)
SELECT e.variant AS variant,
       COUNT(*) AS exposed_users,
       COUNT(DISTINCT g.user_id) AS converted_users
FROM exposed e
LEFT JOIN events_fact g
  ON g.user_id = e.user_id
 AND g.event_name = ?
 AND g.event_time >= ? AND g.event_time < ?
 %s
GROUP BY e.variant`, tenantCond, goalTenantCond)

	rows, err := r.db.DB().QueryxContext(ctx, q, args...)
	if err != nil {
		r.log.Errorf("failed to query flag experiment result: %v", err)
		return nil, ubaV1.ErrorInternalServerError("failed to query flag experiment result")
	}
	defer rows.Close()

	stats := make([]FlagVariantStat, 0, 8)
	for rows.Next() {
		var stat FlagVariantStat
		if err := rows.Scan(&stat.Variant, &stat.ExposedUsers, &stat.ConvertedUsers); err != nil {
			r.log.Errorf("failed to scan flag variant stat: %v", err)
			continue
		}
		stats = append(stats, stat)
	}
	if err := rows.Err(); err != nil {
		r.log.Errorf("iterate flag variant stats failed: %v", err)
		return nil, ubaV1.ErrorInternalServerError("iterate flag variant stats failed")
	}

	return stats, nil
}
