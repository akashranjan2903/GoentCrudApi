package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").MaxLen(255).Optional(),
		field.String("last_name").MaxLen(255).Optional(),
		field.String("email").Unique().MaxLen(255),
		field.String("password"),
		field.Bool("is_active").Default(false).Nillable(),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("user").Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
