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

// DataPlatformEquipmentTypeEquipmentTypeDatum is an object representing the database table.
type DataPlatformEquipmentTypeEquipmentTypeDatum struct {
	EquipmentType string `boil:"EquipmentType" json:"EquipmentType" toml:"EquipmentType" yaml:"EquipmentType"`

	R *dataPlatformEquipmentTypeEquipmentTypeDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformEquipmentTypeEquipmentTypeDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformEquipmentTypeEquipmentTypeDatumColumns = struct {
	EquipmentType string
}{
	EquipmentType: "EquipmentType",
}

var DataPlatformEquipmentTypeEquipmentTypeDatumTableColumns = struct {
	EquipmentType string
}{
	EquipmentType: "data_platform_equipment_type_equipment_type_data.EquipmentType",
}

// Generated where

var DataPlatformEquipmentTypeEquipmentTypeDatumWhere = struct {
	EquipmentType whereHelperstring
}{
	EquipmentType: whereHelperstring{field: "`data_platform_equipment_type_equipment_type_data`.`EquipmentType`"},
}

// DataPlatformEquipmentTypeEquipmentTypeDatumRels is where relationship names are stored.
var DataPlatformEquipmentTypeEquipmentTypeDatumRels = struct {
	EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData string
}{
	EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData: "EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData",
}

// dataPlatformEquipmentTypeEquipmentTypeDatumR is where relationships are stored.
type dataPlatformEquipmentTypeEquipmentTypeDatumR struct {
	EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData DataPlatformEquipmentTypeEquipmentTypeTextDatumSlice `boil:"EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData" json:"EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData" toml:"EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData" yaml:"EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformEquipmentTypeEquipmentTypeDatumR) NewStruct() *dataPlatformEquipmentTypeEquipmentTypeDatumR {
	return &dataPlatformEquipmentTypeEquipmentTypeDatumR{}
}

func (r *dataPlatformEquipmentTypeEquipmentTypeDatumR) GetEquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData() DataPlatformEquipmentTypeEquipmentTypeTextDatumSlice {
	if r == nil {
		return nil
	}
	return r.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData
}

// dataPlatformEquipmentTypeEquipmentTypeDatumL is where Load methods for each relationship are stored.
type dataPlatformEquipmentTypeEquipmentTypeDatumL struct{}

var (
	dataPlatformEquipmentTypeEquipmentTypeDatumAllColumns            = []string{"EquipmentType"}
	dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithoutDefault = []string{"EquipmentType"}
	dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithDefault    = []string{}
	dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns     = []string{"EquipmentType"}
	dataPlatformEquipmentTypeEquipmentTypeDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformEquipmentTypeEquipmentTypeDatumSlice is an alias for a slice of pointers to DataPlatformEquipmentTypeEquipmentTypeDatum.
	// This should almost always be used instead of []DataPlatformEquipmentTypeEquipmentTypeDatum.
	DataPlatformEquipmentTypeEquipmentTypeDatumSlice []*DataPlatformEquipmentTypeEquipmentTypeDatum

	dataPlatformEquipmentTypeEquipmentTypeDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformEquipmentTypeEquipmentTypeDatumType                 = reflect.TypeOf(&DataPlatformEquipmentTypeEquipmentTypeDatum{})
	dataPlatformEquipmentTypeEquipmentTypeDatumMapping              = queries.MakeStructMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType)
	dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns)
	dataPlatformEquipmentTypeEquipmentTypeDatumInsertCacheMut       sync.RWMutex
	dataPlatformEquipmentTypeEquipmentTypeDatumInsertCache          = make(map[string]insertCache)
	dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCacheMut       sync.RWMutex
	dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCacheMut       sync.RWMutex
	dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformEquipmentTypeEquipmentTypeDatum record from the query.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformEquipmentTypeEquipmentTypeDatum, error) {
	o := &DataPlatformEquipmentTypeEquipmentTypeDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_equipment_type_equipment_type_data")
	}

	return o, nil
}

// All returns all DataPlatformEquipmentTypeEquipmentTypeDatum records from the query.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformEquipmentTypeEquipmentTypeDatumSlice, error) {
	var o []*DataPlatformEquipmentTypeEquipmentTypeDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformEquipmentTypeEquipmentTypeDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformEquipmentTypeEquipmentTypeDatum records in the query.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_equipment_type_equipment_type_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_equipment_type_equipment_type_data exists")
	}

	return count > 0, nil
}

// EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData retrieves all the data_platform_equipment_type_equipment_type_text_datum's DataPlatformEquipmentTypeEquipmentTypeTextData with an executor via EquipmentType column.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData(mods ...qm.QueryMod) dataPlatformEquipmentTypeEquipmentTypeTextDatumQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`data_platform_equipment_type_equipment_type_text_data`.`EquipmentType`=?", o.EquipmentType),
	)

	return DataPlatformEquipmentTypeEquipmentTypeTextData(queryMods...)
}

// LoadEquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (dataPlatformEquipmentTypeEquipmentTypeDatumL) LoadEquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformEquipmentTypeEquipmentTypeDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformEquipmentTypeEquipmentTypeDatum
	var object *DataPlatformEquipmentTypeEquipmentTypeDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformEquipmentTypeEquipmentTypeDatum.(*DataPlatformEquipmentTypeEquipmentTypeDatum)
		if !ok {
			object = new(DataPlatformEquipmentTypeEquipmentTypeDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformEquipmentTypeEquipmentTypeDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformEquipmentTypeEquipmentTypeDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformEquipmentTypeEquipmentTypeDatum.(*[]*DataPlatformEquipmentTypeEquipmentTypeDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformEquipmentTypeEquipmentTypeDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformEquipmentTypeEquipmentTypeDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformEquipmentTypeEquipmentTypeDatumR{}
		}
		args = append(args, object.EquipmentType)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformEquipmentTypeEquipmentTypeDatumR{}
			}

			for _, a := range args {
				if a == obj.EquipmentType {
					continue Outer
				}
			}

			args = append(args, obj.EquipmentType)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_equipment_type_equipment_type_text_data`),
		qm.WhereIn(`data_platform_equipment_type_equipment_type_text_data.EquipmentType in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load data_platform_equipment_type_equipment_type_text_data")
	}

	var resultSlice []*DataPlatformEquipmentTypeEquipmentTypeTextDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice data_platform_equipment_type_equipment_type_text_data")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on data_platform_equipment_type_equipment_type_text_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_equipment_type_equipment_type_text_data")
	}

	if singular {
		object.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EquipmentType == foreign.EquipmentType {
				local.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData = append(local.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData, foreign)
				break
			}
		}
	}

	return nil
}

// AddEquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData adds the given related objects to the existing relationships
// of the data_platform_equipment_type_equipment_type_datum, optionally inserting them as new records.
// Appends related to o.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) AddEquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DataPlatformEquipmentTypeEquipmentTypeTextDatum) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EquipmentType = o.EquipmentType
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `data_platform_equipment_type_equipment_type_text_data` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"EquipmentType"}),
				strmangle.WhereClause("`", "`", 0, dataPlatformEquipmentTypeEquipmentTypeTextDatumPrimaryKeyColumns),
			)
			values := []interface{}{o.EquipmentType, rel.EquipmentType, rel.Language}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EquipmentType = o.EquipmentType
		}
	}

	if o.R == nil {
		o.R = &dataPlatformEquipmentTypeEquipmentTypeDatumR{
			EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData: related,
		}
	} else {
		o.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData = append(o.R.EquipmentTypeDataPlatformEquipmentTypeEquipmentTypeTextData, related...)
	}

	return nil
}

// DataPlatformEquipmentTypeEquipmentTypeData retrieves all the records using an executor.
func DataPlatformEquipmentTypeEquipmentTypeData(mods ...qm.QueryMod) dataPlatformEquipmentTypeEquipmentTypeDatumQuery {
	mods = append(mods, qm.From("`data_platform_equipment_type_equipment_type_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_equipment_type_equipment_type_data`.*"})
	}

	return dataPlatformEquipmentTypeEquipmentTypeDatumQuery{q}
}

// FindDataPlatformEquipmentTypeEquipmentTypeDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformEquipmentTypeEquipmentTypeDatum(ctx context.Context, exec boil.ContextExecutor, equipmentType string, selectCols ...string) (*DataPlatformEquipmentTypeEquipmentTypeDatum, error) {
	dataPlatformEquipmentTypeEquipmentTypeDatumObj := &DataPlatformEquipmentTypeEquipmentTypeDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_equipment_type_equipment_type_data` where `EquipmentType`=?", sel,
	)

	q := queries.Raw(query, equipmentType)

	err := q.Bind(ctx, exec, dataPlatformEquipmentTypeEquipmentTypeDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_equipment_type_equipment_type_data")
	}

	return dataPlatformEquipmentTypeEquipmentTypeDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_equipment_type_equipment_type_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformEquipmentTypeEquipmentTypeDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformEquipmentTypeEquipmentTypeDatumInsertCache[key]
	dataPlatformEquipmentTypeEquipmentTypeDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformEquipmentTypeEquipmentTypeDatumAllColumns,
			dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithDefault,
			dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_equipment_type_equipment_type_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_equipment_type_equipment_type_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_equipment_type_equipment_type_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_equipment_type_equipment_type_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.EquipmentType,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_equipment_type_equipment_type_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformEquipmentTypeEquipmentTypeDatumInsertCacheMut.Lock()
		dataPlatformEquipmentTypeEquipmentTypeDatumInsertCache[key] = cache
		dataPlatformEquipmentTypeEquipmentTypeDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformEquipmentTypeEquipmentTypeDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCache[key]
	dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformEquipmentTypeEquipmentTypeDatumAllColumns,
			dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_equipment_type_equipment_type_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_equipment_type_equipment_type_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, append(wl, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_equipment_type_equipment_type_data row")
	}

	if !cached {
		dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCacheMut.Lock()
		dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCache[key] = cache
		dataPlatformEquipmentTypeEquipmentTypeDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_equipment_type_equipment_type_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformEquipmentTypeEquipmentTypeDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_equipment_type_equipment_type_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformEquipmentTypeEquipmentTypeDatum slice")
	}

	return nil
}

var mySQLDataPlatformEquipmentTypeEquipmentTypeDatumUniqueColumns = []string{
	"EquipmentType",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_equipment_type_equipment_type_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformEquipmentTypeEquipmentTypeDatumUniqueColumns, o)

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

	dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCache[key]
	dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformEquipmentTypeEquipmentTypeDatumAllColumns,
			dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithDefault,
			dataPlatformEquipmentTypeEquipmentTypeDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformEquipmentTypeEquipmentTypeDatumAllColumns,
			dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_equipment_type_equipment_type_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_equipment_type_equipment_type_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_equipment_type_equipment_type_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_equipment_type_equipment_type_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformEquipmentTypeEquipmentTypeDatumType, dataPlatformEquipmentTypeEquipmentTypeDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_equipment_type_equipment_type_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_equipment_type_equipment_type_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCacheMut.Lock()
		dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCache[key] = cache
		dataPlatformEquipmentTypeEquipmentTypeDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformEquipmentTypeEquipmentTypeDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformEquipmentTypeEquipmentTypeDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_equipment_type_equipment_type_data` WHERE `EquipmentType`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_equipment_type_equipment_type_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformEquipmentTypeEquipmentTypeDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformEquipmentTypeEquipmentTypeDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_equipment_type_equipment_type_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformEquipmentTypeEquipmentTypeDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_equipment_type_equipment_type_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformEquipmentTypeEquipmentTypeDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformEquipmentTypeEquipmentTypeDatum(ctx, exec, o.EquipmentType)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformEquipmentTypeEquipmentTypeDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_equipment_type_equipment_type_data`.* FROM `data_platform_equipment_type_equipment_type_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformEquipmentTypeEquipmentTypeDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformEquipmentTypeEquipmentTypeDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformEquipmentTypeEquipmentTypeDatumExists checks if the DataPlatformEquipmentTypeEquipmentTypeDatum row exists.
func DataPlatformEquipmentTypeEquipmentTypeDatumExists(ctx context.Context, exec boil.ContextExecutor, equipmentType string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_equipment_type_equipment_type_data` where `EquipmentType`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, equipmentType)
	}
	row := exec.QueryRowContext(ctx, sql, equipmentType)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_equipment_type_equipment_type_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformEquipmentTypeEquipmentTypeDatum row exists.
func (o *DataPlatformEquipmentTypeEquipmentTypeDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformEquipmentTypeEquipmentTypeDatumExists(ctx, exec, o.EquipmentType)
}
