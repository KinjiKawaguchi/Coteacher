// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/student"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"
	"github.com/google/uuid"
)

// StudentClass is the model entity for the StudentClass schema.
type StudentClass struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StudentID holds the value of the "student_id" field.
	StudentID uuid.UUID `json:"student_id,omitempty"`
	// ClassID holds the value of the "class_id" field.
	ClassID uuid.UUID `json:"class_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentClassQuery when eager-loading is set.
	Edges        StudentClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StudentClassEdges holds the relations/edges for other nodes in the graph.
type StudentClassEdges struct {
	// Student holds the value of the student edge.
	Student *Student `json:"student,omitempty"`
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentClassEdges) StudentOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentClassEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[1] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StudentClass) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case studentclass.FieldID:
			values[i] = new(sql.NullInt64)
		case studentclass.FieldCreatedAt, studentclass.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case studentclass.FieldStudentID, studentclass.FieldClassID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudentClass fields.
func (sc *StudentClass) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studentclass.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case studentclass.FieldStudentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field student_id", values[i])
			} else if value != nil {
				sc.StudentID = *value
			}
		case studentclass.FieldClassID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value != nil {
				sc.ClassID = *value
			}
		case studentclass.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case studentclass.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StudentClass.
// This includes values selected through modifiers, order, etc.
func (sc *StudentClass) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QueryStudent queries the "student" edge of the StudentClass entity.
func (sc *StudentClass) QueryStudent() *StudentQuery {
	return NewStudentClassClient(sc.config).QueryStudent(sc)
}

// QueryClass queries the "class" edge of the StudentClass entity.
func (sc *StudentClass) QueryClass() *ClassQuery {
	return NewStudentClassClient(sc.config).QueryClass(sc)
}

// Update returns a builder for updating this StudentClass.
// Note that you need to call StudentClass.Unwrap() before calling this method if this StudentClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *StudentClass) Update() *StudentClassUpdateOne {
	return NewStudentClassClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the StudentClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *StudentClass) Unwrap() *StudentClass {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: StudentClass is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *StudentClass) String() string {
	var builder strings.Builder
	builder.WriteString("StudentClass(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("student_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.StudentID))
	builder.WriteString(", ")
	builder.WriteString("class_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ClassID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// StudentClasses is a parsable slice of StudentClass.
type StudentClasses []*StudentClass
