// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/answer"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/form"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/student"
	"github.com/google/uuid"
)

// ResponseUpdate is the builder for updating Response entities.
type ResponseUpdate struct {
	config
	hooks    []Hook
	mutation *ResponseMutation
}

// Where appends a list predicates to the ResponseUpdate builder.
func (ru *ResponseUpdate) Where(ps ...predicate.Response) *ResponseUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetStudentID sets the "student_id" field.
func (ru *ResponseUpdate) SetStudentID(u uuid.UUID) *ResponseUpdate {
	ru.mutation.SetStudentID(u)
	return ru
}

// SetNillableStudentID sets the "student_id" field if the given value is not nil.
func (ru *ResponseUpdate) SetNillableStudentID(u *uuid.UUID) *ResponseUpdate {
	if u != nil {
		ru.SetStudentID(*u)
	}
	return ru
}

// SetFormID sets the "form_id" field.
func (ru *ResponseUpdate) SetFormID(u uuid.UUID) *ResponseUpdate {
	ru.mutation.SetFormID(u)
	return ru
}

// SetNillableFormID sets the "form_id" field if the given value is not nil.
func (ru *ResponseUpdate) SetNillableFormID(u *uuid.UUID) *ResponseUpdate {
	if u != nil {
		ru.SetFormID(*u)
	}
	return ru
}

// SetAiResponse sets the "ai_response" field.
func (ru *ResponseUpdate) SetAiResponse(s string) *ResponseUpdate {
	ru.mutation.SetAiResponse(s)
	return ru
}

// SetNillableAiResponse sets the "ai_response" field if the given value is not nil.
func (ru *ResponseUpdate) SetNillableAiResponse(s *string) *ResponseUpdate {
	if s != nil {
		ru.SetAiResponse(*s)
	}
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ResponseUpdate) SetCreatedAt(t time.Time) *ResponseUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ResponseUpdate) SetNillableCreatedAt(t *time.Time) *ResponseUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ResponseUpdate) SetUpdatedAt(t time.Time) *ResponseUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetStudent sets the "student" edge to the Student entity.
func (ru *ResponseUpdate) SetStudent(s *Student) *ResponseUpdate {
	return ru.SetStudentID(s.ID)
}

// SetForm sets the "form" edge to the Form entity.
func (ru *ResponseUpdate) SetForm(f *Form) *ResponseUpdate {
	return ru.SetFormID(f.ID)
}

// AddAnswerIDs adds the "answer" edge to the Answer entity by IDs.
func (ru *ResponseUpdate) AddAnswerIDs(ids ...uuid.UUID) *ResponseUpdate {
	ru.mutation.AddAnswerIDs(ids...)
	return ru
}

// AddAnswer adds the "answer" edges to the Answer entity.
func (ru *ResponseUpdate) AddAnswer(a ...*Answer) *ResponseUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAnswerIDs(ids...)
}

// Mutation returns the ResponseMutation object of the builder.
func (ru *ResponseUpdate) Mutation() *ResponseMutation {
	return ru.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (ru *ResponseUpdate) ClearStudent() *ResponseUpdate {
	ru.mutation.ClearStudent()
	return ru
}

// ClearForm clears the "form" edge to the Form entity.
func (ru *ResponseUpdate) ClearForm() *ResponseUpdate {
	ru.mutation.ClearForm()
	return ru
}

// ClearAnswer clears all "answer" edges to the Answer entity.
func (ru *ResponseUpdate) ClearAnswer() *ResponseUpdate {
	ru.mutation.ClearAnswer()
	return ru
}

// RemoveAnswerIDs removes the "answer" edge to Answer entities by IDs.
func (ru *ResponseUpdate) RemoveAnswerIDs(ids ...uuid.UUID) *ResponseUpdate {
	ru.mutation.RemoveAnswerIDs(ids...)
	return ru
}

// RemoveAnswer removes "answer" edges to Answer entities.
func (ru *ResponseUpdate) RemoveAnswer(a ...*Answer) *ResponseUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResponseUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResponseUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResponseUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResponseUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ResponseUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := response.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResponseUpdate) check() error {
	if _, ok := ru.mutation.StudentID(); ru.mutation.StudentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Response.student"`)
	}
	if _, ok := ru.mutation.FormID(); ru.mutation.FormCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Response.form"`)
	}
	return nil
}

func (ru *ResponseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(response.Table, response.Columns, sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.AiResponse(); ok {
		_spec.SetField(response.FieldAiResponse, field.TypeString, value)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(response.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(response.FieldUpdatedAt, field.TypeTime, value)
	}
	if ru.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.StudentTable,
			Columns: []string{response.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.StudentTable,
			Columns: []string{response.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.FormCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.FormTable,
			Columns: []string{response.FormColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(form.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.FormIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.FormTable,
			Columns: []string{response.FormColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(form.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !ru.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{response.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResponseUpdateOne is the builder for updating a single Response entity.
type ResponseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResponseMutation
}

// SetStudentID sets the "student_id" field.
func (ruo *ResponseUpdateOne) SetStudentID(u uuid.UUID) *ResponseUpdateOne {
	ruo.mutation.SetStudentID(u)
	return ruo
}

// SetNillableStudentID sets the "student_id" field if the given value is not nil.
func (ruo *ResponseUpdateOne) SetNillableStudentID(u *uuid.UUID) *ResponseUpdateOne {
	if u != nil {
		ruo.SetStudentID(*u)
	}
	return ruo
}

// SetFormID sets the "form_id" field.
func (ruo *ResponseUpdateOne) SetFormID(u uuid.UUID) *ResponseUpdateOne {
	ruo.mutation.SetFormID(u)
	return ruo
}

// SetNillableFormID sets the "form_id" field if the given value is not nil.
func (ruo *ResponseUpdateOne) SetNillableFormID(u *uuid.UUID) *ResponseUpdateOne {
	if u != nil {
		ruo.SetFormID(*u)
	}
	return ruo
}

// SetAiResponse sets the "ai_response" field.
func (ruo *ResponseUpdateOne) SetAiResponse(s string) *ResponseUpdateOne {
	ruo.mutation.SetAiResponse(s)
	return ruo
}

// SetNillableAiResponse sets the "ai_response" field if the given value is not nil.
func (ruo *ResponseUpdateOne) SetNillableAiResponse(s *string) *ResponseUpdateOne {
	if s != nil {
		ruo.SetAiResponse(*s)
	}
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ResponseUpdateOne) SetCreatedAt(t time.Time) *ResponseUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ResponseUpdateOne) SetNillableCreatedAt(t *time.Time) *ResponseUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ResponseUpdateOne) SetUpdatedAt(t time.Time) *ResponseUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetStudent sets the "student" edge to the Student entity.
func (ruo *ResponseUpdateOne) SetStudent(s *Student) *ResponseUpdateOne {
	return ruo.SetStudentID(s.ID)
}

// SetForm sets the "form" edge to the Form entity.
func (ruo *ResponseUpdateOne) SetForm(f *Form) *ResponseUpdateOne {
	return ruo.SetFormID(f.ID)
}

// AddAnswerIDs adds the "answer" edge to the Answer entity by IDs.
func (ruo *ResponseUpdateOne) AddAnswerIDs(ids ...uuid.UUID) *ResponseUpdateOne {
	ruo.mutation.AddAnswerIDs(ids...)
	return ruo
}

// AddAnswer adds the "answer" edges to the Answer entity.
func (ruo *ResponseUpdateOne) AddAnswer(a ...*Answer) *ResponseUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAnswerIDs(ids...)
}

// Mutation returns the ResponseMutation object of the builder.
func (ruo *ResponseUpdateOne) Mutation() *ResponseMutation {
	return ruo.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (ruo *ResponseUpdateOne) ClearStudent() *ResponseUpdateOne {
	ruo.mutation.ClearStudent()
	return ruo
}

// ClearForm clears the "form" edge to the Form entity.
func (ruo *ResponseUpdateOne) ClearForm() *ResponseUpdateOne {
	ruo.mutation.ClearForm()
	return ruo
}

// ClearAnswer clears all "answer" edges to the Answer entity.
func (ruo *ResponseUpdateOne) ClearAnswer() *ResponseUpdateOne {
	ruo.mutation.ClearAnswer()
	return ruo
}

// RemoveAnswerIDs removes the "answer" edge to Answer entities by IDs.
func (ruo *ResponseUpdateOne) RemoveAnswerIDs(ids ...uuid.UUID) *ResponseUpdateOne {
	ruo.mutation.RemoveAnswerIDs(ids...)
	return ruo
}

// RemoveAnswer removes "answer" edges to Answer entities.
func (ruo *ResponseUpdateOne) RemoveAnswer(a ...*Answer) *ResponseUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAnswerIDs(ids...)
}

// Where appends a list predicates to the ResponseUpdate builder.
func (ruo *ResponseUpdateOne) Where(ps ...predicate.Response) *ResponseUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResponseUpdateOne) Select(field string, fields ...string) *ResponseUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Response entity.
func (ruo *ResponseUpdateOne) Save(ctx context.Context) (*Response, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResponseUpdateOne) SaveX(ctx context.Context) *Response {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResponseUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResponseUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ResponseUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := response.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResponseUpdateOne) check() error {
	if _, ok := ruo.mutation.StudentID(); ruo.mutation.StudentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Response.student"`)
	}
	if _, ok := ruo.mutation.FormID(); ruo.mutation.FormCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Response.form"`)
	}
	return nil
}

func (ruo *ResponseUpdateOne) sqlSave(ctx context.Context) (_node *Response, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(response.Table, response.Columns, sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Response.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, response.FieldID)
		for _, f := range fields {
			if !response.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != response.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.AiResponse(); ok {
		_spec.SetField(response.FieldAiResponse, field.TypeString, value)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(response.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(response.FieldUpdatedAt, field.TypeTime, value)
	}
	if ruo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.StudentTable,
			Columns: []string{response.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.StudentTable,
			Columns: []string{response.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.FormCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.FormTable,
			Columns: []string{response.FormColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(form.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.FormIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.FormTable,
			Columns: []string{response.FormColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(form.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !ruo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.AnswerTable,
			Columns: []string{response.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Response{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{response.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
