// Code generated by ent, DO NOT EDIT.

package realtimewarehousing

import (
	"kratos-bi/app/report/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldDeleteTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldStatus, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldAppID, v))
}

// AppKey applies equality check predicate on the "app_key" field. It's identical to AppKeyEQ.
func AppKey(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldAppKey, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldRemark, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldCreatorID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v int64) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldDeleteTime))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContainsFold(FieldStatus, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContainsFold(FieldAppID, v))
}

// AppKeyEQ applies the EQ predicate on the "app_key" field.
func AppKeyEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldAppKey, v))
}

// AppKeyNEQ applies the NEQ predicate on the "app_key" field.
func AppKeyNEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldAppKey, v))
}

// AppKeyIn applies the In predicate on the "app_key" field.
func AppKeyIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldAppKey, vs...))
}

// AppKeyNotIn applies the NotIn predicate on the "app_key" field.
func AppKeyNotIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldAppKey, vs...))
}

// AppKeyGT applies the GT predicate on the "app_key" field.
func AppKeyGT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldAppKey, v))
}

// AppKeyGTE applies the GTE predicate on the "app_key" field.
func AppKeyGTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldAppKey, v))
}

// AppKeyLT applies the LT predicate on the "app_key" field.
func AppKeyLT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldAppKey, v))
}

// AppKeyLTE applies the LTE predicate on the "app_key" field.
func AppKeyLTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldAppKey, v))
}

// AppKeyContains applies the Contains predicate on the "app_key" field.
func AppKeyContains(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContains(FieldAppKey, v))
}

// AppKeyHasPrefix applies the HasPrefix predicate on the "app_key" field.
func AppKeyHasPrefix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasPrefix(FieldAppKey, v))
}

// AppKeyHasSuffix applies the HasSuffix predicate on the "app_key" field.
func AppKeyHasSuffix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasSuffix(FieldAppKey, v))
}

// AppKeyIsNil applies the IsNil predicate on the "app_key" field.
func AppKeyIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldAppKey))
}

// AppKeyNotNil applies the NotNil predicate on the "app_key" field.
func AppKeyNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldAppKey))
}

// AppKeyEqualFold applies the EqualFold predicate on the "app_key" field.
func AppKeyEqualFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEqualFold(FieldAppKey, v))
}

// AppKeyContainsFold applies the ContainsFold predicate on the "app_key" field.
func AppKeyContainsFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContainsFold(FieldAppKey, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldContainsFold(FieldRemark, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotIn(FieldCreatorID, vs...))
}

// CreatorIDGT applies the GT predicate on the "creator_id" field.
func CreatorIDGT(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGT(FieldCreatorID, v))
}

// CreatorIDGTE applies the GTE predicate on the "creator_id" field.
func CreatorIDGTE(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldGTE(FieldCreatorID, v))
}

// CreatorIDLT applies the LT predicate on the "creator_id" field.
func CreatorIDLT(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLT(FieldCreatorID, v))
}

// CreatorIDLTE applies the LTE predicate on the "creator_id" field.
func CreatorIDLTE(v uint32) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldLTE(FieldCreatorID, v))
}

// CreatorIDIsNil applies the IsNil predicate on the "creator_id" field.
func CreatorIDIsNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldIsNull(FieldCreatorID))
}

// CreatorIDNotNil applies the NotNil predicate on the "creator_id" field.
func CreatorIDNotNil() predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(sql.FieldNotNull(FieldCreatorID))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RealtimeWarehousing) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RealtimeWarehousing) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(func(s *sql.Selector) {
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
func Not(p predicate.RealtimeWarehousing) predicate.RealtimeWarehousing {
	return predicate.RealtimeWarehousing(func(s *sql.Selector) {
		p(s.Not())
	})
}
