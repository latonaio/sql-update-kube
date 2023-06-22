// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DataPlatformIndustryIndustryDatum is an object representing the database table.
type DataPlatformIndustryIndustryDatum struct {
	Industry string `boil:"Industry" json:"Industry" toml:"Industry" yaml:"Industry"`

	R *dataPlatformIndustryIndustryDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformIndustryIndustryDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformIndustryIndustryDatumColumns = struct {
	Industry string
}{
	Industry: "Industry",
}

var DataPlatformIndustryIndustryDatumTableColumns = struct {
	Industry string
}{
	Industry: "data_platform_industry_industry_data.Industry",
}

// Generated where

var DataPlatformIndustryIndustryDatumWhere = struct {
	Industry whereHelperstring
}{
	Industry: whereHelperstring{field: "`data_platform_industry_industry_data`.`Industry`"},
}

// DataPlatformIndustryIndustryDatumRels is where relationship names are stored.
var DataPlatformIndustryIndustryDatumRels = struct {
}{}

// dataPlatformIndustryIndustryDatumR is where relationships are stored.
type dataPlatformIndustryIndustryDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformIndustryIndustryDatumR) NewStruct() *dataPlatformIndustryIndustryDatumR {
	return &dataPlatformIndustryIndustryDatumR{}
}

// dataPlatformIndustryIndustryDatumL is where Load methods for each relationship are stored.
type dataPlatformIndustryIndustryDatumL struct{}

var (
	dataPlatformIndustryIndustryDatumAllColumns            = []string{"Industry"}
	dataPlatformIndustryIndustryDatumColumnsWithoutDefault = []string{"Industry"}
	dataPlatformIndustryIndustryDatumColumnsWithDefault    = []string{}
	dataPlatformIndustryIndustryDatumPrimaryKeyColumns     = []string{"Industry"}
	dataPlatformIndustryIndustryDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformIndustryIndustryDatumSlice is an alias for a slice of pointers to DataPlatformIndustryIndustryDatum.
	// This should almost always be used instead of []DataPlatformIndustryIndustryDatum.
	DataPlatformIndustryIndustryDatumSlice []*DataPlatformIndustryIndustryDatum

	dataPlatformIndustryIndustryDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformIndustryIndustryDatumType                 = reflect.TypeOf(&DataPlatformIndustryIndustryDatum{})
	dataPlatformIndustryIndustryDatumMapping              = queries.MakeStructMapping(dataPlatformIndustryIndustryDatumType)
	dataPlatformIndustryIndustryDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, dataPlatformIndustryIndustryDatumPrimaryKeyColumns)
	dataPlatformIndustryIndustryDatumInsertCacheMut       sync.RWMutex
	dataPlatformIndustryIndustryDatumInsertCache          = make(map[string]insertCache)
	dataPlatformIndustryIndustryDatumUpdateCacheMut       sync.RWMutex
	dataPlatformIndustryIndustryDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformIndustryIndustryDatumUpsertCacheMut       sync.RWMutex
	dataPlatformIndustryIndustryDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformIndustryIndustryDatum record from the query.
func (q dataPlatformIndustryIndustryDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformIndustryIndustryDatum, error) {
	o := &DataPlatformIndustryIndustryDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_industry_industry_data")
	}

	return o, nil
}

// All returns all DataPlatformIndustryIndustryDatum records from the query.
func (q dataPlatformIndustryIndustryDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformIndustryIndustryDatumSlice, error) {
	var o []*DataPlatformIndustryIndustryDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformIndustryIndustryDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformIndustryIndustryDatum records in the query.
func (q dataPlatformIndustryIndustryDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_industry_industry_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformIndustryIndustryDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_industry_industry_data exists")
	}

	return count > 0, nil
}

// DataPlatformIndustryIndustryData retrieves all the records using an executor.
func DataPlatformIndustryIndustryData(mods ...qm.QueryMod) dataPlatformIndustryIndustryDatumQuery {
	mods = append(mods, qm.From("`data_platform_industry_industry_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_industry_industry_data`.*"})
	}

	return dataPlatformIndustryIndustryDatumQuery{q}
}

// FindDataPlatformIndustryIndustryDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformIndustryIndustryDatum(ctx context.Context, exec boil.ContextExecutor, industry string, selectCols ...string) (*DataPlatformIndustryIndustryDatum, error) {
	dataPlatformIndustryIndustryDatumObj := &DataPlatformIndustryIndustryDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_industry_industry_data` where `Industry`=?", sel,
	)

	q := queries.Raw(query, industry)

	err := q.Bind(ctx, exec, dataPlatformIndustryIndustryDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_industry_industry_data")
	}

	return dataPlatformIndustryIndustryDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformIndustryIndustryDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_industry_industry_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIndustryIndustryDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformIndustryIndustryDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformIndustryIndustryDatumInsertCache[key]
	dataPlatformIndustryIndustryDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformIndustryIndustryDatumAllColumns,
			dataPlatformIndustryIndustryDatumColumnsWithDefault,
			dataPlatformIndustryIndustryDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_industry_industry_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_industry_industry_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_industry_industry_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformIndustryIndustryDatumPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into data_platform_industry_industry_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Industry,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_industry_industry_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIndustryIndustryDatumInsertCacheMut.Lock()
		dataPlatformIndustryIndustryDatumInsertCache[key] = cache
		dataPlatformIndustryIndustryDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformIndustryIndustryDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformIndustryIndustryDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformIndustryIndustryDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformIndustryIndustryDatumUpdateCache[key]
	dataPlatformIndustryIndustryDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformIndustryIndustryDatumAllColumns,
			dataPlatformIndustryIndustryDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_industry_industry_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_industry_industry_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformIndustryIndustryDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, append(wl, dataPlatformIndustryIndustryDatumPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update data_platform_industry_industry_data row")
	}

	if !cached {
		dataPlatformIndustryIndustryDatumUpdateCacheMut.Lock()
		dataPlatformIndustryIndustryDatumUpdateCache[key] = cache
		dataPlatformIndustryIndustryDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformIndustryIndustryDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_industry_industry_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformIndustryIndustryDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryIndustryDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_industry_industry_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryIndustryDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformIndustryIndustryDatum slice")
	}

	return nil
}

var mySQLDataPlatformIndustryIndustryDatumUniqueColumns = []string{
	"Industry",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformIndustryIndustryDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_industry_industry_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIndustryIndustryDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformIndustryIndustryDatumUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dataPlatformIndustryIndustryDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformIndustryIndustryDatumUpsertCache[key]
	dataPlatformIndustryIndustryDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformIndustryIndustryDatumAllColumns,
			dataPlatformIndustryIndustryDatumColumnsWithDefault,
			dataPlatformIndustryIndustryDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformIndustryIndustryDatumAllColumns,
			dataPlatformIndustryIndustryDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_industry_industry_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_industry_industry_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_industry_industry_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for data_platform_industry_industry_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformIndustryIndustryDatumType, dataPlatformIndustryIndustryDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_industry_industry_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_industry_industry_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIndustryIndustryDatumUpsertCacheMut.Lock()
		dataPlatformIndustryIndustryDatumUpsertCache[key] = cache
		dataPlatformIndustryIndustryDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformIndustryIndustryDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformIndustryIndustryDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformIndustryIndustryDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformIndustryIndustryDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_industry_industry_data` WHERE `Industry`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_industry_industry_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformIndustryIndustryDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformIndustryIndustryDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_industry_industry_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformIndustryIndustryDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryIndustryDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_industry_industry_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryIndustryDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformIndustryIndustryDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformIndustryIndustryDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformIndustryIndustryDatum(ctx, exec, o.Industry)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformIndustryIndustryDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformIndustryIndustryDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryIndustryDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_industry_industry_data`.* FROM `data_platform_industry_industry_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryIndustryDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformIndustryIndustryDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformIndustryIndustryDatumExists checks if the DataPlatformIndustryIndustryDatum row exists.
func DataPlatformIndustryIndustryDatumExists(ctx context.Context, exec boil.ContextExecutor, industry string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_industry_industry_data` where `Industry`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, industry)
	}
	row := exec.QueryRowContext(ctx, sql, industry)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_industry_industry_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformIndustryIndustryDatum row exists.
func (o *DataPlatformIndustryIndustryDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformIndustryIndustryDatumExists(ctx, exec, o.Industry)
}
