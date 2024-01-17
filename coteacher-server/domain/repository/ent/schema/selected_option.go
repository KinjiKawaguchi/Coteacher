package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"github.com/google/uuid"
)

// SelectedOption holds the schema definition for the SelectedOption entity.
type SelectedOption struct {
	ent.Schema
}

// Fields of the SelectedOption.
func (SelectedOption) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("option_id", uuid.UUID{}).StorageKey("option_id"),
		field.UUID("answer_id", uuid.UUID{}).StorageKey("answer_id"),
	}
}

// Edges of the SelectedOption.
func (SelectedOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("answer", Answer.Type).
			Ref("selectedOption").
			Field("answer_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.From("option", QuestionOption.Type).
			Ref("selectedOption").
			Field("option_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
	}
}
