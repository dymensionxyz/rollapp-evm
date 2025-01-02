// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {EventManager} from "./EventManager.sol";
import {Whitelist} from "./Whitelist.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title AIOracle
 * @dev A contract for interacting with an AI system through prompts and answers
 */
contract AIOracle is Whitelist, EventManager {
    // Store answers by prompt ID
    mapping(uint64 => string) public answers;
    uint64 public latestPromptId;

    // Types of stateful events used by EventManager.
    // Don't change the order of the entries in enum declaration.
    // Backend relies on integer number under the enum.
    enum EventType {
        PromptSubmitted
    }

    struct UnprocessedPrompt {
        uint64 promptId;
        string prompt;
    }

    event PromptSubmitted(uint64 promptId, string prompt);
    event AnswerSubmitted(uint64 promptId, string answer);

    /**
     * @dev Sets the owner during deployment
     */
    constructor(address initialOwner) Whitelist(initialOwner) EventManager(10240) {}

    /**
     * @dev Creates a new prompt and emits an event
     * @return The ID of the newly created prompt
     */
    function submitPrompt(string memory prompt) external onlyWhitelisted returns (uint64) {
        require(bytes(prompt).length > 0, "AIOracle: prompt cannot be empty");

        latestPromptId++;

        insertEvent(
            latestPromptId,
            uint16(EventType.PromptSubmitted),
            encodeUnprocessedPrompt(latestPromptId, prompt)
        );
        emit PromptSubmitted(latestPromptId, prompt);

        return latestPromptId;
    }

    /**
     * @dev Submits an answer for a specific prompt ID
     */
    function submitAnswer(uint64 promptId, string memory answer) external onlyOwner {
        require(promptId > 0 && promptId <= latestPromptId, "AIOracle: invalid prompt ID");
        require(bytes(answers[promptId]).length == 0, "AIOracle: answer already exists");
        require(bytes(answer).length > 0, "AIOracle: answer cannot be empty");

        answers[promptId] = answer;

        eraseEvent(promptId, uint16(EventType.PromptSubmitted));
        emit AnswerSubmitted(promptId, answer);
    }

    /**
     * @dev Retrieves the answer for a specific prompt ID
     */
    function getAnswer(uint64 promptId) external view returns (string memory) {
        string memory answer = answers[promptId];
        require(bytes(answer).length > 0, "AIOracle: answer does not exist");
        return answer;
    }

    /**
     * @dev Encode prompt data
     */
    function encodeUnprocessedPrompt(uint64 promptId, string memory prompt) internal pure returns (bytes memory) {
        return abi.encode(promptId, prompt);
    }

    /**
     * @dev Decode prompt data
     */
    function decodeUnprocessedPrompt(bytes memory data) internal pure returns (UnprocessedPrompt memory) {
        (uint64 promptId, string memory prompt) = abi.decode(data, (uint64, string));
        return UnprocessedPrompt(promptId, prompt);
    }

    /**
     * @dev Get all unprocessed prompts
     */
    function getUnprocessedPrompts() external view returns (UnprocessedPrompt[] memory) {
        Event[] memory events = getEvents(uint16(EventType.PromptSubmitted));
        UnprocessedPrompt[] memory res = new UnprocessedPrompt[](events.length);
        for (uint64 i; i < events.length; i++) {
            res[i] = decodeUnprocessedPrompt(events[i].data);
        }
        return res;
    }
}