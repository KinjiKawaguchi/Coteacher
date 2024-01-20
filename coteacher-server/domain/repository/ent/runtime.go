// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/answer"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/classinvitationcode"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/form"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/questionoption"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/schema"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/textquestion"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescID is the schema descriptor for id field.
	answerDescID := answerFields[0].Descriptor()
	// answer.DefaultID holds the default value on creation for the id field.
	answer.DefaultID = answerDescID.Default.(func() uuid.UUID)
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescCreatedAt is the schema descriptor for created_at field.
	classDescCreatedAt := classFields[3].Descriptor()
	// class.DefaultCreatedAt holds the default value on creation for the created_at field.
	class.DefaultCreatedAt = classDescCreatedAt.Default.(func() time.Time)
	// classDescUpdatedAt is the schema descriptor for updated_at field.
	classDescUpdatedAt := classFields[4].Descriptor()
	// class.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	class.DefaultUpdatedAt = classDescUpdatedAt.Default.(func() time.Time)
	// class.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	class.UpdateDefaultUpdatedAt = classDescUpdatedAt.UpdateDefault.(func() time.Time)
	// classDescID is the schema descriptor for id field.
	classDescID := classFields[0].Descriptor()
	// class.DefaultID holds the default value on creation for the id field.
	class.DefaultID = classDescID.Default.(func() uuid.UUID)
	classinvitationcodeFields := schema.ClassInvitationCode{}.Fields()
	_ = classinvitationcodeFields
	// classinvitationcodeDescIsActive is the schema descriptor for is_active field.
	classinvitationcodeDescIsActive := classinvitationcodeFields[4].Descriptor()
	// classinvitationcode.DefaultIsActive holds the default value on creation for the is_active field.
	classinvitationcode.DefaultIsActive = classinvitationcodeDescIsActive.Default.(bool)
	// classinvitationcodeDescCreatedAt is the schema descriptor for created_at field.
	classinvitationcodeDescCreatedAt := classinvitationcodeFields[5].Descriptor()
	// classinvitationcode.DefaultCreatedAt holds the default value on creation for the created_at field.
	classinvitationcode.DefaultCreatedAt = classinvitationcodeDescCreatedAt.Default.(func() time.Time)
	// classinvitationcodeDescUpdatedAt is the schema descriptor for updated_at field.
	classinvitationcodeDescUpdatedAt := classinvitationcodeFields[6].Descriptor()
	// classinvitationcode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	classinvitationcode.DefaultUpdatedAt = classinvitationcodeDescUpdatedAt.Default.(func() time.Time)
	// classinvitationcode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	classinvitationcode.UpdateDefaultUpdatedAt = classinvitationcodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// classinvitationcodeDescID is the schema descriptor for id field.
	classinvitationcodeDescID := classinvitationcodeFields[0].Descriptor()
	// classinvitationcode.DefaultID holds the default value on creation for the id field.
	classinvitationcode.DefaultID = classinvitationcodeDescID.Default.(func() uuid.UUID)
	formFields := schema.Form{}.Fields()
	_ = formFields
	// formDescCreatedAt is the schema descriptor for created_at field.
	formDescCreatedAt := formFields[5].Descriptor()
	// form.DefaultCreatedAt holds the default value on creation for the created_at field.
	form.DefaultCreatedAt = formDescCreatedAt.Default.(func() time.Time)
	// formDescUpdatedAt is the schema descriptor for updated_at field.
	formDescUpdatedAt := formFields[6].Descriptor()
	// form.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	form.DefaultUpdatedAt = formDescUpdatedAt.Default.(func() time.Time)
	// form.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	form.UpdateDefaultUpdatedAt = formDescUpdatedAt.UpdateDefault.(func() time.Time)
	// formDescID is the schema descriptor for id field.
	formDescID := formFields[0].Descriptor()
	// form.DefaultID holds the default value on creation for the id field.
	form.DefaultID = formDescID.Default.(func() uuid.UUID)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescCreatedAt is the schema descriptor for created_at field.
	questionDescCreatedAt := questionFields[7].Descriptor()
	// question.DefaultCreatedAt holds the default value on creation for the created_at field.
	question.DefaultCreatedAt = questionDescCreatedAt.Default.(func() time.Time)
	// questionDescUpdatedAt is the schema descriptor for updated_at field.
	questionDescUpdatedAt := questionFields[8].Descriptor()
	// question.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	question.DefaultUpdatedAt = questionDescUpdatedAt.Default.(func() time.Time)
	// question.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	question.UpdateDefaultUpdatedAt = questionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// questionDescID is the schema descriptor for id field.
	questionDescID := questionFields[0].Descriptor()
	// question.DefaultID holds the default value on creation for the id field.
	question.DefaultID = questionDescID.Default.(func() uuid.UUID)
	questionoptionFields := schema.QuestionOption{}.Fields()
	_ = questionoptionFields
	// questionoptionDescID is the schema descriptor for id field.
	questionoptionDescID := questionoptionFields[0].Descriptor()
	// questionoption.DefaultID holds the default value on creation for the id field.
	questionoption.DefaultID = questionoptionDescID.Default.(func() uuid.UUID)
	responseFields := schema.Response{}.Fields()
	_ = responseFields
	// responseDescCreatedAt is the schema descriptor for created_at field.
	responseDescCreatedAt := responseFields[4].Descriptor()
	// response.DefaultCreatedAt holds the default value on creation for the created_at field.
	response.DefaultCreatedAt = responseDescCreatedAt.Default.(func() time.Time)
	// responseDescUpdatedAt is the schema descriptor for updated_at field.
	responseDescUpdatedAt := responseFields[5].Descriptor()
	// response.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	response.DefaultUpdatedAt = responseDescUpdatedAt.Default.(func() time.Time)
	// response.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	response.UpdateDefaultUpdatedAt = responseDescUpdatedAt.UpdateDefault.(func() time.Time)
	// responseDescID is the schema descriptor for id field.
	responseDescID := responseFields[0].Descriptor()
	// response.DefaultID holds the default value on creation for the id field.
	response.DefaultID = responseDescID.Default.(func() uuid.UUID)
	studentclassFields := schema.StudentClass{}.Fields()
	_ = studentclassFields
	// studentclassDescCreatedAt is the schema descriptor for created_at field.
	studentclassDescCreatedAt := studentclassFields[2].Descriptor()
	// studentclass.DefaultCreatedAt holds the default value on creation for the created_at field.
	studentclass.DefaultCreatedAt = studentclassDescCreatedAt.Default.(func() time.Time)
	// studentclassDescUpdatedAt is the schema descriptor for updated_at field.
	studentclassDescUpdatedAt := studentclassFields[3].Descriptor()
	// studentclass.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	studentclass.DefaultUpdatedAt = studentclassDescUpdatedAt.Default.(func() time.Time)
	// studentclass.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	studentclass.UpdateDefaultUpdatedAt = studentclassDescUpdatedAt.UpdateDefault.(func() time.Time)
	textquestionFields := schema.TextQuestion{}.Fields()
	_ = textquestionFields
	// textquestionDescID is the schema descriptor for id field.
	textquestionDescID := textquestionFields[0].Descriptor()
	// textquestion.DefaultID holds the default value on creation for the id field.
	textquestion.DefaultID = textquestionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
