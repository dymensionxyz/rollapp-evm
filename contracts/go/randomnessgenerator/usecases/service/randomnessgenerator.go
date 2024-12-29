package service

import (
	"crypto/rand"
	"example1/usecases"
	"fmt"
	"math/big"
)

type RandomnessGenerator struct {
}

func NewRandomnessGenerator() (usecases.RandomnessGenerator, error) {
	return &RandomnessGenerator{}, nil
}

func (g *RandomnessGenerator) GenerateUInt256() (*big.Int, error) {
	maxInt := new(big.Int).Lsh(big.NewInt(1), 256)
	maxInt.Sub(maxInt, big.NewInt(1))

	randomNumber, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random uint256: %v", err)
	}

	return randomNumber, nil
}
