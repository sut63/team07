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
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/user"
)

// CarInspectionCreate is the builder for creating a CarInspection entity.
type CarInspectionCreate struct {
	config
	mutation *CarInspectionMutation
	hooks    []Hook
}

// SetWheelCenter sets the wheel_center field.
func (cic *CarInspectionCreate) SetWheelCenter(f float64) *CarInspectionCreate {
	cic.mutation.SetWheelCenter(f)
	return cic
}

// SetSoundLevel sets the sound_level field.
func (cic *CarInspectionCreate) SetSoundLevel(f float64) *CarInspectionCreate {
	cic.mutation.SetSoundLevel(f)
	return cic
}

// SetBlacksmoke sets the blacksmoke field.
func (cic *CarInspectionCreate) SetBlacksmoke(f float64) *CarInspectionCreate {
	cic.mutation.SetBlacksmoke(f)
	return cic
}

// SetDatetime sets the datetime field.
func (cic *CarInspectionCreate) SetDatetime(t time.Time) *CarInspectionCreate {
	cic.mutation.SetDatetime(t)
	return cic
}

// SetNote sets the note field.
func (cic *CarInspectionCreate) SetNote(s string) *CarInspectionCreate {
	cic.mutation.SetNote(s)
	return cic
}

// SetUserID sets the user edge to User by id.
func (cic *CarInspectionCreate) SetUserID(id int) *CarInspectionCreate {
	cic.mutation.SetUserID(id)
	return cic
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (cic *CarInspectionCreate) SetNillableUserID(id *int) *CarInspectionCreate {
	if id != nil {
		cic = cic.SetUserID(*id)
	}
	return cic
}

// SetUser sets the user edge to User.
func (cic *CarInspectionCreate) SetUser(u *User) *CarInspectionCreate {
	return cic.SetUserID(u.ID)
}

// SetAmbulanceID sets the ambulance edge to Ambulance by id.
func (cic *CarInspectionCreate) SetAmbulanceID(id int) *CarInspectionCreate {
	cic.mutation.SetAmbulanceID(id)
	return cic
}

// SetNillableAmbulanceID sets the ambulance edge to Ambulance by id if the given value is not nil.
func (cic *CarInspectionCreate) SetNillableAmbulanceID(id *int) *CarInspectionCreate {
	if id != nil {
		cic = cic.SetAmbulanceID(*id)
	}
	return cic
}

// SetAmbulance sets the ambulance edge to Ambulance.
func (cic *CarInspectionCreate) SetAmbulance(a *Ambulance) *CarInspectionCreate {
	return cic.SetAmbulanceID(a.ID)
}

// SetInspectionresultID sets the inspectionresult edge to InspectionResult by id.
func (cic *CarInspectionCreate) SetInspectionresultID(id int) *CarInspectionCreate {
	cic.mutation.SetInspectionresultID(id)
	return cic
}

// SetNillableInspectionresultID sets the inspectionresult edge to InspectionResult by id if the given value is not nil.
func (cic *CarInspectionCreate) SetNillableInspectionresultID(id *int) *CarInspectionCreate {
	if id != nil {
		cic = cic.SetInspectionresultID(*id)
	}
	return cic
}

// SetInspectionresult sets the inspectionresult edge to InspectionResult.
func (cic *CarInspectionCreate) SetInspectionresult(i *InspectionResult) *CarInspectionCreate {
	return cic.SetInspectionresultID(i.ID)
}

// AddCarrepairrecordIDs adds the carrepairrecords edge to CarRepairrecord by ids.
func (cic *CarInspectionCreate) AddCarrepairrecordIDs(ids ...int) *CarInspectionCreate {
	cic.mutation.AddCarrepairrecordIDs(ids...)
	return cic
}

// AddCarrepairrecords adds the carrepairrecords edges to CarRepairrecord.
func (cic *CarInspectionCreate) AddCarrepairrecords(c ...*CarRepairrecord) *CarInspectionCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cic.AddCarrepairrecordIDs(ids...)
}

// Mutation returns the CarInspectionMutation object of the builder.
func (cic *CarInspectionCreate) Mutation() *CarInspectionMutation {
	return cic.mutation
}

// Save creates the CarInspection in the database.
func (cic *CarInspectionCreate) Save(ctx context.Context) (*CarInspection, error) {
	if _, ok := cic.mutation.WheelCenter(); !ok {
		return nil, &ValidationError{Name: "wheel_center", err: errors.New("ent: missing required field \"wheel_center\"")}
	}
	if v, ok := cic.mutation.WheelCenter(); ok {
		if err := carinspection.WheelCenterValidator(v); err != nil {
			return nil, &ValidationError{Name: "wheel_center", err: fmt.Errorf("ent: validator failed for field \"wheel_center\": %w", err)}
		}
	}
	if _, ok := cic.mutation.SoundLevel(); !ok {
		return nil, &ValidationError{Name: "sound_level", err: errors.New("ent: missing required field \"sound_level\"")}
	}
	if v, ok := cic.mutation.SoundLevel(); ok {
		if err := carinspection.SoundLevelValidator(v); err != nil {
			return nil, &ValidationError{Name: "sound_level", err: fmt.Errorf("ent: validator failed for field \"sound_level\": %w", err)}
		}
	}
	if _, ok := cic.mutation.Blacksmoke(); !ok {
		return nil, &ValidationError{Name: "blacksmoke", err: errors.New("ent: missing required field \"blacksmoke\"")}
	}
	if v, ok := cic.mutation.Blacksmoke(); ok {
		if err := carinspection.BlacksmokeValidator(v); err != nil {
			return nil, &ValidationError{Name: "blacksmoke", err: fmt.Errorf("ent: validator failed for field \"blacksmoke\": %w", err)}
		}
	}
	if _, ok := cic.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	if _, ok := cic.mutation.Note(); !ok {
		return nil, &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	var (
		err  error
		node *CarInspection
	)
	if len(cic.hooks) == 0 {
		node, err = cic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarInspectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cic.mutation = mutation
			node, err = cic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cic.hooks) - 1; i >= 0; i-- {
			mut = cic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CarInspectionCreate) SaveX(ctx context.Context) *CarInspection {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cic *CarInspectionCreate) sqlSave(ctx context.Context) (*CarInspection, error) {
	ci, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ci.ID = int(id)
	return ci, nil
}

func (cic *CarInspectionCreate) createSpec() (*CarInspection, *sqlgraph.CreateSpec) {
	var (
		ci    = &CarInspection{config: cic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: carinspection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carinspection.FieldID,
			},
		}
	)
	if value, ok := cic.mutation.WheelCenter(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: carinspection.FieldWheelCenter,
		})
		ci.WheelCenter = value
	}
	if value, ok := cic.mutation.SoundLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: carinspection.FieldSoundLevel,
		})
		ci.SoundLevel = value
	}
	if value, ok := cic.mutation.Blacksmoke(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: carinspection.FieldBlacksmoke,
		})
		ci.Blacksmoke = value
	}
	if value, ok := cic.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: carinspection.FieldDatetime,
		})
		ci.Datetime = value
	}
	if value, ok := cic.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: carinspection.FieldNote,
		})
		ci.Note = value
	}
	if nodes := cic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carinspection.UserTable,
			Columns: []string{carinspection.UserColumn},
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
	if nodes := cic.mutation.AmbulanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carinspection.AmbulanceTable,
			Columns: []string{carinspection.AmbulanceColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.InspectionresultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carinspection.InspectionresultTable,
			Columns: []string{carinspection.InspectionresultColumn},
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
	if nodes := cic.mutation.CarrepairrecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   carinspection.CarrepairrecordsTable,
			Columns: []string{carinspection.CarrepairrecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrepairrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ci, _spec
}
