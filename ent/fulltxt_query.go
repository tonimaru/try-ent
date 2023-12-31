// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonimaru/try-ent/ent/fulltxt"
	"github.com/tonimaru/try-ent/ent/predicate"
)

// FulltxtQuery is the builder for querying Fulltxt entities.
type FulltxtQuery struct {
	config
	ctx        *QueryContext
	order      []fulltxt.OrderOption
	inters     []Interceptor
	predicates []predicate.Fulltxt
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FulltxtQuery builder.
func (fq *FulltxtQuery) Where(ps ...predicate.Fulltxt) *FulltxtQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FulltxtQuery) Limit(limit int) *FulltxtQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FulltxtQuery) Offset(offset int) *FulltxtQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FulltxtQuery) Unique(unique bool) *FulltxtQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FulltxtQuery) Order(o ...fulltxt.OrderOption) *FulltxtQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Fulltxt entity from the query.
// Returns a *NotFoundError when no Fulltxt was found.
func (fq *FulltxtQuery) First(ctx context.Context) (*Fulltxt, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fulltxt.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FulltxtQuery) FirstX(ctx context.Context) *Fulltxt {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Fulltxt ID from the query.
// Returns a *NotFoundError when no Fulltxt ID was found.
func (fq *FulltxtQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fulltxt.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FulltxtQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Fulltxt entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Fulltxt entity is found.
// Returns a *NotFoundError when no Fulltxt entities are found.
func (fq *FulltxtQuery) Only(ctx context.Context) (*Fulltxt, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fulltxt.Label}
	default:
		return nil, &NotSingularError{fulltxt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FulltxtQuery) OnlyX(ctx context.Context) *Fulltxt {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Fulltxt ID in the query.
// Returns a *NotSingularError when more than one Fulltxt ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FulltxtQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fulltxt.Label}
	default:
		err = &NotSingularError{fulltxt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FulltxtQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Fulltxts.
func (fq *FulltxtQuery) All(ctx context.Context) ([]*Fulltxt, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Fulltxt, *FulltxtQuery]()
	return withInterceptors[[]*Fulltxt](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FulltxtQuery) AllX(ctx context.Context) []*Fulltxt {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Fulltxt IDs.
func (fq *FulltxtQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(fulltxt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FulltxtQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FulltxtQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FulltxtQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FulltxtQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FulltxtQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FulltxtQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FulltxtQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FulltxtQuery) Clone() *FulltxtQuery {
	if fq == nil {
		return nil
	}
	return &FulltxtQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]fulltxt.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Fulltxt{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Txt string `json:"txt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Fulltxt.Query().
//		GroupBy(fulltxt.FieldTxt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FulltxtQuery) GroupBy(field string, fields ...string) *FulltxtGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FulltxtGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = fulltxt.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Txt string `json:"txt,omitempty"`
//	}
//
//	client.Fulltxt.Query().
//		Select(fulltxt.FieldTxt).
//		Scan(ctx, &v)
func (fq *FulltxtQuery) Select(fields ...string) *FulltxtSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FulltxtSelect{FulltxtQuery: fq}
	sbuild.label = fulltxt.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FulltxtSelect configured with the given aggregations.
func (fq *FulltxtQuery) Aggregate(fns ...AggregateFunc) *FulltxtSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FulltxtQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !fulltxt.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FulltxtQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Fulltxt, error) {
	var (
		nodes = []*Fulltxt{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Fulltxt).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Fulltxt{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fq *FulltxtQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FulltxtQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fulltxt.Table, fulltxt.Columns, sqlgraph.NewFieldSpec(fulltxt.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fulltxt.FieldID)
		for i := range fields {
			if fields[i] != fulltxt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FulltxtQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(fulltxt.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = fulltxt.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range fq.modifiers {
		m(selector)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fq *FulltxtQuery) Modify(modifiers ...func(s *sql.Selector)) *FulltxtSelect {
	fq.modifiers = append(fq.modifiers, modifiers...)
	return fq.Select()
}

// FulltxtGroupBy is the group-by builder for Fulltxt entities.
type FulltxtGroupBy struct {
	selector
	build *FulltxtQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FulltxtGroupBy) Aggregate(fns ...AggregateFunc) *FulltxtGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FulltxtGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FulltxtQuery, *FulltxtGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FulltxtGroupBy) sqlScan(ctx context.Context, root *FulltxtQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FulltxtSelect is the builder for selecting fields of Fulltxt entities.
type FulltxtSelect struct {
	*FulltxtQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FulltxtSelect) Aggregate(fns ...AggregateFunc) *FulltxtSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FulltxtSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FulltxtQuery, *FulltxtSelect](ctx, fs.FulltxtQuery, fs, fs.inters, v)
}

func (fs *FulltxtSelect) sqlScan(ctx context.Context, root *FulltxtQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fs *FulltxtSelect) Modify(modifiers ...func(s *sql.Selector)) *FulltxtSelect {
	fs.modifiers = append(fs.modifiers, modifiers...)
	return fs
}
