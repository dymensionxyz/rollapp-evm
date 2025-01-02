# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.ts
```

## Generate Go bindings

Generate ABI for the contract:
```shell
solc @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ --abi contracts/AIOracle.sol -o build --overwrite
```

Generate Go bindings:
```shell
abigen --abi build/AIOracle.abi --pkg main --type AIOracle --out go/aiagent/AIOracle.sol.go
```
