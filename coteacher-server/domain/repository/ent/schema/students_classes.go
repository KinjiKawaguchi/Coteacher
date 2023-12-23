package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentClass holds the schema definition for the StudentClass entity.
type StudentClass struct {
	ent.Schema
}

// Fields of the StudentClass.
func (StudentClass) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("student_id", uuid.UUID{}),
		field.UUID("class_id", uuid.UUID{}),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the StudentClass.
func (StudentClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("studentClasses").
			Field("student_id"). // ここでリンクする
			Required().
			Unique(),
		edge.From("class", Class.Type).
			Ref("classStudents"). // Class から StudentClass への参照、以前の "classStudentClasses" から変更
			Field("class_id").
			Unique().
			Required(),
	}
}
