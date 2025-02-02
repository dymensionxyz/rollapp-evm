// Import ethers from Hardhat
const { ethers } = require("hardhat");

async function main() {
    [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'), // Adjust as needed
        maxPriorityFeePerGas: ethers.parseUnits('20', 'gwei'), // Adjust as needed
    };

    try {
        // Address of already deployed RandomnessGenerator contract
        const randomnessGeneratorAddress = "0x45b625A16B1304f67488b5D9c2FBB8AA4d649CA5"; // Replace with actual address

        // Deploy CoinFlip contract
        const CoinFlip = await ethers.getContractFactory("LotteryAgent", deployer);
        const writerAddress = deployer.address; // assuming deployer is the writer for testing
        const coinFlip = await CoinFlip.deploy(writerAddress, randomnessGeneratorAddress);
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
