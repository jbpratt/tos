package main

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

func TestOrder_SubmitOrder(t *testing.T) {
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

func TestOrder_CompleteOrder(t *testing.T) {
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

/*func TestOrder_QueryCorrectOrders(t *testing.T) {

}

func TestMenu_GetMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s occured when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	categories := []*mookiespb.Category{
		&mookiespb.Category{
			Id:   1,
			Name: "Sandwiches",
			/*Items: []*mookiespb.Item{
				&mookiespb.Item{
					Id:    1,
					Name:  "Large Hamburger",
					Price: 495,
					Options: []*mookiespb.Option{
						&mookiespb.Option{
							Id:       1,
							Name:     "pickles",
							Price:    0,
							Selected: true,
						},
					},
				},
			},
		},
		&mookiespb.Category{
			Id:   2,
			Name: "Plates",
			/*Items: []*mookiespb.Item{
				&mookiespb.Item{
					Id:    2,
					Name:  "Smoked Pulled Pork",
					Price: 990,
					Options: []*mookiespb.Option{
						&mookiespb.Option{
							Id:       109,
							Name:     "fries",
							Price:    0,
							Selected: false,
						},
					},
				},
			},
		},
	}
	categoryRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(categories[0].GetId(), categories[0].GetName()).
		AddRow(categories[1].GetId(), categories[1].GetName())

	mock.ExpectQuery("^SELECT (.+) FROM categories").WillReturnRows(categoryRows)

	for _, category := range categories {
		itemRows := sqlmock.NewRows([]string{"id", "name", "price", "category_id"}).
			AddRow(category.Items[0].GetId(), category.Items[0].GetName(),
				category.Items[0].GetPrice(), category.Items[0].GetCategoryID())
		mock.ExpectQuery("^SELECT (.+) FROM items").WithArgs("category_id").WillReturnRows(itemRows)

		for _, item := range category.GetItems() {
			itemOptionRows := sqlmock.NewRows([]string{"name", "price", "selected", "id"}).
				AddRow(item.Options[0].GetName(), item.Options[0].GetPrice(),
					item.Options[0].GetSelected(), item.Options[0].GetId())

			mock.ExpectQuery("^SELECT (.+) FROM options JOIN item_options*").WithArgs("item_id").
				WillReturnRows(itemOptionRows)
		}
	}

	expectedMenu := &mookiespb.Menu{
		Categories: categories,
	}

	menu, err := getMenu(sqlxDB)
	if err != nil {
		t.Fatalf("getMenu() failed with: %v", err)
	}

	fmt.Println(menu)

	assert.Equal(t, expectedMenu, menu)
}
*/
