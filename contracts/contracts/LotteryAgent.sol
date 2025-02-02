// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./RandomnessGenerator.sol";

contract LotteryAgent is Ownable {
    // Struct to represent a lottery ticket

    struct Ticket {
        address player;
        bool[] chosenNumbers;
        bool claimed;
        bool winner;
        uint256 id; // unique idx in draw
    }

    // Struct to represent a lottery draw
    struct Draw {
        uint[] randomnessIDs;
        bool[] winningNumbers; // the representation of bitset. example: uint [2, 3, 5] == bool [0, 0, 1, 1, 0, 1]
        uint totalWinnings;
        uint winnersCount;
        uint ticketRevenue;
        uint stackersPoolDistributionRatio;
        Ticket[] tickets;
        bool prepareFinalizeCalled;
    }


    function getDraw(uint256 idx) public view returns (Draw memory) {
        require(idx <= drawHistory.length, "Draw index out of bounds");
        if (idx == drawHistory.length) {
            return curDraw;
        } else {
            return drawHistory[idx];
        }
    }

    function getDrawShortInfo(uint256 idx) external view returns (uint256 totalWinnings, uint256[] memory winningNumbers, uint256 winnersCount, uint256 ticketCount) {
        Draw memory draw = getDraw(idx);

        uint256[] memory winningIndices = new uint256[](draw.winningNumbers.length);
        uint256 count = 0;

        for (uint256 i = 0; i < draw.winningNumbers.length; i++) {
            if (draw.winningNumbers[i]) {
                winningIndices[count] = i;
                count++;
            }
        }

        uint256[] memory result = new uint256[](count);
        for (uint256 j = 0; j < count; j++) {
            result[j] = winningIndices[j];
        }

        totalWinnings = draw.totalWinnings;
        winningNumbers = result;
        winnersCount = draw.winnersCount;

        return (totalWinnings, winningNumbers, winnersCount, draw.tickets.length);
    }

    function depositSupply() external payable {
        curDraw.totalWinnings += msg.value;
    }

    receive() external payable {
        curDraw.totalWinnings += msg.value;
    }

    function activeBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function withdrawSupply(uint256 amount, address receiver) external onlyOwner {
        require(amount <= activeBalance(), "House: insufficient non-withdrawal balance");
        payable(receiver).transfer(amount);
    }

    mapping (uint => mapping (address => uint[])) ticketIdsByUserByDrawId;

    uint constant public NUMBER_TO_CHOOSE = 10;
    uint constant public NUMBERS_COUNT = 20;

    RandomnessGenerator public randomnessGenerator;

    // Lottery parameters
    uint public ticketPrice = 1 * 10 ** 18; // 1 Ether by default (adjustable)
    uint public drawFrequency = 1 days; // Default to one draw per day
    uint public drawBeginTime;
    uint public stackersPoolDistributionRatio = 50; // 50% to prize pool, 50% to staking pool

    uint public ticketCounter = 0;
    Draw public curDraw;

    Draw[] public drawHistory;

    event TicketPurchased(address indexed player, uint ticketId, uint[] chosenNumbers);
    event DrawFinalized(uint indexed drawId, bool[] winningNumbers);
    event PrizeClaimed(address indexed player, uint prizeAmount);

    constructor(address _owner, address _randomnessGenerator) Ownable(_owner) {
        randomnessGenerator = RandomnessGenerator(_randomnessGenerator);
        drawBeginTime = block.timestamp;
        curDraw.stackersPoolDistributionRatio = stackersPoolDistributionRatio;
    }

    function validateTicket(uint[] memory _chosenNumbers) internal pure  {
        require(_chosenNumbers.length == NUMBER_TO_CHOOSE, "You must pick 10 numbers");

        bool[] memory numberPresence = new bool[](NUMBERS_COUNT);
        for (uint i = 0; i < _chosenNumbers.length; i++) {
            uint number = _chosenNumbers[i];
            require(number < NUMBERS_COUNT, "Number out of range");
            require(numberPresence[number] == false, "Duplicate number found");
            numberPresence[number] = true;
        }
    }

    function toSet(uint[] memory _chosenNumbers) internal pure returns (bool[] memory) {
        bool[] memory set = new bool[](NUMBERS_COUNT);
        for (uint i = 0; i < _chosenNumbers.length; i++) {
            uint number = _chosenNumbers[i];
            set[number] = true;
        }
        return set;
    }

    function purchaseTicket(uint[] calldata _chosenNumbers) external payable {
        require(curDraw.prepareFinalizeCalled == false, "Can't purchase tickets to draw, which was prepared to finish");
        validateTicket(_chosenNumbers);
        require(msg.value == ticketPrice, "Incorrect Ether value sent. Ticket costs different prize");

        uint ticketId = curDraw.tickets.length;

        curDraw.tickets.push(
            Ticket({
                player: msg.sender,
                chosenNumbers: toSet(_chosenNumbers),
                claimed: false,
                winner : false,
                id: curDraw.tickets.length
            })
        );

        ticketIdsByUserByDrawId[drawHistory.length][msg.sender].push(curDraw.tickets.length - 1);

        uint stackersFee = ticketPrice * curDraw.stackersPoolDistributionRatio / 100;
        // TODO: SEND TO STACKERS
        uint ticketRevenue = ticketPrice - stackersFee;
        curDraw.ticketRevenue += ticketRevenue;

        emit TicketPurchased(msg.sender, ticketId, _chosenNumbers);
    }

    // Function to check if all randomness has been posted
    function allRandomnessPostedForCurDraw() public view returns (bool) {
        for (uint i = 0; i < curDraw.randomnessIDs.length; i++) {
            if (randomnessGenerator.getRandomness(curDraw.randomnessIDs[i]) == 0) {
                return false;
            }
        }
        return true;
    }

    function prepareFinalizeDraw() external {
        require(block.timestamp >= drawBeginTime + drawFrequency, "It's not time for the draw yet");
        require(curDraw.prepareFinalizeCalled == false, "prepareFinalizeCalled was already called");

        curDraw.prepareFinalizeCalled = true;

        for (uint i = 0; i < NUMBER_TO_CHOOSE; i++) {
            curDraw.randomnessIDs.push(randomnessGenerator.requestRandomness());
        }
    }

    // New function to generate winning numbers
    function generateWinningNumbers(uint[] memory randomNumbers) internal pure returns (bool[] memory) {
        uint[] memory lotteryDrum = new uint[](NUMBERS_COUNT);

        // Initialize the array with numbers from 0 to NUMBERS_COUNT-1
        for (uint i = 0; i < lotteryDrum.length; i++) {
            lotteryDrum[i] = i;
        }

        bool[] memory winningNumbers = new bool[](NUMBERS_COUNT);
        for (uint i = 0; i < winningNumbers.length; i++) {
            winningNumbers[i] = false;
        }

        uint numbersLeft = lotteryDrum.length;
        // Remove elements from lotteryDrum
        for (uint i = 0; i < randomNumbers.length; i++) {
            uint winningNumberIdx = randomNumbers[i] % numbersLeft;
            winningNumbers[lotteryDrum[winningNumberIdx]] = true;
            lotteryDrum[winningNumberIdx] = lotteryDrum[numbersLeft - 1];
            numbersLeft--; // Simulate pop operation
        }

        return winningNumbers;
    }

    function finalizeDraw() external {
        require(curDraw.prepareFinalizeCalled == true, "prepareFinalizeCalled wasn't called, call it first");
        require(block.timestamp >= drawBeginTime + drawFrequency, "It's not time for the draw yet");

        // Check if all randomness has been posted
        require(allRandomnessPostedForCurDraw(), "Not all randomness has been fulfilled");

        uint[] memory randomnessIDs = curDraw.randomnessIDs;
        uint[] memory randomNumbers = new uint[](randomnessIDs.length);
        for (uint i = 0; i < randomnessIDs.length; i++) {
            randomNumbers[i] = randomnessGenerator.getRandomness(randomnessIDs[i]);
        }

        // Generate the winning numbers using the new function
        bool[] memory winningNumbers = generateWinningNumbers(randomNumbers);

        curDraw.winningNumbers = winningNumbers;

        // Check for winners
        for (uint i = 0; i < curDraw.tickets.length; i++) {
            if (checkIfWinner(curDraw.tickets[i].chosenNumbers, curDraw.winningNumbers)) {
                curDraw.winnersCount++;
                curDraw.tickets[i].winner = true;
            }
        }

        emit DrawFinalized(drawHistory.length, winningNumbers);
        drawHistory.push(curDraw);

        uint prevDrawWinnersCount = curDraw.winnersCount;
        uint prevDrawTotalWinnings = curDraw.totalWinnings;
        uint prevDrawTicketRevenue = curDraw.ticketRevenue;

        // Handle the next draw's winnings
        resetCurDraw();
        curDraw.totalWinnings += prevDrawTicketRevenue;
        if (prevDrawWinnersCount == 0) {
            curDraw.totalWinnings += prevDrawTotalWinnings;
        }

        curDraw.stackersPoolDistributionRatio = stackersPoolDistributionRatio;
        drawBeginTime = block.timestamp;
    }

    function resetCurDraw() internal {
        delete curDraw.randomnessIDs;
        delete curDraw.winningNumbers;
        delete curDraw.tickets;

        curDraw.totalWinnings = 0;
        curDraw.winnersCount = 0;
        curDraw.ticketRevenue = 0;
        curDraw.stackersPoolDistributionRatio = 0;
        curDraw.prepareFinalizeCalled = false;
    }

    function claimPrize(uint drawId, uint ticketId) external {
        Ticket storage ticket = drawHistory[drawId].tickets[ticketId];
        require(ticket.player == msg.sender, "You are not the owner of this ticket");
        require(!ticket.claimed, "Prize already claimed");
        require(ticket.winner, "The ticket is not winning one!");

        uint prizeAmount = drawHistory[drawId].totalWinnings / drawHistory[drawId].winnersCount;
        payable(msg.sender).transfer(prizeAmount);
        ticket.claimed = true;
        emit PrizeClaimed(msg.sender, prizeAmount);
    }

    function checkIfWinner(bool[] memory chosenNumbers, bool[] memory winningNumbers) internal pure returns (bool) {
        if (chosenNumbers.length != winningNumbers.length) {
            return false;
        }

        for (uint i = 0; i < chosenNumbers.length; i++) {
            if (chosenNumbers[i] != winningNumbers[i]) {
                return false;
            }
        }

        return true;
    }

    // Admin functions to adjust contract parameters
    function setTicketPrice(uint newTicketPrice) external onlyOwner { // TODO: change it to onlyGov
        ticketPrice = newTicketPrice;
    }

    function setDrawFrequency(uint newFrequency) external onlyOwner {
        drawFrequency = newFrequency;
    }

    function setStackersPoolDistributionRatio(uint newRatio) external onlyOwner {
        stackersPoolDistributionRatio = newRatio;
    }

    function setNewRngAddress(address addr) external onlyOwner {
        randomnessGenerator = RandomnessGenerator(addr);
    }

    // Public function for users to see their ticket IDs
    function getUserTickets(uint drawId, address user) external view returns (Ticket[] memory) {
        Ticket[] memory drawTickets = getDraw(drawId).tickets;

        uint[] memory ticketIds = ticketIdsByUserByDrawId[drawId][user];
        Ticket[] memory res = new Ticket[](ticketIds.length);
        for (uint i = 0; i < ticketIds.length; ++i) {
            res[i] = drawTickets[ticketIds[i]];
        }
        return res;
    }

    function getCurDrawTotalWinnings() public view returns (uint256) {
        return curDraw.totalWinnings;
    }

    function getCurDrawRemainingTime() public view returns (uint256) {
        return drawBeginTime + drawFrequency;
    }

    function getDrawCount() external view returns (uint256) {
        return drawHistory.length + 1;
    }
}
