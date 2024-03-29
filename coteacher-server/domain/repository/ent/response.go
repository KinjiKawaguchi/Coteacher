// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/form"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/student"
	"github.com/google/uuid"
)

// Response is the model entity for the Response schema.
type Response struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// StudentID holds the value of the "student_id" field.
	StudentID uuid.UUID `json:"student_id,omitempty"`
	// FormID holds the value of the "form_id" field.
	FormID uuid.UUID `json:"form_id,omitempty"`
	// AiResponse holds the value of the "ai_response" field.
	AiResponse string `json:"ai_response,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResponseQuery when eager-loading is set.
	Edges        ResponseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ResponseEdges holds the relations/edges for other nodes in the graph.
type ResponseEdges struct {
	// Student holds the value of the student edge.
	Student *Student `json:"student,omitempty"`
	// Form holds the value of the form edge.
	Form *Form `json:"form,omitempty"`
	// Answer holds the value of the answer edge.
	Answer []*Answer `json:"answer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResponseEdges) StudentOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// FormOrErr returns the Form value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResponseEdges) FormOrErr() (*Form, error) {
	if e.loadedTypes[1] {
		if e.Form == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: form.Label}
		}
		return e.Form, nil
	}
	return nil, &NotLoadedError{edge: "form"}
}

// AnswerOrErr returns the Answer value or an error if the edge
// was not loaded in eager-loading.
func (e ResponseEdges) AnswerOrErr() ([]*Answer, error) {
	if e.loadedTypes[2] {
		return e.Answer, nil
	}
	return nil, &NotLoadedError{edge: "answer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Response) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case response.FieldAiResponse:
			values[i] = new(sql.NullString)
		case response.FieldCreatedAt, response.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case response.FieldID, response.FieldStudentID, response.FieldFormID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Response fields.
func (r *Response) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case response.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case response.FieldStudentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field student_id", values[i])
			} else if value != nil {
				r.StudentID = *value
			}
		case response.FieldFormID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field form_id", values[i])
			} else if value != nil {
				r.FormID = *value
			}
		case response.FieldAiResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ai_response", values[i])
			} else if value.Valid {
				r.AiResponse = value.String
			}
		case response.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case response.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Response.
// This includes values selected through modifiers, order, etc.
func (r *Response) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryStudent queries the "student" edge of the Response entity.
func (r *Response) QueryStudent() *StudentQuery {
	return NewResponseClient(r.config).QueryStudent(r)
}

// QueryForm queries the "form" edge of the Response entity.
func (r *Response) QueryForm() *FormQuery {
	return NewResponseClient(r.config).QueryForm(r)
}

// QueryAnswer queries the "answer" edge of the Response entity.
func (r *Response) QueryAnswer() *AnswerQuery {
	return NewResponseClient(r.config).QueryAnswer(r)
}

// Update returns a builder for updating this Response.
// Note that you need to call Response.Unwrap() before calling this method if this Response
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Response) Update() *ResponseUpdateOne {
	return NewResponseClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Response entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Response) Unwrap() *Response {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Response is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Response) String() string {
	var builder strings.Builder
	builder.WriteString("Response(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("student_id=")
	builder.WriteString(fmt.Sprintf("%v", r.StudentID))
	builder.WriteString(", ")
	builder.WriteString("form_id=")
	builder.WriteString(fmt.Sprintf("%v", r.FormID))
	builder.WriteString(", ")
	builder.WriteString("ai_response=")
	builder.WriteString(r.AiResponse)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Responses is a parsable slice of Response.
type Responses []*Response
