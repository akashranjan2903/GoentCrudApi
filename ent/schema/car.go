package schema

import (
	"context"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"

	// "entgo.io/ent/entc/integration/edgeschema/ent/hook"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	gen "github.com/gocrud/ent"
	"github.com/gocrud/ent/hook"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{

		field.String("model"),
	}
}
func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}

// Index
func (Car) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("model").
			Edges("user").
			Unique(),
	}
}
func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CAR"},
	}
}

func (Car) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.CarFunc(
				func(ctx context.Context, m *gen.CarMutation) (ent.Value, error) {
					r, _ := m.Model()
					m.SetModel(strings.ToLower(r))
					return next.Mutate(ctx, m)
				},
			)
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}
