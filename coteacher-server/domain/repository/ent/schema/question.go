package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Question struct {
	ent.Schema
}

func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("form_id", uuid.UUID{}).StorageKey("form_id"),
		field.Enum("question_type").
			Values("checkbox", "list", "radio", "multiple_choice", "paragraph_text", "text", "unspecified"),
		//field.Enum("question_type").
		//Values("checkbox", "checkbox_grid", "date", "datetime", "duration", "grid", "list",
		//"multiple_choice", "paragraph_text", "scale", "text", "time", "file_upload"),
		field.Text("question_text"),
		field.Bool("is_required"),
		field.Bool("for_ai_processing"),
		field.Int("order"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("form", Form.Type).
			Ref("questions").
			Field("form_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("textQuestion", TextQuestion.Type),
		edge.To("questionOption", QuestionOption.Type),
		edge.To("answer", Answer.Type),
	}
}
