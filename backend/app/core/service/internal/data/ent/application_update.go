// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-bi/app/core/service/internal/data/ent/application"
	"kratos-bi/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApplicationUpdate is the builder for updating Application entities.
type ApplicationUpdate struct {
	config
	hooks     []Hook
	mutation  *ApplicationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (au *ApplicationUpdate) Where(ps ...predicate.Application) *ApplicationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdateTime sets the "update_time" field.
func (au *ApplicationUpdate) SetUpdateTime(i int64) *ApplicationUpdate {
	au.mutation.ResetUpdateTime()
	au.mutation.SetUpdateTime(i)
	return au
}

// AddUpdateTime adds i to the "update_time" field.
func (au *ApplicationUpdate) AddUpdateTime(i int64) *ApplicationUpdate {
	au.mutation.AddUpdateTime(i)
	return au
}

// ClearUpdateTime clears the value of the "update_time" field.
func (au *ApplicationUpdate) ClearUpdateTime() *ApplicationUpdate {
	au.mutation.ClearUpdateTime()
	return au
}

// SetDeleteTime sets the "delete_time" field.
func (au *ApplicationUpdate) SetDeleteTime(i int64) *ApplicationUpdate {
	au.mutation.ResetDeleteTime()
	au.mutation.SetDeleteTime(i)
	return au
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableDeleteTime(i *int64) *ApplicationUpdate {
	if i != nil {
		au.SetDeleteTime(*i)
	}
	return au
}

// AddDeleteTime adds i to the "delete_time" field.
func (au *ApplicationUpdate) AddDeleteTime(i int64) *ApplicationUpdate {
	au.mutation.AddDeleteTime(i)
	return au
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (au *ApplicationUpdate) ClearDeleteTime() *ApplicationUpdate {
	au.mutation.ClearDeleteTime()
	return au
}

// SetName sets the "name" field.
func (au *ApplicationUpdate) SetName(s string) *ApplicationUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableName(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *ApplicationUpdate) ClearName() *ApplicationUpdate {
	au.mutation.ClearName()
	return au
}

// SetStatus sets the "status" field.
func (au *ApplicationUpdate) SetStatus(s string) *ApplicationUpdate {
	au.mutation.SetStatus(s)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableStatus(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetStatus(*s)
	}
	return au
}

// ClearStatus clears the value of the "status" field.
func (au *ApplicationUpdate) ClearStatus() *ApplicationUpdate {
	au.mutation.ClearStatus()
	return au
}

// SetAppID sets the "app_id" field.
func (au *ApplicationUpdate) SetAppID(s string) *ApplicationUpdate {
	au.mutation.SetAppID(s)
	return au
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableAppID(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetAppID(*s)
	}
	return au
}

// ClearAppID clears the value of the "app_id" field.
func (au *ApplicationUpdate) ClearAppID() *ApplicationUpdate {
	au.mutation.ClearAppID()
	return au
}

// SetAppKey sets the "app_key" field.
func (au *ApplicationUpdate) SetAppKey(s string) *ApplicationUpdate {
	au.mutation.SetAppKey(s)
	return au
}

// SetNillableAppKey sets the "app_key" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableAppKey(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetAppKey(*s)
	}
	return au
}

// ClearAppKey clears the value of the "app_key" field.
func (au *ApplicationUpdate) ClearAppKey() *ApplicationUpdate {
	au.mutation.ClearAppKey()
	return au
}

// SetRemark sets the "remark" field.
func (au *ApplicationUpdate) SetRemark(s string) *ApplicationUpdate {
	au.mutation.SetRemark(s)
	return au
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableRemark(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetRemark(*s)
	}
	return au
}

// ClearRemark clears the value of the "remark" field.
func (au *ApplicationUpdate) ClearRemark() *ApplicationUpdate {
	au.mutation.ClearRemark()
	return au
}

// SetCreatorID sets the "creator_id" field.
func (au *ApplicationUpdate) SetCreatorID(u uint32) *ApplicationUpdate {
	au.mutation.ResetCreatorID()
	au.mutation.SetCreatorID(u)
	return au
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableCreatorID(u *uint32) *ApplicationUpdate {
	if u != nil {
		au.SetCreatorID(*u)
	}
	return au
}

// AddCreatorID adds u to the "creator_id" field.
func (au *ApplicationUpdate) AddCreatorID(u int32) *ApplicationUpdate {
	au.mutation.AddCreatorID(u)
	return au
}

// ClearCreatorID clears the value of the "creator_id" field.
func (au *ApplicationUpdate) ClearCreatorID() *ApplicationUpdate {
	au.mutation.ClearCreatorID()
	return au
}

// Mutation returns the ApplicationMutation object of the builder.
func (au *ApplicationUpdate) Mutation() *ApplicationMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApplicationUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks[int, ApplicationMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApplicationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApplicationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ApplicationUpdate) defaults() {
	if _, ok := au.mutation.UpdateTime(); !ok && !au.mutation.UpdateTimeCleared() {
		v := application.UpdateDefaultUpdateTime()
		au.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ApplicationUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := application.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Application.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *ApplicationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *ApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUint32))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if au.mutation.CreateTimeCleared() {
		_spec.ClearField(application.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := au.mutation.UpdateTime(); ok {
		_spec.SetField(application.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUpdateTime(); ok {
		_spec.AddField(application.FieldUpdateTime, field.TypeInt64, value)
	}
	if au.mutation.UpdateTimeCleared() {
		_spec.ClearField(application.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := au.mutation.DeleteTime(); ok {
		_spec.SetField(application.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedDeleteTime(); ok {
		_spec.AddField(application.FieldDeleteTime, field.TypeInt64, value)
	}
	if au.mutation.DeleteTimeCleared() {
		_spec.ClearField(application.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(application.FieldName, field.TypeString, value)
	}
	if au.mutation.NameCleared() {
		_spec.ClearField(application.FieldName, field.TypeString)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(application.FieldStatus, field.TypeString, value)
	}
	if au.mutation.StatusCleared() {
		_spec.ClearField(application.FieldStatus, field.TypeString)
	}
	if value, ok := au.mutation.AppID(); ok {
		_spec.SetField(application.FieldAppID, field.TypeString, value)
	}
	if au.mutation.AppIDCleared() {
		_spec.ClearField(application.FieldAppID, field.TypeString)
	}
	if value, ok := au.mutation.AppKey(); ok {
		_spec.SetField(application.FieldAppKey, field.TypeString, value)
	}
	if au.mutation.AppKeyCleared() {
		_spec.ClearField(application.FieldAppKey, field.TypeString)
	}
	if value, ok := au.mutation.Remark(); ok {
		_spec.SetField(application.FieldRemark, field.TypeString, value)
	}
	if au.mutation.RemarkCleared() {
		_spec.ClearField(application.FieldRemark, field.TypeString)
	}
	if value, ok := au.mutation.CreatorID(); ok {
		_spec.SetField(application.FieldCreatorID, field.TypeUint32, value)
	}
	if value, ok := au.mutation.AddedCreatorID(); ok {
		_spec.AddField(application.FieldCreatorID, field.TypeUint32, value)
	}
	if au.mutation.CreatorIDCleared() {
		_spec.ClearField(application.FieldCreatorID, field.TypeUint32)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApplicationUpdateOne is the builder for updating a single Application entity.
type ApplicationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApplicationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (auo *ApplicationUpdateOne) SetUpdateTime(i int64) *ApplicationUpdateOne {
	auo.mutation.ResetUpdateTime()
	auo.mutation.SetUpdateTime(i)
	return auo
}

// AddUpdateTime adds i to the "update_time" field.
func (auo *ApplicationUpdateOne) AddUpdateTime(i int64) *ApplicationUpdateOne {
	auo.mutation.AddUpdateTime(i)
	return auo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (auo *ApplicationUpdateOne) ClearUpdateTime() *ApplicationUpdateOne {
	auo.mutation.ClearUpdateTime()
	return auo
}

// SetDeleteTime sets the "delete_time" field.
func (auo *ApplicationUpdateOne) SetDeleteTime(i int64) *ApplicationUpdateOne {
	auo.mutation.ResetDeleteTime()
	auo.mutation.SetDeleteTime(i)
	return auo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableDeleteTime(i *int64) *ApplicationUpdateOne {
	if i != nil {
		auo.SetDeleteTime(*i)
	}
	return auo
}

// AddDeleteTime adds i to the "delete_time" field.
func (auo *ApplicationUpdateOne) AddDeleteTime(i int64) *ApplicationUpdateOne {
	auo.mutation.AddDeleteTime(i)
	return auo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (auo *ApplicationUpdateOne) ClearDeleteTime() *ApplicationUpdateOne {
	auo.mutation.ClearDeleteTime()
	return auo
}

// SetName sets the "name" field.
func (auo *ApplicationUpdateOne) SetName(s string) *ApplicationUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableName(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *ApplicationUpdateOne) ClearName() *ApplicationUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetStatus sets the "status" field.
func (auo *ApplicationUpdateOne) SetStatus(s string) *ApplicationUpdateOne {
	auo.mutation.SetStatus(s)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableStatus(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetStatus(*s)
	}
	return auo
}

// ClearStatus clears the value of the "status" field.
func (auo *ApplicationUpdateOne) ClearStatus() *ApplicationUpdateOne {
	auo.mutation.ClearStatus()
	return auo
}

// SetAppID sets the "app_id" field.
func (auo *ApplicationUpdateOne) SetAppID(s string) *ApplicationUpdateOne {
	auo.mutation.SetAppID(s)
	return auo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableAppID(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetAppID(*s)
	}
	return auo
}

// ClearAppID clears the value of the "app_id" field.
func (auo *ApplicationUpdateOne) ClearAppID() *ApplicationUpdateOne {
	auo.mutation.ClearAppID()
	return auo
}

// SetAppKey sets the "app_key" field.
func (auo *ApplicationUpdateOne) SetAppKey(s string) *ApplicationUpdateOne {
	auo.mutation.SetAppKey(s)
	return auo
}

// SetNillableAppKey sets the "app_key" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableAppKey(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetAppKey(*s)
	}
	return auo
}

// ClearAppKey clears the value of the "app_key" field.
func (auo *ApplicationUpdateOne) ClearAppKey() *ApplicationUpdateOne {
	auo.mutation.ClearAppKey()
	return auo
}

// SetRemark sets the "remark" field.
func (auo *ApplicationUpdateOne) SetRemark(s string) *ApplicationUpdateOne {
	auo.mutation.SetRemark(s)
	return auo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableRemark(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetRemark(*s)
	}
	return auo
}

// ClearRemark clears the value of the "remark" field.
func (auo *ApplicationUpdateOne) ClearRemark() *ApplicationUpdateOne {
	auo.mutation.ClearRemark()
	return auo
}

// SetCreatorID sets the "creator_id" field.
func (auo *ApplicationUpdateOne) SetCreatorID(u uint32) *ApplicationUpdateOne {
	auo.mutation.ResetCreatorID()
	auo.mutation.SetCreatorID(u)
	return auo
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableCreatorID(u *uint32) *ApplicationUpdateOne {
	if u != nil {
		auo.SetCreatorID(*u)
	}
	return auo
}

// AddCreatorID adds u to the "creator_id" field.
func (auo *ApplicationUpdateOne) AddCreatorID(u int32) *ApplicationUpdateOne {
	auo.mutation.AddCreatorID(u)
	return auo
}

// ClearCreatorID clears the value of the "creator_id" field.
func (auo *ApplicationUpdateOne) ClearCreatorID() *ApplicationUpdateOne {
	auo.mutation.ClearCreatorID()
	return auo
}

// Mutation returns the ApplicationMutation object of the builder.
func (auo *ApplicationUpdateOne) Mutation() *ApplicationMutation {
	return auo.mutation
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (auo *ApplicationUpdateOne) Where(ps ...predicate.Application) *ApplicationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApplicationUpdateOne) Select(field string, fields ...string) *ApplicationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Application entity.
func (auo *ApplicationUpdateOne) Save(ctx context.Context) (*Application, error) {
	auo.defaults()
	return withHooks[*Application, ApplicationMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApplicationUpdateOne) SaveX(ctx context.Context) *Application {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ApplicationUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateTime(); !ok && !auo.mutation.UpdateTimeCleared() {
		v := application.UpdateDefaultUpdateTime()
		auo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ApplicationUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := application.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Application.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *ApplicationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *ApplicationUpdateOne) sqlSave(ctx context.Context) (_node *Application, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUint32))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Application.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, application.FieldID)
		for _, f := range fields {
			if !application.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != application.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auo.mutation.CreateTimeCleared() {
		_spec.ClearField(application.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := auo.mutation.UpdateTime(); ok {
		_spec.SetField(application.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(application.FieldUpdateTime, field.TypeInt64, value)
	}
	if auo.mutation.UpdateTimeCleared() {
		_spec.ClearField(application.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := auo.mutation.DeleteTime(); ok {
		_spec.SetField(application.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(application.FieldDeleteTime, field.TypeInt64, value)
	}
	if auo.mutation.DeleteTimeCleared() {
		_spec.ClearField(application.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(application.FieldName, field.TypeString, value)
	}
	if auo.mutation.NameCleared() {
		_spec.ClearField(application.FieldName, field.TypeString)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(application.FieldStatus, field.TypeString, value)
	}
	if auo.mutation.StatusCleared() {
		_spec.ClearField(application.FieldStatus, field.TypeString)
	}
	if value, ok := auo.mutation.AppID(); ok {
		_spec.SetField(application.FieldAppID, field.TypeString, value)
	}
	if auo.mutation.AppIDCleared() {
		_spec.ClearField(application.FieldAppID, field.TypeString)
	}
	if value, ok := auo.mutation.AppKey(); ok {
		_spec.SetField(application.FieldAppKey, field.TypeString, value)
	}
	if auo.mutation.AppKeyCleared() {
		_spec.ClearField(application.FieldAppKey, field.TypeString)
	}
	if value, ok := auo.mutation.Remark(); ok {
		_spec.SetField(application.FieldRemark, field.TypeString, value)
	}
	if auo.mutation.RemarkCleared() {
		_spec.ClearField(application.FieldRemark, field.TypeString)
	}
	if value, ok := auo.mutation.CreatorID(); ok {
		_spec.SetField(application.FieldCreatorID, field.TypeUint32, value)
	}
	if value, ok := auo.mutation.AddedCreatorID(); ok {
		_spec.AddField(application.FieldCreatorID, field.TypeUint32, value)
	}
	if auo.mutation.CreatorIDCleared() {
		_spec.ClearField(application.FieldCreatorID, field.TypeUint32)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Application{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}