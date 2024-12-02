// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/debugdevice"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DebugDeviceCreate is the builder for creating a DebugDevice entity.
type DebugDeviceCreate struct {
	config
	mutation *DebugDeviceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ddc *DebugDeviceCreate) SetCreateTime(t time.Time) *DebugDeviceCreate {
	ddc.mutation.SetCreateTime(t)
	return ddc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableCreateTime(t *time.Time) *DebugDeviceCreate {
	if t != nil {
		ddc.SetCreateTime(*t)
	}
	return ddc
}

// SetUpdateTime sets the "update_time" field.
func (ddc *DebugDeviceCreate) SetUpdateTime(t time.Time) *DebugDeviceCreate {
	ddc.mutation.SetUpdateTime(t)
	return ddc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableUpdateTime(t *time.Time) *DebugDeviceCreate {
	if t != nil {
		ddc.SetUpdateTime(*t)
	}
	return ddc
}

// SetDeleteTime sets the "delete_time" field.
func (ddc *DebugDeviceCreate) SetDeleteTime(t time.Time) *DebugDeviceCreate {
	ddc.mutation.SetDeleteTime(t)
	return ddc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableDeleteTime(t *time.Time) *DebugDeviceCreate {
	if t != nil {
		ddc.SetDeleteTime(*t)
	}
	return ddc
}

// SetDeviceID sets the "device_id" field.
func (ddc *DebugDeviceCreate) SetDeviceID(s string) *DebugDeviceCreate {
	ddc.mutation.SetDeviceID(s)
	return ddc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableDeviceID(s *string) *DebugDeviceCreate {
	if s != nil {
		ddc.SetDeviceID(*s)
	}
	return ddc
}

// SetAppID sets the "app_id" field.
func (ddc *DebugDeviceCreate) SetAppID(u uint32) *DebugDeviceCreate {
	ddc.mutation.SetAppID(u)
	return ddc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableAppID(u *uint32) *DebugDeviceCreate {
	if u != nil {
		ddc.SetAppID(*u)
	}
	return ddc
}

// SetCreatorID sets the "creator_id" field.
func (ddc *DebugDeviceCreate) SetCreatorID(u uint32) *DebugDeviceCreate {
	ddc.mutation.SetCreatorID(u)
	return ddc
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableCreatorID(u *uint32) *DebugDeviceCreate {
	if u != nil {
		ddc.SetCreatorID(*u)
	}
	return ddc
}

// SetRemark sets the "remark" field.
func (ddc *DebugDeviceCreate) SetRemark(s string) *DebugDeviceCreate {
	ddc.mutation.SetRemark(s)
	return ddc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ddc *DebugDeviceCreate) SetNillableRemark(s *string) *DebugDeviceCreate {
	if s != nil {
		ddc.SetRemark(*s)
	}
	return ddc
}

// SetID sets the "id" field.
func (ddc *DebugDeviceCreate) SetID(u uint32) *DebugDeviceCreate {
	ddc.mutation.SetID(u)
	return ddc
}

// Mutation returns the DebugDeviceMutation object of the builder.
func (ddc *DebugDeviceCreate) Mutation() *DebugDeviceMutation {
	return ddc.mutation
}

// Save creates the DebugDevice in the database.
func (ddc *DebugDeviceCreate) Save(ctx context.Context) (*DebugDevice, error) {
	return withHooks(ctx, ddc.sqlSave, ddc.mutation, ddc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddc *DebugDeviceCreate) SaveX(ctx context.Context) *DebugDevice {
	v, err := ddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddc *DebugDeviceCreate) Exec(ctx context.Context) error {
	_, err := ddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddc *DebugDeviceCreate) ExecX(ctx context.Context) {
	if err := ddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddc *DebugDeviceCreate) check() error {
	if v, ok := ddc.mutation.DeviceID(); ok {
		if err := debugdevice.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`ent: validator failed for field "DebugDevice.device_id": %w`, err)}
		}
	}
	if v, ok := ddc.mutation.ID(); ok {
		if err := debugdevice.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "DebugDevice.id": %w`, err)}
		}
	}
	return nil
}

func (ddc *DebugDeviceCreate) sqlSave(ctx context.Context) (*DebugDevice, error) {
	if err := ddc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	ddc.mutation.id = &_node.ID
	ddc.mutation.done = true
	return _node, nil
}

func (ddc *DebugDeviceCreate) createSpec() (*DebugDevice, *sqlgraph.CreateSpec) {
	var (
		_node = &DebugDevice{config: ddc.config}
		_spec = sqlgraph.NewCreateSpec(debugdevice.Table, sqlgraph.NewFieldSpec(debugdevice.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = ddc.conflict
	if id, ok := ddc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ddc.mutation.CreateTime(); ok {
		_spec.SetField(debugdevice.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := ddc.mutation.UpdateTime(); ok {
		_spec.SetField(debugdevice.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := ddc.mutation.DeleteTime(); ok {
		_spec.SetField(debugdevice.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := ddc.mutation.DeviceID(); ok {
		_spec.SetField(debugdevice.FieldDeviceID, field.TypeString, value)
		_node.DeviceID = &value
	}
	if value, ok := ddc.mutation.AppID(); ok {
		_spec.SetField(debugdevice.FieldAppID, field.TypeUint32, value)
		_node.AppID = &value
	}
	if value, ok := ddc.mutation.CreatorID(); ok {
		_spec.SetField(debugdevice.FieldCreatorID, field.TypeUint32, value)
		_node.CreatorID = &value
	}
	if value, ok := ddc.mutation.Remark(); ok {
		_spec.SetField(debugdevice.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DebugDevice.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DebugDeviceUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ddc *DebugDeviceCreate) OnConflict(opts ...sql.ConflictOption) *DebugDeviceUpsertOne {
	ddc.conflict = opts
	return &DebugDeviceUpsertOne{
		create: ddc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ddc *DebugDeviceCreate) OnConflictColumns(columns ...string) *DebugDeviceUpsertOne {
	ddc.conflict = append(ddc.conflict, sql.ConflictColumns(columns...))
	return &DebugDeviceUpsertOne{
		create: ddc,
	}
}

type (
	// DebugDeviceUpsertOne is the builder for "upsert"-ing
	//  one DebugDevice node.
	DebugDeviceUpsertOne struct {
		create *DebugDeviceCreate
	}

	// DebugDeviceUpsert is the "OnConflict" setter.
	DebugDeviceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *DebugDeviceUpsert) SetUpdateTime(v time.Time) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateUpdateTime() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *DebugDeviceUpsert) ClearUpdateTime() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *DebugDeviceUpsert) SetDeleteTime(v time.Time) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateDeleteTime() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *DebugDeviceUpsert) ClearDeleteTime() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldDeleteTime)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *DebugDeviceUpsert) SetDeviceID(v string) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateDeviceID() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldDeviceID)
	return u
}

// ClearDeviceID clears the value of the "device_id" field.
func (u *DebugDeviceUpsert) ClearDeviceID() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldDeviceID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *DebugDeviceUpsert) SetAppID(v uint32) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateAppID() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldAppID)
	return u
}

// AddAppID adds v to the "app_id" field.
func (u *DebugDeviceUpsert) AddAppID(v uint32) *DebugDeviceUpsert {
	u.Add(debugdevice.FieldAppID, v)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *DebugDeviceUpsert) ClearAppID() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldAppID)
	return u
}

// SetCreatorID sets the "creator_id" field.
func (u *DebugDeviceUpsert) SetCreatorID(v uint32) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldCreatorID, v)
	return u
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateCreatorID() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldCreatorID)
	return u
}

// AddCreatorID adds v to the "creator_id" field.
func (u *DebugDeviceUpsert) AddCreatorID(v uint32) *DebugDeviceUpsert {
	u.Add(debugdevice.FieldCreatorID, v)
	return u
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *DebugDeviceUpsert) ClearCreatorID() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldCreatorID)
	return u
}

// SetRemark sets the "remark" field.
func (u *DebugDeviceUpsert) SetRemark(v string) *DebugDeviceUpsert {
	u.Set(debugdevice.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DebugDeviceUpsert) UpdateRemark() *DebugDeviceUpsert {
	u.SetExcluded(debugdevice.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *DebugDeviceUpsert) ClearRemark() *DebugDeviceUpsert {
	u.SetNull(debugdevice.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(debugdevice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DebugDeviceUpsertOne) UpdateNewValues() *DebugDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(debugdevice.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(debugdevice.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DebugDeviceUpsertOne) Ignore() *DebugDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DebugDeviceUpsertOne) DoNothing() *DebugDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DebugDeviceCreate.OnConflict
// documentation for more info.
func (u *DebugDeviceUpsertOne) Update(set func(*DebugDeviceUpsert)) *DebugDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DebugDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DebugDeviceUpsertOne) SetUpdateTime(v time.Time) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateUpdateTime() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *DebugDeviceUpsertOne) ClearUpdateTime() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *DebugDeviceUpsertOne) SetDeleteTime(v time.Time) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateDeleteTime() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *DebugDeviceUpsertOne) ClearDeleteTime() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearDeleteTime()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DebugDeviceUpsertOne) SetDeviceID(v string) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateDeviceID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// ClearDeviceID clears the value of the "device_id" field.
func (u *DebugDeviceUpsertOne) ClearDeviceID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearDeviceID()
	})
}

// SetAppID sets the "app_id" field.
func (u *DebugDeviceUpsertOne) SetAppID(v uint32) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *DebugDeviceUpsertOne) AddAppID(v uint32) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateAppID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *DebugDeviceUpsertOne) ClearAppID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearAppID()
	})
}

// SetCreatorID sets the "creator_id" field.
func (u *DebugDeviceUpsertOne) SetCreatorID(v uint32) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetCreatorID(v)
	})
}

// AddCreatorID adds v to the "creator_id" field.
func (u *DebugDeviceUpsertOne) AddCreatorID(v uint32) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.AddCreatorID(v)
	})
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateCreatorID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateCreatorID()
	})
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *DebugDeviceUpsertOne) ClearCreatorID() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearCreatorID()
	})
}

// SetRemark sets the "remark" field.
func (u *DebugDeviceUpsertOne) SetRemark(v string) *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DebugDeviceUpsertOne) UpdateRemark() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DebugDeviceUpsertOne) ClearRemark() *DebugDeviceUpsertOne {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *DebugDeviceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DebugDeviceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DebugDeviceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DebugDeviceUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DebugDeviceUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DebugDeviceCreateBulk is the builder for creating many DebugDevice entities in bulk.
type DebugDeviceCreateBulk struct {
	config
	err      error
	builders []*DebugDeviceCreate
	conflict []sql.ConflictOption
}

// Save creates the DebugDevice entities in the database.
func (ddcb *DebugDeviceCreateBulk) Save(ctx context.Context) ([]*DebugDevice, error) {
	if ddcb.err != nil {
		return nil, ddcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ddcb.builders))
	nodes := make([]*DebugDevice, len(ddcb.builders))
	mutators := make([]Mutator, len(ddcb.builders))
	for i := range ddcb.builders {
		func(i int, root context.Context) {
			builder := ddcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DebugDeviceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ddcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddcb *DebugDeviceCreateBulk) SaveX(ctx context.Context) []*DebugDevice {
	v, err := ddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddcb *DebugDeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := ddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddcb *DebugDeviceCreateBulk) ExecX(ctx context.Context) {
	if err := ddcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DebugDevice.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DebugDeviceUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ddcb *DebugDeviceCreateBulk) OnConflict(opts ...sql.ConflictOption) *DebugDeviceUpsertBulk {
	ddcb.conflict = opts
	return &DebugDeviceUpsertBulk{
		create: ddcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ddcb *DebugDeviceCreateBulk) OnConflictColumns(columns ...string) *DebugDeviceUpsertBulk {
	ddcb.conflict = append(ddcb.conflict, sql.ConflictColumns(columns...))
	return &DebugDeviceUpsertBulk{
		create: ddcb,
	}
}

// DebugDeviceUpsertBulk is the builder for "upsert"-ing
// a bulk of DebugDevice nodes.
type DebugDeviceUpsertBulk struct {
	create *DebugDeviceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(debugdevice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DebugDeviceUpsertBulk) UpdateNewValues() *DebugDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(debugdevice.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(debugdevice.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DebugDevice.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DebugDeviceUpsertBulk) Ignore() *DebugDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DebugDeviceUpsertBulk) DoNothing() *DebugDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DebugDeviceCreateBulk.OnConflict
// documentation for more info.
func (u *DebugDeviceUpsertBulk) Update(set func(*DebugDeviceUpsert)) *DebugDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DebugDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DebugDeviceUpsertBulk) SetUpdateTime(v time.Time) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateUpdateTime() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *DebugDeviceUpsertBulk) ClearUpdateTime() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *DebugDeviceUpsertBulk) SetDeleteTime(v time.Time) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateDeleteTime() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *DebugDeviceUpsertBulk) ClearDeleteTime() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearDeleteTime()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DebugDeviceUpsertBulk) SetDeviceID(v string) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateDeviceID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// ClearDeviceID clears the value of the "device_id" field.
func (u *DebugDeviceUpsertBulk) ClearDeviceID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearDeviceID()
	})
}

// SetAppID sets the "app_id" field.
func (u *DebugDeviceUpsertBulk) SetAppID(v uint32) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *DebugDeviceUpsertBulk) AddAppID(v uint32) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateAppID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *DebugDeviceUpsertBulk) ClearAppID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearAppID()
	})
}

// SetCreatorID sets the "creator_id" field.
func (u *DebugDeviceUpsertBulk) SetCreatorID(v uint32) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetCreatorID(v)
	})
}

// AddCreatorID adds v to the "creator_id" field.
func (u *DebugDeviceUpsertBulk) AddCreatorID(v uint32) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.AddCreatorID(v)
	})
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateCreatorID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateCreatorID()
	})
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *DebugDeviceUpsertBulk) ClearCreatorID() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearCreatorID()
	})
}

// SetRemark sets the "remark" field.
func (u *DebugDeviceUpsertBulk) SetRemark(v string) *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DebugDeviceUpsertBulk) UpdateRemark() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DebugDeviceUpsertBulk) ClearRemark() *DebugDeviceUpsertBulk {
	return u.Update(func(s *DebugDeviceUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *DebugDeviceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DebugDeviceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DebugDeviceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DebugDeviceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
