// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RandomnessGenerator {
    event EventNewRandomnessRequest(uint256 randomnessId);
    event EventRandomnessProvided(uint256 randomnessId, uint256 randomnessValue);

    address public writer;
    uint256 public randomnessId;
    mapping(uint256 => uint256) public randomnessJobs;

    constructor(address _writer) {
        writer = _writer;
        randomnessId = 0;
    }

    function requestRandomness() external returns (uint256) {
        randomnessId += 1;
        emit EventNewRandomnessRequest(randomnessId);
        return randomnessId;
    }

    function postRandomness(uint256 id, uint256 randomness) external {
        require(msg.sender == writer, "Only writer can post randomness");
        require(randomnessJobs[id] == 0, "Randomness already posted");

        randomnessJobs[id] = randomness;
        emit EventRandomnessProvided(id, randomness);
    }

    function getRandomness(uint256 id) external view returns (uint256) {
        uint256 storedRandomness = randomnessJobs[id];
        require(storedRandomness != 0, "Randomness not posted");
        return storedRandomness;
    }
}