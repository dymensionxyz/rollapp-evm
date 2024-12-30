package usecases

import "math/big"

type RandomnessGenerator interface {
	GenerateUInt256() (*big.Int, error)
}

