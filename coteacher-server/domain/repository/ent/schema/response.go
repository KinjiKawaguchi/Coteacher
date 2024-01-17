package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Response holds the schema definition for the Response entity.
type Response struct {
	ent.Schema
}

// Fields of the Response.
func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("student_id", uuid.UUID{}).StorageKey("student_id"),
		field.UUID("form_id", uuid.UUID{}).StorageKey("form_id"),
		field.Text("ai_response"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Response.
func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("responses").
			Field("student_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.From("form", Form.Type).
			Ref("responses").
			Field("form_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("answer", Answer.Type), // このエッジを追加
	}
}
