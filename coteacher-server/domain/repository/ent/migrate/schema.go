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
		{Name: "user_classes", Type: field.TypeString, Nullable: true},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_users_classes",
				Columns:    []*schema.Column{ClassesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassInvitationCodesColumns holds the columns for the "class_invitation_codes" table.
	ClassInvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "invitation_code", Type: field.TypeString, Unique: true},
		{Name: "expiration_date", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "class_class_invitation_codes", Type: field.TypeString, Nullable: true},
	}
	// ClassInvitationCodesTable holds the schema information for the "class_invitation_codes" table.
	ClassInvitationCodesTable = &schema.Table{
		Name:       "class_invitation_codes",
		Columns:    ClassInvitationCodesColumns,
		PrimaryKey: []*schema.Column{ClassInvitationCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_invitation_codes_classes_class_invitation_codes",
				Columns:    []*schema.Column{ClassInvitationCodesColumns[4]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudentClassesColumns holds the columns for the "student_classes" table.
	StudentClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "class_student_classes", Type: field.TypeString, Nullable: true},
		{Name: "user_student_classes", Type: field.TypeString, Nullable: true},
	}
	// StudentClassesTable holds the schema information for the "student_classes" table.
	StudentClassesTable = &schema.Table{
		Name:       "student_classes",
		Columns:    StudentClassesColumns,
		PrimaryKey: []*schema.Column{StudentClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_classes_classes_student_classes",
				Columns:    []*schema.Column{StudentClassesColumns[1]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "student_classes_users_student_classes",
				Columns:    []*schema.Column{StudentClassesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "user_type", Type: field.TypeEnum, Enums: []string{"teacher", "student"}},
		{Name: "class_users", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_classes_users",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassesTable,
		ClassInvitationCodesTable,
		StudentClassesTable,
		UsersTable,
	}
)

func init() {
	ClassesTable.ForeignKeys[0].RefTable = UsersTable
	ClassInvitationCodesTable.ForeignKeys[0].RefTable = ClassesTable
	StudentClassesTable.ForeignKeys[0].RefTable = ClassesTable
	StudentClassesTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = ClassesTable
}
