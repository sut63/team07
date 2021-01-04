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
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/repairing"
	"github.com/team07/app/ent/user"
)

// CarRepairrecordQuery is the builder for querying CarRepairrecord entities.
type CarRepairrecordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.CarRepairrecord
	// eager-loading edges.
	withKeeper        *RepairingQuery
	withUser          *UserQuery
	withCarinspection *CarInspectionQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (crq *CarRepairrecordQuery) Where(ps ...predicate.CarRepairrecord) *CarRepairrecordQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit adds a limit step to the query.
func (crq *CarRepairrecordQuery) Limit(limit int) *CarRepairrecordQuery {
	crq.limit = &limit
	return crq
}

// Offset adds an offset step to the query.
func (crq *CarRepairrecordQuery) Offset(offset int) *CarRepairrecordQuery {
	crq.offset = &offset
	return crq
}

// Order adds an order step to the query.
func (crq *CarRepairrecordQuery) Order(o ...OrderFunc) *CarRepairrecordQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryKeeper chains the current query on the keeper edge.
func (crq *CarRepairrecordQuery) QueryKeeper() *RepairingQuery {
	query := &RepairingQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carrepairrecord.Table, carrepairrecord.FieldID, crq.sqlQuery()),
			sqlgraph.To(repairing.Table, repairing.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carrepairrecord.KeeperTable, carrepairrecord.KeeperColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the user edge.
func (crq *CarRepairrecordQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carrepairrecord.Table, carrepairrecord.FieldID, crq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carrepairrecord.UserTable, carrepairrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCarinspection chains the current query on the carinspection edge.
func (crq *CarRepairrecordQuery) QueryCarinspection() *CarInspectionQuery {
	query := &CarInspectionQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(carrepairrecord.Table, carrepairrecord.FieldID, crq.sqlQuery()),
			sqlgraph.To(carinspection.Table, carinspection.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, carrepairrecord.CarinspectionTable, carrepairrecord.CarinspectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CarRepairrecord entity in the query. Returns *NotFoundError when no carrepairrecord was found.
func (crq *CarRepairrecordQuery) First(ctx context.Context) (*CarRepairrecord, error) {
	crs, err := crq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(crs) == 0 {
		return nil, &NotFoundError{carrepairrecord.Label}
	}
	return crs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CarRepairrecordQuery) FirstX(ctx context.Context) *CarRepairrecord {
	cr, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return cr
}

// FirstID returns the first CarRepairrecord id in the query. Returns *NotFoundError when no id was found.
func (crq *CarRepairrecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{carrepairrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (crq *CarRepairrecordQuery) FirstXID(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CarRepairrecord entity in the query, returns an error if not exactly one entity was returned.
func (crq *CarRepairrecordQuery) Only(ctx context.Context) (*CarRepairrecord, error) {
	crs, err := crq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(crs) {
	case 1:
		return crs[0], nil
	case 0:
		return nil, &NotFoundError{carrepairrecord.Label}
	default:
		return nil, &NotSingularError{carrepairrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CarRepairrecordQuery) OnlyX(ctx context.Context) *CarRepairrecord {
	cr, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return cr
}

// OnlyID returns the only CarRepairrecord id in the query, returns an error if not exactly one id was returned.
func (crq *CarRepairrecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = &NotSingularError{carrepairrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CarRepairrecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CarRepairrecords.
func (crq *CarRepairrecordQuery) All(ctx context.Context) ([]*CarRepairrecord, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return crq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (crq *CarRepairrecordQuery) AllX(ctx context.Context) []*CarRepairrecord {
	crs, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return crs
}

// IDs executes the query and returns a list of CarRepairrecord ids.
func (crq *CarRepairrecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := crq.Select(carrepairrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CarRepairrecordQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CarRepairrecordQuery) Count(ctx context.Context) (int, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return crq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CarRepairrecordQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CarRepairrecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return crq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CarRepairrecordQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CarRepairrecordQuery) Clone() *CarRepairrecordQuery {
	return &CarRepairrecordQuery{
		config:     crq.config,
		limit:      crq.limit,
		offset:     crq.offset,
		order:      append([]OrderFunc{}, crq.order...),
		unique:     append([]string{}, crq.unique...),
		predicates: append([]predicate.CarRepairrecord{}, crq.predicates...),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

//  WithKeeper tells the query-builder to eager-loads the nodes that are connected to
// the "keeper" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CarRepairrecordQuery) WithKeeper(opts ...func(*RepairingQuery)) *CarRepairrecordQuery {
	query := &RepairingQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withKeeper = query
	return crq
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CarRepairrecordQuery) WithUser(opts ...func(*UserQuery)) *CarRepairrecordQuery {
	query := &UserQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withUser = query
	return crq
}

//  WithCarinspection tells the query-builder to eager-loads the nodes that are connected to
// the "carinspection" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CarRepairrecordQuery) WithCarinspection(opts ...func(*CarInspectionQuery)) *CarRepairrecordQuery {
	query := &CarInspectionQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withCarinspection = query
	return crq
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
//	client.CarRepairrecord.Query().
//		GroupBy(carrepairrecord.FieldDatetime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (crq *CarRepairrecordQuery) GroupBy(field string, fields ...string) *CarRepairrecordGroupBy {
	group := &CarRepairrecordGroupBy{config: crq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return crq.sqlQuery(), nil
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
//	client.CarRepairrecord.Query().
//		Select(carrepairrecord.FieldDatetime).
//		Scan(ctx, &v)
//
func (crq *CarRepairrecordQuery) Select(field string, fields ...string) *CarRepairrecordSelect {
	selector := &CarRepairrecordSelect{config: crq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return crq.sqlQuery(), nil
	}
	return selector
}

func (crq *CarRepairrecordQuery) prepareQuery(ctx context.Context) error {
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CarRepairrecordQuery) sqlAll(ctx context.Context) ([]*CarRepairrecord, error) {
	var (
		nodes       = []*CarRepairrecord{}
		withFKs     = crq.withFKs
		_spec       = crq.querySpec()
		loadedTypes = [3]bool{
			crq.withKeeper != nil,
			crq.withUser != nil,
			crq.withCarinspection != nil,
		}
	)
	if crq.withKeeper != nil || crq.withUser != nil || crq.withCarinspection != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, carrepairrecord.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &CarRepairrecord{config: crq.config}
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
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := crq.withKeeper; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarRepairrecord)
		for i := range nodes {
			if fk := nodes[i].repairing_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(repairing.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "repairing_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Keeper = n
			}
		}
	}

	if query := crq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarRepairrecord)
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

	if query := crq.withCarinspection; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CarRepairrecord)
		for i := range nodes {
			if fk := nodes[i].carinspection_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(carinspection.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "carinspection_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Carinspection = n
			}
		}
	}

	return nodes, nil
}

func (crq *CarRepairrecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CarRepairrecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := crq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (crq *CarRepairrecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   carrepairrecord.Table,
			Columns: carrepairrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carrepairrecord.FieldID,
			},
		},
		From:   crq.sql,
		Unique: true,
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CarRepairrecordQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(carrepairrecord.Table)
	selector := builder.Select(t1.Columns(carrepairrecord.Columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(carrepairrecord.Columns...)...)
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CarRepairrecordGroupBy is the builder for group-by CarRepairrecord entities.
type CarRepairrecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CarRepairrecordGroupBy) Aggregate(fns ...AggregateFunc) *CarRepairrecordGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the group-by query and scan the result into the given value.
func (crgb *CarRepairrecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := crgb.path(ctx)
	if err != nil {
		return err
	}
	crgb.sql = query
	return crgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := crgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := crgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = crgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) StringX(ctx context.Context) string {
	v, err := crgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := crgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = crgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) IntX(ctx context.Context) int {
	v, err := crgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := crgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = crgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := crgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := crgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (crgb *CarRepairrecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = crgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (crgb *CarRepairrecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := crgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crgb *CarRepairrecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crgb.sqlQuery().Query()
	if err := crgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crgb *CarRepairrecordGroupBy) sqlQuery() *sql.Selector {
	selector := crgb.sql
	columns := make([]string, 0, len(crgb.fields)+len(crgb.fns))
	columns = append(columns, crgb.fields...)
	for _, fn := range crgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(crgb.fields...)
}

// CarRepairrecordSelect is the builder for select fields of CarRepairrecord entities.
type CarRepairrecordSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (crs *CarRepairrecordSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := crs.path(ctx)
	if err != nil {
		return err
	}
	crs.sql = query
	return crs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crs *CarRepairrecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := crs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crs *CarRepairrecordSelect) StringsX(ctx context.Context) []string {
	v, err := crs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = crs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (crs *CarRepairrecordSelect) StringX(ctx context.Context) string {
	v, err := crs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crs *CarRepairrecordSelect) IntsX(ctx context.Context) []int {
	v, err := crs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = crs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (crs *CarRepairrecordSelect) IntX(ctx context.Context) int {
	v, err := crs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crs *CarRepairrecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := crs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = crs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (crs *CarRepairrecordSelect) Float64X(ctx context.Context) float64 {
	v, err := crs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CarRepairrecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crs *CarRepairrecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := crs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (crs *CarRepairrecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = crs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{carrepairrecord.Label}
	default:
		err = fmt.Errorf("ent: CarRepairrecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (crs *CarRepairrecordSelect) BoolX(ctx context.Context) bool {
	v, err := crs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crs *CarRepairrecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crs.sqlQuery().Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crs *CarRepairrecordSelect) sqlQuery() sql.Querier {
	selector := crs.sql
	selector.Select(selector.Columns(crs.fields...)...)
	return selector
}
