// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"education/ent/predicate"
	"education/ent/upload"
	"education/ent/user"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UploadQuery is the builder for querying Upload entities.
type UploadQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Upload
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UploadQuery builder.
func (uq *UploadQuery) Where(ps ...predicate.Upload) *UploadQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UploadQuery) Limit(limit int) *UploadQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UploadQuery) Offset(offset int) *UploadQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UploadQuery) Unique(unique bool) *UploadQuery {
	uq.unique = &unique
	return uq
}

// Order adds an order step to the query.
func (uq *UploadQuery) Order(o ...OrderFunc) *UploadQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryUser chains the current query on the "user" edge.
func (uq *UploadQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upload.Table, upload.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upload.UserTable, upload.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Upload entity from the query.
// Returns a *NotFoundError when no Upload was found.
func (uq *UploadQuery) First(ctx context.Context) (*Upload, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upload.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UploadQuery) FirstX(ctx context.Context) *Upload {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Upload ID from the query.
// Returns a *NotFoundError when no Upload ID was found.
func (uq *UploadQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upload.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UploadQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Upload entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Upload entity is found.
// Returns a *NotFoundError when no Upload entities are found.
func (uq *UploadQuery) Only(ctx context.Context) (*Upload, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upload.Label}
	default:
		return nil, &NotSingularError{upload.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UploadQuery) OnlyX(ctx context.Context) *Upload {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Upload ID in the query.
// Returns a *NotSingularError when more than one Upload ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UploadQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upload.Label}
	default:
		err = &NotSingularError{upload.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UploadQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Uploads.
func (uq *UploadQuery) All(ctx context.Context) ([]*Upload, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UploadQuery) AllX(ctx context.Context) []*Upload {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Upload IDs.
func (uq *UploadQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(upload.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UploadQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UploadQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UploadQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UploadQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UploadQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UploadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UploadQuery) Clone() *UploadQuery {
	if uq == nil {
		return nil
	}
	return &UploadQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]OrderFunc{}, uq.order...),
		predicates: append([]predicate.Upload{}, uq.predicates...),
		withUser:   uq.withUser.Clone(),
		// clone intermediate query.
		sql:    uq.sql.Clone(),
		path:   uq.path,
		unique: uq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UploadQuery) WithUser(opts ...func(*UserQuery)) *UploadQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withUser = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Upload.Query().
//		GroupBy(upload.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UploadQuery) GroupBy(field string, fields ...string) *UploadGroupBy {
	grbuild := &UploadGroupBy{config: uq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(ctx), nil
	}
	grbuild.label = upload.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.Upload.Query().
//		Select(upload.FieldUserID).
//		Scan(ctx, &v)
func (uq *UploadQuery) Select(fields ...string) *UploadSelect {
	uq.fields = append(uq.fields, fields...)
	selbuild := &UploadSelect{UploadQuery: uq}
	selbuild.label = upload.Label
	selbuild.flds, selbuild.scan = &uq.fields, selbuild.Scan
	return selbuild
}

func (uq *UploadQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !upload.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UploadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Upload, error) {
	var (
		nodes       = []*Upload{}
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Upload).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Upload{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withUser; query != nil {
		if err := uq.loadUser(ctx, query, nodes, nil,
			func(n *Upload, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UploadQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Upload, init func(*Upload), assign func(*Upload, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Upload)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uq *UploadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.fields
	if len(uq.fields) > 0 {
		_spec.Unique = uq.unique != nil && *uq.unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UploadQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uq *UploadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upload.Table,
			Columns: upload.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upload.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upload.FieldID)
		for i := range fields {
			if fields[i] != upload.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UploadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(upload.Table)
	columns := uq.fields
	if len(columns) == 0 {
		columns = upload.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.unique != nil && *uq.unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UploadGroupBy is the group-by builder for Upload entities.
type UploadGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UploadGroupBy) Aggregate(fns ...AggregateFunc) *UploadGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scans the result into the given value.
func (ugb *UploadGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

func (ugb *UploadGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ugb.fields {
		if !upload.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UploadGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql.Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
		for _, f := range ugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ugb.fields...)...)
}

// UploadSelect is the builder for selecting fields of Upload entities.
type UploadSelect struct {
	*UploadQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (us *UploadSelect) Scan(ctx context.Context, v interface{}) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.sql = us.UploadQuery.sqlQuery(ctx)
	return us.sqlScan(ctx, v)
}

func (us *UploadSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sql.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
