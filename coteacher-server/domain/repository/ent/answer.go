// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/answer"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/google/uuid"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// QuestionID holds the value of the "question_id" field.
	QuestionID uuid.UUID `json:"question_id,omitempty"`
	// ResponseID holds the value of the "response_id" field.
	ResponseID uuid.UUID `json:"response_id,omitempty"`
	// AnswerText holds the value of the "answer_text" field.
	AnswerText string `json:"answer_text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnswerQuery when eager-loading is set.
	Edges        AnswerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AnswerEdges holds the relations/edges for other nodes in the graph.
type AnswerEdges struct {
	// Question holds the value of the question edge.
	Question *Question `json:"question,omitempty"`
	// Response holds the value of the response edge.
	Response *Response `json:"response,omitempty"`
	// SelectedOption holds the value of the selectedOption edge.
	SelectedOption []*SelectedOption `json:"selectedOption,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) QuestionOrErr() (*Question, error) {
	if e.loadedTypes[0] {
		if e.Question == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: question.Label}
		}
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "question"}
}

// ResponseOrErr returns the Response value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) ResponseOrErr() (*Response, error) {
	if e.loadedTypes[1] {
		if e.Response == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: response.Label}
		}
		return e.Response, nil
	}
	return nil, &NotLoadedError{edge: "response"}
}

// SelectedOptionOrErr returns the SelectedOption value or an error if the edge
// was not loaded in eager-loading.
func (e AnswerEdges) SelectedOptionOrErr() ([]*SelectedOption, error) {
	if e.loadedTypes[2] {
		return e.SelectedOption, nil
	}
	return nil, &NotLoadedError{edge: "selectedOption"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case answer.FieldAnswerText:
			values[i] = new(sql.NullString)
		case answer.FieldID, answer.FieldQuestionID, answer.FieldResponseID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case answer.FieldQuestionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value != nil {
				a.QuestionID = *value
			}
		case answer.FieldResponseID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field response_id", values[i])
			} else if value != nil {
				a.ResponseID = *value
			}
		case answer.FieldAnswerText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer_text", values[i])
			} else if value.Valid {
				a.AnswerText = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Answer.
// This includes values selected through modifiers, order, etc.
func (a *Answer) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryQuestion queries the "question" edge of the Answer entity.
func (a *Answer) QueryQuestion() *QuestionQuery {
	return NewAnswerClient(a.config).QueryQuestion(a)
}

// QueryResponse queries the "response" edge of the Answer entity.
func (a *Answer) QueryResponse() *ResponseQuery {
	return NewAnswerClient(a.config).QueryResponse(a)
}

// QuerySelectedOption queries the "selectedOption" edge of the Answer entity.
func (a *Answer) QuerySelectedOption() *SelectedOptionQuery {
	return NewAnswerClient(a.config).QuerySelectedOption(a)
}

// Update returns a builder for updating this Answer.
// Note that you need to call Answer.Unwrap() before calling this method if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return NewAnswerClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Answer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("question_id=")
	builder.WriteString(fmt.Sprintf("%v", a.QuestionID))
	builder.WriteString(", ")
	builder.WriteString("response_id=")
	builder.WriteString(fmt.Sprintf("%v", a.ResponseID))
	builder.WriteString(", ")
	builder.WriteString("answer_text=")
	builder.WriteString(a.AnswerText)
	builder.WriteByte(')')
	return builder.String()
}

// Answers is a parsable slice of Answer.
type Answers []*Answer
