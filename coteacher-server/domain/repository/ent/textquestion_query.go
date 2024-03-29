// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/question"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/textquestion"
	"github.com/google/uuid"
)

// TextQuestionQuery is the builder for querying TextQuestion entities.
type TextQuestionQuery struct {
	config
	ctx          *QueryContext
	order        []textquestion.OrderOption
	inters       []Interceptor
	predicates   []predicate.TextQuestion
	withQuestion *QuestionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TextQuestionQuery builder.
func (tqq *TextQuestionQuery) Where(ps ...predicate.TextQuestion) *TextQuestionQuery {
	tqq.predicates = append(tqq.predicates, ps...)
	return tqq
}

// Limit the number of records to be returned by this query.
func (tqq *TextQuestionQuery) Limit(limit int) *TextQuestionQuery {
	tqq.ctx.Limit = &limit
	return tqq
}

// Offset to start from.
func (tqq *TextQuestionQuery) Offset(offset int) *TextQuestionQuery {
	tqq.ctx.Offset = &offset
	return tqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tqq *TextQuestionQuery) Unique(unique bool) *TextQuestionQuery {
	tqq.ctx.Unique = &unique
	return tqq
}

// Order specifies how the records should be ordered.
func (tqq *TextQuestionQuery) Order(o ...textquestion.OrderOption) *TextQuestionQuery {
	tqq.order = append(tqq.order, o...)
	return tqq
}

// QueryQuestion chains the current query on the "question" edge.
func (tqq *TextQuestionQuery) QueryQuestion() *QuestionQuery {
	query := (&QuestionClient{config: tqq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(textquestion.Table, textquestion.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, textquestion.QuestionTable, textquestion.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TextQuestion entity from the query.
// Returns a *NotFoundError when no TextQuestion was found.
func (tqq *TextQuestionQuery) First(ctx context.Context) (*TextQuestion, error) {
	nodes, err := tqq.Limit(1).All(setContextOp(ctx, tqq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{textquestion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tqq *TextQuestionQuery) FirstX(ctx context.Context) *TextQuestion {
	node, err := tqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TextQuestion ID from the query.
// Returns a *NotFoundError when no TextQuestion ID was found.
func (tqq *TextQuestionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tqq.Limit(1).IDs(setContextOp(ctx, tqq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{textquestion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tqq *TextQuestionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TextQuestion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TextQuestion entity is found.
// Returns a *NotFoundError when no TextQuestion entities are found.
func (tqq *TextQuestionQuery) Only(ctx context.Context) (*TextQuestion, error) {
	nodes, err := tqq.Limit(2).All(setContextOp(ctx, tqq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{textquestion.Label}
	default:
		return nil, &NotSingularError{textquestion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tqq *TextQuestionQuery) OnlyX(ctx context.Context) *TextQuestion {
	node, err := tqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TextQuestion ID in the query.
// Returns a *NotSingularError when more than one TextQuestion ID is found.
// Returns a *NotFoundError when no entities are found.
func (tqq *TextQuestionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tqq.Limit(2).IDs(setContextOp(ctx, tqq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{textquestion.Label}
	default:
		err = &NotSingularError{textquestion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tqq *TextQuestionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TextQuestions.
func (tqq *TextQuestionQuery) All(ctx context.Context) ([]*TextQuestion, error) {
	ctx = setContextOp(ctx, tqq.ctx, "All")
	if err := tqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TextQuestion, *TextQuestionQuery]()
	return withInterceptors[[]*TextQuestion](ctx, tqq, qr, tqq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tqq *TextQuestionQuery) AllX(ctx context.Context) []*TextQuestion {
	nodes, err := tqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TextQuestion IDs.
func (tqq *TextQuestionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tqq.ctx.Unique == nil && tqq.path != nil {
		tqq.Unique(true)
	}
	ctx = setContextOp(ctx, tqq.ctx, "IDs")
	if err = tqq.Select(textquestion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tqq *TextQuestionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tqq *TextQuestionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tqq.ctx, "Count")
	if err := tqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tqq, querierCount[*TextQuestionQuery](), tqq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tqq *TextQuestionQuery) CountX(ctx context.Context) int {
	count, err := tqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tqq *TextQuestionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tqq.ctx, "Exist")
	switch _, err := tqq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tqq *TextQuestionQuery) ExistX(ctx context.Context) bool {
	exist, err := tqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TextQuestionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tqq *TextQuestionQuery) Clone() *TextQuestionQuery {
	if tqq == nil {
		return nil
	}
	return &TextQuestionQuery{
		config:       tqq.config,
		ctx:          tqq.ctx.Clone(),
		order:        append([]textquestion.OrderOption{}, tqq.order...),
		inters:       append([]Interceptor{}, tqq.inters...),
		predicates:   append([]predicate.TextQuestion{}, tqq.predicates...),
		withQuestion: tqq.withQuestion.Clone(),
		// clone intermediate query.
		sql:  tqq.sql.Clone(),
		path: tqq.path,
	}
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (tqq *TextQuestionQuery) WithQuestion(opts ...func(*QuestionQuery)) *TextQuestionQuery {
	query := (&QuestionClient{config: tqq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tqq.withQuestion = query
	return tqq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		QuestionID uuid.UUID `json:"question_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TextQuestion.Query().
//		GroupBy(textquestion.FieldQuestionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tqq *TextQuestionQuery) GroupBy(field string, fields ...string) *TextQuestionGroupBy {
	tqq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TextQuestionGroupBy{build: tqq}
	grbuild.flds = &tqq.ctx.Fields
	grbuild.label = textquestion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		QuestionID uuid.UUID `json:"question_id,omitempty"`
//	}
//
//	client.TextQuestion.Query().
//		Select(textquestion.FieldQuestionID).
//		Scan(ctx, &v)
func (tqq *TextQuestionQuery) Select(fields ...string) *TextQuestionSelect {
	tqq.ctx.Fields = append(tqq.ctx.Fields, fields...)
	sbuild := &TextQuestionSelect{TextQuestionQuery: tqq}
	sbuild.label = textquestion.Label
	sbuild.flds, sbuild.scan = &tqq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TextQuestionSelect configured with the given aggregations.
func (tqq *TextQuestionQuery) Aggregate(fns ...AggregateFunc) *TextQuestionSelect {
	return tqq.Select().Aggregate(fns...)
}

func (tqq *TextQuestionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tqq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tqq); err != nil {
				return err
			}
		}
	}
	for _, f := range tqq.ctx.Fields {
		if !textquestion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tqq.path != nil {
		prev, err := tqq.path(ctx)
		if err != nil {
			return err
		}
		tqq.sql = prev
	}
	return nil
}

func (tqq *TextQuestionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TextQuestion, error) {
	var (
		nodes       = []*TextQuestion{}
		_spec       = tqq.querySpec()
		loadedTypes = [1]bool{
			tqq.withQuestion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TextQuestion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TextQuestion{config: tqq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tqq.withQuestion; query != nil {
		if err := tqq.loadQuestion(ctx, query, nodes, nil,
			func(n *TextQuestion, e *Question) { n.Edges.Question = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tqq *TextQuestionQuery) loadQuestion(ctx context.Context, query *QuestionQuery, nodes []*TextQuestion, init func(*TextQuestion), assign func(*TextQuestion, *Question)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TextQuestion)
	for i := range nodes {
		fk := nodes[i].QuestionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(question.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "question_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tqq *TextQuestionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tqq.querySpec()
	_spec.Node.Columns = tqq.ctx.Fields
	if len(tqq.ctx.Fields) > 0 {
		_spec.Unique = tqq.ctx.Unique != nil && *tqq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tqq.driver, _spec)
}

func (tqq *TextQuestionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(textquestion.Table, textquestion.Columns, sqlgraph.NewFieldSpec(textquestion.FieldID, field.TypeUUID))
	_spec.From = tqq.sql
	if unique := tqq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tqq.path != nil {
		_spec.Unique = true
	}
	if fields := tqq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, textquestion.FieldID)
		for i := range fields {
			if fields[i] != textquestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tqq.withQuestion != nil {
			_spec.Node.AddColumnOnce(textquestion.FieldQuestionID)
		}
	}
	if ps := tqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tqq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tqq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tqq *TextQuestionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tqq.driver.Dialect())
	t1 := builder.Table(textquestion.Table)
	columns := tqq.ctx.Fields
	if len(columns) == 0 {
		columns = textquestion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tqq.sql != nil {
		selector = tqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tqq.ctx.Unique != nil && *tqq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tqq.predicates {
		p(selector)
	}
	for _, p := range tqq.order {
		p(selector)
	}
	if offset := tqq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tqq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TextQuestionGroupBy is the group-by builder for TextQuestion entities.
type TextQuestionGroupBy struct {
	selector
	build *TextQuestionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tqgb *TextQuestionGroupBy) Aggregate(fns ...AggregateFunc) *TextQuestionGroupBy {
	tqgb.fns = append(tqgb.fns, fns...)
	return tqgb
}

// Scan applies the selector query and scans the result into the given value.
func (tqgb *TextQuestionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tqgb.build.ctx, "GroupBy")
	if err := tqgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TextQuestionQuery, *TextQuestionGroupBy](ctx, tqgb.build, tqgb, tqgb.build.inters, v)
}

func (tqgb *TextQuestionGroupBy) sqlScan(ctx context.Context, root *TextQuestionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tqgb.fns))
	for _, fn := range tqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tqgb.flds)+len(tqgb.fns))
		for _, f := range *tqgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tqgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tqgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TextQuestionSelect is the builder for selecting fields of TextQuestion entities.
type TextQuestionSelect struct {
	*TextQuestionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tqs *TextQuestionSelect) Aggregate(fns ...AggregateFunc) *TextQuestionSelect {
	tqs.fns = append(tqs.fns, fns...)
	return tqs
}

// Scan applies the selector query and scans the result into the given value.
func (tqs *TextQuestionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tqs.ctx, "Select")
	if err := tqs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TextQuestionQuery, *TextQuestionSelect](ctx, tqs.TextQuestionQuery, tqs, tqs.inters, v)
}

func (tqs *TextQuestionSelect) sqlScan(ctx context.Context, root *TextQuestionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tqs.fns))
	for _, fn := range tqs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tqs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
