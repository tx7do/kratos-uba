// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-bi/app/core/service/internal/data/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *int64 `json:"delete_time,omitempty"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 登陆密码
	Password *string `json:"password,omitempty"`
	// 昵称
	NickName *string `json:"nick_name,omitempty"`
	// 姓名
	RealName *string `json:"real_name,omitempty"`
	// 电子邮箱
	Email *string `json:"email,omitempty"`
	// 手机号码
	Phone *string `json:"phone,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`
	// 个人说明
	Description *string `json:"description,omitempty"`
	// 角色ID
	RoleID *uint32 `json:"role_id,omitempty"`
	// 授权
	Authority *user.Authority `json:"authority,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldCreateTime, user.FieldUpdateTime, user.FieldDeleteTime, user.FieldRoleID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserName, user.FieldPassword, user.FieldNickName, user.FieldRealName, user.FieldEmail, user.FieldPhone, user.FieldAvatar, user.FieldDescription, user.FieldAuthority:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint32(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = new(int64)
				*u.CreateTime = value.Int64
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = new(int64)
				*u.UpdateTime = value.Int64
			}
		case user.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				u.DeleteTime = new(int64)
				*u.DeleteTime = value.Int64
			}
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				u.UserName = new(string)
				*u.UserName = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = new(string)
				*u.Password = value.String
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = new(string)
				*u.NickName = value.String
			}
		case user.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				u.RealName = new(string)
				*u.RealName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = new(string)
				*u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = new(string)
				*u.Phone = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = new(string)
				*u.Avatar = value.String
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = new(string)
				*u.Description = value.String
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = new(uint32)
				*u.RoleID = uint32(value.Int64)
			}
		case user.FieldAuthority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authority", values[i])
			} else if value.Valid {
				u.Authority = new(user.Authority)
				*u.Authority = user.Authority(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	if v := u.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.UserName; v != nil {
		builder.WriteString("user_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.NickName; v != nil {
		builder.WriteString("nick_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.RealName; v != nil {
		builder.WriteString("real_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.RoleID; v != nil {
		builder.WriteString("role_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.Authority; v != nil {
		builder.WriteString("authority=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User