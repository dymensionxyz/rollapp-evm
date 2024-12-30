// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./EventManager.sol";

contract RandomnessGenerator is EventManager {
    uint256 public randomnessId;
    mapping(uint256 => uint256) public randomnessJobs;

    // Don't change the order of the entries in enum declaration. Backend relies on integer number under the enum
    enum EventType {
        RandomnessRequested,
        RandomnessProvided
    }

    constructor(address _writer) EventManager(10240, _writer) {
        randomnessId = 0;
    }

    function requestRandomness() external returns (uint256) {
        randomnessId += 1;
        bytes memory requestData = abi.encode(randomnessId);
        insertEvent(randomnessId, uint16(EventType.RandomnessRequested), requestData);
        return randomnessId;
    }

    function postRandomness(uint256 id, uint256 randomness) external {
        require(msg.sender == writer, "Only writer can post randomness");
        require(randomnessJobs[id] == 0, "Randomness already posted");

        randomnessJobs[id] = randomness;
    }

    function getRandomness(uint256 id) external view returns (uint256) {
        uint256 storedRandomness = randomnessJobs[id];
        require(storedRandomness != 0, "Randomness not posted");
        return storedRandomness;
    }
}