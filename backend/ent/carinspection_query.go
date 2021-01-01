// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/user"
)

// CarInspectionQuery is the builder for querying CarInspection entities.
type CarInspectionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.CarInspection
	// eager-loading edges.
	withUser             *UserQuery
	withAmbulance        *AmbulanceQuery
	withInspectionresult *InspectionResultQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ciq *CarInspectionQuery) Where(ps ...predicate.CarInspection) *CarInspectionQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CarInspectionQuery) Limit(limit int) *CarInspectionQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CarInspectionQuery) Offset(offset int) *CarInspectionQuery {
	ciq.offset = &offset
	return ciq
}

// Order adds an order step to the query.
func (ciq *CarInspectionQuery) Order(o ...OrderFunc) *CarInspectionQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryUser chains the current query on the user edge.
func (ciq *CarInspectionQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carinspection.Table, carinspection.FieldID, ciq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carinspection.UserTable, carinspection.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAmbulance chains the current query on the ambulance edge.
func (ciq *CarInspectionQuery) QueryAmbulance() *AmbulanceQuery {
	query := &AmbulanceQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carinspection.Table, carinspection.FieldID, ciq.sqlQuery()),
			sqlgraph.To(ambulance.Table, ambulance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carinspection.AmbulanceTable, carinspection.AmbulanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInspectionresult chains the current query on the inspectionresult edge.
func (ciq *CarInspectionQuery) QueryInspectionresult() *InspectionResultQuery {
	query := &InspectionResultQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carinspection.Table, carinspection.FieldID, ciq.sqlQuery()),
			sqlgraph.To(inspectionresult.Table, inspectionresult.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carinspection.InspectionresultTable, carinspection.InspectionresultColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CarInspection entity in the query. Returns *NotFoundError when no carinspection was found.
func (ciq *CarInspectionQuery) First(ctx context.Context) (*CarInspection, error) {
	cis, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cis) == 0 {
		return nil, &NotFoundError{carinspection.Label}
	}
	return cis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CarInspectionQuery) FirstX(ctx context.Context) *CarInspection {
	ci, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ci
}

// FirstID returns the first CarInspection id in the query. Returns *NotFoundError when no id was found.
func (ciq *CarInspectionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{carinspection.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ciq *CarInspectionQuery) FirstXID(ctx context.Context) int {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CarInspection entity in the query, returns an error if not exactly one entity was returned.
func (ciq *CarInspectionQuery) Only(ctx context.Context) (*CarInspection, error) {
	cis, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cis) {
	case 1:
		return cis[0], nil
	case 0:
		return nil, &NotFoundError{carinspection.Label}
	default:
		return nil, &NotSingularError{carinspection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CarInspectionQuery) OnlyX(ctx context.Context) *CarInspection {
	ci, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ci
}

// OnlyID returns the only CarInspection id in the query, returns an error if not exactly one id was returned.
func (ciq *CarInspectionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = &NotSingularError{carinspection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CarInspectionQuery) OnlyIDX(ctx context.Context) int {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CarInspections.
func (ciq *CarInspectionQuery) All(ctx context.Context) ([]*CarInspection, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CarInspectionQuery) AllX(ctx context.Context) []*CarInspection {
	cis, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cis
}

// IDs executes the query and returns a list of CarInspection ids.
func (ciq *CarInspectionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ciq.Select(carinspection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CarInspectionQuery) IDsX(ctx context.Context) []int {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CarInspectionQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CarInspectionQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CarInspectionQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CarInspectionQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CarInspectionQuery) Clone() *CarInspectionQuery {
	return &CarInspectionQuery{
		config:     ciq.config,
		limit:      ciq.limit,
		offset:     ciq.offset,
		order:      append([]OrderFunc{}, ciq.order...),
		unique:     append([]string{}, ciq.unique...),
		predicates: append([]predicate.CarInspection{}, ciq.predicates...),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CarInspectionQuery) WithUser(opts ...func(*UserQuery)) *CarInspectionQuery {
	query := &UserQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withUser = query
	return ciq
}

//  WithAmbulance tells the query-builder to eager-loads the nodes that are connected to
// the "ambulance" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CarInspectionQuery) WithAmbulance(opts ...func(*AmbulanceQuery)) *CarInspectionQuery {
	query := &AmbulanceQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withAmbulance = query
	return ciq
}

//  WithInspectionresult tells the query-builder to eager-loads the nodes that are connected to
// the "inspectionresult" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CarInspectionQuery) WithInspectionresult(opts ...func(*InspectionResultQuery)) *CarInspectionQuery {
	query := &InspectionResultQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withInspectionresult = query
	return ciq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Datetime time.Time `json:"datetime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CarInspection.Query().
//		GroupBy(carinspection.FieldDatetime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ciq *CarInspectionQuery) GroupBy(field string, fields ...string) *CarInspectionGroupBy {
	group := &CarInspectionGroupBy{config: ciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Datetime time.Time `json:"datetime,omitempty"`
//	}
//
//	client.CarInspection.Query().
//		Select(carinspection.FieldDatetime).
//		Scan(ctx, &v)
//
func (ciq *CarInspectionQuery) Select(field string, fields ...string) *CarInspectionSelect {
	selector := &CarInspectionSelect{config: ciq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(), nil
	}
	return selector
}

func (ciq *CarInspectionQuery) prepareQuery(ctx context.Context) error {
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CarInspectionQuery) sqlAll(ctx context.Context) ([]*CarInspection, error) {
	var (
		nodes       = []*CarInspection{}
		withFKs     = ciq.withFKs
		_spec       = ciq.querySpec()
		loadedTypes = [3]bool{
			ciq.withUser != nil,
			ciq.withAmbulance != nil,
			ciq.withInspectionresult != nil,
		}
	)
	if ciq.withUser != nil || ciq.withAmbulance != nil || ciq.withInspectionresult != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, carinspection.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &CarInspection{config: ciq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ciq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarInspection)
		for i := range nodes {
			if fk := nodes[i].user_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := ciq.withAmbulance; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarInspection)
		for i := range nodes {
			if fk := nodes[i].ambulance_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(ambulance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ambulance_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Ambulance = n
			}
		}
	}

	if query := ciq.withInspectionresult; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarInspection)
		for i := range nodes {
			if fk := nodes[i].inspectionresult_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(inspectionresult.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "inspectionresult_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Inspectionresult = n
			}
		}
	}

	return nodes, nil
}

func (ciq *CarInspectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CarInspectionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ciq *CarInspectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   carinspection.Table,
			Columns: carinspection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carinspection.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CarInspectionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(carinspection.Table)
	selector := builder.Select(t1.Columns(carinspection.Columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(carinspection.Columns...)...)
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CarInspectionGroupBy is the builder for group-by CarInspection entities.
type CarInspectionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CarInspectionGroupBy) Aggregate(fns ...AggregateFunc) *CarInspectionGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scan the result into the given value.
func (cigb *CarInspectionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CarInspectionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) StringsX(ctx context.Context) []string {
	v, err := cigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) StringX(ctx context.Context) string {
	v, err := cigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CarInspectionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) IntsX(ctx context.Context) []int {
	v, err := cigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) IntX(ctx context.Context) int {
	v, err := cigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CarInspectionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CarInspectionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (cigb *CarInspectionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cigb *CarInspectionGroupBy) BoolX(ctx context.Context) bool {
	v, err := cigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cigb *CarInspectionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cigb.sqlQuery().Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CarInspectionGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql
	columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
	columns = append(columns, cigb.fields...)
	for _, fn := range cigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cigb.fields...)
}

// CarInspectionSelect is the builder for select fields of CarInspection entities.
type CarInspectionSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (cis *CarInspectionSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := cis.path(ctx)
	if err != nil {
		return err
	}
	cis.sql = query
	return cis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cis *CarInspectionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CarInspectionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cis *CarInspectionSelect) StringsX(ctx context.Context) []string {
	v, err := cis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cis *CarInspectionSelect) StringX(ctx context.Context) string {
	v, err := cis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CarInspectionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cis *CarInspectionSelect) IntsX(ctx context.Context) []int {
	v, err := cis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cis *CarInspectionSelect) IntX(ctx context.Context) int {
	v, err := cis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CarInspectionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cis *CarInspectionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cis *CarInspectionSelect) Float64X(ctx context.Context) float64 {
	v, err := cis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CarInspectionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cis *CarInspectionSelect) BoolsX(ctx context.Context) []bool {
	v, err := cis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (cis *CarInspectionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carinspection.Label}
	default:
		err = fmt.Errorf("ent: CarInspectionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cis *CarInspectionSelect) BoolX(ctx context.Context) bool {
	v, err := cis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cis *CarInspectionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cis.sqlQuery().Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cis *CarInspectionSelect) sqlQuery() sql.Querier {
	selector := cis.sql
	selector.Select(selector.Columns(cis.fields...)...)
	return selector
}
