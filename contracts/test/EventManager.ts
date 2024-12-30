import hre from "hardhat";
import { expect } from "chai";

describe("EventManagerMock", function () {
    let eventManagerMock: any;
    let owner: any;
    let otherAccount: any;
    const bufferSize = 5;

    beforeEach(async function () {
        const [deployer, _otherAccount] = await hre.ethers.getSigners();
        owner = deployer;
        otherAccount = _otherAccount;

        const EventManagerMock = await hre.ethers.getContractFactory("EventManagerMock");
        eventManagerMock = await EventManagerMock.deploy(bufferSize);
    });

    describe("Inserting Events", function () {
        it("Should allow events to be inserted until buffer is full", async function () {
            const eventId = 1;
            const eventType = 1;
            const data = "0x1234";

            for (let i = 0; i < bufferSize; i++) {
                await eventManagerMock.insertEventPublic(eventId + i, eventType, data);
            }

            await expect(
                eventManagerMock.insertEventPublic(eventId + bufferSize, eventType, data)
            ).to.be.revertedWith("Event buffer is full");
        });

        it("Should revert if inserting an event when buffer is full", async function () {
            const eventId = 1;
            const eventType = 1;
            const data = "0x1234";

            for (let i = 0; i < bufferSize; i++) {
                await eventManagerMock.insertEventPublic(eventId + i, eventType, data);
            }

            await expect(
                eventManagerMock.insertEventPublic(eventId + bufferSize, eventType, data)
            ).to.be.revertedWith("Event buffer is full");
        });
    });

    describe("Erasing Events", function () {
        it("Should allow events to be erased", async function () {
            const eventId = 1;
            const eventType = 1;
            const data = "0x1234";

            await eventManagerMock.insertEventPublic(eventId, eventType, data);
            await eventManagerMock.eraseEventPublic(eventId, eventType);

            const events = await eventManagerMock.pollEvents(eventType);
            expect(events.length).to.equal(0);
        });

        it("Should revert if trying to erase an event that doesn't exist", async function () {
            const eventId = 999;
            const eventType = 1;

            await expect(
                eventManagerMock.eraseEventPublic(eventId, eventType)
            ).to.be.revertedWith("Event does not exist");
        });
    });

    describe("Polling Events", function () {
        it("Should allow fetching events by event type", async function () {
            const eventId = 1;
            const eventType = 1;
            const data = "0x1234";

            await eventManagerMock.insertEventPublic(eventId, eventType, data);
            const events = await eventManagerMock.pollEvents(eventType);

            expect(events.length).to.equal(1);
            expect(events[0].eventId).to.equal(eventId);
            expect(events[0].eventType).to.equal(eventType);
            expect(events[0].data).to.equal(data);
        });

        it("Should return an empty array for an event type with no events", async function () {
            const eventType = 999;

            const events = await eventManagerMock.pollEvents(eventType);

            expect(events.length).to.equal(0);
        });
    });
});
