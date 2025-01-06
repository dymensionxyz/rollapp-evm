const { ethers } = require("hardhat");

async function main() {
    const contractAddress = "0xFf2d86ddb73f35Db4a33c3d36a660E67C534B17d";

    const coinFlip = await ethers.getContractAt("CoinFlip", contractAddress);

    // Fetch the current gas price dynamically
    const { maxFeePerGas, maxPriorityFeePerGas } = await ethers.provider.getFeeData();

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('300000', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('200000', 'gwei'),
    };

    try {
        // await coinFlip.startGame(0, deployOptions);
        // await coinFlip.completeGame(deployOptions)

        const res = await coinFlip.getLastGameResult()
        console.log(res)

        // if (gameCreatedEvent) {
        //     const gameId = gameCreatedEvent.args?.gameId;
        //     console.log('Game ID:', gameId);
        // } else {
        //     console.log(await coinFlip.gameId());
        // }
        // for (let i = 55; i <= 55; i++) {
        //     try {
        //         const result = await coinFlip.getGameResult(i);
        //         console.log(`Game ID: ${i}`);
        //         console.log(`Player: ${result.player}`);
        //         console.log(`Player Choice: ${result.playerChoice === 0 ? 'HEADS' : 'TAILS'}`);
        //         console.log(`Status: ${result.status === 0 ? 'PENDING' : 'COMPLETED'}`);
        //         console.log(`Won: ${result.won}`);
        //         console.log('-----------------------------');
        //     } catch (error) {
        //         console.log(`Game ID: ${i} does not exist or another error occurred.`);
        //     }
        // }

        //
        // const tx1 = await coinFlip.gameId()
        // console.log(tx1)

        // const tx2 = await coinFlip.completeGame(tx1)
        // tx2.wait()
        //
        // console.log(tx2)

        // const tx3 = await coinFlip.getGameResult(1, deployOptions)
        // console.logg(tx3)

    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
