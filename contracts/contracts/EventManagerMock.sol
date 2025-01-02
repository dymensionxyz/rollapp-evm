// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./EventManager.sol";

contract EventManagerMock is EventManager {
    constructor(uint bufferSize) EventManager(bufferSize) {}

    function insertEventPublic(uint64 eventId, uint16 eventType, bytes memory data) public {
        insertEvent(eventId, eventType, data);
    }

    function eraseEventPublic(uint64 eventId, uint16 eventType) public {
        eraseEvent(eventId, eventType);
    }
}
