package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
		field.Int64("id").
			Comment("id").
			StructTag(`json:"id,omitempty"`).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigserial",
			}).
			Positive().
			Immutable().
			Unique(),

		field.String("data_name").
			Comment("数据名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("report_type").
			Comment("报告类型").
			Optional().
			Nillable(),

		field.String("report_data").
			Comment("报告数据").
			Optional().
			Nillable(),

		field.Int32("status").
			Comment("状态").
			Optional().
			Nillable(),

		field.String("error_reason").
			Comment("错误原因").
			Optional().
			Nillable(),

		field.String("error_handling").
			Comment("错误处理").
			Optional().
			Nillable(),

		field.String("part_event").
			Comment("").
			Optional().
			Nillable(),

		field.Time("part_date").
			Comment("记录时间").
			Optional().
			Nillable(),
	}
}
