package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// DebugDevice holds the schema definition for the DebugDevice entity.
type DebugDevice struct {
	ent.Schema
}

func (DebugDevice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "debug_device",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the DebugDevice.
func (DebugDevice) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_id").
			Comment("设备ID").
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("app_id").
			Comment("应用ID").
			Optional().
			Nillable(),

		field.Uint32("creator_id").
			Comment("创建者ID").
			Optional().
			Nillable(),

		field.String("remark").
			Comment("备注").
			Optional().
			Nillable(),
	}
}

// Mixin of the DebugDevice.
func (DebugDevice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the DebugDevice.
func (DebugDevice) Edges() []ent.Edge {
	return nil
}
