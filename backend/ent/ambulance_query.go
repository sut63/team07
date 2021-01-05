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
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carbrand"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/insurance"
	"github.com/team07/app/ent/predicate"
	"github.com/team07/app/ent/user"
)

// AmbulanceQuery is the builder for querying Ambulance entities.
type AmbulanceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Ambulance
	// eager-loading edges.
	withHasbrand       *CarbrandQuery
	withHasinsurance   *InsuranceQuery
	withHasstatus      *InspectionResultQuery
	withHasuser        *UserQuery
	withCarinspections *CarInspectionQuery
	withCarcheckinout  *CarCheckInOutQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (aq *AmbulanceQuery) Where(ps ...predicate.Ambulance) *AmbulanceQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AmbulanceQuery) Limit(limit int) *AmbulanceQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AmbulanceQuery) Offset(offset int) *AmbulanceQuery {
	aq.offset = &offset
	return aq
}

// Order adds an order step to the query.
func (aq *AmbulanceQuery) Order(o ...OrderFunc) *AmbulanceQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryHasbrand chains the current query on the hasbrand edge.
func (aq *AmbulanceQuery) QueryHasbrand() *CarbrandQuery {
	query := &CarbrandQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(carbrand.Table, carbrand.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ambulance.HasbrandTable, ambulance.HasbrandColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasinsurance chains the current query on the hasinsurance edge.
func (aq *AmbulanceQuery) QueryHasinsurance() *InsuranceQuery {
	query := &InsuranceQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(insurance.Table, insurance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ambulance.HasinsuranceTable, ambulance.HasinsuranceColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasstatus chains the current query on the hasstatus edge.
func (aq *AmbulanceQuery) QueryHasstatus() *InspectionResultQuery {
	query := &InspectionResultQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(inspectionresult.Table, inspectionresult.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ambulance.HasstatusTable, ambulance.HasstatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasuser chains the current query on the hasuser edge.
func (aq *AmbulanceQuery) QueryHasuser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ambulance.HasuserTable, ambulance.HasuserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCarinspections chains the current query on the carinspections edge.
func (aq *AmbulanceQuery) QueryCarinspections() *CarInspectionQuery {
	query := &CarInspectionQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(carinspection.Table, carinspection.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ambulance.CarinspectionsTable, ambulance.CarinspectionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCarcheckinout chains the current query on the carcheckinout edge.
func (aq *AmbulanceQuery) QueryCarcheckinout() *CarCheckInOutQuery {
	query := &CarCheckInOutQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambulance.Table, ambulance.FieldID, aq.sqlQuery()),
			sqlgraph.To(carcheckinout.Table, carcheckinout.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ambulance.CarcheckinoutTable, ambulance.CarcheckinoutColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Ambulance entity in the query. Returns *NotFoundError when no ambulance was found.
func (aq *AmbulanceQuery) First(ctx context.Context) (*Ambulance, error) {
	as, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(as) == 0 {
		return nil, &NotFoundError{ambulance.Label}
	}
	return as[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AmbulanceQuery) FirstX(ctx context.Context) *Ambulance {
	a, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return a
}

// FirstID returns the first Ambulance id in the query. Returns *NotFoundError when no id was found.
func (aq *AmbulanceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ambulance.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (aq *AmbulanceQuery) FirstXID(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Ambulance entity in the query, returns an error if not exactly one entity was returned.
func (aq *AmbulanceQuery) Only(ctx context.Context) (*Ambulance, error) {
	as, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(as) {
	case 1:
		return as[0], nil
	case 0:
		return nil, &NotFoundError{ambulance.Label}
	default:
		return nil, &NotSingularError{ambulance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AmbulanceQuery) OnlyX(ctx context.Context) *Ambulance {
	a, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// OnlyID returns the only Ambulance id in the query, returns an error if not exactly one id was returned.
func (aq *AmbulanceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = &NotSingularError{ambulance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AmbulanceQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ambulances.
func (aq *AmbulanceQuery) All(ctx context.Context) ([]*Ambulance, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AmbulanceQuery) AllX(ctx context.Context) []*Ambulance {
	as, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return as
}

// IDs executes the query and returns a list of Ambulance ids.
func (aq *AmbulanceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(ambulance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AmbulanceQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AmbulanceQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AmbulanceQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AmbulanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AmbulanceQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AmbulanceQuery) Clone() *AmbulanceQuery {
	return &AmbulanceQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		unique:     append([]string{}, aq.unique...),
		predicates: append([]predicate.Ambulance{}, aq.predicates...),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

//  WithHasbrand tells the query-builder to eager-loads the nodes that are connected to
// the "hasbrand" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithHasbrand(opts ...func(*CarbrandQuery)) *AmbulanceQuery {
	query := &CarbrandQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withHasbrand = query
	return aq
}

//  WithHasinsurance tells the query-builder to eager-loads the nodes that are connected to
// the "hasinsurance" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithHasinsurance(opts ...func(*InsuranceQuery)) *AmbulanceQuery {
	query := &InsuranceQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withHasinsurance = query
	return aq
}

//  WithHasstatus tells the query-builder to eager-loads the nodes that are connected to
// the "hasstatus" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithHasstatus(opts ...func(*InspectionResultQuery)) *AmbulanceQuery {
	query := &InspectionResultQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withHasstatus = query
	return aq
}

//  WithHasuser tells the query-builder to eager-loads the nodes that are connected to
// the "hasuser" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithHasuser(opts ...func(*UserQuery)) *AmbulanceQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withHasuser = query
	return aq
}

//  WithCarinspections tells the query-builder to eager-loads the nodes that are connected to
// the "carinspections" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithCarinspections(opts ...func(*CarInspectionQuery)) *AmbulanceQuery {
	query := &CarInspectionQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCarinspections = query
	return aq
}

//  WithCarcheckinout tells the query-builder to eager-loads the nodes that are connected to
// the "carcheckinout" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AmbulanceQuery) WithCarcheckinout(opts ...func(*CarCheckInOutQuery)) *AmbulanceQuery {
	query := &CarCheckInOutQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCarcheckinout = query
	return aq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Carregistration string `json:"carregistration,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ambulance.Query().
//		GroupBy(ambulance.FieldCarregistration).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AmbulanceQuery) GroupBy(field string, fields ...string) *AmbulanceGroupBy {
	group := &AmbulanceGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Carregistration string `json:"carregistration,omitempty"`
//	}
//
//	client.Ambulance.Query().
//		Select(ambulance.FieldCarregistration).
//		Scan(ctx, &v)
//
func (aq *AmbulanceQuery) Select(field string, fields ...string) *AmbulanceSelect {
	selector := &AmbulanceSelect{config: aq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(), nil
	}
	return selector
}

func (aq *AmbulanceQuery) prepareQuery(ctx context.Context) error {
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AmbulanceQuery) sqlAll(ctx context.Context) ([]*Ambulance, error) {
	var (
		nodes       = []*Ambulance{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [6]bool{
			aq.withHasbrand != nil,
			aq.withHasinsurance != nil,
			aq.withHasstatus != nil,
			aq.withHasuser != nil,
			aq.withCarinspections != nil,
			aq.withCarcheckinout != nil,
		}
	)
	if aq.withHasbrand != nil || aq.withHasinsurance != nil || aq.withHasstatus != nil || aq.withHasuser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ambulance.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Ambulance{config: aq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withHasbrand; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Ambulance)
		for i := range nodes {
			if fk := nodes[i].brand_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(carbrand.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "brand_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Hasbrand = n
			}
		}
	}

	if query := aq.withHasinsurance; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Ambulance)
		for i := range nodes {
			if fk := nodes[i].insurance_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(insurance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "insurance_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Hasinsurance = n
			}
		}
	}

	if query := aq.withHasstatus; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Ambulance)
		for i := range nodes {
			if fk := nodes[i].carstatus_id; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "carstatus_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Hasstatus = n
			}
		}
	}

	if query := aq.withHasuser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Ambulance)
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
				nodes[i].Edges.Hasuser = n
			}
		}
	}

	if query := aq.withCarinspections; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Ambulance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CarInspection(func(s *sql.Selector) {
			s.Where(sql.InValues(ambulance.CarinspectionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ambulance_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "ambulance_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ambulance_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Carinspections = append(node.Edges.Carinspections, n)
		}
	}

	if query := aq.withCarcheckinout; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Ambulance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CarCheckInOut(func(s *sql.Selector) {
			s.Where(sql.InValues(ambulance.CarcheckinoutColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ambulance
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "ambulance" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ambulance" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Carcheckinout = append(node.Edges.Carcheckinout, n)
		}
	}

	return nodes, nil
}

func (aq *AmbulanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AmbulanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aq *AmbulanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ambulance.Table,
			Columns: ambulance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ambulance.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AmbulanceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(ambulance.Table)
	selector := builder.Select(t1.Columns(ambulance.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(ambulance.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AmbulanceGroupBy is the builder for group-by Ambulance entities.
type AmbulanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AmbulanceGroupBy) Aggregate(fns ...AggregateFunc) *AmbulanceGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scan the result into the given value.
func (agb *AmbulanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AmbulanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AmbulanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AmbulanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *AmbulanceGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AmbulanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AmbulanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *AmbulanceGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AmbulanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AmbulanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *AmbulanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AmbulanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AmbulanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (agb *AmbulanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *AmbulanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AmbulanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agb.sqlQuery().Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AmbulanceGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// AmbulanceSelect is the builder for select fields of Ambulance entities.
type AmbulanceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (as *AmbulanceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := as.path(ctx)
	if err != nil {
		return err
	}
	as.sql = query
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AmbulanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AmbulanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AmbulanceSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *AmbulanceSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AmbulanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AmbulanceSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *AmbulanceSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AmbulanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AmbulanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *AmbulanceSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AmbulanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AmbulanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (as *AmbulanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambulance.Label}
	default:
		err = fmt.Errorf("ent: AmbulanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *AmbulanceSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AmbulanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *AmbulanceSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}
