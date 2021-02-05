package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jbpratt/tos/internal/pb"
)

func TestOrderServiceFull(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	orderService, err := NewOrderService(db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewMenuService(db)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []*pb.Order{
		{Name: "test", Total: 799, Status: "active", TimeOrdered: time.Now().Format("2006-01-02 15:04:05")},
		{Name: "13089 lfak", Total: 1000, Status: "active", TimeOrdered: time.Now().Format("2006-01-02 15:04:05")},
		{Name: "q", Total: 1, Status: "active", TimeOrdered: time.Now().Format("2006-01-02 15:04:05")},
		{Name: "", Total: 9182091809182, Status: "active", TimeOrdered: time.Now().Format("2006-01-02 15:04:05")},
		{
			Name: "majora", Status: "active", TimeOrdered: time.Now().Format("2006-01-03 15:04:05"),
			Items: []*pb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Options: []*pb.Option{
					{Name: "pickles", Price: 0, Selected: true, Id: 1},
				}},
			},
			Total: 495,
		},
	}

	for _, order := range testCases {
		if err = orderService.SubmitOrder(ctx, order); err != nil {
			spew.Dump(order)
			t.Errorf("SubmitOrder(%q) = %q", order.Name, err)
		}
	}

	got, err := orderService.GetOrders(ctx)
	if err != nil {
		spew.Dump(got)
		t.Fatalf("GetOrders() = %q; want nil", err)
	}

	if len(got) != len(testCases) {
		spew.Dump(got)
		t.Fatalf("GetOrders() = %d; want %d", len(got), len(testCases))
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
