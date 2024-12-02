package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// AcceptanceStatus holds the schema definition for the AcceptanceStatus entity.
type AcceptanceStatus struct {
	ent.Schema
}

func (AcceptanceStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "acceptance_status",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the AcceptanceStatus.
func (AcceptanceStatus) Fields() []ent.Field {
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

// Mixin of the AcceptanceStatus.
func (AcceptanceStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the AcceptanceStatus.
func (AcceptanceStatus) Edges() []ent.Edge {
	return nil
}
