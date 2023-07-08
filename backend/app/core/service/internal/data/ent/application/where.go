// Code generated by ent, DO NOT EDIT.

package application

import (
	"kratos-uba/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldDeleteTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldStatus, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldAppID, v))
}

// AppKey applies equality check predicate on the "app_key" field. It's identical to AppKeyEQ.
func AppKey(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldAppKey, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldRemark, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreatorID, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldOwnerID, v))
}

// KeepMonth applies equality check predicate on the "keep_month" field. It's identical to KeepMonthEQ.
func KeepMonth(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldKeepMonth, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v int64) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...int64) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v int64) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v int64) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v int64) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldDeleteTime))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Application {
	return predicate.Application(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Application {
	return predicate.Application(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Application {
	return predicate.Application(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Application {
	return predicate.Application(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Application {
	return predicate.Application(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Application {
	return predicate.Application(sql.FieldContainsFold(FieldStatus, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.Application {
	return predicate.Application(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.Application {
	return predicate.Application(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.Application {
	return predicate.Application(sql.FieldContainsFold(FieldAppID, v))
}

// AppKeyEQ applies the EQ predicate on the "app_key" field.
func AppKeyEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldAppKey, v))
}

// AppKeyNEQ applies the NEQ predicate on the "app_key" field.
func AppKeyNEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldAppKey, v))
}

// AppKeyIn applies the In predicate on the "app_key" field.
func AppKeyIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldAppKey, vs...))
}

// AppKeyNotIn applies the NotIn predicate on the "app_key" field.
func AppKeyNotIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldAppKey, vs...))
}

// AppKeyGT applies the GT predicate on the "app_key" field.
func AppKeyGT(v string) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldAppKey, v))
}

// AppKeyGTE applies the GTE predicate on the "app_key" field.
func AppKeyGTE(v string) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldAppKey, v))
}

// AppKeyLT applies the LT predicate on the "app_key" field.
func AppKeyLT(v string) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldAppKey, v))
}

// AppKeyLTE applies the LTE predicate on the "app_key" field.
func AppKeyLTE(v string) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldAppKey, v))
}

// AppKeyContains applies the Contains predicate on the "app_key" field.
func AppKeyContains(v string) predicate.Application {
	return predicate.Application(sql.FieldContains(FieldAppKey, v))
}

// AppKeyHasPrefix applies the HasPrefix predicate on the "app_key" field.
func AppKeyHasPrefix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasPrefix(FieldAppKey, v))
}

// AppKeyHasSuffix applies the HasSuffix predicate on the "app_key" field.
func AppKeyHasSuffix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasSuffix(FieldAppKey, v))
}

// AppKeyIsNil applies the IsNil predicate on the "app_key" field.
func AppKeyIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldAppKey))
}

// AppKeyNotNil applies the NotNil predicate on the "app_key" field.
func AppKeyNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldAppKey))
}

// AppKeyEqualFold applies the EqualFold predicate on the "app_key" field.
func AppKeyEqualFold(v string) predicate.Application {
	return predicate.Application(sql.FieldEqualFold(FieldAppKey, v))
}

// AppKeyContainsFold applies the ContainsFold predicate on the "app_key" field.
func AppKeyContainsFold(v string) predicate.Application {
	return predicate.Application(sql.FieldContainsFold(FieldAppKey, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Application {
	return predicate.Application(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Application {
	return predicate.Application(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Application {
	return predicate.Application(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Application {
	return predicate.Application(sql.FieldContainsFold(FieldRemark, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldCreatorID, vs...))
}

// CreatorIDGT applies the GT predicate on the "creator_id" field.
func CreatorIDGT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldCreatorID, v))
}

// CreatorIDGTE applies the GTE predicate on the "creator_id" field.
func CreatorIDGTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldCreatorID, v))
}

// CreatorIDLT applies the LT predicate on the "creator_id" field.
func CreatorIDLT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldCreatorID, v))
}

// CreatorIDLTE applies the LTE predicate on the "creator_id" field.
func CreatorIDLTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldCreatorID, v))
}

// CreatorIDIsNil applies the IsNil predicate on the "creator_id" field.
func CreatorIDIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldCreatorID))
}

// CreatorIDNotNil applies the NotNil predicate on the "creator_id" field.
func CreatorIDNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldCreatorID))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldOwnerID))
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldOwnerID))
}

// KeepMonthEQ applies the EQ predicate on the "keep_month" field.
func KeepMonthEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldKeepMonth, v))
}

// KeepMonthNEQ applies the NEQ predicate on the "keep_month" field.
func KeepMonthNEQ(v uint32) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldKeepMonth, v))
}

// KeepMonthIn applies the In predicate on the "keep_month" field.
func KeepMonthIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldKeepMonth, vs...))
}

// KeepMonthNotIn applies the NotIn predicate on the "keep_month" field.
func KeepMonthNotIn(vs ...uint32) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldKeepMonth, vs...))
}

// KeepMonthGT applies the GT predicate on the "keep_month" field.
func KeepMonthGT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldKeepMonth, v))
}

// KeepMonthGTE applies the GTE predicate on the "keep_month" field.
func KeepMonthGTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldKeepMonth, v))
}

// KeepMonthLT applies the LT predicate on the "keep_month" field.
func KeepMonthLT(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldKeepMonth, v))
}

// KeepMonthLTE applies the LTE predicate on the "keep_month" field.
func KeepMonthLTE(v uint32) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldKeepMonth, v))
}

// KeepMonthIsNil applies the IsNil predicate on the "keep_month" field.
func KeepMonthIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldKeepMonth))
}

// KeepMonthNotNil applies the NotNil predicate on the "keep_month" field.
func KeepMonthNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldKeepMonth))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
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
func Not(p predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		p(s.Not())
	})
}
