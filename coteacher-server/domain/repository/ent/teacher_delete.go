// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coteacher/domain/repository/ent/predicate"
	"coteacher/domain/repository/ent/teacher"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeacherDelete is the builder for deleting a Teacher entity.
type TeacherDelete struct {
	config
	hooks    []Hook
	mutation *TeacherMutation
}

// Where appends a list predicates to the TeacherDelete builder.
func (td *TeacherDelete) Where(ps ...predicate.Teacher) *TeacherDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TeacherDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TeacherDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TeacherDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(teacher.Table, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeString))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TeacherDeleteOne is the builder for deleting a single Teacher entity.
type TeacherDeleteOne struct {
	td *TeacherDelete
}

// Where appends a list predicates to the TeacherDelete builder.
func (tdo *TeacherDeleteOne) Where(ps ...predicate.Teacher) *TeacherDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TeacherDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teacher.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TeacherDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
