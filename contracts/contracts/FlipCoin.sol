// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RandomnessGenerator.sol";

contract CoinFlip {
    enum CoinSide { HEADS, TAILS }
    enum GameStatus { PENDING, COMPLETED }

    struct Game {
        address player;
        CoinSide playerChoice;
        uint256 randomnessId;
        GameStatus status;
        bool won;
    }

    event GameCreated(uint256 gameId, address player, CoinSide choice);
    event GameCompleted(uint256 gameId, address player, bool won);

    RandomnessGenerator public randomnessGenerator;
    uint256 public gameId;
    mapping(uint256 => Game) public games;

    constructor(address _randomnessGenerator) {
        randomnessGenerator = RandomnessGenerator(_randomnessGenerator);
        gameId = 0;
    }

    function createGame(CoinSide choice) external returns (uint256) {
        gameId += 1;

        uint256 randomnessId = randomnessGenerator.requestRandomness();

        games[gameId] = Game({
            player: msg.sender,
            playerChoice: choice,
            randomnessId: randomnessId,
            status: GameStatus.PENDING,
            won: false
        });

        emit GameCreated(gameId, msg.sender, choice);
        return gameId;
    }

    function completeGame(uint256 _gameId) external {
        Game storage game = games[_gameId];
        require(game.player != address(0), "Game does not exist");
        require(game.status == GameStatus.PENDING, "Game already completed");

        uint256 randomness = randomnessGenerator.getRandomness(game.randomnessId);
        CoinSide result = CoinSide(randomness % 2);

        game.status = GameStatus.COMPLETED;
        game.won = (result == game.playerChoice);

        emit GameCompleted(_gameId, game.player, game.won);
    }

    function getGameResult(uint256 _gameId) external view returns (
        address player,
        CoinSide playerChoice,
        GameStatus status,
        bool won
    ) {
        Game storage game = games[_gameId];
        require(game.player != address(0), "Game does not exist");

        return (
            game.player,
            game.playerChoice,
            game.status,
            game.won
        );
    }
}
