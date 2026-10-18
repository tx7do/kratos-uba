package service

import (
	"testing"
	"time"
)

// TestClampToolDays 工具时间窗钳制：缺省 7、上限 90、非法回落 7。
func TestClampToolDays(t *testing.T) {
	cases := []struct {
		in   int32
		want int32
	}{
		{0, 7},
		{-5, 7},
		{1, 1},
		{7, 7},
		{89, 89},
		{90, 90},
		{91, 90},
		{1000, 90},
	}
	for _, c := range cases {
		if got := clampToolDays(c.in); got != c.want {
			t.Fatalf("clampToolDays(%d) = %d, want %d", c.in, got, c.want)
		}
	}
}

// TestStartOfDaysWindow 最近 N 天窗口起点：今天 0 点往前推 N-1 天。
func TestStartOfDaysWindow(t *testing.T) {
	now := time.Date(2026, 10, 18, 15, 30, 0, 0, time.Local)

	got := startOfDaysWindow(now, 1)
	want := time.Date(2026, 10, 18, 0, 0, 0, 0, time.Local)
	if !got.Equal(want) {
		t.Fatalf("startOfDaysWindow(1 day) = %v, want %v", got, want)
	}

	got = startOfDaysWindow(now, 7)
	want = time.Date(2026, 10, 12, 0, 0, 0, 0, time.Local)
	if !got.Equal(want) {
		t.Fatalf("startOfDaysWindow(7 days) = %v, want %v", got, want)
	}
}

// TestToolAppIdPtr appId=0 视为未指定转 nil，非 0 原值透传。
func TestToolAppIdPtr(t *testing.T) {
	if p := toolAppIdPtr(0); p != nil {
		t.Fatalf("toolAppIdPtr(0) = %v, want nil", p)
	}
	p := toolAppIdPtr(42)
	if p == nil || *p != 42 {
		t.Fatalf("toolAppIdPtr(42) = %v, want 42", p)
	}
}
