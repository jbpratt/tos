package models

import (
	"testing"
	"time"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

func TestSubmitOrderFull(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	orderService, err := NewOrderService(db)
	if err != nil {
		t.Fatal(err)
	}

	_ = orderService
	// need items

	t.Fatal("not implemented")
}

func TestSubmitOrderOnly(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	orderService, err := NewOrderService(db)
	if err != nil {
		t.Fatal(err)
	}

	order := &mookiespb.Order{
		Name:        "test",
		Total:       799,
		Status:      "active",
		TimeOrdered: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = orderService.SubmitOrder(order)
	if err != nil {
		t.Fatal(err)
	}

	// want to select it out
}
