package models

import (
	"testing"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestSeedMenu(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	err = menuService.SeedMenu()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMenu(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	err = menuService.SeedMenu()
	if err != nil {
		t.Fatal(err)
	}

	menu, err := menuService.GetMenu()
	if err != nil {
		t.Fatal(err)
	}

	if menu == nil {
		t.Error(err)
	}
}

func TestCreateMenuItem(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &mookiespb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	err = menuService.CreateMenuItem(item)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMenuItem(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &mookiespb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	err = menuService.CreateMenuItem(item)
	if err != nil {
		t.Fatal(err)
	}

	err = menuService.DeleteMenuItem(1)
	if err != nil {
		t.Fatal(err)
	}
}
