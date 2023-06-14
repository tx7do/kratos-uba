// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-bi/app/logger/service/internal/data/ent/acceptancestatus"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AcceptanceStatus is the model entity for the AcceptanceStatus schema.
type AcceptanceStatus struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID int64 `json:"id,omitempty"`
	// 数据名称
	DataName *string `json:"data_name,omitempty"`
	// 报告类型
	ReportType *string `json:"report_type,omitempty"`
	// 报告数据
	ReportData *string `json:"report_data,omitempty"`
	// 状态
	Status *int32 `json:"status,omitempty"`
	// 错误原因
	ErrorReason *string `json:"error_reason,omitempty"`
	// 错误处理
	ErrorHandling *string `json:"error_handling,omitempty"`
	// PartEvent holds the value of the "part_event" field.
	PartEvent *string `json:"part_event,omitempty"`
	// 记录时间
	PartDate *time.Time `json:"part_date,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AcceptanceStatus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case acceptancestatus.FieldID, acceptancestatus.FieldStatus:
			values[i] = new(sql.NullInt64)
		case acceptancestatus.FieldDataName, acceptancestatus.FieldReportType, acceptancestatus.FieldReportData, acceptancestatus.FieldErrorReason, acceptancestatus.FieldErrorHandling, acceptancestatus.FieldPartEvent:
			values[i] = new(sql.NullString)
		case acceptancestatus.FieldPartDate:
			values[i] = new(sql.NullTime)
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
			as.ID = int64(value.Int64)
		case acceptancestatus.FieldDataName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data_name", values[i])
			} else if value.Valid {
				as.DataName = new(string)
				*as.DataName = value.String
			}
		case acceptancestatus.FieldReportType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field report_type", values[i])
			} else if value.Valid {
				as.ReportType = new(string)
				*as.ReportType = value.String
			}
		case acceptancestatus.FieldReportData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field report_data", values[i])
			} else if value.Valid {
				as.ReportData = new(string)
				*as.ReportData = value.String
			}
		case acceptancestatus.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				as.Status = new(int32)
				*as.Status = int32(value.Int64)
			}
		case acceptancestatus.FieldErrorReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_reason", values[i])
			} else if value.Valid {
				as.ErrorReason = new(string)
				*as.ErrorReason = value.String
			}
		case acceptancestatus.FieldErrorHandling:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_handling", values[i])
			} else if value.Valid {
				as.ErrorHandling = new(string)
				*as.ErrorHandling = value.String
			}
		case acceptancestatus.FieldPartEvent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field part_event", values[i])
			} else if value.Valid {
				as.PartEvent = new(string)
				*as.PartEvent = value.String
			}
		case acceptancestatus.FieldPartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field part_date", values[i])
			} else if value.Valid {
				as.PartDate = new(time.Time)
				*as.PartDate = value.Time
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
	if v := as.DataName; v != nil {
		builder.WriteString("data_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.ReportType; v != nil {
		builder.WriteString("report_type=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.ReportData; v != nil {
		builder.WriteString("report_data=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := as.ErrorReason; v != nil {
		builder.WriteString("error_reason=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.ErrorHandling; v != nil {
		builder.WriteString("error_handling=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.PartEvent; v != nil {
		builder.WriteString("part_event=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := as.PartDate; v != nil {
		builder.WriteString("part_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// AcceptanceStatusSlice is a parsable slice of AcceptanceStatus.
type AcceptanceStatusSlice []*AcceptanceStatus
