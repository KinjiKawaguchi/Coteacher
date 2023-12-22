package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
		field.String("email").Unique(),
		field.Enum("UserType").Values("teacher", "student"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student_classes", StudentClass.Type),
		edge.To("classes", Class.Type),
	}
}
