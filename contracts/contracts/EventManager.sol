// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract EventManager {
    struct Event {
        uint64 eventId;
        uint16 eventType;
        bytes data;
    }

    struct EventEntries {
        Event[] data;
        mapping(uint64 => uint64) dataIdxByEventId;
    }

    mapping(uint16 => EventEntries) private _eventsByType;

    uint private _eventBufferSize;

    constructor(uint bufferSize) {
        _eventBufferSize = bufferSize;
    }

    function insertEvent(uint64 eventId, uint16 eventType, bytes memory data) internal {
        Event[] storage events = _eventsByType[eventType].data;
        require(events.length < _eventBufferSize, "Event buffer is full");

        events.push(Event(eventId, eventType, data));
        _eventsByType[eventType].dataIdxByEventId[eventId] = uint64(events.length) - 1;
    }

    function eraseEvent(uint64 eventId, uint16 eventType) internal {
        EventEntries storage entries = _eventsByType[eventType];

        uint64 index = entries.dataIdxByEventId[eventId];
        require(index < entries.data.length, "Event does not exist");

        // Swap the last event with the one to delete and then pop the last
        uint64 lastIndex = uint64(entries.data.length) - 1;
        if (index != lastIndex) {
            Event storage lastEvent = entries.data[lastIndex];
            entries.data[index] = lastEvent;
            entries.dataIdxByEventId[lastEvent.eventId] = index;
        }

        entries.data.pop();
        delete entries.dataIdxByEventId[eventId];
    }

    function pollEvents(uint16 eventType) external view returns (Event[] memory) {
        EventEntries storage entries = _eventsByType[eventType];
        return entries.data;
    }

    function getEvents(uint16 eventType) internal view returns (Event[] memory) {
        EventEntries storage entries = _eventsByType[eventType];
        return entries.data;
    }
}
