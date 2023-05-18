// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-bi/app/core/service/internal/data/ent/predicate"
	"kratos-bi/app/core/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(i int64) *UserUpdate {
	uu.mutation.ResetUpdateTime()
	uu.mutation.SetUpdateTime(i)
	return uu
}

// AddUpdateTime adds i to the "update_time" field.
func (uu *UserUpdate) AddUpdateTime(i int64) *UserUpdate {
	uu.mutation.AddUpdateTime(i)
	return uu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (uu *UserUpdate) ClearUpdateTime() *UserUpdate {
	uu.mutation.ClearUpdateTime()
	return uu
}

// SetDeleteTime sets the "delete_time" field.
func (uu *UserUpdate) SetDeleteTime(i int64) *UserUpdate {
	uu.mutation.ResetDeleteTime()
	uu.mutation.SetDeleteTime(i)
	return uu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeleteTime(i *int64) *UserUpdate {
	if i != nil {
		uu.SetDeleteTime(*i)
	}
	return uu
}

// AddDeleteTime adds i to the "delete_time" field.
func (uu *UserUpdate) AddDeleteTime(i int64) *UserUpdate {
	uu.mutation.AddDeleteTime(i)
	return uu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uu *UserUpdate) ClearDeleteTime() *UserUpdate {
	uu.mutation.ClearDeleteTime()
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetNickName sets the "nick_name" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// ClearNickName clears the value of the "nick_name" field.
func (uu *UserUpdate) ClearNickName() *UserUpdate {
	uu.mutation.ClearNickName()
	return uu
}

// SetRealName sets the "real_name" field.
func (uu *UserUpdate) SetRealName(s string) *UserUpdate {
	uu.mutation.SetRealName(s)
	return uu
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRealName(s *string) *UserUpdate {
	if s != nil {
		uu.SetRealName(*s)
	}
	return uu
}

// ClearRealName clears the value of the "real_name" field.
func (uu *UserUpdate) ClearRealName() *UserUpdate {
	uu.mutation.ClearRealName()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDescription(s *string) *UserUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UserUpdate) ClearDescription() *UserUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetRoleID sets the "role_id" field.
func (uu *UserUpdate) SetRoleID(u uint32) *UserUpdate {
	uu.mutation.ResetRoleID()
	uu.mutation.SetRoleID(u)
	return uu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetRoleID(*u)
	}
	return uu
}

// AddRoleID adds u to the "role_id" field.
func (uu *UserUpdate) AddRoleID(u int32) *UserUpdate {
	uu.mutation.AddRoleID(u)
	return uu
}

// ClearRoleID clears the value of the "role_id" field.
func (uu *UserUpdate) ClearRoleID() *UserUpdate {
	uu.mutation.ClearRoleID()
	return uu
}

// SetAuthority sets the "authority" field.
func (uu *UserUpdate) SetAuthority(u user.Authority) *UserUpdate {
	uu.mutation.SetAuthority(u)
	return uu
}

// SetNillableAuthority sets the "authority" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAuthority(u *user.Authority) *UserUpdate {
	if u != nil {
		uu.SetAuthority(*u)
	}
	return uu
}

// ClearAuthority clears the value of the "authority" field.
func (uu *UserUpdate) ClearAuthority() *UserUpdate {
	uu.mutation.ClearAuthority()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok && !uu.mutation.UpdateTimeCleared() {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "User.nick_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.RealName(); ok {
		if err := user.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "User.real_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "User.description": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Authority(); ok {
		if err := user.AuthorityValidator(v); err != nil {
			return &ValidationError{Name: "authority", err: fmt.Errorf(`ent: validator failed for field "User.authority": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uu.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if uu.mutation.UpdateTimeCleared() {
		_spec.ClearField(user.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := uu.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedDeleteTime(); ok {
		_spec.AddField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if uu.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeInt64)
	}
	if uu.mutation.UserNameCleared() {
		_spec.ClearField(user.FieldUserName, field.TypeString)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uu.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uu.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if uu.mutation.RealNameCleared() {
		_spec.ClearField(user.FieldRealName, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uu.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uu.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uu.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint32, value)
	}
	if value, ok := uu.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint32, value)
	}
	if uu.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint32)
	}
	if value, ok := uu.mutation.Authority(); ok {
		_spec.SetField(user.FieldAuthority, field.TypeEnum, value)
	}
	if uu.mutation.AuthorityCleared() {
		_spec.ClearField(user.FieldAuthority, field.TypeEnum)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdateTime()
	uuo.mutation.SetUpdateTime(i)
	return uuo
}

// AddUpdateTime adds i to the "update_time" field.
func (uuo *UserUpdateOne) AddUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdateTime(i)
	return uuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (uuo *UserUpdateOne) ClearUpdateTime() *UserUpdateOne {
	uuo.mutation.ClearUpdateTime()
	return uuo
}

// SetDeleteTime sets the "delete_time" field.
func (uuo *UserUpdateOne) SetDeleteTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetDeleteTime()
	uuo.mutation.SetDeleteTime(i)
	return uuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeleteTime(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetDeleteTime(*i)
	}
	return uuo
}

// AddDeleteTime adds i to the "delete_time" field.
func (uuo *UserUpdateOne) AddDeleteTime(i int64) *UserUpdateOne {
	uuo.mutation.AddDeleteTime(i)
	return uuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uuo *UserUpdateOne) ClearDeleteTime() *UserUpdateOne {
	uuo.mutation.ClearDeleteTime()
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetNickName sets the "nick_name" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// ClearNickName clears the value of the "nick_name" field.
func (uuo *UserUpdateOne) ClearNickName() *UserUpdateOne {
	uuo.mutation.ClearNickName()
	return uuo
}

// SetRealName sets the "real_name" field.
func (uuo *UserUpdateOne) SetRealName(s string) *UserUpdateOne {
	uuo.mutation.SetRealName(s)
	return uuo
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRealName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRealName(*s)
	}
	return uuo
}

// ClearRealName clears the value of the "real_name" field.
func (uuo *UserUpdateOne) ClearRealName() *UserUpdateOne {
	uuo.mutation.ClearRealName()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDescription(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UserUpdateOne) ClearDescription() *UserUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetRoleID sets the "role_id" field.
func (uuo *UserUpdateOne) SetRoleID(u uint32) *UserUpdateOne {
	uuo.mutation.ResetRoleID()
	uuo.mutation.SetRoleID(u)
	return uuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetRoleID(*u)
	}
	return uuo
}

// AddRoleID adds u to the "role_id" field.
func (uuo *UserUpdateOne) AddRoleID(u int32) *UserUpdateOne {
	uuo.mutation.AddRoleID(u)
	return uuo
}

// ClearRoleID clears the value of the "role_id" field.
func (uuo *UserUpdateOne) ClearRoleID() *UserUpdateOne {
	uuo.mutation.ClearRoleID()
	return uuo
}

// SetAuthority sets the "authority" field.
func (uuo *UserUpdateOne) SetAuthority(u user.Authority) *UserUpdateOne {
	uuo.mutation.SetAuthority(u)
	return uuo
}

// SetNillableAuthority sets the "authority" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAuthority(u *user.Authority) *UserUpdateOne {
	if u != nil {
		uuo.SetAuthority(*u)
	}
	return uuo
}

// ClearAuthority clears the value of the "authority" field.
func (uuo *UserUpdateOne) ClearAuthority() *UserUpdateOne {
	uuo.mutation.ClearAuthority()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok && !uuo.mutation.UpdateTimeCleared() {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "User.nick_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.RealName(); ok {
		if err := user.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "User.real_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "User.description": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Authority(); ok {
		if err := user.AuthorityValidator(v); err != nil {
			return &ValidationError{Name: "authority", err: fmt.Errorf(`ent: validator failed for field "User.authority": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uuo.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if uuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(user.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := uuo.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if uuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeInt64)
	}
	if uuo.mutation.UserNameCleared() {
		_spec.ClearField(user.FieldUserName, field.TypeString)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uuo.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uuo.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if uuo.mutation.RealNameCleared() {
		_spec.ClearField(user.FieldRealName, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uuo.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uuo.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint32, value)
	}
	if value, ok := uuo.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint32, value)
	}
	if uuo.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint32)
	}
	if value, ok := uuo.mutation.Authority(); ok {
		_spec.SetField(user.FieldAuthority, field.TypeEnum, value)
	}
	if uuo.mutation.AuthorityCleared() {
		_spec.ClearField(user.FieldAuthority, field.TypeEnum)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
