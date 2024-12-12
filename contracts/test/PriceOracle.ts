import { loadFixture } from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { expect } from "chai";
import hre from "hardhat";

describe("PriceOracle", function () {
  // Reusable Constants
  const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
  const DEFAULT_PRICE_EXPIRY = 3600; // 1 hour
  const BTC_ERC20_ADDRESS = "0x1234567890123456789012345678901234567890";
  const USDC_ERC20_ADDRESS = "0x2170Ed0880ac9A755fd29B2688956BD959F933F8";

  const SCALE_FACTOR = 10n ** 18n;

  // Fixture to reuse the same setup in every test
  async function deployPriceOracleFixture(targetBlockNumber: number = 150) {
    const [owner, otherAccount] = await hre.ethers.getSigners();

    const assetInfos = [
      {
        localNetworkName: BTC_ERC20_ADDRESS, // Token address
        oracleNetworkName: "btc",    // Corresponding token name in oracle
        localNetworkPrecision: 8    // Decimal precision for the token
      },
      {
        localNetworkName: USDC_ERC20_ADDRESS,
        oracleNetworkName: "usdc",
        localNetworkPrecision: 18
      }
    ];

    const PriceOracle = await hre.ethers.getContractFactory("PriceOracle");
    const priceOracle = await PriceOracle.deploy(DEFAULT_PRICE_EXPIRY, assetInfos);

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
    it("Should correctly map asset info", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);

      expect(await priceOracle.localNetworkToOracleNetworkDenoms(BTC_ERC20_ADDRESS)).to.equal("btc");
      expect(await priceOracle.precissionMapping(BTC_ERC20_ADDRESS)).to.equal(8);
      expect(await priceOracle.localNetworkToOracleNetworkDenoms(USDC_ERC20_ADDRESS)).to.equal("usdc");
      expect(await priceOracle.precissionMapping(USDC_ERC20_ADDRESS)).to.equal(18);
    });

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
          BTC_ERC20_ADDRESS,
          USDC_ERC20_ADDRESS,
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
              BTC_ERC20_ADDRESS,
              "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
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
              BTC_ERC20_ADDRESS,
              "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
              olderPriceWithProof,
          )
      ).to.be.revertedWith("PriceOracle: cannot update with an older price");
    })

    it("Should reject when base token is not registered", async function () {
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

      // Using an unregistered base token address
      const unregisteredAddress = "0x9999999999999999999999999999999999999999";

      await expect(
          priceOracle.updatePrice(
              unregisteredAddress,
              USDC_ERC20_ADDRESS,
              priceWithProof
          )
      ).to.be.revertedWith("PriceOracle: base denom not found in local_network_to_oracle_network_denoms");
    });

    it("Should reject when quote token is not registered", async function () {
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

      // Using registered base but unregistered quote token address
      const unregisteredAddress = "0x9999999999999999999999999999999999999999";

      await expect(
          priceOracle.updatePrice(
              BTC_ERC20_ADDRESS,
              unregisteredAddress,
              priceWithProof
          )
      ).to.be.revertedWith("PriceOracle: quote denom not found in local_network_to_oracle_network_denoms");
    });

    describe("Price Precision Adjustment", function () {
      it("Should correctly adjust price when quote precision > base precision", async function () {
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

        const baseToken = BTC_ERC20_ADDRESS; // 8 decimals
        const quoteToken = USDC_ERC20_ADDRESS; // 18 decimals
        const price = 1000n;

        await priceOracle.updatePrice(
            baseToken,
            quoteToken,
            {
              price: price,
              proof: priceProof,
            }
        );

        const storedPrice = await priceOracle.getPrice(baseToken, quoteToken);
        expect(storedPrice.price).to.equal(price * SCALE_FACTOR * (10n ** 10n)); // Adjusted for 18 - 6 decimals
      });

      it("Should correctly adjust price when quote precision < base precision", async function () {
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

        const baseToken = USDC_ERC20_ADDRESS; // 18 decimals
        const quoteToken = BTC_ERC20_ADDRESS; // 8 decimals
        const price = 1000n;

        await priceOracle.updatePrice(
            baseToken,
            quoteToken,
            {
              price: price,
              proof: priceProof,
            }
        );

        const storedPrice = await priceOracle.getPrice(baseToken, quoteToken);
        expect(storedPrice.price).to.equal(price * SCALE_FACTOR / (10n ** 10n)); // Adjusted for 6 - 18 decimals
      });
    });
  });

  describe("GetPrice", function () {
    it("should return the correct price when the price is set and has not expired", async function () {
      const { priceOracle, owner } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);

      const base = BTC_ERC20_ADDRESS;
      const quote = USDC_ERC20_ADDRESS;
      const price = 1000n;
      const currentBlock = await hre.ethers.provider.getBlock("latest");
      const creationTimeUnixMs = currentBlock!.timestamp * 1000;
      const expiration = creationTimeUnixMs + DEFAULT_PRICE_EXPIRY * 1000;

      const priceProof = {
        creationHeight: currentBlock!.number,
        creationTimeUnixMs: creationTimeUnixMs,
        height: currentBlock!.number,
        revision: 1,
        merkleProof: "0x42",
      };

      const priceWithProof = {
        price: price,
        proof: priceProof,
      };

      await expect(priceOracle.updatePrice(base, quote, priceWithProof))
          .to.emit(priceOracle, "PriceUpdated")
          .withArgs(base, quote, price);

      const fetchedPrice = await priceOracle.getPrice(base, quote);
      expect(fetchedPrice.price).to.equal(price * 10n ** (18n - 8n)); // Adjusted for precision
      expect(fetchedPrice.is_inverse).to.be.false;

      const inverseFetchedPrice = await priceOracle.getPrice(quote, base);
      expect(inverseFetchedPrice.price).to.equal(price * 10n ** (18n - 8n));
      expect(inverseFetchedPrice.is_inverse).to.be.true;
    });

    it("should revert if the price does not exist", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);

      const base = BTC_ERC20_ADDRESS; // 0x1234567890123456789012345678901234567890
      const quote = USDC_ERC20_ADDRESS; // 0x2170Ed0880ac9A755fd29B2688956BD959F933F8

      await expect(priceOracle.getPrice(base, quote)).to.be.revertedWith(
          "PriceOracle: price not found"
      );
    });

    it("should revert if the price has expired", async function () {
      const { priceOracle } = await loadFixture(deployPriceOracleFixture);
      await initializePriceOracle(priceOracle);

      const base = BTC_ERC20_ADDRESS;
      const quote = USDC_ERC20_ADDRESS;
      const price = 1000n;

      const currentBlock = await hre.ethers.provider.getBlock("latest");
      const creationTimeUnixMs = currentBlock!.timestamp * 1000;

      const priceProof = {
        creationHeight: currentBlock!.number,
        creationTimeUnixMs: creationTimeUnixMs,
        height: currentBlock!.number,
        revision: 1,
        merkleProof: "0x42",
      };

      const priceWithProof = {
        price: price,
        proof: priceProof,
      };

      await expect(priceOracle.updatePrice(base, quote, priceWithProof))
          .to.emit(priceOracle, "PriceUpdated")
          .withArgs(base, quote, price);

      // Advance time to expire the price
      await hre.network.provider.send("evm_increaseTime", [DEFAULT_PRICE_EXPIRY + 1]);
      await hre.network.provider.send("evm_mine", []);

      await expect(priceOracle.getPrice(base, quote)).to.be.revertedWith(
          "PriceOracle: price expired"
      );
    });
  });
});

