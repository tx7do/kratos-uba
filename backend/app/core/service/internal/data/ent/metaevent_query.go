// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/metaevent"
	"kratos-uba/app/core/service/internal/data/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MetaEventQuery is the builder for querying MetaEvent entities.
type MetaEventQuery struct {
	config
	ctx        *QueryContext
	order      []metaevent.OrderOption
	inters     []Interceptor
	predicates []predicate.MetaEvent
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetaEventQuery builder.
func (meq *MetaEventQuery) Where(ps ...predicate.MetaEvent) *MetaEventQuery {
	meq.predicates = append(meq.predicates, ps...)
	return meq
}

// Limit the number of records to be returned by this query.
func (meq *MetaEventQuery) Limit(limit int) *MetaEventQuery {
	meq.ctx.Limit = &limit
	return meq
}

// Offset to start from.
func (meq *MetaEventQuery) Offset(offset int) *MetaEventQuery {
	meq.ctx.Offset = &offset
	return meq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (meq *MetaEventQuery) Unique(unique bool) *MetaEventQuery {
	meq.ctx.Unique = &unique
	return meq
}

// Order specifies how the records should be ordered.
func (meq *MetaEventQuery) Order(o ...metaevent.OrderOption) *MetaEventQuery {
	meq.order = append(meq.order, o...)
	return meq
}

// First returns the first MetaEvent entity from the query.
// Returns a *NotFoundError when no MetaEvent was found.
func (meq *MetaEventQuery) First(ctx context.Context) (*MetaEvent, error) {
	nodes, err := meq.Limit(1).All(setContextOp(ctx, meq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metaevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (meq *MetaEventQuery) FirstX(ctx context.Context) *MetaEvent {
	node, err := meq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MetaEvent ID from the query.
// Returns a *NotFoundError when no MetaEvent ID was found.
func (meq *MetaEventQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = meq.Limit(1).IDs(setContextOp(ctx, meq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metaevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (meq *MetaEventQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := meq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MetaEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MetaEvent entity is found.
// Returns a *NotFoundError when no MetaEvent entities are found.
func (meq *MetaEventQuery) Only(ctx context.Context) (*MetaEvent, error) {
	nodes, err := meq.Limit(2).All(setContextOp(ctx, meq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metaevent.Label}
	default:
		return nil, &NotSingularError{metaevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (meq *MetaEventQuery) OnlyX(ctx context.Context) *MetaEvent {
	node, err := meq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MetaEvent ID in the query.
// Returns a *NotSingularError when more than one MetaEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (meq *MetaEventQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = meq.Limit(2).IDs(setContextOp(ctx, meq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metaevent.Label}
	default:
		err = &NotSingularError{metaevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (meq *MetaEventQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := meq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetaEvents.
func (meq *MetaEventQuery) All(ctx context.Context) ([]*MetaEvent, error) {
	ctx = setContextOp(ctx, meq.ctx, ent.OpQueryAll)
	if err := meq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MetaEvent, *MetaEventQuery]()
	return withInterceptors[[]*MetaEvent](ctx, meq, qr, meq.inters)
}

// AllX is like All, but panics if an error occurs.
func (meq *MetaEventQuery) AllX(ctx context.Context) []*MetaEvent {
	nodes, err := meq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MetaEvent IDs.
func (meq *MetaEventQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if meq.ctx.Unique == nil && meq.path != nil {
		meq.Unique(true)
	}
	ctx = setContextOp(ctx, meq.ctx, ent.OpQueryIDs)
	if err = meq.Select(metaevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (meq *MetaEventQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := meq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (meq *MetaEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, meq.ctx, ent.OpQueryCount)
	if err := meq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, meq, querierCount[*MetaEventQuery](), meq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (meq *MetaEventQuery) CountX(ctx context.Context) int {
	count, err := meq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (meq *MetaEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, meq.ctx, ent.OpQueryExist)
	switch _, err := meq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (meq *MetaEventQuery) ExistX(ctx context.Context) bool {
	exist, err := meq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetaEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (meq *MetaEventQuery) Clone() *MetaEventQuery {
	if meq == nil {
		return nil
	}
	return &MetaEventQuery{
		config:     meq.config,
		ctx:        meq.ctx.Clone(),
		order:      append([]metaevent.OrderOption{}, meq.order...),
		inters:     append([]Interceptor{}, meq.inters...),
		predicates: append([]predicate.MetaEvent{}, meq.predicates...),
		// clone intermediate query.
		sql:       meq.sql.Clone(),
		path:      meq.path,
		modifiers: append([]func(*sql.Selector){}, meq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MetaEvent.Query().
//		GroupBy(metaevent.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (meq *MetaEventQuery) GroupBy(field string, fields ...string) *MetaEventGroupBy {
	meq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MetaEventGroupBy{build: meq}
	grbuild.flds = &meq.ctx.Fields
	grbuild.label = metaevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.MetaEvent.Query().
//		Select(metaevent.FieldCreateTime).
//		Scan(ctx, &v)
func (meq *MetaEventQuery) Select(fields ...string) *MetaEventSelect {
	meq.ctx.Fields = append(meq.ctx.Fields, fields...)
	sbuild := &MetaEventSelect{MetaEventQuery: meq}
	sbuild.label = metaevent.Label
	sbuild.flds, sbuild.scan = &meq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MetaEventSelect configured with the given aggregations.
func (meq *MetaEventQuery) Aggregate(fns ...AggregateFunc) *MetaEventSelect {
	return meq.Select().Aggregate(fns...)
}

func (meq *MetaEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range meq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, meq); err != nil {
				return err
			}
		}
	}
	for _, f := range meq.ctx.Fields {
		if !metaevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if meq.path != nil {
		prev, err := meq.path(ctx)
		if err != nil {
			return err
		}
		meq.sql = prev
	}
	return nil
}

func (meq *MetaEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MetaEvent, error) {
	var (
		nodes = []*MetaEvent{}
		_spec = meq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MetaEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MetaEvent{config: meq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(meq.modifiers) > 0 {
		_spec.Modifiers = meq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, meq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (meq *MetaEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := meq.querySpec()
	if len(meq.modifiers) > 0 {
		_spec.Modifiers = meq.modifiers
	}
	_spec.Node.Columns = meq.ctx.Fields
	if len(meq.ctx.Fields) > 0 {
		_spec.Unique = meq.ctx.Unique != nil && *meq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, meq.driver, _spec)
}

func (meq *MetaEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(metaevent.Table, metaevent.Columns, sqlgraph.NewFieldSpec(metaevent.FieldID, field.TypeUint32))
	_spec.From = meq.sql
	if unique := meq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if meq.path != nil {
		_spec.Unique = true
	}
	if fields := meq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metaevent.FieldID)
		for i := range fields {
			if fields[i] != metaevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := meq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := meq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := meq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := meq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (meq *MetaEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(meq.driver.Dialect())
	t1 := builder.Table(metaevent.Table)
	columns := meq.ctx.Fields
	if len(columns) == 0 {
		columns = metaevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if meq.sql != nil {
		selector = meq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if meq.ctx.Unique != nil && *meq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range meq.modifiers {
		m(selector)
	}
	for _, p := range meq.predicates {
		p(selector)
	}
	for _, p := range meq.order {
		p(selector)
	}
	if offset := meq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := meq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (meq *MetaEventQuery) Modify(modifiers ...func(s *sql.Selector)) *MetaEventSelect {
	meq.modifiers = append(meq.modifiers, modifiers...)
	return meq.Select()
}

// MetaEventGroupBy is the group-by builder for MetaEvent entities.
type MetaEventGroupBy struct {
	selector
	build *MetaEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (megb *MetaEventGroupBy) Aggregate(fns ...AggregateFunc) *MetaEventGroupBy {
	megb.fns = append(megb.fns, fns...)
	return megb
}

// Scan applies the selector query and scans the result into the given value.
func (megb *MetaEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, megb.build.ctx, ent.OpQueryGroupBy)
	if err := megb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MetaEventQuery, *MetaEventGroupBy](ctx, megb.build, megb, megb.build.inters, v)
}

func (megb *MetaEventGroupBy) sqlScan(ctx context.Context, root *MetaEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(megb.fns))
	for _, fn := range megb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*megb.flds)+len(megb.fns))
		for _, f := range *megb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*megb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := megb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MetaEventSelect is the builder for selecting fields of MetaEvent entities.
type MetaEventSelect struct {
	*MetaEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mes *MetaEventSelect) Aggregate(fns ...AggregateFunc) *MetaEventSelect {
	mes.fns = append(mes.fns, fns...)
	return mes
}

// Scan applies the selector query and scans the result into the given value.
func (mes *MetaEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mes.ctx, ent.OpQuerySelect)
	if err := mes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MetaEventQuery, *MetaEventSelect](ctx, mes.MetaEventQuery, mes, mes.inters, v)
}

func (mes *MetaEventSelect) sqlScan(ctx context.Context, root *MetaEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mes.fns))
	for _, fn := range mes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mes *MetaEventSelect) Modify(modifiers ...func(s *sql.Selector)) *MetaEventSelect {
	mes.modifiers = append(mes.modifiers, modifiers...)
	return mes
}
