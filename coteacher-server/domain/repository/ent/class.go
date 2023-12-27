// Code generated by ent, DO NOT EDIT.

package ent

import (
	"coteacher/domain/repository/ent/class"
	"coteacher/domain/repository/ent/teacher"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TeacherID holds the value of the "teacher_id" field.
	TeacherID string `json:"teacher_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges        ClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Teacher holds the value of the teacher edge.
	Teacher *Teacher `json:"teacher,omitempty"`
	// ClassStudents holds the value of the classStudents edge.
	ClassStudents []*StudentClass `json:"classStudents,omitempty"`
	// InvitationCodes holds the value of the invitationCodes edge.
	InvitationCodes []*ClassInvitationCode `json:"invitationCodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TeacherOrErr returns the Teacher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) TeacherOrErr() (*Teacher, error) {
	if e.loadedTypes[0] {
		if e.Teacher == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.Teacher, nil
	}
	return nil, &NotLoadedError{edge: "teacher"}
}

// ClassStudentsOrErr returns the ClassStudents value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) ClassStudentsOrErr() ([]*StudentClass, error) {
	if e.loadedTypes[1] {
		return e.ClassStudents, nil
	}
	return nil, &NotLoadedError{edge: "classStudents"}
}

// InvitationCodesOrErr returns the InvitationCodes value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) InvitationCodesOrErr() ([]*ClassInvitationCode, error) {
	if e.loadedTypes[2] {
		return e.InvitationCodes, nil
	}
	return nil, &NotLoadedError{edge: "invitationCodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldID, class.FieldName, class.FieldTeacherID:
			values[i] = new(sql.NullString)
		case class.FieldCreatedAt, class.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case class.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case class.FieldTeacherID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_id", values[i])
			} else if value.Valid {
				c.TeacherID = value.String
			}
		case class.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case class.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Class.
// This includes values selected through modifiers, order, etc.
func (c *Class) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryTeacher queries the "teacher" edge of the Class entity.
func (c *Class) QueryTeacher() *TeacherQuery {
	return NewClassClient(c.config).QueryTeacher(c)
}

// QueryClassStudents queries the "classStudents" edge of the Class entity.
func (c *Class) QueryClassStudents() *StudentClassQuery {
	return NewClassClient(c.config).QueryClassStudents(c)
}

// QueryInvitationCodes queries the "invitationCodes" edge of the Class entity.
func (c *Class) QueryInvitationCodes() *ClassInvitationCodeQuery {
	return NewClassClient(c.config).QueryInvitationCodes(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("teacher_id=")
	builder.WriteString(c.TeacherID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class
