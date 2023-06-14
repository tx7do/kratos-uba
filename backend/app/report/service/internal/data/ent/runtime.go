// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-bi/app/report/service/internal/data/ent/acceptancestatus"
	"kratos-bi/app/report/service/internal/data/ent/realtimewarehousing"
	"kratos-bi/app/report/service/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	acceptancestatusMixin := schema.AcceptanceStatus{}.Mixin()
	acceptancestatusMixinFields0 := acceptancestatusMixin[0].Fields()
	_ = acceptancestatusMixinFields0
	acceptancestatusMixinFields1 := acceptancestatusMixin[1].Fields()
	_ = acceptancestatusMixinFields1
	acceptancestatusFields := schema.AcceptanceStatus{}.Fields()
	_ = acceptancestatusFields
	// acceptancestatusDescCreateTime is the schema descriptor for create_time field.
	acceptancestatusDescCreateTime := acceptancestatusMixinFields1[0].Descriptor()
	// acceptancestatus.DefaultCreateTime holds the default value on creation for the create_time field.
	acceptancestatus.DefaultCreateTime = acceptancestatusDescCreateTime.Default.(func() int64)
	// acceptancestatusDescUpdateTime is the schema descriptor for update_time field.
	acceptancestatusDescUpdateTime := acceptancestatusMixinFields1[1].Descriptor()
	// acceptancestatus.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	acceptancestatus.UpdateDefaultUpdateTime = acceptancestatusDescUpdateTime.UpdateDefault.(func() int64)
	// acceptancestatusDescName is the schema descriptor for name field.
	acceptancestatusDescName := acceptancestatusFields[0].Descriptor()
	// acceptancestatus.NameValidator is a validator for the "name" field. It is called by the builders before save.
	acceptancestatus.NameValidator = acceptancestatusDescName.Validators[0].(func(string) error)
	// acceptancestatusDescID is the schema descriptor for id field.
	acceptancestatusDescID := acceptancestatusMixinFields0[0].Descriptor()
	// acceptancestatus.IDValidator is a validator for the "id" field. It is called by the builders before save.
	acceptancestatus.IDValidator = acceptancestatusDescID.Validators[0].(func(uint32) error)
	realtimewarehousingMixin := schema.RealtimeWarehousing{}.Mixin()
	realtimewarehousingMixinFields0 := realtimewarehousingMixin[0].Fields()
	_ = realtimewarehousingMixinFields0
	realtimewarehousingMixinFields1 := realtimewarehousingMixin[1].Fields()
	_ = realtimewarehousingMixinFields1
	realtimewarehousingFields := schema.RealtimeWarehousing{}.Fields()
	_ = realtimewarehousingFields
	// realtimewarehousingDescCreateTime is the schema descriptor for create_time field.
	realtimewarehousingDescCreateTime := realtimewarehousingMixinFields1[0].Descriptor()
	// realtimewarehousing.DefaultCreateTime holds the default value on creation for the create_time field.
	realtimewarehousing.DefaultCreateTime = realtimewarehousingDescCreateTime.Default.(func() int64)
	// realtimewarehousingDescUpdateTime is the schema descriptor for update_time field.
	realtimewarehousingDescUpdateTime := realtimewarehousingMixinFields1[1].Descriptor()
	// realtimewarehousing.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	realtimewarehousing.UpdateDefaultUpdateTime = realtimewarehousingDescUpdateTime.UpdateDefault.(func() int64)
	// realtimewarehousingDescName is the schema descriptor for name field.
	realtimewarehousingDescName := realtimewarehousingFields[0].Descriptor()
	// realtimewarehousing.NameValidator is a validator for the "name" field. It is called by the builders before save.
	realtimewarehousing.NameValidator = realtimewarehousingDescName.Validators[0].(func(string) error)
	// realtimewarehousingDescID is the schema descriptor for id field.
	realtimewarehousingDescID := realtimewarehousingMixinFields0[0].Descriptor()
	// realtimewarehousing.IDValidator is a validator for the "id" field. It is called by the builders before save.
	realtimewarehousing.IDValidator = realtimewarehousingDescID.Validators[0].(func(uint32) error)
}
