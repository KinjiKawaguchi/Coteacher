package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Teacher holds the schema definition for the Teacher entity.
type Student struct {
	ent.Schema
}

// Fields of the Teacher.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("id").Unique(),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("student").
			Unique().
			Required(),
		edge.To("studentClasses", StudentClass.Type), // このエッジを追加
	}
}
