package schema

import (
	"time"

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
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.UUID("class_id", uuid.UUID{}).StorageKey("class_id"),
		field.String("invitation_code").Unique(),
		field.Time("expiration_date").Default(nil),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
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
