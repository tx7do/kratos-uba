package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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

		field.String("event_name").
			Comment("事件的名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("report_data").
			Comment("报告的数据").
			Optional().
			Nillable(),

		field.Time("create_time").
			Comment("记录创建时间").
			Optional().
			Nillable(),
	}
}
