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
        const RandomnessGenerator = await ethers.getContractFactory("RandomnessGenerator", deployer);

        // Setup for RandomnessGenerator
        const writerAddress = deployer.address; // assuming deployer is the writer for testing

        const randomnessGenerator = await RandomnessGenerator.deploy(writerAddress, deployOptions);

        await randomnessGenerator.waitForDeployment();
        console.log("RandomnessGenerator deployed at:", randomnessGenerator.target);
    } catch (error) {
        console.error("Error during RandomnessGenerator deployment:", error);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Unhandled error:", error);
        process.exit(1);
    });
