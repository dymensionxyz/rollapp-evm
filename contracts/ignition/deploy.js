// Import ethers from Hardhat
const { ethers } = require("hardhat");

async function main() {
    [ deployer ] = await ethers.getSigners();
    console.log("Deploying PriceOracle contract with the account:", deployer.address);

    const PriceOracle = await ethers.getContractFactory("PriceOracle", deployer);

    const expirationOffset = 3600; // 1 hour in seconds
    const assetInfos = [
        {
            localNetworkName: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
            oracleNetworkName: "WBTC",
            localNetworkPrecision: 8,
        },
        {
            localNetworkName: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
            oracleNetworkName: "USDC",
            localNetworkPrecision: 6,
        }
    ];
    const boundThreshold = 500;

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'), // Adjust as needed
        maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei'), // Adjust as needed
    };

    try {
        const priceOracle = await PriceOracle.deploy(
            expirationOffset,
            assetInfos,
            boundThreshold,
            deployOptions
        );

        await priceOracle.waitForDeployment();

        console.log("Price Oracle deployed at:", priceOracle.target);
    } catch (error) {
        console.error("Error during deployment:", error);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });