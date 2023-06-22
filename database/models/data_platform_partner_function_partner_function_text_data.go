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

// DataPlatformPartnerFunctionPartnerFunctionTextDatum is an object representing the database table.
type DataPlatformPartnerFunctionPartnerFunctionTextDatum struct {
	PartnerFunction     string      `boil:"PartnerFunction" json:"PartnerFunction" toml:"PartnerFunction" yaml:"PartnerFunction"`
	Language            string      `boil:"Language" json:"Language" toml:"Language" yaml:"Language"`
	PartnerFunctionName null.String `boil:"PartnerFunctionName" json:"PartnerFunctionName,omitempty" toml:"PartnerFunctionName" yaml:"PartnerFunctionName,omitempty"`

	R *dataPlatformPartnerFunctionPartnerFunctionTextDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformPartnerFunctionPartnerFunctionTextDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformPartnerFunctionPartnerFunctionTextDatumColumns = struct {
	PartnerFunction     string
	Language            string
	PartnerFunctionName string
}{
	PartnerFunction:     "PartnerFunction",
	Language:            "Language",
	PartnerFunctionName: "PartnerFunctionName",
}

var DataPlatformPartnerFunctionPartnerFunctionTextDatumTableColumns = struct {
	PartnerFunction     string
	Language            string
	PartnerFunctionName string
}{
	PartnerFunction:     "data_platform_partner_function_partner_function_text_data.PartnerFunction",
	Language:            "data_platform_partner_function_partner_function_text_data.Language",
	PartnerFunctionName: "data_platform_partner_function_partner_function_text_data.PartnerFunctionName",
}

// Generated where

var DataPlatformPartnerFunctionPartnerFunctionTextDatumWhere = struct {
	PartnerFunction     whereHelperstring
	Language            whereHelperstring
	PartnerFunctionName whereHelpernull_String
}{
	PartnerFunction:     whereHelperstring{field: "`data_platform_partner_function_partner_function_text_data`.`PartnerFunction`"},
	Language:            whereHelperstring{field: "`data_platform_partner_function_partner_function_text_data`.`Language`"},
	PartnerFunctionName: whereHelpernull_String{field: "`data_platform_partner_function_partner_function_text_data`.`PartnerFunctionName`"},
}

// DataPlatformPartnerFunctionPartnerFunctionTextDatumRels is where relationship names are stored.
var DataPlatformPartnerFunctionPartnerFunctionTextDatumRels = struct {
	LanguageDataPlatformLanguageLanguageDatum string
}{
	LanguageDataPlatformLanguageLanguageDatum: "LanguageDataPlatformLanguageLanguageDatum",
}

// dataPlatformPartnerFunctionPartnerFunctionTextDatumR is where relationships are stored.
type dataPlatformPartnerFunctionPartnerFunctionTextDatumR struct {
	LanguageDataPlatformLanguageLanguageDatum *DataPlatformLanguageLanguageDatum `boil:"LanguageDataPlatformLanguageLanguageDatum" json:"LanguageDataPlatformLanguageLanguageDatum" toml:"LanguageDataPlatformLanguageLanguageDatum" yaml:"LanguageDataPlatformLanguageLanguageDatum"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformPartnerFunctionPartnerFunctionTextDatumR) NewStruct() *dataPlatformPartnerFunctionPartnerFunctionTextDatumR {
	return &dataPlatformPartnerFunctionPartnerFunctionTextDatumR{}
}

func (r *dataPlatformPartnerFunctionPartnerFunctionTextDatumR) GetLanguageDataPlatformLanguageLanguageDatum() *DataPlatformLanguageLanguageDatum {
	if r == nil {
		return nil
	}
	return r.LanguageDataPlatformLanguageLanguageDatum
}

// dataPlatformPartnerFunctionPartnerFunctionTextDatumL is where Load methods for each relationship are stored.
type dataPlatformPartnerFunctionPartnerFunctionTextDatumL struct{}

var (
	dataPlatformPartnerFunctionPartnerFunctionTextDatumAllColumns            = []string{"PartnerFunction", "Language", "PartnerFunctionName"}
	dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithoutDefault = []string{"PartnerFunction", "Language", "PartnerFunctionName"}
	dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithDefault    = []string{}
	dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns     = []string{"PartnerFunction", "Language"}
	dataPlatformPartnerFunctionPartnerFunctionTextDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice is an alias for a slice of pointers to DataPlatformPartnerFunctionPartnerFunctionTextDatum.
	// This should almost always be used instead of []DataPlatformPartnerFunctionPartnerFunctionTextDatum.
	DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice []*DataPlatformPartnerFunctionPartnerFunctionTextDatum

	dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformPartnerFunctionPartnerFunctionTextDatumType                 = reflect.TypeOf(&DataPlatformPartnerFunctionPartnerFunctionTextDatum{})
	dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping              = queries.MakeStructMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCacheMut       sync.RWMutex
	dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCache          = make(map[string]insertCache)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCacheMut       sync.RWMutex
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCacheMut       sync.RWMutex
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformPartnerFunctionPartnerFunctionTextDatum record from the query.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformPartnerFunctionPartnerFunctionTextDatum, error) {
	o := &DataPlatformPartnerFunctionPartnerFunctionTextDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_partner_function_partner_function_text_data")
	}

	return o, nil
}

// All returns all DataPlatformPartnerFunctionPartnerFunctionTextDatum records from the query.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice, error) {
	var o []*DataPlatformPartnerFunctionPartnerFunctionTextDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformPartnerFunctionPartnerFunctionTextDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformPartnerFunctionPartnerFunctionTextDatum records in the query.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_partner_function_partner_function_text_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_partner_function_partner_function_text_data exists")
	}

	return count > 0, nil
}

// LanguageDataPlatformLanguageLanguageDatum pointed to by the foreign key.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) LanguageDataPlatformLanguageLanguageDatum(mods ...qm.QueryMod) dataPlatformLanguageLanguageDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`Language` = ?", o.Language),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformLanguageLanguageData(queryMods...)
}

// LoadLanguageDataPlatformLanguageLanguageDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformPartnerFunctionPartnerFunctionTextDatumL) LoadLanguageDataPlatformLanguageLanguageDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformPartnerFunctionPartnerFunctionTextDatum
	var object *DataPlatformPartnerFunctionPartnerFunctionTextDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum.(*DataPlatformPartnerFunctionPartnerFunctionTextDatum)
		if !ok {
			object = new(DataPlatformPartnerFunctionPartnerFunctionTextDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum.(*[]*DataPlatformPartnerFunctionPartnerFunctionTextDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformPartnerFunctionPartnerFunctionTextDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformPartnerFunctionPartnerFunctionTextDatumR{}
		}
		args = append(args, object.Language)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformPartnerFunctionPartnerFunctionTextDatumR{}
			}

			for _, a := range args {
				if a == obj.Language {
					continue Outer
				}
			}

			args = append(args, obj.Language)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_language_language_data`),
		qm.WhereIn(`data_platform_language_language_data.Language in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformLanguageLanguageDatum")
	}

	var resultSlice []*DataPlatformLanguageLanguageDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformLanguageLanguageDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_language_language_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_language_language_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.LanguageDataPlatformLanguageLanguageDatum = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Language == foreign.Language {
				local.R.LanguageDataPlatformLanguageLanguageDatum = foreign
				break
			}
		}
	}

	return nil
}

// SetLanguageDataPlatformLanguageLanguageDatum of the dataPlatformPartnerFunctionPartnerFunctionTextDatum to the related item.
// Sets o.R.LanguageDataPlatformLanguageLanguageDatum to related.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) SetLanguageDataPlatformLanguageLanguageDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformLanguageLanguageDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_partner_function_partner_function_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"Language"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.Language, o.PartnerFunction, o.Language}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Language = related.Language
	if o.R == nil {
		o.R = &dataPlatformPartnerFunctionPartnerFunctionTextDatumR{
			LanguageDataPlatformLanguageLanguageDatum: related,
		}
	} else {
		o.R.LanguageDataPlatformLanguageLanguageDatum = related
	}

	return nil
}

// DataPlatformPartnerFunctionPartnerFunctionTextData retrieves all the records using an executor.
func DataPlatformPartnerFunctionPartnerFunctionTextData(mods ...qm.QueryMod) dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery {
	mods = append(mods, qm.From("`data_platform_partner_function_partner_function_text_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_partner_function_partner_function_text_data`.*"})
	}

	return dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery{q}
}

// FindDataPlatformPartnerFunctionPartnerFunctionTextDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformPartnerFunctionPartnerFunctionTextDatum(ctx context.Context, exec boil.ContextExecutor, partnerFunction string, language string, selectCols ...string) (*DataPlatformPartnerFunctionPartnerFunctionTextDatum, error) {
	dataPlatformPartnerFunctionPartnerFunctionTextDatumObj := &DataPlatformPartnerFunctionPartnerFunctionTextDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_partner_function_partner_function_text_data` where `PartnerFunction`=? AND `Language`=?", sel,
	)

	q := queries.Raw(query, partnerFunction, language)

	err := q.Bind(ctx, exec, dataPlatformPartnerFunctionPartnerFunctionTextDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_partner_function_partner_function_text_data")
	}

	return dataPlatformPartnerFunctionPartnerFunctionTextDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_partner_function_partner_function_text_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCache[key]
	dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformPartnerFunctionPartnerFunctionTextDatumAllColumns,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithDefault,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_partner_function_partner_function_text_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_partner_function_partner_function_text_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_partner_function_partner_function_text_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_partner_function_partner_function_text_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PartnerFunction,
		o.Language,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_partner_function_partner_function_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCacheMut.Lock()
		dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCache[key] = cache
		dataPlatformPartnerFunctionPartnerFunctionTextDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformPartnerFunctionPartnerFunctionTextDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCache[key]
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformPartnerFunctionPartnerFunctionTextDatumAllColumns,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_partner_function_partner_function_text_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_partner_function_partner_function_text_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, append(wl, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_partner_function_partner_function_text_data row")
	}

	if !cached {
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCacheMut.Lock()
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCache[key] = cache
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_partner_function_partner_function_text_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_partner_function_partner_function_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformPartnerFunctionPartnerFunctionTextDatum slice")
	}

	return nil
}

var mySQLDataPlatformPartnerFunctionPartnerFunctionTextDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_partner_function_partner_function_text_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformPartnerFunctionPartnerFunctionTextDatumUniqueColumns, o)

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

	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCache[key]
	dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformPartnerFunctionPartnerFunctionTextDatumAllColumns,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithDefault,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformPartnerFunctionPartnerFunctionTextDatumAllColumns,
			dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_partner_function_partner_function_text_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_partner_function_partner_function_text_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_partner_function_partner_function_text_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_partner_function_partner_function_text_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformPartnerFunctionPartnerFunctionTextDatumType, dataPlatformPartnerFunctionPartnerFunctionTextDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_partner_function_partner_function_text_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_partner_function_partner_function_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCacheMut.Lock()
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCache[key] = cache
		dataPlatformPartnerFunctionPartnerFunctionTextDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformPartnerFunctionPartnerFunctionTextDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformPartnerFunctionPartnerFunctionTextDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_partner_function_partner_function_text_data` WHERE `PartnerFunction`=? AND `Language`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_partner_function_partner_function_text_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformPartnerFunctionPartnerFunctionTextDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_partner_function_partner_function_text_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_partner_function_partner_function_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformPartnerFunctionPartnerFunctionTextDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformPartnerFunctionPartnerFunctionTextDatum(ctx, exec, o.PartnerFunction, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_partner_function_partner_function_text_data`.* FROM `data_platform_partner_function_partner_function_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformPartnerFunctionPartnerFunctionTextDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformPartnerFunctionPartnerFunctionTextDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformPartnerFunctionPartnerFunctionTextDatumExists checks if the DataPlatformPartnerFunctionPartnerFunctionTextDatum row exists.
func DataPlatformPartnerFunctionPartnerFunctionTextDatumExists(ctx context.Context, exec boil.ContextExecutor, partnerFunction string, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_partner_function_partner_function_text_data` where `PartnerFunction`=? AND `Language`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, partnerFunction, language)
	}
	row := exec.QueryRowContext(ctx, sql, partnerFunction, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_partner_function_partner_function_text_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformPartnerFunctionPartnerFunctionTextDatum row exists.
func (o *DataPlatformPartnerFunctionPartnerFunctionTextDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformPartnerFunctionPartnerFunctionTextDatumExists(ctx, exec, o.PartnerFunction, o.Language)
}
