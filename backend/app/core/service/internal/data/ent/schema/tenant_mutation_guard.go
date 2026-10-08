package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"

	"github.com/tx7do/go-crud/viewer"
)

// TenantMutationGuardPolicy 补齐 go-crud TenantPrivacy 的缺口：
// TenantPrivacy.EvalMutation 仅覆盖 Create（强制覆盖 tenant_id），
// 对 Update/UpdateOne/Delete/DeleteOne 直接放行——知道 ID 即可跨租户篡改/删除。
//
// 本 rule 对非 Create 的变更注入 tenant_id 过滤（mutation.WhereP）：
//   - 租户上下文：只能作用于本租户行，跨租户变更命中 0 行；
//   - 平台 / 系统上下文：放行（平台管理员全权）。
//
// 评估顺序：追加在各 schema Policy() 链尾部，位于 go-crud TenantPrivacy
// （Query 过滤 + Create 覆盖）之后，不改变既有行为。
type TenantMutationGuardPolicy struct{}

func (TenantMutationGuardPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	// 查询隔离由 go-crud TenantPrivacy 负责
	return nil
}

func (TenantMutationGuardPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m.Op().Is(ent.OpCreate) {
		return nil
	}

	vc, exist := viewer.FromContext(ctx)
	if !exist {
		return fmt.Errorf("security: missing ViewerContext in context")
	}

	// 平台管理 / 系统上下文放行
	if vc.IsPlatformContext() || vc.IsSystemContext() {
		return nil
	}

	tid := vc.TenantID()

	// 通过 entql feature 生成的 WhereP 注入 tenant_id 谓词，
	// 覆盖 Update / UpdateOne / Delete / DeleteOne 全部变更形态。
	if wm, ok := m.(interface {
		WhereP(...func(*sql.Selector))
	}); ok {
		tid := uint64(tid)
		wm.WhereP(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C("tenant_id"), tid))
		})
	}

	return nil
}
