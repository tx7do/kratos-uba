package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kratos-bi/pkg/util/entgo/mixin"
)

// RealtimeWarehousing holds the schema definition for the RealtimeWarehousing entity.
type RealtimeWarehousing struct {
	ent.Schema
}

func (RealtimeWarehousing) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "realtime_warehousing",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the RealtimeWarehousing.
func (RealtimeWarehousing) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("应用名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("status").
			Comment("应用状态").
			Optional().
			Nillable(),

		field.String("app_id").
			Comment("应用ID").
			Optional().
			Nillable(),

		field.String("app_key").
			Comment("应用密钥").
			Optional().
			Nillable(),

		field.String("remark").
			Comment("备注").
			Optional().
			Nillable(),

		field.Uint32("creator_id").
			Comment("创建者ID").
			Optional().
			Nillable(),
	}
}

// Mixin of the RealtimeWarehousing.
func (RealtimeWarehousing) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the RealtimeWarehousing.
func (RealtimeWarehousing) Edges() []ent.Edge {
	return nil
}
