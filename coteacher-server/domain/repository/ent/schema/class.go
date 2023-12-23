package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }), field.String("name").NotEmpty(),
		field.UUID("teacher_id", uuid.UUID{}),
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
