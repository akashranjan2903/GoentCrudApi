package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
		field.UUID("created_by", uuid.UUID{}).Optional().Nillable(),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable(),
		field.UUID("deleted_by", uuid.UUID{}).Optional().Nillable(),
	}
}
