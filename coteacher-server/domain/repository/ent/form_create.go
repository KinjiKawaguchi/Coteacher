// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/form"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/google/uuid"
)

// FormCreate is the builder for creating a Form entity.
type FormCreate struct {
	config
	mutation *FormMutation
	hooks    []Hook
}

// SetClassID sets the "class_id" field.
func (fc *FormCreate) SetClassID(u uuid.UUID) *FormCreate {
	fc.mutation.SetClassID(u)
	return fc
}

// SetName sets the "name" field.
func (fc *FormCreate) SetName(s string) *FormCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetDescription sets the "description" field.
func (fc *FormCreate) SetDescription(s string) *FormCreate {
	fc.mutation.SetDescription(s)
	return fc
}

// SetSystemPrompt sets the "system__prompt" field.
func (fc *FormCreate) SetSystemPrompt(s string) *FormCreate {
	fc.mutation.SetSystemPrompt(s)
	return fc
}

// SetNillableSystemPrompt sets the "system__prompt" field if the given value is not nil.
func (fc *FormCreate) SetNillableSystemPrompt(s *string) *FormCreate {
	if s != nil {
		fc.SetSystemPrompt(*s)
	}
	return fc
}

// SetUsageLimit sets the "usage_limit" field.
func (fc *FormCreate) SetUsageLimit(i int) *FormCreate {
	fc.mutation.SetUsageLimit(i)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FormCreate) SetCreatedAt(t time.Time) *FormCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FormCreate) SetNillableCreatedAt(t *time.Time) *FormCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FormCreate) SetUpdatedAt(t time.Time) *FormCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FormCreate) SetNillableUpdatedAt(t *time.Time) *FormCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FormCreate) SetID(u uuid.UUID) *FormCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FormCreate) SetNillableID(u *uuid.UUID) *FormCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetClass sets the "class" edge to the Class entity.
func (fc *FormCreate) SetClass(c *Class) *FormCreate {
	return fc.SetClassID(c.ID)
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (fc *FormCreate) AddQuestionIDs(ids ...uuid.UUID) *FormCreate {
	fc.mutation.AddQuestionIDs(ids...)
	return fc
}

// AddQuestions adds the "questions" edges to the Question entity.
func (fc *FormCreate) AddQuestions(q ...*Question) *FormCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return fc.AddQuestionIDs(ids...)
}

// AddResponseIDs adds the "responses" edge to the Response entity by IDs.
func (fc *FormCreate) AddResponseIDs(ids ...uuid.UUID) *FormCreate {
	fc.mutation.AddResponseIDs(ids...)
	return fc
}

// AddResponses adds the "responses" edges to the Response entity.
func (fc *FormCreate) AddResponses(r ...*Response) *FormCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fc.AddResponseIDs(ids...)
}

// Mutation returns the FormMutation object of the builder.
func (fc *FormCreate) Mutation() *FormMutation {
	return fc.mutation
}

// Save creates the Form in the database.
func (fc *FormCreate) Save(ctx context.Context) (*Form, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FormCreate) SaveX(ctx context.Context) *Form {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FormCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FormCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FormCreate) defaults() {
	if _, ok := fc.mutation.SystemPrompt(); !ok {
		v := form.DefaultSystemPrompt
		fc.mutation.SetSystemPrompt(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := form.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := form.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := form.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FormCreate) check() error {
	if _, ok := fc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class_id", err: errors.New(`ent: missing required field "Form.class_id"`)}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Form.name"`)}
	}
	if _, ok := fc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Form.description"`)}
	}
	if _, ok := fc.mutation.SystemPrompt(); !ok {
		return &ValidationError{Name: "system__prompt", err: errors.New(`ent: missing required field "Form.system__prompt"`)}
	}
	if _, ok := fc.mutation.UsageLimit(); !ok {
		return &ValidationError{Name: "usage_limit", err: errors.New(`ent: missing required field "Form.usage_limit"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Form.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Form.updated_at"`)}
	}
	if _, ok := fc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required edge "Form.class"`)}
	}
	return nil
}

func (fc *FormCreate) sqlSave(ctx context.Context) (*Form, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FormCreate) createSpec() (*Form, *sqlgraph.CreateSpec) {
	var (
		_node = &Form{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(form.Table, sqlgraph.NewFieldSpec(form.FieldID, field.TypeUUID))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(form.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Description(); ok {
		_spec.SetField(form.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := fc.mutation.SystemPrompt(); ok {
		_spec.SetField(form.FieldSystemPrompt, field.TypeString, value)
		_node.SystemPrompt = value
	}
	if value, ok := fc.mutation.UsageLimit(); ok {
		_spec.SetField(form.FieldUsageLimit, field.TypeInt, value)
		_node.UsageLimit = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(form.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(form.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   form.ClassTable,
			Columns: []string{form.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   form.QuestionsTable,
			Columns: []string{form.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   form.ResponsesTable,
			Columns: []string{form.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FormCreateBulk is the builder for creating many Form entities in bulk.
type FormCreateBulk struct {
	config
	err      error
	builders []*FormCreate
}

// Save creates the Form entities in the database.
func (fcb *FormCreateBulk) Save(ctx context.Context) ([]*Form, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Form, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FormMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FormCreateBulk) SaveX(ctx context.Context) []*Form {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FormCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FormCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
