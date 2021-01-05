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
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/user"
)

// JobPositionQuery is the builder for querying JobPosition entities.
type JobPositionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.JobPosition
	// eager-loading edges.
	withUsers             *UserQuery
	withInspectionresults *InspectionResultQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (jpq *JobPositionQuery) Where(ps ...predicate.JobPosition) *JobPositionQuery {
	jpq.predicates = append(jpq.predicates, ps...)
	return jpq
}

// Limit adds a limit step to the query.
func (jpq *JobPositionQuery) Limit(limit int) *JobPositionQuery {
	jpq.limit = &limit
	return jpq
}

// Offset adds an offset step to the query.
func (jpq *JobPositionQuery) Offset(offset int) *JobPositionQuery {
	jpq.offset = &offset
	return jpq
}

// Order adds an order step to the query.
func (jpq *JobPositionQuery) Order(o ...OrderFunc) *JobPositionQuery {
	jpq.order = append(jpq.order, o...)
	return jpq
}

// QueryUsers chains the current query on the users edge.
func (jpq *JobPositionQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: jpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobposition.Table, jobposition.FieldID, jpq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, jobposition.UsersTable, jobposition.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(jpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInspectionresults chains the current query on the inspectionresults edge.
func (jpq *JobPositionQuery) QueryInspectionresults() *InspectionResultQuery {
	query := &InspectionResultQuery{config: jpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobposition.Table, jobposition.FieldID, jpq.sqlQuery()),
			sqlgraph.To(inspectionresult.Table, inspectionresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, jobposition.InspectionresultsTable, jobposition.InspectionresultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(jpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobPosition entity in the query. Returns *NotFoundError when no jobposition was found.
func (jpq *JobPositionQuery) First(ctx context.Context) (*JobPosition, error) {
	jps, err := jpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(jps) == 0 {
		return nil, &NotFoundError{jobposition.Label}
	}
	return jps[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jpq *JobPositionQuery) FirstX(ctx context.Context) *JobPosition {
	jp, err := jpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return jp
}

// FirstID returns the first JobPosition id in the query. Returns *NotFoundError when no id was found.
func (jpq *JobPositionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobposition.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (jpq *JobPositionQuery) FirstXID(ctx context.Context) int {
	id, err := jpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only JobPosition entity in the query, returns an error if not exactly one entity was returned.
func (jpq *JobPositionQuery) Only(ctx context.Context) (*JobPosition, error) {
	jps, err := jpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(jps) {
	case 1:
		return jps[0], nil
	case 0:
		return nil, &NotFoundError{jobposition.Label}
	default:
		return nil, &NotSingularError{jobposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jpq *JobPositionQuery) OnlyX(ctx context.Context) *JobPosition {
	jp, err := jpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return jp
}

// OnlyID returns the only JobPosition id in the query, returns an error if not exactly one id was returned.
func (jpq *JobPositionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = &NotSingularError{jobposition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jpq *JobPositionQuery) OnlyIDX(ctx context.Context) int {
	id, err := jpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobPositions.
func (jpq *JobPositionQuery) All(ctx context.Context) ([]*JobPosition, error) {
	if err := jpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jpq *JobPositionQuery) AllX(ctx context.Context) []*JobPosition {
	jps, err := jpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return jps
}

// IDs executes the query and returns a list of JobPosition ids.
func (jpq *JobPositionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jpq.Select(jobposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jpq *JobPositionQuery) IDsX(ctx context.Context) []int {
	ids, err := jpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jpq *JobPositionQuery) Count(ctx context.Context) (int, error) {
	if err := jpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jpq *JobPositionQuery) CountX(ctx context.Context) int {
	count, err := jpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jpq *JobPositionQuery) Exist(ctx context.Context) (bool, error) {
	if err := jpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jpq *JobPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := jpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jpq *JobPositionQuery) Clone() *JobPositionQuery {
	return &JobPositionQuery{
		config:     jpq.config,
		limit:      jpq.limit,
		offset:     jpq.offset,
		order:      append([]OrderFunc{}, jpq.order...),
		unique:     append([]string{}, jpq.unique...),
		predicates: append([]predicate.JobPosition{}, jpq.predicates...),
		// clone intermediate query.
		sql:  jpq.sql.Clone(),
		path: jpq.path,
	}
}

//  WithUsers tells the query-builder to eager-loads the nodes that are connected to
// the "users" edge. The optional arguments used to configure the query builder of the edge.
func (jpq *JobPositionQuery) WithUsers(opts ...func(*UserQuery)) *JobPositionQuery {
	query := &UserQuery{config: jpq.config}
	for _, opt := range opts {
		opt(query)
	}
	jpq.withUsers = query
	return jpq
}

//  WithInspectionresults tells the query-builder to eager-loads the nodes that are connected to
// the "inspectionresults" edge. The optional arguments used to configure the query builder of the edge.
func (jpq *JobPositionQuery) WithInspectionresults(opts ...func(*InspectionResultQuery)) *JobPositionQuery {
	query := &InspectionResultQuery{config: jpq.config}
	for _, opt := range opts {
		opt(query)
	}
	jpq.withInspectionresults = query
	return jpq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PositionName string `json:"position_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobPosition.Query().
//		GroupBy(jobposition.FieldPositionName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (jpq *JobPositionQuery) GroupBy(field string, fields ...string) *JobPositionGroupBy {
	group := &JobPositionGroupBy{config: jpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jpq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		PositionName string `json:"position_name,omitempty"`
//	}
//
//	client.JobPosition.Query().
//		Select(jobposition.FieldPositionName).
//		Scan(ctx, &v)
//
func (jpq *JobPositionQuery) Select(field string, fields ...string) *JobPositionSelect {
	selector := &JobPositionSelect{config: jpq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jpq.sqlQuery(), nil
	}
	return selector
}

func (jpq *JobPositionQuery) prepareQuery(ctx context.Context) error {
	if jpq.path != nil {
		prev, err := jpq.path(ctx)
		if err != nil {
			return err
		}
		jpq.sql = prev
	}
	return nil
}

func (jpq *JobPositionQuery) sqlAll(ctx context.Context) ([]*JobPosition, error) {
	var (
		nodes       = []*JobPosition{}
		_spec       = jpq.querySpec()
		loadedTypes = [2]bool{
			jpq.withUsers != nil,
			jpq.withInspectionresults != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &JobPosition{config: jpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, jpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := jpq.withUsers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*JobPosition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(jobposition.UsersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.jobposition_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "jobposition_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "jobposition_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Users = append(node.Edges.Users, n)
		}
	}

	if query := jpq.withInspectionresults; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*JobPosition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.InspectionResult(func(s *sql.Selector) {
			s.Where(sql.InValues(jobposition.InspectionresultsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.jobposition_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "jobposition_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "jobposition_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Inspectionresults = append(node.Edges.Inspectionresults, n)
		}
	}

	return nodes, nil
}

func (jpq *JobPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jpq.querySpec()
	return sqlgraph.CountNodes(ctx, jpq.driver, _spec)
}

func (jpq *JobPositionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := jpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (jpq *JobPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobposition.Table,
			Columns: jobposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobposition.FieldID,
			},
		},
		From:   jpq.sql,
		Unique: true,
	}
	if ps := jpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jpq *JobPositionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(jpq.driver.Dialect())
	t1 := builder.Table(jobposition.Table)
	selector := builder.Select(t1.Columns(jobposition.Columns...)...).From(t1)
	if jpq.sql != nil {
		selector = jpq.sql
		selector.Select(selector.Columns(jobposition.Columns...)...)
	}
	for _, p := range jpq.predicates {
		p(selector)
	}
	for _, p := range jpq.order {
		p(selector)
	}
	if offset := jpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobPositionGroupBy is the builder for group-by JobPosition entities.
type JobPositionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jpgb *JobPositionGroupBy) Aggregate(fns ...AggregateFunc) *JobPositionGroupBy {
	jpgb.fns = append(jpgb.fns, fns...)
	return jpgb
}

// Scan applies the group-by query and scan the result into the given value.
func (jpgb *JobPositionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jpgb.path(ctx)
	if err != nil {
		return err
	}
	jpgb.sql = query
	return jpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jpgb.fields) > 1 {
		return nil, errors.New("ent: JobPositionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) StringsX(ctx context.Context) []string {
	v, err := jpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) StringX(ctx context.Context) string {
	v, err := jpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jpgb.fields) > 1 {
		return nil, errors.New("ent: JobPositionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) IntsX(ctx context.Context) []int {
	v, err := jpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) IntX(ctx context.Context) int {
	v, err := jpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jpgb.fields) > 1 {
		return nil, errors.New("ent: JobPositionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jpgb.fields) > 1 {
		return nil, errors.New("ent: JobPositionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (jpgb *JobPositionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jpgb *JobPositionGroupBy) BoolX(ctx context.Context) bool {
	v, err := jpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jpgb *JobPositionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := jpgb.sqlQuery().Query()
	if err := jpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jpgb *JobPositionGroupBy) sqlQuery() *sql.Selector {
	selector := jpgb.sql
	columns := make([]string, 0, len(jpgb.fields)+len(jpgb.fns))
	columns = append(columns, jpgb.fields...)
	for _, fn := range jpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(jpgb.fields...)
}

// JobPositionSelect is the builder for select fields of JobPosition entities.
type JobPositionSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (jps *JobPositionSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := jps.path(ctx)
	if err != nil {
		return err
	}
	jps.sql = query
	return jps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jps *JobPositionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := jps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(jps.fields) > 1 {
		return nil, errors.New("ent: JobPositionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := jps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jps *JobPositionSelect) StringsX(ctx context.Context) []string {
	v, err := jps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jps *JobPositionSelect) StringX(ctx context.Context) string {
	v, err := jps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(jps.fields) > 1 {
		return nil, errors.New("ent: JobPositionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := jps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jps *JobPositionSelect) IntsX(ctx context.Context) []int {
	v, err := jps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jps *JobPositionSelect) IntX(ctx context.Context) int {
	v, err := jps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(jps.fields) > 1 {
		return nil, errors.New("ent: JobPositionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := jps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jps *JobPositionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := jps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jps *JobPositionSelect) Float64X(ctx context.Context) float64 {
	v, err := jps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(jps.fields) > 1 {
		return nil, errors.New("ent: JobPositionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := jps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jps *JobPositionSelect) BoolsX(ctx context.Context) []bool {
	v, err := jps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (jps *JobPositionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobposition.Label}
	default:
		err = fmt.Errorf("ent: JobPositionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jps *JobPositionSelect) BoolX(ctx context.Context) bool {
	v, err := jps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jps *JobPositionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := jps.sqlQuery().Query()
	if err := jps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jps *JobPositionSelect) sqlQuery() sql.Querier {
	selector := jps.sql
	selector.Select(selector.Columns(jps.fields...)...)
	return selector
}
