package abi

import (
	_ "embed"
	"encoding/json"
	"github.com/evmos/evmos/v12/x/evm/abi"
)

var (
	//go:embed price_feed.json
	priceFeedJson []byte

	PriceFeedCpcInfo abi.CustomPrecompiledContractInfo
)

func init() {
	var err error

	err = json.Unmarshal(priceFeedJson, &PriceFeedCpcInfo)
	if err != nil {
		panic(err)
	}
	PriceFeedCpcInfo.Name = "Price Feed"
}
