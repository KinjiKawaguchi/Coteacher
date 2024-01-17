package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Form struct {
	ent.Schema
}

func (Form) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("class_id", uuid.UUID{}).StorageKey("class_id"),
		field.String("name"),
		field.Text("description"),
		field.Int("usage_limit"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (Form) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("forms").
			Field("class_id"). // ここでリンクする
			Unique().
			Required(), // ここで必須にする
		edge.To("questions", Question.Type),
		edge.To("responses", Response.Type),
	}
}
