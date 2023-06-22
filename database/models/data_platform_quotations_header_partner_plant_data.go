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

// DataPlatformQuotationsHeaderPartnerPlantDatum is an object representing the database table.
type DataPlatformQuotationsHeaderPartnerPlantDatum struct {
	Quotation       int         `boil:"Quotation" json:"Quotation" toml:"Quotation" yaml:"Quotation"`
	PartnerFunction string      `boil:"PartnerFunction" json:"PartnerFunction" toml:"PartnerFunction" yaml:"PartnerFunction"`
	BusinessPartner int         `boil:"BusinessPartner" json:"BusinessPartner" toml:"BusinessPartner" yaml:"BusinessPartner"`
	Plant           null.String `boil:"Plant" json:"Plant,omitempty" toml:"Plant" yaml:"Plant,omitempty"`

	R *dataPlatformQuotationsHeaderPartnerPlantDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformQuotationsHeaderPartnerPlantDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformQuotationsHeaderPartnerPlantDatumColumns = struct {
	Quotation       string
	PartnerFunction string
	BusinessPartner string
	Plant           string
}{
	Quotation:       "Quotation",
	PartnerFunction: "PartnerFunction",
	BusinessPartner: "BusinessPartner",
	Plant:           "Plant",
}

var DataPlatformQuotationsHeaderPartnerPlantDatumTableColumns = struct {
	Quotation       string
	PartnerFunction string
	BusinessPartner string
	Plant           string
}{
	Quotation:       "data_platform_quotations_header_partner_plant_data.Quotation",
	PartnerFunction: "data_platform_quotations_header_partner_plant_data.PartnerFunction",
	BusinessPartner: "data_platform_quotations_header_partner_plant_data.BusinessPartner",
	Plant:           "data_platform_quotations_header_partner_plant_data.Plant",
}

// Generated where

var DataPlatformQuotationsHeaderPartnerPlantDatumWhere = struct {
	Quotation       whereHelperint
	PartnerFunction whereHelperstring
	BusinessPartner whereHelperint
	Plant           whereHelpernull_String
}{
	Quotation:       whereHelperint{field: "`data_platform_quotations_header_partner_plant_data`.`Quotation`"},
	PartnerFunction: whereHelperstring{field: "`data_platform_quotations_header_partner_plant_data`.`PartnerFunction`"},
	BusinessPartner: whereHelperint{field: "`data_platform_quotations_header_partner_plant_data`.`BusinessPartner`"},
	Plant:           whereHelpernull_String{field: "`data_platform_quotations_header_partner_plant_data`.`Plant`"},
}

// DataPlatformQuotationsHeaderPartnerPlantDatumRels is where relationship names are stored.
var DataPlatformQuotationsHeaderPartnerPlantDatumRels = struct {
	PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum string
}{
	PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum: "PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum",
}

// dataPlatformQuotationsHeaderPartnerPlantDatumR is where relationships are stored.
type dataPlatformQuotationsHeaderPartnerPlantDatumR struct {
	PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum *DataPlatformQuotationsHeaderPartnerDatum `boil:"PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum" json:"PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum" toml:"PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum" yaml:"PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformQuotationsHeaderPartnerPlantDatumR) NewStruct() *dataPlatformQuotationsHeaderPartnerPlantDatumR {
	return &dataPlatformQuotationsHeaderPartnerPlantDatumR{}
}

func (r *dataPlatformQuotationsHeaderPartnerPlantDatumR) GetPartnerFunctionDataPlatformQuotationsHeaderPartnerDatum() *DataPlatformQuotationsHeaderPartnerDatum {
	if r == nil {
		return nil
	}
	return r.PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum
}

// dataPlatformQuotationsHeaderPartnerPlantDatumL is where Load methods for each relationship are stored.
type dataPlatformQuotationsHeaderPartnerPlantDatumL struct{}

var (
	dataPlatformQuotationsHeaderPartnerPlantDatumAllColumns            = []string{"Quotation", "PartnerFunction", "BusinessPartner", "Plant"}
	dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithoutDefault = []string{"Quotation", "PartnerFunction", "BusinessPartner", "Plant"}
	dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithDefault    = []string{}
	dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns     = []string{"Quotation", "PartnerFunction", "BusinessPartner"}
	dataPlatformQuotationsHeaderPartnerPlantDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformQuotationsHeaderPartnerPlantDatumSlice is an alias for a slice of pointers to DataPlatformQuotationsHeaderPartnerPlantDatum.
	// This should almost always be used instead of []DataPlatformQuotationsHeaderPartnerPlantDatum.
	DataPlatformQuotationsHeaderPartnerPlantDatumSlice []*DataPlatformQuotationsHeaderPartnerPlantDatum

	dataPlatformQuotationsHeaderPartnerPlantDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformQuotationsHeaderPartnerPlantDatumType                 = reflect.TypeOf(&DataPlatformQuotationsHeaderPartnerPlantDatum{})
	dataPlatformQuotationsHeaderPartnerPlantDatumMapping              = queries.MakeStructMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType)
	dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns)
	dataPlatformQuotationsHeaderPartnerPlantDatumInsertCacheMut       sync.RWMutex
	dataPlatformQuotationsHeaderPartnerPlantDatumInsertCache          = make(map[string]insertCache)
	dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCacheMut       sync.RWMutex
	dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCacheMut       sync.RWMutex
	dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformQuotationsHeaderPartnerPlantDatum record from the query.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformQuotationsHeaderPartnerPlantDatum, error) {
	o := &DataPlatformQuotationsHeaderPartnerPlantDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_quotations_header_partner_plant_data")
	}

	return o, nil
}

// All returns all DataPlatformQuotationsHeaderPartnerPlantDatum records from the query.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformQuotationsHeaderPartnerPlantDatumSlice, error) {
	var o []*DataPlatformQuotationsHeaderPartnerPlantDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformQuotationsHeaderPartnerPlantDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformQuotationsHeaderPartnerPlantDatum records in the query.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_quotations_header_partner_plant_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_quotations_header_partner_plant_data exists")
	}

	return count > 0, nil
}

// PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum pointed to by the foreign key.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum(mods ...qm.QueryMod) dataPlatformQuotationsHeaderPartnerDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`PartnerFunction` = ?", o.PartnerFunction),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformQuotationsHeaderPartnerData(queryMods...)
}

// LoadPartnerFunctionDataPlatformQuotationsHeaderPartnerDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformQuotationsHeaderPartnerPlantDatumL) LoadPartnerFunctionDataPlatformQuotationsHeaderPartnerDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformQuotationsHeaderPartnerPlantDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformQuotationsHeaderPartnerPlantDatum
	var object *DataPlatformQuotationsHeaderPartnerPlantDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformQuotationsHeaderPartnerPlantDatum.(*DataPlatformQuotationsHeaderPartnerPlantDatum)
		if !ok {
			object = new(DataPlatformQuotationsHeaderPartnerPlantDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformQuotationsHeaderPartnerPlantDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformQuotationsHeaderPartnerPlantDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformQuotationsHeaderPartnerPlantDatum.(*[]*DataPlatformQuotationsHeaderPartnerPlantDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformQuotationsHeaderPartnerPlantDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformQuotationsHeaderPartnerPlantDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformQuotationsHeaderPartnerPlantDatumR{}
		}
		args = append(args, object.PartnerFunction)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformQuotationsHeaderPartnerPlantDatumR{}
			}

			for _, a := range args {
				if a == obj.PartnerFunction {
					continue Outer
				}
			}

			args = append(args, obj.PartnerFunction)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_quotations_header_partner_data`),
		qm.WhereIn(`data_platform_quotations_header_partner_data.PartnerFunction in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformQuotationsHeaderPartnerDatum")
	}

	var resultSlice []*DataPlatformQuotationsHeaderPartnerDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformQuotationsHeaderPartnerDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_quotations_header_partner_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_quotations_header_partner_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PartnerFunction == foreign.PartnerFunction {
				local.R.PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum = foreign
				break
			}
		}
	}

	return nil
}

// SetPartnerFunctionDataPlatformQuotationsHeaderPartnerDatum of the dataPlatformQuotationsHeaderPartnerPlantDatum to the related item.
// Sets o.R.PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum to related.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) SetPartnerFunctionDataPlatformQuotationsHeaderPartnerDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformQuotationsHeaderPartnerDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_quotations_header_partner_plant_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"PartnerFunction"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.PartnerFunction, o.Quotation, o.PartnerFunction, o.BusinessPartner}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PartnerFunction = related.PartnerFunction
	if o.R == nil {
		o.R = &dataPlatformQuotationsHeaderPartnerPlantDatumR{
			PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum: related,
		}
	} else {
		o.R.PartnerFunctionDataPlatformQuotationsHeaderPartnerDatum = related
	}

	return nil
}

// DataPlatformQuotationsHeaderPartnerPlantData retrieves all the records using an executor.
func DataPlatformQuotationsHeaderPartnerPlantData(mods ...qm.QueryMod) dataPlatformQuotationsHeaderPartnerPlantDatumQuery {
	mods = append(mods, qm.From("`data_platform_quotations_header_partner_plant_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_quotations_header_partner_plant_data`.*"})
	}

	return dataPlatformQuotationsHeaderPartnerPlantDatumQuery{q}
}

// FindDataPlatformQuotationsHeaderPartnerPlantDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformQuotationsHeaderPartnerPlantDatum(ctx context.Context, exec boil.ContextExecutor, quotation int, partnerFunction string, businessPartner int, selectCols ...string) (*DataPlatformQuotationsHeaderPartnerPlantDatum, error) {
	dataPlatformQuotationsHeaderPartnerPlantDatumObj := &DataPlatformQuotationsHeaderPartnerPlantDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_quotations_header_partner_plant_data` where `Quotation`=? AND `PartnerFunction`=? AND `BusinessPartner`=?", sel,
	)

	q := queries.Raw(query, quotation, partnerFunction, businessPartner)

	err := q.Bind(ctx, exec, dataPlatformQuotationsHeaderPartnerPlantDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_quotations_header_partner_plant_data")
	}

	return dataPlatformQuotationsHeaderPartnerPlantDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_quotations_header_partner_plant_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformQuotationsHeaderPartnerPlantDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformQuotationsHeaderPartnerPlantDatumInsertCache[key]
	dataPlatformQuotationsHeaderPartnerPlantDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformQuotationsHeaderPartnerPlantDatumAllColumns,
			dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithDefault,
			dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_quotations_header_partner_plant_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_quotations_header_partner_plant_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_quotations_header_partner_plant_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_quotations_header_partner_plant_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Quotation,
		o.PartnerFunction,
		o.BusinessPartner,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_quotations_header_partner_plant_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformQuotationsHeaderPartnerPlantDatumInsertCacheMut.Lock()
		dataPlatformQuotationsHeaderPartnerPlantDatumInsertCache[key] = cache
		dataPlatformQuotationsHeaderPartnerPlantDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformQuotationsHeaderPartnerPlantDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCache[key]
	dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformQuotationsHeaderPartnerPlantDatumAllColumns,
			dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_quotations_header_partner_plant_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_quotations_header_partner_plant_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, append(wl, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_quotations_header_partner_plant_data row")
	}

	if !cached {
		dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCacheMut.Lock()
		dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCache[key] = cache
		dataPlatformQuotationsHeaderPartnerPlantDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_quotations_header_partner_plant_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformQuotationsHeaderPartnerPlantDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_quotations_header_partner_plant_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformQuotationsHeaderPartnerPlantDatum slice")
	}

	return nil
}

var mySQLDataPlatformQuotationsHeaderPartnerPlantDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_quotations_header_partner_plant_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformQuotationsHeaderPartnerPlantDatumUniqueColumns, o)

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

	dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCache[key]
	dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformQuotationsHeaderPartnerPlantDatumAllColumns,
			dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithDefault,
			dataPlatformQuotationsHeaderPartnerPlantDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformQuotationsHeaderPartnerPlantDatumAllColumns,
			dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_quotations_header_partner_plant_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_quotations_header_partner_plant_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_quotations_header_partner_plant_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_quotations_header_partner_plant_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformQuotationsHeaderPartnerPlantDatumType, dataPlatformQuotationsHeaderPartnerPlantDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_quotations_header_partner_plant_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_quotations_header_partner_plant_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCacheMut.Lock()
		dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCache[key] = cache
		dataPlatformQuotationsHeaderPartnerPlantDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformQuotationsHeaderPartnerPlantDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformQuotationsHeaderPartnerPlantDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_quotations_header_partner_plant_data` WHERE `Quotation`=? AND `PartnerFunction`=? AND `BusinessPartner`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_quotations_header_partner_plant_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformQuotationsHeaderPartnerPlantDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformQuotationsHeaderPartnerPlantDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_quotations_header_partner_plant_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformQuotationsHeaderPartnerPlantDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_quotations_header_partner_plant_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformQuotationsHeaderPartnerPlantDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformQuotationsHeaderPartnerPlantDatum(ctx, exec, o.Quotation, o.PartnerFunction, o.BusinessPartner)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformQuotationsHeaderPartnerPlantDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_quotations_header_partner_plant_data`.* FROM `data_platform_quotations_header_partner_plant_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformQuotationsHeaderPartnerPlantDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformQuotationsHeaderPartnerPlantDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformQuotationsHeaderPartnerPlantDatumExists checks if the DataPlatformQuotationsHeaderPartnerPlantDatum row exists.
func DataPlatformQuotationsHeaderPartnerPlantDatumExists(ctx context.Context, exec boil.ContextExecutor, quotation int, partnerFunction string, businessPartner int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_quotations_header_partner_plant_data` where `Quotation`=? AND `PartnerFunction`=? AND `BusinessPartner`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, quotation, partnerFunction, businessPartner)
	}
	row := exec.QueryRowContext(ctx, sql, quotation, partnerFunction, businessPartner)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_quotations_header_partner_plant_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformQuotationsHeaderPartnerPlantDatum row exists.
func (o *DataPlatformQuotationsHeaderPartnerPlantDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformQuotationsHeaderPartnerPlantDatumExists(ctx, exec, o.Quotation, o.PartnerFunction, o.BusinessPartner)
}
