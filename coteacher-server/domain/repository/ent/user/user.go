// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUserType holds the string denoting the usertype field in the database.
	FieldUserType = "user_type"
	// EdgeStudentClasses holds the string denoting the student_classes edge name in mutations.
	EdgeStudentClasses = "student_classes"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// StudentClassesTable is the table that holds the student_classes relation/edge.
	StudentClassesTable = "student_classes"
	// StudentClassesInverseTable is the table name for the StudentClass entity.
	// It exists in this package in order to avoid circular dependency with the "studentclass" package.
	StudentClassesInverseTable = "student_classes"
	// StudentClassesColumn is the table column denoting the student_classes relation/edge.
	StudentClassesColumn = "user_student_classes"
	// ClassesTable is the table that holds the classes relation/edge.
	ClassesTable = "classes"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "classes"
	// ClassesColumn is the table column denoting the classes relation/edge.
	ClassesColumn = "user_classes"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldUserType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_users",
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

// UserType defines the type for the "UserType" enum field.
type UserType string

// UserType values.
const (
	UserTypeTeacher UserType = "teacher"
	UserTypeStudent UserType = "student"
)

func (_usertype UserType) String() string {
	return string(_usertype)
}

// UserTypeValidator is a validator for the "UserType" field enum values. It is called by the builders before save.
func UserTypeValidator(_usertype UserType) error {
	switch _usertype {
	case UserTypeTeacher, UserTypeStudent:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for UserType field: %q", _usertype)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUserType orders the results by the UserType field.
func ByUserType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserType, opts...).ToFunc()
}

// ByStudentClassesCount orders the results by student_classes count.
func ByStudentClassesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudentClassesStep(), opts...)
	}
}

// ByStudentClasses orders the results by student_classes terms.
func ByStudentClasses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentClassesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByClassesCount orders the results by classes count.
func ByClassesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassesStep(), opts...)
	}
}

// ByClasses orders the results by classes terms.
func ByClasses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStudentClassesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentClassesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudentClassesTable, StudentClassesColumn),
	)
}
func newClassesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
	)
}
