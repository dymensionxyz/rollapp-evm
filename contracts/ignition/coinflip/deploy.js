// Import ethers from Hardhat
const { ethers } = require("hardhat");

async function main() {
    [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'), // Adjust as needed
        maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei'), // Adjust as needed
    };

    try {
        // Address of already deployed RandomnessGenerator contract
        const randomnessGeneratorAddress = "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0"; // Replace with actual address

        // Deploy CoinFlip contract
        const CoinFlip = await ethers.getContractFactory("CoinFlip", deployer);
        const coinFlip = await CoinFlip.deploy(randomnessGeneratorAddress, deployOptions);
        await coinFlip.waitForDeployment();
        console.log("CoinFlip deployed at:", coinFlip.target);

    } catch (error) {
        console.error("Error during CoinFlip deployment:", error);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Unhandled error:", error);
        process.exit(1);
    });
