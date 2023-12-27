package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
		field.String("teacher_id"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).
			Ref("classes").
			Field("teacher_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("classStudents", StudentClass.Type), // このエッジを追加
		edge.To("invitationCodes", ClassInvitationCode.Type),
	}
}
