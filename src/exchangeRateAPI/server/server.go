package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"exchangeRateAPI/src/exchangeRateAPI/db"
	pb "exchangeRateAPI/src/exchangeRateAPI/proto/exchange_rate"
	"exchangeRateAPI/src/exchangeRateAPI/services"

	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Init()

	log.Println("Starting server...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()

	exchangeRateService := &services.ExchangeRateService{}

	pb.RegisterExchangeRateServiceServer(grpcServer, exchangeRateService)

	reflection.Register(grpcServer)

	log.Println("Server is running on port 8080...")

	sendDailyExchangeRateEmail()
	c := cron.New()
	c.AddFunc("@daily", func() {
		sendDailyExchangeRateEmail()
	})
	c.Start()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8080: %v", err)
	}
}

// Addictional function to process daily email sending
func sendDailyExchangeRateEmail() {
	exchangeRateService := &services.ExchangeRateService{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rateResp, err := exchangeRateService.GetCurrentRate(ctx, &pb.GetRateRequest{})
	if err != nil {
		log.Printf("Failed to get current exchange rate: %v", err)
		return
	}

	services.SendCurrentRateToSubscribers(rateResp.Rate)
}
