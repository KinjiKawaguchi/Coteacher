// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coteacher/domain/repository/ent/class"
	"coteacher/domain/repository/ent/classinvitationcode"
	"coteacher/domain/repository/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClassInvitationCodeQuery is the builder for querying ClassInvitationCode entities.
type ClassInvitationCodeQuery struct {
	config
	ctx        *QueryContext
	order      []classinvitationcode.OrderOption
	inters     []Interceptor
	predicates []predicate.ClassInvitationCode
	withClass  *ClassQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassInvitationCodeQuery builder.
func (cicq *ClassInvitationCodeQuery) Where(ps ...predicate.ClassInvitationCode) *ClassInvitationCodeQuery {
	cicq.predicates = append(cicq.predicates, ps...)
	return cicq
}

// Limit the number of records to be returned by this query.
func (cicq *ClassInvitationCodeQuery) Limit(limit int) *ClassInvitationCodeQuery {
	cicq.ctx.Limit = &limit
	return cicq
}

// Offset to start from.
func (cicq *ClassInvitationCodeQuery) Offset(offset int) *ClassInvitationCodeQuery {
	cicq.ctx.Offset = &offset
	return cicq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cicq *ClassInvitationCodeQuery) Unique(unique bool) *ClassInvitationCodeQuery {
	cicq.ctx.Unique = &unique
	return cicq
}

// Order specifies how the records should be ordered.
func (cicq *ClassInvitationCodeQuery) Order(o ...classinvitationcode.OrderOption) *ClassInvitationCodeQuery {
	cicq.order = append(cicq.order, o...)
	return cicq
}

// QueryClass chains the current query on the "class" edge.
func (cicq *ClassInvitationCodeQuery) QueryClass() *ClassQuery {
	query := (&ClassClient{config: cicq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cicq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cicq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classinvitationcode.Table, classinvitationcode.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classinvitationcode.ClassTable, classinvitationcode.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(cicq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClassInvitationCode entity from the query.
// Returns a *NotFoundError when no ClassInvitationCode was found.
func (cicq *ClassInvitationCodeQuery) First(ctx context.Context) (*ClassInvitationCode, error) {
	nodes, err := cicq.Limit(1).All(setContextOp(ctx, cicq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{classinvitationcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) FirstX(ctx context.Context) *ClassInvitationCode {
	node, err := cicq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClassInvitationCode ID from the query.
// Returns a *NotFoundError when no ClassInvitationCode ID was found.
func (cicq *ClassInvitationCodeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cicq.Limit(1).IDs(setContextOp(ctx, cicq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{classinvitationcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cicq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClassInvitationCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClassInvitationCode entity is found.
// Returns a *NotFoundError when no ClassInvitationCode entities are found.
func (cicq *ClassInvitationCodeQuery) Only(ctx context.Context) (*ClassInvitationCode, error) {
	nodes, err := cicq.Limit(2).All(setContextOp(ctx, cicq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{classinvitationcode.Label}
	default:
		return nil, &NotSingularError{classinvitationcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) OnlyX(ctx context.Context) *ClassInvitationCode {
	node, err := cicq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClassInvitationCode ID in the query.
// Returns a *NotSingularError when more than one ClassInvitationCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (cicq *ClassInvitationCodeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cicq.Limit(2).IDs(setContextOp(ctx, cicq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{classinvitationcode.Label}
	default:
		err = &NotSingularError{classinvitationcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cicq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClassInvitationCodes.
func (cicq *ClassInvitationCodeQuery) All(ctx context.Context) ([]*ClassInvitationCode, error) {
	ctx = setContextOp(ctx, cicq.ctx, "All")
	if err := cicq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ClassInvitationCode, *ClassInvitationCodeQuery]()
	return withInterceptors[[]*ClassInvitationCode](ctx, cicq, qr, cicq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) AllX(ctx context.Context) []*ClassInvitationCode {
	nodes, err := cicq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClassInvitationCode IDs.
func (cicq *ClassInvitationCodeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cicq.ctx.Unique == nil && cicq.path != nil {
		cicq.Unique(true)
	}
	ctx = setContextOp(ctx, cicq.ctx, "IDs")
	if err = cicq.Select(classinvitationcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cicq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cicq *ClassInvitationCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cicq.ctx, "Count")
	if err := cicq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cicq, querierCount[*ClassInvitationCodeQuery](), cicq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) CountX(ctx context.Context) int {
	count, err := cicq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cicq *ClassInvitationCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cicq.ctx, "Exist")
	switch _, err := cicq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cicq *ClassInvitationCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := cicq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassInvitationCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cicq *ClassInvitationCodeQuery) Clone() *ClassInvitationCodeQuery {
	if cicq == nil {
		return nil
	}
	return &ClassInvitationCodeQuery{
		config:     cicq.config,
		ctx:        cicq.ctx.Clone(),
		order:      append([]classinvitationcode.OrderOption{}, cicq.order...),
		inters:     append([]Interceptor{}, cicq.inters...),
		predicates: append([]predicate.ClassInvitationCode{}, cicq.predicates...),
		withClass:  cicq.withClass.Clone(),
		// clone intermediate query.
		sql:  cicq.sql.Clone(),
		path: cicq.path,
	}
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (cicq *ClassInvitationCodeQuery) WithClass(opts ...func(*ClassQuery)) *ClassInvitationCodeQuery {
	query := (&ClassClient{config: cicq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cicq.withClass = query
	return cicq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClassID uuid.UUID `json:"class_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClassInvitationCode.Query().
//		GroupBy(classinvitationcode.FieldClassID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cicq *ClassInvitationCodeQuery) GroupBy(field string, fields ...string) *ClassInvitationCodeGroupBy {
	cicq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassInvitationCodeGroupBy{build: cicq}
	grbuild.flds = &cicq.ctx.Fields
	grbuild.label = classinvitationcode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClassID uuid.UUID `json:"class_id,omitempty"`
//	}
//
//	client.ClassInvitationCode.Query().
//		Select(classinvitationcode.FieldClassID).
//		Scan(ctx, &v)
func (cicq *ClassInvitationCodeQuery) Select(fields ...string) *ClassInvitationCodeSelect {
	cicq.ctx.Fields = append(cicq.ctx.Fields, fields...)
	sbuild := &ClassInvitationCodeSelect{ClassInvitationCodeQuery: cicq}
	sbuild.label = classinvitationcode.Label
	sbuild.flds, sbuild.scan = &cicq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassInvitationCodeSelect configured with the given aggregations.
func (cicq *ClassInvitationCodeQuery) Aggregate(fns ...AggregateFunc) *ClassInvitationCodeSelect {
	return cicq.Select().Aggregate(fns...)
}

func (cicq *ClassInvitationCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cicq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cicq); err != nil {
				return err
			}
		}
	}
	for _, f := range cicq.ctx.Fields {
		if !classinvitationcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cicq.path != nil {
		prev, err := cicq.path(ctx)
		if err != nil {
			return err
		}
		cicq.sql = prev
	}
	return nil
}

func (cicq *ClassInvitationCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClassInvitationCode, error) {
	var (
		nodes       = []*ClassInvitationCode{}
		_spec       = cicq.querySpec()
		loadedTypes = [1]bool{
			cicq.withClass != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ClassInvitationCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ClassInvitationCode{config: cicq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cicq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cicq.withClass; query != nil {
		if err := cicq.loadClass(ctx, query, nodes, nil,
			func(n *ClassInvitationCode, e *Class) { n.Edges.Class = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cicq *ClassInvitationCodeQuery) loadClass(ctx context.Context, query *ClassQuery, nodes []*ClassInvitationCode, init func(*ClassInvitationCode), assign func(*ClassInvitationCode, *Class)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ClassInvitationCode)
	for i := range nodes {
		fk := nodes[i].ClassID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(class.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cicq *ClassInvitationCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cicq.querySpec()
	_spec.Node.Columns = cicq.ctx.Fields
	if len(cicq.ctx.Fields) > 0 {
		_spec.Unique = cicq.ctx.Unique != nil && *cicq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cicq.driver, _spec)
}

func (cicq *ClassInvitationCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(classinvitationcode.Table, classinvitationcode.Columns, sqlgraph.NewFieldSpec(classinvitationcode.FieldID, field.TypeUUID))
	_spec.From = cicq.sql
	if unique := cicq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cicq.path != nil {
		_spec.Unique = true
	}
	if fields := cicq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classinvitationcode.FieldID)
		for i := range fields {
			if fields[i] != classinvitationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cicq.withClass != nil {
			_spec.Node.AddColumnOnce(classinvitationcode.FieldClassID)
		}
	}
	if ps := cicq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cicq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cicq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cicq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cicq *ClassInvitationCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cicq.driver.Dialect())
	t1 := builder.Table(classinvitationcode.Table)
	columns := cicq.ctx.Fields
	if len(columns) == 0 {
		columns = classinvitationcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cicq.sql != nil {
		selector = cicq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cicq.ctx.Unique != nil && *cicq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cicq.predicates {
		p(selector)
	}
	for _, p := range cicq.order {
		p(selector)
	}
	if offset := cicq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cicq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassInvitationCodeGroupBy is the group-by builder for ClassInvitationCode entities.
type ClassInvitationCodeGroupBy struct {
	selector
	build *ClassInvitationCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cicgb *ClassInvitationCodeGroupBy) Aggregate(fns ...AggregateFunc) *ClassInvitationCodeGroupBy {
	cicgb.fns = append(cicgb.fns, fns...)
	return cicgb
}

// Scan applies the selector query and scans the result into the given value.
func (cicgb *ClassInvitationCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cicgb.build.ctx, "GroupBy")
	if err := cicgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassInvitationCodeQuery, *ClassInvitationCodeGroupBy](ctx, cicgb.build, cicgb, cicgb.build.inters, v)
}

func (cicgb *ClassInvitationCodeGroupBy) sqlScan(ctx context.Context, root *ClassInvitationCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cicgb.fns))
	for _, fn := range cicgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cicgb.flds)+len(cicgb.fns))
		for _, f := range *cicgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cicgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cicgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClassInvitationCodeSelect is the builder for selecting fields of ClassInvitationCode entities.
type ClassInvitationCodeSelect struct {
	*ClassInvitationCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cics *ClassInvitationCodeSelect) Aggregate(fns ...AggregateFunc) *ClassInvitationCodeSelect {
	cics.fns = append(cics.fns, fns...)
	return cics
}

// Scan applies the selector query and scans the result into the given value.
func (cics *ClassInvitationCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cics.ctx, "Select")
	if err := cics.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassInvitationCodeQuery, *ClassInvitationCodeSelect](ctx, cics.ClassInvitationCodeQuery, cics, cics.inters, v)
}

func (cics *ClassInvitationCodeSelect) sqlScan(ctx context.Context, root *ClassInvitationCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cics.fns))
	for _, fn := range cics.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cics.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cics.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
