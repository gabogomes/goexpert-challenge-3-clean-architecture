package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gabogomes/goexpert-challenge-3-clean-architecture/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// Create 3 orders
	for i := 1; i <= 3; i++ {
		orderID := fmt.Sprintf("order-%d", i)
		createOrder(client, orderID, float32(100*i), float32(10*i))
	}

	// Query all orders
	getAllOrders(client)
}

func createOrder(client pb.OrderServiceClient, id string, price float32, tax float32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create the order via gRPC
	req := &pb.CreateOrderRequest{
		Id:    id,
		Price: price,
		Tax:   tax,
	}
	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error while creating order: %v", err)
	}
	fmt.Printf("Created Order: ID: %s, Price: %.2f, Tax: %.2f, Final Price: %.2f\n", res.Id, res.Price, res.Tax, res.FinalPrice)
}

func getAllOrders(client pb.OrderServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Query all orders via gRPC
	req := &pb.EmptyRequest{}
	res, err := client.GetAllOrders(ctx, req)
	if err != nil {
		log.Fatalf("Error while getting all orders: %v", err)
	}
	fmt.Println("All Orders:")
	for _, order := range res.Orders {
		fmt.Printf("ID: %s, Price: %.2f, Tax: %.2f, Final Price: %.2f\n", order.Id, order.Price, order.Tax, order.FinalPrice)
	}
}
