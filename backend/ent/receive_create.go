// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/receive"
	"github.com/team07/app/ent/transport"
)

// ReceiveCreate is the builder for creating a Receive entity.
type ReceiveCreate struct {
	config
	mutation *ReceiveMutation
	hooks    []Hook
}

// SetSendname sets the sendname field.
func (rc *ReceiveCreate) SetSendname(s string) *ReceiveCreate {
	rc.mutation.SetSendname(s)
	return rc
}

// AddReceiveidIDs adds the receiveid edge to Transport by ids.
func (rc *ReceiveCreate) AddReceiveidIDs(ids ...int) *ReceiveCreate {
	rc.mutation.AddReceiveidIDs(ids...)
	return rc
}

// AddReceiveid adds the receiveid edges to Transport.
func (rc *ReceiveCreate) AddReceiveid(t ...*Transport) *ReceiveCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddReceiveidIDs(ids...)
}

// Mutation returns the ReceiveMutation object of the builder.
func (rc *ReceiveCreate) Mutation() *ReceiveMutation {
	return rc.mutation
}

// Save creates the Receive in the database.
func (rc *ReceiveCreate) Save(ctx context.Context) (*Receive, error) {
	if _, ok := rc.mutation.Sendname(); !ok {
		return nil, &ValidationError{Name: "sendname", err: errors.New("ent: missing required field \"sendname\"")}
	}
	var (
		err  error
		node *Receive
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReceiveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReceiveCreate) SaveX(ctx context.Context) *Receive {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *ReceiveCreate) sqlSave(ctx context.Context) (*Receive, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *ReceiveCreate) createSpec() (*Receive, *sqlgraph.CreateSpec) {
	var (
		r     = &Receive{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: receive.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: receive.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Sendname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: receive.FieldSendname,
		})
		r.Sendname = value
	}
	if nodes := rc.mutation.ReceiveidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   receive.ReceiveidTable,
			Columns: []string{receive.ReceiveidColumn},
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
	return r, _spec
}
