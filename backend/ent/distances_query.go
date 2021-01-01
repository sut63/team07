// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/distances"
	"github.com/team07/app/ent/predicate"
)

// DistancesQuery is the builder for querying Distances entities.
type DistancesQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Distances
	// eager-loading edges.
	withDisid *CarserviceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (dq *DistancesQuery) Where(ps ...predicate.Distances) *DistancesQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DistancesQuery) Limit(limit int) *DistancesQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DistancesQuery) Offset(offset int) *DistancesQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DistancesQuery) Order(o ...OrderFunc) *DistancesQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDisid chains the current query on the disid edge.
func (dq *DistancesQuery) QueryDisid() *CarserviceQuery {
	query := &CarserviceQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(distances.Table, distances.FieldID, dq.sqlQuery()),
			sqlgraph.To(carservice.Table, carservice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, distances.DisidTable, distances.DisidColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Distances entity in the query. Returns *NotFoundError when no distances was found.
func (dq *DistancesQuery) First(ctx context.Context) (*Distances, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{distances.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DistancesQuery) FirstX(ctx context.Context) *Distances {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Distances id in the query. Returns *NotFoundError when no id was found.
func (dq *DistancesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{distances.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DistancesQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Distances entity in the query, returns an error if not exactly one entity was returned.
func (dq *DistancesQuery) Only(ctx context.Context) (*Distances, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{distances.Label}
	default:
		return nil, &NotSingularError{distances.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DistancesQuery) OnlyX(ctx context.Context) *Distances {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Distances id in the query, returns an error if not exactly one id was returned.
func (dq *DistancesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = &NotSingularError{distances.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DistancesQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DistancesSlice.
func (dq *DistancesQuery) All(ctx context.Context) ([]*Distances, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DistancesQuery) AllX(ctx context.Context) []*Distances {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Distances ids.
func (dq *DistancesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(distances.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DistancesQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DistancesQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DistancesQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DistancesQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DistancesQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DistancesQuery) Clone() *DistancesQuery {
	return &DistancesQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Distances{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

//  WithDisid tells the query-builder to eager-loads the nodes that are connected to
// the "disid" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DistancesQuery) WithDisid(opts ...func(*CarserviceQuery)) *DistancesQuery {
	query := &CarserviceQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDisid = query
	return dq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Distances string `json:"Distances,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Distances.Query().
//		GroupBy(distances.FieldDistances).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DistancesQuery) GroupBy(field string, fields ...string) *DistancesGroupBy {
	group := &DistancesGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Distances string `json:"Distances,omitempty"`
//	}
//
//	client.Distances.Query().
//		Select(distances.FieldDistances).
//		Scan(ctx, &v)
//
func (dq *DistancesQuery) Select(field string, fields ...string) *DistancesSelect {
	selector := &DistancesSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return selector
}

func (dq *DistancesQuery) prepareQuery(ctx context.Context) error {
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DistancesQuery) sqlAll(ctx context.Context) ([]*Distances, error) {
	var (
		nodes       = []*Distances{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDisid != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Distances{config: dq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withDisid; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Distances)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Carservice(func(s *sql.Selector) {
			s.Where(sql.InValues(distances.DisidColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.distances_disid
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "distances_disid" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "distances_disid" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Disid = append(node.Edges.Disid, n)
		}
	}

	return nodes, nil
}

func (dq *DistancesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DistancesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DistancesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distances.Table,
			Columns: distances.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distances.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DistancesQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(distances.Table)
	selector := builder.Select(t1.Columns(distances.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(distances.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DistancesGroupBy is the builder for group-by Distances entities.
type DistancesGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DistancesGroupBy) Aggregate(fns ...AggregateFunc) *DistancesGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DistancesGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DistancesGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistancesGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DistancesGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DistancesGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistancesGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DistancesGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DistancesGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistancesGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DistancesGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DistancesGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistancesGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DistancesGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistancesGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DistancesGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DistancesGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DistancesGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DistancesSelect is the builder for select fields of Distances entities.
type DistancesSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DistancesSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ds.path(ctx)
	if err != nil {
		return err
	}
	ds.sql = query
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DistancesSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistancesSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DistancesSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DistancesSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistancesSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DistancesSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DistancesSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistancesSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DistancesSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DistancesSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistancesSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DistancesSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ds *DistancesSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distances.Label}
	default:
		err = fmt.Errorf("ent: DistancesSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DistancesSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DistancesSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DistancesSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
