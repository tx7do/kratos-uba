// Code generated by ent, DO NOT EDIT.

package acceptancestatus

import (
	"kratos-bi/app/logger/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldID, id))
}

// DataName applies equality check predicate on the "data_name" field. It's identical to DataNameEQ.
func DataName(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldDataName, v))
}

// ReportType applies equality check predicate on the "report_type" field. It's identical to ReportTypeEQ.
func ReportType(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldReportType, v))
}

// ReportData applies equality check predicate on the "report_data" field. It's identical to ReportDataEQ.
func ReportData(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldReportData, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldStatus, v))
}

// ErrorReason applies equality check predicate on the "error_reason" field. It's identical to ErrorReasonEQ.
func ErrorReason(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldErrorReason, v))
}

// ErrorHandling applies equality check predicate on the "error_handling" field. It's identical to ErrorHandlingEQ.
func ErrorHandling(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldErrorHandling, v))
}

// PartEvent applies equality check predicate on the "part_event" field. It's identical to PartEventEQ.
func PartEvent(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldPartEvent, v))
}

// PartDate applies equality check predicate on the "part_date" field. It's identical to PartDateEQ.
func PartDate(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldPartDate, v))
}

// DataNameEQ applies the EQ predicate on the "data_name" field.
func DataNameEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldDataName, v))
}

// DataNameNEQ applies the NEQ predicate on the "data_name" field.
func DataNameNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldDataName, v))
}

// DataNameIn applies the In predicate on the "data_name" field.
func DataNameIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldDataName, vs...))
}

// DataNameNotIn applies the NotIn predicate on the "data_name" field.
func DataNameNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldDataName, vs...))
}

// DataNameGT applies the GT predicate on the "data_name" field.
func DataNameGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldDataName, v))
}

// DataNameGTE applies the GTE predicate on the "data_name" field.
func DataNameGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldDataName, v))
}

// DataNameLT applies the LT predicate on the "data_name" field.
func DataNameLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldDataName, v))
}

// DataNameLTE applies the LTE predicate on the "data_name" field.
func DataNameLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldDataName, v))
}

// DataNameContains applies the Contains predicate on the "data_name" field.
func DataNameContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldDataName, v))
}

// DataNameHasPrefix applies the HasPrefix predicate on the "data_name" field.
func DataNameHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldDataName, v))
}

// DataNameHasSuffix applies the HasSuffix predicate on the "data_name" field.
func DataNameHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldDataName, v))
}

// DataNameIsNil applies the IsNil predicate on the "data_name" field.
func DataNameIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldDataName))
}

// DataNameNotNil applies the NotNil predicate on the "data_name" field.
func DataNameNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldDataName))
}

// DataNameEqualFold applies the EqualFold predicate on the "data_name" field.
func DataNameEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldDataName, v))
}

// DataNameContainsFold applies the ContainsFold predicate on the "data_name" field.
func DataNameContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldDataName, v))
}

// ReportTypeEQ applies the EQ predicate on the "report_type" field.
func ReportTypeEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldReportType, v))
}

// ReportTypeNEQ applies the NEQ predicate on the "report_type" field.
func ReportTypeNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldReportType, v))
}

// ReportTypeIn applies the In predicate on the "report_type" field.
func ReportTypeIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldReportType, vs...))
}

// ReportTypeNotIn applies the NotIn predicate on the "report_type" field.
func ReportTypeNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldReportType, vs...))
}

// ReportTypeGT applies the GT predicate on the "report_type" field.
func ReportTypeGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldReportType, v))
}

// ReportTypeGTE applies the GTE predicate on the "report_type" field.
func ReportTypeGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldReportType, v))
}

// ReportTypeLT applies the LT predicate on the "report_type" field.
func ReportTypeLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldReportType, v))
}

// ReportTypeLTE applies the LTE predicate on the "report_type" field.
func ReportTypeLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldReportType, v))
}

// ReportTypeContains applies the Contains predicate on the "report_type" field.
func ReportTypeContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldReportType, v))
}

// ReportTypeHasPrefix applies the HasPrefix predicate on the "report_type" field.
func ReportTypeHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldReportType, v))
}

// ReportTypeHasSuffix applies the HasSuffix predicate on the "report_type" field.
func ReportTypeHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldReportType, v))
}

// ReportTypeIsNil applies the IsNil predicate on the "report_type" field.
func ReportTypeIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldReportType))
}

// ReportTypeNotNil applies the NotNil predicate on the "report_type" field.
func ReportTypeNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldReportType))
}

// ReportTypeEqualFold applies the EqualFold predicate on the "report_type" field.
func ReportTypeEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldReportType, v))
}

// ReportTypeContainsFold applies the ContainsFold predicate on the "report_type" field.
func ReportTypeContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldReportType, v))
}

// ReportDataEQ applies the EQ predicate on the "report_data" field.
func ReportDataEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldReportData, v))
}

// ReportDataNEQ applies the NEQ predicate on the "report_data" field.
func ReportDataNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldReportData, v))
}

// ReportDataIn applies the In predicate on the "report_data" field.
func ReportDataIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldReportData, vs...))
}

// ReportDataNotIn applies the NotIn predicate on the "report_data" field.
func ReportDataNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldReportData, vs...))
}

// ReportDataGT applies the GT predicate on the "report_data" field.
func ReportDataGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldReportData, v))
}

// ReportDataGTE applies the GTE predicate on the "report_data" field.
func ReportDataGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldReportData, v))
}

// ReportDataLT applies the LT predicate on the "report_data" field.
func ReportDataLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldReportData, v))
}

// ReportDataLTE applies the LTE predicate on the "report_data" field.
func ReportDataLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldReportData, v))
}

// ReportDataContains applies the Contains predicate on the "report_data" field.
func ReportDataContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldReportData, v))
}

// ReportDataHasPrefix applies the HasPrefix predicate on the "report_data" field.
func ReportDataHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldReportData, v))
}

// ReportDataHasSuffix applies the HasSuffix predicate on the "report_data" field.
func ReportDataHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldReportData, v))
}

// ReportDataIsNil applies the IsNil predicate on the "report_data" field.
func ReportDataIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldReportData))
}

// ReportDataNotNil applies the NotNil predicate on the "report_data" field.
func ReportDataNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldReportData))
}

// ReportDataEqualFold applies the EqualFold predicate on the "report_data" field.
func ReportDataEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldReportData, v))
}

// ReportDataContainsFold applies the ContainsFold predicate on the "report_data" field.
func ReportDataContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldReportData, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int32) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldStatus))
}

// ErrorReasonEQ applies the EQ predicate on the "error_reason" field.
func ErrorReasonEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldErrorReason, v))
}

// ErrorReasonNEQ applies the NEQ predicate on the "error_reason" field.
func ErrorReasonNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldErrorReason, v))
}

// ErrorReasonIn applies the In predicate on the "error_reason" field.
func ErrorReasonIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldErrorReason, vs...))
}

// ErrorReasonNotIn applies the NotIn predicate on the "error_reason" field.
func ErrorReasonNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldErrorReason, vs...))
}

// ErrorReasonGT applies the GT predicate on the "error_reason" field.
func ErrorReasonGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldErrorReason, v))
}

// ErrorReasonGTE applies the GTE predicate on the "error_reason" field.
func ErrorReasonGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldErrorReason, v))
}

// ErrorReasonLT applies the LT predicate on the "error_reason" field.
func ErrorReasonLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldErrorReason, v))
}

// ErrorReasonLTE applies the LTE predicate on the "error_reason" field.
func ErrorReasonLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldErrorReason, v))
}

// ErrorReasonContains applies the Contains predicate on the "error_reason" field.
func ErrorReasonContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldErrorReason, v))
}

// ErrorReasonHasPrefix applies the HasPrefix predicate on the "error_reason" field.
func ErrorReasonHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldErrorReason, v))
}

// ErrorReasonHasSuffix applies the HasSuffix predicate on the "error_reason" field.
func ErrorReasonHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldErrorReason, v))
}

// ErrorReasonIsNil applies the IsNil predicate on the "error_reason" field.
func ErrorReasonIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldErrorReason))
}

// ErrorReasonNotNil applies the NotNil predicate on the "error_reason" field.
func ErrorReasonNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldErrorReason))
}

// ErrorReasonEqualFold applies the EqualFold predicate on the "error_reason" field.
func ErrorReasonEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldErrorReason, v))
}

// ErrorReasonContainsFold applies the ContainsFold predicate on the "error_reason" field.
func ErrorReasonContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldErrorReason, v))
}

// ErrorHandlingEQ applies the EQ predicate on the "error_handling" field.
func ErrorHandlingEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldErrorHandling, v))
}

// ErrorHandlingNEQ applies the NEQ predicate on the "error_handling" field.
func ErrorHandlingNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldErrorHandling, v))
}

// ErrorHandlingIn applies the In predicate on the "error_handling" field.
func ErrorHandlingIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldErrorHandling, vs...))
}

// ErrorHandlingNotIn applies the NotIn predicate on the "error_handling" field.
func ErrorHandlingNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldErrorHandling, vs...))
}

// ErrorHandlingGT applies the GT predicate on the "error_handling" field.
func ErrorHandlingGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldErrorHandling, v))
}

// ErrorHandlingGTE applies the GTE predicate on the "error_handling" field.
func ErrorHandlingGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldErrorHandling, v))
}

// ErrorHandlingLT applies the LT predicate on the "error_handling" field.
func ErrorHandlingLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldErrorHandling, v))
}

// ErrorHandlingLTE applies the LTE predicate on the "error_handling" field.
func ErrorHandlingLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldErrorHandling, v))
}

// ErrorHandlingContains applies the Contains predicate on the "error_handling" field.
func ErrorHandlingContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldErrorHandling, v))
}

// ErrorHandlingHasPrefix applies the HasPrefix predicate on the "error_handling" field.
func ErrorHandlingHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldErrorHandling, v))
}

// ErrorHandlingHasSuffix applies the HasSuffix predicate on the "error_handling" field.
func ErrorHandlingHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldErrorHandling, v))
}

// ErrorHandlingIsNil applies the IsNil predicate on the "error_handling" field.
func ErrorHandlingIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldErrorHandling))
}

// ErrorHandlingNotNil applies the NotNil predicate on the "error_handling" field.
func ErrorHandlingNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldErrorHandling))
}

// ErrorHandlingEqualFold applies the EqualFold predicate on the "error_handling" field.
func ErrorHandlingEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldErrorHandling, v))
}

// ErrorHandlingContainsFold applies the ContainsFold predicate on the "error_handling" field.
func ErrorHandlingContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldErrorHandling, v))
}

// PartEventEQ applies the EQ predicate on the "part_event" field.
func PartEventEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldPartEvent, v))
}

// PartEventNEQ applies the NEQ predicate on the "part_event" field.
func PartEventNEQ(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldPartEvent, v))
}

// PartEventIn applies the In predicate on the "part_event" field.
func PartEventIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldPartEvent, vs...))
}

// PartEventNotIn applies the NotIn predicate on the "part_event" field.
func PartEventNotIn(vs ...string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldPartEvent, vs...))
}

// PartEventGT applies the GT predicate on the "part_event" field.
func PartEventGT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldPartEvent, v))
}

// PartEventGTE applies the GTE predicate on the "part_event" field.
func PartEventGTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldPartEvent, v))
}

// PartEventLT applies the LT predicate on the "part_event" field.
func PartEventLT(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldPartEvent, v))
}

// PartEventLTE applies the LTE predicate on the "part_event" field.
func PartEventLTE(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldPartEvent, v))
}

// PartEventContains applies the Contains predicate on the "part_event" field.
func PartEventContains(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContains(FieldPartEvent, v))
}

// PartEventHasPrefix applies the HasPrefix predicate on the "part_event" field.
func PartEventHasPrefix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasPrefix(FieldPartEvent, v))
}

// PartEventHasSuffix applies the HasSuffix predicate on the "part_event" field.
func PartEventHasSuffix(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldHasSuffix(FieldPartEvent, v))
}

// PartEventIsNil applies the IsNil predicate on the "part_event" field.
func PartEventIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldPartEvent))
}

// PartEventNotNil applies the NotNil predicate on the "part_event" field.
func PartEventNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldPartEvent))
}

// PartEventEqualFold applies the EqualFold predicate on the "part_event" field.
func PartEventEqualFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEqualFold(FieldPartEvent, v))
}

// PartEventContainsFold applies the ContainsFold predicate on the "part_event" field.
func PartEventContainsFold(v string) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldContainsFold(FieldPartEvent, v))
}

// PartDateEQ applies the EQ predicate on the "part_date" field.
func PartDateEQ(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldEQ(FieldPartDate, v))
}

// PartDateNEQ applies the NEQ predicate on the "part_date" field.
func PartDateNEQ(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNEQ(FieldPartDate, v))
}

// PartDateIn applies the In predicate on the "part_date" field.
func PartDateIn(vs ...time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIn(FieldPartDate, vs...))
}

// PartDateNotIn applies the NotIn predicate on the "part_date" field.
func PartDateNotIn(vs ...time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotIn(FieldPartDate, vs...))
}

// PartDateGT applies the GT predicate on the "part_date" field.
func PartDateGT(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGT(FieldPartDate, v))
}

// PartDateGTE applies the GTE predicate on the "part_date" field.
func PartDateGTE(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldGTE(FieldPartDate, v))
}

// PartDateLT applies the LT predicate on the "part_date" field.
func PartDateLT(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLT(FieldPartDate, v))
}

// PartDateLTE applies the LTE predicate on the "part_date" field.
func PartDateLTE(v time.Time) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldLTE(FieldPartDate, v))
}

// PartDateIsNil applies the IsNil predicate on the "part_date" field.
func PartDateIsNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldIsNull(FieldPartDate))
}

// PartDateNotNil applies the NotNil predicate on the "part_date" field.
func PartDateNotNil() predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(sql.FieldNotNull(FieldPartDate))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AcceptanceStatus) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AcceptanceStatus) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AcceptanceStatus) predicate.AcceptanceStatus {
	return predicate.AcceptanceStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
