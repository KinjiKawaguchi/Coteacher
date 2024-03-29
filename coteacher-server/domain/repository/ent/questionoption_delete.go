// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/questionoption"
)

// QuestionOptionDelete is the builder for deleting a QuestionOption entity.
type QuestionOptionDelete struct {
	config
	hooks    []Hook
	mutation *QuestionOptionMutation
}

// Where appends a list predicates to the QuestionOptionDelete builder.
func (qod *QuestionOptionDelete) Where(ps ...predicate.QuestionOption) *QuestionOptionDelete {
	qod.mutation.Where(ps...)
	return qod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qod *QuestionOptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qod.sqlExec, qod.mutation, qod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qod *QuestionOptionDelete) ExecX(ctx context.Context) int {
	n, err := qod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qod *QuestionOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(questionoption.Table, sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeUUID))
	if ps := qod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qod.mutation.done = true
	return affected, err
}

// QuestionOptionDeleteOne is the builder for deleting a single QuestionOption entity.
type QuestionOptionDeleteOne struct {
	qod *QuestionOptionDelete
}

// Where appends a list predicates to the QuestionOptionDelete builder.
func (qodo *QuestionOptionDeleteOne) Where(ps ...predicate.QuestionOption) *QuestionOptionDeleteOne {
	qodo.qod.mutation.Where(ps...)
	return qodo
}

// Exec executes the deletion query.
func (qodo *QuestionOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := qodo.qod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{questionoption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qodo *QuestionOptionDeleteOne) ExecX(ctx context.Context) {
	if err := qodo.Exec(ctx); err != nil {
		panic(err)
	}
}
