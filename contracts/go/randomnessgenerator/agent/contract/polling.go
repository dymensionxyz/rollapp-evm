package contract

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (a *RNGClient) ListenSmartContractEvents(ctx context.Context) <-chan []RandomnessGeneratorUnprocessedRandomness {
	prompts := make(chan []RandomnessGeneratorUnprocessedRandomness)

	ticker := time.NewTicker(a.config.PollInterval)

	go func() {
		defer close(prompts)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				a.logger.With("error", ctx.Err()).
					Info("Context done, exiting event loop")
				return

			case <-ticker.C:
				p, err := a.contractAPI.GetUnprocessedRandomness(&bind.CallOpts{Context: ctx})
				if err != nil {
					a.logger.Error("Error polling events from contract", "error", err)
					continue
				}

				prompts <- p
			}
		}
	}()

	return prompts
}

func (a *RNGClient) PostRandomness(ctx context.Context, randID uint64, randomness *big.Int) error {
	tx, err := a.contractAPI.PostRandomness(a.txAuth, randID, randomness)
	if err != nil {
		return fmt.Errorf("submit answer: %w", err)
	}
	err = a.waitForTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("wait for transaction: %w", err)
	}
	return nil
}

func (a *RNGClient) waitForTransaction(ctx context.Context, tx *types.Transaction) error {
	receipt, err := bind.WaitMined(ctx, a.ethClient, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction confirmation: %w", err)
	}

	if receipt.Status == 1 {
		return nil
	}

	revertReason, err := a.getRevertReason(ctx, tx, receipt)
	if err != nil {
		return fmt.Errorf("error getting revert reason: %w", err)
	}

	return fmt.Errorf("tx reverted: hash: %s, reason: %s", tx.Hash().String(), revertReason)
}

func (a *RNGClient) getRevertReason(ctx context.Context, tx *types.Transaction, receipt *types.Receipt) (string, error) {
	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
	}

	res, err := a.ethClient.CallContract(ctx, msg, receipt.BlockNumber)
	if err != nil {
		return "", fmt.Errorf("call contract: %w", err)
	}

	if len(res) < 4 {
		return "", fmt.Errorf("no revert reason found")
	}

	const errorMethodID = "0x08c379a0"
	if fmt.Sprintf("0x%x", res[:4]) != errorMethodID {
		return "", fmt.Errorf("no revert reason found")
	}

	abiError, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"reason","type":"string"}],"name":"Error","type":"function"}]`))
	if err != nil {
		return "", fmt.Errorf("parse revert reason ABI: %w", err)
	}

	var errorMsg string
	err = abiError.UnpackIntoInterface(&errorMsg, "Error", res[4:])
	if err != nil {
		return "", fmt.Errorf("unpack revert reason: %w", err)
	}

	return errorMsg, nil
}
