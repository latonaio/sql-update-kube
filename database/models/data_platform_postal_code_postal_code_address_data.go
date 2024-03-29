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

// DataPlatformPostalCodePostalCodeAddressDatum is an object representing the database table.
type DataPlatformPostalCodePostalCodeAddressDatum struct {
	PostalCode                  string      `boil:"PostalCode" json:"PostalCode" toml:"PostalCode" yaml:"PostalCode"`
	LocalRegion                 string      `boil:"LocalRegion" json:"LocalRegion" toml:"LocalRegion" yaml:"LocalRegion"`
	Country                     string      `boil:"Country" json:"Country" toml:"Country" yaml:"Country"`
	GlobalRegion                null.String `boil:"GlobalRegion" json:"GlobalRegion,omitempty" toml:"GlobalRegion" yaml:"GlobalRegion,omitempty"`
	PostalCodeAddressDetailText null.String `boil:"PostalCodeAddressDetailText" json:"PostalCodeAddressDetailText,omitempty" toml:"PostalCodeAddressDetailText" yaml:"PostalCodeAddressDetailText,omitempty"`
	CityName                    null.String `boil:"CityName" json:"CityName,omitempty" toml:"CityName" yaml:"CityName,omitempty"`
	Builiding                   null.String `boil:"Builiding" json:"Builiding,omitempty" toml:"Builiding" yaml:"Builiding,omitempty"`
	Floor                       null.Int    `boil:"Floor" json:"Floor,omitempty" toml:"Floor" yaml:"Floor,omitempty"`
	PostalCodeAddressTotalText  null.String `boil:"PostalCodeAddressTotalText" json:"PostalCodeAddressTotalText,omitempty" toml:"PostalCodeAddressTotalText" yaml:"PostalCodeAddressTotalText,omitempty"`

	R *dataPlatformPostalCodePostalCodeAddressDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformPostalCodePostalCodeAddressDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformPostalCodePostalCodeAddressDatumColumns = struct {
	PostalCode                  string
	LocalRegion                 string
	Country                     string
	GlobalRegion                string
	PostalCodeAddressDetailText string
	CityName                    string
	Builiding                   string
	Floor                       string
	PostalCodeAddressTotalText  string
}{
	PostalCode:                  "PostalCode",
	LocalRegion:                 "LocalRegion",
	Country:                     "Country",
	GlobalRegion:                "GlobalRegion",
	PostalCodeAddressDetailText: "PostalCodeAddressDetailText",
	CityName:                    "CityName",
	Builiding:                   "Builiding",
	Floor:                       "Floor",
	PostalCodeAddressTotalText:  "PostalCodeAddressTotalText",
}

var DataPlatformPostalCodePostalCodeAddressDatumTableColumns = struct {
	PostalCode                  string
	LocalRegion                 string
	Country                     string
	GlobalRegion                string
	PostalCodeAddressDetailText string
	CityName                    string
	Builiding                   string
	Floor                       string
	PostalCodeAddressTotalText  string
}{
	PostalCode:                  "data_platform_postal_code_postal_code_address_data.PostalCode",
	LocalRegion:                 "data_platform_postal_code_postal_code_address_data.LocalRegion",
	Country:                     "data_platform_postal_code_postal_code_address_data.Country",
	GlobalRegion:                "data_platform_postal_code_postal_code_address_data.GlobalRegion",
	PostalCodeAddressDetailText: "data_platform_postal_code_postal_code_address_data.PostalCodeAddressDetailText",
	CityName:                    "data_platform_postal_code_postal_code_address_data.CityName",
	Builiding:                   "data_platform_postal_code_postal_code_address_data.Builiding",
	Floor:                       "data_platform_postal_code_postal_code_address_data.Floor",
	PostalCodeAddressTotalText:  "data_platform_postal_code_postal_code_address_data.PostalCodeAddressTotalText",
}

// Generated where

var DataPlatformPostalCodePostalCodeAddressDatumWhere = struct {
	PostalCode                  whereHelperstring
	LocalRegion                 whereHelperstring
	Country                     whereHelperstring
	GlobalRegion                whereHelpernull_String
	PostalCodeAddressDetailText whereHelpernull_String
	CityName                    whereHelpernull_String
	Builiding                   whereHelpernull_String
	Floor                       whereHelpernull_Int
	PostalCodeAddressTotalText  whereHelpernull_String
}{
	PostalCode:                  whereHelperstring{field: "`data_platform_postal_code_postal_code_address_data`.`PostalCode`"},
	LocalRegion:                 whereHelperstring{field: "`data_platform_postal_code_postal_code_address_data`.`LocalRegion`"},
	Country:                     whereHelperstring{field: "`data_platform_postal_code_postal_code_address_data`.`Country`"},
	GlobalRegion:                whereHelpernull_String{field: "`data_platform_postal_code_postal_code_address_data`.`GlobalRegion`"},
	PostalCodeAddressDetailText: whereHelpernull_String{field: "`data_platform_postal_code_postal_code_address_data`.`PostalCodeAddressDetailText`"},
	CityName:                    whereHelpernull_String{field: "`data_platform_postal_code_postal_code_address_data`.`CityName`"},
	Builiding:                   whereHelpernull_String{field: "`data_platform_postal_code_postal_code_address_data`.`Builiding`"},
	Floor:                       whereHelpernull_Int{field: "`data_platform_postal_code_postal_code_address_data`.`Floor`"},
	PostalCodeAddressTotalText:  whereHelpernull_String{field: "`data_platform_postal_code_postal_code_address_data`.`PostalCodeAddressTotalText`"},
}

// DataPlatformPostalCodePostalCodeAddressDatumRels is where relationship names are stored.
var DataPlatformPostalCodePostalCodeAddressDatumRels = struct {
}{}

// dataPlatformPostalCodePostalCodeAddressDatumR is where relationships are stored.
type dataPlatformPostalCodePostalCodeAddressDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dataPlatformPostalCodePostalCodeAddressDatumR) NewStruct() *dataPlatformPostalCodePostalCodeAddressDatumR {
	return &dataPlatformPostalCodePostalCodeAddressDatumR{}
}

// dataPlatformPostalCodePostalCodeAddressDatumL is where Load methods for each relationship are stored.
type dataPlatformPostalCodePostalCodeAddressDatumL struct{}

var (
	dataPlatformPostalCodePostalCodeAddressDatumAllColumns            = []string{"PostalCode", "LocalRegion", "Country", "GlobalRegion", "PostalCodeAddressDetailText", "CityName", "Builiding", "Floor", "PostalCodeAddressTotalText"}
	dataPlatformPostalCodePostalCodeAddressDatumColumnsWithoutDefault = []string{"PostalCode", "LocalRegion", "Country", "GlobalRegion", "PostalCodeAddressDetailText", "CityName", "Builiding", "Floor", "PostalCodeAddressTotalText"}
	dataPlatformPostalCodePostalCodeAddressDatumColumnsWithDefault    = []string{}
	dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns     = []string{"PostalCode", "LocalRegion", "Country"}
	dataPlatformPostalCodePostalCodeAddressDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformPostalCodePostalCodeAddressDatumSlice is an alias for a slice of pointers to DataPlatformPostalCodePostalCodeAddressDatum.
	// This should almost always be used instead of []DataPlatformPostalCodePostalCodeAddressDatum.
	DataPlatformPostalCodePostalCodeAddressDatumSlice []*DataPlatformPostalCodePostalCodeAddressDatum

	dataPlatformPostalCodePostalCodeAddressDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformPostalCodePostalCodeAddressDatumType                 = reflect.TypeOf(&DataPlatformPostalCodePostalCodeAddressDatum{})
	dataPlatformPostalCodePostalCodeAddressDatumMapping              = queries.MakeStructMapping(dataPlatformPostalCodePostalCodeAddressDatumType)
	dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns)
	dataPlatformPostalCodePostalCodeAddressDatumInsertCacheMut       sync.RWMutex
	dataPlatformPostalCodePostalCodeAddressDatumInsertCache          = make(map[string]insertCache)
	dataPlatformPostalCodePostalCodeAddressDatumUpdateCacheMut       sync.RWMutex
	dataPlatformPostalCodePostalCodeAddressDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformPostalCodePostalCodeAddressDatumUpsertCacheMut       sync.RWMutex
	dataPlatformPostalCodePostalCodeAddressDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformPostalCodePostalCodeAddressDatum record from the query.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformPostalCodePostalCodeAddressDatum, error) {
	o := &DataPlatformPostalCodePostalCodeAddressDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_postal_code_postal_code_address_data")
	}

	return o, nil
}

// All returns all DataPlatformPostalCodePostalCodeAddressDatum records from the query.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformPostalCodePostalCodeAddressDatumSlice, error) {
	var o []*DataPlatformPostalCodePostalCodeAddressDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformPostalCodePostalCodeAddressDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformPostalCodePostalCodeAddressDatum records in the query.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_postal_code_postal_code_address_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_postal_code_postal_code_address_data exists")
	}

	return count > 0, nil
}

// DataPlatformPostalCodePostalCodeAddressData retrieves all the records using an executor.
func DataPlatformPostalCodePostalCodeAddressData(mods ...qm.QueryMod) dataPlatformPostalCodePostalCodeAddressDatumQuery {
	mods = append(mods, qm.From("`data_platform_postal_code_postal_code_address_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_postal_code_postal_code_address_data`.*"})
	}

	return dataPlatformPostalCodePostalCodeAddressDatumQuery{q}
}

// FindDataPlatformPostalCodePostalCodeAddressDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformPostalCodePostalCodeAddressDatum(ctx context.Context, exec boil.ContextExecutor, postalCode string, localRegion string, country string, selectCols ...string) (*DataPlatformPostalCodePostalCodeAddressDatum, error) {
	dataPlatformPostalCodePostalCodeAddressDatumObj := &DataPlatformPostalCodePostalCodeAddressDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_postal_code_postal_code_address_data` where `PostalCode`=? AND `LocalRegion`=? AND `Country`=?", sel,
	)

	q := queries.Raw(query, postalCode, localRegion, country)

	err := q.Bind(ctx, exec, dataPlatformPostalCodePostalCodeAddressDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_postal_code_postal_code_address_data")
	}

	return dataPlatformPostalCodePostalCodeAddressDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_postal_code_postal_code_address_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformPostalCodePostalCodeAddressDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformPostalCodePostalCodeAddressDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformPostalCodePostalCodeAddressDatumInsertCache[key]
	dataPlatformPostalCodePostalCodeAddressDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformPostalCodePostalCodeAddressDatumAllColumns,
			dataPlatformPostalCodePostalCodeAddressDatumColumnsWithDefault,
			dataPlatformPostalCodePostalCodeAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_postal_code_postal_code_address_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_postal_code_postal_code_address_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_postal_code_postal_code_address_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_postal_code_postal_code_address_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PostalCode,
		o.LocalRegion,
		o.Country,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_postal_code_postal_code_address_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformPostalCodePostalCodeAddressDatumInsertCacheMut.Lock()
		dataPlatformPostalCodePostalCodeAddressDatumInsertCache[key] = cache
		dataPlatformPostalCodePostalCodeAddressDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformPostalCodePostalCodeAddressDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformPostalCodePostalCodeAddressDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformPostalCodePostalCodeAddressDatumUpdateCache[key]
	dataPlatformPostalCodePostalCodeAddressDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformPostalCodePostalCodeAddressDatumAllColumns,
			dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_postal_code_postal_code_address_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_postal_code_postal_code_address_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, append(wl, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_postal_code_postal_code_address_data row")
	}

	if !cached {
		dataPlatformPostalCodePostalCodeAddressDatumUpdateCacheMut.Lock()
		dataPlatformPostalCodePostalCodeAddressDatumUpdateCache[key] = cache
		dataPlatformPostalCodePostalCodeAddressDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_postal_code_postal_code_address_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformPostalCodePostalCodeAddressDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_postal_code_postal_code_address_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformPostalCodePostalCodeAddressDatum slice")
	}

	return nil
}

var mySQLDataPlatformPostalCodePostalCodeAddressDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_postal_code_postal_code_address_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformPostalCodePostalCodeAddressDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformPostalCodePostalCodeAddressDatumUniqueColumns, o)

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

	dataPlatformPostalCodePostalCodeAddressDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformPostalCodePostalCodeAddressDatumUpsertCache[key]
	dataPlatformPostalCodePostalCodeAddressDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformPostalCodePostalCodeAddressDatumAllColumns,
			dataPlatformPostalCodePostalCodeAddressDatumColumnsWithDefault,
			dataPlatformPostalCodePostalCodeAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformPostalCodePostalCodeAddressDatumAllColumns,
			dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_postal_code_postal_code_address_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_postal_code_postal_code_address_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_postal_code_postal_code_address_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_postal_code_postal_code_address_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformPostalCodePostalCodeAddressDatumType, dataPlatformPostalCodePostalCodeAddressDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_postal_code_postal_code_address_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_postal_code_postal_code_address_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformPostalCodePostalCodeAddressDatumUpsertCacheMut.Lock()
		dataPlatformPostalCodePostalCodeAddressDatumUpsertCache[key] = cache
		dataPlatformPostalCodePostalCodeAddressDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformPostalCodePostalCodeAddressDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformPostalCodePostalCodeAddressDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_postal_code_postal_code_address_data` WHERE `PostalCode`=? AND `LocalRegion`=? AND `Country`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_postal_code_postal_code_address_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformPostalCodePostalCodeAddressDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformPostalCodePostalCodeAddressDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_postal_code_postal_code_address_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformPostalCodePostalCodeAddressDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_postal_code_postal_code_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformPostalCodePostalCodeAddressDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformPostalCodePostalCodeAddressDatum(ctx, exec, o.PostalCode, o.LocalRegion, o.Country)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformPostalCodePostalCodeAddressDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformPostalCodePostalCodeAddressDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_postal_code_postal_code_address_data`.* FROM `data_platform_postal_code_postal_code_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPostalCodePostalCodeAddressDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformPostalCodePostalCodeAddressDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformPostalCodePostalCodeAddressDatumExists checks if the DataPlatformPostalCodePostalCodeAddressDatum row exists.
func DataPlatformPostalCodePostalCodeAddressDatumExists(ctx context.Context, exec boil.ContextExecutor, postalCode string, localRegion string, country string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_postal_code_postal_code_address_data` where `PostalCode`=? AND `LocalRegion`=? AND `Country`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, postalCode, localRegion, country)
	}
	row := exec.QueryRowContext(ctx, sql, postalCode, localRegion, country)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_postal_code_postal_code_address_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformPostalCodePostalCodeAddressDatum row exists.
func (o *DataPlatformPostalCodePostalCodeAddressDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformPostalCodePostalCodeAddressDatumExists(ctx, exec, o.PostalCode, o.LocalRegion, o.Country)
}
