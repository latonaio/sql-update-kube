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

// DataPlatformProductMasterBusinessPartnerDatum is an object representing the database table.
type DataPlatformProductMasterBusinessPartnerDatum struct {
	Product                string      `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
	BusinessPartner        int         `boil:"BusinessPartner" json:"BusinessPartner" toml:"BusinessPartner" yaml:"BusinessPartner"`
	ValidityEndDate        string      `boil:"ValidityEndDate" json:"ValidityEndDate" toml:"ValidityEndDate" yaml:"ValidityEndDate"`
	ValidityStartDate      string      `boil:"ValidityStartDate" json:"ValidityStartDate" toml:"ValidityStartDate" yaml:"ValidityStartDate"`
	BusinessPartnerProduct null.String `boil:"BusinessPartnerProduct" json:"BusinessPartnerProduct,omitempty" toml:"BusinessPartnerProduct" yaml:"BusinessPartnerProduct,omitempty"`
	IsMarkedForDeletion    null.Bool   `boil:"IsMarkedForDeletion" json:"IsMarkedForDeletion,omitempty" toml:"IsMarkedForDeletion" yaml:"IsMarkedForDeletion,omitempty"`

	R *dataPlatformProductMasterBusinessPartnerDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformProductMasterBusinessPartnerDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformProductMasterBusinessPartnerDatumColumns = struct {
	Product                string
	BusinessPartner        string
	ValidityEndDate        string
	ValidityStartDate      string
	BusinessPartnerProduct string
	IsMarkedForDeletion    string
}{
	Product:                "Product",
	BusinessPartner:        "BusinessPartner",
	ValidityEndDate:        "ValidityEndDate",
	ValidityStartDate:      "ValidityStartDate",
	BusinessPartnerProduct: "BusinessPartnerProduct",
	IsMarkedForDeletion:    "IsMarkedForDeletion",
}

var DataPlatformProductMasterBusinessPartnerDatumTableColumns = struct {
	Product                string
	BusinessPartner        string
	ValidityEndDate        string
	ValidityStartDate      string
	BusinessPartnerProduct string
	IsMarkedForDeletion    string
}{
	Product:                "data_platform_product_master_business_partner_data.Product",
	BusinessPartner:        "data_platform_product_master_business_partner_data.BusinessPartner",
	ValidityEndDate:        "data_platform_product_master_business_partner_data.ValidityEndDate",
	ValidityStartDate:      "data_platform_product_master_business_partner_data.ValidityStartDate",
	BusinessPartnerProduct: "data_platform_product_master_business_partner_data.BusinessPartnerProduct",
	IsMarkedForDeletion:    "data_platform_product_master_business_partner_data.IsMarkedForDeletion",
}

// Generated where

var DataPlatformProductMasterBusinessPartnerDatumWhere = struct {
	Product                whereHelperstring
	BusinessPartner        whereHelperint
	ValidityEndDate        whereHelperstring
	ValidityStartDate      whereHelperstring
	BusinessPartnerProduct whereHelpernull_String
	IsMarkedForDeletion    whereHelpernull_Bool
}{
	Product:                whereHelperstring{field: "`data_platform_product_master_business_partner_data`.`Product`"},
	BusinessPartner:        whereHelperint{field: "`data_platform_product_master_business_partner_data`.`BusinessPartner`"},
	ValidityEndDate:        whereHelperstring{field: "`data_platform_product_master_business_partner_data`.`ValidityEndDate`"},
	ValidityStartDate:      whereHelperstring{field: "`data_platform_product_master_business_partner_data`.`ValidityStartDate`"},
	BusinessPartnerProduct: whereHelpernull_String{field: "`data_platform_product_master_business_partner_data`.`BusinessPartnerProduct`"},
	IsMarkedForDeletion:    whereHelpernull_Bool{field: "`data_platform_product_master_business_partner_data`.`IsMarkedForDeletion`"},
}

// DataPlatformProductMasterBusinessPartnerDatumRels is where relationship names are stored.
var DataPlatformProductMasterBusinessPartnerDatumRels = struct {
}{}

// dataPlatformProductMasterBusinessPartnerDatumR is where relationships are stored.
type dataPlatformProductMasterBusinessPartnerDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformProductMasterBusinessPartnerDatumR) NewStruct() *dataPlatformProductMasterBusinessPartnerDatumR {
	return &dataPlatformProductMasterBusinessPartnerDatumR{}
}

// dataPlatformProductMasterBusinessPartnerDatumL is where Load methods for each relationship are stored.
type dataPlatformProductMasterBusinessPartnerDatumL struct{}

var (
	dataPlatformProductMasterBusinessPartnerDatumAllColumns            = []string{"Product", "BusinessPartner", "ValidityEndDate", "ValidityStartDate", "BusinessPartnerProduct", "IsMarkedForDeletion"}
	dataPlatformProductMasterBusinessPartnerDatumColumnsWithoutDefault = []string{"Product", "BusinessPartner", "ValidityEndDate", "ValidityStartDate", "BusinessPartnerProduct", "IsMarkedForDeletion"}
	dataPlatformProductMasterBusinessPartnerDatumColumnsWithDefault    = []string{}
	dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns     = []string{"Product", "BusinessPartner", "ValidityEndDate", "ValidityStartDate"}
	dataPlatformProductMasterBusinessPartnerDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformProductMasterBusinessPartnerDatumSlice is an alias for a slice of pointers to DataPlatformProductMasterBusinessPartnerDatum.
	// This should almost always be used instead of []DataPlatformProductMasterBusinessPartnerDatum.
	DataPlatformProductMasterBusinessPartnerDatumSlice []*DataPlatformProductMasterBusinessPartnerDatum

	dataPlatformProductMasterBusinessPartnerDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformProductMasterBusinessPartnerDatumType                 = reflect.TypeOf(&DataPlatformProductMasterBusinessPartnerDatum{})
	dataPlatformProductMasterBusinessPartnerDatumMapping              = queries.MakeStructMapping(dataPlatformProductMasterBusinessPartnerDatumType)
	dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns)
	dataPlatformProductMasterBusinessPartnerDatumInsertCacheMut       sync.RWMutex
	dataPlatformProductMasterBusinessPartnerDatumInsertCache          = make(map[string]insertCache)
	dataPlatformProductMasterBusinessPartnerDatumUpdateCacheMut       sync.RWMutex
	dataPlatformProductMasterBusinessPartnerDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformProductMasterBusinessPartnerDatumUpsertCacheMut       sync.RWMutex
	dataPlatformProductMasterBusinessPartnerDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformProductMasterBusinessPartnerDatum record from the query.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformProductMasterBusinessPartnerDatum, error) {
	o := &DataPlatformProductMasterBusinessPartnerDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_product_master_business_partner_data")
	}

	return o, nil
}

// All returns all DataPlatformProductMasterBusinessPartnerDatum records from the query.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformProductMasterBusinessPartnerDatumSlice, error) {
	var o []*DataPlatformProductMasterBusinessPartnerDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformProductMasterBusinessPartnerDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformProductMasterBusinessPartnerDatum records in the query.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_product_master_business_partner_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_product_master_business_partner_data exists")
	}

	return count > 0, nil
}

// DataPlatformProductMasterBusinessPartnerData retrieves all the records using an executor.
func DataPlatformProductMasterBusinessPartnerData(mods ...qm.QueryMod) dataPlatformProductMasterBusinessPartnerDatumQuery {
	mods = append(mods, qm.From("`data_platform_product_master_business_partner_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_product_master_business_partner_data`.*"})
	}

	return dataPlatformProductMasterBusinessPartnerDatumQuery{q}
}

// FindDataPlatformProductMasterBusinessPartnerDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformProductMasterBusinessPartnerDatum(ctx context.Context, exec boil.ContextExecutor, product string, businessPartner int, validityEndDate string, validityStartDate string, selectCols ...string) (*DataPlatformProductMasterBusinessPartnerDatum, error) {
	dataPlatformProductMasterBusinessPartnerDatumObj := &DataPlatformProductMasterBusinessPartnerDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_product_master_business_partner_data` where `Product`=? AND `BusinessPartner`=? AND `ValidityEndDate`=? AND `ValidityStartDate`=?", sel,
	)

	q := queries.Raw(query, product, businessPartner, validityEndDate, validityStartDate)

	err := q.Bind(ctx, exec, dataPlatformProductMasterBusinessPartnerDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_product_master_business_partner_data")
	}

	return dataPlatformProductMasterBusinessPartnerDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_master_business_partner_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductMasterBusinessPartnerDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformProductMasterBusinessPartnerDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformProductMasterBusinessPartnerDatumInsertCache[key]
	dataPlatformProductMasterBusinessPartnerDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformProductMasterBusinessPartnerDatumAllColumns,
			dataPlatformProductMasterBusinessPartnerDatumColumnsWithDefault,
			dataPlatformProductMasterBusinessPartnerDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_product_master_business_partner_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_product_master_business_partner_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_product_master_business_partner_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_product_master_business_partner_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Product,
		o.BusinessPartner,
		o.ValidityEndDate,
		o.ValidityStartDate,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_master_business_partner_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductMasterBusinessPartnerDatumInsertCacheMut.Lock()
		dataPlatformProductMasterBusinessPartnerDatumInsertCache[key] = cache
		dataPlatformProductMasterBusinessPartnerDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformProductMasterBusinessPartnerDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformProductMasterBusinessPartnerDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformProductMasterBusinessPartnerDatumUpdateCache[key]
	dataPlatformProductMasterBusinessPartnerDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformProductMasterBusinessPartnerDatumAllColumns,
			dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_product_master_business_partner_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_product_master_business_partner_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, append(wl, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_product_master_business_partner_data row")
	}

	if !cached {
		dataPlatformProductMasterBusinessPartnerDatumUpdateCacheMut.Lock()
		dataPlatformProductMasterBusinessPartnerDatumUpdateCache[key] = cache
		dataPlatformProductMasterBusinessPartnerDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_product_master_business_partner_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformProductMasterBusinessPartnerDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_product_master_business_partner_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformProductMasterBusinessPartnerDatum slice")
	}

	return nil
}

var mySQLDataPlatformProductMasterBusinessPartnerDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_master_business_partner_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductMasterBusinessPartnerDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformProductMasterBusinessPartnerDatumUniqueColumns, o)

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

	dataPlatformProductMasterBusinessPartnerDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformProductMasterBusinessPartnerDatumUpsertCache[key]
	dataPlatformProductMasterBusinessPartnerDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformProductMasterBusinessPartnerDatumAllColumns,
			dataPlatformProductMasterBusinessPartnerDatumColumnsWithDefault,
			dataPlatformProductMasterBusinessPartnerDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformProductMasterBusinessPartnerDatumAllColumns,
			dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_product_master_business_partner_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_product_master_business_partner_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_product_master_business_partner_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_product_master_business_partner_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformProductMasterBusinessPartnerDatumType, dataPlatformProductMasterBusinessPartnerDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_product_master_business_partner_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_master_business_partner_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductMasterBusinessPartnerDatumUpsertCacheMut.Lock()
		dataPlatformProductMasterBusinessPartnerDatumUpsertCache[key] = cache
		dataPlatformProductMasterBusinessPartnerDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformProductMasterBusinessPartnerDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformProductMasterBusinessPartnerDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_product_master_business_partner_data` WHERE `Product`=? AND `BusinessPartner`=? AND `ValidityEndDate`=? AND `ValidityStartDate`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_product_master_business_partner_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformProductMasterBusinessPartnerDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformProductMasterBusinessPartnerDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_product_master_business_partner_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformProductMasterBusinessPartnerDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_product_master_business_partner_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformProductMasterBusinessPartnerDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformProductMasterBusinessPartnerDatum(ctx, exec, o.Product, o.BusinessPartner, o.ValidityEndDate, o.ValidityStartDate)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformProductMasterBusinessPartnerDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformProductMasterBusinessPartnerDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_product_master_business_partner_data`.* FROM `data_platform_product_master_business_partner_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductMasterBusinessPartnerDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformProductMasterBusinessPartnerDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformProductMasterBusinessPartnerDatumExists checks if the DataPlatformProductMasterBusinessPartnerDatum row exists.
func DataPlatformProductMasterBusinessPartnerDatumExists(ctx context.Context, exec boil.ContextExecutor, product string, businessPartner int, validityEndDate string, validityStartDate string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_product_master_business_partner_data` where `Product`=? AND `BusinessPartner`=? AND `ValidityEndDate`=? AND `ValidityStartDate`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, product, businessPartner, validityEndDate, validityStartDate)
	}
	row := exec.QueryRowContext(ctx, sql, product, businessPartner, validityEndDate, validityStartDate)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_product_master_business_partner_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformProductMasterBusinessPartnerDatum row exists.
func (o *DataPlatformProductMasterBusinessPartnerDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformProductMasterBusinessPartnerDatumExists(ctx, exec, o.Product, o.BusinessPartner, o.ValidityEndDate, o.ValidityStartDate)
}
