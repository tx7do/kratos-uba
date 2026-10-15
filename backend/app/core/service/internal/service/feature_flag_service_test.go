package service

import (
	"testing"

	"go-wind-uba/app/core/service/internal/data/ent"
	"go-wind-uba/app/core/service/internal/data/ent/schema"
)

func flagOf(key string, rollout int32, weights ...int32) *ent.FeatureFlag {
	variants := make([]schema.FeatureVariant, 0, len(weights))
	keys := []string{"control", "treatment_a", "treatment_b"}
	for i, w := range weights {
		variants = append(variants, schema.FeatureVariant{Key: keys[i], Weight: w})
	}
	return &ent.FeatureFlag{
		Enabled:           true,
		Key:               &key,
		RolloutPercentage: rollout,
		Variants:          variants,
	}
}

func TestAssignFlagVariant_Basic(t *testing.T) {
	flag := flagOf("exp", 100, 500, 500)

	seen := map[string]int{}
	for i := 0; i < 1000; i++ {
		distinctID := "u:" + string(rune('a'+i%26)) + string(rune('0'+i%10)) + string(rune(i))
		variant := AssignFlagVariant(flag, distinctID)
		if variant == "" {
			t.Fatalf("rollout=100 should never miss, got miss for %q", distinctID)
		}
		if variant != "control" && variant != "treatment_a" {
			t.Fatalf("unexpected variant %q", variant)
		}
		seen[variant]++

		// 一致性：同一身份重复评估结果必须相同
		if again := AssignFlagVariant(flag, distinctID); again != variant {
			t.Fatalf("assignment not stable for %q: %q vs %q", distinctID, variant, again)
		}
	}

	// 五五开：允许明显偏差但不应全偏一边
	if seen["control"] == 0 || seen["treatment_a"] == 0 {
		t.Fatalf("expected both variants to be assigned, got %+v", seen)
	}
}

func TestAssignFlagVariant_RolloutBoundaries(t *testing.T) {
	if v := AssignFlagVariant(flagOf("exp", 0, 500, 500), "u:1"); v != "" {
		t.Fatalf("rollout=0 should always miss, got %q", v)
	}
	if v := AssignFlagVariant(flagOf("exp", 100, 500, 500), "u:1"); v == "" {
		t.Fatal("rollout=100 should never miss")
	}
	// rollout 超界钳到 100
	if v := AssignFlagVariant(flagOf("exp", 150, 500, 500), "u:1"); v == "" {
		t.Fatal("rollout>100 should clamp to 100")
	}
	// rollout=50：恰好一半身份应命中（一致性哈希的确定性使我们能精确断言同一身份的行为）
	flag := flagOf("exp", 50, 500, 500)
	hit, miss := 0, 0
	for i := 0; i < 100; i++ {
		if AssignFlagVariant(flag, string(rune(i))) == "" {
			miss++
		} else {
			hit++
		}
	}
	if hit == 0 || miss == 0 {
		t.Fatalf("rollout=50 should split identities, hit=%d miss=%d", hit, miss)
	}
}

func TestAssignFlagVariant_EdgeCases(t *testing.T) {
	disabled := flagOf("exp", 100, 500, 500)
	disabled.Enabled = false
	if v := AssignFlagVariant(disabled, "u:1"); v != "" {
		t.Fatalf("disabled flag should miss, got %q", v)
	}

	noVariants := flagOf("exp", 100)
	if v := AssignFlagVariant(noVariants, "u:1"); v != "" {
		t.Fatalf("flag without variants should miss, got %q", v)
	}

	// 权重全 0：均分回退
	zeroWeights := flagOf("exp", 100, 0, 0)
	if v := AssignFlagVariant(zeroWeights, "u:1"); v != "control" && v != "treatment_a" {
		t.Fatalf("zero weights should fall back to even split, got %q", v)
	}

	// 负权重被跳过
	negWeights := flagOf("exp", 100, -1, 1000)
	if v := AssignFlagVariant(negWeights, "u:1"); v != "treatment_a" {
		t.Fatalf("negative weight should be skipped, got %q", v)
	}

	if v := AssignFlagVariant(flagOf("exp", 100, 500, 500), ""); v != "" {
		t.Fatalf("empty distinctID should miss, got %q", v)
	}
}

func TestDistinctIDOf(t *testing.T) {
	if got := distinctIDOf(7, "device"); got != "u:7" {
		t.Fatalf("user takes precedence: got %q", got)
	}
	if got := distinctIDOf(0, "device-abc"); got != "d:device-abc" {
		t.Fatalf("device fallback: got %q", got)
	}
	if got := distinctIDOf(0, ""); got != "" {
		t.Fatalf("both empty: got %q", got)
	}
}
