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

var order = &mookiespb.Order{}
var menu = &mookiespb.Menu{}
var empty = &mookiespb.Empty{}

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mMenuClient := mpb.NewMockMenuServiceClient(ctrl)

	mMenuClient.EXPECT().GetMenu(
		gomock.Any(),
		empty,
	).Return(menu, nil)

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

func testGetMenu(client mookiespb.MenuServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := menu

	req := empty
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
