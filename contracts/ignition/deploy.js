async function main() {
    const [deployer] = await ethers.getSigners();

    console.log("Deploying contract with the account:", deployer.address);

    const PriceOracle = await ethers.getContractFactory("PriceOracle");

    // Constructor parameters
    const expirationOffset = 3600; // For example, 1 hour
    const assetInfos = [
        {
            localNetworkName: "0xTokenAddress1",
            oracleNetworkName: "OracleDenom1",
            localNetworkPrecision: 18,
        },
        // Add more assets as needed
    ];
    const boundThreshold = 500; // For example

    const priceOracle = await PriceOracle.deploy(
        expirationOffset,
        assetInfos,
        boundThreshold
    );

    console.log("Price Oracle deployed at:", priceOracle.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });