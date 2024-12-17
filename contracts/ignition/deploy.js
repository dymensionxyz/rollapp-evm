// Import ethers from Hardhat
const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();

    if (!deployer) {
        throw new Error("No deployer account found. Check your configuration.");
    }

    const PriceOracle = await ethers.getContractFactory("PriceOracle");

    const expirationOffset = 3600; // 1 hour in seconds
    const assetInfos = [
        {
            localNetworkName: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
            oracleNetworkName: "OracleDenom1",
            localNetworkPrecision: 18,
        },
        {
            localNetworkName: "0x0000000000000000000000000000000000000000",
            oracleNetworkName: "OracleDenom2",
            localNetworkPrecision: 8,
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