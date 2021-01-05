// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/user"
)

// JobPositionUpdate is the builder for updating JobPosition entities.
type JobPositionUpdate struct {
	config
	hooks      []Hook
	mutation   *JobPositionMutation
	predicates []predicate.JobPosition
}

// Where adds a new predicate for the builder.
func (jpu *JobPositionUpdate) Where(ps ...predicate.JobPosition) *JobPositionUpdate {
	jpu.predicates = append(jpu.predicates, ps...)
	return jpu
}

// SetPositionName sets the position_name field.
func (jpu *JobPositionUpdate) SetPositionName(s string) *JobPositionUpdate {
	jpu.mutation.SetPositionName(s)
	return jpu
}

// AddUserIDs adds the users edge to User by ids.
func (jpu *JobPositionUpdate) AddUserIDs(ids ...int) *JobPositionUpdate {
	jpu.mutation.AddUserIDs(ids...)
	return jpu
}

// AddUsers adds the users edges to User.
func (jpu *JobPositionUpdate) AddUsers(u ...*User) *JobPositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jpu.AddUserIDs(ids...)
}

// AddInspectionresultIDs adds the inspectionresults edge to InspectionResult by ids.
func (jpu *JobPositionUpdate) AddInspectionresultIDs(ids ...int) *JobPositionUpdate {
	jpu.mutation.AddInspectionresultIDs(ids...)
	return jpu
}

// AddInspectionresults adds the inspectionresults edges to InspectionResult.
func (jpu *JobPositionUpdate) AddInspectionresults(i ...*InspectionResult) *JobPositionUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return jpu.AddInspectionresultIDs(ids...)
}

// Mutation returns the JobPositionMutation object of the builder.
func (jpu *JobPositionUpdate) Mutation() *JobPositionMutation {
	return jpu.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (jpu *JobPositionUpdate) RemoveUserIDs(ids ...int) *JobPositionUpdate {
	jpu.mutation.RemoveUserIDs(ids...)
	return jpu
}

// RemoveUsers removes users edges to User.
func (jpu *JobPositionUpdate) RemoveUsers(u ...*User) *JobPositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jpu.RemoveUserIDs(ids...)
}

// RemoveInspectionresultIDs removes the inspectionresults edge to InspectionResult by ids.
func (jpu *JobPositionUpdate) RemoveInspectionresultIDs(ids ...int) *JobPositionUpdate {
	jpu.mutation.RemoveInspectionresultIDs(ids...)
	return jpu
}

// RemoveInspectionresults removes inspectionresults edges to InspectionResult.
func (jpu *JobPositionUpdate) RemoveInspectionresults(i ...*InspectionResult) *JobPositionUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return jpu.RemoveInspectionresultIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (jpu *JobPositionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := jpu.mutation.PositionName(); ok {
		if err := jobposition.PositionNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "position_name", err: fmt.Errorf("ent: validator failed for field \"position_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(jpu.hooks) == 0 {
		affected, err = jpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jpu.mutation = mutation
			affected, err = jpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jpu.hooks) - 1; i >= 0; i-- {
			mut = jpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (jpu *JobPositionUpdate) SaveX(ctx context.Context) int {
	affected, err := jpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jpu *JobPositionUpdate) Exec(ctx context.Context) error {
	_, err := jpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpu *JobPositionUpdate) ExecX(ctx context.Context) {
	if err := jpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jpu *JobPositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobposition.Table,
			Columns: jobposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobposition.FieldID,
			},
		},
	}
	if ps := jpu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jpu.mutation.PositionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: jobposition.FieldPositionName,
		})
	}
	if nodes := jpu.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.UsersTable,
			Columns: []string{jobposition.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.UsersTable,
			Columns: []string{jobposition.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := jpu.mutation.RemovedInspectionresultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.InspectionresultsTable,
			Columns: []string{jobposition.InspectionresultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpu.mutation.InspectionresultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.InspectionresultsTable,
			Columns: []string{jobposition.InspectionresultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobposition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// JobPositionUpdateOne is the builder for updating a single JobPosition entity.
type JobPositionUpdateOne struct {
	config
	hooks    []Hook
	mutation *JobPositionMutation
}

// SetPositionName sets the position_name field.
func (jpuo *JobPositionUpdateOne) SetPositionName(s string) *JobPositionUpdateOne {
	jpuo.mutation.SetPositionName(s)
	return jpuo
}

// AddUserIDs adds the users edge to User by ids.
func (jpuo *JobPositionUpdateOne) AddUserIDs(ids ...int) *JobPositionUpdateOne {
	jpuo.mutation.AddUserIDs(ids...)
	return jpuo
}

// AddUsers adds the users edges to User.
func (jpuo *JobPositionUpdateOne) AddUsers(u ...*User) *JobPositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jpuo.AddUserIDs(ids...)
}

// AddInspectionresultIDs adds the inspectionresults edge to InspectionResult by ids.
func (jpuo *JobPositionUpdateOne) AddInspectionresultIDs(ids ...int) *JobPositionUpdateOne {
	jpuo.mutation.AddInspectionresultIDs(ids...)
	return jpuo
}

// AddInspectionresults adds the inspectionresults edges to InspectionResult.
func (jpuo *JobPositionUpdateOne) AddInspectionresults(i ...*InspectionResult) *JobPositionUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return jpuo.AddInspectionresultIDs(ids...)
}

// Mutation returns the JobPositionMutation object of the builder.
func (jpuo *JobPositionUpdateOne) Mutation() *JobPositionMutation {
	return jpuo.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (jpuo *JobPositionUpdateOne) RemoveUserIDs(ids ...int) *JobPositionUpdateOne {
	jpuo.mutation.RemoveUserIDs(ids...)
	return jpuo
}

// RemoveUsers removes users edges to User.
func (jpuo *JobPositionUpdateOne) RemoveUsers(u ...*User) *JobPositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jpuo.RemoveUserIDs(ids...)
}

// RemoveInspectionresultIDs removes the inspectionresults edge to InspectionResult by ids.
func (jpuo *JobPositionUpdateOne) RemoveInspectionresultIDs(ids ...int) *JobPositionUpdateOne {
	jpuo.mutation.RemoveInspectionresultIDs(ids...)
	return jpuo
}

// RemoveInspectionresults removes inspectionresults edges to InspectionResult.
func (jpuo *JobPositionUpdateOne) RemoveInspectionresults(i ...*InspectionResult) *JobPositionUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return jpuo.RemoveInspectionresultIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (jpuo *JobPositionUpdateOne) Save(ctx context.Context) (*JobPosition, error) {
	if v, ok := jpuo.mutation.PositionName(); ok {
		if err := jobposition.PositionNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "position_name", err: fmt.Errorf("ent: validator failed for field \"position_name\": %w", err)}
		}
	}

	var (
		err  error
		node *JobPosition
	)
	if len(jpuo.hooks) == 0 {
		node, err = jpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jpuo.mutation = mutation
			node, err = jpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jpuo.hooks) - 1; i >= 0; i-- {
			mut = jpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (jpuo *JobPositionUpdateOne) SaveX(ctx context.Context) *JobPosition {
	jp, err := jpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return jp
}

// Exec executes the query on the entity.
func (jpuo *JobPositionUpdateOne) Exec(ctx context.Context) error {
	_, err := jpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpuo *JobPositionUpdateOne) ExecX(ctx context.Context) {
	if err := jpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jpuo *JobPositionUpdateOne) sqlSave(ctx context.Context) (jp *JobPosition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobposition.Table,
			Columns: jobposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobposition.FieldID,
			},
		},
	}
	id, ok := jpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing JobPosition.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := jpuo.mutation.PositionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: jobposition.FieldPositionName,
		})
	}
	if nodes := jpuo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.UsersTable,
			Columns: []string{jobposition.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.UsersTable,
			Columns: []string{jobposition.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := jpuo.mutation.RemovedInspectionresultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.InspectionresultsTable,
			Columns: []string{jobposition.InspectionresultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpuo.mutation.InspectionresultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.InspectionresultsTable,
			Columns: []string{jobposition.InspectionresultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	jp = &JobPosition{config: jpuo.config}
	_spec.Assign = jp.assignValues
	_spec.ScanValues = jp.scanValues()
	if err = sqlgraph.UpdateNode(ctx, jpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobposition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return jp, nil
}
