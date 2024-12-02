// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-uba/app/core/service/internal/data/ent/application"
	"kratos-uba/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApplicationDelete is the builder for deleting a Application entity.
type ApplicationDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// Where appends a list predicates to the ApplicationDelete builder.
func (ad *ApplicationDelete) Where(ps ...predicate.Application) *ApplicationDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ApplicationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ApplicationDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ApplicationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(application.Table, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUint32))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// ApplicationDeleteOne is the builder for deleting a single Application entity.
type ApplicationDeleteOne struct {
	ad *ApplicationDelete
}

// Where appends a list predicates to the ApplicationDelete builder.
func (ado *ApplicationDeleteOne) Where(ps ...predicate.Application) *ApplicationDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *ApplicationDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{application.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ApplicationDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
