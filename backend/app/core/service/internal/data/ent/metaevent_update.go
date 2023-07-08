// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/metaevent"
	"kratos-uba/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MetaEventUpdate is the builder for updating MetaEvent entities.
type MetaEventUpdate struct {
	config
	hooks     []Hook
	mutation  *MetaEventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MetaEventUpdate builder.
func (meu *MetaEventUpdate) Where(ps ...predicate.MetaEvent) *MetaEventUpdate {
	meu.mutation.Where(ps...)
	return meu
}

// SetUpdateTime sets the "update_time" field.
func (meu *MetaEventUpdate) SetUpdateTime(i int64) *MetaEventUpdate {
	meu.mutation.ResetUpdateTime()
	meu.mutation.SetUpdateTime(i)
	return meu
}

// AddUpdateTime adds i to the "update_time" field.
func (meu *MetaEventUpdate) AddUpdateTime(i int64) *MetaEventUpdate {
	meu.mutation.AddUpdateTime(i)
	return meu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (meu *MetaEventUpdate) ClearUpdateTime() *MetaEventUpdate {
	meu.mutation.ClearUpdateTime()
	return meu
}

// SetDeleteTime sets the "delete_time" field.
func (meu *MetaEventUpdate) SetDeleteTime(i int64) *MetaEventUpdate {
	meu.mutation.ResetDeleteTime()
	meu.mutation.SetDeleteTime(i)
	return meu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (meu *MetaEventUpdate) SetNillableDeleteTime(i *int64) *MetaEventUpdate {
	if i != nil {
		meu.SetDeleteTime(*i)
	}
	return meu
}

// AddDeleteTime adds i to the "delete_time" field.
func (meu *MetaEventUpdate) AddDeleteTime(i int64) *MetaEventUpdate {
	meu.mutation.AddDeleteTime(i)
	return meu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (meu *MetaEventUpdate) ClearDeleteTime() *MetaEventUpdate {
	meu.mutation.ClearDeleteTime()
	return meu
}

// SetEventName sets the "event_name" field.
func (meu *MetaEventUpdate) SetEventName(s string) *MetaEventUpdate {
	meu.mutation.SetEventName(s)
	return meu
}

// SetNillableEventName sets the "event_name" field if the given value is not nil.
func (meu *MetaEventUpdate) SetNillableEventName(s *string) *MetaEventUpdate {
	if s != nil {
		meu.SetEventName(*s)
	}
	return meu
}

// ClearEventName clears the value of the "event_name" field.
func (meu *MetaEventUpdate) ClearEventName() *MetaEventUpdate {
	meu.mutation.ClearEventName()
	return meu
}

// SetShowName sets the "show_name" field.
func (meu *MetaEventUpdate) SetShowName(s string) *MetaEventUpdate {
	meu.mutation.SetShowName(s)
	return meu
}

// SetNillableShowName sets the "show_name" field if the given value is not nil.
func (meu *MetaEventUpdate) SetNillableShowName(s *string) *MetaEventUpdate {
	if s != nil {
		meu.SetShowName(*s)
	}
	return meu
}

// ClearShowName clears the value of the "show_name" field.
func (meu *MetaEventUpdate) ClearShowName() *MetaEventUpdate {
	meu.mutation.ClearShowName()
	return meu
}

// SetAppID sets the "app_id" field.
func (meu *MetaEventUpdate) SetAppID(u uint32) *MetaEventUpdate {
	meu.mutation.ResetAppID()
	meu.mutation.SetAppID(u)
	return meu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (meu *MetaEventUpdate) SetNillableAppID(u *uint32) *MetaEventUpdate {
	if u != nil {
		meu.SetAppID(*u)
	}
	return meu
}

// AddAppID adds u to the "app_id" field.
func (meu *MetaEventUpdate) AddAppID(u int32) *MetaEventUpdate {
	meu.mutation.AddAppID(u)
	return meu
}

// ClearAppID clears the value of the "app_id" field.
func (meu *MetaEventUpdate) ClearAppID() *MetaEventUpdate {
	meu.mutation.ClearAppID()
	return meu
}

// SetYesterdayCount sets the "yesterday_count" field.
func (meu *MetaEventUpdate) SetYesterdayCount(u uint32) *MetaEventUpdate {
	meu.mutation.ResetYesterdayCount()
	meu.mutation.SetYesterdayCount(u)
	return meu
}

// SetNillableYesterdayCount sets the "yesterday_count" field if the given value is not nil.
func (meu *MetaEventUpdate) SetNillableYesterdayCount(u *uint32) *MetaEventUpdate {
	if u != nil {
		meu.SetYesterdayCount(*u)
	}
	return meu
}

// AddYesterdayCount adds u to the "yesterday_count" field.
func (meu *MetaEventUpdate) AddYesterdayCount(u int32) *MetaEventUpdate {
	meu.mutation.AddYesterdayCount(u)
	return meu
}

// ClearYesterdayCount clears the value of the "yesterday_count" field.
func (meu *MetaEventUpdate) ClearYesterdayCount() *MetaEventUpdate {
	meu.mutation.ClearYesterdayCount()
	return meu
}

// Mutation returns the MetaEventMutation object of the builder.
func (meu *MetaEventUpdate) Mutation() *MetaEventMutation {
	return meu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (meu *MetaEventUpdate) Save(ctx context.Context) (int, error) {
	meu.defaults()
	return withHooks[int, MetaEventMutation](ctx, meu.sqlSave, meu.mutation, meu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meu *MetaEventUpdate) SaveX(ctx context.Context) int {
	affected, err := meu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (meu *MetaEventUpdate) Exec(ctx context.Context) error {
	_, err := meu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meu *MetaEventUpdate) ExecX(ctx context.Context) {
	if err := meu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (meu *MetaEventUpdate) defaults() {
	if _, ok := meu.mutation.UpdateTime(); !ok && !meu.mutation.UpdateTimeCleared() {
		v := metaevent.UpdateDefaultUpdateTime()
		meu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meu *MetaEventUpdate) check() error {
	if v, ok := meu.mutation.EventName(); ok {
		if err := metaevent.EventNameValidator(v); err != nil {
			return &ValidationError{Name: "event_name", err: fmt.Errorf(`ent: validator failed for field "MetaEvent.event_name": %w`, err)}
		}
	}
	if v, ok := meu.mutation.ShowName(); ok {
		if err := metaevent.ShowNameValidator(v); err != nil {
			return &ValidationError{Name: "show_name", err: fmt.Errorf(`ent: validator failed for field "MetaEvent.show_name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (meu *MetaEventUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MetaEventUpdate {
	meu.modifiers = append(meu.modifiers, modifiers...)
	return meu
}

func (meu *MetaEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := meu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(metaevent.Table, metaevent.Columns, sqlgraph.NewFieldSpec(metaevent.FieldID, field.TypeUint32))
	if ps := meu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if meu.mutation.CreateTimeCleared() {
		_spec.ClearField(metaevent.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := meu.mutation.UpdateTime(); ok {
		_spec.SetField(metaevent.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := meu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(metaevent.FieldUpdateTime, field.TypeInt64, value)
	}
	if meu.mutation.UpdateTimeCleared() {
		_spec.ClearField(metaevent.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := meu.mutation.DeleteTime(); ok {
		_spec.SetField(metaevent.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := meu.mutation.AddedDeleteTime(); ok {
		_spec.AddField(metaevent.FieldDeleteTime, field.TypeInt64, value)
	}
	if meu.mutation.DeleteTimeCleared() {
		_spec.ClearField(metaevent.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := meu.mutation.EventName(); ok {
		_spec.SetField(metaevent.FieldEventName, field.TypeString, value)
	}
	if meu.mutation.EventNameCleared() {
		_spec.ClearField(metaevent.FieldEventName, field.TypeString)
	}
	if value, ok := meu.mutation.ShowName(); ok {
		_spec.SetField(metaevent.FieldShowName, field.TypeString, value)
	}
	if meu.mutation.ShowNameCleared() {
		_spec.ClearField(metaevent.FieldShowName, field.TypeString)
	}
	if value, ok := meu.mutation.AppID(); ok {
		_spec.SetField(metaevent.FieldAppID, field.TypeUint32, value)
	}
	if value, ok := meu.mutation.AddedAppID(); ok {
		_spec.AddField(metaevent.FieldAppID, field.TypeUint32, value)
	}
	if meu.mutation.AppIDCleared() {
		_spec.ClearField(metaevent.FieldAppID, field.TypeUint32)
	}
	if value, ok := meu.mutation.YesterdayCount(); ok {
		_spec.SetField(metaevent.FieldYesterdayCount, field.TypeUint32, value)
	}
	if value, ok := meu.mutation.AddedYesterdayCount(); ok {
		_spec.AddField(metaevent.FieldYesterdayCount, field.TypeUint32, value)
	}
	if meu.mutation.YesterdayCountCleared() {
		_spec.ClearField(metaevent.FieldYesterdayCount, field.TypeUint32)
	}
	_spec.AddModifiers(meu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, meu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metaevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	meu.mutation.done = true
	return n, nil
}

// MetaEventUpdateOne is the builder for updating a single MetaEvent entity.
type MetaEventUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MetaEventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (meuo *MetaEventUpdateOne) SetUpdateTime(i int64) *MetaEventUpdateOne {
	meuo.mutation.ResetUpdateTime()
	meuo.mutation.SetUpdateTime(i)
	return meuo
}

// AddUpdateTime adds i to the "update_time" field.
func (meuo *MetaEventUpdateOne) AddUpdateTime(i int64) *MetaEventUpdateOne {
	meuo.mutation.AddUpdateTime(i)
	return meuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (meuo *MetaEventUpdateOne) ClearUpdateTime() *MetaEventUpdateOne {
	meuo.mutation.ClearUpdateTime()
	return meuo
}

// SetDeleteTime sets the "delete_time" field.
func (meuo *MetaEventUpdateOne) SetDeleteTime(i int64) *MetaEventUpdateOne {
	meuo.mutation.ResetDeleteTime()
	meuo.mutation.SetDeleteTime(i)
	return meuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (meuo *MetaEventUpdateOne) SetNillableDeleteTime(i *int64) *MetaEventUpdateOne {
	if i != nil {
		meuo.SetDeleteTime(*i)
	}
	return meuo
}

// AddDeleteTime adds i to the "delete_time" field.
func (meuo *MetaEventUpdateOne) AddDeleteTime(i int64) *MetaEventUpdateOne {
	meuo.mutation.AddDeleteTime(i)
	return meuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (meuo *MetaEventUpdateOne) ClearDeleteTime() *MetaEventUpdateOne {
	meuo.mutation.ClearDeleteTime()
	return meuo
}

// SetEventName sets the "event_name" field.
func (meuo *MetaEventUpdateOne) SetEventName(s string) *MetaEventUpdateOne {
	meuo.mutation.SetEventName(s)
	return meuo
}

// SetNillableEventName sets the "event_name" field if the given value is not nil.
func (meuo *MetaEventUpdateOne) SetNillableEventName(s *string) *MetaEventUpdateOne {
	if s != nil {
		meuo.SetEventName(*s)
	}
	return meuo
}

// ClearEventName clears the value of the "event_name" field.
func (meuo *MetaEventUpdateOne) ClearEventName() *MetaEventUpdateOne {
	meuo.mutation.ClearEventName()
	return meuo
}

// SetShowName sets the "show_name" field.
func (meuo *MetaEventUpdateOne) SetShowName(s string) *MetaEventUpdateOne {
	meuo.mutation.SetShowName(s)
	return meuo
}

// SetNillableShowName sets the "show_name" field if the given value is not nil.
func (meuo *MetaEventUpdateOne) SetNillableShowName(s *string) *MetaEventUpdateOne {
	if s != nil {
		meuo.SetShowName(*s)
	}
	return meuo
}

// ClearShowName clears the value of the "show_name" field.
func (meuo *MetaEventUpdateOne) ClearShowName() *MetaEventUpdateOne {
	meuo.mutation.ClearShowName()
	return meuo
}

// SetAppID sets the "app_id" field.
func (meuo *MetaEventUpdateOne) SetAppID(u uint32) *MetaEventUpdateOne {
	meuo.mutation.ResetAppID()
	meuo.mutation.SetAppID(u)
	return meuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (meuo *MetaEventUpdateOne) SetNillableAppID(u *uint32) *MetaEventUpdateOne {
	if u != nil {
		meuo.SetAppID(*u)
	}
	return meuo
}

// AddAppID adds u to the "app_id" field.
func (meuo *MetaEventUpdateOne) AddAppID(u int32) *MetaEventUpdateOne {
	meuo.mutation.AddAppID(u)
	return meuo
}

// ClearAppID clears the value of the "app_id" field.
func (meuo *MetaEventUpdateOne) ClearAppID() *MetaEventUpdateOne {
	meuo.mutation.ClearAppID()
	return meuo
}

// SetYesterdayCount sets the "yesterday_count" field.
func (meuo *MetaEventUpdateOne) SetYesterdayCount(u uint32) *MetaEventUpdateOne {
	meuo.mutation.ResetYesterdayCount()
	meuo.mutation.SetYesterdayCount(u)
	return meuo
}

// SetNillableYesterdayCount sets the "yesterday_count" field if the given value is not nil.
func (meuo *MetaEventUpdateOne) SetNillableYesterdayCount(u *uint32) *MetaEventUpdateOne {
	if u != nil {
		meuo.SetYesterdayCount(*u)
	}
	return meuo
}

// AddYesterdayCount adds u to the "yesterday_count" field.
func (meuo *MetaEventUpdateOne) AddYesterdayCount(u int32) *MetaEventUpdateOne {
	meuo.mutation.AddYesterdayCount(u)
	return meuo
}

// ClearYesterdayCount clears the value of the "yesterday_count" field.
func (meuo *MetaEventUpdateOne) ClearYesterdayCount() *MetaEventUpdateOne {
	meuo.mutation.ClearYesterdayCount()
	return meuo
}

// Mutation returns the MetaEventMutation object of the builder.
func (meuo *MetaEventUpdateOne) Mutation() *MetaEventMutation {
	return meuo.mutation
}

// Where appends a list predicates to the MetaEventUpdate builder.
func (meuo *MetaEventUpdateOne) Where(ps ...predicate.MetaEvent) *MetaEventUpdateOne {
	meuo.mutation.Where(ps...)
	return meuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (meuo *MetaEventUpdateOne) Select(field string, fields ...string) *MetaEventUpdateOne {
	meuo.fields = append([]string{field}, fields...)
	return meuo
}

// Save executes the query and returns the updated MetaEvent entity.
func (meuo *MetaEventUpdateOne) Save(ctx context.Context) (*MetaEvent, error) {
	meuo.defaults()
	return withHooks[*MetaEvent, MetaEventMutation](ctx, meuo.sqlSave, meuo.mutation, meuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meuo *MetaEventUpdateOne) SaveX(ctx context.Context) *MetaEvent {
	node, err := meuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (meuo *MetaEventUpdateOne) Exec(ctx context.Context) error {
	_, err := meuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meuo *MetaEventUpdateOne) ExecX(ctx context.Context) {
	if err := meuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (meuo *MetaEventUpdateOne) defaults() {
	if _, ok := meuo.mutation.UpdateTime(); !ok && !meuo.mutation.UpdateTimeCleared() {
		v := metaevent.UpdateDefaultUpdateTime()
		meuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meuo *MetaEventUpdateOne) check() error {
	if v, ok := meuo.mutation.EventName(); ok {
		if err := metaevent.EventNameValidator(v); err != nil {
			return &ValidationError{Name: "event_name", err: fmt.Errorf(`ent: validator failed for field "MetaEvent.event_name": %w`, err)}
		}
	}
	if v, ok := meuo.mutation.ShowName(); ok {
		if err := metaevent.ShowNameValidator(v); err != nil {
			return &ValidationError{Name: "show_name", err: fmt.Errorf(`ent: validator failed for field "MetaEvent.show_name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (meuo *MetaEventUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MetaEventUpdateOne {
	meuo.modifiers = append(meuo.modifiers, modifiers...)
	return meuo
}

func (meuo *MetaEventUpdateOne) sqlSave(ctx context.Context) (_node *MetaEvent, err error) {
	if err := meuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(metaevent.Table, metaevent.Columns, sqlgraph.NewFieldSpec(metaevent.FieldID, field.TypeUint32))
	id, ok := meuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MetaEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := meuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metaevent.FieldID)
		for _, f := range fields {
			if !metaevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metaevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := meuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if meuo.mutation.CreateTimeCleared() {
		_spec.ClearField(metaevent.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := meuo.mutation.UpdateTime(); ok {
		_spec.SetField(metaevent.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := meuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(metaevent.FieldUpdateTime, field.TypeInt64, value)
	}
	if meuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(metaevent.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := meuo.mutation.DeleteTime(); ok {
		_spec.SetField(metaevent.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := meuo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(metaevent.FieldDeleteTime, field.TypeInt64, value)
	}
	if meuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(metaevent.FieldDeleteTime, field.TypeInt64)
	}
	if value, ok := meuo.mutation.EventName(); ok {
		_spec.SetField(metaevent.FieldEventName, field.TypeString, value)
	}
	if meuo.mutation.EventNameCleared() {
		_spec.ClearField(metaevent.FieldEventName, field.TypeString)
	}
	if value, ok := meuo.mutation.ShowName(); ok {
		_spec.SetField(metaevent.FieldShowName, field.TypeString, value)
	}
	if meuo.mutation.ShowNameCleared() {
		_spec.ClearField(metaevent.FieldShowName, field.TypeString)
	}
	if value, ok := meuo.mutation.AppID(); ok {
		_spec.SetField(metaevent.FieldAppID, field.TypeUint32, value)
	}
	if value, ok := meuo.mutation.AddedAppID(); ok {
		_spec.AddField(metaevent.FieldAppID, field.TypeUint32, value)
	}
	if meuo.mutation.AppIDCleared() {
		_spec.ClearField(metaevent.FieldAppID, field.TypeUint32)
	}
	if value, ok := meuo.mutation.YesterdayCount(); ok {
		_spec.SetField(metaevent.FieldYesterdayCount, field.TypeUint32, value)
	}
	if value, ok := meuo.mutation.AddedYesterdayCount(); ok {
		_spec.AddField(metaevent.FieldYesterdayCount, field.TypeUint32, value)
	}
	if meuo.mutation.YesterdayCountCleared() {
		_spec.ClearField(metaevent.FieldYesterdayCount, field.TypeUint32)
	}
	_spec.AddModifiers(meuo.modifiers...)
	_node = &MetaEvent{config: meuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, meuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metaevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	meuo.mutation.done = true
	return _node, nil
}
