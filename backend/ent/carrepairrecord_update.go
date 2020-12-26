// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/predicate"
)

// CarRepairrecordUpdate is the builder for updating CarRepairrecord entities.
type CarRepairrecordUpdate struct {
	config
	hooks      []Hook
	mutation   *CarRepairrecordMutation
	predicates []predicate.CarRepairrecord
}

// Where adds a new predicate for the builder.
func (cru *CarRepairrecordUpdate) Where(ps ...predicate.CarRepairrecord) *CarRepairrecordUpdate {
	cru.predicates = append(cru.predicates, ps...)
	return cru
}

// Mutation returns the CarRepairrecordMutation object of the builder.
func (cru *CarRepairrecordUpdate) Mutation() *CarRepairrecordMutation {
	return cru.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cru *CarRepairrecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cru.hooks) == 0 {
		affected, err = cru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarRepairrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cru.mutation = mutation
			affected, err = cru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cru.hooks) - 1; i >= 0; i-- {
			mut = cru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CarRepairrecordUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CarRepairrecordUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CarRepairrecordUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cru *CarRepairrecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   carrepairrecord.Table,
			Columns: carrepairrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carrepairrecord.FieldID,
			},
		},
	}
	if ps := cru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{carrepairrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CarRepairrecordUpdateOne is the builder for updating a single CarRepairrecord entity.
type CarRepairrecordUpdateOne struct {
	config
	hooks    []Hook
	mutation *CarRepairrecordMutation
}

// Mutation returns the CarRepairrecordMutation object of the builder.
func (cruo *CarRepairrecordUpdateOne) Mutation() *CarRepairrecordMutation {
	return cruo.mutation
}

// Save executes the query and returns the updated entity.
func (cruo *CarRepairrecordUpdateOne) Save(ctx context.Context) (*CarRepairrecord, error) {
	var (
		err  error
		node *CarRepairrecord
	)
	if len(cruo.hooks) == 0 {
		node, err = cruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarRepairrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cruo.mutation = mutation
			node, err = cruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cruo.hooks) - 1; i >= 0; i-- {
			mut = cruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CarRepairrecordUpdateOne) SaveX(ctx context.Context) *CarRepairrecord {
	cr, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cr
}

// Exec executes the query on the entity.
func (cruo *CarRepairrecordUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CarRepairrecordUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cruo *CarRepairrecordUpdateOne) sqlSave(ctx context.Context) (cr *CarRepairrecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   carrepairrecord.Table,
			Columns: carrepairrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carrepairrecord.FieldID,
			},
		},
	}
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CarRepairrecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	cr = &CarRepairrecord{config: cruo.config}
	_spec.Assign = cr.assignValues
	_spec.ScanValues = cr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{carrepairrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cr, nil
}