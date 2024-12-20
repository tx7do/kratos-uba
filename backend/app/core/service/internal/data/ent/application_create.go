// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/application"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ac *ApplicationCreate) SetCreateTime(t time.Time) *ApplicationCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreateTime(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *ApplicationCreate) SetUpdateTime(t time.Time) *ApplicationCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableUpdateTime(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetDeleteTime sets the "delete_time" field.
func (ac *ApplicationCreate) SetDeleteTime(t time.Time) *ApplicationCreate {
	ac.mutation.SetDeleteTime(t)
	return ac
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableDeleteTime(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetDeleteTime(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *ApplicationCreate) SetName(s string) *ApplicationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableName(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *ApplicationCreate) SetStatus(s string) *ApplicationCreate {
	ac.mutation.SetStatus(s)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableStatus(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetStatus(*s)
	}
	return ac
}

// SetAppID sets the "app_id" field.
func (ac *ApplicationCreate) SetAppID(s string) *ApplicationCreate {
	ac.mutation.SetAppID(s)
	return ac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableAppID(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetAppID(*s)
	}
	return ac
}

// SetAppKey sets the "app_key" field.
func (ac *ApplicationCreate) SetAppKey(s string) *ApplicationCreate {
	ac.mutation.SetAppKey(s)
	return ac
}

// SetNillableAppKey sets the "app_key" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableAppKey(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetAppKey(*s)
	}
	return ac
}

// SetRemark sets the "remark" field.
func (ac *ApplicationCreate) SetRemark(s string) *ApplicationCreate {
	ac.mutation.SetRemark(s)
	return ac
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableRemark(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetRemark(*s)
	}
	return ac
}

// SetCreatorID sets the "creator_id" field.
func (ac *ApplicationCreate) SetCreatorID(u uint32) *ApplicationCreate {
	ac.mutation.SetCreatorID(u)
	return ac
}

// SetNillableCreatorID sets the "creator_id" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreatorID(u *uint32) *ApplicationCreate {
	if u != nil {
		ac.SetCreatorID(*u)
	}
	return ac
}

// SetOwnerID sets the "owner_id" field.
func (ac *ApplicationCreate) SetOwnerID(u uint32) *ApplicationCreate {
	ac.mutation.SetOwnerID(u)
	return ac
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableOwnerID(u *uint32) *ApplicationCreate {
	if u != nil {
		ac.SetOwnerID(*u)
	}
	return ac
}

// SetKeepMonth sets the "keep_month" field.
func (ac *ApplicationCreate) SetKeepMonth(u uint32) *ApplicationCreate {
	ac.mutation.SetKeepMonth(u)
	return ac
}

// SetNillableKeepMonth sets the "keep_month" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableKeepMonth(u *uint32) *ApplicationCreate {
	if u != nil {
		ac.SetKeepMonth(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ApplicationCreate) SetID(u uint32) *ApplicationCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApplicationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApplicationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApplicationCreate) check() error {
	if v, ok := ac.mutation.Name(); ok {
		if err := application.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Application.name": %w`, err)}
		}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := application.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Application.id": %w`, err)}
		}
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		_node = &Application{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(application.Table, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(application.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(application.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := ac.mutation.DeleteTime(); ok {
		_spec.SetField(application.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(application.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(application.FieldStatus, field.TypeString, value)
		_node.Status = &value
	}
	if value, ok := ac.mutation.AppID(); ok {
		_spec.SetField(application.FieldAppID, field.TypeString, value)
		_node.AppID = &value
	}
	if value, ok := ac.mutation.AppKey(); ok {
		_spec.SetField(application.FieldAppKey, field.TypeString, value)
		_node.AppKey = &value
	}
	if value, ok := ac.mutation.Remark(); ok {
		_spec.SetField(application.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := ac.mutation.CreatorID(); ok {
		_spec.SetField(application.FieldCreatorID, field.TypeUint32, value)
		_node.CreatorID = &value
	}
	if value, ok := ac.mutation.OwnerID(); ok {
		_spec.SetField(application.FieldOwnerID, field.TypeUint32, value)
		_node.OwnerID = &value
	}
	if value, ok := ac.mutation.KeepMonth(); ok {
		_spec.SetField(application.FieldKeepMonth, field.TypeUint32, value)
		_node.KeepMonth = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Application.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApplicationUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ac *ApplicationCreate) OnConflict(opts ...sql.ConflictOption) *ApplicationUpsertOne {
	ac.conflict = opts
	return &ApplicationUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Application.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *ApplicationCreate) OnConflictColumns(columns ...string) *ApplicationUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ApplicationUpsertOne{
		create: ac,
	}
}

type (
	// ApplicationUpsertOne is the builder for "upsert"-ing
	//  one Application node.
	ApplicationUpsertOne struct {
		create *ApplicationCreate
	}

	// ApplicationUpsert is the "OnConflict" setter.
	ApplicationUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *ApplicationUpsert) SetUpdateTime(v time.Time) *ApplicationUpsert {
	u.Set(application.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateUpdateTime() *ApplicationUpsert {
	u.SetExcluded(application.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApplicationUpsert) ClearUpdateTime() *ApplicationUpsert {
	u.SetNull(application.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *ApplicationUpsert) SetDeleteTime(v time.Time) *ApplicationUpsert {
	u.Set(application.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateDeleteTime() *ApplicationUpsert {
	u.SetExcluded(application.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *ApplicationUpsert) ClearDeleteTime() *ApplicationUpsert {
	u.SetNull(application.FieldDeleteTime)
	return u
}

// SetName sets the "name" field.
func (u *ApplicationUpsert) SetName(v string) *ApplicationUpsert {
	u.Set(application.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateName() *ApplicationUpsert {
	u.SetExcluded(application.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *ApplicationUpsert) ClearName() *ApplicationUpsert {
	u.SetNull(application.FieldName)
	return u
}

// SetStatus sets the "status" field.
func (u *ApplicationUpsert) SetStatus(v string) *ApplicationUpsert {
	u.Set(application.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateStatus() *ApplicationUpsert {
	u.SetExcluded(application.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *ApplicationUpsert) ClearStatus() *ApplicationUpsert {
	u.SetNull(application.FieldStatus)
	return u
}

// SetAppID sets the "app_id" field.
func (u *ApplicationUpsert) SetAppID(v string) *ApplicationUpsert {
	u.Set(application.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateAppID() *ApplicationUpsert {
	u.SetExcluded(application.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *ApplicationUpsert) ClearAppID() *ApplicationUpsert {
	u.SetNull(application.FieldAppID)
	return u
}

// SetAppKey sets the "app_key" field.
func (u *ApplicationUpsert) SetAppKey(v string) *ApplicationUpsert {
	u.Set(application.FieldAppKey, v)
	return u
}

// UpdateAppKey sets the "app_key" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateAppKey() *ApplicationUpsert {
	u.SetExcluded(application.FieldAppKey)
	return u
}

// ClearAppKey clears the value of the "app_key" field.
func (u *ApplicationUpsert) ClearAppKey() *ApplicationUpsert {
	u.SetNull(application.FieldAppKey)
	return u
}

// SetRemark sets the "remark" field.
func (u *ApplicationUpsert) SetRemark(v string) *ApplicationUpsert {
	u.Set(application.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateRemark() *ApplicationUpsert {
	u.SetExcluded(application.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *ApplicationUpsert) ClearRemark() *ApplicationUpsert {
	u.SetNull(application.FieldRemark)
	return u
}

// SetCreatorID sets the "creator_id" field.
func (u *ApplicationUpsert) SetCreatorID(v uint32) *ApplicationUpsert {
	u.Set(application.FieldCreatorID, v)
	return u
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateCreatorID() *ApplicationUpsert {
	u.SetExcluded(application.FieldCreatorID)
	return u
}

// AddCreatorID adds v to the "creator_id" field.
func (u *ApplicationUpsert) AddCreatorID(v uint32) *ApplicationUpsert {
	u.Add(application.FieldCreatorID, v)
	return u
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *ApplicationUpsert) ClearCreatorID() *ApplicationUpsert {
	u.SetNull(application.FieldCreatorID)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *ApplicationUpsert) SetOwnerID(v uint32) *ApplicationUpsert {
	u.Set(application.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateOwnerID() *ApplicationUpsert {
	u.SetExcluded(application.FieldOwnerID)
	return u
}

// AddOwnerID adds v to the "owner_id" field.
func (u *ApplicationUpsert) AddOwnerID(v uint32) *ApplicationUpsert {
	u.Add(application.FieldOwnerID, v)
	return u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *ApplicationUpsert) ClearOwnerID() *ApplicationUpsert {
	u.SetNull(application.FieldOwnerID)
	return u
}

// SetKeepMonth sets the "keep_month" field.
func (u *ApplicationUpsert) SetKeepMonth(v uint32) *ApplicationUpsert {
	u.Set(application.FieldKeepMonth, v)
	return u
}

// UpdateKeepMonth sets the "keep_month" field to the value that was provided on create.
func (u *ApplicationUpsert) UpdateKeepMonth() *ApplicationUpsert {
	u.SetExcluded(application.FieldKeepMonth)
	return u
}

// AddKeepMonth adds v to the "keep_month" field.
func (u *ApplicationUpsert) AddKeepMonth(v uint32) *ApplicationUpsert {
	u.Add(application.FieldKeepMonth, v)
	return u
}

// ClearKeepMonth clears the value of the "keep_month" field.
func (u *ApplicationUpsert) ClearKeepMonth() *ApplicationUpsert {
	u.SetNull(application.FieldKeepMonth)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Application.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(application.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApplicationUpsertOne) UpdateNewValues() *ApplicationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(application.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(application.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Application.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApplicationUpsertOne) Ignore() *ApplicationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApplicationUpsertOne) DoNothing() *ApplicationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApplicationCreate.OnConflict
// documentation for more info.
func (u *ApplicationUpsertOne) Update(set func(*ApplicationUpsert)) *ApplicationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApplicationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ApplicationUpsertOne) SetUpdateTime(v time.Time) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateUpdateTime() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApplicationUpsertOne) ClearUpdateTime() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *ApplicationUpsertOne) SetDeleteTime(v time.Time) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateDeleteTime() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *ApplicationUpsertOne) ClearDeleteTime() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearDeleteTime()
	})
}

// SetName sets the "name" field.
func (u *ApplicationUpsertOne) SetName(v string) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateName() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ApplicationUpsertOne) ClearName() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearName()
	})
}

// SetStatus sets the "status" field.
func (u *ApplicationUpsertOne) SetStatus(v string) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateStatus() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ApplicationUpsertOne) ClearStatus() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearStatus()
	})
}

// SetAppID sets the "app_id" field.
func (u *ApplicationUpsertOne) SetAppID(v string) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateAppID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *ApplicationUpsertOne) ClearAppID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearAppID()
	})
}

// SetAppKey sets the "app_key" field.
func (u *ApplicationUpsertOne) SetAppKey(v string) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetAppKey(v)
	})
}

// UpdateAppKey sets the "app_key" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateAppKey() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateAppKey()
	})
}

// ClearAppKey clears the value of the "app_key" field.
func (u *ApplicationUpsertOne) ClearAppKey() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearAppKey()
	})
}

// SetRemark sets the "remark" field.
func (u *ApplicationUpsertOne) SetRemark(v string) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateRemark() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *ApplicationUpsertOne) ClearRemark() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearRemark()
	})
}

// SetCreatorID sets the "creator_id" field.
func (u *ApplicationUpsertOne) SetCreatorID(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetCreatorID(v)
	})
}

// AddCreatorID adds v to the "creator_id" field.
func (u *ApplicationUpsertOne) AddCreatorID(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddCreatorID(v)
	})
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateCreatorID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateCreatorID()
	})
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *ApplicationUpsertOne) ClearCreatorID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearCreatorID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *ApplicationUpsertOne) SetOwnerID(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetOwnerID(v)
	})
}

// AddOwnerID adds v to the "owner_id" field.
func (u *ApplicationUpsertOne) AddOwnerID(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateOwnerID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateOwnerID()
	})
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *ApplicationUpsertOne) ClearOwnerID() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearOwnerID()
	})
}

// SetKeepMonth sets the "keep_month" field.
func (u *ApplicationUpsertOne) SetKeepMonth(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetKeepMonth(v)
	})
}

// AddKeepMonth adds v to the "keep_month" field.
func (u *ApplicationUpsertOne) AddKeepMonth(v uint32) *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddKeepMonth(v)
	})
}

// UpdateKeepMonth sets the "keep_month" field to the value that was provided on create.
func (u *ApplicationUpsertOne) UpdateKeepMonth() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateKeepMonth()
	})
}

// ClearKeepMonth clears the value of the "keep_month" field.
func (u *ApplicationUpsertOne) ClearKeepMonth() *ApplicationUpsertOne {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearKeepMonth()
	})
}

// Exec executes the query.
func (u *ApplicationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApplicationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApplicationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApplicationUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApplicationUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ApplicationCreateBulk is the builder for creating many Application entities in bulk.
type ApplicationCreateBulk struct {
	config
	err      error
	builders []*ApplicationCreate
	conflict []sql.ConflictOption
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApplicationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApplicationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Application.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApplicationUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (acb *ApplicationCreateBulk) OnConflict(opts ...sql.ConflictOption) *ApplicationUpsertBulk {
	acb.conflict = opts
	return &ApplicationUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Application.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *ApplicationCreateBulk) OnConflictColumns(columns ...string) *ApplicationUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ApplicationUpsertBulk{
		create: acb,
	}
}

// ApplicationUpsertBulk is the builder for "upsert"-ing
// a bulk of Application nodes.
type ApplicationUpsertBulk struct {
	create *ApplicationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Application.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(application.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApplicationUpsertBulk) UpdateNewValues() *ApplicationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(application.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(application.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Application.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApplicationUpsertBulk) Ignore() *ApplicationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApplicationUpsertBulk) DoNothing() *ApplicationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApplicationCreateBulk.OnConflict
// documentation for more info.
func (u *ApplicationUpsertBulk) Update(set func(*ApplicationUpsert)) *ApplicationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApplicationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ApplicationUpsertBulk) SetUpdateTime(v time.Time) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateUpdateTime() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *ApplicationUpsertBulk) ClearUpdateTime() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *ApplicationUpsertBulk) SetDeleteTime(v time.Time) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateDeleteTime() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *ApplicationUpsertBulk) ClearDeleteTime() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearDeleteTime()
	})
}

// SetName sets the "name" field.
func (u *ApplicationUpsertBulk) SetName(v string) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateName() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ApplicationUpsertBulk) ClearName() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearName()
	})
}

// SetStatus sets the "status" field.
func (u *ApplicationUpsertBulk) SetStatus(v string) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateStatus() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ApplicationUpsertBulk) ClearStatus() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearStatus()
	})
}

// SetAppID sets the "app_id" field.
func (u *ApplicationUpsertBulk) SetAppID(v string) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateAppID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *ApplicationUpsertBulk) ClearAppID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearAppID()
	})
}

// SetAppKey sets the "app_key" field.
func (u *ApplicationUpsertBulk) SetAppKey(v string) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetAppKey(v)
	})
}

// UpdateAppKey sets the "app_key" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateAppKey() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateAppKey()
	})
}

// ClearAppKey clears the value of the "app_key" field.
func (u *ApplicationUpsertBulk) ClearAppKey() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearAppKey()
	})
}

// SetRemark sets the "remark" field.
func (u *ApplicationUpsertBulk) SetRemark(v string) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateRemark() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *ApplicationUpsertBulk) ClearRemark() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearRemark()
	})
}

// SetCreatorID sets the "creator_id" field.
func (u *ApplicationUpsertBulk) SetCreatorID(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetCreatorID(v)
	})
}

// AddCreatorID adds v to the "creator_id" field.
func (u *ApplicationUpsertBulk) AddCreatorID(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddCreatorID(v)
	})
}

// UpdateCreatorID sets the "creator_id" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateCreatorID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateCreatorID()
	})
}

// ClearCreatorID clears the value of the "creator_id" field.
func (u *ApplicationUpsertBulk) ClearCreatorID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearCreatorID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *ApplicationUpsertBulk) SetOwnerID(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetOwnerID(v)
	})
}

// AddOwnerID adds v to the "owner_id" field.
func (u *ApplicationUpsertBulk) AddOwnerID(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateOwnerID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateOwnerID()
	})
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *ApplicationUpsertBulk) ClearOwnerID() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearOwnerID()
	})
}

// SetKeepMonth sets the "keep_month" field.
func (u *ApplicationUpsertBulk) SetKeepMonth(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.SetKeepMonth(v)
	})
}

// AddKeepMonth adds v to the "keep_month" field.
func (u *ApplicationUpsertBulk) AddKeepMonth(v uint32) *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.AddKeepMonth(v)
	})
}

// UpdateKeepMonth sets the "keep_month" field to the value that was provided on create.
func (u *ApplicationUpsertBulk) UpdateKeepMonth() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.UpdateKeepMonth()
	})
}

// ClearKeepMonth clears the value of the "keep_month" field.
func (u *ApplicationUpsertBulk) ClearKeepMonth() *ApplicationUpsertBulk {
	return u.Update(func(s *ApplicationUpsert) {
		s.ClearKeepMonth()
	})
}

// Exec executes the query.
func (u *ApplicationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ApplicationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApplicationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApplicationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
