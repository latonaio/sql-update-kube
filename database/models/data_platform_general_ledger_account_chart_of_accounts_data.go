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

// DataPlatformGeneralLedgerAccountChartOfAccountsDatum is an object representing the database table.
type DataPlatformGeneralLedgerAccountChartOfAccountsDatum struct {
	ChartOfAccounts            string      `boil:"ChartOfAccounts" json:"ChartOfAccounts" toml:"ChartOfAccounts" yaml:"ChartOfAccounts"`
	GLAccount                  string      `boil:"GLAccount" json:"GLAccount" toml:"GLAccount" yaml:"GLAccount"`
	IsBalanceSheetAccount      null.Bool   `boil:"IsBalanceSheetAccount" json:"IsBalanceSheetAccount,omitempty" toml:"IsBalanceSheetAccount" yaml:"IsBalanceSheetAccount,omitempty"`
	IsProfitLossAccount        null.Bool   `boil:"IsProfitLossAccount" json:"IsProfitLossAccount,omitempty" toml:"IsProfitLossAccount" yaml:"IsProfitLossAccount,omitempty"`
	GLAccountGroup             null.String `boil:"GLAccountGroup" json:"GLAccountGroup,omitempty" toml:"GLAccountGroup" yaml:"GLAccountGroup,omitempty"`
	AccountIsMarkedForDeletion null.Bool   `boil:"AccountIsMarkedForDeletion" json:"AccountIsMarkedForDeletion,omitempty" toml:"AccountIsMarkedForDeletion" yaml:"AccountIsMarkedForDeletion,omitempty"`
	AccountIsBlockedForPosting null.Bool   `boil:"AccountIsBlockedForPosting" json:"AccountIsBlockedForPosting,omitempty" toml:"AccountIsBlockedForPosting" yaml:"AccountIsBlockedForPosting,omitempty"`
	CreationDate               null.String `boil:"CreationDate" json:"CreationDate,omitempty" toml:"CreationDate" yaml:"CreationDate,omitempty"`
	LastChangeDateTime         null.String `boil:"LastChangeDateTime" json:"LastChangeDateTime,omitempty" toml:"LastChangeDateTime" yaml:"LastChangeDateTime,omitempty"`

	R *dataPlatformGeneralLedgerAccountChartOfAccountsDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformGeneralLedgerAccountChartOfAccountsDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformGeneralLedgerAccountChartOfAccountsDatumColumns = struct {
	ChartOfAccounts            string
	GLAccount                  string
	IsBalanceSheetAccount      string
	IsProfitLossAccount        string
	GLAccountGroup             string
	AccountIsMarkedForDeletion string
	AccountIsBlockedForPosting string
	CreationDate               string
	LastChangeDateTime         string
}{
	ChartOfAccounts:            "ChartOfAccounts",
	GLAccount:                  "GLAccount",
	IsBalanceSheetAccount:      "IsBalanceSheetAccount",
	IsProfitLossAccount:        "IsProfitLossAccount",
	GLAccountGroup:             "GLAccountGroup",
	AccountIsMarkedForDeletion: "AccountIsMarkedForDeletion",
	AccountIsBlockedForPosting: "AccountIsBlockedForPosting",
	CreationDate:               "CreationDate",
	LastChangeDateTime:         "LastChangeDateTime",
}

var DataPlatformGeneralLedgerAccountChartOfAccountsDatumTableColumns = struct {
	ChartOfAccounts            string
	GLAccount                  string
	IsBalanceSheetAccount      string
	IsProfitLossAccount        string
	GLAccountGroup             string
	AccountIsMarkedForDeletion string
	AccountIsBlockedForPosting string
	CreationDate               string
	LastChangeDateTime         string
}{
	ChartOfAccounts:            "data_platform_general_ledger_account_chart_of_accounts_data.ChartOfAccounts",
	GLAccount:                  "data_platform_general_ledger_account_chart_of_accounts_data.GLAccount",
	IsBalanceSheetAccount:      "data_platform_general_ledger_account_chart_of_accounts_data.IsBalanceSheetAccount",
	IsProfitLossAccount:        "data_platform_general_ledger_account_chart_of_accounts_data.IsProfitLossAccount",
	GLAccountGroup:             "data_platform_general_ledger_account_chart_of_accounts_data.GLAccountGroup",
	AccountIsMarkedForDeletion: "data_platform_general_ledger_account_chart_of_accounts_data.AccountIsMarkedForDeletion",
	AccountIsBlockedForPosting: "data_platform_general_ledger_account_chart_of_accounts_data.AccountIsBlockedForPosting",
	CreationDate:               "data_platform_general_ledger_account_chart_of_accounts_data.CreationDate",
	LastChangeDateTime:         "data_platform_general_ledger_account_chart_of_accounts_data.LastChangeDateTime",
}

// Generated where

var DataPlatformGeneralLedgerAccountChartOfAccountsDatumWhere = struct {
	ChartOfAccounts            whereHelperstring
	GLAccount                  whereHelperstring
	IsBalanceSheetAccount      whereHelpernull_Bool
	IsProfitLossAccount        whereHelpernull_Bool
	GLAccountGroup             whereHelpernull_String
	AccountIsMarkedForDeletion whereHelpernull_Bool
	AccountIsBlockedForPosting whereHelpernull_Bool
	CreationDate               whereHelpernull_String
	LastChangeDateTime         whereHelpernull_String
}{
	ChartOfAccounts:            whereHelperstring{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`ChartOfAccounts`"},
	GLAccount:                  whereHelperstring{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`GLAccount`"},
	IsBalanceSheetAccount:      whereHelpernull_Bool{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`IsBalanceSheetAccount`"},
	IsProfitLossAccount:        whereHelpernull_Bool{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`IsProfitLossAccount`"},
	GLAccountGroup:             whereHelpernull_String{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`GLAccountGroup`"},
	AccountIsMarkedForDeletion: whereHelpernull_Bool{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`AccountIsMarkedForDeletion`"},
	AccountIsBlockedForPosting: whereHelpernull_Bool{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`AccountIsBlockedForPosting`"},
	CreationDate:               whereHelpernull_String{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`CreationDate`"},
	LastChangeDateTime:         whereHelpernull_String{field: "`data_platform_general_ledger_account_chart_of_accounts_data`.`LastChangeDateTime`"},
}

// DataPlatformGeneralLedgerAccountChartOfAccountsDatumRels is where relationship names are stored.
var DataPlatformGeneralLedgerAccountChartOfAccountsDatumRels = struct {
	ChartOfAccountDataPlatformGeneralLedgerAccountTextData string
}{
	ChartOfAccountDataPlatformGeneralLedgerAccountTextData: "ChartOfAccountDataPlatformGeneralLedgerAccountTextData",
}

// dataPlatformGeneralLedgerAccountChartOfAccountsDatumR is where relationships are stored.
type dataPlatformGeneralLedgerAccountChartOfAccountsDatumR struct {
	ChartOfAccountDataPlatformGeneralLedgerAccountTextData DataPlatformGeneralLedgerAccountTextDatumSlice `boil:"ChartOfAccountDataPlatformGeneralLedgerAccountTextData" json:"ChartOfAccountDataPlatformGeneralLedgerAccountTextData" toml:"ChartOfAccountDataPlatformGeneralLedgerAccountTextData" yaml:"ChartOfAccountDataPlatformGeneralLedgerAccountTextData"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformGeneralLedgerAccountChartOfAccountsDatumR) NewStruct() *dataPlatformGeneralLedgerAccountChartOfAccountsDatumR {
	return &dataPlatformGeneralLedgerAccountChartOfAccountsDatumR{}
}

func (r *dataPlatformGeneralLedgerAccountChartOfAccountsDatumR) GetChartOfAccountDataPlatformGeneralLedgerAccountTextData() DataPlatformGeneralLedgerAccountTextDatumSlice {
	if r == nil {
		return nil
	}
	return r.ChartOfAccountDataPlatformGeneralLedgerAccountTextData
}

// dataPlatformGeneralLedgerAccountChartOfAccountsDatumL is where Load methods for each relationship are stored.
type dataPlatformGeneralLedgerAccountChartOfAccountsDatumL struct{}

var (
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumAllColumns            = []string{"ChartOfAccounts", "GLAccount", "IsBalanceSheetAccount", "IsProfitLossAccount", "GLAccountGroup", "AccountIsMarkedForDeletion", "AccountIsBlockedForPosting", "CreationDate", "LastChangeDateTime"}
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithoutDefault = []string{"ChartOfAccounts", "GLAccount", "IsBalanceSheetAccount", "IsProfitLossAccount", "GLAccountGroup", "AccountIsMarkedForDeletion", "AccountIsBlockedForPosting", "CreationDate", "LastChangeDateTime"}
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithDefault    = []string{}
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns     = []string{"ChartOfAccounts", "GLAccount"}
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice is an alias for a slice of pointers to DataPlatformGeneralLedgerAccountChartOfAccountsDatum.
	// This should almost always be used instead of []DataPlatformGeneralLedgerAccountChartOfAccountsDatum.
	DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice []*DataPlatformGeneralLedgerAccountChartOfAccountsDatum

	dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumType                 = reflect.TypeOf(&DataPlatformGeneralLedgerAccountChartOfAccountsDatum{})
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping              = queries.MakeStructMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCacheMut       sync.RWMutex
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCache          = make(map[string]insertCache)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCacheMut       sync.RWMutex
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCacheMut       sync.RWMutex
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformGeneralLedgerAccountChartOfAccountsDatum record from the query.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformGeneralLedgerAccountChartOfAccountsDatum, error) {
	o := &DataPlatformGeneralLedgerAccountChartOfAccountsDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_general_ledger_account_chart_of_accounts_data")
	}

	return o, nil
}

// All returns all DataPlatformGeneralLedgerAccountChartOfAccountsDatum records from the query.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice, error) {
	var o []*DataPlatformGeneralLedgerAccountChartOfAccountsDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformGeneralLedgerAccountChartOfAccountsDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformGeneralLedgerAccountChartOfAccountsDatum records in the query.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_general_ledger_account_chart_of_accounts_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_general_ledger_account_chart_of_accounts_data exists")
	}

	return count > 0, nil
}

// ChartOfAccountDataPlatformGeneralLedgerAccountTextData retrieves all the data_platform_general_ledger_account_text_datum's DataPlatformGeneralLedgerAccountTextData with an executor via ChartOfAccounts column.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) ChartOfAccountDataPlatformGeneralLedgerAccountTextData(mods ...qm.QueryMod) dataPlatformGeneralLedgerAccountTextDatumQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`data_platform_general_ledger_account_text_data`.`ChartOfAccounts`=?", o.ChartOfAccounts),
	)

	return DataPlatformGeneralLedgerAccountTextData(queryMods...)
}

// LoadChartOfAccountDataPlatformGeneralLedgerAccountTextData allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (dataPlatformGeneralLedgerAccountChartOfAccountsDatumL) LoadChartOfAccountDataPlatformGeneralLedgerAccountTextData(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformGeneralLedgerAccountChartOfAccountsDatum
	var object *DataPlatformGeneralLedgerAccountChartOfAccountsDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum.(*DataPlatformGeneralLedgerAccountChartOfAccountsDatum)
		if !ok {
			object = new(DataPlatformGeneralLedgerAccountChartOfAccountsDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum.(*[]*DataPlatformGeneralLedgerAccountChartOfAccountsDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformGeneralLedgerAccountChartOfAccountsDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformGeneralLedgerAccountChartOfAccountsDatumR{}
		}
		args = append(args, object.ChartOfAccounts)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformGeneralLedgerAccountChartOfAccountsDatumR{}
			}

			for _, a := range args {
				if a == obj.ChartOfAccounts {
					continue Outer
				}
			}

			args = append(args, obj.ChartOfAccounts)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_general_ledger_account_text_data`),
		qm.WhereIn(`data_platform_general_ledger_account_text_data.ChartOfAccounts in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load data_platform_general_ledger_account_text_data")
	}

	var resultSlice []*DataPlatformGeneralLedgerAccountTextDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice data_platform_general_ledger_account_text_data")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on data_platform_general_ledger_account_text_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_general_ledger_account_text_data")
	}

	if singular {
		object.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ChartOfAccounts == foreign.ChartOfAccounts {
				local.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData = append(local.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData, foreign)
				break
			}
		}
	}

	return nil
}

// AddChartOfAccountDataPlatformGeneralLedgerAccountTextData adds the given related objects to the existing relationships
// of the data_platform_general_ledger_account_chart_of_accounts_datum, optionally inserting them as new records.
// Appends related to o.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) AddChartOfAccountDataPlatformGeneralLedgerAccountTextData(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DataPlatformGeneralLedgerAccountTextDatum) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChartOfAccounts = o.ChartOfAccounts
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `data_platform_general_ledger_account_text_data` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"ChartOfAccounts"}),
				strmangle.WhereClause("`", "`", 0, dataPlatformGeneralLedgerAccountTextDatumPrimaryKeyColumns),
			)
			values := []interface{}{o.ChartOfAccounts, rel.ChartOfAccounts, rel.GLAccount, rel.Language}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChartOfAccounts = o.ChartOfAccounts
		}
	}

	if o.R == nil {
		o.R = &dataPlatformGeneralLedgerAccountChartOfAccountsDatumR{
			ChartOfAccountDataPlatformGeneralLedgerAccountTextData: related,
		}
	} else {
		o.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData = append(o.R.ChartOfAccountDataPlatformGeneralLedgerAccountTextData, related...)
	}

	return nil
}

// DataPlatformGeneralLedgerAccountChartOfAccountsData retrieves all the records using an executor.
func DataPlatformGeneralLedgerAccountChartOfAccountsData(mods ...qm.QueryMod) dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery {
	mods = append(mods, qm.From("`data_platform_general_ledger_account_chart_of_accounts_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_general_ledger_account_chart_of_accounts_data`.*"})
	}

	return dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery{q}
}

// FindDataPlatformGeneralLedgerAccountChartOfAccountsDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformGeneralLedgerAccountChartOfAccountsDatum(ctx context.Context, exec boil.ContextExecutor, chartOfAccounts string, gLAccount string, selectCols ...string) (*DataPlatformGeneralLedgerAccountChartOfAccountsDatum, error) {
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumObj := &DataPlatformGeneralLedgerAccountChartOfAccountsDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_general_ledger_account_chart_of_accounts_data` where `ChartOfAccounts`=? AND `GLAccount`=?", sel,
	)

	q := queries.Raw(query, chartOfAccounts, gLAccount)

	err := q.Bind(ctx, exec, dataPlatformGeneralLedgerAccountChartOfAccountsDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_general_ledger_account_chart_of_accounts_data")
	}

	return dataPlatformGeneralLedgerAccountChartOfAccountsDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_general_ledger_account_chart_of_accounts_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCache[key]
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumAllColumns,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithDefault,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_general_ledger_account_chart_of_accounts_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_general_ledger_account_chart_of_accounts_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_general_ledger_account_chart_of_accounts_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_general_ledger_account_chart_of_accounts_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ChartOfAccounts,
		o.GLAccount,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_general_ledger_account_chart_of_accounts_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCacheMut.Lock()
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCache[key] = cache
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformGeneralLedgerAccountChartOfAccountsDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCache[key]
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumAllColumns,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_general_ledger_account_chart_of_accounts_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_general_ledger_account_chart_of_accounts_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, append(wl, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_general_ledger_account_chart_of_accounts_data row")
	}

	if !cached {
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCacheMut.Lock()
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCache[key] = cache
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_general_ledger_account_chart_of_accounts_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_general_ledger_account_chart_of_accounts_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformGeneralLedgerAccountChartOfAccountsDatum slice")
	}

	return nil
}

var mySQLDataPlatformGeneralLedgerAccountChartOfAccountsDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_general_ledger_account_chart_of_accounts_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformGeneralLedgerAccountChartOfAccountsDatumUniqueColumns, o)

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

	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCache[key]
	dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumAllColumns,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithDefault,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumAllColumns,
			dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_general_ledger_account_chart_of_accounts_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_general_ledger_account_chart_of_accounts_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_general_ledger_account_chart_of_accounts_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_general_ledger_account_chart_of_accounts_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformGeneralLedgerAccountChartOfAccountsDatumType, dataPlatformGeneralLedgerAccountChartOfAccountsDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_general_ledger_account_chart_of_accounts_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_general_ledger_account_chart_of_accounts_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCacheMut.Lock()
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCache[key] = cache
		dataPlatformGeneralLedgerAccountChartOfAccountsDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformGeneralLedgerAccountChartOfAccountsDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformGeneralLedgerAccountChartOfAccountsDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_general_ledger_account_chart_of_accounts_data` WHERE `ChartOfAccounts`=? AND `GLAccount`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_general_ledger_account_chart_of_accounts_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformGeneralLedgerAccountChartOfAccountsDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_general_ledger_account_chart_of_accounts_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_general_ledger_account_chart_of_accounts_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformGeneralLedgerAccountChartOfAccountsDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformGeneralLedgerAccountChartOfAccountsDatum(ctx, exec, o.ChartOfAccounts, o.GLAccount)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_general_ledger_account_chart_of_accounts_data`.* FROM `data_platform_general_ledger_account_chart_of_accounts_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformGeneralLedgerAccountChartOfAccountsDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformGeneralLedgerAccountChartOfAccountsDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformGeneralLedgerAccountChartOfAccountsDatumExists checks if the DataPlatformGeneralLedgerAccountChartOfAccountsDatum row exists.
func DataPlatformGeneralLedgerAccountChartOfAccountsDatumExists(ctx context.Context, exec boil.ContextExecutor, chartOfAccounts string, gLAccount string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_general_ledger_account_chart_of_accounts_data` where `ChartOfAccounts`=? AND `GLAccount`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, chartOfAccounts, gLAccount)
	}
	row := exec.QueryRowContext(ctx, sql, chartOfAccounts, gLAccount)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_general_ledger_account_chart_of_accounts_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformGeneralLedgerAccountChartOfAccountsDatum row exists.
func (o *DataPlatformGeneralLedgerAccountChartOfAccountsDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformGeneralLedgerAccountChartOfAccountsDatumExists(ctx, exec, o.ChartOfAccounts, o.GLAccount)
}
