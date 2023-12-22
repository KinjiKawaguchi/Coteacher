package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ClassInvitationCode struct {
	ent.Schema
}

func (ClassInvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("invitation_code").Unique(),
		field.Time("expiration_date"),
		field.Bool("is_active"),
	}
}

func (ClassInvitationCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("class_invitation_codes").Unique(),
	}
}
