// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/answer"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/response"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/selectedoption"
	"github.com/google/uuid"
)

// AnswerCreate is the builder for creating a Answer entity.
type AnswerCreate struct {
	config
	mutation *AnswerMutation
	hooks    []Hook
}

// SetQuestionID sets the "question_id" field.
func (ac *AnswerCreate) SetQuestionID(u uuid.UUID) *AnswerCreate {
	ac.mutation.SetQuestionID(u)
	return ac
}

// SetResponseID sets the "response_id" field.
func (ac *AnswerCreate) SetResponseID(u uuid.UUID) *AnswerCreate {
	ac.mutation.SetResponseID(u)
	return ac
}

// SetOrder sets the "order" field.
func (ac *AnswerCreate) SetOrder(i int) *AnswerCreate {
	ac.mutation.SetOrder(i)
	return ac
}

// SetAnswerText sets the "answer_text" field.
func (ac *AnswerCreate) SetAnswerText(s string) *AnswerCreate {
	ac.mutation.SetAnswerText(s)
	return ac
}

// SetNillableAnswerText sets the "answer_text" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableAnswerText(s *string) *AnswerCreate {
	if s != nil {
		ac.SetAnswerText(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AnswerCreate) SetID(u uuid.UUID) *AnswerCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableID(u *uuid.UUID) *AnswerCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetQuestion sets the "question" edge to the Question entity.
func (ac *AnswerCreate) SetQuestion(q *Question) *AnswerCreate {
	return ac.SetQuestionID(q.ID)
}

// SetResponse sets the "response" edge to the Response entity.
func (ac *AnswerCreate) SetResponse(r *Response) *AnswerCreate {
	return ac.SetResponseID(r.ID)
}

// AddSelectedOptionIDs adds the "selectedOption" edge to the SelectedOption entity by IDs.
func (ac *AnswerCreate) AddSelectedOptionIDs(ids ...int) *AnswerCreate {
	ac.mutation.AddSelectedOptionIDs(ids...)
	return ac
}

// AddSelectedOption adds the "selectedOption" edges to the SelectedOption entity.
func (ac *AnswerCreate) AddSelectedOption(s ...*SelectedOption) *AnswerCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddSelectedOptionIDs(ids...)
}

// Mutation returns the AnswerMutation object of the builder.
func (ac *AnswerCreate) Mutation() *AnswerMutation {
	return ac.mutation
}

// Save creates the Answer in the database.
func (ac *AnswerCreate) Save(ctx context.Context) (*Answer, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnswerCreate) SaveX(ctx context.Context) *Answer {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnswerCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnswerCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnswerCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := answer.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnswerCreate) check() error {
	if _, ok := ac.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question_id", err: errors.New(`ent: missing required field "Answer.question_id"`)}
	}
	if _, ok := ac.mutation.ResponseID(); !ok {
		return &ValidationError{Name: "response_id", err: errors.New(`ent: missing required field "Answer.response_id"`)}
	}
	if _, ok := ac.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Answer.order"`)}
	}
	if _, ok := ac.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`ent: missing required edge "Answer.question"`)}
	}
	if _, ok := ac.mutation.ResponseID(); !ok {
		return &ValidationError{Name: "response", err: errors.New(`ent: missing required edge "Answer.response"`)}
	}
	return nil
}

func (ac *AnswerCreate) sqlSave(ctx context.Context) (*Answer, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AnswerCreate) createSpec() (*Answer, *sqlgraph.CreateSpec) {
	var (
		_node = &Answer{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(answer.Table, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Order(); ok {
		_spec.SetField(answer.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := ac.mutation.AnswerText(); ok {
		_spec.SetField(answer.FieldAnswerText, field.TypeString, value)
		_node.AnswerText = value
	}
	if nodes := ac.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.QuestionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.ResponseTable,
			Columns: []string{answer.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ResponseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SelectedOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   answer.SelectedOptionTable,
			Columns: []string{answer.SelectedOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(selectedoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AnswerCreateBulk is the builder for creating many Answer entities in bulk.
type AnswerCreateBulk struct {
	config
	err      error
	builders []*AnswerCreate
}

// Save creates the Answer entities in the database.
func (acb *AnswerCreateBulk) Save(ctx context.Context) ([]*Answer, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Answer, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnswerMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnswerCreateBulk) SaveX(ctx context.Context) []*Answer {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnswerCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnswerCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
