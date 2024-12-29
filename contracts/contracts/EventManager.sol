// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract EventManager {
    struct Event {
        uint256 eventId;
        uint16 eventType;
        bytes data;
    }

    mapping(uint => Event[]) private _eventsByType;
    uint private _eventBufferSize;
    address public writer;

    constructor(uint bufferSize, address _writer) {
        _eventBufferSize = bufferSize;
        writer = _writer;
    }

    function insertEvent(uint256 eventId, uint16 eventType, bytes memory data) internal {
        require(_eventsByType[eventType].length < _eventBufferSize, "Event buffer is full");
        _eventsByType[eventType].push(Event(eventId, eventType, data));
    }

    function eraseEvents(uint256[] memory eventIds, uint16 eventType) external {
        require(msg.sender == writer, "Only writer can erase events");

        Event[] storage events = _eventsByType[eventType];
        uint256 i = 0;

        while (i < events.length) {
            bool found = false;
            for (uint256 j = 0; j < eventIds.length; j++) {
                if (events[i].eventId == eventIds[j]) {
                    events[i] = events[events.length - 1];
                    events.pop();
                    found = true;
                    break;
                }
            }
            if (!found) {
                i++;
            }
        }
    }

    function pollEvents(uint16 eventType) external view returns (Event[] memory) {
        return _eventsByType[eventType];
    }
}
