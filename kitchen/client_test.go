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

func TestSubscribeToOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := mpb.NewMockOrderService_SubscribeToOrdersClient(ctrl)

	stream.EXPECT().SendMsg(
		order,
	).Return(nil)

	stream.EXPECT().Recv().Return(order, nil)
	stream.EXPECT().CloseSend().Return(nil)

	mOrderClient := mpb.NewMockOrderServiceClient(ctrl)
	mOrderClient.EXPECT().SubscribeToOrders(
		gomock.Any(),
		&mookiespb.Empty{},
	).Return(stream, nil)

	if err := testSubscribeToOrders(mOrderClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func TestCompleteOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mOrderClient := mpb.NewMockOrderServiceClient(ctrl)
	mOrderClient.EXPECT().CompleteOrder(
		gomock.Any(),
		&mookiespb.CompleteOrderRequest{Id: 1},
	).Return(&mookiespb.CompleteOrderResponse{Result: "Order marked as complete"}, nil)

	if err := testCompleteOrder(mOrderClient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testSubscribeToOrders(client mookiespb.OrderServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &mookiespb.Empty{}

	stream, err := client.SubscribeToOrders(ctx, req)
	if err != nil {
		return err
	}

	if err := stream.SendMsg(order); err != nil {
		return err
	}

	if err := stream.CloseSend(); err != nil {
		return err
	}

	want := order
	got, err := stream.Recv()
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, want)
	}

	return nil
}

func testCompleteOrder(client mookiespb.OrderServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	want := &mookiespb.CompleteOrderResponse{Result: "Order marked as complete"}

	req := &mookiespb.CompleteOrderRequest{Id: 1}
	got, err := client.CompleteOrder(ctx, req)
	if err != nil {
		return err
	}

	if !proto.Equal(got, want) {
		return fmt.Errorf("CompleteOrder() = %v, want %v", got, want)
	}

	return nil
}
