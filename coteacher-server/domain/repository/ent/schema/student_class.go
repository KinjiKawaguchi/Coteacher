package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type StudentClass struct {
	ent.Schema
}

func (StudentClass) Fields() []ent.Field {
	return nil
}

func (StudentClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("student_classes").Unique(),
		edge.From("user", User.Type).Ref("student_classes").Unique(),
	}
}
