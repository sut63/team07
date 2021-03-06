// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carbrand"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/insurance"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/user"
)

// AmbulanceCreate is the builder for creating a Ambulance entity.
type AmbulanceCreate struct {
	config
	mutation *AmbulanceMutation
	hooks    []Hook
}

// SetCarregistration sets the carregistration field.
func (ac *AmbulanceCreate) SetCarregistration(s string) *AmbulanceCreate {
	ac.mutation.SetCarregistration(s)
	return ac
}

// SetEnginepower sets the enginepower field.
func (ac *AmbulanceCreate) SetEnginepower(i int) *AmbulanceCreate {
	ac.mutation.SetEnginepower(i)
	return ac
}

// SetDisplacement sets the displacement field.
func (ac *AmbulanceCreate) SetDisplacement(i int) *AmbulanceCreate {
	ac.mutation.SetDisplacement(i)
	return ac
}

// SetRegisterat sets the registerat field.
func (ac *AmbulanceCreate) SetRegisterat(t time.Time) *AmbulanceCreate {
	ac.mutation.SetRegisterat(t)
	return ac
}

// SetHasbrandID sets the hasbrand edge to Carbrand by id.
func (ac *AmbulanceCreate) SetHasbrandID(id int) *AmbulanceCreate {
	ac.mutation.SetHasbrandID(id)
	return ac
}

// SetNillableHasbrandID sets the hasbrand edge to Carbrand by id if the given value is not nil.
func (ac *AmbulanceCreate) SetNillableHasbrandID(id *int) *AmbulanceCreate {
	if id != nil {
		ac = ac.SetHasbrandID(*id)
	}
	return ac
}

// SetHasbrand sets the hasbrand edge to Carbrand.
func (ac *AmbulanceCreate) SetHasbrand(c *Carbrand) *AmbulanceCreate {
	return ac.SetHasbrandID(c.ID)
}

// SetHasinsuranceID sets the hasinsurance edge to Insurance by id.
func (ac *AmbulanceCreate) SetHasinsuranceID(id int) *AmbulanceCreate {
	ac.mutation.SetHasinsuranceID(id)
	return ac
}

// SetNillableHasinsuranceID sets the hasinsurance edge to Insurance by id if the given value is not nil.
func (ac *AmbulanceCreate) SetNillableHasinsuranceID(id *int) *AmbulanceCreate {
	if id != nil {
		ac = ac.SetHasinsuranceID(*id)
	}
	return ac
}

// SetHasinsurance sets the hasinsurance edge to Insurance.
func (ac *AmbulanceCreate) SetHasinsurance(i *Insurance) *AmbulanceCreate {
	return ac.SetHasinsuranceID(i.ID)
}

// SetHasstatusID sets the hasstatus edge to InspectionResult by id.
func (ac *AmbulanceCreate) SetHasstatusID(id int) *AmbulanceCreate {
	ac.mutation.SetHasstatusID(id)
	return ac
}

// SetNillableHasstatusID sets the hasstatus edge to InspectionResult by id if the given value is not nil.
func (ac *AmbulanceCreate) SetNillableHasstatusID(id *int) *AmbulanceCreate {
	if id != nil {
		ac = ac.SetHasstatusID(*id)
	}
	return ac
}

// SetHasstatus sets the hasstatus edge to InspectionResult.
func (ac *AmbulanceCreate) SetHasstatus(i *InspectionResult) *AmbulanceCreate {
	return ac.SetHasstatusID(i.ID)
}

// SetHasuserID sets the hasuser edge to User by id.
func (ac *AmbulanceCreate) SetHasuserID(id int) *AmbulanceCreate {
	ac.mutation.SetHasuserID(id)
	return ac
}

// SetNillableHasuserID sets the hasuser edge to User by id if the given value is not nil.
func (ac *AmbulanceCreate) SetNillableHasuserID(id *int) *AmbulanceCreate {
	if id != nil {
		ac = ac.SetHasuserID(*id)
	}
	return ac
}

// SetHasuser sets the hasuser edge to User.
func (ac *AmbulanceCreate) SetHasuser(u *User) *AmbulanceCreate {
	return ac.SetHasuserID(u.ID)
}

// AddCarinspectionIDs adds the carinspections edge to CarInspection by ids.
func (ac *AmbulanceCreate) AddCarinspectionIDs(ids ...int) *AmbulanceCreate {
	ac.mutation.AddCarinspectionIDs(ids...)
	return ac
}

// AddCarinspections adds the carinspections edges to CarInspection.
func (ac *AmbulanceCreate) AddCarinspections(c ...*CarInspection) *AmbulanceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddCarinspectionIDs(ids...)
}

// AddCarcheckinoutIDs adds the carcheckinout edge to CarCheckInOut by ids.
func (ac *AmbulanceCreate) AddCarcheckinoutIDs(ids ...int) *AmbulanceCreate {
	ac.mutation.AddCarcheckinoutIDs(ids...)
	return ac
}

// AddCarcheckinout adds the carcheckinout edges to CarCheckInOut.
func (ac *AmbulanceCreate) AddCarcheckinout(c ...*CarCheckInOut) *AmbulanceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddCarcheckinoutIDs(ids...)
}

// AddAmbulanceIDs adds the ambulance edge to Transport by ids.
func (ac *AmbulanceCreate) AddAmbulanceIDs(ids ...int) *AmbulanceCreate {
	ac.mutation.AddAmbulanceIDs(ids...)
	return ac
}

// AddAmbulance adds the ambulance edges to Transport.
func (ac *AmbulanceCreate) AddAmbulance(t ...*Transport) *AmbulanceCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddAmbulanceIDs(ids...)
}

// Mutation returns the AmbulanceMutation object of the builder.
func (ac *AmbulanceCreate) Mutation() *AmbulanceMutation {
	return ac.mutation
}

// Save creates the Ambulance in the database.
func (ac *AmbulanceCreate) Save(ctx context.Context) (*Ambulance, error) {
	if _, ok := ac.mutation.Carregistration(); !ok {
		return nil, &ValidationError{Name: "carregistration", err: errors.New("ent: missing required field \"carregistration\"")}
	}
	if v, ok := ac.mutation.Carregistration(); ok {
		if err := ambulance.CarregistrationValidator(v); err != nil {
			return nil, &ValidationError{Name: "carregistration", err: fmt.Errorf("ent: validator failed for field \"carregistration\": %w", err)}
		}
	}
	if _, ok := ac.mutation.Enginepower(); !ok {
		return nil, &ValidationError{Name: "enginepower", err: errors.New("ent: missing required field \"enginepower\"")}
	}
	if v, ok := ac.mutation.Enginepower(); ok {
		if err := ambulance.EnginepowerValidator(v); err != nil {
			return nil, &ValidationError{Name: "enginepower", err: fmt.Errorf("ent: validator failed for field \"enginepower\": %w", err)}
		}
	}
	if _, ok := ac.mutation.Displacement(); !ok {
		return nil, &ValidationError{Name: "displacement", err: errors.New("ent: missing required field \"displacement\"")}
	}
	if v, ok := ac.mutation.Displacement(); ok {
		if err := ambulance.DisplacementValidator(v); err != nil {
			return nil, &ValidationError{Name: "displacement", err: fmt.Errorf("ent: validator failed for field \"displacement\": %w", err)}
		}
	}
	if _, ok := ac.mutation.Registerat(); !ok {
		return nil, &ValidationError{Name: "registerat", err: errors.New("ent: missing required field \"registerat\"")}
	}
	var (
		err  error
		node *Ambulance
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmbulanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AmbulanceCreate) SaveX(ctx context.Context) *Ambulance {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AmbulanceCreate) sqlSave(ctx context.Context) (*Ambulance, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AmbulanceCreate) createSpec() (*Ambulance, *sqlgraph.CreateSpec) {
	var (
		a     = &Ambulance{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ambulance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ambulance.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Carregistration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ambulance.FieldCarregistration,
		})
		a.Carregistration = value
	}
	if value, ok := ac.mutation.Enginepower(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ambulance.FieldEnginepower,
		})
		a.Enginepower = value
	}
	if value, ok := ac.mutation.Displacement(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ambulance.FieldDisplacement,
		})
		a.Displacement = value
	}
	if value, ok := ac.mutation.Registerat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ambulance.FieldRegisterat,
		})
		a.Registerat = value
	}
	if nodes := ac.mutation.HasbrandIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.HasinsuranceIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.HasstatusIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.HasuserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CarinspectionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CarcheckinoutIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AmbulanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ambulance.AmbulanceTable,
			Columns: []string{ambulance.AmbulanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
