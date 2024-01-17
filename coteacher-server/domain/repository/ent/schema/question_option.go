package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QuestionOption holds the schema definition for the QuestionOption entity.
type QuestionOption struct {
	ent.Schema
}

// Fields of the QuestionOption.
func (QuestionOption) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("question_id", uuid.UUID{}).StorageKey("question_id"),
		field.Text("option_text"),
		field.Int("order"),
	}
}

// Edges of the QuestionOption.
func (QuestionOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).
			Ref("questionOption").
			Field("question_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("selectedOption", SelectedOption.Type),
	}
}
