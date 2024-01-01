// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/classinvitationcode"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/studentclass"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/teacher"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClassUpdate is the builder for updating Class entities.
type ClassUpdate struct {
	config
	hooks    []Hook
	mutation *ClassMutation
}

// Where appends a list predicates to the ClassUpdate builder.
func (cu *ClassUpdate) Where(ps ...predicate.Class) *ClassUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ClassUpdate) SetName(s string) *ClassUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableName(s *string) *ClassUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetTeacherID sets the "teacher_id" field.
func (cu *ClassUpdate) SetTeacherID(u uuid.UUID) *ClassUpdate {
	cu.mutation.SetTeacherID(u)
	return cu
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableTeacherID(u *uuid.UUID) *ClassUpdate {
	if u != nil {
		cu.SetTeacherID(*u)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ClassUpdate) SetCreatedAt(t time.Time) *ClassUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableCreatedAt(t *time.Time) *ClassUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ClassUpdate) SetUpdatedAt(t time.Time) *ClassUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableUpdatedAt(t *time.Time) *ClassUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (cu *ClassUpdate) SetTeacher(t *Teacher) *ClassUpdate {
	return cu.SetTeacherID(t.ID)
}

// AddClassStudentIDs adds the "classStudents" edge to the StudentClass entity by IDs.
func (cu *ClassUpdate) AddClassStudentIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddClassStudentIDs(ids...)
	return cu
}

// AddClassStudents adds the "classStudents" edges to the StudentClass entity.
func (cu *ClassUpdate) AddClassStudents(s ...*StudentClass) *ClassUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddClassStudentIDs(ids...)
}

// AddInvitationCodeIDs adds the "invitationCodes" edge to the ClassInvitationCode entity by IDs.
func (cu *ClassUpdate) AddInvitationCodeIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.AddInvitationCodeIDs(ids...)
	return cu
}

// AddInvitationCodes adds the "invitationCodes" edges to the ClassInvitationCode entity.
func (cu *ClassUpdate) AddInvitationCodes(c ...*ClassInvitationCode) *ClassUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddInvitationCodeIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cu *ClassUpdate) Mutation() *ClassMutation {
	return cu.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (cu *ClassUpdate) ClearTeacher() *ClassUpdate {
	cu.mutation.ClearTeacher()
	return cu
}

// ClearClassStudents clears all "classStudents" edges to the StudentClass entity.
func (cu *ClassUpdate) ClearClassStudents() *ClassUpdate {
	cu.mutation.ClearClassStudents()
	return cu
}

// RemoveClassStudentIDs removes the "classStudents" edge to StudentClass entities by IDs.
func (cu *ClassUpdate) RemoveClassStudentIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveClassStudentIDs(ids...)
	return cu
}

// RemoveClassStudents removes "classStudents" edges to StudentClass entities.
func (cu *ClassUpdate) RemoveClassStudents(s ...*StudentClass) *ClassUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveClassStudentIDs(ids...)
}

// ClearInvitationCodes clears all "invitationCodes" edges to the ClassInvitationCode entity.
func (cu *ClassUpdate) ClearInvitationCodes() *ClassUpdate {
	cu.mutation.ClearInvitationCodes()
	return cu
}

// RemoveInvitationCodeIDs removes the "invitationCodes" edge to ClassInvitationCode entities by IDs.
func (cu *ClassUpdate) RemoveInvitationCodeIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.RemoveInvitationCodeIDs(ids...)
	return cu
}

// RemoveInvitationCodes removes "invitationCodes" edges to ClassInvitationCode entities.
func (cu *ClassUpdate) RemoveInvitationCodes(c ...*ClassInvitationCode) *ClassUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveInvitationCodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClassUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClassUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClassUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClassUpdate) check() error {
	if _, ok := cu.mutation.TeacherID(); cu.mutation.TeacherCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Class.teacher"`)
	}
	return nil
}

func (cu *ClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(class.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(class.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClassStudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClassStudentsIDs(); len(nodes) > 0 && !cu.mutation.ClassStudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClassStudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.InvitationCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedInvitationCodesIDs(); len(nodes) > 0 && !cu.mutation.InvitationCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InvitationCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ClassUpdateOne is the builder for updating a single Class entity.
type ClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassMutation
}

// SetName sets the "name" field.
func (cuo *ClassUpdateOne) SetName(s string) *ClassUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableName(s *string) *ClassUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetTeacherID sets the "teacher_id" field.
func (cuo *ClassUpdateOne) SetTeacherID(u uuid.UUID) *ClassUpdateOne {
	cuo.mutation.SetTeacherID(u)
	return cuo
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableTeacherID(u *uuid.UUID) *ClassUpdateOne {
	if u != nil {
		cuo.SetTeacherID(*u)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ClassUpdateOne) SetCreatedAt(t time.Time) *ClassUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableCreatedAt(t *time.Time) *ClassUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ClassUpdateOne) SetUpdatedAt(t time.Time) *ClassUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableUpdatedAt(t *time.Time) *ClassUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (cuo *ClassUpdateOne) SetTeacher(t *Teacher) *ClassUpdateOne {
	return cuo.SetTeacherID(t.ID)
}

// AddClassStudentIDs adds the "classStudents" edge to the StudentClass entity by IDs.
func (cuo *ClassUpdateOne) AddClassStudentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddClassStudentIDs(ids...)
	return cuo
}

// AddClassStudents adds the "classStudents" edges to the StudentClass entity.
func (cuo *ClassUpdateOne) AddClassStudents(s ...*StudentClass) *ClassUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddClassStudentIDs(ids...)
}

// AddInvitationCodeIDs adds the "invitationCodes" edge to the ClassInvitationCode entity by IDs.
func (cuo *ClassUpdateOne) AddInvitationCodeIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.AddInvitationCodeIDs(ids...)
	return cuo
}

// AddInvitationCodes adds the "invitationCodes" edges to the ClassInvitationCode entity.
func (cuo *ClassUpdateOne) AddInvitationCodes(c ...*ClassInvitationCode) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddInvitationCodeIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cuo *ClassUpdateOne) Mutation() *ClassMutation {
	return cuo.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (cuo *ClassUpdateOne) ClearTeacher() *ClassUpdateOne {
	cuo.mutation.ClearTeacher()
	return cuo
}

// ClearClassStudents clears all "classStudents" edges to the StudentClass entity.
func (cuo *ClassUpdateOne) ClearClassStudents() *ClassUpdateOne {
	cuo.mutation.ClearClassStudents()
	return cuo
}

// RemoveClassStudentIDs removes the "classStudents" edge to StudentClass entities by IDs.
func (cuo *ClassUpdateOne) RemoveClassStudentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveClassStudentIDs(ids...)
	return cuo
}

// RemoveClassStudents removes "classStudents" edges to StudentClass entities.
func (cuo *ClassUpdateOne) RemoveClassStudents(s ...*StudentClass) *ClassUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveClassStudentIDs(ids...)
}

// ClearInvitationCodes clears all "invitationCodes" edges to the ClassInvitationCode entity.
func (cuo *ClassUpdateOne) ClearInvitationCodes() *ClassUpdateOne {
	cuo.mutation.ClearInvitationCodes()
	return cuo
}

// RemoveInvitationCodeIDs removes the "invitationCodes" edge to ClassInvitationCode entities by IDs.
func (cuo *ClassUpdateOne) RemoveInvitationCodeIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.RemoveInvitationCodeIDs(ids...)
	return cuo
}

// RemoveInvitationCodes removes "invitationCodes" edges to ClassInvitationCode entities.
func (cuo *ClassUpdateOne) RemoveInvitationCodes(c ...*ClassInvitationCode) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveInvitationCodeIDs(ids...)
}

// Where appends a list predicates to the ClassUpdate builder.
func (cuo *ClassUpdateOne) Where(ps ...predicate.Class) *ClassUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClassUpdateOne) Select(field string, fields ...string) *ClassUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Class entity.
func (cuo *ClassUpdateOne) Save(ctx context.Context) (*Class, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClassUpdateOne) SaveX(ctx context.Context) *Class {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClassUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClassUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClassUpdateOne) check() error {
	if _, ok := cuo.mutation.TeacherID(); cuo.mutation.TeacherCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Class.teacher"`)
	}
	return nil
}

func (cuo *ClassUpdateOne) sqlSave(ctx context.Context) (_node *Class, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Class.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for _, f := range fields {
			if !class.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(class.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(class.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClassStudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClassStudentsIDs(); len(nodes) > 0 && !cuo.mutation.ClassStudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClassStudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ClassStudentsTable,
			Columns: []string{class.ClassStudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studentclass.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.InvitationCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedInvitationCodesIDs(); len(nodes) > 0 && !cuo.mutation.InvitationCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InvitationCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.InvitationCodesTable,
			Columns: []string{class.InvitationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Class{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
