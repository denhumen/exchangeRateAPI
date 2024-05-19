package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"exchangeRateAPI/src/exchangeRateAPI/db"
	pb "exchangeRateAPI/src/exchangeRateAPI/proto/exchange_rate"
)

// Struct that implements the gRPC service interface
type ExchangeRateService struct {
	pb.UnimplementedExchangeRateServiceServer
}

// Struct to parse the response from the external exchange rate API
type ExchangeRateAPIResponse struct {
	Rates struct {
		UAH float64 `json:"UAH"`
	} `json:"rates"`
}

// GetCurrentRate fetches the current exchange rate from USD to UAH from an external openexchangerates API
func (s *ExchangeRateService) GetCurrentRate(ctx context.Context, req *pb.GetRateRequest) (*pb.GetRateResponse, error) {
	apiKey := os.Getenv("OPENEXCHANGERATES_API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing API key for Open Exchange Rates")
	}

	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&base=USD", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	var apiResponse ExchangeRateAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	log.Printf("Successfully returned rate")

	rate := apiResponse.Rates.UAH
	return &pb.GetRateResponse{Rate: rate}, nil
}

// SubscribeEmail subscribes an email to receive daily exchange rate updates.
func (s *ExchangeRateService) SubscribeEmail(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	subscriber := db.Subscriber{Email: email}
	if err := db.DB.Create(&subscriber).Error; err != nil {
		return nil, err
	}

	log.Printf("Subscription sucessfull for %s", email)

	return &pb.SubscribeResponse{Message: "Subscription successful for " + email}, nil
}

// UnsubscribeEmail unsubscribes an email not to receive daily exchange rate updates.
func (s *ExchangeRateService) UnsubscribeEmail(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if err := db.DB.Where("email = ?", email).Delete(&db.Subscriber{}).Error; err != nil {
		return nil, err
	}

	log.Printf("Unsubscription sucessfull for %s", email)

	return &pb.UnsubscribeResponse{Message: "Unsubscription successful for " + email}, nil
}
