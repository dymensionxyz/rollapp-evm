const { ethers } = require("hardhat");

async function main() {
    const randomnessGeneratorAddress = "0x22A1E4163fbD0dc09C717B81AEEa83A68AD41451";

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
        const tx1 = await randomnessGenerator.requestRandomness(deployOptions);
        console.log("Randomness request sent. ID:", tx1.hash.toString());
        // await tx1.wait();

        // const tx2 = await randomnessGenerator.getRandomness(1)
        // console.log(tx2.toString())
        // const tx3 = await randomnessGenerator.postRandomness(10, 10)

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
