const { ethers } = require("hardhat");

async function main() {
    const randomnessGeneratorAddress = "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0";

    const randomnessGenerator = await ethers.getContractAt("RandomnessGenerator", randomnessGeneratorAddress);

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('3000', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('2000', 'gwei'),
    };

    try {
        const storageSlot = await randomnessGenerator.randomnessId(); // Start at storage slot 0
        // const value = await randomnessGenerator.randomnessId(randomnessGeneratorAddress, storageSlot);
        // console.log(`Value at slot ${storageSlot}:`, value);
        console.log(`${storageSlot}`)

        const tx = await randomnessGenerator.requestRandomness(deployOptions);
        console.log("Randomness request sent. ID:", tx.hash.toString());
        await tx.wait();
        // //
        // const randomnessValue = 123456789;
        // const writerAddress = "0xYourWriterAddress";
        //
        // const signer = await ethers.getSigner(writerAddress);
        // const tx = await randomnessGenerator.connect(signer).postRandomness(randomnessId, randomnessValue, deployOptions);
        //
        // await tx.wait();
        // console.log(`Randomness for ID ${randomnessId.toString()} posted successfully`);
        //
        // const storedRandomness = await randomnessGenerator.getRandomness(3);
        // console.log("Stored randomness:", storedRandomness.toString());

    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
