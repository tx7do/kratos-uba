package service

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/types/known/structpb"
)

// TestStringifyStructValue 问数结果单元格的结构值 → 展示字符串转换。
// nil/Null → 空串，字符串原样，数字 %v，布尔 true/false，复合值走 JSON 形态。
func TestStringifyStructValue(t *testing.T) {
	nullVal, _ := structpb.NewValue(nil)
	listVal, _ := structpb.NewValue([]any{1, "a"})

	cases := []struct {
		name string
		in   *structpb.Value
		want string
	}{
		{"nil", nil, ""},
		{"null", nullVal, ""},
		{"string", structpb.NewStringValue("hello"), "hello"},
		{"number-int", structpb.NewNumberValue(42), "42"},
		{"number-float", structpb.NewNumberValue(1.5), "1.5"},
		{"bool-true", structpb.NewBoolValue(true), "true"},
		{"bool-false", structpb.NewBoolValue(false), "false"},
		// 复合值走 fmt %v 的 Go 默认形态（展示用，不追求 JSON 规范形）
		{"list", listVal, "[1 a]"},
	}
	for _, c := range cases {
		if got := stringifyStructValue(c.in); got != c.want {
			t.Fatalf("%s: stringifyStructValue(%v) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}

// TestBuildReportAiPrompt 报表 AI 摘要提示词组装：报表名/模型/确定性摘要
// 必须原样进入提示词；CSV 超过 20 行截断到表头+20 行并追加省略号。
func TestBuildReportAiPrompt(t *testing.T) {
	csv := "timestamp,value\n"
	for i := 0; i < 30; i++ {
		csv += "1690000000000,1\n"
	}

	got := buildReportAiPrompt("周报", "group-by", "总量 30。", csv)

	if !strings.Contains(got, "报表：周报") || !strings.Contains(got, "分析模型：group-by") {
		t.Fatalf("prompt missing report name/model: %q", got)
	}
	if !strings.Contains(got, "总量 30。") {
		t.Fatalf("prompt missing deterministic summary: %q", got)
	}
	if strings.Count(got, "1690000000000,1") != 20 {
		t.Fatalf("csv rows not clamped to 20: count=%d", strings.Count(got, "1690000000000,1"))
	}
	if !strings.Contains(got, "...") {
		t.Fatalf("truncated csv missing ellipsis marker")
	}
}
