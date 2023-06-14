package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kratos-bi/pkg/util/entgo/mixin"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

func (Attribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "attribute",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("属性的名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("show_name").
			Comment("展示的名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("status").
			Comment("属性状态").
			Optional().
			Nillable(),

		field.Uint8("attribute_type").
			Comment("属性类型").
			Optional().
			Nillable(),

		field.Uint8("attribute_source").
			Comment("属性源").
			Optional().
			Nillable(),

		field.Uint32("app_id").
			Comment("应用ID").
			Optional().
			Nillable(),

		field.String("data_type").
			Comment("数据类型").
			Optional().
			Nillable(),
	}
}

// Mixin of the Attribute.
func (Attribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return nil
}
