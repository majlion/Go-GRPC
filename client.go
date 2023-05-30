package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/example/calculator"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)

	// Call the Add RPC
	addReq := &calculator.AddRequest{Num1: 10, Num2: 5}
	addRes, err := client.Add(context.Background(), addReq)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	fmt.Printf("Add result: %d\n", addRes.Result)

	// Call the Subtract RPC
	subtractReq := &calculator.SubtractRequest{Num1: 10, Num2: 5}
	subtractRes, err := client.Subtract(context.Background(), subtractReq)
	if err != nil {
		log.Fatalf("Subtract failed: %v", err)
	}
	fmt.Printf("Subtract result: %d\n", subtractRes.Result)
}
