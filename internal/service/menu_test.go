package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jbpratt/tos/internal/pb"
	_ "github.com/mattn/go-sqlite3"
)

func TestSeedMenu(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.SeedMenu(ctx); err != nil {
		t.Fatal(err)
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestGetMenu(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.SeedMenu(ctx); err != nil {
		t.Fatal(err)
	}

	menu, err := menuService.GetMenu(ctx)
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &pb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	_, err = menuService.CreateMenuItem(ctx, item)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMenuItem(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	menuService, err := NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	item := &pb.Item{
		Name:       "test",
		CategoryID: 1,
		Price:      399,
	}

	_, err = menuService.CreateMenuItem(ctx, item)
	if err != nil {
		t.Fatal(err)
	}

	if err = menuService.DeleteMenuItem(ctx, 1); err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
