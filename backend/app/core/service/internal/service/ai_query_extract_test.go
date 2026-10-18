package service

import (
	"strings"
	"testing"
)

// 复现线上"no SQL statement found"：DeepSeek 在 SQL 前后带说明文字与围栏时的提取。
func TestExtractSQL_PrefixedFence(t *testing.T) {
	raw := "您这个问题我没办法直接帮您查——我这边没有连接到任何数据库或日志系统，看不到您环境里的实际记录。\n\n" +
		"如果您是在自己系统上排查，可以参考下面这个思路：\n" +
		"```sql\n" +
		"SELECT risk_event_id, user_id, risk_type, occur_time\n" +
		"FROM risk_events\n" +
		"WHERE risk_level = 'HIGH'\n" +
		"  AND occur_time >= now() - INTERVAL 24 HOUR\n" +
		"ORDER BY occur_time DESC\n" +
		"LIMIT 100;\n" +
		"```\n" +
		"几点说明，避免你直接跑出问题。"

	got, err := sanitizeSQLFor(raw, nil)
	if err != nil {
		t.Fatalf("sanitizeSQL rejected: %v", err)
	}
	if !aiQuerySelectPattern.MatchString(got) {
		t.Fatalf("extracted SQL invalid: %q", got)
	}
	t.Logf("extracted: %s", got)
}

// 租户用户生成的 SQL 缺 tenant_id 过滤必须被拒（fail-closed）。
func TestSanitizeSQL_MissingTenantFilter(t *testing.T) {
	raw := "SELECT event_name, COUNT(*) AS cnt FROM events_fact GROUP BY event_name LIMIT 10"
	tenantPattern := tenantIDPatternFor(7)
	if _, err := sanitizeSQLFor(raw, tenantPattern); err == nil {
		t.Fatalf("expected rejection for missing tenant_id filter")
	}

	withTenant := "SELECT event_name, COUNT(*) AS cnt FROM events_fact WHERE tenant_id = 7 GROUP BY event_name LIMIT 10"
	if _, err := sanitizeSQLFor(withTenant, tenantPattern); err != nil {
		t.Fatalf("sanitizeSQL rejected tenant-scoped sql: %v", err)
	}
}

// 白名单外表必须被拒；写操作必须被拒。
// （注释与多语句不在此列：extractSQL 剥行注释、截分号只留首条语句，
// 属预期内的消毒行为，由 TestSanitizeSQL_MultiStatementTruncated 钉住。）
func TestSanitizeSQL_Rejects(t *testing.T) {
	cases := []string{
		"SELECT * FROM pg_catalog.pg_tables LIMIT 10",
		"SELECT * FROM some_unknown_table LIMIT 10",
		"DELETE FROM events_fact",
	}
	for _, raw := range cases {
		if _, err := sanitizeSQLFor(raw, nil); err == nil {
			t.Fatalf("expected rejection for: %s", raw)
		}
	}
}

// 多语句在提取期截断为首条只读语句（后续语句全部丢弃）。
func TestSanitizeSQL_MultiStatementTruncated(t *testing.T) {
	got, err := sanitizeSQLFor("SELECT event_name FROM events_fact LIMIT 1; DROP TABLE events_fact", nil)
	if err != nil {
		t.Fatalf("sanitizeSQL rejected: %v", err)
	}
	if strings.Contains(got, "DROP") || strings.Contains(got, ";") {
		t.Fatalf("second statement not truncated: %q", got)
	}
}
