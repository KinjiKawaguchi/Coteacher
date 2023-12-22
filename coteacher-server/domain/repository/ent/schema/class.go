package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Class struct {
	ent.Schema
}

func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("class_invitation_codes", ClassInvitationCode.Type),
		edge.To("student_classes", StudentClass.Type), // Add this line
	}
}
