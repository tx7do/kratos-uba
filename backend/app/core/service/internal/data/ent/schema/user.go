package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kratos-uba/pkg/util/entgo/mixin"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").
			Comment("用户名").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),

		field.String("password").
			Comment("登陆密码").
			MaxLen(255).
			Optional().
			Nillable().
			NotEmpty(),

		field.String("nick_name").
			Comment("昵称").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("real_name").
			Comment("姓名").
			MaxLen(128).
			Optional().
			Nillable(),

		field.String("email").
			Comment("电子邮箱").
			MaxLen(127).
			Optional().
			Nillable(),

		field.String("phone").
			Comment("手机号码").
			MaxLen(127).
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("头像").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Uint32("role_id").
			Comment("角色ID").
			Default(0).
			Optional().
			Nillable(),

		field.Enum("authority").
			Comment("授权").
			Optional().
			Nillable().
			Values(
				"SYS_ADMIN",
				"CUSTOMER_USER",
				"GUEST_USER",
				"REFRESH_TOKEN",
			).
			Default("CUSTOMER_USER"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "user_name").Unique(),
	}
}
