// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carbrand"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/insurance"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/user"
)

// AmbulanceUpdate is the builder for updating Ambulance entities.
type AmbulanceUpdate struct {
	config
	hooks      []Hook
	mutation   *AmbulanceMutation
	predicates []predicate.Ambulance
}

// Where adds a new predicate for the builder.
func (au *AmbulanceUpdate) Where(ps ...predicate.Ambulance) *AmbulanceUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetCarregistration sets the carregistration field.
func (au *AmbulanceUpdate) SetCarregistration(s string) *AmbulanceUpdate {
	au.mutation.SetCarregistration(s)
	return au
}

// SetRegisterat sets the registerat field.
func (au *AmbulanceUpdate) SetRegisterat(t time.Time) *AmbulanceUpdate {
	au.mutation.SetRegisterat(t)
	return au
}

// SetNillableRegisterat sets the registerat field if the given value is not nil.
func (au *AmbulanceUpdate) SetNillableRegisterat(t *time.Time) *AmbulanceUpdate {
	if t != nil {
		au.SetRegisterat(*t)
	}
	return au
}

// SetHasbrandID sets the hasbrand edge to Carbrand by id.
func (au *AmbulanceUpdate) SetHasbrandID(id int) *AmbulanceUpdate {
	au.mutation.SetHasbrandID(id)
	return au
}

// SetNillableHasbrandID sets the hasbrand edge to Carbrand by id if the given value is not nil.
func (au *AmbulanceUpdate) SetNillableHasbrandID(id *int) *AmbulanceUpdate {
	if id != nil {
		au = au.SetHasbrandID(*id)
	}
	return au
}

// SetHasbrand sets the hasbrand edge to Carbrand.
func (au *AmbulanceUpdate) SetHasbrand(c *Carbrand) *AmbulanceUpdate {
	return au.SetHasbrandID(c.ID)
}

// SetHasinsuranceID sets the hasinsurance edge to Insurance by id.
func (au *AmbulanceUpdate) SetHasinsuranceID(id int) *AmbulanceUpdate {
	au.mutation.SetHasinsuranceID(id)
	return au
}

// SetNillableHasinsuranceID sets the hasinsurance edge to Insurance by id if the given value is not nil.
func (au *AmbulanceUpdate) SetNillableHasinsuranceID(id *int) *AmbulanceUpdate {
	if id != nil {
		au = au.SetHasinsuranceID(*id)
	}
	return au
}

// SetHasinsurance sets the hasinsurance edge to Insurance.
func (au *AmbulanceUpdate) SetHasinsurance(i *Insurance) *AmbulanceUpdate {
	return au.SetHasinsuranceID(i.ID)
}

// SetHasstatusID sets the hasstatus edge to InspectionResult by id.
func (au *AmbulanceUpdate) SetHasstatusID(id int) *AmbulanceUpdate {
	au.mutation.SetHasstatusID(id)
	return au
}

// SetNillableHasstatusID sets the hasstatus edge to InspectionResult by id if the given value is not nil.
func (au *AmbulanceUpdate) SetNillableHasstatusID(id *int) *AmbulanceUpdate {
	if id != nil {
		au = au.SetHasstatusID(*id)
	}
	return au
}

// SetHasstatus sets the hasstatus edge to InspectionResult.
func (au *AmbulanceUpdate) SetHasstatus(i *InspectionResult) *AmbulanceUpdate {
	return au.SetHasstatusID(i.ID)
}

// SetHasuserID sets the hasuser edge to User by id.
func (au *AmbulanceUpdate) SetHasuserID(id int) *AmbulanceUpdate {
	au.mutation.SetHasuserID(id)
	return au
}

// SetNillableHasuserID sets the hasuser edge to User by id if the given value is not nil.
func (au *AmbulanceUpdate) SetNillableHasuserID(id *int) *AmbulanceUpdate {
	if id != nil {
		au = au.SetHasuserID(*id)
	}
	return au
}

// SetHasuser sets the hasuser edge to User.
func (au *AmbulanceUpdate) SetHasuser(u *User) *AmbulanceUpdate {
	return au.SetHasuserID(u.ID)
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (au *AmbulanceUpdate) AddCarinspectionIDs(ids ...int) *AmbulanceUpdate {
	au.mutation.AddCarinspectionIDs(ids...)
	return au
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (au *AmbulanceUpdate) AddCarinspections(c ...*CarInspection) *AmbulanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCarinspectionIDs(ids...)
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (au *AmbulanceUpdate) AddCarcheckinoutIDs(ids ...int) *AmbulanceUpdate {
	au.mutation.AddCarcheckinoutIDs(ids...)
	return au
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (au *AmbulanceUpdate) AddCarcheckinout(c ...*CarCheckInOut) *AmbulanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCarcheckinoutIDs(ids...)
}

// Mutation returns the AmbulanceMutation object of the builder.
func (au *AmbulanceUpdate) Mutation() *AmbulanceMutation {
	return au.mutation
}

// ClearHasbrand clears the hasbrand edge to Carbrand.
func (au *AmbulanceUpdate) ClearHasbrand() *AmbulanceUpdate {
	au.mutation.ClearHasbrand()
	return au
}

// ClearHasinsurance clears the hasinsurance edge to Insurance.
func (au *AmbulanceUpdate) ClearHasinsurance() *AmbulanceUpdate {
	au.mutation.ClearHasinsurance()
	return au
}

// ClearHasstatus clears the hasstatus edge to InspectionResult.
func (au *AmbulanceUpdate) ClearHasstatus() *AmbulanceUpdate {
	au.mutation.ClearHasstatus()
	return au
}

// ClearHasuser clears the hasuser edge to User.
func (au *AmbulanceUpdate) ClearHasuser() *AmbulanceUpdate {
	au.mutation.ClearHasuser()
	return au
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (au *AmbulanceUpdate) RemoveCarinspectionIDs(ids ...int) *AmbulanceUpdate {
	au.mutation.RemoveCarinspectionIDs(ids...)
	return au
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (au *AmbulanceUpdate) RemoveCarinspections(c ...*CarInspection) *AmbulanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCarinspectionIDs(ids...)
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (au *AmbulanceUpdate) RemoveCarcheckinoutIDs(ids ...int) *AmbulanceUpdate {
	au.mutation.RemoveCarcheckinoutIDs(ids...)
	return au
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (au *AmbulanceUpdate) RemoveCarcheckinout(c ...*CarCheckInOut) *AmbulanceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCarcheckinoutIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AmbulanceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmbulanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AmbulanceUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AmbulanceUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AmbulanceUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AmbulanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ambulance.Table,
			Columns: ambulance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ambulance.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Carregistration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ambulance.FieldCarregistration,
		})
	}
	if value, ok := au.mutation.Registerat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ambulance.FieldRegisterat,
		})
	}
	if au.mutation.HasbrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasbrandTable,
			Columns: []string{ambulance.HasbrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carbrand.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.HasbrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasbrandTable,
			Columns: []string{ambulance.HasbrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carbrand.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.HasinsuranceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasinsuranceTable,
			Columns: []string{ambulance.HasinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.HasinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasinsuranceTable,
			Columns: []string{ambulance.HasinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.HasstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasstatusTable,
			Columns: []string{ambulance.HasstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.HasstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasstatusTable,
			Columns: []string{ambulance.HasstatusColumn},
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
	if au.mutation.HasuserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasuserTable,
			Columns: []string{ambulance.HasuserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.HasuserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasuserTable,
			Columns: []string{ambulance.HasuserColumn},
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
	if nodes := au.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarinspectionsTable,
			Columns: []string{ambulance.CarinspectionsColumn},
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
	if nodes := au.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarinspectionsTable,
			Columns: []string{ambulance.CarinspectionsColumn},
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
	if nodes := au.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarcheckinoutTable,
			Columns: []string{ambulance.CarcheckinoutColumn},
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
	if nodes := au.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarcheckinoutTable,
			Columns: []string{ambulance.CarcheckinoutColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ambulance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AmbulanceUpdateOne is the builder for updating a single Ambulance entity.
type AmbulanceUpdateOne struct {
	config
	hooks    []Hook
	mutation *AmbulanceMutation
}

// SetCarregistration sets the carregistration field.
func (auo *AmbulanceUpdateOne) SetCarregistration(s string) *AmbulanceUpdateOne {
	auo.mutation.SetCarregistration(s)
	return auo
}

// SetRegisterat sets the registerat field.
func (auo *AmbulanceUpdateOne) SetRegisterat(t time.Time) *AmbulanceUpdateOne {
	auo.mutation.SetRegisterat(t)
	return auo
}

// SetNillableRegisterat sets the registerat field if the given value is not nil.
func (auo *AmbulanceUpdateOne) SetNillableRegisterat(t *time.Time) *AmbulanceUpdateOne {
	if t != nil {
		auo.SetRegisterat(*t)
	}
	return auo
}

// SetHasbrandID sets the hasbrand edge to Carbrand by id.
func (auo *AmbulanceUpdateOne) SetHasbrandID(id int) *AmbulanceUpdateOne {
	auo.mutation.SetHasbrandID(id)
	return auo
}

// SetNillableHasbrandID sets the hasbrand edge to Carbrand by id if the given value is not nil.
func (auo *AmbulanceUpdateOne) SetNillableHasbrandID(id *int) *AmbulanceUpdateOne {
	if id != nil {
		auo = auo.SetHasbrandID(*id)
	}
	return auo
}

// SetHasbrand sets the hasbrand edge to Carbrand.
func (auo *AmbulanceUpdateOne) SetHasbrand(c *Carbrand) *AmbulanceUpdateOne {
	return auo.SetHasbrandID(c.ID)
}

// SetHasinsuranceID sets the hasinsurance edge to Insurance by id.
func (auo *AmbulanceUpdateOne) SetHasinsuranceID(id int) *AmbulanceUpdateOne {
	auo.mutation.SetHasinsuranceID(id)
	return auo
}

// SetNillableHasinsuranceID sets the hasinsurance edge to Insurance by id if the given value is not nil.
func (auo *AmbulanceUpdateOne) SetNillableHasinsuranceID(id *int) *AmbulanceUpdateOne {
	if id != nil {
		auo = auo.SetHasinsuranceID(*id)
	}
	return auo
}

// SetHasinsurance sets the hasinsurance edge to Insurance.
func (auo *AmbulanceUpdateOne) SetHasinsurance(i *Insurance) *AmbulanceUpdateOne {
	return auo.SetHasinsuranceID(i.ID)
}

// SetHasstatusID sets the hasstatus edge to InspectionResult by id.
func (auo *AmbulanceUpdateOne) SetHasstatusID(id int) *AmbulanceUpdateOne {
	auo.mutation.SetHasstatusID(id)
	return auo
}

// SetNillableHasstatusID sets the hasstatus edge to InspectionResult by id if the given value is not nil.
func (auo *AmbulanceUpdateOne) SetNillableHasstatusID(id *int) *AmbulanceUpdateOne {
	if id != nil {
		auo = auo.SetHasstatusID(*id)
	}
	return auo
}

// SetHasstatus sets the hasstatus edge to InspectionResult.
func (auo *AmbulanceUpdateOne) SetHasstatus(i *InspectionResult) *AmbulanceUpdateOne {
	return auo.SetHasstatusID(i.ID)
}

// SetHasuserID sets the hasuser edge to User by id.
func (auo *AmbulanceUpdateOne) SetHasuserID(id int) *AmbulanceUpdateOne {
	auo.mutation.SetHasuserID(id)
	return auo
}

// SetNillableHasuserID sets the hasuser edge to User by id if the given value is not nil.
func (auo *AmbulanceUpdateOne) SetNillableHasuserID(id *int) *AmbulanceUpdateOne {
	if id != nil {
		auo = auo.SetHasuserID(*id)
	}
	return auo
}

// SetHasuser sets the hasuser edge to User.
func (auo *AmbulanceUpdateOne) SetHasuser(u *User) *AmbulanceUpdateOne {
	return auo.SetHasuserID(u.ID)
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (auo *AmbulanceUpdateOne) AddCarinspectionIDs(ids ...int) *AmbulanceUpdateOne {
	auo.mutation.AddCarinspectionIDs(ids...)
	return auo
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (auo *AmbulanceUpdateOne) AddCarinspections(c ...*CarInspection) *AmbulanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCarinspectionIDs(ids...)
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (auo *AmbulanceUpdateOne) AddCarcheckinoutIDs(ids ...int) *AmbulanceUpdateOne {
	auo.mutation.AddCarcheckinoutIDs(ids...)
	return auo
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (auo *AmbulanceUpdateOne) AddCarcheckinout(c ...*CarCheckInOut) *AmbulanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCarcheckinoutIDs(ids...)
}

// Mutation returns the AmbulanceMutation object of the builder.
func (auo *AmbulanceUpdateOne) Mutation() *AmbulanceMutation {
	return auo.mutation
}

// ClearHasbrand clears the hasbrand edge to Carbrand.
func (auo *AmbulanceUpdateOne) ClearHasbrand() *AmbulanceUpdateOne {
	auo.mutation.ClearHasbrand()
	return auo
}

// ClearHasinsurance clears the hasinsurance edge to Insurance.
func (auo *AmbulanceUpdateOne) ClearHasinsurance() *AmbulanceUpdateOne {
	auo.mutation.ClearHasinsurance()
	return auo
}

// ClearHasstatus clears the hasstatus edge to InspectionResult.
func (auo *AmbulanceUpdateOne) ClearHasstatus() *AmbulanceUpdateOne {
	auo.mutation.ClearHasstatus()
	return auo
}

// ClearHasuser clears the hasuser edge to User.
func (auo *AmbulanceUpdateOne) ClearHasuser() *AmbulanceUpdateOne {
	auo.mutation.ClearHasuser()
	return auo
}

// RemoveCarinspectionIDs removes the carinspections edge to CarInspection by ids.
func (auo *AmbulanceUpdateOne) RemoveCarinspectionIDs(ids ...int) *AmbulanceUpdateOne {
	auo.mutation.RemoveCarinspectionIDs(ids...)
	return auo
}

// RemoveCarinspections removes carinspections edges to CarInspection.
func (auo *AmbulanceUpdateOne) RemoveCarinspections(c ...*CarInspection) *AmbulanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCarinspectionIDs(ids...)
}

// RemoveCarcheckinoutIDs removes the carcheckinout edge to CarCheckInOut by ids.
func (auo *AmbulanceUpdateOne) RemoveCarcheckinoutIDs(ids ...int) *AmbulanceUpdateOne {
	auo.mutation.RemoveCarcheckinoutIDs(ids...)
	return auo
}

// RemoveCarcheckinout removes carcheckinout edges to CarCheckInOut.
func (auo *AmbulanceUpdateOne) RemoveCarcheckinout(c ...*CarCheckInOut) *AmbulanceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCarcheckinoutIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AmbulanceUpdateOne) Save(ctx context.Context) (*Ambulance, error) {

	var (
		err  error
		node *Ambulance
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmbulanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AmbulanceUpdateOne) SaveX(ctx context.Context) *Ambulance {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AmbulanceUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AmbulanceUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AmbulanceUpdateOne) sqlSave(ctx context.Context) (a *Ambulance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ambulance.Table,
			Columns: ambulance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ambulance.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Ambulance.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Carregistration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ambulance.FieldCarregistration,
		})
	}
	if value, ok := auo.mutation.Registerat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ambulance.FieldRegisterat,
		})
	}
	if auo.mutation.HasbrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasbrandTable,
			Columns: []string{ambulance.HasbrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carbrand.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.HasbrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasbrandTable,
			Columns: []string{ambulance.HasbrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carbrand.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.HasinsuranceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasinsuranceTable,
			Columns: []string{ambulance.HasinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.HasinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasinsuranceTable,
			Columns: []string{ambulance.HasinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.HasstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasstatusTable,
			Columns: []string{ambulance.HasstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inspectionresult.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.HasstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasstatusTable,
			Columns: []string{ambulance.HasstatusColumn},
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
	if auo.mutation.HasuserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasuserTable,
			Columns: []string{ambulance.HasuserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.HasuserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ambulance.HasuserTable,
			Columns: []string{ambulance.HasuserColumn},
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
	if nodes := auo.mutation.RemovedCarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarinspectionsTable,
			Columns: []string{ambulance.CarinspectionsColumn},
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
	if nodes := auo.mutation.CarinspectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarinspectionsTable,
			Columns: []string{ambulance.CarinspectionsColumn},
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
	if nodes := auo.mutation.RemovedCarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarcheckinoutTable,
			Columns: []string{ambulance.CarcheckinoutColumn},
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
	if nodes := auo.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.CarcheckinoutTable,
			Columns: []string{ambulance.CarcheckinoutColumn},
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
	a = &Ambulance{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ambulance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
