package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"math/big"
	"net/http"
)

type RandomnessGenerator struct {
	apiURL string
}

type RawRandomnessResponse struct {
	RequestID  string `json:"requestID"`
	Randomness string `json:"randomness"`
}

type RandomnessResponse struct {
	RequestID  string   `json:"requestID"`
	Randomness *big.Int `json:"randomness"`
}

func NewRandomnessGenerator(apiURL string) (*RandomnessGenerator, error) {
	return &RandomnessGenerator{apiURL: apiURL}, nil
}

func (g *RandomnessGenerator) GenerateUInt256() (*RandomnessResponse, error) {
	resp, err := http.Get(g.apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var rawResponse RawRandomnessResponse
	err = json.Unmarshal(body, &rawResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response JSON: %v", err)
	}

	if _, err := uuid.Parse(rawResponse.RequestID); err != nil {
		return nil, fmt.Errorf("invalid RequestID, not a valid UUID: %v", err)
	}

	randomNumber := new(big.Int)
	randomNumber, ok := randomNumber.SetString(rawResponse.Randomness, 10)
	if !ok {
		return nil, fmt.Errorf("invalid Randomness value, not a valid uint256")
	}

	response := &RandomnessResponse{
		RequestID:  rawResponse.RequestID,
		Randomness: randomNumber,
	}

	return response, nil
}
