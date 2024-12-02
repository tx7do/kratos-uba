package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

func (Application) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "application",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
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

		field.Uint32("owner_id").
			Comment("拥有者ID").
			Optional().
			Nillable(),

		field.Uint32("keep_month").
			Comment("数据保存多少个月").
			Optional().
			Nillable(),
	}
}

// Mixin of the Application.
func (Application) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return nil
}
