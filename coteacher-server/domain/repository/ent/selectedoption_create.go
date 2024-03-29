// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/answer"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/questionoption"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/selectedoption"
	"github.com/google/uuid"
)

// SelectedOptionCreate is the builder for creating a SelectedOption entity.
type SelectedOptionCreate struct {
	config
	mutation *SelectedOptionMutation
	hooks    []Hook
}

// SetOptionID sets the "option_id" field.
func (soc *SelectedOptionCreate) SetOptionID(u uuid.UUID) *SelectedOptionCreate {
	soc.mutation.SetOptionID(u)
	return soc
}

// SetAnswerID sets the "answer_id" field.
func (soc *SelectedOptionCreate) SetAnswerID(u uuid.UUID) *SelectedOptionCreate {
	soc.mutation.SetAnswerID(u)
	return soc
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (soc *SelectedOptionCreate) SetAnswer(a *Answer) *SelectedOptionCreate {
	return soc.SetAnswerID(a.ID)
}

// SetOption sets the "option" edge to the QuestionOption entity.
func (soc *SelectedOptionCreate) SetOption(q *QuestionOption) *SelectedOptionCreate {
	return soc.SetOptionID(q.ID)
}

// Mutation returns the SelectedOptionMutation object of the builder.
func (soc *SelectedOptionCreate) Mutation() *SelectedOptionMutation {
	return soc.mutation
}

// Save creates the SelectedOption in the database.
func (soc *SelectedOptionCreate) Save(ctx context.Context) (*SelectedOption, error) {
	return withHooks(ctx, soc.sqlSave, soc.mutation, soc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (soc *SelectedOptionCreate) SaveX(ctx context.Context) *SelectedOption {
	v, err := soc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (soc *SelectedOptionCreate) Exec(ctx context.Context) error {
	_, err := soc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (soc *SelectedOptionCreate) ExecX(ctx context.Context) {
	if err := soc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (soc *SelectedOptionCreate) check() error {
	if _, ok := soc.mutation.OptionID(); !ok {
		return &ValidationError{Name: "option_id", err: errors.New(`ent: missing required field "SelectedOption.option_id"`)}
	}
	if _, ok := soc.mutation.AnswerID(); !ok {
		return &ValidationError{Name: "answer_id", err: errors.New(`ent: missing required field "SelectedOption.answer_id"`)}
	}
	if _, ok := soc.mutation.AnswerID(); !ok {
		return &ValidationError{Name: "answer", err: errors.New(`ent: missing required edge "SelectedOption.answer"`)}
	}
	if _, ok := soc.mutation.OptionID(); !ok {
		return &ValidationError{Name: "option", err: errors.New(`ent: missing required edge "SelectedOption.option"`)}
	}
	return nil
}

func (soc *SelectedOptionCreate) sqlSave(ctx context.Context) (*SelectedOption, error) {
	if err := soc.check(); err != nil {
		return nil, err
	}
	_node, _spec := soc.createSpec()
	if err := sqlgraph.CreateNode(ctx, soc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	soc.mutation.id = &_node.ID
	soc.mutation.done = true
	return _node, nil
}

func (soc *SelectedOptionCreate) createSpec() (*SelectedOption, *sqlgraph.CreateSpec) {
	var (
		_node = &SelectedOption{config: soc.config}
		_spec = sqlgraph.NewCreateSpec(selectedoption.Table, sqlgraph.NewFieldSpec(selectedoption.FieldID, field.TypeInt))
	)
	if nodes := soc.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   selectedoption.AnswerTable,
			Columns: []string{selectedoption.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AnswerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := soc.mutation.OptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   selectedoption.OptionTable,
			Columns: []string{selectedoption.OptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OptionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SelectedOptionCreateBulk is the builder for creating many SelectedOption entities in bulk.
type SelectedOptionCreateBulk struct {
	config
	err      error
	builders []*SelectedOptionCreate
}

// Save creates the SelectedOption entities in the database.
func (socb *SelectedOptionCreateBulk) Save(ctx context.Context) ([]*SelectedOption, error) {
	if socb.err != nil {
		return nil, socb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(socb.builders))
	nodes := make([]*SelectedOption, len(socb.builders))
	mutators := make([]Mutator, len(socb.builders))
	for i := range socb.builders {
		func(i int, root context.Context) {
			builder := socb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SelectedOptionMutation)
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
					_, err = mutators[i+1].Mutate(root, socb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, socb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, socb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (socb *SelectedOptionCreateBulk) SaveX(ctx context.Context) []*SelectedOption {
	v, err := socb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (socb *SelectedOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := socb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (socb *SelectedOptionCreateBulk) ExecX(ctx context.Context) {
	if err := socb.Exec(ctx); err != nil {
		panic(err)
	}
}
