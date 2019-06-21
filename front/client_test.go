package main_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	mpb "github.com/jbpratt78/tos/mock"
	mookiespb "github.com/jbpratt78/tos/protofiles"
)

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)

	mMenuClient.EXPECT().GetMenu(
		gomock.Any(),
		&mookiespb.Empty{},
	).Return(&mookiespb.Menu{}, nil)

	if err := testGetMenu(mMenuClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func TestSubmitOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mOrderClient := mpb.NewMockOrderServiceClient(ctrl)

	req := &mookiespb.SubmitOrderRequest{
		Order: &mookiespb.Order{
			Name: "test",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, Selected: true, Id: 1},
				}},
			},
			Total: 495,
		},
	}

	mOrderClient.EXPECT().SubmitOrder(
		gomock.Any(),
		req,
	).Return(&mookiespb.SubmitOrderResponse{Result: "Order has been placed.."}, nil)

	if err := testSubmitOrder(mOrderClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func TestCreateMenuItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)

	item := &mookiespb.Item{
		Name: "Create item test", Price: 695,
	}

	req := &mookiespb.CreateMenuItemRequest{Item: item}

	mMenuClient.EXPECT().CreateMenuItem(
		gomock.Any(),
		req,
	).Return(&mookiespb.CreateMenuItemResponse{Result: "Item has been created"}, nil)

	if err := testCreateMenuItem(mMenuClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}

}

func TestDeleteMenuItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)

	req := &mookiespb.DeleteMenuItemRequest{Id: 123}

	mMenuClient.EXPECT().DeleteMenuItem(
		gomock.Any(),
		req,
	).Return(&mookiespb.DeleteMenuItemResponse{Result: "Item was deleted"}, nil)

	if err := testDeleteMenuItem(mMenuClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func TestUpdateMenuItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)
	req := &mookiespb.UpdateMenuItemRequest{
		Item: &mookiespb.Item{},
	}

	mMenuClient.EXPECT().UpdateMenuItem(
		gomock.Any(),
		req,
	).Return(&mookiespb.UpdateMenuItemResponse{Result: "Item was updated"}, nil)

	if err := testUpdateMenuItem(mMenuClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func TestCreateOptionToItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)
	req := &mookiespb.CreateMenuItemOptionRequest{
		Option: &mookiespb.Option{},
	}

	mMenuClient.EXPECT().CreateMenuItemOption(
		gomock.Any(),
		req,
	).Return(&mookiespb.CreateMenuItemOptionResponse{Result: "Option was added"}, nil)

	if err := testCreateMenuItemOption(mMenuClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testGetMenu(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.Menu{}

	req := &mookiespb.Empty{}
	got, err := client.GetMenu(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("GetMenu() = %v, want %v", got, want)
	}

	return nil
}

func testSubmitOrder(client mookiespb.OrderServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &mookiespb.SubmitOrderRequest{
		Order: &mookiespb.Order{
			Name: "test",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, Selected: true, Id: 1},
				}},
			},
			Total: 495,
		},
	}

	want := &mookiespb.SubmitOrderResponse{
		Result: "Order has been placed..",
	}
	got, err := client.SubmitOrder(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("SubmitOrder() = %v, want %v", got, want)
	}

	return nil
}

func testCreateMenuItem(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.CreateMenuItemResponse{Result: "Item has been created"}
	req := &mookiespb.CreateMenuItemRequest{
		Item: &mookiespb.Item{Name: "Create item test", Price: 695},
	}

	got, err := client.CreateMenuItem(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("CreateMenuItem() = %v, want %v", got, want)
	}

	return nil
}

func testDeleteMenuItem(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.DeleteMenuItemResponse{Result: "Item was deleted"}
	req := &mookiespb.DeleteMenuItemRequest{Id: 123}

	got, err := client.DeleteMenuItem(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("DeleteMenuItem() = %v, want %v", got, want)
	}

	return nil
}

func testUpdateMenuItem(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.UpdateMenuItemResponse{Result: "Item was updated"}
	req := &mookiespb.UpdateMenuItemRequest{
		Item: &mookiespb.Item{},
	}

	got, err := client.UpdateMenuItem(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("UpdateMenuItemRequest() = %v, want %v", got, want)
	}

	return nil
}

func testCreateMenuItemOption(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.CreateMenuItemOptionResponse{Result: "Option was added"}
	req := &mookiespb.CreateMenuItemOptionRequest{
		Option: &mookiespb.Option{},
	}

	got, err := client.CreateMenuItemOption(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("CreateMenuItemOption() = %v, want %v", got, want)
	}

	return nil
}
