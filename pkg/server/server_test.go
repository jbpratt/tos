package server

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jbpratt/tos/pkg/pb"
)

func TestServer(t *testing.T) {
	*dbp = ":memory:"
	*prnt = false

	s, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer() failed with %v", err)
	}

	defer s.services.Close()

	if err = s.loadData(); err != nil {
		t.Fatalf("server.loadData() failed with %v", err)
	}

	menu, err := s.GetMenu(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("server.GetMenu() failed with %v", err)
	}

	testItem := &pb.Item{Name: "Test item", Price: 495, CategoryID: 1}

	res, err := s.CreateMenuItem(context.Background(), testItem)
	if err != nil {
		t.Errorf("server.CreateMenuItem(%v) failed with %v", testItem, err)
	}

	if len(menu.Categories[0].Items)+1 != len(s.menu.Categories[0].Items) {
		spew.Dump(s.menu)
		t.Error("CreateMenuItem() failed to add a new item in the category")
	}

	testItem = &pb.Item{Id: res.GetId(), Name: "New test item", Price: 555, CategoryID: 2}

	_, err = s.UpdateMenuItem(context.Background(), testItem)
	if err != nil {
		t.Fatalf("server.UpdateMenuItem(%v) failed with %v", testItem, err)
	}

	if len(menu.Categories[1].Items)+1 != len(s.menu.Categories[1].Items) {
		spew.Dump(s.menu)
		t.Error("UpdateMenuItem() failed to update the item")
	}

	_, err = s.DeleteMenuItem(context.Background(),
		&pb.DeleteMenuItemRequest{Id: res.GetId()})
	if err != nil {
		t.Errorf("server.DeleteMenuItem() failed with %v", err)
	}

	if len(menu.Categories[1].Items) != len(s.menu.Categories[1].Items) {
		spew.Dump(s.menu)
		t.Errorf("DeleteMenuItem() failed to delete the item")
	}

	// subscribe to orders

	testOrders := []*pb.Order{
		{Items: []*pb.Item{testItem}, Total: 555, Name: "order test"},
		{Items: []*pb.Item{testItem}, Total: 999, Name: "mfsjo813ma"},
		{Items: []*pb.Item{testItem}, Total: 1, Name: "majora"},
	}

	for _, o := range testOrders {
		_, err = s.SubmitOrder(context.Background(), o)
		if err != nil {
			t.Errorf("server.SubmitOrder() failed with %v", err)
		}
	}

	// compare order from here to order received on subscribe
	x, err := s.ActiveOrders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("ActiveOrders() failed with %v", err)
	}

	for _, o := range x.GetOrders() {
		_, err = s.CompleteOrder(context.Background(),
			&pb.CompleteOrderRequest{Id: o.GetId()})
		if err != nil {
			spew.Dump(o)
			t.Errorf("server.CompleteOrder() failed with %v", err)
		}
	}
}
