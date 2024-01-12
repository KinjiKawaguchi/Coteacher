package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TextQuestion holds the schema definition for the TextQuestion entity.
type TextQuestion struct {
	ent.Schema
}

// Fields of the TextQuestion.
func (TextQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("question_id", uuid.UUID{}).StorageKey("question_id"),
		field.Int("max_length"),
	}
}

// Edges of the TextQuestion.
func (TextQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).
			Ref("textQuestion").
			Field("question_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
	}
}
