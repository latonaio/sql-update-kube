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

// DataPlatformOrdersAddressDatum is an object representing the database table.
type DataPlatformOrdersAddressDatum struct {
	OrderID     int         `boil:"OrderID" json:"OrderID" toml:"OrderID" yaml:"OrderID"`
	AddressID   int         `boil:"AddressID" json:"AddressID" toml:"AddressID" yaml:"AddressID"`
	PostalCode  null.String `boil:"PostalCode" json:"PostalCode,omitempty" toml:"PostalCode" yaml:"PostalCode,omitempty"`
	LocalRegion null.String `boil:"LocalRegion" json:"LocalRegion,omitempty" toml:"LocalRegion" yaml:"LocalRegion,omitempty"`
	Country     null.String `boil:"Country" json:"Country,omitempty" toml:"Country" yaml:"Country,omitempty"`
	District    null.String `boil:"District" json:"District,omitempty" toml:"District" yaml:"District,omitempty"`
	StreetName  null.String `boil:"StreetName" json:"StreetName,omitempty" toml:"StreetName" yaml:"StreetName,omitempty"`
	CityName    null.String `boil:"CityName" json:"CityName,omitempty" toml:"CityName" yaml:"CityName,omitempty"`
	Building    null.String `boil:"Building" json:"Building,omitempty" toml:"Building" yaml:"Building,omitempty"`
	Floor       null.Int    `boil:"Floor" json:"Floor,omitempty" toml:"Floor" yaml:"Floor,omitempty"`
	Room        null.Int    `boil:"Room" json:"Room,omitempty" toml:"Room" yaml:"Room,omitempty"`

	R *dataPlatformOrdersAddressDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformOrdersAddressDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformOrdersAddressDatumColumns = struct {
	OrderID     string
	AddressID   string
	PostalCode  string
	LocalRegion string
	Country     string
	District    string
	StreetName  string
	CityName    string
	Building    string
	Floor       string
	Room        string
}{
	OrderID:     "OrderID",
	AddressID:   "AddressID",
	PostalCode:  "PostalCode",
	LocalRegion: "LocalRegion",
	Country:     "Country",
	District:    "District",
	StreetName:  "StreetName",
	CityName:    "CityName",
	Building:    "Building",
	Floor:       "Floor",
	Room:        "Room",
}

var DataPlatformOrdersAddressDatumTableColumns = struct {
	OrderID     string
	AddressID   string
	PostalCode  string
	LocalRegion string
	Country     string
	District    string
	StreetName  string
	CityName    string
	Building    string
	Floor       string
	Room        string
}{
	OrderID:     "data_platform_orders_address_data.OrderID",
	AddressID:   "data_platform_orders_address_data.AddressID",
	PostalCode:  "data_platform_orders_address_data.PostalCode",
	LocalRegion: "data_platform_orders_address_data.LocalRegion",
	Country:     "data_platform_orders_address_data.Country",
	District:    "data_platform_orders_address_data.District",
	StreetName:  "data_platform_orders_address_data.StreetName",
	CityName:    "data_platform_orders_address_data.CityName",
	Building:    "data_platform_orders_address_data.Building",
	Floor:       "data_platform_orders_address_data.Floor",
	Room:        "data_platform_orders_address_data.Room",
}

// Generated where

var DataPlatformOrdersAddressDatumWhere = struct {
	OrderID     whereHelperint
	AddressID   whereHelperint
	PostalCode  whereHelpernull_String
	LocalRegion whereHelpernull_String
	Country     whereHelpernull_String
	District    whereHelpernull_String
	StreetName  whereHelpernull_String
	CityName    whereHelpernull_String
	Building    whereHelpernull_String
	Floor       whereHelpernull_Int
	Room        whereHelpernull_Int
}{
	OrderID:     whereHelperint{field: "`data_platform_orders_address_data`.`OrderID`"},
	AddressID:   whereHelperint{field: "`data_platform_orders_address_data`.`AddressID`"},
	PostalCode:  whereHelpernull_String{field: "`data_platform_orders_address_data`.`PostalCode`"},
	LocalRegion: whereHelpernull_String{field: "`data_platform_orders_address_data`.`LocalRegion`"},
	Country:     whereHelpernull_String{field: "`data_platform_orders_address_data`.`Country`"},
	District:    whereHelpernull_String{field: "`data_platform_orders_address_data`.`District`"},
	StreetName:  whereHelpernull_String{field: "`data_platform_orders_address_data`.`StreetName`"},
	CityName:    whereHelpernull_String{field: "`data_platform_orders_address_data`.`CityName`"},
	Building:    whereHelpernull_String{field: "`data_platform_orders_address_data`.`Building`"},
	Floor:       whereHelpernull_Int{field: "`data_platform_orders_address_data`.`Floor`"},
	Room:        whereHelpernull_Int{field: "`data_platform_orders_address_data`.`Room`"},
}

// DataPlatformOrdersAddressDatumRels is where relationship names are stored.
var DataPlatformOrdersAddressDatumRels = struct {
	AddressIDDataPlatformAddressAddressDatum string
	OrderIDDataPlatformOrdersHeaderDatum     string
}{
	AddressIDDataPlatformAddressAddressDatum: "AddressIDDataPlatformAddressAddressDatum",
	OrderIDDataPlatformOrdersHeaderDatum:     "OrderIDDataPlatformOrdersHeaderDatum",
}

// dataPlatformOrdersAddressDatumR is where relationships are stored.
type dataPlatformOrdersAddressDatumR struct {
	AddressIDDataPlatformAddressAddressDatum *DataPlatformAddressAddressDatum `boil:"AddressIDDataPlatformAddressAddressDatum" json:"AddressIDDataPlatformAddressAddressDatum" toml:"AddressIDDataPlatformAddressAddressDatum" yaml:"AddressIDDataPlatformAddressAddressDatum"`
	OrderIDDataPlatformOrdersHeaderDatum     *DataPlatformOrdersHeaderDatum   `boil:"OrderIDDataPlatformOrdersHeaderDatum" json:"OrderIDDataPlatformOrdersHeaderDatum" toml:"OrderIDDataPlatformOrdersHeaderDatum" yaml:"OrderIDDataPlatformOrdersHeaderDatum"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformOrdersAddressDatumR) NewStruct() *dataPlatformOrdersAddressDatumR {
	return &dataPlatformOrdersAddressDatumR{}
}

func (r *dataPlatformOrdersAddressDatumR) GetAddressIDDataPlatformAddressAddressDatum() *DataPlatformAddressAddressDatum {
	if r == nil {
		return nil
	}
	return r.AddressIDDataPlatformAddressAddressDatum
}

func (r *dataPlatformOrdersAddressDatumR) GetOrderIDDataPlatformOrdersHeaderDatum() *DataPlatformOrdersHeaderDatum {
	if r == nil {
		return nil
	}
	return r.OrderIDDataPlatformOrdersHeaderDatum
}

// dataPlatformOrdersAddressDatumL is where Load methods for each relationship are stored.
type dataPlatformOrdersAddressDatumL struct{}

var (
	dataPlatformOrdersAddressDatumAllColumns            = []string{"OrderID", "AddressID", "PostalCode", "LocalRegion", "Country", "District", "StreetName", "CityName", "Building", "Floor", "Room"}
	dataPlatformOrdersAddressDatumColumnsWithoutDefault = []string{"OrderID", "AddressID", "PostalCode", "LocalRegion", "Country", "District", "StreetName", "CityName", "Building", "Floor", "Room"}
	dataPlatformOrdersAddressDatumColumnsWithDefault    = []string{}
	dataPlatformOrdersAddressDatumPrimaryKeyColumns     = []string{"OrderID", "AddressID"}
	dataPlatformOrdersAddressDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformOrdersAddressDatumSlice is an alias for a slice of pointers to DataPlatformOrdersAddressDatum.
	// This should almost always be used instead of []DataPlatformOrdersAddressDatum.
	DataPlatformOrdersAddressDatumSlice []*DataPlatformOrdersAddressDatum

	dataPlatformOrdersAddressDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformOrdersAddressDatumType                 = reflect.TypeOf(&DataPlatformOrdersAddressDatum{})
	dataPlatformOrdersAddressDatumMapping              = queries.MakeStructMapping(dataPlatformOrdersAddressDatumType)
	dataPlatformOrdersAddressDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, dataPlatformOrdersAddressDatumPrimaryKeyColumns)
	dataPlatformOrdersAddressDatumInsertCacheMut       sync.RWMutex
	dataPlatformOrdersAddressDatumInsertCache          = make(map[string]insertCache)
	dataPlatformOrdersAddressDatumUpdateCacheMut       sync.RWMutex
	dataPlatformOrdersAddressDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformOrdersAddressDatumUpsertCacheMut       sync.RWMutex
	dataPlatformOrdersAddressDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dataPlatformOrdersAddressDatum record from the query.
func (q dataPlatformOrdersAddressDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformOrdersAddressDatum, error) {
	o := &DataPlatformOrdersAddressDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_orders_address_data")
	}

	return o, nil
}

// All returns all DataPlatformOrdersAddressDatum records from the query.
func (q dataPlatformOrdersAddressDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformOrdersAddressDatumSlice, error) {
	var o []*DataPlatformOrdersAddressDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformOrdersAddressDatum slice")
	}

	return o, nil
}

// Count returns the count of all DataPlatformOrdersAddressDatum records in the query.
func (q dataPlatformOrdersAddressDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_orders_address_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformOrdersAddressDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_orders_address_data exists")
	}

	return count > 0, nil
}

// AddressIDDataPlatformAddressAddressDatum pointed to by the foreign key.
func (o *DataPlatformOrdersAddressDatum) AddressIDDataPlatformAddressAddressDatum(mods ...qm.QueryMod) dataPlatformAddressAddressDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`AddressID` = ?", o.AddressID),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformAddressAddressData(queryMods...)
}

// OrderIDDataPlatformOrdersHeaderDatum pointed to by the foreign key.
func (o *DataPlatformOrdersAddressDatum) OrderIDDataPlatformOrdersHeaderDatum(mods ...qm.QueryMod) dataPlatformOrdersHeaderDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`OrderID` = ?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformOrdersHeaderData(queryMods...)
}

// LoadAddressIDDataPlatformAddressAddressDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformOrdersAddressDatumL) LoadAddressIDDataPlatformAddressAddressDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformOrdersAddressDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformOrdersAddressDatum
	var object *DataPlatformOrdersAddressDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformOrdersAddressDatum.(*DataPlatformOrdersAddressDatum)
		if !ok {
			object = new(DataPlatformOrdersAddressDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformOrdersAddressDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformOrdersAddressDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformOrdersAddressDatum.(*[]*DataPlatformOrdersAddressDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformOrdersAddressDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformOrdersAddressDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformOrdersAddressDatumR{}
		}
		args = append(args, object.AddressID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformOrdersAddressDatumR{}
			}

			for _, a := range args {
				if a == obj.AddressID {
					continue Outer
				}
			}

			args = append(args, obj.AddressID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_address_address_data`),
		qm.WhereIn(`data_platform_address_address_data.AddressID in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformAddressAddressDatum")
	}

	var resultSlice []*DataPlatformAddressAddressDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformAddressAddressDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_address_address_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_address_address_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.AddressIDDataPlatformAddressAddressDatum = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AddressID == foreign.AddressID {
				local.R.AddressIDDataPlatformAddressAddressDatum = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrderIDDataPlatformOrdersHeaderDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformOrdersAddressDatumL) LoadOrderIDDataPlatformOrdersHeaderDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformOrdersAddressDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformOrdersAddressDatum
	var object *DataPlatformOrdersAddressDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformOrdersAddressDatum.(*DataPlatformOrdersAddressDatum)
		if !ok {
			object = new(DataPlatformOrdersAddressDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformOrdersAddressDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformOrdersAddressDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformOrdersAddressDatum.(*[]*DataPlatformOrdersAddressDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformOrdersAddressDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformOrdersAddressDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformOrdersAddressDatumR{}
		}
		args = append(args, object.OrderID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformOrdersAddressDatumR{}
			}

			for _, a := range args {
				if a == obj.OrderID {
					continue Outer
				}
			}

			args = append(args, obj.OrderID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_orders_header_data`),
		qm.WhereIn(`data_platform_orders_header_data.OrderID in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformOrdersHeaderDatum")
	}

	var resultSlice []*DataPlatformOrdersHeaderDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformOrdersHeaderDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_orders_header_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_orders_header_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.OrderIDDataPlatformOrdersHeaderDatum = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrderID == foreign.OrderID {
				local.R.OrderIDDataPlatformOrdersHeaderDatum = foreign
				break
			}
		}
	}

	return nil
}

// SetAddressIDDataPlatformAddressAddressDatum of the dataPlatformOrdersAddressDatum to the related item.
// Sets o.R.AddressIDDataPlatformAddressAddressDatum to related.
func (o *DataPlatformOrdersAddressDatum) SetAddressIDDataPlatformAddressAddressDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformAddressAddressDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_orders_address_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"AddressID"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.AddressID, o.OrderID, o.AddressID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AddressID = related.AddressID
	if o.R == nil {
		o.R = &dataPlatformOrdersAddressDatumR{
			AddressIDDataPlatformAddressAddressDatum: related,
		}
	} else {
		o.R.AddressIDDataPlatformAddressAddressDatum = related
	}

	return nil
}

// SetOrderIDDataPlatformOrdersHeaderDatum of the dataPlatformOrdersAddressDatum to the related item.
// Sets o.R.OrderIDDataPlatformOrdersHeaderDatum to related.
func (o *DataPlatformOrdersAddressDatum) SetOrderIDDataPlatformOrdersHeaderDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformOrdersHeaderDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_orders_address_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"OrderID"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.OrderID, o.OrderID, o.AddressID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.OrderID
	if o.R == nil {
		o.R = &dataPlatformOrdersAddressDatumR{
			OrderIDDataPlatformOrdersHeaderDatum: related,
		}
	} else {
		o.R.OrderIDDataPlatformOrdersHeaderDatum = related
	}

	return nil
}

// DataPlatformOrdersAddressData retrieves all the records using an executor.
func DataPlatformOrdersAddressData(mods ...qm.QueryMod) dataPlatformOrdersAddressDatumQuery {
	mods = append(mods, qm.From("`data_platform_orders_address_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_orders_address_data`.*"})
	}

	return dataPlatformOrdersAddressDatumQuery{q}
}

// FindDataPlatformOrdersAddressDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformOrdersAddressDatum(ctx context.Context, exec boil.ContextExecutor, orderID int, addressID int, selectCols ...string) (*DataPlatformOrdersAddressDatum, error) {
	dataPlatformOrdersAddressDatumObj := &DataPlatformOrdersAddressDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_orders_address_data` where `OrderID`=? AND `AddressID`=?", sel,
	)

	q := queries.Raw(query, orderID, addressID)

	err := q.Bind(ctx, exec, dataPlatformOrdersAddressDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_orders_address_data")
	}

	return dataPlatformOrdersAddressDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformOrdersAddressDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_orders_address_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformOrdersAddressDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformOrdersAddressDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformOrdersAddressDatumInsertCache[key]
	dataPlatformOrdersAddressDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformOrdersAddressDatumAllColumns,
			dataPlatformOrdersAddressDatumColumnsWithDefault,
			dataPlatformOrdersAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_orders_address_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_orders_address_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_orders_address_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into data_platform_orders_address_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OrderID,
		o.AddressID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_orders_address_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformOrdersAddressDatumInsertCacheMut.Lock()
		dataPlatformOrdersAddressDatumInsertCache[key] = cache
		dataPlatformOrdersAddressDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DataPlatformOrdersAddressDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformOrdersAddressDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	dataPlatformOrdersAddressDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformOrdersAddressDatumUpdateCache[key]
	dataPlatformOrdersAddressDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformOrdersAddressDatumAllColumns,
			dataPlatformOrdersAddressDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update data_platform_orders_address_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_orders_address_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, append(wl, dataPlatformOrdersAddressDatumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update data_platform_orders_address_data row")
	}

	if !cached {
		dataPlatformOrdersAddressDatumUpdateCacheMut.Lock()
		dataPlatformOrdersAddressDatumUpdateCache[key] = cache
		dataPlatformOrdersAddressDatumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformOrdersAddressDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for data_platform_orders_address_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformOrdersAddressDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformOrdersAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_orders_address_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dataPlatformOrdersAddressDatum slice")
	}

	return nil
}

var mySQLDataPlatformOrdersAddressDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformOrdersAddressDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_orders_address_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformOrdersAddressDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformOrdersAddressDatumUniqueColumns, o)

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

	dataPlatformOrdersAddressDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformOrdersAddressDatumUpsertCache[key]
	dataPlatformOrdersAddressDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformOrdersAddressDatumAllColumns,
			dataPlatformOrdersAddressDatumColumnsWithDefault,
			dataPlatformOrdersAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformOrdersAddressDatumAllColumns,
			dataPlatformOrdersAddressDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_orders_address_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_orders_address_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_orders_address_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for data_platform_orders_address_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformOrdersAddressDatumType, dataPlatformOrdersAddressDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_orders_address_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_orders_address_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformOrdersAddressDatumUpsertCacheMut.Lock()
		dataPlatformOrdersAddressDatumUpsertCache[key] = cache
		dataPlatformOrdersAddressDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DataPlatformOrdersAddressDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformOrdersAddressDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no DataPlatformOrdersAddressDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformOrdersAddressDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_orders_address_data` WHERE `OrderID`=? AND `AddressID`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from data_platform_orders_address_data")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformOrdersAddressDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no dataPlatformOrdersAddressDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from data_platform_orders_address_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformOrdersAddressDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformOrdersAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_orders_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dataPlatformOrdersAddressDatum slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformOrdersAddressDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformOrdersAddressDatum(ctx, exec, o.OrderID, o.AddressID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformOrdersAddressDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformOrdersAddressDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformOrdersAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_orders_address_data`.* FROM `data_platform_orders_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformOrdersAddressDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformOrdersAddressDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformOrdersAddressDatumExists checks if the DataPlatformOrdersAddressDatum row exists.
func DataPlatformOrdersAddressDatumExists(ctx context.Context, exec boil.ContextExecutor, orderID int, addressID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_orders_address_data` where `OrderID`=? AND `AddressID`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, orderID, addressID)
	}
	row := exec.QueryRowContext(ctx, sql, orderID, addressID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_orders_address_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformOrdersAddressDatum row exists.
func (o *DataPlatformOrdersAddressDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformOrdersAddressDatumExists(ctx, exec, o.OrderID, o.AddressID)
}
