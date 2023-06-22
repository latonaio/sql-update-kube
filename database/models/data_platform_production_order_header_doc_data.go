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

// DataPlatformProductionOrderHeaderDocDatum is an object representing the database table.
type DataPlatformProductionOrderHeaderDocDatum struct {
	ProductionOrder          int         `boil:"ProductionOrder" json:"ProductionOrder" toml:"ProductionOrder" yaml:"ProductionOrder"`
	DocType                  string      `boil:"DocType" json:"DocType" toml:"DocType" yaml:"DocType"`
	DocVersionID             int         `boil:"DocVersionID" json:"DocVersionID" toml:"DocVersionID" yaml:"DocVersionID"`
	DocID                    string      `boil:"DocID" json:"DocID" toml:"DocID" yaml:"DocID"`
	FileExtension            string      `boil:"FileExtension" json:"FileExtension" toml:"FileExtension" yaml:"FileExtension"`
	FileName                 null.String `boil:"FileName" json:"FileName,omitempty" toml:"FileName" yaml:"FileName,omitempty"`
	FilePath                 null.String `boil:"FilePath" json:"FilePath,omitempty" toml:"FilePath" yaml:"FilePath,omitempty"`
	DocIssuerBusinessPartner null.Int    `boil:"DocIssuerBusinessPartner" json:"DocIssuerBusinessPartner,omitempty" toml:"DocIssuerBusinessPartner" yaml:"DocIssuerBusinessPartner,omitempty"`

	R *dataPlatformProductionOrderHeaderDocDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformProductionOrderHeaderDocDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformProductionOrderHeaderDocDatumColumns = struct {
	ProductionOrder          string
	DocType                  string
	DocVersionID             string
	DocID                    string
	FileExtension            string
	FileName                 string
	FilePath                 string
	DocIssuerBusinessPartner string
}{
	ProductionOrder:          "ProductionOrder",
	DocType:                  "DocType",
	DocVersionID:             "DocVersionID",
	DocID:                    "DocID",
	FileExtension:            "FileExtension",
	FileName:                 "FileName",
	FilePath:                 "FilePath",
	DocIssuerBusinessPartner: "DocIssuerBusinessPartner",
}

var DataPlatformProductionOrderHeaderDocDatumTableColumns = struct {
	ProductionOrder          string
	DocType                  string
	DocVersionID             string
	DocID                    string
	FileExtension            string
	FileName                 string
	FilePath                 string
	DocIssuerBusinessPartner string
}{
	ProductionOrder:          "data_platform_production_order_header_doc_data.ProductionOrder",
	DocType:                  "data_platform_production_order_header_doc_data.DocType",
	DocVersionID:             "data_platform_production_order_header_doc_data.DocVersionID",
	DocID:                    "data_platform_production_order_header_doc_data.DocID",
	FileExtension:            "data_platform_production_order_header_doc_data.FileExtension",
	FileName:                 "data_platform_production_order_header_doc_data.FileName",
	FilePath:                 "data_platform_production_order_header_doc_data.FilePath",
	DocIssuerBusinessPartner: "data_platform_production_order_header_doc_data.DocIssuerBusinessPartner",
}

// Generated where

var DataPlatformProductionOrderHeaderDocDatumWhere = struct {
	ProductionOrder          whereHelperint
	DocType                  whereHelperstring
	DocVersionID             whereHelperint
	DocID                    whereHelperstring
	FileExtension            whereHelperstring
	FileName                 whereHelpernull_String
	FilePath                 whereHelpernull_String
	DocIssuerBusinessPartner whereHelpernull_Int
}{
	ProductionOrder:          whereHelperint{field: "`data_platform_production_order_header_doc_data`.`ProductionOrder`"},
	DocType:                  whereHelperstring{field: "`data_platform_production_order_header_doc_data`.`DocType`"},
	DocVersionID:             whereHelperint{field: "`data_platform_production_order_header_doc_data`.`DocVersionID`"},
	DocID:                    whereHelperstring{field: "`data_platform_production_order_header_doc_data`.`DocID`"},
	FileExtension:            whereHelperstring{field: "`data_platform_production_order_header_doc_data`.`FileExtension`"},
	FileName:                 whereHelpernull_String{field: "`data_platform_production_order_header_doc_data`.`FileName`"},
	FilePath:                 whereHelpernull_String{field: "`data_platform_production_order_header_doc_data`.`FilePath`"},
	DocIssuerBusinessPartner: whereHelpernull_Int{field: "`data_platform_production_order_header_doc_data`.`DocIssuerBusinessPartner`"},
}

// DataPlatformProductionOrderHeaderDocDatumRels is where relationship names are stored.
var DataPlatformProductionOrderHeaderDocDatumRels = struct {
}{}

// dataPlatformProductionOrderHeaderDocDatumR is where relationships are stored.
type dataPlatformProductionOrderHeaderDocDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformProductionOrderHeaderDocDatumR) NewStruct() *dataPlatformProductionOrderHeaderDocDatumR {
	return &dataPlatformProductionOrderHeaderDocDatumR{}
}

// dataPlatformProductionOrderHeaderDocDatumL is where Load methods for each relationship are stored.
type dataPlatformProductionOrderHeaderDocDatumL struct{}

var (
	dataPlatformProductionOrderHeaderDocDatumAllColumns            = []string{"ProductionOrder", "DocType", "DocVersionID", "DocID", "FileExtension", "FileName", "FilePath", "DocIssuerBusinessPartner"}
	dataPlatformProductionOrderHeaderDocDatumColumnsWithoutDefault = []string{"ProductionOrder", "DocType", "DocVersionID", "DocID", "FileExtension", "FileName", "FilePath", "DocIssuerBusinessPartner"}
	dataPlatformProductionOrderHeaderDocDatumColumnsWithDefault    = []string{}
	dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns     = []string{"ProductionOrder", "DocType", "DocVersionID", "DocID"}
	dataPlatformProductionOrderHeaderDocDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformProductionOrderHeaderDocDatumSlice is an alias for a slice of pointers to DataPlatformProductionOrderHeaderDocDatum.
	// This should almost always be used instead of []DataPlatformProductionOrderHeaderDocDatum.
	DataPlatformProductionOrderHeaderDocDatumSlice []*DataPlatformProductionOrderHeaderDocDatum

	dataPlatformProductionOrderHeaderDocDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformProductionOrderHeaderDocDatumType                 = reflect.TypeOf(&DataPlatformProductionOrderHeaderDocDatum{})
	dataPlatformProductionOrderHeaderDocDatumMapping              = queries.MakeStructMapping(dataPlatformProductionOrderHeaderDocDatumType)
	dataPlatformProductionOrderHeaderDocDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns)
	dataPlatformProductionOrderHeaderDocDatumInsertCacheMut       sync.RWMutex
	dataPlatformProductionOrderHeaderDocDatumInsertCache          = make(map[string]insertCache)
	dataPlatformProductionOrderHeaderDocDatumUpdateCacheMut       sync.RWMutex
	dataPlatformProductionOrderHeaderDocDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformProductionOrderHeaderDocDatumUpsertCacheMut       sync.RWMutex
	dataPlatformProductionOrderHeaderDocDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformProductionOrderHeaderDocDatum record from the query.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformProductionOrderHeaderDocDatum, error) {
	o := &DataPlatformProductionOrderHeaderDocDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_production_order_header_doc_data")
	}

	return o, nil
}

// All returns all DataPlatformProductionOrderHeaderDocDatum records from the query.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformProductionOrderHeaderDocDatumSlice, error) {
	var o []*DataPlatformProductionOrderHeaderDocDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformProductionOrderHeaderDocDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformProductionOrderHeaderDocDatum records in the query.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_production_order_header_doc_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_production_order_header_doc_data exists")
	}

	return count > 0, nil
}

// DataPlatformProductionOrderHeaderDocData retrieves all the records using an executor.
func DataPlatformProductionOrderHeaderDocData(mods ...qm.QueryMod) dataPlatformProductionOrderHeaderDocDatumQuery {
	mods = append(mods, qm.From("`data_platform_production_order_header_doc_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_production_order_header_doc_data`.*"})
	}

	return dataPlatformProductionOrderHeaderDocDatumQuery{q}
}

// FindDataPlatformProductionOrderHeaderDocDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformProductionOrderHeaderDocDatum(ctx context.Context, exec boil.ContextExecutor, productionOrder int, docType string, docVersionID int, docID string, selectCols ...string) (*DataPlatformProductionOrderHeaderDocDatum, error) {
	dataPlatformProductionOrderHeaderDocDatumObj := &DataPlatformProductionOrderHeaderDocDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_production_order_header_doc_data` where `ProductionOrder`=? AND `DocType`=? AND `DocVersionID`=? AND `DocID`=?", sel,
	)

	q := queries.Raw(query, productionOrder, docType, docVersionID, docID)

	err := q.Bind(ctx, exec, dataPlatformProductionOrderHeaderDocDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_production_order_header_doc_data")
	}

	return dataPlatformProductionOrderHeaderDocDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformProductionOrderHeaderDocDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_production_order_header_doc_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductionOrderHeaderDocDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformProductionOrderHeaderDocDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformProductionOrderHeaderDocDatumInsertCache[key]
	dataPlatformProductionOrderHeaderDocDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformProductionOrderHeaderDocDatumAllColumns,
			dataPlatformProductionOrderHeaderDocDatumColumnsWithDefault,
			dataPlatformProductionOrderHeaderDocDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_production_order_header_doc_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_production_order_header_doc_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_production_order_header_doc_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_production_order_header_doc_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ProductionOrder,
		o.DocType,
		o.DocVersionID,
		o.DocID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_production_order_header_doc_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductionOrderHeaderDocDatumInsertCacheMut.Lock()
		dataPlatformProductionOrderHeaderDocDatumInsertCache[key] = cache
		dataPlatformProductionOrderHeaderDocDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformProductionOrderHeaderDocDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformProductionOrderHeaderDocDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformProductionOrderHeaderDocDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformProductionOrderHeaderDocDatumUpdateCache[key]
	dataPlatformProductionOrderHeaderDocDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformProductionOrderHeaderDocDatumAllColumns,
			dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_production_order_header_doc_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_production_order_header_doc_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, append(wl, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_production_order_header_doc_data row")
	}

	if !cached {
		dataPlatformProductionOrderHeaderDocDatumUpdateCacheMut.Lock()
		dataPlatformProductionOrderHeaderDocDatumUpdateCache[key] = cache
		dataPlatformProductionOrderHeaderDocDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_production_order_header_doc_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformProductionOrderHeaderDocDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductionOrderHeaderDocDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_production_order_header_doc_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformProductionOrderHeaderDocDatum slice")
	}

	return nil
}

var mySQLDataPlatformProductionOrderHeaderDocDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformProductionOrderHeaderDocDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_production_order_header_doc_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductionOrderHeaderDocDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformProductionOrderHeaderDocDatumUniqueColumns, o)

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

	dataPlatformProductionOrderHeaderDocDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformProductionOrderHeaderDocDatumUpsertCache[key]
	dataPlatformProductionOrderHeaderDocDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformProductionOrderHeaderDocDatumAllColumns,
			dataPlatformProductionOrderHeaderDocDatumColumnsWithDefault,
			dataPlatformProductionOrderHeaderDocDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformProductionOrderHeaderDocDatumAllColumns,
			dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_production_order_header_doc_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_production_order_header_doc_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_production_order_header_doc_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_production_order_header_doc_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformProductionOrderHeaderDocDatumType, dataPlatformProductionOrderHeaderDocDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_production_order_header_doc_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_production_order_header_doc_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductionOrderHeaderDocDatumUpsertCacheMut.Lock()
		dataPlatformProductionOrderHeaderDocDatumUpsertCache[key] = cache
		dataPlatformProductionOrderHeaderDocDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformProductionOrderHeaderDocDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformProductionOrderHeaderDocDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformProductionOrderHeaderDocDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformProductionOrderHeaderDocDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_production_order_header_doc_data` WHERE `ProductionOrder`=? AND `DocType`=? AND `DocVersionID`=? AND `DocID`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_production_order_header_doc_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformProductionOrderHeaderDocDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformProductionOrderHeaderDocDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_production_order_header_doc_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformProductionOrderHeaderDocDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductionOrderHeaderDocDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_production_order_header_doc_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformProductionOrderHeaderDocDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformProductionOrderHeaderDocDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformProductionOrderHeaderDocDatum(ctx, exec, o.ProductionOrder, o.DocType, o.DocVersionID, o.DocID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformProductionOrderHeaderDocDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformProductionOrderHeaderDocDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductionOrderHeaderDocDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_production_order_header_doc_data`.* FROM `data_platform_production_order_header_doc_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductionOrderHeaderDocDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformProductionOrderHeaderDocDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformProductionOrderHeaderDocDatumExists checks if the DataPlatformProductionOrderHeaderDocDatum row exists.
func DataPlatformProductionOrderHeaderDocDatumExists(ctx context.Context, exec boil.ContextExecutor, productionOrder int, docType string, docVersionID int, docID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_production_order_header_doc_data` where `ProductionOrder`=? AND `DocType`=? AND `DocVersionID`=? AND `DocID`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, productionOrder, docType, docVersionID, docID)
	}
	row := exec.QueryRowContext(ctx, sql, productionOrder, docType, docVersionID, docID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_production_order_header_doc_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformProductionOrderHeaderDocDatum row exists.
func (o *DataPlatformProductionOrderHeaderDocDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformProductionOrderHeaderDocDatumExists(ctx, exec, o.ProductionOrder, o.DocType, o.DocVersionID, o.DocID)
}
