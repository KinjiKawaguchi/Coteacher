package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("question_id", uuid.UUID{}).StorageKey("question_id"),
		field.UUID("response_id", uuid.UUID{}).StorageKey("response_id"),
		field.Int("order"),
		field.Text("answer_text").Optional(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).
			Ref("answer").
			Field("question_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.From("response", Response.Type).
			Ref("answer").
			Field("response_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("selectedOption", SelectedOption.Type), // このエッジを追加
	}
}
