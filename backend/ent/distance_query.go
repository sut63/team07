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
	"github.com/team07/app/ent/distance"
	"github.com/team07/app/ent/predicate"
)

// DistanceQuery is the builder for querying Distance entities.
type DistanceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Distance
	// eager-loading edges.
	withDisid *CarserviceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (dq *DistanceQuery) Where(ps ...predicate.Distance) *DistanceQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DistanceQuery) Limit(limit int) *DistanceQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DistanceQuery) Offset(offset int) *DistanceQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DistanceQuery) Order(o ...OrderFunc) *DistanceQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDisid chains the current query on the disid edge.
func (dq *DistanceQuery) QueryDisid() *CarserviceQuery {
	query := &CarserviceQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(distance.Table, distance.FieldID, dq.sqlQuery()),
			sqlgraph.To(carservice.Table, carservice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, distance.DisidTable, distance.DisidColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Distance entity in the query. Returns *NotFoundError when no distance was found.
func (dq *DistanceQuery) First(ctx context.Context) (*Distance, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{distance.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DistanceQuery) FirstX(ctx context.Context) *Distance {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Distance id in the query. Returns *NotFoundError when no id was found.
func (dq *DistanceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{distance.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DistanceQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Distance entity in the query, returns an error if not exactly one entity was returned.
func (dq *DistanceQuery) Only(ctx context.Context) (*Distance, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{distance.Label}
	default:
		return nil, &NotSingularError{distance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DistanceQuery) OnlyX(ctx context.Context) *Distance {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Distance id in the query, returns an error if not exactly one id was returned.
func (dq *DistanceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = &NotSingularError{distance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DistanceQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Distances.
func (dq *DistanceQuery) All(ctx context.Context) ([]*Distance, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DistanceQuery) AllX(ctx context.Context) []*Distance {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Distance ids.
func (dq *DistanceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(distance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DistanceQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DistanceQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DistanceQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DistanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DistanceQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DistanceQuery) Clone() *DistanceQuery {
	return &DistanceQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Distance{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

//  WithDisid tells the query-builder to eager-loads the nodes that are connected to
// the "disid" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DistanceQuery) WithDisid(opts ...func(*CarserviceQuery)) *DistanceQuery {
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
//		Distance string `json:"Distance,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Distance.Query().
//		GroupBy(distance.FieldDistance).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DistanceQuery) GroupBy(field string, fields ...string) *DistanceGroupBy {
	group := &DistanceGroupBy{config: dq.config}
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
//		Distance string `json:"Distance,omitempty"`
//	}
//
//	client.Distance.Query().
//		Select(distance.FieldDistance).
//		Scan(ctx, &v)
//
func (dq *DistanceQuery) Select(field string, fields ...string) *DistanceSelect {
	selector := &DistanceSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return selector
}

func (dq *DistanceQuery) prepareQuery(ctx context.Context) error {
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DistanceQuery) sqlAll(ctx context.Context) ([]*Distance, error) {
	var (
		nodes       = []*Distance{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDisid != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Distance{config: dq.config}
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
		nodeids := make(map[int]*Distance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Carservice(func(s *sql.Selector) {
			s.Where(sql.InValues(distance.DisidColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.distance_disid
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "distance_disid" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "distance_disid" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Disid = append(node.Edges.Disid, n)
		}
	}

	return nodes, nil
}

func (dq *DistanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DistanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DistanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   distance.Table,
			Columns: distance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: distance.FieldID,
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

func (dq *DistanceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(distance.Table)
	selector := builder.Select(t1.Columns(distance.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(distance.Columns...)...)
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

// DistanceGroupBy is the builder for group-by Distance entities.
type DistanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DistanceGroupBy) Aggregate(fns ...AggregateFunc) *DistanceGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DistanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DistanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DistanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DistanceGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DistanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DistanceGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DistanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DistanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DistanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DistanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dgb *DistanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DistanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DistanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DistanceGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DistanceSelect is the builder for select fields of Distance entities.
type DistanceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DistanceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ds.path(ctx)
	if err != nil {
		return err
	}
	ds.sql = query
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DistanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DistanceSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DistanceSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DistanceSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DistanceSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DistanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DistanceSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DistanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DistanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ds *DistanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{distance.Label}
	default:
		err = fmt.Errorf("ent: DistanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DistanceSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DistanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DistanceSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
