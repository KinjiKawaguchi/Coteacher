package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClassInvitationCode holds the schema definition for the ClassInvitationCode entity.
type ClassInvitationCode struct {
	ent.Schema
}

// Fields of the ClassInvitationCode.
func (ClassInvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.UUID("class_id", uuid.UUID{}),
		field.String("invitation_code").Unique().NotEmpty(),
		field.Time("expiration_date").Optional(),
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
