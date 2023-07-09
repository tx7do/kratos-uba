package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kratos-uba/pkg/util/entgo/mixin"
)

// MetaEvent holds the schema definition for the MetaEvent entity.
type MetaEvent struct {
	ent.Schema
}

func (MetaEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "meta_event",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the MetaEvent.
func (MetaEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_name").
			Comment("事件名").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("show_name").
			Comment("显示名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("app_id").
			Comment("应用ID").
			Optional().
			Nillable(),

		field.Uint32("yesterday_count").
			Comment("计数").
			Optional().
			Nillable(),
	}
}

// Mixin of the MetaEvent.
func (MetaEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}
