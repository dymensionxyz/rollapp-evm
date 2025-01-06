package external

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"math/big"
	"randomnessgenerator/agent/config"

	"github.com/go-resty/resty/v2"
)

type RNGServiceClient struct {
	logger *slog.Logger
	http   *resty.Client
	config config.RNGConfig
}

// NewRNGServiceClient creates and returns a new instance of RNGServiceClient.
func NewRNGServiceClient(logger *slog.Logger, config config.RNGConfig) *RNGServiceClient {
	return &RNGServiceClient{
		logger: logger,
		http: resty.New().
			SetBaseURL(config.BaseURL).
			SetHeader("Content-Type", "application/json"),
		config: config,
	}
}

type RandomnessResponse struct {
	RequestID  uuid.UUID `json:"requestID"`
	Randomness *big.Int  `json:"randomness"`
}

type rawRandomnessResponse struct {
	RequestID  string `json:"requestID"`
	Randomness string `json:"randomness"`
}

func (RandomnessResponse) IsResponse() {}

func (c *RNGServiceClient) GetRandomness(ctx context.Context) (RandomnessResponse, error) {
	var rawResp rawRandomnessResponse

	resp, err := c.http.R().
		SetContext(ctx).
		SetResult(&rawResp).
		Get("/generate")

	if err != nil {
		return RandomnessResponse{}, fmt.Errorf("failed to create message: %w", err)
	}
	if resp.IsError() {
		return RandomnessResponse{}, fmt.Errorf("failed to create message: %s", resp.Error())
	}

	uid, err := uuid.Parse(rawResp.RequestID)
	if err != nil {
		return RandomnessResponse{}, fmt.Errorf("invalid RequestID, not a valid UUID: %v", err)
	}

	randomNumber := new(big.Int)
	randomNumber, ok := randomNumber.SetString(rawResp.Randomness, 10)
	if !ok {
		return RandomnessResponse{}, fmt.Errorf("invalid Randomness value, not a valid uint256")
	}

	response := RandomnessResponse{
		RequestID:  uid,
		Randomness: randomNumber,
	}

	return response, nil
}
