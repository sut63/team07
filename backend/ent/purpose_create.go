// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/purpose"
)

// PurposeCreate is the builder for creating a Purpose entity.
type PurposeCreate struct {
	config
	mutation *PurposeMutation
	hooks    []Hook
}

// Mutation returns the PurposeMutation object of the builder.
func (pc *PurposeCreate) Mutation() *PurposeMutation {
	return pc.mutation
}

// Save creates the Purpose in the database.
func (pc *PurposeCreate) Save(ctx context.Context) (*Purpose, error) {
	var (
		err  error
		node *Purpose
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurposeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PurposeCreate) SaveX(ctx context.Context) *Purpose {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PurposeCreate) sqlSave(ctx context.Context) (*Purpose, error) {
	pu, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pu.ID = int(id)
	return pu, nil
}

func (pc *PurposeCreate) createSpec() (*Purpose, *sqlgraph.CreateSpec) {
	var (
		pu    = &Purpose{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: purpose.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: purpose.FieldID,
			},
		}
	)
	return pu, _spec
}
