// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/repairing"
	"github.com/team07/app/ent/user"
)

// CarRepairrecordCreate is the builder for creating a CarRepairrecord entity.
type CarRepairrecordCreate struct {
	config
	mutation *CarRepairrecordMutation
	hooks    []Hook
}

// SetDatetime sets the datetime field.
func (crc *CarRepairrecordCreate) SetDatetime(t time.Time) *CarRepairrecordCreate {
	crc.mutation.SetDatetime(t)
	return crc
}

// SetKeeperID sets the keeper edge to Repairing by id.
func (crc *CarRepairrecordCreate) SetKeeperID(id int) *CarRepairrecordCreate {
	crc.mutation.SetKeeperID(id)
	return crc
}

// SetNillableKeeperID sets the keeper edge to Repairing by id if the given value is not nil.
func (crc *CarRepairrecordCreate) SetNillableKeeperID(id *int) *CarRepairrecordCreate {
	if id != nil {
		crc = crc.SetKeeperID(*id)
	}
	return crc
}

// SetKeeper sets the keeper edge to Repairing.
func (crc *CarRepairrecordCreate) SetKeeper(r *Repairing) *CarRepairrecordCreate {
	return crc.SetKeeperID(r.ID)
}

// SetUserID sets the user edge to User by id.
func (crc *CarRepairrecordCreate) SetUserID(id int) *CarRepairrecordCreate {
	crc.mutation.SetUserID(id)
	return crc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (crc *CarRepairrecordCreate) SetNillableUserID(id *int) *CarRepairrecordCreate {
	if id != nil {
		crc = crc.SetUserID(*id)
	}
	return crc
}

// SetUser sets the user edge to User.
func (crc *CarRepairrecordCreate) SetUser(u *User) *CarRepairrecordCreate {
	return crc.SetUserID(u.ID)
}

// SetCarinspectionID sets the carinspection edge to CarInspection by id.
func (crc *CarRepairrecordCreate) SetCarinspectionID(id int) *CarRepairrecordCreate {
	crc.mutation.SetCarinspectionID(id)
	return crc
}

// SetNillableCarinspectionID sets the carinspection edge to CarInspection by id if the given value is not nil.
func (crc *CarRepairrecordCreate) SetNillableCarinspectionID(id *int) *CarRepairrecordCreate {
	if id != nil {
		crc = crc.SetCarinspectionID(*id)
	}
	return crc
}

// SetCarinspection sets the carinspection edge to CarInspection.
func (crc *CarRepairrecordCreate) SetCarinspection(c *CarInspection) *CarRepairrecordCreate {
	return crc.SetCarinspectionID(c.ID)
}

// Mutation returns the CarRepairrecordMutation object of the builder.
func (crc *CarRepairrecordCreate) Mutation() *CarRepairrecordMutation {
	return crc.mutation
}

// Save creates the CarRepairrecord in the database.
func (crc *CarRepairrecordCreate) Save(ctx context.Context) (*CarRepairrecord, error) {
	if _, ok := crc.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	var (
		err  error
		node *CarRepairrecord
	)
	if len(crc.hooks) == 0 {
		node, err = crc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarRepairrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			crc.mutation = mutation
			node, err = crc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(crc.hooks) - 1; i >= 0; i-- {
			mut = crc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, crc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CarRepairrecordCreate) SaveX(ctx context.Context) *CarRepairrecord {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crc *CarRepairrecordCreate) sqlSave(ctx context.Context) (*CarRepairrecord, error) {
	cr, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	cr.ID = int(id)
	return cr, nil
}

func (crc *CarRepairrecordCreate) createSpec() (*CarRepairrecord, *sqlgraph.CreateSpec) {
	var (
		cr    = &CarRepairrecord{config: crc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: carrepairrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carrepairrecord.FieldID,
			},
		}
	)
	if value, ok := crc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: carrepairrecord.FieldDatetime,
		})
		cr.Datetime = value
	}
	if nodes := crc.mutation.KeeperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carrepairrecord.KeeperTable,
			Columns: []string{carrepairrecord.KeeperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairing.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := crc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carrepairrecord.UserTable,
			Columns: []string{carrepairrecord.UserColumn},
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
	if nodes := crc.mutation.CarinspectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carrepairrecord.CarinspectionTable,
			Columns: []string{carrepairrecord.CarinspectionColumn},
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
	return cr, _spec
}
