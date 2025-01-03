// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {AIOracle} from "./AIOracle.sol";
import {House} from "./House.sol";

contract AIGambling is House {
    struct Bet {
        uint64 promptId;
        uint256 amount;
        string guessedNumber;
        bool resolved;
        bool won;
    }

    event BetPlaced(address indexed user, uint256 betAmount, uint256 guessedNumber);
    event BetResult(address indexed user, uint256 guessedNumber, uint256 correctNumber, bool won, uint256 reward);

    AIOracle private aiOracle;

    mapping(address => Bet) public bets;

    uint256 public minBetAmount = 0.01 ether;
    uint256 public maxBetAmountPercentage = 1; // Max bet amount is 1% of the house balance

    uint256 public houseFeePercentage = 5; // House takes 5% of winnings

    string public constant PROMPT = "Generate a number between 1 and 10, inclusive";

    constructor(address _aiOracle) {
        require(_aiOracle != address(0), "Invalid AIOracle address");
        aiOracle = AIOracle(_aiOracle);
    }

    function placeBet(string memory guessedNumber) external payable {
        require(bets[msg.sender].amount == 0 || bets[msg.sender].resolved, "Resolve your current bet first");
        require(msg.value >= minBetAmount, "Bet amount is too low");
        require(msg.value <= address(this).balance * (100 - maxBetAmountPercentage) / 100, "Bet amount is too high");

        uint64 promptId = aiOracle.submitPrompt(PROMPT);

        bets[msg.sender] = Bet({
            promptId: promptId,
            amount: msg.value,
            guessedNumber: guessedNumber,
            resolved: false,
            won: false
        });

        emit BetPlaced(msg.sender, msg.value, guessedNumber);
    }

    function resolveBet() external {
        Bet storage bet = bets[msg.sender];
        require(!bet.resolved, "Bet already resolved");

        string memory correctNumber = aiOracle.getAnswer(bet.promptId);

        bet.resolved = true;

        if (equalStrings(bet.guessedNumber, correctNumber)) {
            // We generate a number between 1 and 10, inclusive. So the win chance is 1 of 10.
            // The reward multiplier is 10 minus the house fee percentage (F).
            // So the reward is 10B - F, where B is a bet amount.
            uint256 reward = (10 * bet.amount * (100 - houseFeePercentage)) / 100;
            bet.won = true;
            addBalance(msg.sender, reward);
            emit BetResult(msg.sender, bet.guessedNumber, correctNumber, true, reward);
        } else {
            emit BetResult(msg.sender, bet.guessedNumber, correctNumber, false, 0);
        }
    }

    function equalStrings(string memory a, string memory b) internal pure returns (bool) {
        if (bytes(a).length != bytes(b).length) {
            return false;
        }

        for (uint i = 0; i < bytes(a).length; i++) {
            if (bytes(a)[i] != bytes(b)[i]) {
                return false;
            }
        }
        return true;
    }

    function updateMinBetAmount(uint256 newAmount) external onlyOwner {
        minBetAmount = newAmount;
    }

    function updateHouseFeePercentage(uint256 newPercentage) external onlyOwner {
        require(newPercentage <= 10, "House fee too high");
        houseFeePercentage = newPercentage;
    }
}
