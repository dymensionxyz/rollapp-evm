// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

abstract contract EventManager {
    struct Event {
        uint256 eventId;
        uint16 eventType;
        bytes data;
    }

    mapping(uint => Event[]) private _eventsByType;
    uint private _eventBufferSize;

    constructor(uint bufferSize) {
        _eventBufferSize = bufferSize;
    }

    function insertEvent(uint256 eventId, uint16 eventType, bytes memory data) internal {
        require(_eventsByType[eventType].length < _eventBufferSize, "Event buffer is full");
        _eventsByType[eventType].push(Event(eventId, eventType, data));
    }

    function eraseEvent(uint256 eventId, uint16 eventType) internal {
        Event[] storage events = _eventsByType[eventType];
        for (uint256 i = 0; i < events.length; i++) {
            if (events[i].eventId == eventId) {
                events[i] = events[events.length - 1];
                events.pop();
                return;
            }
        }

        revert("Event with provided ID not found");
    }

    function pollEvents(uint16 eventType) external view returns (Event[] memory) {
        return _eventsByType[eventType];
    }
}
