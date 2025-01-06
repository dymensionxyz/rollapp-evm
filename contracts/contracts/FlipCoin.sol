// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RandomnessGenerator.sol";

contract CoinFlip {
    enum CoinSide { HEADS, TAILS }
    enum GameStatus { NULL, PENDING, COMPLETED }

    struct Game {
        CoinSide playerChoice;
        uint256 randomnessId;
        GameStatus status;
        bool won;
    }

    event GameCreated(address player, CoinSide choice);
    event GameCompleted(address player, bool won);

    RandomnessGenerator public randomnessGenerator;
    mapping(address => Game) public gameByPlayer;

    constructor(address _randomnessGenerator) {
        randomnessGenerator = RandomnessGenerator(_randomnessGenerator);
    }

    function startGame(CoinSide choice) external {
        uint256 randomnessId = randomnessGenerator.requestRandomness();

        gameByPlayer[msg.sender] = Game({
            playerChoice: choice,
            randomnessId: randomnessId,
            status: GameStatus.PENDING,
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
        game.won = (result == game.playerChoice);

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
}
