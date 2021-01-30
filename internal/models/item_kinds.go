// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
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

// ItemKind is an object representing the database table.
type ItemKind struct {
	ID      null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Deleted null.Int64 `boil:"deleted" json:"deleted,omitempty" toml:"deleted" yaml:"deleted,omitempty"`
	Name    string     `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *itemKindR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L itemKindL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ItemKindColumns = struct {
	ID      string
	Deleted string
	Name    string
}{
	ID:      "id",
	Deleted: "deleted",
	Name:    "name",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ItemKindWhere = struct {
	ID      whereHelpernull_Int64
	Deleted whereHelpernull_Int64
	Name    whereHelperstring
}{
	ID:      whereHelpernull_Int64{field: "\"item_kinds\".\"id\""},
	Deleted: whereHelpernull_Int64{field: "\"item_kinds\".\"deleted\""},
	Name:    whereHelperstring{field: "\"item_kinds\".\"name\""},
}

// ItemKindRels is where relationship names are stored.
var ItemKindRels = struct {
	KindItems string
}{
	KindItems: "KindItems",
}

// itemKindR is where relationships are stored.
type itemKindR struct {
	KindItems ItemSlice `boil:"KindItems" json:"KindItems" toml:"KindItems" yaml:"KindItems"`
}

// NewStruct creates a new relationship struct
func (*itemKindR) NewStruct() *itemKindR {
	return &itemKindR{}
}

// itemKindL is where Load methods for each relationship are stored.
type itemKindL struct{}

var (
	itemKindAllColumns            = []string{"id", "deleted", "name"}
	itemKindColumnsWithoutDefault = []string{}
	itemKindColumnsWithDefault    = []string{"id", "deleted", "name"}
	itemKindPrimaryKeyColumns     = []string{"id"}
)

type (
	// ItemKindSlice is an alias for a slice of pointers to ItemKind.
	// This should generally be used opposed to []ItemKind.
	ItemKindSlice []*ItemKind

	itemKindQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	itemKindType                 = reflect.TypeOf(&ItemKind{})
	itemKindMapping              = queries.MakeStructMapping(itemKindType)
	itemKindPrimaryKeyMapping, _ = queries.BindMapping(itemKindType, itemKindMapping, itemKindPrimaryKeyColumns)
	itemKindInsertCacheMut       sync.RWMutex
	itemKindInsertCache          = make(map[string]insertCache)
	itemKindUpdateCacheMut       sync.RWMutex
	itemKindUpdateCache          = make(map[string]updateCache)
	itemKindUpsertCacheMut       sync.RWMutex
	itemKindUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single itemKind record from the query using the global executor.
func (q itemKindQuery) OneG(ctx context.Context) (*ItemKind, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single itemKind record from the query.
func (q itemKindQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ItemKind, error) {
	o := &ItemKind{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for item_kinds")
	}

	return o, nil
}

// AllG returns all ItemKind records from the query using the global executor.
func (q itemKindQuery) AllG(ctx context.Context) (ItemKindSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ItemKind records from the query.
func (q itemKindQuery) All(ctx context.Context, exec boil.ContextExecutor) (ItemKindSlice, error) {
	var o []*ItemKind

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ItemKind slice")
	}

	return o, nil
}

// CountG returns the count of all ItemKind records in the query, and panics on error.
func (q itemKindQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ItemKind records in the query.
func (q itemKindQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count item_kinds rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q itemKindQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q itemKindQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if item_kinds exists")
	}

	return count > 0, nil
}

// KindItems retrieves all the item's Items with an executor via kind_id column.
func (o *ItemKind) KindItems(mods ...qm.QueryMod) itemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"items\".\"kind_id\"=?", o.ID),
	)

	query := Items(queryMods...)
	queries.SetFrom(query.Query, "\"items\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"items\".*"})
	}

	return query
}

// LoadKindItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (itemKindL) LoadKindItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeItemKind interface{}, mods queries.Applicator) error {
	var slice []*ItemKind
	var object *ItemKind

	if singular {
		object = maybeItemKind.(*ItemKind)
	} else {
		slice = *maybeItemKind.(*[]*ItemKind)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &itemKindR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &itemKindR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`items`),
		qm.WhereIn(`items.kind_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load items")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if singular {
		object.R.KindItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &itemR{}
			}
			foreign.R.Kind = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.KindID) {
				local.R.KindItems = append(local.R.KindItems, foreign)
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.Kind = local
				break
			}
		}
	}

	return nil
}

// AddKindItemsG adds the given related objects to the existing relationships
// of the item_kind, optionally inserting them as new records.
// Appends related to o.R.KindItems.
// Sets related.R.Kind appropriately.
// Uses the global database handle.
func (o *ItemKind) AddKindItemsG(ctx context.Context, insert bool, related ...*Item) error {
	return o.AddKindItems(ctx, boil.GetContextDB(), insert, related...)
}

// AddKindItems adds the given related objects to the existing relationships
// of the item_kind, optionally inserting them as new records.
// Appends related to o.R.KindItems.
// Sets related.R.Kind appropriately.
func (o *ItemKind) AddKindItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Item) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.KindID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"kind_id"}),
				strmangle.WhereClause("\"", "\"", 0, itemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.KindID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &itemKindR{
			KindItems: related,
		}
	} else {
		o.R.KindItems = append(o.R.KindItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &itemR{
				Kind: o,
			}
		} else {
			rel.R.Kind = o
		}
	}
	return nil
}

// ItemKinds retrieves all the records using an executor.
func ItemKinds(mods ...qm.QueryMod) itemKindQuery {
	mods = append(mods, qm.From("\"item_kinds\""))
	return itemKindQuery{NewQuery(mods...)}
}

// FindItemKindG retrieves a single record by ID.
func FindItemKindG(ctx context.Context, iD null.Int64, selectCols ...string) (*ItemKind, error) {
	return FindItemKind(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindItemKind retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindItemKind(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*ItemKind, error) {
	itemKindObj := &ItemKind{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"item_kinds\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, itemKindObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from item_kinds")
	}

	return itemKindObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ItemKind) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ItemKind) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no item_kinds provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(itemKindColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	itemKindInsertCacheMut.RLock()
	cache, cached := itemKindInsertCache[key]
	itemKindInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			itemKindAllColumns,
			itemKindColumnsWithDefault,
			itemKindColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(itemKindType, itemKindMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(itemKindType, itemKindMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"item_kinds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"item_kinds\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"item_kinds\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, itemKindPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into item_kinds")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for item_kinds")
	}

CacheNoHooks:
	if !cached {
		itemKindInsertCacheMut.Lock()
		itemKindInsertCache[key] = cache
		itemKindInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ItemKind record using the global executor.
// See Update for more documentation.
func (o *ItemKind) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ItemKind.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ItemKind) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	itemKindUpdateCacheMut.RLock()
	cache, cached := itemKindUpdateCache[key]
	itemKindUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			itemKindAllColumns,
			itemKindPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update item_kinds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"item_kinds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, itemKindPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(itemKindType, itemKindMapping, append(wl, itemKindPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update item_kinds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for item_kinds")
	}

	if !cached {
		itemKindUpdateCacheMut.Lock()
		itemKindUpdateCache[key] = cache
		itemKindUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q itemKindQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q itemKindQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for item_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for item_kinds")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ItemKindSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ItemKindSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"item_kinds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemKindPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in itemKind slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all itemKind")
	}
	return rowsAff, nil
}

// DeleteG deletes a single ItemKind record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ItemKind) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ItemKind record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ItemKind) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ItemKind provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), itemKindPrimaryKeyMapping)
	sql := "DELETE FROM \"item_kinds\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from item_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for item_kinds")
	}

	return rowsAff, nil
}

func (q itemKindQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q itemKindQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no itemKindQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from item_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for item_kinds")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ItemKindSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ItemKindSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"item_kinds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemKindPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from itemKind slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for item_kinds")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ItemKind) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ItemKind provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ItemKind) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindItemKind(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemKindSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ItemKindSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemKindSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ItemKindSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"item_kinds\".* FROM \"item_kinds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemKindPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ItemKindSlice")
	}

	*o = slice

	return nil
}

// ItemKindExistsG checks if the ItemKind row exists.
func ItemKindExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return ItemKindExists(ctx, boil.GetContextDB(), iD)
}

// ItemKindExists checks if the ItemKind row exists.
func ItemKindExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"item_kinds\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if item_kinds exists")
	}

	return exists, nil
}
