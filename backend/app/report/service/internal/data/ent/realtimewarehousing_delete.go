// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-bi/app/report/service/internal/data/ent/predicate"
	"kratos-bi/app/report/service/internal/data/ent/realtimewarehousing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RealtimeWarehousingDelete is the builder for deleting a RealtimeWarehousing entity.
type RealtimeWarehousingDelete struct {
	config
	hooks    []Hook
	mutation *RealtimeWarehousingMutation
}

// Where appends a list predicates to the RealtimeWarehousingDelete builder.
func (rwd *RealtimeWarehousingDelete) Where(ps ...predicate.RealtimeWarehousing) *RealtimeWarehousingDelete {
	rwd.mutation.Where(ps...)
	return rwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rwd *RealtimeWarehousingDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RealtimeWarehousingMutation](ctx, rwd.sqlExec, rwd.mutation, rwd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rwd *RealtimeWarehousingDelete) ExecX(ctx context.Context) int {
	n, err := rwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rwd *RealtimeWarehousingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(realtimewarehousing.Table, sqlgraph.NewFieldSpec(realtimewarehousing.FieldID, field.TypeUint32))
	if ps := rwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rwd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rwd.mutation.done = true
	return affected, err
}

// RealtimeWarehousingDeleteOne is the builder for deleting a single RealtimeWarehousing entity.
type RealtimeWarehousingDeleteOne struct {
	rwd *RealtimeWarehousingDelete
}

// Where appends a list predicates to the RealtimeWarehousingDelete builder.
func (rwdo *RealtimeWarehousingDeleteOne) Where(ps ...predicate.RealtimeWarehousing) *RealtimeWarehousingDeleteOne {
	rwdo.rwd.mutation.Where(ps...)
	return rwdo
}

// Exec executes the deletion query.
func (rwdo *RealtimeWarehousingDeleteOne) Exec(ctx context.Context) error {
	n, err := rwdo.rwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{realtimewarehousing.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rwdo *RealtimeWarehousingDeleteOne) ExecX(ctx context.Context) {
	if err := rwdo.Exec(ctx); err != nil {
		panic(err)
	}
}