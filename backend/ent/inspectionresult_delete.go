// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/predicate"
)

// InspectionResultDelete is the builder for deleting a InspectionResult entity.
type InspectionResultDelete struct {
	config
	hooks      []Hook
	mutation   *InspectionResultMutation
	predicates []predicate.InspectionResult
}

// Where adds a new predicate to the delete builder.
func (ird *InspectionResultDelete) Where(ps ...predicate.InspectionResult) *InspectionResultDelete {
	ird.predicates = append(ird.predicates, ps...)
	return ird
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ird *InspectionResultDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ird.hooks) == 0 {
		affected, err = ird.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InspectionResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ird.mutation = mutation
			affected, err = ird.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ird.hooks) - 1; i >= 0; i-- {
			mut = ird.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ird.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ird *InspectionResultDelete) ExecX(ctx context.Context) int {
	n, err := ird.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ird *InspectionResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: inspectionresult.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inspectionresult.FieldID,
			},
		},
	}
	if ps := ird.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ird.driver, _spec)
}

// InspectionResultDeleteOne is the builder for deleting a single InspectionResult entity.
type InspectionResultDeleteOne struct {
	ird *InspectionResultDelete
}

// Exec executes the deletion query.
func (irdo *InspectionResultDeleteOne) Exec(ctx context.Context) error {
	n, err := irdo.ird.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inspectionresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (irdo *InspectionResultDeleteOne) ExecX(ctx context.Context) {
	irdo.ird.ExecX(ctx)
}