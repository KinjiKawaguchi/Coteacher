// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/textquestion"
	"github.com/google/uuid"
)

// TextQuestionCreate is the builder for creating a TextQuestion entity.
type TextQuestionCreate struct {
	config
	mutation *TextQuestionMutation
	hooks    []Hook
}

// SetQuestionID sets the "question_id" field.
func (tqc *TextQuestionCreate) SetQuestionID(u uuid.UUID) *TextQuestionCreate {
	tqc.mutation.SetQuestionID(u)
	return tqc
}

// SetMaxLength sets the "max_length" field.
func (tqc *TextQuestionCreate) SetMaxLength(i int) *TextQuestionCreate {
	tqc.mutation.SetMaxLength(i)
	return tqc
}

// SetID sets the "id" field.
func (tqc *TextQuestionCreate) SetID(u uuid.UUID) *TextQuestionCreate {
	tqc.mutation.SetID(u)
	return tqc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tqc *TextQuestionCreate) SetNillableID(u *uuid.UUID) *TextQuestionCreate {
	if u != nil {
		tqc.SetID(*u)
	}
	return tqc
}

// SetQuestion sets the "question" edge to the Question entity.
func (tqc *TextQuestionCreate) SetQuestion(q *Question) *TextQuestionCreate {
	return tqc.SetQuestionID(q.ID)
}

// Mutation returns the TextQuestionMutation object of the builder.
func (tqc *TextQuestionCreate) Mutation() *TextQuestionMutation {
	return tqc.mutation
}

// Save creates the TextQuestion in the database.
func (tqc *TextQuestionCreate) Save(ctx context.Context) (*TextQuestion, error) {
	tqc.defaults()
	return withHooks(ctx, tqc.sqlSave, tqc.mutation, tqc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tqc *TextQuestionCreate) SaveX(ctx context.Context) *TextQuestion {
	v, err := tqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqc *TextQuestionCreate) Exec(ctx context.Context) error {
	_, err := tqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqc *TextQuestionCreate) ExecX(ctx context.Context) {
	if err := tqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tqc *TextQuestionCreate) defaults() {
	if _, ok := tqc.mutation.ID(); !ok {
		v := textquestion.DefaultID()
		tqc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tqc *TextQuestionCreate) check() error {
	if _, ok := tqc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question_id", err: errors.New(`ent: missing required field "TextQuestion.question_id"`)}
	}
	if _, ok := tqc.mutation.MaxLength(); !ok {
		return &ValidationError{Name: "max_length", err: errors.New(`ent: missing required field "TextQuestion.max_length"`)}
	}
	if _, ok := tqc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`ent: missing required edge "TextQuestion.question"`)}
	}
	return nil
}

func (tqc *TextQuestionCreate) sqlSave(ctx context.Context) (*TextQuestion, error) {
	if err := tqc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tqc.driver, _spec); err != nil {
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
	tqc.mutation.id = &_node.ID
	tqc.mutation.done = true
	return _node, nil
}

func (tqc *TextQuestionCreate) createSpec() (*TextQuestion, *sqlgraph.CreateSpec) {
	var (
		_node = &TextQuestion{config: tqc.config}
		_spec = sqlgraph.NewCreateSpec(textquestion.Table, sqlgraph.NewFieldSpec(textquestion.FieldID, field.TypeUUID))
	)
	if id, ok := tqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tqc.mutation.MaxLength(); ok {
		_spec.SetField(textquestion.FieldMaxLength, field.TypeInt, value)
		_node.MaxLength = value
	}
	if nodes := tqc.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   textquestion.QuestionTable,
			Columns: []string{textquestion.QuestionColumn},
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
	return _node, _spec
}

// TextQuestionCreateBulk is the builder for creating many TextQuestion entities in bulk.
type TextQuestionCreateBulk struct {
	config
	err      error
	builders []*TextQuestionCreate
}

// Save creates the TextQuestion entities in the database.
func (tqcb *TextQuestionCreateBulk) Save(ctx context.Context) ([]*TextQuestion, error) {
	if tqcb.err != nil {
		return nil, tqcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tqcb.builders))
	nodes := make([]*TextQuestion, len(tqcb.builders))
	mutators := make([]Mutator, len(tqcb.builders))
	for i := range tqcb.builders {
		func(i int, root context.Context) {
			builder := tqcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TextQuestionMutation)
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
					_, err = mutators[i+1].Mutate(root, tqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tqcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tqcb *TextQuestionCreateBulk) SaveX(ctx context.Context) []*TextQuestion {
	v, err := tqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tqcb *TextQuestionCreateBulk) Exec(ctx context.Context) error {
	_, err := tqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqcb *TextQuestionCreateBulk) ExecX(ctx context.Context) {
	if err := tqcb.Exec(ctx); err != nil {
		panic(err)
	}
}
