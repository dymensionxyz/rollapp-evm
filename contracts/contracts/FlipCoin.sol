// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RandomnessGenerator.sol";
import "./House.sol";

contract CoinFlip is House {
    enum CoinSide { HEADS, TAILS }
    enum GameStatus { NULL, PENDING, COMPLETED }

    struct Game {
        CoinSide playerChoice;
        uint256 randomnessId;
        GameStatus status;
        uint256 betAmount;
        bool won;
    }

    event GameCreated(address player, CoinSide choice);
    event GameCompleted(address player, bool won);

    RandomnessGenerator public randomnessGenerator;
    mapping(address => Game) public gameByPlayer;

    // Minimum bet amount
    uint256 public minBetAmount = 0.01 ether;

    // Maximum bet amount as a percentage of the house balance
    uint256 public maxBetAmountPercentage = 1;

    // House fee percentage on winnings
    uint256 public houseFeePercentage = 5;

    constructor(address _initialOwner, address _randomnessGenerator) House(_initialOwner) {
        randomnessGenerator = RandomnessGenerator(_randomnessGenerator);
    }

    function startGame(CoinSide choice) external payable {
        require(msg.value >= minBetAmount, "Bet amount is too low");
        require(msg.value <= calculateMaxBetAmount(), "Bet amount is too high");

        uint256 randomnessId = randomnessGenerator.requestRandomness();

        gameByPlayer[msg.sender] = Game({
            playerChoice: choice,
            randomnessId: randomnessId,
            status: GameStatus.PENDING,
            betAmount: msg.value,
            won: false
        });

        emit GameCreated(msg.sender, choice);
    }

    function completeGame() external {
        Game storage game = gameByPlayer[msg.sender];
        require(game.status == GameStatus.PENDING, "Game already completed");

        uint256 randomness = randomnessGenerator.getRandomness(game.randomnessId);
        CoinSide result = CoinSide(randomness % 2);

        game.status = GameStatus.COMPLETED;

        if (result == game.playerChoice) {
            game.won = true;
            uint256 reward = calculateReward(game.betAmount);
            addBalance(msg.sender, reward);
        }

        emit GameCompleted(msg.sender, game.won);
    }

    function getPlayerLastGameResult() external view returns (
        CoinSide playerChoice,
        GameStatus status,
        bool won
    ) {
        Game storage game = gameByPlayer[msg.sender];

        return (
            game.playerChoice,
            game.status,
            game.won
        );
    }

    function estimateReward(uint256 betAmount) external view returns (uint256) {
        return calculateReward(betAmount);
    }

    function calculateReward(uint256 betAmount) internal view returns (uint256) {
        return (2 * betAmount * (100 - houseFeePercentage)) / 100;
    }

    function calculateMaxBetAmount() internal view returns (uint256) {
        return calculateNonWithdrawalBalance() * maxBetAmountPercentage / 100;
    }

    function updateMinBetAmount(uint256 newAmount) external onlyOwner {
        minBetAmount = newAmount;
    }

    function updateHouseFeePercentage(uint256 newPercentage) external onlyOwner {
        houseFeePercentage = newPercentage;
    }
}
