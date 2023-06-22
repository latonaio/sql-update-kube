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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DataPlatformProductGroupProductGroupTextDatum is an object representing the database table.
type DataPlatformProductGroupProductGroupTextDatum struct {
	ProductGroup     string      `boil:"ProductGroup" json:"ProductGroup" toml:"ProductGroup" yaml:"ProductGroup"`
	Language         string      `boil:"Language" json:"Language" toml:"Language" yaml:"Language"`
	ProductGroupName null.String `boil:"ProductGroupName" json:"ProductGroupName,omitempty" toml:"ProductGroupName" yaml:"ProductGroupName,omitempty"`

	R *dataPlatformProductGroupProductGroupTextDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformProductGroupProductGroupTextDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformProductGroupProductGroupTextDatumColumns = struct {
	ProductGroup     string
	Language         string
	ProductGroupName string
}{
	ProductGroup:     "ProductGroup",
	Language:         "Language",
	ProductGroupName: "ProductGroupName",
}

var DataPlatformProductGroupProductGroupTextDatumTableColumns = struct {
	ProductGroup     string
	Language         string
	ProductGroupName string
}{
	ProductGroup:     "data_platform_product_group_product_group_text_data.ProductGroup",
	Language:         "data_platform_product_group_product_group_text_data.Language",
	ProductGroupName: "data_platform_product_group_product_group_text_data.ProductGroupName",
}

// Generated where

var DataPlatformProductGroupProductGroupTextDatumWhere = struct {
	ProductGroup     whereHelperstring
	Language         whereHelperstring
	ProductGroupName whereHelpernull_String
}{
	ProductGroup:     whereHelperstring{field: "`data_platform_product_group_product_group_text_data`.`ProductGroup`"},
	Language:         whereHelperstring{field: "`data_platform_product_group_product_group_text_data`.`Language`"},
	ProductGroupName: whereHelpernull_String{field: "`data_platform_product_group_product_group_text_data`.`ProductGroupName`"},
}

// DataPlatformProductGroupProductGroupTextDatumRels is where relationship names are stored.
var DataPlatformProductGroupProductGroupTextDatumRels = struct {
}{}

// dataPlatformProductGroupProductGroupTextDatumR is where relationships are stored.
type dataPlatformProductGroupProductGroupTextDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformProductGroupProductGroupTextDatumR) NewStruct() *dataPlatformProductGroupProductGroupTextDatumR {
	return &dataPlatformProductGroupProductGroupTextDatumR{}
}

// dataPlatformProductGroupProductGroupTextDatumL is where Load methods for each relationship are stored.
type dataPlatformProductGroupProductGroupTextDatumL struct{}

var (
	dataPlatformProductGroupProductGroupTextDatumAllColumns            = []string{"ProductGroup", "Language", "ProductGroupName"}
	dataPlatformProductGroupProductGroupTextDatumColumnsWithoutDefault = []string{"ProductGroup", "Language", "ProductGroupName"}
	dataPlatformProductGroupProductGroupTextDatumColumnsWithDefault    = []string{}
	dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns     = []string{"ProductGroup", "Language"}
	dataPlatformProductGroupProductGroupTextDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformProductGroupProductGroupTextDatumSlice is an alias for a slice of pointers to DataPlatformProductGroupProductGroupTextDatum.
	// This should almost always be used instead of []DataPlatformProductGroupProductGroupTextDatum.
	DataPlatformProductGroupProductGroupTextDatumSlice []*DataPlatformProductGroupProductGroupTextDatum

	dataPlatformProductGroupProductGroupTextDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformProductGroupProductGroupTextDatumType                 = reflect.TypeOf(&DataPlatformProductGroupProductGroupTextDatum{})
	dataPlatformProductGroupProductGroupTextDatumMapping              = queries.MakeStructMapping(dataPlatformProductGroupProductGroupTextDatumType)
	dataPlatformProductGroupProductGroupTextDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns)
	dataPlatformProductGroupProductGroupTextDatumInsertCacheMut       sync.RWMutex
	dataPlatformProductGroupProductGroupTextDatumInsertCache          = make(map[string]insertCache)
	dataPlatformProductGroupProductGroupTextDatumUpdateCacheMut       sync.RWMutex
	dataPlatformProductGroupProductGroupTextDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformProductGroupProductGroupTextDatumUpsertCacheMut       sync.RWMutex
	dataPlatformProductGroupProductGroupTextDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformProductGroupProductGroupTextDatum record from the query.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformProductGroupProductGroupTextDatum, error) {
	o := &DataPlatformProductGroupProductGroupTextDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_product_group_product_group_text_data")
	}

	return o, nil
}

// All returns all DataPlatformProductGroupProductGroupTextDatum records from the query.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformProductGroupProductGroupTextDatumSlice, error) {
	var o []*DataPlatformProductGroupProductGroupTextDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformProductGroupProductGroupTextDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformProductGroupProductGroupTextDatum records in the query.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_product_group_product_group_text_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_product_group_product_group_text_data exists")
	}

	return count > 0, nil
}

// DataPlatformProductGroupProductGroupTextData retrieves all the records using an executor.
func DataPlatformProductGroupProductGroupTextData(mods ...qm.QueryMod) dataPlatformProductGroupProductGroupTextDatumQuery {
	mods = append(mods, qm.From("`data_platform_product_group_product_group_text_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_product_group_product_group_text_data`.*"})
	}

	return dataPlatformProductGroupProductGroupTextDatumQuery{q}
}

// FindDataPlatformProductGroupProductGroupTextDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformProductGroupProductGroupTextDatum(ctx context.Context, exec boil.ContextExecutor, productGroup string, language string, selectCols ...string) (*DataPlatformProductGroupProductGroupTextDatum, error) {
	dataPlatformProductGroupProductGroupTextDatumObj := &DataPlatformProductGroupProductGroupTextDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_product_group_product_group_text_data` where `ProductGroup`=? AND `Language`=?", sel,
	)

	q := queries.Raw(query, productGroup, language)

	err := q.Bind(ctx, exec, dataPlatformProductGroupProductGroupTextDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_product_group_product_group_text_data")
	}

	return dataPlatformProductGroupProductGroupTextDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformProductGroupProductGroupTextDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_group_product_group_text_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductGroupProductGroupTextDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformProductGroupProductGroupTextDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformProductGroupProductGroupTextDatumInsertCache[key]
	dataPlatformProductGroupProductGroupTextDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformProductGroupProductGroupTextDatumAllColumns,
			dataPlatformProductGroupProductGroupTextDatumColumnsWithDefault,
			dataPlatformProductGroupProductGroupTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_product_group_product_group_text_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_product_group_product_group_text_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_product_group_product_group_text_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_product_group_product_group_text_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ProductGroup,
		o.Language,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_group_product_group_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductGroupProductGroupTextDatumInsertCacheMut.Lock()
		dataPlatformProductGroupProductGroupTextDatumInsertCache[key] = cache
		dataPlatformProductGroupProductGroupTextDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformProductGroupProductGroupTextDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformProductGroupProductGroupTextDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformProductGroupProductGroupTextDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformProductGroupProductGroupTextDatumUpdateCache[key]
	dataPlatformProductGroupProductGroupTextDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformProductGroupProductGroupTextDatumAllColumns,
			dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_product_group_product_group_text_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_product_group_product_group_text_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, append(wl, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_product_group_product_group_text_data row")
	}

	if !cached {
		dataPlatformProductGroupProductGroupTextDatumUpdateCacheMut.Lock()
		dataPlatformProductGroupProductGroupTextDatumUpdateCache[key] = cache
		dataPlatformProductGroupProductGroupTextDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_product_group_product_group_text_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformProductGroupProductGroupTextDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductGroupProductGroupTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_product_group_product_group_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformProductGroupProductGroupTextDatum slice")
	}

	return nil
}

var mySQLDataPlatformProductGroupProductGroupTextDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformProductGroupProductGroupTextDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_group_product_group_text_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductGroupProductGroupTextDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformProductGroupProductGroupTextDatumUniqueColumns, o)

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

	dataPlatformProductGroupProductGroupTextDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformProductGroupProductGroupTextDatumUpsertCache[key]
	dataPlatformProductGroupProductGroupTextDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformProductGroupProductGroupTextDatumAllColumns,
			dataPlatformProductGroupProductGroupTextDatumColumnsWithDefault,
			dataPlatformProductGroupProductGroupTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformProductGroupProductGroupTextDatumAllColumns,
			dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_product_group_product_group_text_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_product_group_product_group_text_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_product_group_product_group_text_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_product_group_product_group_text_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformProductGroupProductGroupTextDatumType, dataPlatformProductGroupProductGroupTextDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_product_group_product_group_text_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_group_product_group_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductGroupProductGroupTextDatumUpsertCacheMut.Lock()
		dataPlatformProductGroupProductGroupTextDatumUpsertCache[key] = cache
		dataPlatformProductGroupProductGroupTextDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformProductGroupProductGroupTextDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformProductGroupProductGroupTextDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformProductGroupProductGroupTextDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformProductGroupProductGroupTextDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_product_group_product_group_text_data` WHERE `ProductGroup`=? AND `Language`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_product_group_product_group_text_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformProductGroupProductGroupTextDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformProductGroupProductGroupTextDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_product_group_product_group_text_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformProductGroupProductGroupTextDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductGroupProductGroupTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_product_group_product_group_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformProductGroupProductGroupTextDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformProductGroupProductGroupTextDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformProductGroupProductGroupTextDatum(ctx, exec, o.ProductGroup, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformProductGroupProductGroupTextDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformProductGroupProductGroupTextDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductGroupProductGroupTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_product_group_product_group_text_data`.* FROM `data_platform_product_group_product_group_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductGroupProductGroupTextDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformProductGroupProductGroupTextDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformProductGroupProductGroupTextDatumExists checks if the DataPlatformProductGroupProductGroupTextDatum row exists.
func DataPlatformProductGroupProductGroupTextDatumExists(ctx context.Context, exec boil.ContextExecutor, productGroup string, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_product_group_product_group_text_data` where `ProductGroup`=? AND `Language`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, productGroup, language)
	}
	row := exec.QueryRowContext(ctx, sql, productGroup, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_product_group_product_group_text_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformProductGroupProductGroupTextDatum row exists.
func (o *DataPlatformProductGroupProductGroupTextDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformProductGroupProductGroupTextDatumExists(ctx, exec, o.ProductGroup, o.Language)
}
