const { ethers } = require("hardhat");

async function main() {
    const PriceOracle = await ethers.getContractAt("PriceOracle", "0x9E6c736B3ddc1fa57ab97F402eC24B72056DB72F");

    const currentTimeUnixMs = Date.now();

    const deployOptions = {
        maxFeePerGas: ethers.parseUnits('30', 'gwei'),
        maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei'),
    };

    try {
        const tx = await PriceOracle.updatePrice(
            "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
            "0x0000000000000000000000000000000000000000",
            {
                price: 1000,
                proof: {
                    creationHeight: 123456,
                    creationTimeUnixMs: currentTimeUnixMs,
                    height: 12345678,
                    revision: 1,
                    merkleProof: "0x12345678"
                }
            },
            deployOptions,
        );

        await tx.wait();
        console.log("Price updated successfully: ", tx.hash);
    } catch (error) {
        console.error("Error:", error.message);
    }
}

main().catch((error) => {
    console.error("Error in script execution:", error);
    process.exitCode = 1;
});