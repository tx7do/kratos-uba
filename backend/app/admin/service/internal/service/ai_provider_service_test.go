package service

import "testing"

// TestMaskApiKey 密钥脱敏：保留前 3 后 4，其余以 *** 概括；过短整体 ***。
func TestMaskApiKey(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", "***"},
		{"sk-1", "***"},
		{"sk-123", "***"},
		{"sk-12345", "sk-***2345"},
		{"sk-abcdefghijklmnop", "sk-***mnop"},
	}
	for _, c := range cases {
		if got := maskApiKey(c.in); got != c.want {
			t.Fatalf("maskApiKey(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
