// Code generated by ent, DO NOT EDIT.

package student

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeStudentClasses holds the string denoting the studentclasses edge name in mutations.
	EdgeStudentClasses = "studentClasses"
	// EdgeResponses holds the string denoting the responses edge name in mutations.
	EdgeResponses = "responses"
	// Table holds the table name of the student in the database.
	Table = "students"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "students"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_student"
	// StudentClassesTable is the table that holds the studentClasses relation/edge.
	StudentClassesTable = "student_classes"
	// StudentClassesInverseTable is the table name for the StudentClass entity.
	// It exists in this package in order to avoid circular dependency with the "studentclass" package.
	StudentClassesInverseTable = "student_classes"
	// StudentClassesColumn is the table column denoting the studentClasses relation/edge.
	StudentClassesColumn = "student_id"
	// ResponsesTable is the table that holds the responses relation/edge.
	ResponsesTable = "responses"
	// ResponsesInverseTable is the table name for the Response entity.
	// It exists in this package in order to avoid circular dependency with the "response" package.
	ResponsesInverseTable = "responses"
	// ResponsesColumn is the table column denoting the responses relation/edge.
	ResponsesColumn = "student_id"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "students"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_student",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Student queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByStudentClassesCount orders the results by studentClasses count.
func ByStudentClassesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudentClassesStep(), opts...)
	}
}

// ByStudentClasses orders the results by studentClasses terms.
func ByStudentClasses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentClassesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByResponsesCount orders the results by responses count.
func ByResponsesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResponsesStep(), opts...)
	}
}

// ByResponses orders the results by responses terms.
func ByResponses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResponsesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newStudentClassesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentClassesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudentClassesTable, StudentClassesColumn),
	)
}
func newResponsesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResponsesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResponsesTable, ResponsesColumn),
	)
}
