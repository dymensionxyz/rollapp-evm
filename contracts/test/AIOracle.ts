import { loadFixture } from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { expect } from "chai";
import hre from "hardhat";

describe("AIOracle", function () {
    // Fixture to deploy and set up the contract
    async function deployAIOracleFixture() {
        const [owner, aiAgent, prompter1, prompter2] = await hre.ethers.getSigners();

        const AIOracle = await hre.ethers.getContractFactory("AIOracle");
        const aiOracle = await AIOracle.deploy(aiAgent.address);

        return { aiOracle, owner, aiAgent, prompter1, prompter2 };
    }

    describe("Deployment", function () {
        it("Should set the correct AI agent on deployment", async function () {
            const { aiOracle, aiAgent } = await loadFixture(deployAIOracleFixture);
            expect(await aiOracle.aiAgent()).to.equal(aiAgent.address);
        });

        it("Should initialize with no prompts", async function () {
            const { aiOracle } = await loadFixture(deployAIOracleFixture);
            expect(await aiOracle.latestPromptId()).to.equal(0);
        });
    });

    describe("Whitelisting", function () {
        it("Should allow AI agent to whitelist a prompter", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address))
                .to.emit(aiOracle, "AddWhitelisted")
                .withArgs(prompter1.address);

            expect(await aiOracle.isWhitelistedPrompter(prompter1.address)).to.equal(true);
        });

        it("Should not allow non-AI agent to whitelist a prompter", async function () {
            const { aiOracle, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.addWhitelistAddress(prompter1.address)).to.be.revertedWith(
                "AIOracle: caller is not the AI agent"
            );
        });

        it("Should allow AI agent to remove a prompter from the whitelist", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await expect(aiOracle.connect(aiAgent).removeWhitelistAddress(prompter1.address))
                .to.emit(aiOracle, "RemoveWhitelisted")
                .withArgs(prompter1.address);

            expect(await aiOracle.isWhitelistedPrompter(prompter1.address)).to.equal(false);
        });

        it("Should not allow removing a non-whitelisted address", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(aiAgent).removeWhitelistAddress(prompter1.address)).to.be.revertedWith(
                "AIOracle: address not whitelisted"
            );
        });
    });

    describe("Prompt Submission", function () {
        it("Should allow a whitelisted prompter to submit a prompt", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await expect(aiOracle.connect(prompter1).submitPrompt("What is the capital of France?"))
                .to.emit(aiOracle, "PromptSubmitted")
                .withArgs(1, "What is the capital of France?");

            expect(await aiOracle.latestPromptId()).to.equal(1);
        });

        it("Should not allow a non-whitelisted prompter to submit a prompt", async function () {
            const { aiOracle, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(prompter1).submitPrompt("What is the capital of France?")).to.be.revertedWith(
                "AIOracle: caller is not a whitelisted prompter"
            );
        });

        it("Should revert when submitting an empty prompt", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await expect(aiOracle.connect(prompter1).submitPrompt("")).to.be.revertedWith("AIOracle: prompt cannot be empty");
        });
    });

    describe("Answer Submission", function () {
        it("Should allow the AI agent to submit an answer for a valid prompt", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await aiOracle.connect(prompter1).submitPrompt("What is the capital of France?");
            const promptId = await aiOracle.latestPromptId();

            await expect(aiOracle.connect(aiAgent).submitAnswer(promptId, "Paris"))
                .to.emit(aiOracle, "AnswerSubmitted")
                .withArgs(promptId, "Paris");

            const answer = await aiOracle.getAnswer(promptId);
            expect(answer).to.equal("Paris");
        });

        it("Should not allow a non-AI agent to submit an answer", async function () {
            const { aiOracle, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(prompter1).submitAnswer(1, "Paris")).to.be.revertedWith(
                "AIOracle: caller is not the AI agent"
            );
        });

        it("Should not allow an answer for an invalid prompt ID", async function () {
            const { aiOracle, aiAgent } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(aiAgent).submitAnswer(1, "Paris")).to.be.revertedWith(
                "AIOracle: invalid prompt ID"
            );
        });

        it("Should not allow an empty answer", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await aiOracle.connect(prompter1).submitPrompt("What is the capital of France?");
            const promptId = await aiOracle.latestPromptId();

            await expect(aiOracle.connect(aiAgent).submitAnswer(promptId, "")).to.be.revertedWith(
                "AIOracle: answer cannot be empty"
            );
        });

        it("Should not allow multiple answers for the same prompt", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await aiOracle.connect(prompter1).submitPrompt("What is the capital of France?");
            const promptId = await aiOracle.latestPromptId();

            await aiOracle.connect(aiAgent).submitAnswer(promptId, "Paris");

            await expect(aiOracle.connect(aiAgent).submitAnswer(promptId, "London")).to.be.revertedWith(
                "AIOracle: answer already exists"
            );
        });
    });

    describe("Answer Retrieval", function () {
        it("Should allow anyone to retrieve an answer for a valid prompt ID", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await aiOracle.connect(aiAgent).addWhitelistAddress(prompter1.address);

            await aiOracle.connect(prompter1).submitPrompt("What is the capital of France?");
            const promptId = await aiOracle.latestPromptId();

            await aiOracle.connect(aiAgent).submitAnswer(promptId, "Paris");

            const answer = await aiOracle.getAnswer(promptId);
            expect(answer).to.equal("Paris");
        });

        it("Should revert when retrieving an answer for a non-existent prompt ID", async function () {
            const { aiOracle } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.getAnswer(1)).to.be.revertedWith("AIOracle: answer does not exist");
        });
    });

    describe("Ownership Transfer", function () {
        it("Should allow the AI agent to transfer ownership", async function () {
            const { aiOracle, aiAgent, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(aiAgent).transferOwnership(prompter1.address))
                .to.emit(aiOracle, "OwnershipTransferred")
                .withArgs(aiAgent.address, prompter1.address);

            expect(await aiOracle.aiAgent()).to.equal(prompter1.address);
        });

        it("Should not allow a non-AI agent to transfer ownership", async function () {
            const { aiOracle, prompter1 } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(prompter1).transferOwnership(prompter1.address)).to.be.revertedWith(
                "AIOracle: caller is not the AI agent"
            );
        });

        it("Should revert when transferring ownership to the zero address", async function () {
            const { aiOracle, aiAgent } = await loadFixture(deployAIOracleFixture);

            await expect(aiOracle.connect(aiAgent).transferOwnership(hre.ethers.ZeroAddress)).to.be.revertedWith(
                "AIOracle: new owner is the zero address"
            );
        });
    });
});
