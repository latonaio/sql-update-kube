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

// DataPlatformProductStockProductStockAvailByBTCH is an object representing the database table.
type DataPlatformProductStockProductStockAvailByBTCH struct {
	Product                      string      `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
	BusinessPartner              int         `boil:"BusinessPartner" json:"BusinessPartner" toml:"BusinessPartner" yaml:"BusinessPartner"`
	Plant                        string      `boil:"Plant" json:"Plant" toml:"Plant" yaml:"Plant"`
	Batch                        string      `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	ProductStockAvailabilityDate string      `boil:"ProductStockAvailabilityDate" json:"ProductStockAvailabilityDate" toml:"ProductStockAvailabilityDate" yaml:"ProductStockAvailabilityDate"`
	InventoryStockType           null.String `boil:"InventoryStockType" json:"InventoryStockType,omitempty" toml:"InventoryStockType" yaml:"InventoryStockType,omitempty"`
	InventorySpecialStockType    null.String `boil:"InventorySpecialStockType" json:"InventorySpecialStockType,omitempty" toml:"InventorySpecialStockType" yaml:"InventorySpecialStockType,omitempty"`
	AvailableProductStock        float32     `boil:"AvailableProductStock" json:"AvailableProductStock" toml:"AvailableProductStock" yaml:"AvailableProductStock"`

	R *dataPlatformProductStockProductStockAvailByBTCHR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformProductStockProductStockAvailByBTCHL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformProductStockProductStockAvailByBTCHColumns = struct {
	Product                      string
	BusinessPartner              string
	Plant                        string
	Batch                        string
	ProductStockAvailabilityDate string
	InventoryStockType           string
	InventorySpecialStockType    string
	AvailableProductStock        string
}{
	Product:                      "Product",
	BusinessPartner:              "BusinessPartner",
	Plant:                        "Plant",
	Batch:                        "Batch",
	ProductStockAvailabilityDate: "ProductStockAvailabilityDate",
	InventoryStockType:           "InventoryStockType",
	InventorySpecialStockType:    "InventorySpecialStockType",
	AvailableProductStock:        "AvailableProductStock",
}

var DataPlatformProductStockProductStockAvailByBTCHTableColumns = struct {
	Product                      string
	BusinessPartner              string
	Plant                        string
	Batch                        string
	ProductStockAvailabilityDate string
	InventoryStockType           string
	InventorySpecialStockType    string
	AvailableProductStock        string
}{
	Product:                      "data_platform_product_stock_product_stock_avail_by_btch.Product",
	BusinessPartner:              "data_platform_product_stock_product_stock_avail_by_btch.BusinessPartner",
	Plant:                        "data_platform_product_stock_product_stock_avail_by_btch.Plant",
	Batch:                        "data_platform_product_stock_product_stock_avail_by_btch.Batch",
	ProductStockAvailabilityDate: "data_platform_product_stock_product_stock_avail_by_btch.ProductStockAvailabilityDate",
	InventoryStockType:           "data_platform_product_stock_product_stock_avail_by_btch.InventoryStockType",
	InventorySpecialStockType:    "data_platform_product_stock_product_stock_avail_by_btch.InventorySpecialStockType",
	AvailableProductStock:        "data_platform_product_stock_product_stock_avail_by_btch.AvailableProductStock",
}

// Generated where

var DataPlatformProductStockProductStockAvailByBTCHWhere = struct {
	Product                      whereHelperstring
	BusinessPartner              whereHelperint
	Plant                        whereHelperstring
	Batch                        whereHelperstring
	ProductStockAvailabilityDate whereHelperstring
	InventoryStockType           whereHelpernull_String
	InventorySpecialStockType    whereHelpernull_String
	AvailableProductStock        whereHelperfloat32
}{
	Product:                      whereHelperstring{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`Product`"},
	BusinessPartner:              whereHelperint{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`BusinessPartner`"},
	Plant:                        whereHelperstring{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`Plant`"},
	Batch:                        whereHelperstring{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`Batch`"},
	ProductStockAvailabilityDate: whereHelperstring{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`ProductStockAvailabilityDate`"},
	InventoryStockType:           whereHelpernull_String{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`InventoryStockType`"},
	InventorySpecialStockType:    whereHelpernull_String{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`InventorySpecialStockType`"},
	AvailableProductStock:        whereHelperfloat32{field: "`data_platform_product_stock_product_stock_avail_by_btch`.`AvailableProductStock`"},
}

// DataPlatformProductStockProductStockAvailByBTCHRels is where relationship names are stored.
var DataPlatformProductStockProductStockAvailByBTCHRels = struct {
	BatchDataPlatformProductStockProductStockByBTCH string
}{
	BatchDataPlatformProductStockProductStockByBTCH: "BatchDataPlatformProductStockProductStockByBTCH",
}

// dataPlatformProductStockProductStockAvailByBTCHR is where relationships are stored.
type dataPlatformProductStockProductStockAvailByBTCHR struct {
	BatchDataPlatformProductStockProductStockByBTCH *DataPlatformProductStockProductStockByBTCH `boil:"BatchDataPlatformProductStockProductStockByBTCH" json:"BatchDataPlatformProductStockProductStockByBTCH" toml:"BatchDataPlatformProductStockProductStockByBTCH" yaml:"BatchDataPlatformProductStockProductStockByBTCH"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformProductStockProductStockAvailByBTCHR) NewStruct() *dataPlatformProductStockProductStockAvailByBTCHR {
	return &dataPlatformProductStockProductStockAvailByBTCHR{}
}

func (r *dataPlatformProductStockProductStockAvailByBTCHR) GetBatchDataPlatformProductStockProductStockByBTCH() *DataPlatformProductStockProductStockByBTCH {
	if r == nil {
		return nil
	}
	return r.BatchDataPlatformProductStockProductStockByBTCH
}

// dataPlatformProductStockProductStockAvailByBTCHL is where Load methods for each relationship are stored.
type dataPlatformProductStockProductStockAvailByBTCHL struct{}

var (
	dataPlatformProductStockProductStockAvailByBTCHAllColumns            = []string{"Product", "BusinessPartner", "Plant", "Batch", "ProductStockAvailabilityDate", "InventoryStockType", "InventorySpecialStockType", "AvailableProductStock"}
	dataPlatformProductStockProductStockAvailByBTCHColumnsWithoutDefault = []string{"Product", "BusinessPartner", "Plant", "Batch", "ProductStockAvailabilityDate", "InventoryStockType", "InventorySpecialStockType", "AvailableProductStock"}
	dataPlatformProductStockProductStockAvailByBTCHColumnsWithDefault    = []string{}
	dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns     = []string{"Product", "BusinessPartner", "Plant", "Batch", "ProductStockAvailabilityDate"}
	dataPlatformProductStockProductStockAvailByBTCHGeneratedColumns      = []string{}
)

type (
	// DataPlatformProductStockProductStockAvailByBTCHSlice is an alias for a slice of pointers to DataPlatformProductStockProductStockAvailByBTCH.
	// This should almost always be used instead of []DataPlatformProductStockProductStockAvailByBTCH.
	DataPlatformProductStockProductStockAvailByBTCHSlice []*DataPlatformProductStockProductStockAvailByBTCH

	dataPlatformProductStockProductStockAvailByBTCHQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformProductStockProductStockAvailByBTCHType                 = reflect.TypeOf(&DataPlatformProductStockProductStockAvailByBTCH{})
	dataPlatformProductStockProductStockAvailByBTCHMapping              = queries.MakeStructMapping(dataPlatformProductStockProductStockAvailByBTCHType)
	dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns)
	dataPlatformProductStockProductStockAvailByBTCHInsertCacheMut       sync.RWMutex
	dataPlatformProductStockProductStockAvailByBTCHInsertCache          = make(map[string]insertCache)
	dataPlatformProductStockProductStockAvailByBTCHUpdateCacheMut       sync.RWMutex
	dataPlatformProductStockProductStockAvailByBTCHUpdateCache          = make(map[string]updateCache)
	dataPlatformProductStockProductStockAvailByBTCHUpsertCacheMut       sync.RWMutex
	dataPlatformProductStockProductStockAvailByBTCHUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformProductStockProductStockAvailByBTCH record from the query.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformProductStockProductStockAvailByBTCH, error) {
	o := &DataPlatformProductStockProductStockAvailByBTCH{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_product_stock_product_stock_avail_by_btch")
	}

	return o, nil
}

// All returns all DataPlatformProductStockProductStockAvailByBTCH records from the query.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformProductStockProductStockAvailByBTCHSlice, error) {
	var o []*DataPlatformProductStockProductStockAvailByBTCH

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformProductStockProductStockAvailByBTCH slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformProductStockProductStockAvailByBTCH records in the query.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_product_stock_product_stock_avail_by_btch rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_product_stock_product_stock_avail_by_btch exists")
	}

	return count > 0, nil
}

// BatchDataPlatformProductStockProductStockByBTCH pointed to by the foreign key.
func (o *DataPlatformProductStockProductStockAvailByBTCH) BatchDataPlatformProductStockProductStockByBTCH(mods ...qm.QueryMod) dataPlatformProductStockProductStockByBTCHQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`Batch` = ?", o.Batch),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformProductStockProductStockByBtches(queryMods...)
}

// LoadBatchDataPlatformProductStockProductStockByBTCH allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformProductStockProductStockAvailByBTCHL) LoadBatchDataPlatformProductStockProductStockByBTCH(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformProductStockProductStockAvailByBTCH interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformProductStockProductStockAvailByBTCH
	var object *DataPlatformProductStockProductStockAvailByBTCH

	if singular {
		var ok bool
		object, ok = maybeDataPlatformProductStockProductStockAvailByBTCH.(*DataPlatformProductStockProductStockAvailByBTCH)
		if !ok {
			object = new(DataPlatformProductStockProductStockAvailByBTCH)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformProductStockProductStockAvailByBTCH)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformProductStockProductStockAvailByBTCH))
			}
		}
	} else {
		s, ok := maybeDataPlatformProductStockProductStockAvailByBTCH.(*[]*DataPlatformProductStockProductStockAvailByBTCH)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformProductStockProductStockAvailByBTCH)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformProductStockProductStockAvailByBTCH))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformProductStockProductStockAvailByBTCHR{}
		}
		args = append(args, object.Batch)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformProductStockProductStockAvailByBTCHR{}
			}

			for _, a := range args {
				if a == obj.Batch {
					continue Outer
				}
			}

			args = append(args, obj.Batch)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_product_stock_product_stock_by_btch`),
		qm.WhereIn(`data_platform_product_stock_product_stock_by_btch.Batch in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformProductStockProductStockByBTCH")
	}

	var resultSlice []*DataPlatformProductStockProductStockByBTCH
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformProductStockProductStockByBTCH")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_product_stock_product_stock_by_btch")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_product_stock_product_stock_by_btch")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.BatchDataPlatformProductStockProductStockByBTCH = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Batch == foreign.Batch {
				local.R.BatchDataPlatformProductStockProductStockByBTCH = foreign
				break
			}
		}
	}

	return nil
}

// SetBatchDataPlatformProductStockProductStockByBTCH of the dataPlatformProductStockProductStockAvailByBTCH to the related item.
// Sets o.R.BatchDataPlatformProductStockProductStockByBTCH to related.
func (o *DataPlatformProductStockProductStockAvailByBTCH) SetBatchDataPlatformProductStockProductStockByBTCH(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformProductStockProductStockByBTCH) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_product_stock_product_stock_avail_by_btch` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"Batch"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns),
	)
	values := []interface{}{related.Batch, o.Product, o.BusinessPartner, o.Plant, o.Batch, o.ProductStockAvailabilityDate}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Batch = related.Batch
	if o.R == nil {
		o.R = &dataPlatformProductStockProductStockAvailByBTCHR{
			BatchDataPlatformProductStockProductStockByBTCH: related,
		}
	} else {
		o.R.BatchDataPlatformProductStockProductStockByBTCH = related
	}

	return nil
}

// DataPlatformProductStockProductStockAvailByBtches retrieves all the records using an executor.
func DataPlatformProductStockProductStockAvailByBtches(mods ...qm.QueryMod) dataPlatformProductStockProductStockAvailByBTCHQuery {
	mods = append(mods, qm.From("`data_platform_product_stock_product_stock_avail_by_btch`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_product_stock_product_stock_avail_by_btch`.*"})
	}

	return dataPlatformProductStockProductStockAvailByBTCHQuery{q}
}

// FindDataPlatformProductStockProductStockAvailByBTCH retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformProductStockProductStockAvailByBTCH(ctx context.Context, exec boil.ContextExecutor, product string, businessPartner int, plant string, batch string, productStockAvailabilityDate string, selectCols ...string) (*DataPlatformProductStockProductStockAvailByBTCH, error) {
	dataPlatformProductStockProductStockAvailByBTCHObj := &DataPlatformProductStockProductStockAvailByBTCH{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_product_stock_product_stock_avail_by_btch` where `Product`=? AND `BusinessPartner`=? AND `Plant`=? AND `Batch`=? AND `ProductStockAvailabilityDate`=?", sel,
	)

	q := queries.Raw(query, product, businessPartner, plant, batch, productStockAvailabilityDate)

	err := q.Bind(ctx, exec, dataPlatformProductStockProductStockAvailByBTCHObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_product_stock_product_stock_avail_by_btch")
	}

	return dataPlatformProductStockProductStockAvailByBTCHObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_stock_product_stock_avail_by_btch provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductStockProductStockAvailByBTCHColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformProductStockProductStockAvailByBTCHInsertCacheMut.RLock()
	cache, cached := dataPlatformProductStockProductStockAvailByBTCHInsertCache[key]
	dataPlatformProductStockProductStockAvailByBTCHInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformProductStockProductStockAvailByBTCHAllColumns,
			dataPlatformProductStockProductStockAvailByBTCHColumnsWithDefault,
			dataPlatformProductStockProductStockAvailByBTCHColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_product_stock_product_stock_avail_by_btch` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_product_stock_product_stock_avail_by_btch` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_product_stock_product_stock_avail_by_btch` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_product_stock_product_stock_avail_by_btch")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Product,
		o.BusinessPartner,
		o.Plant,
		o.Batch,
		o.ProductStockAvailabilityDate,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_stock_product_stock_avail_by_btch")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductStockProductStockAvailByBTCHInsertCacheMut.Lock()
		dataPlatformProductStockProductStockAvailByBTCHInsertCache[key] = cache
		dataPlatformProductStockProductStockAvailByBTCHInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformProductStockProductStockAvailByBTCH.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformProductStockProductStockAvailByBTCHUpdateCacheMut.RLock()
	cache, cached := dataPlatformProductStockProductStockAvailByBTCHUpdateCache[key]
	dataPlatformProductStockProductStockAvailByBTCHUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformProductStockProductStockAvailByBTCHAllColumns,
			dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_product_stock_product_stock_avail_by_btch, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_product_stock_product_stock_avail_by_btch` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, append(wl, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_product_stock_product_stock_avail_by_btch row")
	}

	if !cached {
		dataPlatformProductStockProductStockAvailByBTCHUpdateCacheMut.Lock()
		dataPlatformProductStockProductStockAvailByBTCHUpdateCache[key] = cache
		dataPlatformProductStockProductStockAvailByBTCHUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_product_stock_product_stock_avail_by_btch")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformProductStockProductStockAvailByBTCHSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_product_stock_product_stock_avail_by_btch` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformProductStockProductStockAvailByBTCH slice")
	}

	return nil
}

var mySQLDataPlatformProductStockProductStockAvailByBTCHUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_product_stock_product_stock_avail_by_btch provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformProductStockProductStockAvailByBTCHColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformProductStockProductStockAvailByBTCHUniqueColumns, o)

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

	dataPlatformProductStockProductStockAvailByBTCHUpsertCacheMut.RLock()
	cache, cached := dataPlatformProductStockProductStockAvailByBTCHUpsertCache[key]
	dataPlatformProductStockProductStockAvailByBTCHUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformProductStockProductStockAvailByBTCHAllColumns,
			dataPlatformProductStockProductStockAvailByBTCHColumnsWithDefault,
			dataPlatformProductStockProductStockAvailByBTCHColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformProductStockProductStockAvailByBTCHAllColumns,
			dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_product_stock_product_stock_avail_by_btch, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_product_stock_product_stock_avail_by_btch`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_product_stock_product_stock_avail_by_btch` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_product_stock_product_stock_avail_by_btch")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformProductStockProductStockAvailByBTCHType, dataPlatformProductStockProductStockAvailByBTCHMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_product_stock_product_stock_avail_by_btch")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_product_stock_product_stock_avail_by_btch")
	}

CacheNoHooks:
	if !cached {
		dataPlatformProductStockProductStockAvailByBTCHUpsertCacheMut.Lock()
		dataPlatformProductStockProductStockAvailByBTCHUpsertCache[key] = cache
		dataPlatformProductStockProductStockAvailByBTCHUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformProductStockProductStockAvailByBTCH record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformProductStockProductStockAvailByBTCH provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_product_stock_product_stock_avail_by_btch` WHERE `Product`=? AND `BusinessPartner`=? AND `Plant`=? AND `Batch`=? AND `ProductStockAvailabilityDate`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_product_stock_product_stock_avail_by_btch")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformProductStockProductStockAvailByBTCHQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformProductStockProductStockAvailByBTCHQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_product_stock_product_stock_avail_by_btch")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformProductStockProductStockAvailByBTCHSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_product_stock_product_stock_avail_by_btch` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformProductStockProductStockAvailByBTCH slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformProductStockProductStockAvailByBTCH(ctx, exec, o.Product, o.BusinessPartner, o.Plant, o.Batch, o.ProductStockAvailabilityDate)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformProductStockProductStockAvailByBTCHSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformProductStockProductStockAvailByBTCHSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_product_stock_product_stock_avail_by_btch`.* FROM `data_platform_product_stock_product_stock_avail_by_btch` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformProductStockProductStockAvailByBTCHPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformProductStockProductStockAvailByBTCHSlice")
	}

	*o = slice

	return nil
}

// DataPlatformProductStockProductStockAvailByBTCHExists checks if the DataPlatformProductStockProductStockAvailByBTCH row exists.
func DataPlatformProductStockProductStockAvailByBTCHExists(ctx context.Context, exec boil.ContextExecutor, product string, businessPartner int, plant string, batch string, productStockAvailabilityDate string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_product_stock_product_stock_avail_by_btch` where `Product`=? AND `BusinessPartner`=? AND `Plant`=? AND `Batch`=? AND `ProductStockAvailabilityDate`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, product, businessPartner, plant, batch, productStockAvailabilityDate)
	}
	row := exec.QueryRowContext(ctx, sql, product, businessPartner, plant, batch, productStockAvailabilityDate)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_product_stock_product_stock_avail_by_btch exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformProductStockProductStockAvailByBTCH row exists.
func (o *DataPlatformProductStockProductStockAvailByBTCH) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformProductStockProductStockAvailByBTCHExists(ctx, exec, o.Product, o.BusinessPartner, o.Plant, o.Batch, o.ProductStockAvailabilityDate)
}
