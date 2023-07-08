// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-uba/app/core/service/internal/data/ent/attribute"
	"kratos-uba/app/core/service/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeQuery is the builder for querying Attribute entities.
type AttributeQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Attribute
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeQuery builder.
func (aq *AttributeQuery) Where(ps ...predicate.Attribute) *AttributeQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AttributeQuery) Limit(limit int) *AttributeQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AttributeQuery) Offset(offset int) *AttributeQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttributeQuery) Unique(unique bool) *AttributeQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AttributeQuery) Order(o ...OrderFunc) *AttributeQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// First returns the first Attribute entity from the query.
// Returns a *NotFoundError when no Attribute was found.
func (aq *AttributeQuery) First(ctx context.Context) (*Attribute, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attribute.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttributeQuery) FirstX(ctx context.Context) *Attribute {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attribute ID from the query.
// Returns a *NotFoundError when no Attribute ID was found.
func (aq *AttributeQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attribute.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttributeQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attribute entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attribute entity is found.
// Returns a *NotFoundError when no Attribute entities are found.
func (aq *AttributeQuery) Only(ctx context.Context) (*Attribute, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attribute.Label}
	default:
		return nil, &NotSingularError{attribute.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttributeQuery) OnlyX(ctx context.Context) *Attribute {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attribute ID in the query.
// Returns a *NotSingularError when more than one Attribute ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttributeQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attribute.Label}
	default:
		err = &NotSingularError{attribute.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttributeQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attributes.
func (aq *AttributeQuery) All(ctx context.Context) ([]*Attribute, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Attribute, *AttributeQuery]()
	return withInterceptors[[]*Attribute](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttributeQuery) AllX(ctx context.Context) []*Attribute {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attribute IDs.
func (aq *AttributeQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(attribute.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttributeQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttributeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AttributeQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttributeQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttributeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttributeQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttributeQuery) Clone() *AttributeQuery {
	if aq == nil {
		return nil
	}
	return &AttributeQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]OrderFunc{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Attribute{}, aq.predicates...),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Attribute.Query().
//		GroupBy(attribute.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttributeQuery) GroupBy(field string, fields ...string) *AttributeGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = attribute.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//	}
//
//	client.Attribute.Query().
//		Select(attribute.FieldCreateTime).
//		Scan(ctx, &v)
func (aq *AttributeQuery) Select(fields ...string) *AttributeSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AttributeSelect{AttributeQuery: aq}
	sbuild.label = attribute.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeSelect configured with the given aggregations.
func (aq *AttributeQuery) Aggregate(fns ...AggregateFunc) *AttributeSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AttributeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !attribute.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttributeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attribute, error) {
	var (
		nodes = []*Attribute{}
		_spec = aq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Attribute).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Attribute{config: aq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aq *AttributeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttributeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeUint32))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attribute.FieldID)
		for i := range fields {
			if fields[i] != attribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
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

func (aq *AttributeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attribute.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = attribute.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AttributeQuery) Modify(modifiers ...func(s *sql.Selector)) *AttributeSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// AttributeGroupBy is the group-by builder for Attribute entities.
type AttributeGroupBy struct {
	selector
	build *AttributeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttributeGroupBy) Aggregate(fns ...AggregateFunc) *AttributeGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AttributeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeQuery, *AttributeGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AttributeGroupBy) sqlScan(ctx context.Context, root *AttributeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeSelect is the builder for selecting fields of Attribute entities.
type AttributeSelect struct {
	*AttributeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AttributeSelect) Aggregate(fns ...AggregateFunc) *AttributeSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttributeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeQuery, *AttributeSelect](ctx, as.AttributeQuery, as, as.inters, v)
}

func (as *AttributeSelect) sqlScan(ctx context.Context, root *AttributeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AttributeSelect) Modify(modifiers ...func(s *sql.Selector)) *AttributeSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}
