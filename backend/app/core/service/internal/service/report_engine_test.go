package service

import (
	"testing"
	"time"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

func TestParseReportParams(t *testing.T) {
	req := &ubaV1.EventTrendRequest{}
	// 空参数：零值保持
	if err := parseReportParams("", req); err != nil {
		t.Fatalf("empty params should pass: %v", err)
	}
	if req.GetGranularity() != ubaV1.AnalyticsGranularity_ANALYTICS_GRANULARITY_UNSPECIFIED {
		t.Fatalf("empty params should leave zero value, got %v", req.GetGranularity())
	}

	// JSON 快照往返（camelCase json_name）
	if err := parseReportParams(`{"granularity":"DAY","eventName":"pay_success"}`, req); err != nil {
		t.Fatalf("valid params should pass: %v", err)
	}
	if req.GetGranularity() != ubaV1.AnalyticsGranularity_DAY || req.GetEventName() != "pay_success" {
		t.Fatalf("params not applied: %+v", req)
	}

	// 未知字段丢弃（DiscardUnknown）
	if err := parseReportParams(`{"unknownField":1}`, req); err != nil {
		t.Fatalf("unknown fields should be discarded: %v", err)
	}

	// 非法 JSON 报错
	if err := parseReportParams(`{bad json}`, &ubaV1.EventTrendRequest{}); err == nil {
		t.Fatal("invalid JSON should fail")
	}
}

func TestEnsureReportRange(t *testing.T) {
	now := time.Now()
	interval := 24 * time.Hour

	// 已配置范围：原样保留
	existing := &ubaV1.TimeRange{StartMs: 1000, EndMs: 2000}
	if got := ensureReportRange(existing, interval, now); got != existing {
		t.Fatal("existing range should be kept")
	}

	// 未配置：按间隔回看
	fresh := ensureReportRange(nil, interval, now)
	if fresh.GetStartMs() != now.Add(-interval).UnixMilli() || fresh.GetEndMs() != now.UnixMilli() {
		t.Fatalf("unexpected default range: %+v", fresh)
	}

	// 零值范围：同样回看
	if got := ensureReportRange(&ubaV1.TimeRange{}, interval, now); got.GetStartMs() == 0 {
		t.Fatal("zero range should be replaced with interval window")
	}
}

func TestCsvFromRows_Basic(t *testing.T) {
	got := csvFromRows(
		[]string{"name", "value"},
		[][]string{
			{"a", "1"},
			{"with,comma", `with"quote`},
			{"multi\nline", "2"},
		},
	)
	want := "\"name\",\"value\"\n" +
		"\"a\",\"1\"\n" +
		"\"with,comma\",\"with\"\"quote\"\n" +
		"\"multi line\",\"2\"\n"
	if got != want {
		t.Fatalf("csv mismatch:\n got: %q\nwant: %q", got, want)
	}
}
