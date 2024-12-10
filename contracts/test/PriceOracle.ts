import { loadFixture } from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { expect } from "chai";
import hre from "hardhat";

describe("PriceOracle", function () {
  // Reusable Constants
  const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
  const DEFAULT_PRICE_EXPIRY = 3600; // 1 hour

  // Fixture to reuse the same setup in every test
  async function deployPriceOracleFixture(targetBlockNumber: number = 150) {
    const [owner, otherAccount] = await hre.ethers.getSigners();

    const PriceOracle = await hre.ethers.getContractFactory("PriceOracle");
    const priceOracle = await PriceOracle.deploy(DEFAULT_PRICE_EXPIRY);

    // Mine additional blocks to reach the target block number
    const currentBlockNumber = await hre.ethers.provider.getBlockNumber();
    if (currentBlockNumber < targetBlockNumber) {
      const blocksToMine = targetBlockNumber - currentBlockNumber;
      for (let i = 0; i < blocksToMine; i++) {
        await hre.ethers.provider.send("evm_mine", []);
      }
    }

    return { priceOracle, owner, otherAccount };
  }

  // Helper function to initialize the contract
  async function initializePriceOracle(priceOracle: any) {
    await priceOracle.initialize();
    expect(await priceOracle.initialized()).to.equal(true);
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { priceOracle, owner } = await loadFixture(deployPriceOracleFixture);
      expect(await priceOracle.owner()).to.equal(owner.address);
    });

    it("Should start with initialized as false", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      expect(await priceOracle.initialized()).to.equal(false);
    });
  });

  describe("Initialization", function () {
    it("Should initialize successfully by the owner", async function () {
      const { priceOracle, owner } = await loadFixture(deployPriceOracleFixture);
      await expect(priceOracle.connect(owner).initialize())
        .to.emit(priceOracle, "OracleInitialized")
        .withArgs(owner.address);
      expect(await priceOracle.initialized()).to.equal(true);
    });

    it("Should fail if a non-owner tries to initialize", async function () {
      const { priceOracle, otherAccount } = await loadFixture(deployPriceOracleFixture);
      await expect(priceOracle.connect(otherAccount).initialize())
        .to.be.revertedWith("PriceOracle: caller is not the owner");
    });

    it("Should fail if already initialized", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);
      await expect(priceOracle.initialize())
        .to.be.revertedWith("PriceOracle: already initialized");
    });
  });

  describe("Ownership", function () {
    it("Should transfer ownership successfully", async function () {
      const { priceOracle, owner, otherAccount } = await loadFixture(deployPriceOracleFixture);
      await expect(priceOracle.transferOwnership(otherAccount.address))
        .to.emit(priceOracle, "OwnershipTransferred")
        .withArgs(owner.address, otherAccount.address);
      expect(await priceOracle.owner()).to.equal(otherAccount.address);
    });

    it("Should fail if a non-owner tries to transfer ownership", async function () {
      const { priceOracle, otherAccount } = await loadFixture(deployPriceOracleFixture);
      await expect(priceOracle.connect(otherAccount).transferOwnership(otherAccount.address))
        .to.be.revertedWith("PriceOracle: caller is not the owner");
    });

    it("Should fail if transferring ownership to the zero address", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await expect(priceOracle.transferOwnership(ZERO_ADDRESS))
        .to.be.revertedWith("PriceOracle: new owner is the zero address");
    });
  });

  describe("UpdatePrice", function () {
    it("Should reject expired prices", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);

      const block = await hre.ethers.provider.getBlock("latest");
      const expiredPriceProof = {
        creationHeight: block!.number - 100,
        creationTimeUnixMs: block!.timestamp * 1000 - 2 * 60 * 60 * 1000,
        height: block!.number,
        revision: 1,
        merkleProof: "0x42",
      };

      const expirecPriceWithProof = {
        price: 1000,
        proof: expiredPriceProof,
      };

      await expect(
        priceOracle.updatePrice(
          "0x0000000000000000000000000000000000000001",
          "0x0000000000000000000000000000000000000002",
          expirecPriceWithProof,
        )
      ).to.be.revertedWith("PriceOracle: price proof expired");
    });

    it("should reject if update with an older price", async function() {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);

      const block = await hre.ethers.provider.getBlock("latest");
      const priceProof = {
        creationHeight: block!.number,
        creationTimeUnixMs: block!.timestamp * 1000,
        height: block!.number,
        revision: 1,
        merkleProof: "0x42",
      };

      const priceWithProof = {
        price: 1000,
        proof: priceProof,
      };

      await expect(
          priceOracle.updatePrice(
              "0x0000000000000000000000000000000000000001",
              "0x0000000000000000000000000000000000000002",
              priceWithProof,
          )
      ).not.to.be.revertedWith("PriceOracle: price proof expired");

      await hre.ethers.provider.send("evm_mine", []); // move one block

      const olderPriceProof = {
        creationHeight: block!.number - 1,
        creationTimeUnixMs: block!.timestamp * 1000 - 60 * 1000, // 1 minute ago
        height: block!.number,
        revision: 1,
        merkleProof: "0x42",
      };

      const olderPriceWithProof = {
        price: 1000,
        proof: olderPriceProof,
      };

      await expect(
          priceOracle.updatePrice(
              "0x0000000000000000000000000000000000000001",
              "0x0000000000000000000000000000000000000002",
              olderPriceWithProof,
          )
      ).to.be.revertedWith("PriceOracle: cannot update with an older price");
    })
  });
});