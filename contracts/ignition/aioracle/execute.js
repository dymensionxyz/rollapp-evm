const { ethers } = require("hardhat");

/**
 * Main function to execute the AIOracle contract interactions.
 * It performs the following steps:
 * 1. Retrieves the AIOracle contract instance.
 * 2. Checks if the prompter is whitelisted.
 * 3. If not whitelisted, whitelists the prompter.
 * 4. Submits a prompt to the AIOracle contract.
 * 5. Retrieves and logs the latest prompt ID.
 */
async function main() {
    const aiOracleAddr = "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0";
    // aiAgent address: 0x84ac82e5Ae41685D76021b909Db4f8E7C4bE279E
    // prompter address: 0x4781200f96791A81684b67D1777BC7Cc66EF5813
    [ aiAgent, prompter ] = await ethers.getSigners();

    console.log(`Executing AIOracle contract with: AIAgent: ${aiAgent.address}, Prompter: ${prompter.address}`);

    const AIOracle = await ethers.getContractAt("AIOracle", aiOracleAddr);

    const txOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei'),
    };

    try {
        // Check if prompter is whitelisted
        const isWhitelisted = await AIOracle.isWhitelisted(prompter.address);
        console.log("Prompter whitelisted: ", isWhitelisted);

        if (!isWhitelisted) {
            console.log("Whitelisting prompter...");
            const whitelistTx = await AIOracle.connect(aiAgent).addWhitelisted(prompter.address, txOptions);
            await whitelistTx.wait();
            console.log("Prompter successfully whitelisted!");
        }

        // Submit a prompt
        const promptTx = await AIOracle.connect(prompter).submitPrompt("Generate a random number between 1 and 10", txOptions);
        const receipt = await promptTx.wait();

        console.log("Prompt submitted. TX receipt: ", receipt);

        // Get the latest prompt ID
        const latestPromptID = await AIOracle.latestPromptId();
        console.log("Latest prompt ID: ", latestPromptID.toString());

    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
