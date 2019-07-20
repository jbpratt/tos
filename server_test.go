package main

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	mookiespb "github.com/jbpratt78/tos/protofiles"
)

func TestServer(t *testing.T) {
	*dbp = ":memory:"

	s, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer() failed with %v", err)
	}

	defer func() {
		s.services.Close()

	}()

	if err = s.loadData(); err != nil {
		t.Fatalf("server.loadData() failed with %v", err)
	}

	menu, err := s.GetMenu(context.Background(), &mookiespb.Empty{})
	if err != nil {
		t.Errorf("server.GetMenu() failed with %v", err)
	}

	testItem := &mookiespb.Item{Name: "Test item", Price: 495, CategoryID: 1}

	res, err := s.CreateMenuItem(context.Background(), testItem)
	if err != nil {
		t.Errorf("server.CreateMenuItem(%v) failed with %v", testItem, err)
	}

	if len(menu.Categories[0].Items)+1 != len(s.menu.Categories[0].Items) {
		spew.Dump(s.menu)
		t.Error("CreateMenuItem() failed to add a new item in the category")
	}

	testItem = &mookiespb.Item{Id: res.GetId(), Name: "New test item", Price: 555, CategoryID: 2}

	_, err = s.UpdateMenuItem(context.Background(), testItem)
	if err != nil {
		t.Fatalf("server.UpdateMenuItem(%v) failed with %v", testItem, err)
	}

	if len(menu.Categories[1].Items)+1 != len(s.menu.Categories[1].Items) {
		spew.Dump(s.menu)
		t.Error("UpdateMenuItem() failed to update the item")
	}

	_, err = s.DeleteMenuItem(context.Background(),
		&mookiespb.DeleteMenuItemRequest{Id: res.GetId()})
	if err != nil {
		t.Errorf("server.DeleteMenuItem() failed with %v", err)
	}

	if len(menu.Categories[1].Items) != len(s.menu.Categories[1].Items) {
		spew.Dump(s.menu)
		t.Errorf("DeleteMenuItem() failed to delete the item")
	}
}
