// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/distance"
	"github.com/team07/app/ent/predicate"
)

// DistanceUpdate is the builder for updating Distance entities.
type DistanceUpdate struct {
	config
	hooks      []Hook
	mutation   *DistanceMutation
	predicates []predicate.Distance
}

// Where adds a new predicate for the builder.
func (du *DistanceUpdate) Where(ps ...predicate.Distance) *DistanceUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDistance sets the Distance field.
func (du *DistanceUpdate) SetDistance(s string) *DistanceUpdate {
	du.mutation.SetDistance(s)
	return du
}

// AddDisidIDs adds the disid edge to Carservice by ids.
func (du *DistanceUpdate) AddDisidIDs(ids ...int) *DistanceUpdate {
	du.mutation.AddDisidIDs(ids...)
	return du
}

// AddDisid adds the disid edges to Carservice.
func (du *DistanceUpdate) AddDisid(c ...*Carservice) *DistanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddDisidIDs(ids...)
}

// Mutation returns the DistanceMutation object of the builder.
func (du *DistanceUpdate) Mutation() *DistanceMutation {
	return du.mutation
}

// RemoveDisidIDs removes the disid edge to Carservice by ids.
func (du *DistanceUpdate) RemoveDisidIDs(ids ...int) *DistanceUpdate {
	du.mutation.RemoveDisidIDs(ids...)
	return du
}

// RemoveDisid removes disid edges to Carservice.
func (du *DistanceUpdate) RemoveDisid(c ...*Carservice) *DistanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveDisidIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DistanceUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Distance(); ok {
		if err := distance.DistanceValidator(v); err != nil {
			return 0, &ValidationError{Name: "Distance", err: fmt.Errorf("ent: validator failed for field \"Distance\": %w", err)}
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
			mutation, ok := m.(*DistanceMutation)
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
func (du *DistanceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DistanceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DistanceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DistanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distance.Table,
			Columns: distance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distance.FieldID,
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
	if value, ok := du.mutation.Distance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: distance.FieldDistance,
		})
	}
	if nodes := du.mutation.RemovedDisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distance.DisidTable,
			Columns: []string{distance.DisidColumn},
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
			Table:   distance.DisidTable,
			Columns: []string{distance.DisidColumn},
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
			err = &NotFoundError{distance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DistanceUpdateOne is the builder for updating a single Distance entity.
type DistanceUpdateOne struct {
	config
	hooks    []Hook
	mutation *DistanceMutation
}

// SetDistance sets the Distance field.
func (duo *DistanceUpdateOne) SetDistance(s string) *DistanceUpdateOne {
	duo.mutation.SetDistance(s)
	return duo
}

// AddDisidIDs adds the disid edge to Carservice by ids.
func (duo *DistanceUpdateOne) AddDisidIDs(ids ...int) *DistanceUpdateOne {
	duo.mutation.AddDisidIDs(ids...)
	return duo
}

// AddDisid adds the disid edges to Carservice.
func (duo *DistanceUpdateOne) AddDisid(c ...*Carservice) *DistanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddDisidIDs(ids...)
}

// Mutation returns the DistanceMutation object of the builder.
func (duo *DistanceUpdateOne) Mutation() *DistanceMutation {
	return duo.mutation
}

// RemoveDisidIDs removes the disid edge to Carservice by ids.
func (duo *DistanceUpdateOne) RemoveDisidIDs(ids ...int) *DistanceUpdateOne {
	duo.mutation.RemoveDisidIDs(ids...)
	return duo
}

// RemoveDisid removes disid edges to Carservice.
func (duo *DistanceUpdateOne) RemoveDisid(c ...*Carservice) *DistanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveDisidIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DistanceUpdateOne) Save(ctx context.Context) (*Distance, error) {
	if v, ok := duo.mutation.Distance(); ok {
		if err := distance.DistanceValidator(v); err != nil {
			return nil, &ValidationError{Name: "Distance", err: fmt.Errorf("ent: validator failed for field \"Distance\": %w", err)}
		}
	}

	var (
		err  error
		node *Distance
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DistanceMutation)
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
func (duo *DistanceUpdateOne) SaveX(ctx context.Context) *Distance {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DistanceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DistanceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DistanceUpdateOne) sqlSave(ctx context.Context) (d *Distance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distance.Table,
			Columns: distance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distance.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Distance.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Distance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: distance.FieldDistance,
		})
	}
	if nodes := duo.mutation.RemovedDisidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   distance.DisidTable,
			Columns: []string{distance.DisidColumn},
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
			Table:   distance.DisidTable,
			Columns: []string{distance.DisidColumn},
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
	d = &Distance{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{distance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
