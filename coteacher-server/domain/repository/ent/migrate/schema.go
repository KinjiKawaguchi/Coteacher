// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "teacher_id", Type: field.TypeString},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_teachers_classes",
				Columns:    []*schema.Column{ClassesColumns[4]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ClassInvitationCodesColumns holds the columns for the "class_invitation_codes" table.
	ClassInvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "invitation_code", Type: field.TypeString, Unique: true},
		{Name: "expiration_date", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "class_id", Type: field.TypeString},
	}
	// ClassInvitationCodesTable holds the schema information for the "class_invitation_codes" table.
	ClassInvitationCodesTable = &schema.Table{
		Name:       "class_invitation_codes",
		Columns:    ClassInvitationCodesColumns,
		PrimaryKey: []*schema.Column{ClassInvitationCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_invitation_codes_classes_invitationCodes",
				Columns:    []*schema.Column{ClassInvitationCodesColumns[6]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_student", Type: field.TypeString, Unique: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_users_student",
				Columns:    []*schema.Column{StudentsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudentClassesColumns holds the columns for the "student_classes" table.
	StudentClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "class_id", Type: field.TypeString},
		{Name: "student_id", Type: field.TypeString},
	}
	// StudentClassesTable holds the schema information for the "student_classes" table.
	StudentClassesTable = &schema.Table{
		Name:       "student_classes",
		Columns:    StudentClassesColumns,
		PrimaryKey: []*schema.Column{StudentClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_classes_classes_classStudents",
				Columns:    []*schema.Column{StudentClassesColumns[3]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "student_classes_students_studentClasses",
				Columns:    []*schema.Column{StudentClassesColumns[4]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_teacher", Type: field.TypeString, Unique: true},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teachers_users_teacher",
				Columns:    []*schema.Column{TeachersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassesTable,
		ClassInvitationCodesTable,
		StudentsTable,
		StudentClassesTable,
		TeachersTable,
		UsersTable,
	}
)

func init() {
	ClassesTable.ForeignKeys[0].RefTable = TeachersTable
	ClassInvitationCodesTable.ForeignKeys[0].RefTable = ClassesTable
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
	StudentClassesTable.ForeignKeys[0].RefTable = ClassesTable
	StudentClassesTable.ForeignKeys[1].RefTable = StudentsTable
	TeachersTable.ForeignKeys[0].RefTable = UsersTable
}
