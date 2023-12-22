// Code generated by ent, DO NOT EDIT.

package ent

import (
	"coteacher/domain/repository/ent/class"
	"coteacher/domain/repository/ent/studentclass"
	"coteacher/domain/repository/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// StudentClass is the model entity for the StudentClass schema.
type StudentClass struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentClassQuery when eager-loading is set.
	Edges                 StudentClassEdges `json:"edges"`
	class_student_classes *string
	user_student_classes  *string
	selectValues          sql.SelectValues
}

// StudentClassEdges holds the relations/edges for other nodes in the graph.
type StudentClassEdges struct {
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentClassEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[0] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentClassEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StudentClass) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case studentclass.FieldID:
			values[i] = new(sql.NullInt64)
		case studentclass.ForeignKeys[0]: // class_student_classes
			values[i] = new(sql.NullString)
		case studentclass.ForeignKeys[1]: // user_student_classes
			values[i] = new(sql.NullString)
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
		case studentclass.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_student_classes", values[i])
			} else if value.Valid {
				sc.class_student_classes = new(string)
				*sc.class_student_classes = value.String
			}
		case studentclass.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_student_classes", values[i])
			} else if value.Valid {
				sc.user_student_classes = new(string)
				*sc.user_student_classes = value.String
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

// QueryClass queries the "class" edge of the StudentClass entity.
func (sc *StudentClass) QueryClass() *ClassQuery {
	return NewStudentClassClient(sc.config).QueryClass(sc)
}

// QueryUser queries the "user" edge of the StudentClass entity.
func (sc *StudentClass) QueryUser() *UserQuery {
	return NewStudentClassClient(sc.config).QueryUser(sc)
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
	builder.WriteString(fmt.Sprintf("id=%v", sc.ID))
	builder.WriteByte(')')
	return builder.String()
}

// StudentClasses is a parsable slice of StudentClass.
type StudentClasses []*StudentClass
