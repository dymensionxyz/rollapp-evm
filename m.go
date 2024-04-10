package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	pathPref = "/home/ubuntu/go/bin/" // TODO: config

	rollapd        = pathPref + "rollapp-evm"
	dymd           = pathPref + "dymd"
	owner          = "rol-user"
	userRoll       = "ralex"
	userHub        = "halex"
	totalAmount    = 1000000000
	transferAmount = "1alxx"

	keyringBackendFlag = "--keyring-backend"
)

var fundAmount = fmt.Sprintf("%dalxx", totalAmount)

type Balance struct {
	Balances []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balances"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	// setup accounts
	if err := setupAccounts(); err != nil {
		log.Fatalf("setupAccounts(): %s", err)
	}

	// get the sender balance before transfer
	senderBalance, err := balanceOfName(rollapd, userRoll)
	if err != nil {
		log.Fatalf("balanceOfName(userRoll): %s", senderBalance)
	}
	log.Printf("Sender balance before: %s\n", senderBalance)

	// get the receiver balance before transfer
	receiverBalance, err := balanceOfName(dymd, userHub)
	if err != nil {
		log.Fatalf("balanceOfName(userHub): %s", receiverBalance)
	}
	log.Printf("Receiver balance before: %s\n", receiverBalance)

	for i := 0; i < totalAmount; i++ {
		time.Sleep(time.Millisecond * 1000)
		// do the ibc transfer
		log.Printf("Transferring 1alxx from %s to %s...; i=%d\n", userRoll, userHub, i)
		output, err := ibcTransfer(rollapd, dymd, userRoll, userHub, transferAmount)
		if err != nil {
			log.Printf("ibcTransfer(rollapd, dymd, userRoll, userHub, transferAmount): %s", output)
		}
		senderBalance, err := balanceOfName(rollapd, userRoll)
		if err != nil {
			log.Printf("balanceOfName(userRoll): %s", senderBalance)
		}
		receiverBalance, err := balanceOfName(dymd, userHub)
		if err != nil {
			log.Printf("balanceOfName(userHub): %s", receiverBalance)
		}
		log.Printf("Transferred; i=%d, sender balance: %s; receiver balance: %s\n", i, senderBalance, receiverBalance)
	}

	// get the sender balance after transfer
	senderBalance, err = balanceOfName(rollapd, userRoll)
	if err != nil {
		log.Fatalf("balanceOfName(userRoll): %s", senderBalance)
	}
	log.Printf("Sender balance after: %s\n", senderBalance)

	// get the receiver balance after transfer
	receiverBalance, err = balanceOfName(dymd, userHub)
	if err != nil {
		log.Fatalf("balanceOfName(userHub): %s", receiverBalance)
	}
	log.Printf("Receiver balance after: %s\n", receiverBalance)
}

// rollapp-evm tx ibc-transfer transfer transfer channel-0 dym1g3w9nvkg70h0mhhmqtm2wutzvfnc7gwzyyf5xt 1rax --from alex --keyring-backend test --broadcast-mode block -y

func ibcTransfer(senderBin, receiverBin, senderName, receiverName, amount string) (string, error) {
	receiverAddress, err := address(receiverBin, receiverName)
	if err != nil {
		return receiverAddress, err
	}
	cmd := exec.Command(
		senderBin, "tx", "ibc-transfer", "transfer", "transfer",
		"channel-0",
		receiverAddress,
		amount,
		"--from", senderName,
		keyringBackendFlag, "test",
		"--broadcast-mode", "block",
		// "--fees", "20000000000000adym",
		"-y")

	output, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		output = eerr.Stderr
	}
	return string(output), err
}

func setupAccounts() error {
	exists, err := keyExists(rollapd, userRoll)
	if err != nil {
		return fmt.Errorf("keyExists(rollapd, userRoll): %w", err)
	}
	if !exists {
		// Add rollapp key
		output, err := addKey(rollapd, userRoll)
		log.Printf("addKey(rollapd, userRoll): %s\n", output)
		if err != nil {
			return fmt.Errorf("addKey(rollapd, userRoll): %w", err)
		}
		// fund the sender account
		output, err = transferToName(rollapd, owner, userRoll, fundAmount)
		log.Printf("transferToName(owner, userRoll, fundAmount): %s\n", output)
		if err != nil {
			return fmt.Errorf("transferToName(owner, userRoll, fundAmount): %w", err)
		}
	}

	exists, err = keyExists(dymd, userHub)
	if err != nil {
		return fmt.Errorf("keyExists(dymd, userHub): %w", err)
	}
	if !exists {
		// Add hub key
		output, err := addKey(dymd, userHub)
		log.Printf("addKey(dymd, userHub): %s\n", output)
		if err != nil {
			return fmt.Errorf("addKey(dymd, userHub): %w", err)
		}
	}
	return nil
}

func balance(bin, address string) (string, error) {
	cmd := exec.Command(
		bin, "query", "bank", "balances", address, "-o", "json")
	output, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		output = eerr.Stderr
	}

	var bal Balance
	if err := json.Unmarshal(output, &bal); err != nil {
		return "", err
	}

	if len(bal.Balances) == 0 {
		return "", nil
	}

	return fmt.Sprintf("%s%s", bal.Balances[0].Amount, bal.Balances[0].Denom), err
}

func addKey(bin, name string) (string, error) {
	cmd := exec.Command(
		bin, "keys", "add",
		name, keyringBackendFlag, "test")
	output, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		output = eerr.Stderr
	}
	return string(output), err
}

func keyExists(bin, name string) (bool, error) {
	cmd := exec.Command(
		bin, "keys", "show",
		name, keyringBackendFlag, "test")
	out, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		out = eerr.Stderr
		log.Printf("keyExists(bin, name): %s\n", out)
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func transfer(bin, senderName, receiverAddress, amount string) (string, error) {
	cmd := exec.Command(
		bin, "tx", "bank", "send",
		senderName,
		receiverAddress,
		amount,
		keyringBackendFlag, "test",
		"--broadcast-mode", "block",
		"-y")
	output, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		output = eerr.Stderr
	}
	return string(output), err
}

// dymd tx bank send bob-genesis --keyring-backend test $(dymd keys show alice-genesis --keyring-backend test -a) 5000ibc/22953CC8562B0A279979E6B49564DA2CC59DFFF074A94E48140AA9FF89ADEE6C --broadcast-mode block --node https://rpc.hwpd.noisnemyd.xyz:443 --chain-id dymension_1405-1 -y

func address(bin, name string) (string, error) {
	cmd := exec.Command(
		bin, "keys", "show",
		name,
		keyringBackendFlag, "test",
		"-a")
	output, err := cmd.Output()
	if eerr, ok := err.(*exec.ExitError); ok {
		output = eerr.Stderr
	} else {
		return strings.Trim(string(output), "\n"), nil
	}
	return string(output), err

}

func balanceOfName(bin, name string) (string, error) {
	addr, err := address(bin, name)
	if err != nil {
		return addr, err
	}
	return balance(bin, addr)
}

func transferToName(bin, senderName, receiverName, amount string) (string, error) {
	receiverAddr, err := address(bin, receiverName)
	if err != nil {
		return receiverAddr, err
	}
	return transfer(bin, senderName, receiverAddr, amount)
}
