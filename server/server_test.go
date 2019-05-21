package main

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/jmoiron/sqlx"
)

func TestSubmitOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s occured when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	order := &mookiespb.Order{
		Name: "Majora",
		Items: []*mookiespb.Item{
			{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true, Id: 1},
			}},
		},
		Total: 495,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO orders").WithArgs(order.GetName(), order.GetTotal(), order.GetStatus(), time.Now().Format("2006-01-02 15:04:05"), "").WillReturnResult(sqlmock.NewResult(1, 1))
	for i, item := range order.GetItems() {
		mock.ExpectExec("INSERT INTO order_items").WithArgs(i+1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		for x, _ := range item.GetOptions() {
			mock.ExpectExec("INSERT INTO order_item_option").WithArgs(x+1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		}
	}
	mock.ExpectCommit()

	if err = submitOrder(sqlxDB, order); err != nil {
		t.Errorf("error encountered while submitting order: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCompleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s occured when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE orders").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("UPDATE orders").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err = completeOrder(sqlxDB, 1); err != nil {
		t.Errorf("error encountered while completing order: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
