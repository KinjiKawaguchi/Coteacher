// Code generated by ent, DO NOT EDIT.

package ent

import (
	"coteacher/domain/repository/ent/teacher"
	"coteacher/domain/repository/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Teacher is the model entity for the Teacher schema.
type Teacher struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeacherQuery when eager-loading is set.
	Edges        TeacherEdges `json:"edges"`
	user_teacher *uuid.UUID
	selectValues sql.SelectValues
}

// TeacherEdges holds the relations/edges for other nodes in the graph.
type TeacherEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Classes holds the value of the classes edge.
	Classes []*Class `json:"classes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeacherEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ClassesOrErr returns the Classes value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) ClassesOrErr() ([]*Class, error) {
	if e.loadedTypes[1] {
		return e.Classes, nil
	}
	return nil, &NotLoadedError{edge: "classes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teacher) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			values[i] = new(uuid.UUID)
		case teacher.ForeignKeys[0]: // user_teacher
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teacher fields.
func (t *Teacher) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case teacher.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_teacher", values[i])
			} else if value.Valid {
				t.user_teacher = new(uuid.UUID)
				*t.user_teacher = *value.S.(*uuid.UUID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Teacher.
// This includes values selected through modifiers, order, etc.
func (t *Teacher) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Teacher entity.
func (t *Teacher) QueryUser() *UserQuery {
	return NewTeacherClient(t.config).QueryUser(t)
}

// QueryClasses queries the "classes" edge of the Teacher entity.
func (t *Teacher) QueryClasses() *ClassQuery {
	return NewTeacherClient(t.config).QueryClasses(t)
}

// Update returns a builder for updating this Teacher.
// Note that you need to call Teacher.Unwrap() before calling this method if this Teacher
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teacher) Update() *TeacherUpdateOne {
	return NewTeacherClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Teacher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teacher) Unwrap() *Teacher {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teacher is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teacher) String() string {
	var builder strings.Builder
	builder.WriteString("Teacher(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Teachers is a parsable slice of Teacher.
type Teachers []*Teacher