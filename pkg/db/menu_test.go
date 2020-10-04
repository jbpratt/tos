package db

import (
	"testing"

	tospb "github.com/jbpratt/tos/protofiles"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestSeedMenu(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.SeedMenu(); err != nil {
		t.Fatal(err)
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestGetMenu(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.SeedMenu(); err != nil {
		t.Fatal(err)
	}

	menu, err := menuService.GetMenu()
	if err != nil {
		t.Fatal(err)
	}

	if menu == nil {
		t.Error(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestCreateMenuItem(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &tospb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	_, err = menuService.CreateMenuItem(item)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMenuItem(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &tospb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	_, err = menuService.CreateMenuItem(item)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.DeleteMenuItem(1); err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
