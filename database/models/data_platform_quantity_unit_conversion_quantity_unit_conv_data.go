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

// DataPlatformQuantityUnitConversionQuantityUnitConvDatum is an object representing the database table.
type DataPlatformQuantityUnitConversionQuantityUnitConvDatum struct {
	QuantityUnitFrom      string  `boil:"QuantityUnitFrom" json:"QuantityUnitFrom" toml:"QuantityUnitFrom" yaml:"QuantityUnitFrom"`
	QuantityUnitTo        string  `boil:"QuantityUnitTo" json:"QuantityUnitTo" toml:"QuantityUnitTo" yaml:"QuantityUnitTo"`
	ConversionCoefficient float32 `boil:"ConversionCoefficient" json:"ConversionCoefficient" toml:"ConversionCoefficient" yaml:"ConversionCoefficient"`

	R *dataPlatformQuantityUnitConversionQuantityUnitConvDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformQuantityUnitConversionQuantityUnitConvDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformQuantityUnitConversionQuantityUnitConvDatumColumns = struct {
	QuantityUnitFrom      string
	QuantityUnitTo        string
	ConversionCoefficient string
}{
	QuantityUnitFrom:      "QuantityUnitFrom",
	QuantityUnitTo:        "QuantityUnitTo",
	ConversionCoefficient: "ConversionCoefficient",
}

var DataPlatformQuantityUnitConversionQuantityUnitConvDatumTableColumns = struct {
	QuantityUnitFrom      string
	QuantityUnitTo        string
	ConversionCoefficient string
}{
	QuantityUnitFrom:      "data_platform_quantity_unit_conversion_quantity_unit_conv_data.QuantityUnitFrom",
	QuantityUnitTo:        "data_platform_quantity_unit_conversion_quantity_unit_conv_data.QuantityUnitTo",
	ConversionCoefficient: "data_platform_quantity_unit_conversion_quantity_unit_conv_data.ConversionCoefficient",
}

// Generated where

var DataPlatformQuantityUnitConversionQuantityUnitConvDatumWhere = struct {
	QuantityUnitFrom      whereHelperstring
	QuantityUnitTo        whereHelperstring
	ConversionCoefficient whereHelperfloat32
}{
	QuantityUnitFrom:      whereHelperstring{field: "`data_platform_quantity_unit_conversion_quantity_unit_conv_data`.`QuantityUnitFrom`"},
	QuantityUnitTo:        whereHelperstring{field: "`data_platform_quantity_unit_conversion_quantity_unit_conv_data`.`QuantityUnitTo`"},
	ConversionCoefficient: whereHelperfloat32{field: "`data_platform_quantity_unit_conversion_quantity_unit_conv_data`.`ConversionCoefficient`"},
}

// DataPlatformQuantityUnitConversionQuantityUnitConvDatumRels is where relationship names are stored.
var DataPlatformQuantityUnitConversionQuantityUnitConvDatumRels = struct {
	QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum string
}{
	QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum: "QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum",
}

// dataPlatformQuantityUnitConversionQuantityUnitConvDatumR is where relationships are stored.
type dataPlatformQuantityUnitConversionQuantityUnitConvDatumR struct {
	QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum *DataPlatformQuantityUnitQuantityUnitDatum `boil:"QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum" json:"QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum" toml:"QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum" yaml:"QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformQuantityUnitConversionQuantityUnitConvDatumR) NewStruct() *dataPlatformQuantityUnitConversionQuantityUnitConvDatumR {
	return &dataPlatformQuantityUnitConversionQuantityUnitConvDatumR{}
}

func (r *dataPlatformQuantityUnitConversionQuantityUnitConvDatumR) GetQuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum() *DataPlatformQuantityUnitQuantityUnitDatum {
	if r == nil {
		return nil
	}
	return r.QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum
}

// dataPlatformQuantityUnitConversionQuantityUnitConvDatumL is where Load methods for each relationship are stored.
type dataPlatformQuantityUnitConversionQuantityUnitConvDatumL struct{}

var (
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumAllColumns            = []string{"QuantityUnitFrom", "QuantityUnitTo", "ConversionCoefficient"}
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithoutDefault = []string{"QuantityUnitFrom", "QuantityUnitTo", "ConversionCoefficient"}
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithDefault    = []string{}
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns     = []string{"QuantityUnitFrom", "QuantityUnitTo"}
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice is an alias for a slice of pointers to DataPlatformQuantityUnitConversionQuantityUnitConvDatum.
	// This should almost always be used instead of []DataPlatformQuantityUnitConversionQuantityUnitConvDatum.
	DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice []*DataPlatformQuantityUnitConversionQuantityUnitConvDatum

	dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumType                 = reflect.TypeOf(&DataPlatformQuantityUnitConversionQuantityUnitConvDatum{})
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping              = queries.MakeStructMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCacheMut       sync.RWMutex
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCache          = make(map[string]insertCache)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCacheMut       sync.RWMutex
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCacheMut       sync.RWMutex
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformQuantityUnitConversionQuantityUnitConvDatum record from the query.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformQuantityUnitConversionQuantityUnitConvDatum, error) {
	o := &DataPlatformQuantityUnitConversionQuantityUnitConvDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	return o, nil
}

// All returns all DataPlatformQuantityUnitConversionQuantityUnitConvDatum records from the query.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice, error) {
	var o []*DataPlatformQuantityUnitConversionQuantityUnitConvDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformQuantityUnitConversionQuantityUnitConvDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformQuantityUnitConversionQuantityUnitConvDatum records in the query.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_quantity_unit_conversion_quantity_unit_conv_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_quantity_unit_conversion_quantity_unit_conv_data exists")
	}

	return count > 0, nil
}

// QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum pointed to by the foreign key.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum(mods ...qm.QueryMod) dataPlatformQuantityUnitQuantityUnitDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`QuantityUnit` = ?", o.QuantityUnitFrom),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformQuantityUnitQuantityUnitData(queryMods...)
}

// LoadQuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformQuantityUnitConversionQuantityUnitConvDatumL) LoadQuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformQuantityUnitConversionQuantityUnitConvDatum
	var object *DataPlatformQuantityUnitConversionQuantityUnitConvDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum.(*DataPlatformQuantityUnitConversionQuantityUnitConvDatum)
		if !ok {
			object = new(DataPlatformQuantityUnitConversionQuantityUnitConvDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum.(*[]*DataPlatformQuantityUnitConversionQuantityUnitConvDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformQuantityUnitConversionQuantityUnitConvDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformQuantityUnitConversionQuantityUnitConvDatumR{}
		}
		args = append(args, object.QuantityUnitFrom)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformQuantityUnitConversionQuantityUnitConvDatumR{}
			}

			for _, a := range args {
				if a == obj.QuantityUnitFrom {
					continue Outer
				}
			}

			args = append(args, obj.QuantityUnitFrom)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_quantity_unit_quantity_unit_data`),
		qm.WhereIn(`data_platform_quantity_unit_quantity_unit_data.QuantityUnit in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformQuantityUnitQuantityUnitDatum")
	}

	var resultSlice []*DataPlatformQuantityUnitQuantityUnitDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformQuantityUnitQuantityUnitDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_quantity_unit_quantity_unit_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_quantity_unit_quantity_unit_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuantityUnitFrom == foreign.QuantityUnit {
				local.R.QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum = foreign
				break
			}
		}
	}

	return nil
}

// SetQuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum of the dataPlatformQuantityUnitConversionQuantityUnitConvDatum to the related item.
// Sets o.R.QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum to related.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) SetQuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformQuantityUnitQuantityUnitDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_quantity_unit_conversion_quantity_unit_conv_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"QuantityUnitFrom"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.QuantityUnit, o.QuantityUnitFrom, o.QuantityUnitTo}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.QuantityUnitFrom = related.QuantityUnit
	if o.R == nil {
		o.R = &dataPlatformQuantityUnitConversionQuantityUnitConvDatumR{
			QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum: related,
		}
	} else {
		o.R.QuantityUnitFromDataPlatformQuantityUnitQuantityUnitDatum = related
	}

	return nil
}

// DataPlatformQuantityUnitConversionQuantityUnitConvData retrieves all the records using an executor.
func DataPlatformQuantityUnitConversionQuantityUnitConvData(mods ...qm.QueryMod) dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery {
	mods = append(mods, qm.From("`data_platform_quantity_unit_conversion_quantity_unit_conv_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_quantity_unit_conversion_quantity_unit_conv_data`.*"})
	}

	return dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery{q}
}

// FindDataPlatformQuantityUnitConversionQuantityUnitConvDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformQuantityUnitConversionQuantityUnitConvDatum(ctx context.Context, exec boil.ContextExecutor, quantityUnitFrom string, quantityUnitTo string, selectCols ...string) (*DataPlatformQuantityUnitConversionQuantityUnitConvDatum, error) {
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumObj := &DataPlatformQuantityUnitConversionQuantityUnitConvDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_quantity_unit_conversion_quantity_unit_conv_data` where `QuantityUnitFrom`=? AND `QuantityUnitTo`=?", sel,
	)

	q := queries.Raw(query, quantityUnitFrom, quantityUnitTo)

	err := q.Bind(ctx, exec, dataPlatformQuantityUnitConversionQuantityUnitConvDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	return dataPlatformQuantityUnitConversionQuantityUnitConvDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_quantity_unit_conversion_quantity_unit_conv_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCache[key]
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumAllColumns,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithDefault,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_quantity_unit_conversion_quantity_unit_conv_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_quantity_unit_conversion_quantity_unit_conv_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_quantity_unit_conversion_quantity_unit_conv_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.QuantityUnitFrom,
		o.QuantityUnitTo,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCacheMut.Lock()
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCache[key] = cache
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformQuantityUnitConversionQuantityUnitConvDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCache[key]
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumAllColumns,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_quantity_unit_conversion_quantity_unit_conv_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_quantity_unit_conversion_quantity_unit_conv_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, append(wl, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_quantity_unit_conversion_quantity_unit_conv_data row")
	}

	if !cached {
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCacheMut.Lock()
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCache[key] = cache
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_quantity_unit_conversion_quantity_unit_conv_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformQuantityUnitConversionQuantityUnitConvDatum slice")
	}

	return nil
}

var mySQLDataPlatformQuantityUnitConversionQuantityUnitConvDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_quantity_unit_conversion_quantity_unit_conv_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformQuantityUnitConversionQuantityUnitConvDatumUniqueColumns, o)

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

	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCache[key]
	dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumAllColumns,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithDefault,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumAllColumns,
			dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_quantity_unit_conversion_quantity_unit_conv_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_quantity_unit_conversion_quantity_unit_conv_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_quantity_unit_conversion_quantity_unit_conv_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformQuantityUnitConversionQuantityUnitConvDatumType, dataPlatformQuantityUnitConversionQuantityUnitConvDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCacheMut.Lock()
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCache[key] = cache
		dataPlatformQuantityUnitConversionQuantityUnitConvDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformQuantityUnitConversionQuantityUnitConvDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformQuantityUnitConversionQuantityUnitConvDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_quantity_unit_conversion_quantity_unit_conv_data` WHERE `QuantityUnitFrom`=? AND `QuantityUnitTo`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformQuantityUnitConversionQuantityUnitConvDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_quantity_unit_conversion_quantity_unit_conv_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_quantity_unit_conversion_quantity_unit_conv_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformQuantityUnitConversionQuantityUnitConvDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformQuantityUnitConversionQuantityUnitConvDatum(ctx, exec, o.QuantityUnitFrom, o.QuantityUnitTo)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_quantity_unit_conversion_quantity_unit_conv_data`.* FROM `data_platform_quantity_unit_conversion_quantity_unit_conv_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuantityUnitConversionQuantityUnitConvDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformQuantityUnitConversionQuantityUnitConvDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformQuantityUnitConversionQuantityUnitConvDatumExists checks if the DataPlatformQuantityUnitConversionQuantityUnitConvDatum row exists.
func DataPlatformQuantityUnitConversionQuantityUnitConvDatumExists(ctx context.Context, exec boil.ContextExecutor, quantityUnitFrom string, quantityUnitTo string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_quantity_unit_conversion_quantity_unit_conv_data` where `QuantityUnitFrom`=? AND `QuantityUnitTo`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, quantityUnitFrom, quantityUnitTo)
	}
	row := exec.QueryRowContext(ctx, sql, quantityUnitFrom, quantityUnitTo)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_quantity_unit_conversion_quantity_unit_conv_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformQuantityUnitConversionQuantityUnitConvDatum row exists.
func (o *DataPlatformQuantityUnitConversionQuantityUnitConvDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformQuantityUnitConversionQuantityUnitConvDatumExists(ctx, exec, o.QuantityUnitFrom, o.QuantityUnitTo)
}
