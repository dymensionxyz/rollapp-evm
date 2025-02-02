const { ethers } = require("hardhat");

async function main() {
    const randomnessGeneratorAddress = "0x45b625A16B1304f67488b5D9c2FBB8AA4d649CA5";

    const randomnessGenerator = await ethers.getContractAt("RandomnessGenerator", randomnessGeneratorAddress);

    // Fetch the current gas price dynamically
    const { maxFeePerGas, maxPriorityFeePerGas } = await ethers.provider.getFeeData();

    // Ensure the gas fees are sufficient
    const MIN_FEE = 1041500000000; // Use the minimum fee reported in the error message
    const adjustedMaxFeePerGas = 104150000000000
    const adjustedMaxPriorityFeePerGas = 104150000000000

    const deployOptions = {
        maxFeePerGas: adjustedMaxFeePerGas,
        maxPriorityFeePerGas: adjustedMaxPriorityFeePerGas,
    };

    try {
        // Example: Request Randomness
        // const tx1 = await randomnessGenerator.requestRandomness(deployOptions);
        // console.log("Randomness request sent. ID:", tx1.hash.toString());
        // await tx1.wait();

        console.log(await randomnessGenerator.getRandomness(10))

        const tx3 = await randomnessGenerator.postRandomness(1, 10)

        // const updatedEvents = await randomnessGenerator.pollEvents(0);
        // console.log("Updated Events:", updatedEvents);

    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});
