import hre from "hardhat";
import { expect } from "chai";

describe("RandomnessGenerator", function () {
    let randomnessGenerator: any;
    let writer: any;
    let otherAccount: any;

    beforeEach(async function () {
        const [owner, _otherAccount] = await hre.ethers.getSigners();
        writer = owner;
        otherAccount = _otherAccount;

        const RandomnessGenerator = await hre.ethers.getContractFactory("RandomnessGenerator");
        randomnessGenerator = await RandomnessGenerator.deploy(writer.address);
    });

    describe("Posting Randomness", function () {
        it("Should allow the writer to post randomness", async function () {
            await randomnessGenerator.requestRandomness()
            const randomnessId = randomnessGenerator.randomnessId();
            await randomnessGenerator.postRandomness(randomnessId, 1234);
            const randomness = await randomnessGenerator.getRandomness(randomnessId);
            expect(randomness).to.equal(1234);
        });

        it("Should revert if non-writer tries to post randomness", async function () {
            const randomnessId = randomnessGenerator.randomnessId();
            await expect(
                randomnessGenerator.connect(otherAccount).postRandomness(randomnessId, 1234)
            ).to.be.revertedWith("Only writer can post randomness");
        });

        it("Should revert if randomness is already posted", async function () {
            await randomnessGenerator.requestRandomness()
            const randomnessId = randomnessGenerator.randomnessId();
            await randomnessGenerator.postRandomness(randomnessId, 1234);
            await expect(
                randomnessGenerator.postRandomness(randomnessId, 5678)
            ).to.be.revertedWith("Randomness already posted");
        });
    });

    describe("Fetching Randomness", function () {
        it("Should fetch the posted randomness correctly", async function () {
            await randomnessGenerator.requestRandomness()
            const randomnessId = randomnessGenerator.randomnessId();
            await randomnessGenerator.postRandomness(randomnessId, 1234);
            const randomness = await randomnessGenerator.getRandomness(randomnessId);
            expect(randomness).to.equal(1234);
        });

        it("Should revert if randomness has not been posted", async function () {
            const randomnessId = randomnessGenerator.randomnessId();
            await expect(
                randomnessGenerator.getRandomness(randomnessId)
            ).to.be.revertedWith("Randomness not posted");
        });
    });
});
