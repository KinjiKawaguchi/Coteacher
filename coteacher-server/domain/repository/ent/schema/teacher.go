package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("id").Unique(),
	}
}

func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("teacher").
			Unique().
			Required(),
		edge.To("classes", Class.Type),
	}
}
