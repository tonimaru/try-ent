// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	entgeo "github.com/tonimaru/try-ent/ent/geo"
	"github.com/tonimaru/try-ent/ent/predicate"
)

// GeoQuery is the builder for querying Geo entities.
type GeoQuery struct {
	config
	ctx        *QueryContext
	order      []entgeo.OrderOption
	inters     []Interceptor
	predicates []predicate.Geo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GeoQuery builder.
func (gq *GeoQuery) Where(ps ...predicate.Geo) *GeoQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GeoQuery) Limit(limit int) *GeoQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GeoQuery) Offset(offset int) *GeoQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GeoQuery) Unique(unique bool) *GeoQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GeoQuery) Order(o ...entgeo.OrderOption) *GeoQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Geo entity from the query.
// Returns a *NotFoundError when no Geo was found.
func (gq *GeoQuery) First(ctx context.Context) (*Geo, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entgeo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GeoQuery) FirstX(ctx context.Context) *Geo {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Geo ID from the query.
// Returns a *NotFoundError when no Geo ID was found.
func (gq *GeoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entgeo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GeoQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Geo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Geo entity is found.
// Returns a *NotFoundError when no Geo entities are found.
func (gq *GeoQuery) Only(ctx context.Context) (*Geo, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entgeo.Label}
	default:
		return nil, &NotSingularError{entgeo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GeoQuery) OnlyX(ctx context.Context) *Geo {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Geo ID in the query.
// Returns a *NotSingularError when more than one Geo ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GeoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entgeo.Label}
	default:
		err = &NotSingularError{entgeo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GeoQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Geos.
func (gq *GeoQuery) All(ctx context.Context) ([]*Geo, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Geo, *GeoQuery]()
	return withInterceptors[[]*Geo](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GeoQuery) AllX(ctx context.Context) []*Geo {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Geo IDs.
func (gq *GeoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(entgeo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GeoQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GeoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GeoQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GeoQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GeoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GeoQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GeoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GeoQuery) Clone() *GeoQuery {
	if gq == nil {
		return nil
	}
	return &GeoQuery{
		config:     gq.config,
		ctx:        gq.ctx.Clone(),
		order:      append([]entgeo.OrderOption{}, gq.order...),
		inters:     append([]Interceptor{}, gq.inters...),
		predicates: append([]predicate.Geo{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Point *geo.Point `json:"point,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Geo.Query().
//		GroupBy(entgeo.FieldPoint).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GeoQuery) GroupBy(field string, fields ...string) *GeoGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GeoGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = entgeo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Point *geo.Point `json:"point,omitempty"`
//	}
//
//	client.Geo.Query().
//		Select(entgeo.FieldPoint).
//		Scan(ctx, &v)
func (gq *GeoQuery) Select(fields ...string) *GeoSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GeoSelect{GeoQuery: gq}
	sbuild.label = entgeo.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GeoSelect configured with the given aggregations.
func (gq *GeoQuery) Aggregate(fns ...AggregateFunc) *GeoSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GeoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !entgeo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GeoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Geo, error) {
	var (
		nodes = []*Geo{}
		_spec = gq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Geo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Geo{config: gq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gq *GeoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GeoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entgeo.Table, entgeo.Columns, sqlgraph.NewFieldSpec(entgeo.FieldID, field.TypeInt))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entgeo.FieldID)
		for i := range fields {
			if fields[i] != entgeo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GeoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(entgeo.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = entgeo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gq *GeoQuery) Modify(modifiers ...func(s *sql.Selector)) *GeoSelect {
	gq.modifiers = append(gq.modifiers, modifiers...)
	return gq.Select()
}

// GeoGroupBy is the group-by builder for Geo entities.
type GeoGroupBy struct {
	selector
	build *GeoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GeoGroupBy) Aggregate(fns ...AggregateFunc) *GeoGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GeoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeoQuery, *GeoGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GeoGroupBy) sqlScan(ctx context.Context, root *GeoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GeoSelect is the builder for selecting fields of Geo entities.
type GeoSelect struct {
	*GeoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GeoSelect) Aggregate(fns ...AggregateFunc) *GeoSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GeoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeoQuery, *GeoSelect](ctx, gs.GeoQuery, gs, gs.inters, v)
}

func (gs *GeoSelect) sqlScan(ctx context.Context, root *GeoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GeoSelect) Modify(modifiers ...func(s *sql.Selector)) *GeoSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}
