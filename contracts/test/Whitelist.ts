import hre from "hardhat";
import { expect } from "chai";

describe("Whitelist", function () {
    let whitelist: any;
    let owner: any;
    let otherAccount: any;
    let anotherAccount: any;

    beforeEach(async function () {
        const [deployer, _otherAccount, _anotherAccount] = await hre.ethers.getSigners();
        owner = deployer;
        otherAccount = _otherAccount;
        anotherAccount = _anotherAccount;

        const Whitelist = await hre.ethers.getContractFactory("Whitelist");
        whitelist = await Whitelist.deploy(owner.address);
    });

    describe("Adding to Whitelist", function () {
        it("Should allow the owner to add an address to the whitelist", async function () {
            await whitelist.addWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.true;
        });

        it("Should emit an event when an address is added to the whitelist", async function () {
            await expect(whitelist.addWhitelisted(otherAccount.address))
                .to.emit(whitelist, "AddWhitelisted")
                .withArgs(otherAccount.address);
        });

        it("Should revert if a non-owner tries to add an address", async function () {
            await expect(
                whitelist.connect(otherAccount).addWhitelisted(anotherAccount.address)
            ).to.be.revertedWithCustomError(whitelist, "OwnableUnauthorizedAccount");
        });
    });

    describe("Removing from Whitelist", function () {
        it("Should allow the owner to remove an address from the whitelist", async function () {
            await whitelist.addWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.true;

            await whitelist.removeWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.false;
        });

        it("Should emit an event when an address is removed from the whitelist", async function () {
            await whitelist.addWhitelisted(otherAccount.address);

            await expect(whitelist.removeWhitelisted(otherAccount.address))
                .to.emit(whitelist, "RemoveWhitelisted")
                .withArgs(otherAccount.address);
        });

        it("Should revert if a non-owner tries to remove an address", async function () {
            await whitelist.addWhitelisted(otherAccount.address);
            await expect(
                whitelist.connect(otherAccount).removeWhitelisted(otherAccount.address)
            ).to.be.revertedWithCustomError(whitelist, "OwnableUnauthorizedAccount");
        });
    });

    describe("Whitelist Checks", function () {
        it("Should correctly return true for whitelisted addresses", async function () {
            await whitelist.addWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.true;
        });

        it("Should correctly return false for non-whitelisted addresses", async function () {
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.false;
        });
    });

    describe("Access Control", function () {
        it("Should only allow the owner to manage the whitelist", async function () {
            await whitelist.addWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.true;

            await expect(
                whitelist.connect(otherAccount).addWhitelisted(anotherAccount.address)
            ).to.be.revertedWithCustomError(whitelist, "OwnableUnauthorizedAccount");

            await whitelist.removeWhitelisted(otherAccount.address);
            expect(await whitelist.isWhitelisted(otherAccount.address)).to.be.false;
        });
    });
});
