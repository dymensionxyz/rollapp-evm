// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title PriceOracle
 * @dev Basic price oracle contract structure for rollapp-evm
 */
contract PriceOracle {
    struct AssetInfo {
        address localNetworkName;
        string oracleNetworkName;
        int localNetworkPrecision;
    }

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
    mapping(address => int) public precissionMapping;
    mapping(address => string) public localNetworkToOracleNetworkDenoms;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event OracleInitialized(address indexed initializer);
    event PriceUpdated(address indexed base, address indexed quote, uint256 price);

    constructor(uint256 _expirationOffset, AssetInfo[] memory _assetInfos) {
        owner = msg.sender;
        expirationOffset = _expirationOffset;
        _populateAssetsInfo(_assetInfos);
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
            !_priceIsExpired(priceWithProof),
            "PriceOracle: price proof expired"
        );

        _ensureNewPriceIsValid(base, quote, priceWithProof);

        _verifyPriceProof(base, quote, priceWithProof);

        _updatePriceCache(base, quote, priceWithProof);
    }

    function _updatePriceCache(address base, address quote, PriceWithProof calldata priceWithProof) internal {
        prices_cache[base][quote] = PriceWithExpiration(
            priceWithProof.price,
            _getProofExpiryDate(priceWithProof.proof),
            true
        );

        emit PriceUpdated(base, quote, priceWithProof.price);
    }

    /**
     * @dev Ensures that the new price's expiration is later than the currently cached price's expiration.
     * This prevents updating the cache with an outdated price.
     * @param base The base token address.
     * @param quote The quote token address.
     * @param priceWithProof The new price along with its proof.
     */
    function _ensureNewPriceIsValid(
        address base,
        address quote,
        PriceWithProof calldata priceWithProof
    ) internal view {
        PriceWithExpiration memory cachedPriceWithExpiration = prices_cache[base][quote];
        if (cachedPriceWithExpiration.exists) {
            require(
                cachedPriceWithExpiration.expiration < _getProofExpiryDate(priceWithProof.proof),
                "PriceOracle: cannot update with an older price"
            );
        }
    }

    function _priceIsExpired(PriceWithProof memory price) internal view returns (bool) {
        return (block.timestamp * 1000) > _getProofExpiryDate(price.proof);
    }

    function _getProofExpiryDate(PriceProof memory proof) internal view returns (uint256) {
        return proof.creationTimeUnixMs + (expirationOffset * 1000);
    }

    function _populateAssetsInfo(AssetInfo[] memory _assetInfos) internal {
        for (uint256 i = 0; i < _assetInfos.length; i++) {
            precissionMapping[_assetInfos[i].localNetworkName] = _assetInfos[i].localNetworkPrecision;
            localNetworkToOracleNetworkDenoms[_assetInfos[i].localNetworkName] = _assetInfos[i].oracleNetworkName;
        }
    }

    function _verifyPriceProof(
        address base,
        address quote,
        PriceWithProof calldata priceWithProof
    ) internal view {
        require(
            bytes(localNetworkToOracleNetworkDenoms[base]).length > 0,
            "PriceOracle: base denom not found in local_network_to_oracle_network_denoms"
        );

        // Check if quote exists in mapping
        require(
            bytes(localNetworkToOracleNetworkDenoms[quote]).length > 0,
            "PriceOracle: quote denom not found in local_network_to_oracle_network_denoms"
        );
    }
}