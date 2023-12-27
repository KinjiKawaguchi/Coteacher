package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClassInvitationCode holds the schema definition for the ClassInvitationCode entity.
type ClassInvitationCode struct {
	ent.Schema
}

// Fields of the ClassInvitationCode.
func (ClassInvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("class_id"),
		field.String("invitation_code").Unique(),
		field.Time("expiration_date"),
		field.Bool("is_active"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the ClassInvitationCode.
func (ClassInvitationCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("invitationCodes").
			Field("class_id").
			Required().
			Unique(),
	}
}
