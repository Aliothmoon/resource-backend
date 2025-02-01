// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MirrorChyan/resource-backend/internal/ent/latestversion"
	"github.com/MirrorChyan/resource-backend/internal/ent/predicate"
	"github.com/MirrorChyan/resource-backend/internal/ent/resource"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// LatestVersionQuery is the builder for querying LatestVersion entities.
type LatestVersionQuery struct {
	config
	ctx          *QueryContext
	order        []latestversion.OrderOption
	inters       []Interceptor
	predicates   []predicate.LatestVersion
	withResource *ResourceQuery
	withVersion  *VersionQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LatestVersionQuery builder.
func (lvq *LatestVersionQuery) Where(ps ...predicate.LatestVersion) *LatestVersionQuery {
	lvq.predicates = append(lvq.predicates, ps...)
	return lvq
}

// Limit the number of records to be returned by this query.
func (lvq *LatestVersionQuery) Limit(limit int) *LatestVersionQuery {
	lvq.ctx.Limit = &limit
	return lvq
}

// Offset to start from.
func (lvq *LatestVersionQuery) Offset(offset int) *LatestVersionQuery {
	lvq.ctx.Offset = &offset
	return lvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lvq *LatestVersionQuery) Unique(unique bool) *LatestVersionQuery {
	lvq.ctx.Unique = &unique
	return lvq
}

// Order specifies how the records should be ordered.
func (lvq *LatestVersionQuery) Order(o ...latestversion.OrderOption) *LatestVersionQuery {
	lvq.order = append(lvq.order, o...)
	return lvq
}

// QueryResource chains the current query on the "resource" edge.
func (lvq *LatestVersionQuery) QueryResource() *ResourceQuery {
	query := (&ResourceClient{config: lvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(latestversion.Table, latestversion.FieldID, selector),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, latestversion.ResourceTable, latestversion.ResourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(lvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVersion chains the current query on the "version" edge.
func (lvq *LatestVersionQuery) QueryVersion() *VersionQuery {
	query := (&VersionClient{config: lvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(latestversion.Table, latestversion.FieldID, selector),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, latestversion.VersionTable, latestversion.VersionColumn),
		)
		fromU = sqlgraph.SetNeighbors(lvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LatestVersion entity from the query.
// Returns a *NotFoundError when no LatestVersion was found.
func (lvq *LatestVersionQuery) First(ctx context.Context) (*LatestVersion, error) {
	nodes, err := lvq.Limit(1).All(setContextOp(ctx, lvq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{latestversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lvq *LatestVersionQuery) FirstX(ctx context.Context) *LatestVersion {
	node, err := lvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LatestVersion ID from the query.
// Returns a *NotFoundError when no LatestVersion ID was found.
func (lvq *LatestVersionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lvq.Limit(1).IDs(setContextOp(ctx, lvq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{latestversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lvq *LatestVersionQuery) FirstIDX(ctx context.Context) int {
	id, err := lvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LatestVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LatestVersion entity is found.
// Returns a *NotFoundError when no LatestVersion entities are found.
func (lvq *LatestVersionQuery) Only(ctx context.Context) (*LatestVersion, error) {
	nodes, err := lvq.Limit(2).All(setContextOp(ctx, lvq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{latestversion.Label}
	default:
		return nil, &NotSingularError{latestversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lvq *LatestVersionQuery) OnlyX(ctx context.Context) *LatestVersion {
	node, err := lvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LatestVersion ID in the query.
// Returns a *NotSingularError when more than one LatestVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (lvq *LatestVersionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lvq.Limit(2).IDs(setContextOp(ctx, lvq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{latestversion.Label}
	default:
		err = &NotSingularError{latestversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lvq *LatestVersionQuery) OnlyIDX(ctx context.Context) int {
	id, err := lvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LatestVersions.
func (lvq *LatestVersionQuery) All(ctx context.Context) ([]*LatestVersion, error) {
	ctx = setContextOp(ctx, lvq.ctx, ent.OpQueryAll)
	if err := lvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LatestVersion, *LatestVersionQuery]()
	return withInterceptors[[]*LatestVersion](ctx, lvq, qr, lvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lvq *LatestVersionQuery) AllX(ctx context.Context) []*LatestVersion {
	nodes, err := lvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LatestVersion IDs.
func (lvq *LatestVersionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lvq.ctx.Unique == nil && lvq.path != nil {
		lvq.Unique(true)
	}
	ctx = setContextOp(ctx, lvq.ctx, ent.OpQueryIDs)
	if err = lvq.Select(latestversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lvq *LatestVersionQuery) IDsX(ctx context.Context) []int {
	ids, err := lvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lvq *LatestVersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lvq.ctx, ent.OpQueryCount)
	if err := lvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lvq, querierCount[*LatestVersionQuery](), lvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lvq *LatestVersionQuery) CountX(ctx context.Context) int {
	count, err := lvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lvq *LatestVersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lvq.ctx, ent.OpQueryExist)
	switch _, err := lvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lvq *LatestVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := lvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LatestVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lvq *LatestVersionQuery) Clone() *LatestVersionQuery {
	if lvq == nil {
		return nil
	}
	return &LatestVersionQuery{
		config:       lvq.config,
		ctx:          lvq.ctx.Clone(),
		order:        append([]latestversion.OrderOption{}, lvq.order...),
		inters:       append([]Interceptor{}, lvq.inters...),
		predicates:   append([]predicate.LatestVersion{}, lvq.predicates...),
		withResource: lvq.withResource.Clone(),
		withVersion:  lvq.withVersion.Clone(),
		// clone intermediate query.
		sql:  lvq.sql.Clone(),
		path: lvq.path,
	}
}

// WithResource tells the query-builder to eager-load the nodes that are connected to
// the "resource" edge. The optional arguments are used to configure the query builder of the edge.
func (lvq *LatestVersionQuery) WithResource(opts ...func(*ResourceQuery)) *LatestVersionQuery {
	query := (&ResourceClient{config: lvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lvq.withResource = query
	return lvq
}

// WithVersion tells the query-builder to eager-load the nodes that are connected to
// the "version" edge. The optional arguments are used to configure the query builder of the edge.
func (lvq *LatestVersionQuery) WithVersion(opts ...func(*VersionQuery)) *LatestVersionQuery {
	query := (&VersionClient{config: lvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lvq.withVersion = query
	return lvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Channel latestversion.Channel `json:"channel,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LatestVersion.Query().
//		GroupBy(latestversion.FieldChannel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lvq *LatestVersionQuery) GroupBy(field string, fields ...string) *LatestVersionGroupBy {
	lvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LatestVersionGroupBy{build: lvq}
	grbuild.flds = &lvq.ctx.Fields
	grbuild.label = latestversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Channel latestversion.Channel `json:"channel,omitempty"`
//	}
//
//	client.LatestVersion.Query().
//		Select(latestversion.FieldChannel).
//		Scan(ctx, &v)
func (lvq *LatestVersionQuery) Select(fields ...string) *LatestVersionSelect {
	lvq.ctx.Fields = append(lvq.ctx.Fields, fields...)
	sbuild := &LatestVersionSelect{LatestVersionQuery: lvq}
	sbuild.label = latestversion.Label
	sbuild.flds, sbuild.scan = &lvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LatestVersionSelect configured with the given aggregations.
func (lvq *LatestVersionQuery) Aggregate(fns ...AggregateFunc) *LatestVersionSelect {
	return lvq.Select().Aggregate(fns...)
}

func (lvq *LatestVersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lvq); err != nil {
				return err
			}
		}
	}
	for _, f := range lvq.ctx.Fields {
		if !latestversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lvq.path != nil {
		prev, err := lvq.path(ctx)
		if err != nil {
			return err
		}
		lvq.sql = prev
	}
	return nil
}

func (lvq *LatestVersionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LatestVersion, error) {
	var (
		nodes       = []*LatestVersion{}
		withFKs     = lvq.withFKs
		_spec       = lvq.querySpec()
		loadedTypes = [2]bool{
			lvq.withResource != nil,
			lvq.withVersion != nil,
		}
	)
	if lvq.withResource != nil || lvq.withVersion != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, latestversion.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LatestVersion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LatestVersion{config: lvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lvq.withResource; query != nil {
		if err := lvq.loadResource(ctx, query, nodes, nil,
			func(n *LatestVersion, e *Resource) { n.Edges.Resource = e }); err != nil {
			return nil, err
		}
	}
	if query := lvq.withVersion; query != nil {
		if err := lvq.loadVersion(ctx, query, nodes, nil,
			func(n *LatestVersion, e *Version) { n.Edges.Version = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lvq *LatestVersionQuery) loadResource(ctx context.Context, query *ResourceQuery, nodes []*LatestVersion, init func(*LatestVersion), assign func(*LatestVersion, *Resource)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*LatestVersion)
	for i := range nodes {
		if nodes[i].resource_latest_versions == nil {
			continue
		}
		fk := *nodes[i].resource_latest_versions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(resource.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "resource_latest_versions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lvq *LatestVersionQuery) loadVersion(ctx context.Context, query *VersionQuery, nodes []*LatestVersion, init func(*LatestVersion), assign func(*LatestVersion, *Version)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*LatestVersion)
	for i := range nodes {
		if nodes[i].latest_version_version == nil {
			continue
		}
		fk := *nodes[i].latest_version_version
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(version.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "latest_version_version" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lvq *LatestVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lvq.querySpec()
	_spec.Node.Columns = lvq.ctx.Fields
	if len(lvq.ctx.Fields) > 0 {
		_spec.Unique = lvq.ctx.Unique != nil && *lvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lvq.driver, _spec)
}

func (lvq *LatestVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(latestversion.Table, latestversion.Columns, sqlgraph.NewFieldSpec(latestversion.FieldID, field.TypeInt))
	_spec.From = lvq.sql
	if unique := lvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lvq.path != nil {
		_spec.Unique = true
	}
	if fields := lvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, latestversion.FieldID)
		for i := range fields {
			if fields[i] != latestversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lvq *LatestVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lvq.driver.Dialect())
	t1 := builder.Table(latestversion.Table)
	columns := lvq.ctx.Fields
	if len(columns) == 0 {
		columns = latestversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lvq.sql != nil {
		selector = lvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lvq.ctx.Unique != nil && *lvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lvq.predicates {
		p(selector)
	}
	for _, p := range lvq.order {
		p(selector)
	}
	if offset := lvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LatestVersionGroupBy is the group-by builder for LatestVersion entities.
type LatestVersionGroupBy struct {
	selector
	build *LatestVersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lvgb *LatestVersionGroupBy) Aggregate(fns ...AggregateFunc) *LatestVersionGroupBy {
	lvgb.fns = append(lvgb.fns, fns...)
	return lvgb
}

// Scan applies the selector query and scans the result into the given value.
func (lvgb *LatestVersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lvgb.build.ctx, ent.OpQueryGroupBy)
	if err := lvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LatestVersionQuery, *LatestVersionGroupBy](ctx, lvgb.build, lvgb, lvgb.build.inters, v)
}

func (lvgb *LatestVersionGroupBy) sqlScan(ctx context.Context, root *LatestVersionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lvgb.fns))
	for _, fn := range lvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lvgb.flds)+len(lvgb.fns))
		for _, f := range *lvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LatestVersionSelect is the builder for selecting fields of LatestVersion entities.
type LatestVersionSelect struct {
	*LatestVersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lvs *LatestVersionSelect) Aggregate(fns ...AggregateFunc) *LatestVersionSelect {
	lvs.fns = append(lvs.fns, fns...)
	return lvs
}

// Scan applies the selector query and scans the result into the given value.
func (lvs *LatestVersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lvs.ctx, ent.OpQuerySelect)
	if err := lvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LatestVersionQuery, *LatestVersionSelect](ctx, lvs.LatestVersionQuery, lvs, lvs.inters, v)
}

func (lvs *LatestVersionSelect) sqlScan(ctx context.Context, root *LatestVersionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lvs.fns))
	for _, fn := range lvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
