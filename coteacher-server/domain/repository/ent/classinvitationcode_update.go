// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/classinvitationcode"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClassInvitationCodeUpdate is the builder for updating ClassInvitationCode entities.
type ClassInvitationCodeUpdate struct {
	config
	hooks    []Hook
	mutation *ClassInvitationCodeMutation
}

// Where appends a list predicates to the ClassInvitationCodeUpdate builder.
func (cicu *ClassInvitationCodeUpdate) Where(ps ...predicate.ClassInvitationCode) *ClassInvitationCodeUpdate {
	cicu.mutation.Where(ps...)
	return cicu
}

// SetClassID sets the "class_id" field.
func (cicu *ClassInvitationCodeUpdate) SetClassID(u uuid.UUID) *ClassInvitationCodeUpdate {
	cicu.mutation.SetClassID(u)
	return cicu
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableClassID(u *uuid.UUID) *ClassInvitationCodeUpdate {
	if u != nil {
		cicu.SetClassID(*u)
	}
	return cicu
}

// SetInvitationCode sets the "invitation_code" field.
func (cicu *ClassInvitationCodeUpdate) SetInvitationCode(s string) *ClassInvitationCodeUpdate {
	cicu.mutation.SetInvitationCode(s)
	return cicu
}

// SetNillableInvitationCode sets the "invitation_code" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableInvitationCode(s *string) *ClassInvitationCodeUpdate {
	if s != nil {
		cicu.SetInvitationCode(*s)
	}
	return cicu
}

// SetExpirationDate sets the "expiration_date" field.
func (cicu *ClassInvitationCodeUpdate) SetExpirationDate(t time.Time) *ClassInvitationCodeUpdate {
	cicu.mutation.SetExpirationDate(t)
	return cicu
}

// SetNillableExpirationDate sets the "expiration_date" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableExpirationDate(t *time.Time) *ClassInvitationCodeUpdate {
	if t != nil {
		cicu.SetExpirationDate(*t)
	}
	return cicu
}

// SetIsActive sets the "is_active" field.
func (cicu *ClassInvitationCodeUpdate) SetIsActive(b bool) *ClassInvitationCodeUpdate {
	cicu.mutation.SetIsActive(b)
	return cicu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableIsActive(b *bool) *ClassInvitationCodeUpdate {
	if b != nil {
		cicu.SetIsActive(*b)
	}
	return cicu
}

// SetCreatedAt sets the "created_at" field.
func (cicu *ClassInvitationCodeUpdate) SetCreatedAt(t time.Time) *ClassInvitationCodeUpdate {
	cicu.mutation.SetCreatedAt(t)
	return cicu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableCreatedAt(t *time.Time) *ClassInvitationCodeUpdate {
	if t != nil {
		cicu.SetCreatedAt(*t)
	}
	return cicu
}

// SetUpdatedAt sets the "updated_at" field.
func (cicu *ClassInvitationCodeUpdate) SetUpdatedAt(t time.Time) *ClassInvitationCodeUpdate {
	cicu.mutation.SetUpdatedAt(t)
	return cicu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cicu *ClassInvitationCodeUpdate) SetNillableUpdatedAt(t *time.Time) *ClassInvitationCodeUpdate {
	if t != nil {
		cicu.SetUpdatedAt(*t)
	}
	return cicu
}

// SetClass sets the "class" edge to the Class entity.
func (cicu *ClassInvitationCodeUpdate) SetClass(c *Class) *ClassInvitationCodeUpdate {
	return cicu.SetClassID(c.ID)
}

// Mutation returns the ClassInvitationCodeMutation object of the builder.
func (cicu *ClassInvitationCodeUpdate) Mutation() *ClassInvitationCodeMutation {
	return cicu.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (cicu *ClassInvitationCodeUpdate) ClearClass() *ClassInvitationCodeUpdate {
	cicu.mutation.ClearClass()
	return cicu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cicu *ClassInvitationCodeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cicu.sqlSave, cicu.mutation, cicu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cicu *ClassInvitationCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := cicu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cicu *ClassInvitationCodeUpdate) Exec(ctx context.Context) error {
	_, err := cicu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicu *ClassInvitationCodeUpdate) ExecX(ctx context.Context) {
	if err := cicu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cicu *ClassInvitationCodeUpdate) check() error {
	if _, ok := cicu.mutation.ClassID(); cicu.mutation.ClassCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ClassInvitationCode.class"`)
	}
	return nil
}

func (cicu *ClassInvitationCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cicu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(classinvitationcode.Table, classinvitationcode.Columns, sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID))
	if ps := cicu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cicu.mutation.InvitationCode(); ok {
		_spec.SetField(classinvitationcode.FieldInvitationCode, field.TypeString, value)
	}
	if value, ok := cicu.mutation.ExpirationDate(); ok {
		_spec.SetField(classinvitationcode.FieldExpirationDate, field.TypeTime, value)
	}
	if value, ok := cicu.mutation.IsActive(); ok {
		_spec.SetField(classinvitationcode.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cicu.mutation.CreatedAt(); ok {
		_spec.SetField(classinvitationcode.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cicu.mutation.UpdatedAt(); ok {
		_spec.SetField(classinvitationcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if cicu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classinvitationcode.ClassTable,
			Columns: []string{classinvitationcode.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cicu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classinvitationcode.ClassTable,
			Columns: []string{classinvitationcode.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cicu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classinvitationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cicu.mutation.done = true
	return n, nil
}

// ClassInvitationCodeUpdateOne is the builder for updating a single ClassInvitationCode entity.
type ClassInvitationCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassInvitationCodeMutation
}

// SetClassID sets the "class_id" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetClassID(u uuid.UUID) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetClassID(u)
	return cicuo
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableClassID(u *uuid.UUID) *ClassInvitationCodeUpdateOne {
	if u != nil {
		cicuo.SetClassID(*u)
	}
	return cicuo
}

// SetInvitationCode sets the "invitation_code" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetInvitationCode(s string) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetInvitationCode(s)
	return cicuo
}

// SetNillableInvitationCode sets the "invitation_code" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableInvitationCode(s *string) *ClassInvitationCodeUpdateOne {
	if s != nil {
		cicuo.SetInvitationCode(*s)
	}
	return cicuo
}

// SetExpirationDate sets the "expiration_date" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetExpirationDate(t time.Time) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetExpirationDate(t)
	return cicuo
}

// SetNillableExpirationDate sets the "expiration_date" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableExpirationDate(t *time.Time) *ClassInvitationCodeUpdateOne {
	if t != nil {
		cicuo.SetExpirationDate(*t)
	}
	return cicuo
}

// SetIsActive sets the "is_active" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetIsActive(b bool) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetIsActive(b)
	return cicuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableIsActive(b *bool) *ClassInvitationCodeUpdateOne {
	if b != nil {
		cicuo.SetIsActive(*b)
	}
	return cicuo
}

// SetCreatedAt sets the "created_at" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetCreatedAt(t time.Time) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetCreatedAt(t)
	return cicuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableCreatedAt(t *time.Time) *ClassInvitationCodeUpdateOne {
	if t != nil {
		cicuo.SetCreatedAt(*t)
	}
	return cicuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cicuo *ClassInvitationCodeUpdateOne) SetUpdatedAt(t time.Time) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.SetUpdatedAt(t)
	return cicuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cicuo *ClassInvitationCodeUpdateOne) SetNillableUpdatedAt(t *time.Time) *ClassInvitationCodeUpdateOne {
	if t != nil {
		cicuo.SetUpdatedAt(*t)
	}
	return cicuo
}

// SetClass sets the "class" edge to the Class entity.
func (cicuo *ClassInvitationCodeUpdateOne) SetClass(c *Class) *ClassInvitationCodeUpdateOne {
	return cicuo.SetClassID(c.ID)
}

// Mutation returns the ClassInvitationCodeMutation object of the builder.
func (cicuo *ClassInvitationCodeUpdateOne) Mutation() *ClassInvitationCodeMutation {
	return cicuo.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (cicuo *ClassInvitationCodeUpdateOne) ClearClass() *ClassInvitationCodeUpdateOne {
	cicuo.mutation.ClearClass()
	return cicuo
}

// Where appends a list predicates to the ClassInvitationCodeUpdate builder.
func (cicuo *ClassInvitationCodeUpdateOne) Where(ps ...predicate.ClassInvitationCode) *ClassInvitationCodeUpdateOne {
	cicuo.mutation.Where(ps...)
	return cicuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cicuo *ClassInvitationCodeUpdateOne) Select(field string, fields ...string) *ClassInvitationCodeUpdateOne {
	cicuo.fields = append([]string{field}, fields...)
	return cicuo
}

// Save executes the query and returns the updated ClassInvitationCode entity.
func (cicuo *ClassInvitationCodeUpdateOne) Save(ctx context.Context) (*ClassInvitationCode, error) {
	return withHooks(ctx, cicuo.sqlSave, cicuo.mutation, cicuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cicuo *ClassInvitationCodeUpdateOne) SaveX(ctx context.Context) *ClassInvitationCode {
	node, err := cicuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cicuo *ClassInvitationCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := cicuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicuo *ClassInvitationCodeUpdateOne) ExecX(ctx context.Context) {
	if err := cicuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cicuo *ClassInvitationCodeUpdateOne) check() error {
	if _, ok := cicuo.mutation.ClassID(); cicuo.mutation.ClassCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ClassInvitationCode.class"`)
	}
	return nil
}

func (cicuo *ClassInvitationCodeUpdateOne) sqlSave(ctx context.Context) (_node *ClassInvitationCode, err error) {
	if err := cicuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(classinvitationcode.Table, classinvitationcode.Columns, sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID))
	id, ok := cicuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ClassInvitationCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cicuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classinvitationcode.FieldID)
		for _, f := range fields {
			if !classinvitationcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != classinvitationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cicuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cicuo.mutation.InvitationCode(); ok {
		_spec.SetField(classinvitationcode.FieldInvitationCode, field.TypeString, value)
	}
	if value, ok := cicuo.mutation.ExpirationDate(); ok {
		_spec.SetField(classinvitationcode.FieldExpirationDate, field.TypeTime, value)
	}
	if value, ok := cicuo.mutation.IsActive(); ok {
		_spec.SetField(classinvitationcode.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cicuo.mutation.CreatedAt(); ok {
		_spec.SetField(classinvitationcode.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cicuo.mutation.UpdatedAt(); ok {
		_spec.SetField(classinvitationcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if cicuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classinvitationcode.ClassTable,
			Columns: []string{classinvitationcode.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cicuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classinvitationcode.ClassTable,
			Columns: []string{classinvitationcode.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClassInvitationCode{config: cicuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cicuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classinvitationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cicuo.mutation.done = true
	return _node, nil
}
