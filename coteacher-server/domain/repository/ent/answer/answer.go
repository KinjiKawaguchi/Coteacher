// Code generated by ent, DO NOT EDIT.

package answer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the answer type in the database.
	Label = "answer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuestionID holds the string denoting the question_id field in the database.
	FieldQuestionID = "question_id"
	// FieldResponseID holds the string denoting the response_id field in the database.
	FieldResponseID = "response_id"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldAnswerText holds the string denoting the answer_text field in the database.
	FieldAnswerText = "answer_text"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeResponse holds the string denoting the response edge name in mutations.
	EdgeResponse = "response"
	// EdgeSelectedOption holds the string denoting the selectedoption edge name in mutations.
	EdgeSelectedOption = "selectedOption"
	// Table holds the table name of the answer in the database.
	Table = "answers"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "answers"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_id"
	// ResponseTable is the table that holds the response relation/edge.
	ResponseTable = "answers"
	// ResponseInverseTable is the table name for the Response entity.
	// It exists in this package in order to avoid circular dependency with the "response" package.
	ResponseInverseTable = "responses"
	// ResponseColumn is the table column denoting the response relation/edge.
	ResponseColumn = "response_id"
	// SelectedOptionTable is the table that holds the selectedOption relation/edge.
	SelectedOptionTable = "selected_options"
	// SelectedOptionInverseTable is the table name for the SelectedOption entity.
	// It exists in this package in order to avoid circular dependency with the "selectedoption" package.
	SelectedOptionInverseTable = "selected_options"
	// SelectedOptionColumn is the table column denoting the selectedOption relation/edge.
	SelectedOptionColumn = "answer_id"
)

// Columns holds all SQL columns for answer fields.
var Columns = []string{
	FieldID,
	FieldQuestionID,
	FieldResponseID,
	FieldOrder,
	FieldAnswerText,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Answer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuestionID orders the results by the question_id field.
func ByQuestionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuestionID, opts...).ToFunc()
}

// ByResponseID orders the results by the response_id field.
func ByResponseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResponseID, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByAnswerText orders the results by the answer_text field.
func ByAnswerText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnswerText, opts...).ToFunc()
}

// ByQuestionField orders the results by question field.
func ByQuestionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), sql.OrderByField(field, opts...))
	}
}

// ByResponseField orders the results by response field.
func ByResponseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResponseStep(), sql.OrderByField(field, opts...))
	}
}

// BySelectedOptionCount orders the results by selectedOption count.
func BySelectedOptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSelectedOptionStep(), opts...)
	}
}

// BySelectedOption orders the results by selectedOption terms.
func BySelectedOption(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSelectedOptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
	)
}
func newResponseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResponseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ResponseTable, ResponseColumn),
	)
}
func newSelectedOptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SelectedOptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SelectedOptionTable, SelectedOptionColumn),
	)
}
