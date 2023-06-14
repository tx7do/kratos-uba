// Code generated by ent, DO NOT EDIT.

package acceptancestatus

const (
	// Label holds the string label denoting the acceptancestatus type in the database.
	Label = "acceptance_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDataName holds the string denoting the data_name field in the database.
	FieldDataName = "data_name"
	// FieldReportType holds the string denoting the report_type field in the database.
	FieldReportType = "report_type"
	// FieldReportData holds the string denoting the report_data field in the database.
	FieldReportData = "report_data"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldErrorReason holds the string denoting the error_reason field in the database.
	FieldErrorReason = "error_reason"
	// FieldErrorHandling holds the string denoting the error_handling field in the database.
	FieldErrorHandling = "error_handling"
	// FieldPartEvent holds the string denoting the part_event field in the database.
	FieldPartEvent = "part_event"
	// FieldPartDate holds the string denoting the part_date field in the database.
	FieldPartDate = "part_date"
	// Table holds the table name of the acceptancestatus in the database.
	Table = "acceptance_status"
)

// Columns holds all SQL columns for acceptancestatus fields.
var Columns = []string{
	FieldID,
	FieldDataName,
	FieldReportType,
	FieldReportData,
	FieldStatus,
	FieldErrorReason,
	FieldErrorHandling,
	FieldPartEvent,
	FieldPartDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DataNameValidator is a validator for the "data_name" field. It is called by the builders before save.
	DataNameValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)