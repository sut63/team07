// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team07/app/ent/playlist_video"
	"github.com/team07/app/ent/resolution"
)

// ResolutionCreate is the builder for creating a Resolution entity.
type ResolutionCreate struct {
	config
	mutation *ResolutionMutation
	hooks    []Hook
}

// SetValue sets the value field.
func (rc *ResolutionCreate) SetValue(i int) *ResolutionCreate {
	rc.mutation.SetValue(i)
	return rc
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (rc *ResolutionCreate) AddPlaylistVideoIDs(ids ...int) *ResolutionCreate {
	rc.mutation.AddPlaylistVideoIDs(ids...)
	return rc
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (rc *ResolutionCreate) AddPlaylistVideos(p ...*Playlist_Video) *ResolutionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the ResolutionMutation object of the builder.
func (rc *ResolutionCreate) Mutation() *ResolutionMutation {
	return rc.mutation
}

// Save creates the Resolution in the database.
func (rc *ResolutionCreate) Save(ctx context.Context) (*Resolution, error) {
	if err := rc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Resolution
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResolutionMutation)
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
func (rc *ResolutionCreate) SaveX(ctx context.Context) *Resolution {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *ResolutionCreate) preSave() error {
	if _, ok := rc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	if v, ok := rc.mutation.Value(); ok {
		if err := resolution.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf("ent: validator failed for field \"value\": %w", err)}
		}
	}
	return nil
}

func (rc *ResolutionCreate) sqlSave(ctx context.Context) (*Resolution, error) {
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

func (rc *ResolutionCreate) createSpec() (*Resolution, *sqlgraph.CreateSpec) {
	var (
		r     = &Resolution{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resolution.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resolution.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resolution.FieldValue,
		})
		r.Value = value
	}
	if nodes := rc.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resolution.PlaylistVideosTable,
			Columns: []string{resolution.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
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

// ResolutionCreateBulk is the builder for creating a bulk of Resolution entities.
type ResolutionCreateBulk struct {
	config
	builders []*ResolutionCreate
}

// Save creates the Resolution entities in the database.
func (rcb *ResolutionCreateBulk) Save(ctx context.Context) ([]*Resolution, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resolution, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ResolutionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rcb *ResolutionCreateBulk) SaveX(ctx context.Context) []*Resolution {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
