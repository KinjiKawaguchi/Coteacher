package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teacher", Teacher.Type).Unique(),
		edge.To("student", Student.Type).Unique(),
	}
}