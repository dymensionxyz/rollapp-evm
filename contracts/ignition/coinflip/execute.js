const { ethers } = require("hardhat");

async function main() {
    const contractAddress = "0xB8c4Ec444AD59a8f7Cc2c3e3F78bb0c367d2cE1d";

    const coinFlip = await ethers.getContractAt("CoinFlip", contractAddress);

    // Fetch the current gas price dynamically
    const { maxFeePerGas, maxPriorityFeePerGas } = await ethers.provider.getFeeData();

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('300000', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('200000', 'gwei'),
        value: ethers.parseEther("100000000.0")
    };

    try {
        const tx = await coinFlip.depositSupply(deployOptions)
        await tx.wait()
        // const tx = await coinFlip.startGame(0, deployOptions)
        // await tx.wait()
        // const tx1 = await coinFlip.completeGame()
        // await tx1.wait()
        // console.log(await coinFlip.getPlayerLastGameResult())
    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
