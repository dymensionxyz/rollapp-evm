const { ethers } = require("hardhat");

async function main() {
    const contractAddress = "0x09d0647B434e6315f20AB0D6Cc87E1A274299b69";

    const coinFlip = await ethers.getContractAt("CoinFlip", contractAddress);

    // Fetch the current gas price dynamically
    const { maxFeePerGas, maxPriorityFeePerGas } = await ethers.provider.getFeeData();

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('300000', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('200000', 'gwei'),
    };

    try {

        const res = await coinFlip.getPlayerLastGameResult()
        console.log(res)
    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
