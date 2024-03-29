// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/questionoption"
	"github.com/google/uuid"
)

// QuestionOption is the model entity for the QuestionOption schema.
type QuestionOption struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// QuestionID holds the value of the "question_id" field.
	QuestionID uuid.UUID `json:"question_id,omitempty"`
	// OptionText holds the value of the "option_text" field.
	OptionText string `json:"option_text,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionOptionQuery when eager-loading is set.
	Edges        QuestionOptionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// QuestionOptionEdges holds the relations/edges for other nodes in the graph.
type QuestionOptionEdges struct {
	// Question holds the value of the question edge.
	Question *Question `json:"question,omitempty"`
	// SelectedOption holds the value of the selectedOption edge.
	SelectedOption []*SelectedOption `json:"selectedOption,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionOptionEdges) QuestionOrErr() (*Question, error) {
	if e.loadedTypes[0] {
		if e.Question == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: question.Label}
		}
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "question"}
}

// SelectedOptionOrErr returns the SelectedOption value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionOptionEdges) SelectedOptionOrErr() ([]*SelectedOption, error) {
	if e.loadedTypes[1] {
		return e.SelectedOption, nil
	}
	return nil, &NotLoadedError{edge: "selectedOption"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QuestionOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case questionoption.FieldOrder:
			values[i] = new(sql.NullInt64)
		case questionoption.FieldOptionText:
			values[i] = new(sql.NullString)
		case questionoption.FieldID, questionoption.FieldQuestionID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QuestionOption fields.
func (qo *QuestionOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case questionoption.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				qo.ID = *value
			}
		case questionoption.FieldQuestionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value != nil {
				qo.QuestionID = *value
			}
		case questionoption.FieldOptionText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field option_text", values[i])
			} else if value.Valid {
				qo.OptionText = value.String
			}
		case questionoption.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				qo.Order = int(value.Int64)
			}
		default:
			qo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the QuestionOption.
// This includes values selected through modifiers, order, etc.
func (qo *QuestionOption) Value(name string) (ent.Value, error) {
	return qo.selectValues.Get(name)
}

// QueryQuestion queries the "question" edge of the QuestionOption entity.
func (qo *QuestionOption) QueryQuestion() *QuestionQuery {
	return NewQuestionOptionClient(qo.config).QueryQuestion(qo)
}

// QuerySelectedOption queries the "selectedOption" edge of the QuestionOption entity.
func (qo *QuestionOption) QuerySelectedOption() *SelectedOptionQuery {
	return NewQuestionOptionClient(qo.config).QuerySelectedOption(qo)
}

// Update returns a builder for updating this QuestionOption.
// Note that you need to call QuestionOption.Unwrap() before calling this method if this QuestionOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (qo *QuestionOption) Update() *QuestionOptionUpdateOne {
	return NewQuestionOptionClient(qo.config).UpdateOne(qo)
}

// Unwrap unwraps the QuestionOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (qo *QuestionOption) Unwrap() *QuestionOption {
	_tx, ok := qo.config.driver.(*txDriver)
	if !ok {
		panic("ent: QuestionOption is not a transactional entity")
	}
	qo.config.driver = _tx.drv
	return qo
}

// String implements the fmt.Stringer.
func (qo *QuestionOption) String() string {
	var builder strings.Builder
	builder.WriteString("QuestionOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", qo.ID))
	builder.WriteString("question_id=")
	builder.WriteString(fmt.Sprintf("%v", qo.QuestionID))
	builder.WriteString(", ")
	builder.WriteString("option_text=")
	builder.WriteString(qo.OptionText)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", qo.Order))
	builder.WriteByte(')')
	return builder.String()
}

// QuestionOptions is a parsable slice of QuestionOption.
type QuestionOptions []*QuestionOption
