// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Shikugawa/potraq/ent/club"
	"github.com/Shikugawa/potraq/ent/music"
	"github.com/Shikugawa/potraq/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MusicQuery is the builder for querying Music entities.
type MusicQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Music
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (mq *MusicQuery) Where(ps ...predicate.Music) *MusicQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MusicQuery) Limit(limit int) *MusicQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MusicQuery) Offset(offset int) *MusicQuery {
	mq.offset = &offset
	return mq
}

// Order adds an order step to the query.
func (mq *MusicQuery) Order(o ...Order) *MusicQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryClub chains the current query on the club edge.
func (mq *MusicQuery) QueryClub() *ClubQuery {
	query := &ClubQuery{config: mq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(music.Table, music.FieldID, mq.sqlQuery()),
		sqlgraph.To(club.Table, club.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, music.ClubTable, music.ClubColumn),
	)
	query.sql = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
	return query
}

// First returns the first Music entity in the query. Returns *ErrNotFound when no music was found.
func (mq *MusicQuery) First(ctx context.Context) (*Music, error) {
	ms, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, &ErrNotFound{music.Label}
	}
	return ms[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MusicQuery) FirstX(ctx context.Context) *Music {
	m, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return m
}

// FirstID returns the first Music id in the query. Returns *ErrNotFound when no id was found.
func (mq *MusicQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{music.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mq *MusicQuery) FirstXID(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Music entity in the query, returns an error if not exactly one entity was returned.
func (mq *MusicQuery) Only(ctx context.Context) (*Music, error) {
	ms, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ms) {
	case 1:
		return ms[0], nil
	case 0:
		return nil, &ErrNotFound{music.Label}
	default:
		return nil, &ErrNotSingular{music.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MusicQuery) OnlyX(ctx context.Context) *Music {
	m, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// OnlyID returns the only Music id in the query, returns an error if not exactly one id was returned.
func (mq *MusicQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{music.Label}
	default:
		err = &ErrNotSingular{music.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (mq *MusicQuery) OnlyXID(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Musics.
func (mq *MusicQuery) All(ctx context.Context) ([]*Music, error) {
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MusicQuery) AllX(ctx context.Context) []*Music {
	ms, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ms
}

// IDs executes the query and returns a list of Music ids.
func (mq *MusicQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(music.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MusicQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MusicQuery) Count(ctx context.Context) (int, error) {
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MusicQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MusicQuery) Exist(ctx context.Context) (bool, error) {
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MusicQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MusicQuery) Clone() *MusicQuery {
	return &MusicQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]Order{}, mq.order...),
		unique:     append([]string{}, mq.unique...),
		predicates: append([]predicate.Music{}, mq.predicates...),
		// clone intermediate query.
		sql: mq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mq *MusicQuery) GroupBy(field string, fields ...string) *MusicGroupBy {
	group := &MusicGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = mq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
func (mq *MusicQuery) Select(field string, fields ...string) *MusicSelect {
	selector := &MusicSelect{config: mq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = mq.sqlQuery()
	return selector
}

func (mq *MusicQuery) sqlAll(ctx context.Context) ([]*Music, error) {
	var (
		nodes []*Music
		spec  = mq.querySpec()
	)
	spec.ScanValues = func() []interface{} {
		node := &Music{config: mq.config}
		nodes = append(nodes, node)
		return node.scanValues()
	}
	spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, spec); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (mq *MusicQuery) sqlCount(ctx context.Context) (int, error) {
	spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, spec)
}

func (mq *MusicQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mq *MusicQuery) querySpec() *sqlgraph.QuerySpec {
	spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   music.Table,
			Columns: music.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: music.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if ps := mq.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return spec
}

func (mq *MusicQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(music.Table)
	selector := builder.Select(t1.Columns(music.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(music.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MusicGroupBy is the builder for group-by Music entities.
type MusicGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MusicGroupBy) Aggregate(fns ...Aggregate) *MusicGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mgb *MusicGroupBy) Scan(ctx context.Context, v interface{}) error {
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MusicGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mgb *MusicGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MusicGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MusicGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mgb *MusicGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MusicGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MusicGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mgb *MusicGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MusicGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MusicGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mgb *MusicGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MusicGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MusicGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MusicGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgb.sqlQuery().Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MusicGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// MusicSelect is the builder for select fields of Music entities.
type MusicSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ms *MusicSelect) Scan(ctx context.Context, v interface{}) error {
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MusicSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ms *MusicSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MusicSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MusicSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ms *MusicSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MusicSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MusicSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ms *MusicSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MusicSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MusicSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ms *MusicSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MusicSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MusicSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MusicSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *MusicSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
