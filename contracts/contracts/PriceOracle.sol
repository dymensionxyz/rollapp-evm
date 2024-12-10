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

    struct PriceWithProof {
        uint256 price;
        PriceProof proof;
    }

    struct PriceWithExpiration {
        uint256 price;
        uint256 expiration;
        bool exists;
    }

    address public owner;
    bool public initialized;

    uint256 public expirationOffset;

    mapping(address => mapping(address => PriceWithExpiration)) public prices_cache;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event OracleInitialized(address indexed initializer);
    event PriceUpdated(address indexed base, address indexed quote, uint256 price);

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

    function updatePrice(address base, address quote, PriceWithProof calldata priceWithProof) external onlyOwner {
        require(
            !priceIsExpired(priceWithProof),
            "PriceOracle: price proof expired"
        );

        _validatePriceExpiration(base, quote, priceWithProof);
        _updatePriceCache(base, quote, priceWithProof);
    }

    function _updatePriceCache(address base, address quote, PriceWithProof calldata priceWithProof) internal {
        prices_cache[base][quote] = PriceWithExpiration(
            priceWithProof.price,
            getProofExpiryDate(priceWithProof.proof),
            true
        );

        emit PriceUpdated(base, quote, priceWithProof.price);
    }

    function _validatePriceExpiration(
        address base,
        address quote,
        PriceWithProof calldata priceWithProof
    ) internal view {
        PriceWithExpiration memory cachedPriceWithExpiration = prices_cache[base][quote];
        if (cachedPriceWithExpiration.exists) {
            require(
                cachedPriceWithExpiration.expiration < getProofExpiryDate(priceWithProof.proof),
                "PriceOracle: cannot update with an older price"
            );
        }
    }

    function priceIsExpired(PriceWithProof memory price) public view returns (bool) {
        return (block.timestamp * 1000) > getProofExpiryDate(price.proof);
    }

    function getProofExpiryDate(PriceProof memory proof) public view returns (uint256) {
        return proof.creationTimeUnixMs + (expirationOffset * 1000);
    }
}
