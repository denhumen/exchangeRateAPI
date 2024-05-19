package main

import (
	"context"
	"log"
	"time"

	pb "exchangeRateAPI/src/exchangeRateAPI/proto/exchange_rate"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*10))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExchangeRateServiceClient(conn)

	// GetCurrentRate
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rateResp, err := client.GetCurrentRate(ctx, &pb.GetRateRequest{})
	if err != nil {
		log.Fatalf("Could not get rate: %v", err)
	}
	log.Printf("Current Rate: %f", rateResp.GetRate())

	// SubscribeEmail
	subscribeResp, err := client.SubscribeEmail(ctx, &pb.SubscribeRequest{Email: "mail05denis@gmail.com"})
	if err != nil {
		log.Fatalf("Could not subscribe email: %v", err)
	}
	log.Printf("Subscribe Response: %s", subscribeResp.GetMessage())

	// UnsubscribeEmail
	unsubscribeResp, err := client.UnsubscribeEmail(ctx, &pb.UnsubscribeRequest{Email: "mail007aris@gmail.com"})
	if err != nil {
		log.Fatalf("Could not unsubscribe email: %v", err)
	}
	log.Printf("Unsubscribe Response: %s", unsubscribeResp.GetMessage())
}
