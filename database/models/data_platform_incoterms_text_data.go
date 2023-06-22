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

// DataPlatformIncotermsTextDatum is an object representing the database table.
type DataPlatformIncotermsTextDatum struct {
	Incoterms     string      `boil:"Incoterms" json:"Incoterms" toml:"Incoterms" yaml:"Incoterms"`
	Language      string      `boil:"Language" json:"Language" toml:"Language" yaml:"Language"`
	IncotermsName null.String `boil:"IncotermsName" json:"IncotermsName,omitempty" toml:"IncotermsName" yaml:"IncotermsName,omitempty"`

	R *dataPlatformIncotermsTextDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformIncotermsTextDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformIncotermsTextDatumColumns = struct {
	Incoterms     string
	Language      string
	IncotermsName string
}{
	Incoterms:     "Incoterms",
	Language:      "Language",
	IncotermsName: "IncotermsName",
}

var DataPlatformIncotermsTextDatumTableColumns = struct {
	Incoterms     string
	Language      string
	IncotermsName string
}{
	Incoterms:     "data_platform_incoterms_text_data.Incoterms",
	Language:      "data_platform_incoterms_text_data.Language",
	IncotermsName: "data_platform_incoterms_text_data.IncotermsName",
}

// Generated where

var DataPlatformIncotermsTextDatumWhere = struct {
	Incoterms     whereHelperstring
	Language      whereHelperstring
	IncotermsName whereHelpernull_String
}{
	Incoterms:     whereHelperstring{field: "`data_platform_incoterms_text_data`.`Incoterms`"},
	Language:      whereHelperstring{field: "`data_platform_incoterms_text_data`.`Language`"},
	IncotermsName: whereHelpernull_String{field: "`data_platform_incoterms_text_data`.`IncotermsName`"},
}

// DataPlatformIncotermsTextDatumRels is where relationship names are stored.
var DataPlatformIncotermsTextDatumRels = struct {
}{}

// dataPlatformIncotermsTextDatumR is where relationships are stored.
type dataPlatformIncotermsTextDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformIncotermsTextDatumR) NewStruct() *dataPlatformIncotermsTextDatumR {
	return &dataPlatformIncotermsTextDatumR{}
}

// dataPlatformIncotermsTextDatumL is where Load methods for each relationship are stored.
type dataPlatformIncotermsTextDatumL struct{}

var (
	dataPlatformIncotermsTextDatumAllColumns            = []string{"Incoterms", "Language", "IncotermsName"}
	dataPlatformIncotermsTextDatumColumnsWithoutDefault = []string{"Incoterms", "Language", "IncotermsName"}
	dataPlatformIncotermsTextDatumColumnsWithDefault    = []string{}
	dataPlatformIncotermsTextDatumPrimaryKeyColumns     = []string{"Incoterms", "Language"}
	dataPlatformIncotermsTextDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformIncotermsTextDatumSlice is an alias for a slice of pointers to DataPlatformIncotermsTextDatum.
	// This should almost always be used instead of []DataPlatformIncotermsTextDatum.
	DataPlatformIncotermsTextDatumSlice []*DataPlatformIncotermsTextDatum

	dataPlatformIncotermsTextDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformIncotermsTextDatumType                 = reflect.TypeOf(&DataPlatformIncotermsTextDatum{})
	dataPlatformIncotermsTextDatumMapping              = queries.MakeStructMapping(dataPlatformIncotermsTextDatumType)
	dataPlatformIncotermsTextDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, dataPlatformIncotermsTextDatumPrimaryKeyColumns)
	dataPlatformIncotermsTextDatumInsertCacheMut       sync.RWMutex
	dataPlatformIncotermsTextDatumInsertCache          = make(map[string]insertCache)
	dataPlatformIncotermsTextDatumUpdateCacheMut       sync.RWMutex
	dataPlatformIncotermsTextDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformIncotermsTextDatumUpsertCacheMut       sync.RWMutex
	dataPlatformIncotermsTextDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformIncotermsTextDatum record from the query.
func (q dataPlatformIncotermsTextDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformIncotermsTextDatum, error) {
	o := &DataPlatformIncotermsTextDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_incoterms_text_data")
	}

	return o, nil
}

// All returns all DataPlatformIncotermsTextDatum records from the query.
func (q dataPlatformIncotermsTextDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformIncotermsTextDatumSlice, error) {
	var o []*DataPlatformIncotermsTextDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformIncotermsTextDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformIncotermsTextDatum records in the query.
func (q dataPlatformIncotermsTextDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_incoterms_text_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformIncotermsTextDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_incoterms_text_data exists")
	}

	return count > 0, nil
}

// DataPlatformIncotermsTextData retrieves all the records using an executor.
func DataPlatformIncotermsTextData(mods ...qm.QueryMod) dataPlatformIncotermsTextDatumQuery {
	mods = append(mods, qm.From("`data_platform_incoterms_text_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_incoterms_text_data`.*"})
	}

	return dataPlatformIncotermsTextDatumQuery{q}
}

// FindDataPlatformIncotermsTextDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformIncotermsTextDatum(ctx context.Context, exec boil.ContextExecutor, incoterms string, language string, selectCols ...string) (*DataPlatformIncotermsTextDatum, error) {
	dataPlatformIncotermsTextDatumObj := &DataPlatformIncotermsTextDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_incoterms_text_data` where `Incoterms`=? AND `Language`=?", sel,
	)

	q := queries.Raw(query, incoterms, language)

	err := q.Bind(ctx, exec, dataPlatformIncotermsTextDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_incoterms_text_data")
	}

	return dataPlatformIncotermsTextDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformIncotermsTextDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_incoterms_text_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIncotermsTextDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformIncotermsTextDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformIncotermsTextDatumInsertCache[key]
	dataPlatformIncotermsTextDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformIncotermsTextDatumAllColumns,
			dataPlatformIncotermsTextDatumColumnsWithDefault,
			dataPlatformIncotermsTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_incoterms_text_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_incoterms_text_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_incoterms_text_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformIncotermsTextDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_incoterms_text_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Incoterms,
		o.Language,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_incoterms_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIncotermsTextDatumInsertCacheMut.Lock()
		dataPlatformIncotermsTextDatumInsertCache[key] = cache
		dataPlatformIncotermsTextDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformIncotermsTextDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformIncotermsTextDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformIncotermsTextDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformIncotermsTextDatumUpdateCache[key]
	dataPlatformIncotermsTextDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformIncotermsTextDatumAllColumns,
			dataPlatformIncotermsTextDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_incoterms_text_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_incoterms_text_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformIncotermsTextDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, append(wl, dataPlatformIncotermsTextDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_incoterms_text_data row")
	}

	if !cached {
		dataPlatformIncotermsTextDatumUpdateCacheMut.Lock()
		dataPlatformIncotermsTextDatumUpdateCache[key] = cache
		dataPlatformIncotermsTextDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformIncotermsTextDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_incoterms_text_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformIncotermsTextDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_incoterms_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsTextDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformIncotermsTextDatum slice")
	}

	return nil
}

var mySQLDataPlatformIncotermsTextDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformIncotermsTextDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_incoterms_text_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIncotermsTextDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformIncotermsTextDatumUniqueColumns, o)

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

	dataPlatformIncotermsTextDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformIncotermsTextDatumUpsertCache[key]
	dataPlatformIncotermsTextDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformIncotermsTextDatumAllColumns,
			dataPlatformIncotermsTextDatumColumnsWithDefault,
			dataPlatformIncotermsTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformIncotermsTextDatumAllColumns,
			dataPlatformIncotermsTextDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_incoterms_text_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_incoterms_text_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_incoterms_text_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_incoterms_text_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformIncotermsTextDatumType, dataPlatformIncotermsTextDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_incoterms_text_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_incoterms_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIncotermsTextDatumUpsertCacheMut.Lock()
		dataPlatformIncotermsTextDatumUpsertCache[key] = cache
		dataPlatformIncotermsTextDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformIncotermsTextDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformIncotermsTextDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformIncotermsTextDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformIncotermsTextDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_incoterms_text_data` WHERE `Incoterms`=? AND `Language`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_incoterms_text_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformIncotermsTextDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformIncotermsTextDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_incoterms_text_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformIncotermsTextDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_incoterms_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsTextDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformIncotermsTextDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformIncotermsTextDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformIncotermsTextDatum(ctx, exec, o.Incoterms, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformIncotermsTextDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformIncotermsTextDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_incoterms_text_data`.* FROM `data_platform_incoterms_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsTextDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformIncotermsTextDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformIncotermsTextDatumExists checks if the DataPlatformIncotermsTextDatum row exists.
func DataPlatformIncotermsTextDatumExists(ctx context.Context, exec boil.ContextExecutor, incoterms string, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_incoterms_text_data` where `Incoterms`=? AND `Language`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, incoterms, language)
	}
	row := exec.QueryRowContext(ctx, sql, incoterms, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_incoterms_text_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformIncotermsTextDatum row exists.
func (o *DataPlatformIncotermsTextDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformIncotermsTextDatumExists(ctx, exec, o.Incoterms, o.Language)
}
