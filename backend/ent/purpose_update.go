// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/purpose"
)

// PurposeUpdate is the builder for updating Purpose entities.
type PurposeUpdate struct {
	config
	hooks      []Hook
	mutation   *PurposeMutation
	predicates []predicate.Purpose
}

// Where adds a new predicate for the builder.
func (pu *PurposeUpdate) Where(ps ...predicate.Purpose) *PurposeUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetObjective sets the objective field.
func (pu *PurposeUpdate) SetObjective(s string) *PurposeUpdate {
	pu.mutation.SetObjective(s)
	return pu
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (pu *PurposeUpdate) AddCarcheckinoutIDs(ids ...int) *PurposeUpdate {
	pu.mutation.AddCarcheckinoutIDs(ids...)
	return pu
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (pu *PurposeUpdate) AddCarcheckinout(c ...*CarCheckInOut) *PurposeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCarcheckinoutIDs(ids...)
}

// Mutation returns the PurposeMutation object of the builder.
func (pu *PurposeUpdate) Mutation() *PurposeMutation {
	return pu.mutation
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (pu *PurposeUpdate) RemoveCarcheckinoutIDs(ids ...int) *PurposeUpdate {
	pu.mutation.RemoveCarcheckinoutIDs(ids...)
	return pu
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (pu *PurposeUpdate) RemoveCarcheckinout(c ...*CarCheckInOut) *PurposeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCarcheckinoutIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PurposeUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurposeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PurposeUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PurposeUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PurposeUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PurposeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   purpose.Table,
			Columns: purpose.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: purpose.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Objective(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: purpose.FieldObjective,
		})
	}
	if nodes := pu.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   purpose.CarcheckinoutTable,
			Columns: []string{purpose.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   purpose.CarcheckinoutTable,
			Columns: []string{purpose.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purpose.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PurposeUpdateOne is the builder for updating a single Purpose entity.
type PurposeUpdateOne struct {
	config
	hooks    []Hook
	mutation *PurposeMutation
}

// SetObjective sets the objective field.
func (puo *PurposeUpdateOne) SetObjective(s string) *PurposeUpdateOne {
	puo.mutation.SetObjective(s)
	return puo
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (puo *PurposeUpdateOne) AddCarcheckinoutIDs(ids ...int) *PurposeUpdateOne {
	puo.mutation.AddCarcheckinoutIDs(ids...)
	return puo
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (puo *PurposeUpdateOne) AddCarcheckinout(c ...*CarCheckInOut) *PurposeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCarcheckinoutIDs(ids...)
}

// Mutation returns the PurposeMutation object of the builder.
func (puo *PurposeUpdateOne) Mutation() *PurposeMutation {
	return puo.mutation
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (puo *PurposeUpdateOne) RemoveCarcheckinoutIDs(ids ...int) *PurposeUpdateOne {
	puo.mutation.RemoveCarcheckinoutIDs(ids...)
	return puo
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (puo *PurposeUpdateOne) RemoveCarcheckinout(c ...*CarCheckInOut) *PurposeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCarcheckinoutIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PurposeUpdateOne) Save(ctx context.Context) (*Purpose, error) {

	var (
		err  error
		node *Purpose
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurposeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PurposeUpdateOne) SaveX(ctx context.Context) *Purpose {
	pu, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pu
}

// Exec executes the query on the entity.
func (puo *PurposeUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PurposeUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PurposeUpdateOne) sqlSave(ctx context.Context) (pu *Purpose, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   purpose.Table,
			Columns: purpose.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: purpose.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Purpose.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Objective(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: purpose.FieldObjective,
		})
	}
	if nodes := puo.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   purpose.CarcheckinoutTable,
			Columns: []string{purpose.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   purpose.CarcheckinoutTable,
			Columns: []string{purpose.CarcheckinoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carcheckinout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pu = &Purpose{config: puo.config}
	_spec.Assign = pu.assignValues
	_spec.ScanValues = pu.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purpose.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pu, nil
}
