const { ethers } = require("hardhat");



async function main() {
    const contractAddress = "0xAEb1cc59bD804DD3ADA20C405e61B4E05e908874";
    [deployer] = await ethers.getSigners();
    const coinFlip = await ethers.getContractAt("LotteryAgent", contractAddress);

    // Fetch the current gas price dynamically
    const { maxFeePerGas, maxPriorityFeePerGas } = await ethers.provider.getFeeData();


    // const deployOptions1 = {
    //     value: ethers.parseEther("1.0")
    // };
    // var chosenNumbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    //
    // try {
    //     for (let i = 0; i < 10; i++) {
    //         if (i === 9) {
    //             chosenNumbers = [1, 11, 12, 13, 14, 15, 16, 17, 18, 19];
    //         }
    //         console.log(`Purchasing ticket number ${i + 1}`);
    //         const tx = await coinFlip.purchaseTicket(chosenNumbers, deployOptions1);
    //         console.log(`Transaction ${i + 1} hash:`, tx.hash);
    //         await tx.wait();
    //     }
    //     console.log("Purchased 10 tickets successfully.");
    // } catch (error) {
    //     console.error("An error occurred during ticket purchases:", error);
    // }
    //
    // const deployOptions2 = {
    //     value: ethers.parseEther("50.0")
    // };
    // const tx3 = await coinFlip.depositSupply(deployOptions2);
    // // await tx3.wait()
    // await coinFlip.setDrawFrequency(0)
    // await coinFlip.prepareFinalizeDraw()
    // const randomnessGeneratorAddress = "0x45b625A16B1304f67488b5D9c2FBB8AA4d649CA5";
    // const randomnessGenerator = await ethers.getContractAt("RandomnessGenerator", randomnessGeneratorAddress);
    // const randId = await randomnessGenerator.randomnessId()
    // for (let i = Number(randId) - 10; i <= randId; i++) {
    //     try {
    //         const tx = await randomnessGenerator.postRandomness(i, 1);
    //         await tx.wait()
    //     } catch (err) {
    //         // console.log(err)
    //     }
    // }


    try {
        //
        // console.log(await coinFlip.curDraw())

        // await coinFlip.depositSupply(deployOptions)
        // console.log(await coinFlip.activeBalance())

        await coinFlip.setDrawFrequency(1000)
// 6000 = 100 = 60 + 40 = 3600 + 2400
        // await coinFlip.finalizeDraw();
        // console.log(await coinFlip.getDrawShortInfo(0))
    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
