// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title AIOracle
 * @dev A contract for interacting with an AI system through prompts and answers
 */
contract AIOracle {
    address public aiAgent; // Address allowed to provide answers and modify prompters
    uint256 public latestPromptId;
    mapping(uint256 => string) public answers; // Stores answers by prompt ID
    mapping(address => bool) private whitelistedPrompters; // Whitelisted addresses that can submit prompts

    event PromptSubmitted(uint256 promptId, string prompt);
    event AnswerSubmitted(uint256 promptId, string answer);

    event AddWhitelisted(address indexed account);
    event RemoveWhitelisted(address indexed account);
    event OwnershipTransferred(address indexed oldAIAgent, address indexed newAIAgent);

    /**
     * @dev Sets the authorized writer during deployment
     * @param _authorizedWriter Address of the AI agent
     */
    constructor(address _authorizedWriter) {
        require(_authorizedWriter != address(0), "AIOracle: invalid writer address");
        aiAgent = _authorizedWriter;
        latestPromptId = 0;
    }

    function transferOwnership(address newAIAgent) external onlyAIAgent {
        require(newAIAgent != address(0), "AIOracle: new owner is the zero address");

        address oldAIAgent = aiAgent;
        aiAgent = newAIAgent;
        emit OwnershipTransferred(oldAIAgent, newAIAgent);
    }

    function isWhitelistedPrompter(address account) public view returns (bool) {
        return whitelistedPrompters[account];
    }

    modifier onlyAIAgent() {
        require(msg.sender == aiAgent, "AIOracle: caller is not the AI agent");
        _;
    }

    modifier onlyWhitelistedPrompter() {
        require(isWhitelistedPrompter(msg.sender), "AIOracle: caller is not a whitelisted prompter");
        _;
    }

    /**
     * @dev Creates a new prompt and emits an event
     * @param prompt The prompt string
     * @return The ID of the newly created prompt
     */
    function submitPrompt(string memory prompt) external onlyWhitelistedPrompter returns (uint256) {
        require(bytes(prompt).length > 0, "AIOracle: prompt cannot be empty");

        latestPromptId++;
        emit PromptSubmitted(latestPromptId, prompt);
        return latestPromptId;
    }

    /**
     * @dev Submits an answer for a specific prompt ID
     * @param id The ID of the prompt
     * @param answer The answer string
     */
    function submitAnswer(uint256 id, string memory answer) external onlyAIAgent {
        require(id > 0 && id <= latestPromptId, "AIOracle: invalid prompt ID");
        require(bytes(answers[id]).length == 0, "AIOracle: answer already exists");
        require(bytes(answer).length > 0, "AIOracle: answer cannot be empty");

        answers[id] = answer;
        emit AnswerSubmitted(id, answer);
    }

    /**
     * @dev Retrieves the answer for a specific prompt ID
     * @param id The ID of the prompt
     * @return The answer string
     */
    function getAnswer(uint256 id) external view returns (string memory) {
        string memory answer = answers[id];
        require(bytes(answer).length > 0, "AIOracle: answer does not exist");
        return answer;
    }

    /**
     * @dev Adds an address to the whitelist.
     * @param account The address to whitelist.
     */
    function addWhitelistAddress(address account) external onlyAIAgent {
        require(account != address(0), "AIOracle: invalid address");
        require(!whitelistedPrompters[account], "AIOracle: address already whitelisted");

        whitelistedPrompters[account] = true;
        emit AddWhitelisted(account);
    }

    /**
     * @dev Removes an address from the whitelist.
     * @param account The address to remove from the whitelist.
     */
    function removeWhitelistAddress(address account) external onlyAIAgent {
        require(whitelistedPrompters[account], "AIOracle: address not whitelisted");

        whitelistedPrompters[account] = false;
        emit RemoveWhitelisted(account);
    }
}