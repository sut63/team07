// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/predicate"
)

// InspectionResultUpdate is the builder for updating InspectionResult entities.
type InspectionResultUpdate struct {
	config
	hooks      []Hook
	mutation   *InspectionResultMutation
	predicates []predicate.InspectionResult
}

// Where adds a new predicate for the builder.
func (iru *InspectionResultUpdate) Where(ps ...predicate.InspectionResult) *InspectionResultUpdate {
	iru.predicates = append(iru.predicates, ps...)
	return iru
}

// SetResultName sets the result_name field.
func (iru *InspectionResultUpdate) SetResultName(s string) *InspectionResultUpdate {
	iru.mutation.SetResultName(s)
	return iru
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (iru *InspectionResultUpdate) AddCarinspectionIDs(ids ...int) *InspectionResultUpdate {
	iru.mutation.AddCarinspectionIDs(ids...)
	return iru
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (iru *InspectionResultUpdate) AddCarinspections(c ...*CarInspection) *InspectionResultUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iru.AddCarinspectionIDs(ids...)
}

// AddStatusofIDs adds the statusof edge to Ambulance by ids.
func (iru *InspectionResultUpdate) AddStatusofIDs(ids ...int) *InspectionResultUpdate {
	iru.mutation.AddStatusofIDs(ids...)
	return iru
}

// AddStatusof adds the statusof edges to Ambulance.
func (iru *InspectionResultUpdate) AddStatusof(a ...*Ambulance) *InspectionResultUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iru.AddStatusofIDs(ids...)
}

// Mutation returns the InspectionResultMutation object of the builder.
func (iru *InspectionResultUpdate) Mutation() *InspectionResultMutation {
	return iru.mutation
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (iru *InspectionResultUpdate) RemoveCarinspectionIDs(ids ...int) *InspectionResultUpdate {
	iru.mutation.RemoveCarinspectionIDs(ids...)
	return iru
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (iru *InspectionResultUpdate) RemoveCarinspections(c ...*CarInspection) *InspectionResultUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iru.RemoveCarinspectionIDs(ids...)
}

// RemoveStatusofIDs removes the statusof edge to Ambulance by ids.
func (iru *InspectionResultUpdate) RemoveStatusofIDs(ids ...int) *InspectionResultUpdate {
	iru.mutation.RemoveStatusofIDs(ids...)
	return iru
}

// RemoveStatusof removes statusof edges to Ambulance.
func (iru *InspectionResultUpdate) RemoveStatusof(a ...*Ambulance) *InspectionResultUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iru.RemoveStatusofIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iru *InspectionResultUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := iru.mutation.ResultName(); ok {
		if err := inspectionresult.ResultNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "result_name", err: fmt.Errorf("ent: validator failed for field \"result_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(iru.hooks) == 0 {
		affected, err = iru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InspectionResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iru.mutation = mutation
			affected, err = iru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iru.hooks) - 1; i >= 0; i-- {
			mut = iru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iru *InspectionResultUpdate) SaveX(ctx context.Context) int {
	affected, err := iru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iru *InspectionResultUpdate) Exec(ctx context.Context) error {
	_, err := iru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iru *InspectionResultUpdate) ExecX(ctx context.Context) {
	if err := iru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iru *InspectionResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inspectionresult.Table,
			Columns: inspectionresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inspectionresult.FieldID,
			},
		},
	}
	if ps := iru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iru.mutation.ResultName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inspectionresult.FieldResultName,
		})
	}
	if nodes := iru.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.CarinspectionsTable,
			Columns: []string{inspectionresult.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iru.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.CarinspectionsTable,
			Columns: []string{inspectionresult.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := iru.mutation.RemovedStatusofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.StatusofTable,
			Columns: []string{inspectionresult.StatusofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iru.mutation.StatusofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.StatusofTable,
			Columns: []string{inspectionresult.StatusofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inspectionresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InspectionResultUpdateOne is the builder for updating a single InspectionResult entity.
type InspectionResultUpdateOne struct {
	config
	hooks    []Hook
	mutation *InspectionResultMutation
}

// SetResultName sets the result_name field.
func (iruo *InspectionResultUpdateOne) SetResultName(s string) *InspectionResultUpdateOne {
	iruo.mutation.SetResultName(s)
	return iruo
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (iruo *InspectionResultUpdateOne) AddCarinspectionIDs(ids ...int) *InspectionResultUpdateOne {
	iruo.mutation.AddCarinspectionIDs(ids...)
	return iruo
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (iruo *InspectionResultUpdateOne) AddCarinspections(c ...*CarInspection) *InspectionResultUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iruo.AddCarinspectionIDs(ids...)
}

// AddStatusofIDs adds the statusof edge to Ambulance by ids.
func (iruo *InspectionResultUpdateOne) AddStatusofIDs(ids ...int) *InspectionResultUpdateOne {
	iruo.mutation.AddStatusofIDs(ids...)
	return iruo
}

// AddStatusof adds the statusof edges to Ambulance.
func (iruo *InspectionResultUpdateOne) AddStatusof(a ...*Ambulance) *InspectionResultUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iruo.AddStatusofIDs(ids...)
}

// Mutation returns the InspectionResultMutation object of the builder.
func (iruo *InspectionResultUpdateOne) Mutation() *InspectionResultMutation {
	return iruo.mutation
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (iruo *InspectionResultUpdateOne) RemoveCarinspectionIDs(ids ...int) *InspectionResultUpdateOne {
	iruo.mutation.RemoveCarinspectionIDs(ids...)
	return iruo
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (iruo *InspectionResultUpdateOne) RemoveCarinspections(c ...*CarInspection) *InspectionResultUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iruo.RemoveCarinspectionIDs(ids...)
}

// RemoveStatusofIDs removes the statusof edge to Ambulance by ids.
func (iruo *InspectionResultUpdateOne) RemoveStatusofIDs(ids ...int) *InspectionResultUpdateOne {
	iruo.mutation.RemoveStatusofIDs(ids...)
	return iruo
}

// RemoveStatusof removes statusof edges to Ambulance.
func (iruo *InspectionResultUpdateOne) RemoveStatusof(a ...*Ambulance) *InspectionResultUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iruo.RemoveStatusofIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (iruo *InspectionResultUpdateOne) Save(ctx context.Context) (*InspectionResult, error) {
	if v, ok := iruo.mutation.ResultName(); ok {
		if err := inspectionresult.ResultNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "result_name", err: fmt.Errorf("ent: validator failed for field \"result_name\": %w", err)}
		}
	}

	var (
		err  error
		node *InspectionResult
	)
	if len(iruo.hooks) == 0 {
		node, err = iruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InspectionResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iruo.mutation = mutation
			node, err = iruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iruo.hooks) - 1; i >= 0; i-- {
			mut = iruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iruo *InspectionResultUpdateOne) SaveX(ctx context.Context) *InspectionResult {
	ir, err := iruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ir
}

// Exec executes the query on the entity.
func (iruo *InspectionResultUpdateOne) Exec(ctx context.Context) error {
	_, err := iruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iruo *InspectionResultUpdateOne) ExecX(ctx context.Context) {
	if err := iruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iruo *InspectionResultUpdateOne) sqlSave(ctx context.Context) (ir *InspectionResult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inspectionresult.Table,
			Columns: inspectionresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inspectionresult.FieldID,
			},
		},
	}
	id, ok := iruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing InspectionResult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iruo.mutation.ResultName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inspectionresult.FieldResultName,
		})
	}
	if nodes := iruo.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.CarinspectionsTable,
			Columns: []string{inspectionresult.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iruo.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.CarinspectionsTable,
			Columns: []string{inspectionresult.CarinspectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carinspection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := iruo.mutation.RemovedStatusofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.StatusofTable,
			Columns: []string{inspectionresult.StatusofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iruo.mutation.StatusofIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inspectionresult.StatusofTable,
			Columns: []string{inspectionresult.StatusofColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ambulance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ir = &InspectionResult{config: iruo.config}
	_spec.Assign = ir.assignValues
	_spec.ScanValues = ir.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inspectionresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ir, nil
}
