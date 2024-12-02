// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/debugdevice"
	"kratos-uba/app/core/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DebugDeviceUpdate is the builder for updating DebugDevice entities.
type DebugDeviceUpdate struct {
	config
	hooks     []Hook
	mutation  *DebugDeviceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DebugDeviceUpdate builder.
func (ddu *DebugDeviceUpdate) Where(ps ...predicate.DebugDevice) *DebugDeviceUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetUpdateTime sets the "update_time" field.
func (ddu *DebugDeviceUpdate) SetUpdateTime(t time.Time) *DebugDeviceUpdate {
	ddu.mutation.SetUpdateTime(t)
	return ddu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableUpdateTime(t *time.Time) *DebugDeviceUpdate {
	if t != nil {
		ddu.SetUpdateTime(*t)
	}
	return ddu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ddu *DebugDeviceUpdate) ClearUpdateTime() *DebugDeviceUpdate {
	ddu.mutation.ClearUpdateTime()
	return ddu
}

// SetDeleteTime sets the "delete_time" field.
func (ddu *DebugDeviceUpdate) SetDeleteTime(t time.Time) *DebugDeviceUpdate {
	ddu.mutation.SetDeleteTime(t)
	return ddu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableDeleteTime(t *time.Time) *DebugDeviceUpdate {
	if t != nil {
		ddu.SetDeleteTime(*t)
	}
	return ddu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ddu *DebugDeviceUpdate) ClearDeleteTime() *DebugDeviceUpdate {
	ddu.mutation.ClearDeleteTime()
	return ddu
}

// SetDeviceID sets the "device_id" field.
func (ddu *DebugDeviceUpdate) SetDeviceID(s string) *DebugDeviceUpdate {
	ddu.mutation.SetDeviceID(s)
	return ddu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableDeviceID(s *string) *DebugDeviceUpdate {
	if s != nil {
		ddu.SetDeviceID(*s)
	}
	return ddu
}

// ClearDeviceID clears the value of the "device_id" field.
func (ddu *DebugDeviceUpdate) ClearDeviceID() *DebugDeviceUpdate {
	ddu.mutation.ClearDeviceID()
	return ddu
}

// SetAppID sets the "app_id" field.
func (ddu *DebugDeviceUpdate) SetAppID(u uint32) *DebugDeviceUpdate {
	ddu.mutation.ResetAppID()
	ddu.mutation.SetAppID(u)
	return ddu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableAppID(u *uint32) *DebugDeviceUpdate {
	if u != nil {
		ddu.SetAppID(*u)
	}
	return ddu
}

// AddAppID adds u to the "app_id" field.
func (ddu *DebugDeviceUpdate) AddAppID(u int32) *DebugDeviceUpdate {
	ddu.mutation.AddAppID(u)
	return ddu
}

// ClearAppID clears the value of the "app_id" field.
func (ddu *DebugDeviceUpdate) ClearAppID() *DebugDeviceUpdate {
	ddu.mutation.ClearAppID()
	return ddu
}

// SetCreatorID sets the "creator_id" field.
func (ddu *DebugDeviceUpdate) SetCreatorID(u uint32) *DebugDeviceUpdate {
	ddu.mutation.ResetCreatorID()
	ddu.mutation.SetCreatorID(u)
	return ddu
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableCreatorID(u *uint32) *DebugDeviceUpdate {
	if u != nil {
		ddu.SetCreatorID(*u)
	}
	return ddu
}

// AddCreatorID adds u to the "creator_id" field.
func (ddu *DebugDeviceUpdate) AddCreatorID(u int32) *DebugDeviceUpdate {
	ddu.mutation.AddCreatorID(u)
	return ddu
}

// ClearCreatorID clears the value of the "creator_id" field.
func (ddu *DebugDeviceUpdate) ClearCreatorID() *DebugDeviceUpdate {
	ddu.mutation.ClearCreatorID()
	return ddu
}

// SetRemark sets the "remark" field.
func (ddu *DebugDeviceUpdate) SetRemark(s string) *DebugDeviceUpdate {
	ddu.mutation.SetRemark(s)
	return ddu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ddu *DebugDeviceUpdate) SetNillableRemark(s *string) *DebugDeviceUpdate {
	if s != nil {
		ddu.SetRemark(*s)
	}
	return ddu
}

// ClearRemark clears the value of the "remark" field.
func (ddu *DebugDeviceUpdate) ClearRemark() *DebugDeviceUpdate {
	ddu.mutation.ClearRemark()
	return ddu
}

// Mutation returns the DebugDeviceMutation object of the builder.
func (ddu *DebugDeviceUpdate) Mutation() *DebugDeviceMutation {
	return ddu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DebugDeviceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ddu.sqlSave, ddu.mutation, ddu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DebugDeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DebugDeviceUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DebugDeviceUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddu *DebugDeviceUpdate) check() error {
	if v, ok := ddu.mutation.DeviceID(); ok {
		if err := debugdevice.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`ent: validator failed for field "DebugDevice.device_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ddu *DebugDeviceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DebugDeviceUpdate {
	ddu.modifiers = append(ddu.modifiers, modifiers...)
	return ddu
}

func (ddu *DebugDeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ddu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(debugdevice.Table, debugdevice.Columns, sqlgraph.NewFieldSpec(debugdevice.FieldID, field.TypeUint32))
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ddu.mutation.CreateTimeCleared() {
		_spec.ClearField(debugdevice.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ddu.mutation.UpdateTime(); ok {
		_spec.SetField(debugdevice.FieldUpdateTime, field.TypeTime, value)
	}
	if ddu.mutation.UpdateTimeCleared() {
		_spec.ClearField(debugdevice.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ddu.mutation.DeleteTime(); ok {
		_spec.SetField(debugdevice.FieldDeleteTime, field.TypeTime, value)
	}
	if ddu.mutation.DeleteTimeCleared() {
		_spec.ClearField(debugdevice.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := ddu.mutation.DeviceID(); ok {
		_spec.SetField(debugdevice.FieldDeviceID, field.TypeString, value)
	}
	if ddu.mutation.DeviceIDCleared() {
		_spec.ClearField(debugdevice.FieldDeviceID, field.TypeString)
	}
	if value, ok := ddu.mutation.AppID(); ok {
		_spec.SetField(debugdevice.FieldAppID, field.TypeUint32, value)
	}
	if value, ok := ddu.mutation.AddedAppID(); ok {
		_spec.AddField(debugdevice.FieldAppID, field.TypeUint32, value)
	}
	if ddu.mutation.AppIDCleared() {
		_spec.ClearField(debugdevice.FieldAppID, field.TypeUint32)
	}
	if value, ok := ddu.mutation.CreatorID(); ok {
		_spec.SetField(debugdevice.FieldCreatorID, field.TypeUint32, value)
	}
	if value, ok := ddu.mutation.AddedCreatorID(); ok {
		_spec.AddField(debugdevice.FieldCreatorID, field.TypeUint32, value)
	}
	if ddu.mutation.CreatorIDCleared() {
		_spec.ClearField(debugdevice.FieldCreatorID, field.TypeUint32)
	}
	if value, ok := ddu.mutation.Remark(); ok {
		_spec.SetField(debugdevice.FieldRemark, field.TypeString, value)
	}
	if ddu.mutation.RemarkCleared() {
		_spec.ClearField(debugdevice.FieldRemark, field.TypeString)
	}
	_spec.AddModifiers(ddu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{debugdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ddu.mutation.done = true
	return n, nil
}

// DebugDeviceUpdateOne is the builder for updating a single DebugDevice entity.
type DebugDeviceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DebugDeviceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (dduo *DebugDeviceUpdateOne) SetUpdateTime(t time.Time) *DebugDeviceUpdateOne {
	dduo.mutation.SetUpdateTime(t)
	return dduo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableUpdateTime(t *time.Time) *DebugDeviceUpdateOne {
	if t != nil {
		dduo.SetUpdateTime(*t)
	}
	return dduo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (dduo *DebugDeviceUpdateOne) ClearUpdateTime() *DebugDeviceUpdateOne {
	dduo.mutation.ClearUpdateTime()
	return dduo
}

// SetDeleteTime sets the "delete_time" field.
func (dduo *DebugDeviceUpdateOne) SetDeleteTime(t time.Time) *DebugDeviceUpdateOne {
	dduo.mutation.SetDeleteTime(t)
	return dduo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableDeleteTime(t *time.Time) *DebugDeviceUpdateOne {
	if t != nil {
		dduo.SetDeleteTime(*t)
	}
	return dduo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (dduo *DebugDeviceUpdateOne) ClearDeleteTime() *DebugDeviceUpdateOne {
	dduo.mutation.ClearDeleteTime()
	return dduo
}

// SetDeviceID sets the "device_id" field.
func (dduo *DebugDeviceUpdateOne) SetDeviceID(s string) *DebugDeviceUpdateOne {
	dduo.mutation.SetDeviceID(s)
	return dduo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableDeviceID(s *string) *DebugDeviceUpdateOne {
	if s != nil {
		dduo.SetDeviceID(*s)
	}
	return dduo
}

// ClearDeviceID clears the value of the "device_id" field.
func (dduo *DebugDeviceUpdateOne) ClearDeviceID() *DebugDeviceUpdateOne {
	dduo.mutation.ClearDeviceID()
	return dduo
}

// SetAppID sets the "app_id" field.
func (dduo *DebugDeviceUpdateOne) SetAppID(u uint32) *DebugDeviceUpdateOne {
	dduo.mutation.ResetAppID()
	dduo.mutation.SetAppID(u)
	return dduo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableAppID(u *uint32) *DebugDeviceUpdateOne {
	if u != nil {
		dduo.SetAppID(*u)
	}
	return dduo
}

// AddAppID adds u to the "app_id" field.
func (dduo *DebugDeviceUpdateOne) AddAppID(u int32) *DebugDeviceUpdateOne {
	dduo.mutation.AddAppID(u)
	return dduo
}

// ClearAppID clears the value of the "app_id" field.
func (dduo *DebugDeviceUpdateOne) ClearAppID() *DebugDeviceUpdateOne {
	dduo.mutation.ClearAppID()
	return dduo
}

// SetCreatorID sets the "creator_id" field.
func (dduo *DebugDeviceUpdateOne) SetCreatorID(u uint32) *DebugDeviceUpdateOne {
	dduo.mutation.ResetCreatorID()
	dduo.mutation.SetCreatorID(u)
	return dduo
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableCreatorID(u *uint32) *DebugDeviceUpdateOne {
	if u != nil {
		dduo.SetCreatorID(*u)
	}
	return dduo
}

// AddCreatorID adds u to the "creator_id" field.
func (dduo *DebugDeviceUpdateOne) AddCreatorID(u int32) *DebugDeviceUpdateOne {
	dduo.mutation.AddCreatorID(u)
	return dduo
}

// ClearCreatorID clears the value of the "creator_id" field.
func (dduo *DebugDeviceUpdateOne) ClearCreatorID() *DebugDeviceUpdateOne {
	dduo.mutation.ClearCreatorID()
	return dduo
}

// SetRemark sets the "remark" field.
func (dduo *DebugDeviceUpdateOne) SetRemark(s string) *DebugDeviceUpdateOne {
	dduo.mutation.SetRemark(s)
	return dduo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dduo *DebugDeviceUpdateOne) SetNillableRemark(s *string) *DebugDeviceUpdateOne {
	if s != nil {
		dduo.SetRemark(*s)
	}
	return dduo
}

// ClearRemark clears the value of the "remark" field.
func (dduo *DebugDeviceUpdateOne) ClearRemark() *DebugDeviceUpdateOne {
	dduo.mutation.ClearRemark()
	return dduo
}

// Mutation returns the DebugDeviceMutation object of the builder.
func (dduo *DebugDeviceUpdateOne) Mutation() *DebugDeviceMutation {
	return dduo.mutation
}

// Where appends a list predicates to the DebugDeviceUpdate builder.
func (dduo *DebugDeviceUpdateOne) Where(ps ...predicate.DebugDevice) *DebugDeviceUpdateOne {
	dduo.mutation.Where(ps...)
	return dduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DebugDeviceUpdateOne) Select(field string, fields ...string) *DebugDeviceUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DebugDevice entity.
func (dduo *DebugDeviceUpdateOne) Save(ctx context.Context) (*DebugDevice, error) {
	return withHooks(ctx, dduo.sqlSave, dduo.mutation, dduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DebugDeviceUpdateOne) SaveX(ctx context.Context) *DebugDevice {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DebugDeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DebugDeviceUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dduo *DebugDeviceUpdateOne) check() error {
	if v, ok := dduo.mutation.DeviceID(); ok {
		if err := debugdevice.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`ent: validator failed for field "DebugDevice.device_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dduo *DebugDeviceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DebugDeviceUpdateOne {
	dduo.modifiers = append(dduo.modifiers, modifiers...)
	return dduo
}

func (dduo *DebugDeviceUpdateOne) sqlSave(ctx context.Context) (_node *DebugDevice, err error) {
	if err := dduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(debugdevice.Table, debugdevice.Columns, sqlgraph.NewFieldSpec(debugdevice.FieldID, field.TypeUint32))
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DebugDevice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, debugdevice.FieldID)
		for _, f := range fields {
			if !debugdevice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != debugdevice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dduo.mutation.CreateTimeCleared() {
		_spec.ClearField(debugdevice.FieldCreateTime, field.TypeTime)
	}
	if value, ok := dduo.mutation.UpdateTime(); ok {
		_spec.SetField(debugdevice.FieldUpdateTime, field.TypeTime, value)
	}
	if dduo.mutation.UpdateTimeCleared() {
		_spec.ClearField(debugdevice.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := dduo.mutation.DeleteTime(); ok {
		_spec.SetField(debugdevice.FieldDeleteTime, field.TypeTime, value)
	}
	if dduo.mutation.DeleteTimeCleared() {
		_spec.ClearField(debugdevice.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := dduo.mutation.DeviceID(); ok {
		_spec.SetField(debugdevice.FieldDeviceID, field.TypeString, value)
	}
	if dduo.mutation.DeviceIDCleared() {
		_spec.ClearField(debugdevice.FieldDeviceID, field.TypeString)
	}
	if value, ok := dduo.mutation.AppID(); ok {
		_spec.SetField(debugdevice.FieldAppID, field.TypeUint32, value)
	}
	if value, ok := dduo.mutation.AddedAppID(); ok {
		_spec.AddField(debugdevice.FieldAppID, field.TypeUint32, value)
	}
	if dduo.mutation.AppIDCleared() {
		_spec.ClearField(debugdevice.FieldAppID, field.TypeUint32)
	}
	if value, ok := dduo.mutation.CreatorID(); ok {
		_spec.SetField(debugdevice.FieldCreatorID, field.TypeUint32, value)
	}
	if value, ok := dduo.mutation.AddedCreatorID(); ok {
		_spec.AddField(debugdevice.FieldCreatorID, field.TypeUint32, value)
	}
	if dduo.mutation.CreatorIDCleared() {
		_spec.ClearField(debugdevice.FieldCreatorID, field.TypeUint32)
	}
	if value, ok := dduo.mutation.Remark(); ok {
		_spec.SetField(debugdevice.FieldRemark, field.TypeString, value)
	}
	if dduo.mutation.RemarkCleared() {
		_spec.ClearField(debugdevice.FieldRemark, field.TypeString)
	}
	_spec.AddModifiers(dduo.modifiers...)
	_node = &DebugDevice{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{debugdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dduo.mutation.done = true
	return _node, nil
}
