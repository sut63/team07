// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/distances"
	"github.com/team07/app/ent/predicate"
)

// DistancesUpdate is the builder for updating Distances entities.
type DistancesUpdate struct {
	config
	hooks      []Hook
	mutation   *DistancesMutation
	predicates []predicate.Distances
}

// Where adds a new predicate for the builder.
func (du *DistancesUpdate) Where(ps ...predicate.Distances) *DistancesUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDistances sets the Distances field.
func (du *DistancesUpdate) SetDistances(s string) *DistancesUpdate {
	du.mutation.SetDistances(s)
	return du
}

// AddDisidIDs adds the disid edge to Carservice by ids.
func (du *DistancesUpdate) AddDisidIDs(ids ...int) *DistancesUpdate {
	du.mutation.AddDisidIDs(ids...)
	return du
}

// AddDisid adds the disid edges to Carservice.
func (du *DistancesUpdate) AddDisid(c ...*Carservice) *DistancesUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddDisidIDs(ids...)
}

// Mutation returns the DistancesMutation object of the builder.
func (du *DistancesUpdate) Mutation() *DistancesMutation {
	return du.mutation
}

// RemoveDisidIDs removes the disid edge to Carservice by ids.
func (du *DistancesUpdate) RemoveDisidIDs(ids ...int) *DistancesUpdate {
	du.mutation.RemoveDisidIDs(ids...)
	return du
}

// RemoveDisid removes disid edges to Carservice.
func (du *DistancesUpdate) RemoveDisid(c ...*Carservice) *DistancesUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveDisidIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DistancesUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Distances(); ok {
		if err := distances.DistancesValidator(v); err != nil {
			return 0, &ValidationError{Name: "Distances", err: fmt.Errorf("ent: validator failed for field \"Distances\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DistancesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DistancesUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DistancesUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DistancesUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DistancesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distances.Table,
			Columns: distances.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distances.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Distances(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: distances.FieldDistances,
		})
	}
	if nodes := du.mutation.RemovedDisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distances.DisidTable,
			Columns: []string{distances.DisidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distances.DisidTable,
			Columns: []string{distances.DisidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{distances.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DistancesUpdateOne is the builder for updating a single Distances entity.
type DistancesUpdateOne struct {
	config
	hooks    []Hook
	mutation *DistancesMutation
}

// SetDistances sets the Distances field.
func (duo *DistancesUpdateOne) SetDistances(s string) *DistancesUpdateOne {
	duo.mutation.SetDistances(s)
	return duo
}

// AddDisidIDs adds the disid edge to Carservice by ids.
func (duo *DistancesUpdateOne) AddDisidIDs(ids ...int) *DistancesUpdateOne {
	duo.mutation.AddDisidIDs(ids...)
	return duo
}

// AddDisid adds the disid edges to Carservice.
func (duo *DistancesUpdateOne) AddDisid(c ...*Carservice) *DistancesUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddDisidIDs(ids...)
}

// Mutation returns the DistancesMutation object of the builder.
func (duo *DistancesUpdateOne) Mutation() *DistancesMutation {
	return duo.mutation
}

// RemoveDisidIDs removes the disid edge to Carservice by ids.
func (duo *DistancesUpdateOne) RemoveDisidIDs(ids ...int) *DistancesUpdateOne {
	duo.mutation.RemoveDisidIDs(ids...)
	return duo
}

// RemoveDisid removes disid edges to Carservice.
func (duo *DistancesUpdateOne) RemoveDisid(c ...*Carservice) *DistancesUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveDisidIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DistancesUpdateOne) Save(ctx context.Context) (*Distances, error) {
	if v, ok := duo.mutation.Distances(); ok {
		if err := distances.DistancesValidator(v); err != nil {
			return nil, &ValidationError{Name: "Distances", err: fmt.Errorf("ent: validator failed for field \"Distances\": %w", err)}
		}
	}

	var (
		err  error
		node *Distances
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DistancesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DistancesUpdateOne) SaveX(ctx context.Context) *Distances {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DistancesUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DistancesUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DistancesUpdateOne) sqlSave(ctx context.Context) (d *Distances, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distances.Table,
			Columns: distances.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distances.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Distances.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Distances(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: distances.FieldDistances,
		})
	}
	if nodes := duo.mutation.RemovedDisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distances.DisidTable,
			Columns: []string{distances.DisidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distances.DisidTable,
			Columns: []string{distances.DisidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Distances{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{distances.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}