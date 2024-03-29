// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-uba/app/report/service/internal/data/ent/acceptancestatus"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AcceptanceStatus is the model entity for the AcceptanceStatus schema.
type AcceptanceStatus struct {
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
	// 应用名称
	Name *string `json:"name,omitempty"`
	// 应用状态
	Status *string `json:"status,omitempty"`
	// 应用ID
	AppID *string `json:"app_id,omitempty"`
	// 应用密钥
	AppKey *string `json:"app_key,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 创建者ID
	CreatorID *uint32 `json:"creator_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AcceptanceStatus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case acceptancestatus.FieldID, acceptancestatus.FieldCreateTime, acceptancestatus.FieldUpdateTime, acceptancestatus.FieldDeleteTime, acceptancestatus.FieldCreatorID:
			values[i] = new(sql.NullInt64)
		case acceptancestatus.FieldName, acceptancestatus.FieldStatus, acceptancestatus.FieldAppID, acceptancestatus.FieldAppKey, acceptancestatus.FieldRemark:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AcceptanceStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AcceptanceStatus fields.
func (as *AcceptanceStatus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case acceptancestatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = uint32(value.Int64)
		case acceptancestatus.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				as.CreateTime = new(int64)
				*as.CreateTime = value.Int64
			}
		case acceptancestatus.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				as.UpdateTime = new(int64)
				*as.UpdateTime = value.Int64
			}
		case acceptancestatus.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				as.DeleteTime = new(int64)
				*as.DeleteTime = value.Int64
			}
		case acceptancestatus.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				as.Name = new(string)
				*as.Name = value.String
			}
		case acceptancestatus.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				as.Status = new(string)
				*as.Status = value.String
			}
		case acceptancestatus.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				as.AppID = new(string)
				*as.AppID = value.String
			}
		case acceptancestatus.FieldAppKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_key", values[i])
			} else if value.Valid {
				as.AppKey = new(string)
				*as.AppKey = value.String
			}
		case acceptancestatus.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				as.Remark = new(string)
				*as.Remark = value.String
			}
		case acceptancestatus.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				as.CreatorID = new(uint32)
				*as.CreatorID = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AcceptanceStatus.
// Note that you need to call AcceptanceStatus.Unwrap() before calling this method if this AcceptanceStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AcceptanceStatus) Update() *AcceptanceStatusUpdateOne {
	return NewAcceptanceStatusClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AcceptanceStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AcceptanceStatus) Unwrap() *AcceptanceStatus {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AcceptanceStatus is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AcceptanceStatus) String() string {
	var builder strings.Builder
	builder.WriteString("AcceptanceStatus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	if v := as.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := as.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := as.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := as.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.AppID; v != nil {
		builder.WriteString("app_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.AppKey; v != nil {
		builder.WriteString("app_key=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.CreatorID; v != nil {
		builder.WriteString("creator_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// AcceptanceStatusSlice is a parsable slice of AcceptanceStatus.
type AcceptanceStatusSlice []*AcceptanceStatus
