// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title PriceOracle
 * @dev Basic price oracle contract structure for rollapp-evm
 */
contract PriceOracle {
    struct PriceProof {
        uint256 creationHeight;
        uint256 creationTimeUnixMs;
        uint256 height;
        uint256 revision;
        bytes merkleProof;
    }

    address public owner;
    bool public initialized;

    uint256 public expirationOffset;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event OracleInitialized(address indexed initializer);

    constructor(uint256 _expirationOffset) {
        owner = msg.sender;
        expirationOffset = _expirationOffset;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "PriceOracle: caller is not the owner");
        _;
    }

    modifier notInitialized() {
        require(!initialized, "PriceOracle: already initialized");
        _;
    }

    function initialize() external onlyOwner notInitialized {
        initialized = true;
        emit OracleInitialized(msg.sender);
    }

    function transferOwnership(address newOwner) external onlyOwner {
        require(newOwner != address(0), "PriceOracle: new owner is the zero address");
        address oldOwner = owner;
        owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }

    function updatePrice(address base, address quote, PriceProof calldata proof) external onlyOwner {
        // validate expiration
        require(
            (block.timestamp * 1000) <= (proof.creationTimeUnixMs + (expirationOffset * 1000)),
            "PriceOracle: price proof expired"
        );
    }
}