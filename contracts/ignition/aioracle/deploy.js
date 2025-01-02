// Import ethers from Hardhat
const { ethers } = require("hardhat");

/**
 * Main function to deploy the AIOracle contract.
 * It performs the following steps:
 * 1. Retrieves the deployer account.
 * 2. Sets up deployment options for gas fees.
 * 3. Deploys the AIOracle contract with the deployer as the writer.
 * 4. Waits for the deployment to complete and logs the contract address.
 */
async function main() {
    [ deployer ] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'), // Adjust as needed
        maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei'), // Adjust as needed
    };

    try {
        const AIOracle = await ethers.getContractFactory("AIOracle", deployer);

        // Setup for RandomnessGenerator
        const writerAddress = deployer.address; // assuming deployer is the writer for testing

        const aiOracle = await AIOracle.deploy(writerAddress, deployOptions);

        await aiOracle.waitForDeployment();
        console.log("AIOracle deployed at:", aiOracle.target);
    } catch (error) {
        console.error("Error during AIOracle deployment:", error);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Unhandled error:", error);
        process.exit(1);
    });
