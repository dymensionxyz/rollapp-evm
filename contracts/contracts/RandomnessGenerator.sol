// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./EventManager.sol";

contract RandomnessGenerator is EventManager {
    uint64 public randomnessId;
    mapping(uint256 => uint256) public randomnessJobs;
    address public writer;

    // Don't change the order of the entries in enum declaration. Backend relies on integer number under the enum
    enum EventType {
        RandomnessRequested
    }

    constructor(address _writer) EventManager(10240) {
        randomnessId = 0;
        writer = _writer;
    }

    function requestRandomness() external returns (uint256) {
        randomnessId += 1;
        bytes memory requestData = abi.encode(randomnessId);
        insertEvent(randomnessId, uint16(EventType.RandomnessRequested), requestData);
        return randomnessId;
    }

    function postRandomness(uint64 id, uint256 randomness) external {
        require(msg.sender == writer, "Only writer can post randomness");
        require(randomnessJobs[id] == 0, "Randomness already posted");

        randomnessJobs[id] = randomness;
        eraseEvent(randomnessId, uint16(EventType.RandomnessRequested));
    }

    function getRandomness(uint256 id) external view returns (uint256) {
        uint256 storedRandomness = randomnessJobs[id];
        require(storedRandomness != 0, "Randomness not posted");
        return storedRandomness;
    }

    struct UnprocessedRandomness {
        uint64 randomnessId;
    }

    function decodeUnprocessedRandomness(bytes memory data) internal pure returns (UnprocessedRandomness memory) {
        uint64 id;
        (id) = abi.decode(data, (uint64));
        return UnprocessedRandomness(id);
    }

    function getUnprocessedRandomness() external view returns (UnprocessedRandomness[] memory) {
        Event[] memory events = getEvents(uint16(EventType.RandomnessRequested));
        UnprocessedRandomness[] memory res = new UnprocessedRandomness[](events.length);
        for (uint64 i = 0; i < events.length; i++) {
            res[i] = decodeUnprocessedRandomness(events[i].data);
        }
        return res;
    }
}