// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testItemSides(t *testing.T) {
	t.Parallel()

	query := ItemSides()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testItemSidesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemSidesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ItemSides().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemSidesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ItemSideSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemSidesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ItemSideExists(ctx, tx, o.ItemID, o.SideItemID)
	if err != nil {
		t.Errorf("Unable to check if ItemSide exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ItemSideExists to return true, but got false.")
	}
}

func testItemSidesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	itemSideFound, err := FindItemSide(ctx, tx, o.ItemID, o.SideItemID)
	if err != nil {
		t.Error(err)
	}

	if itemSideFound == nil {
		t.Error("want a record, got nil")
	}
}

func testItemSidesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ItemSides().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testItemSidesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ItemSides().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testItemSidesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	itemSideOne := &ItemSide{}
	itemSideTwo := &ItemSide{}
	if err = randomize.Struct(seed, itemSideOne, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}
	if err = randomize.Struct(seed, itemSideTwo, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = itemSideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = itemSideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ItemSides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testItemSidesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	itemSideOne := &ItemSide{}
	itemSideTwo := &ItemSide{}
	if err = randomize.Struct(seed, itemSideOne, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}
	if err = randomize.Struct(seed, itemSideTwo, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = itemSideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = itemSideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testItemSidesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testItemSidesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(itemSideColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testItemSideToOneItemUsingSideItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ItemSide
	var foreign Item

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.SideItemID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SideItem().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ItemSideSlice{&local}
	if err = local.L.LoadSideItem(ctx, tx, false, (*[]*ItemSide)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SideItem == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SideItem = nil
	if err = local.L.LoadSideItem(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SideItem == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testItemSideToOneItemUsingItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ItemSide
	var foreign Item

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, itemSideDBTypes, false, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ItemID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Item().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ItemSideSlice{&local}
	if err = local.L.LoadItem(ctx, tx, false, (*[]*ItemSide)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Item = nil
	if err = local.L.LoadItem(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testItemSideToOneSetOpItemUsingSideItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ItemSide
	var b, c Item

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemSideDBTypes, false, strmangle.SetComplement(itemSidePrimaryKeyColumns, itemSideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Item{&b, &c} {
		err = a.SetSideItem(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SideItem != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SideItemItemSide != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.SideItemID, x.ID) {
			t.Error("foreign key was wrong value", a.SideItemID)
		}

		if exists, err := ItemSideExists(ctx, tx, a.ItemID, a.SideItemID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testItemSideToOneSetOpItemUsingItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ItemSide
	var b, c Item

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemSideDBTypes, false, strmangle.SetComplement(itemSidePrimaryKeyColumns, itemSideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Item{&b, &c} {
		err = a.SetItem(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Item != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ItemSide != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ItemID, x.ID) {
			t.Error("foreign key was wrong value", a.ItemID)
		}

		if exists, err := ItemSideExists(ctx, tx, a.ItemID, a.SideItemID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testItemSidesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testItemSidesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ItemSideSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testItemSidesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ItemSides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	itemSideDBTypes = map[string]string{`ItemID`: `INTEGER`, `SideItemID`: `INTEGER`, `IsDefault`: `INTEGER`, `Price`: `INTEGER`}
	_               = bytes.MinRead
)

func testItemSidesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(itemSidePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(itemSideAllColumns) == len(itemSidePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSidePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testItemSidesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(itemSideAllColumns) == len(itemSidePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ItemSide{}
	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ItemSides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, itemSideDBTypes, true, itemSidePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ItemSide struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(itemSideAllColumns, itemSidePrimaryKeyColumns) {
		fields = itemSideAllColumns
	} else {
		fields = strmangle.SetComplement(
			itemSideAllColumns,
			itemSidePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ItemSideSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
