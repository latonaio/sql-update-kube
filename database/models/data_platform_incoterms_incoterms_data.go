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

// DataPlatformIncotermsIncotermsDatum is an object representing the database table.
type DataPlatformIncotermsIncotermsDatum struct {
	Incoterms string `boil:"Incoterms" json:"Incoterms" toml:"Incoterms" yaml:"Incoterms"`

	R *dataPlatformIncotermsIncotermsDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformIncotermsIncotermsDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformIncotermsIncotermsDatumColumns = struct {
	Incoterms string
}{
	Incoterms: "Incoterms",
}

var DataPlatformIncotermsIncotermsDatumTableColumns = struct {
	Incoterms string
}{
	Incoterms: "data_platform_incoterms_incoterms_data.Incoterms",
}

// Generated where

var DataPlatformIncotermsIncotermsDatumWhere = struct {
	Incoterms whereHelperstring
}{
	Incoterms: whereHelperstring{field: "`data_platform_incoterms_incoterms_data`.`Incoterms`"},
}

// DataPlatformIncotermsIncotermsDatumRels is where relationship names are stored.
var DataPlatformIncotermsIncotermsDatumRels = struct {
	IncotermDataPlatformSCRTransactionData string
}{
	IncotermDataPlatformSCRTransactionData: "IncotermDataPlatformSCRTransactionData",
}

// dataPlatformIncotermsIncotermsDatumR is where relationships are stored.
type dataPlatformIncotermsIncotermsDatumR struct {
	IncotermDataPlatformSCRTransactionData DataPlatformSCRTransactionDatumSlice `boil:"IncotermDataPlatformSCRTransactionData" json:"IncotermDataPlatformSCRTransactionData" toml:"IncotermDataPlatformSCRTransactionData" yaml:"IncotermDataPlatformSCRTransactionData"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformIncotermsIncotermsDatumR) NewStruct() *dataPlatformIncotermsIncotermsDatumR {
	return &dataPlatformIncotermsIncotermsDatumR{}
}

func (r *dataPlatformIncotermsIncotermsDatumR) GetIncotermDataPlatformSCRTransactionData() DataPlatformSCRTransactionDatumSlice {
	if r == nil {
		return nil
	}
	return r.IncotermDataPlatformSCRTransactionData
}

// dataPlatformIncotermsIncotermsDatumL is where Load methods for each relationship are stored.
type dataPlatformIncotermsIncotermsDatumL struct{}

var (
	dataPlatformIncotermsIncotermsDatumAllColumns            = []string{"Incoterms"}
	dataPlatformIncotermsIncotermsDatumColumnsWithoutDefault = []string{"Incoterms"}
	dataPlatformIncotermsIncotermsDatumColumnsWithDefault    = []string{}
	dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns     = []string{"Incoterms"}
	dataPlatformIncotermsIncotermsDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformIncotermsIncotermsDatumSlice is an alias for a slice of pointers to DataPlatformIncotermsIncotermsDatum.
	// This should almost always be used instead of []DataPlatformIncotermsIncotermsDatum.
	DataPlatformIncotermsIncotermsDatumSlice []*DataPlatformIncotermsIncotermsDatum

	dataPlatformIncotermsIncotermsDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformIncotermsIncotermsDatumType                 = reflect.TypeOf(&DataPlatformIncotermsIncotermsDatum{})
	dataPlatformIncotermsIncotermsDatumMapping              = queries.MakeStructMapping(dataPlatformIncotermsIncotermsDatumType)
	dataPlatformIncotermsIncotermsDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns)
	dataPlatformIncotermsIncotermsDatumInsertCacheMut       sync.RWMutex
	dataPlatformIncotermsIncotermsDatumInsertCache          = make(map[string]insertCache)
	dataPlatformIncotermsIncotermsDatumUpdateCacheMut       sync.RWMutex
	dataPlatformIncotermsIncotermsDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformIncotermsIncotermsDatumUpsertCacheMut       sync.RWMutex
	dataPlatformIncotermsIncotermsDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformIncotermsIncotermsDatum record from the query.
func (q dataPlatformIncotermsIncotermsDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformIncotermsIncotermsDatum, error) {
	o := &DataPlatformIncotermsIncotermsDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_incoterms_incoterms_data")
	}

	return o, nil
}

// All returns all DataPlatformIncotermsIncotermsDatum records from the query.
func (q dataPlatformIncotermsIncotermsDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformIncotermsIncotermsDatumSlice, error) {
	var o []*DataPlatformIncotermsIncotermsDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformIncotermsIncotermsDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformIncotermsIncotermsDatum records in the query.
func (q dataPlatformIncotermsIncotermsDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_incoterms_incoterms_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformIncotermsIncotermsDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_incoterms_incoterms_data exists")
	}

	return count > 0, nil
}

// IncotermDataPlatformSCRTransactionData retrieves all the data_platform_scr_transaction_datum's DataPlatformSCRTransactionData with an executor via Incoterms column.
func (o *DataPlatformIncotermsIncotermsDatum) IncotermDataPlatformSCRTransactionData(mods ...qm.QueryMod) dataPlatformSCRTransactionDatumQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`data_platform_scr_transaction_data`.`Incoterms`=?", o.Incoterms),
	)

	return DataPlatformSCRTransactionData(queryMods...)
}

// LoadIncotermDataPlatformSCRTransactionData allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (dataPlatformIncotermsIncotermsDatumL) LoadIncotermDataPlatformSCRTransactionData(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformIncotermsIncotermsDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformIncotermsIncotermsDatum
	var object *DataPlatformIncotermsIncotermsDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformIncotermsIncotermsDatum.(*DataPlatformIncotermsIncotermsDatum)
		if !ok {
			object = new(DataPlatformIncotermsIncotermsDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformIncotermsIncotermsDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformIncotermsIncotermsDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformIncotermsIncotermsDatum.(*[]*DataPlatformIncotermsIncotermsDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformIncotermsIncotermsDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformIncotermsIncotermsDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformIncotermsIncotermsDatumR{}
		}
		args = append(args, object.Incoterms)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformIncotermsIncotermsDatumR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Incoterms) {
					continue Outer
				}
			}

			args = append(args, obj.Incoterms)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_scr_transaction_data`),
		qm.WhereIn(`data_platform_scr_transaction_data.Incoterms in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load data_platform_scr_transaction_data")
	}

	var resultSlice []*DataPlatformSCRTransactionDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice data_platform_scr_transaction_data")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on data_platform_scr_transaction_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_scr_transaction_data")
	}

	if singular {
		object.R.IncotermDataPlatformSCRTransactionData = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.Incoterms, foreign.Incoterms) {
				local.R.IncotermDataPlatformSCRTransactionData = append(local.R.IncotermDataPlatformSCRTransactionData, foreign)
				break
			}
		}
	}

	return nil
}

// AddIncotermDataPlatformSCRTransactionData adds the given related objects to the existing relationships
// of the data_platform_incoterms_incoterms_datum, optionally inserting them as new records.
// Appends related to o.R.IncotermDataPlatformSCRTransactionData.
func (o *DataPlatformIncotermsIncotermsDatum) AddIncotermDataPlatformSCRTransactionData(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DataPlatformSCRTransactionDatum) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.Incoterms, o.Incoterms)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `data_platform_scr_transaction_data` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"Incoterms"}),
				strmangle.WhereClause("`", "`", 0, dataPlatformSCRTransactionDatumPrimaryKeyColumns),
			)
			values := []interface{}{o.Incoterms, rel.SupplyChainRelationshipID, rel.Buyer, rel.Seller}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.Incoterms, o.Incoterms)
		}
	}

	if o.R == nil {
		o.R = &dataPlatformIncotermsIncotermsDatumR{
			IncotermDataPlatformSCRTransactionData: related,
		}
	} else {
		o.R.IncotermDataPlatformSCRTransactionData = append(o.R.IncotermDataPlatformSCRTransactionData, related...)
	}

	return nil
}

// SetIncotermDataPlatformSCRTransactionData removes all previously related items of the
// data_platform_incoterms_incoterms_datum replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Incoterm's IncotermDataPlatformSCRTransactionData accordingly.
// Replaces o.R.IncotermDataPlatformSCRTransactionData with related.
func (o *DataPlatformIncotermsIncotermsDatum) SetIncotermDataPlatformSCRTransactionData(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DataPlatformSCRTransactionDatum) error {
	query := "update `data_platform_scr_transaction_data` set `Incoterms` = null where `Incoterms` = ?"
	values := []interface{}{o.Incoterms}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		o.R.IncotermDataPlatformSCRTransactionData = nil
	}

	return o.AddIncotermDataPlatformSCRTransactionData(ctx, exec, insert, related...)
}

// RemoveIncotermDataPlatformSCRTransactionData relationships from objects passed in.
// Removes related items from R.IncotermDataPlatformSCRTransactionData (uses pointer comparison, removal does not keep order)
func (o *DataPlatformIncotermsIncotermsDatum) RemoveIncotermDataPlatformSCRTransactionData(ctx context.Context, exec boil.ContextExecutor, related ...*DataPlatformSCRTransactionDatum) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.Incoterms, nil)
		if err = rel.Update(ctx, exec, boil.Whitelist("Incoterms")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.IncotermDataPlatformSCRTransactionData {
			if rel != ri {
				continue
			}

			ln := len(o.R.IncotermDataPlatformSCRTransactionData)
			if ln > 1 && i < ln-1 {
				o.R.IncotermDataPlatformSCRTransactionData[i] = o.R.IncotermDataPlatformSCRTransactionData[ln-1]
			}
			o.R.IncotermDataPlatformSCRTransactionData = o.R.IncotermDataPlatformSCRTransactionData[:ln-1]
			break
		}
	}

	return nil
}

// DataPlatformIncotermsIncotermsData retrieves all the records using an executor.
func DataPlatformIncotermsIncotermsData(mods ...qm.QueryMod) dataPlatformIncotermsIncotermsDatumQuery {
	mods = append(mods, qm.From("`data_platform_incoterms_incoterms_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_incoterms_incoterms_data`.*"})
	}

	return dataPlatformIncotermsIncotermsDatumQuery{q}
}

// FindDataPlatformIncotermsIncotermsDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformIncotermsIncotermsDatum(ctx context.Context, exec boil.ContextExecutor, incoterms string, selectCols ...string) (*DataPlatformIncotermsIncotermsDatum, error) {
	dataPlatformIncotermsIncotermsDatumObj := &DataPlatformIncotermsIncotermsDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_incoterms_incoterms_data` where `Incoterms`=?", sel,
	)

	q := queries.Raw(query, incoterms)

	err := q.Bind(ctx, exec, dataPlatformIncotermsIncotermsDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_incoterms_incoterms_data")
	}

	return dataPlatformIncotermsIncotermsDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformIncotermsIncotermsDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_incoterms_incoterms_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIncotermsIncotermsDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformIncotermsIncotermsDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformIncotermsIncotermsDatumInsertCache[key]
	dataPlatformIncotermsIncotermsDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformIncotermsIncotermsDatumAllColumns,
			dataPlatformIncotermsIncotermsDatumColumnsWithDefault,
			dataPlatformIncotermsIncotermsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_incoterms_incoterms_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_incoterms_incoterms_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_incoterms_incoterms_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_incoterms_incoterms_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Incoterms,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_incoterms_incoterms_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIncotermsIncotermsDatumInsertCacheMut.Lock()
		dataPlatformIncotermsIncotermsDatumInsertCache[key] = cache
		dataPlatformIncotermsIncotermsDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformIncotermsIncotermsDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformIncotermsIncotermsDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformIncotermsIncotermsDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformIncotermsIncotermsDatumUpdateCache[key]
	dataPlatformIncotermsIncotermsDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformIncotermsIncotermsDatumAllColumns,
			dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_incoterms_incoterms_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_incoterms_incoterms_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, append(wl, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_incoterms_incoterms_data row")
	}

	if !cached {
		dataPlatformIncotermsIncotermsDatumUpdateCacheMut.Lock()
		dataPlatformIncotermsIncotermsDatumUpdateCache[key] = cache
		dataPlatformIncotermsIncotermsDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformIncotermsIncotermsDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_incoterms_incoterms_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformIncotermsIncotermsDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsIncotermsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_incoterms_incoterms_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformIncotermsIncotermsDatum slice")
	}

	return nil
}

var mySQLDataPlatformIncotermsIncotermsDatumUniqueColumns = []string{
	"Incoterms",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformIncotermsIncotermsDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_incoterms_incoterms_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIncotermsIncotermsDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformIncotermsIncotermsDatumUniqueColumns, o)

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

	dataPlatformIncotermsIncotermsDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformIncotermsIncotermsDatumUpsertCache[key]
	dataPlatformIncotermsIncotermsDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformIncotermsIncotermsDatumAllColumns,
			dataPlatformIncotermsIncotermsDatumColumnsWithDefault,
			dataPlatformIncotermsIncotermsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformIncotermsIncotermsDatumAllColumns,
			dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_incoterms_incoterms_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_incoterms_incoterms_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_incoterms_incoterms_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_incoterms_incoterms_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformIncotermsIncotermsDatumType, dataPlatformIncotermsIncotermsDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_incoterms_incoterms_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_incoterms_incoterms_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIncotermsIncotermsDatumUpsertCacheMut.Lock()
		dataPlatformIncotermsIncotermsDatumUpsertCache[key] = cache
		dataPlatformIncotermsIncotermsDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformIncotermsIncotermsDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformIncotermsIncotermsDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformIncotermsIncotermsDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformIncotermsIncotermsDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_incoterms_incoterms_data` WHERE `Incoterms`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_incoterms_incoterms_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformIncotermsIncotermsDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformIncotermsIncotermsDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_incoterms_incoterms_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformIncotermsIncotermsDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsIncotermsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_incoterms_incoterms_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformIncotermsIncotermsDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformIncotermsIncotermsDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformIncotermsIncotermsDatum(ctx, exec, o.Incoterms)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformIncotermsIncotermsDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformIncotermsIncotermsDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIncotermsIncotermsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_incoterms_incoterms_data`.* FROM `data_platform_incoterms_incoterms_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIncotermsIncotermsDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformIncotermsIncotermsDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformIncotermsIncotermsDatumExists checks if the DataPlatformIncotermsIncotermsDatum row exists.
func DataPlatformIncotermsIncotermsDatumExists(ctx context.Context, exec boil.ContextExecutor, incoterms string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_incoterms_incoterms_data` where `Incoterms`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, incoterms)
	}
	row := exec.QueryRowContext(ctx, sql, incoterms)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_incoterms_incoterms_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformIncotermsIncotermsDatum row exists.
func (o *DataPlatformIncotermsIncotermsDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformIncotermsIncotermsDatumExists(ctx, exec, o.Incoterms)
}
