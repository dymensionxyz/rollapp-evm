// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "hardhat/console.sol";

/**
 * @title PriceOracle
 * @dev Basic price oracle contract structure for rollapp-evm
 */
contract PriceOracle {
    uint256 public constant SCALE_FACTOR = 10 ** 18;

    struct AssetInfo {
        address localNetworkName;
        string oracleNetworkName;
        uint256 localNetworkPrecision;
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

    struct GetPriceResponse {
        uint256 price;
        bool is_inverse;
    }

    address public owner;
    bool public initialized;

    uint256 public expirationOffset;
    uint256 public boundThreshold;

    mapping(address => mapping(address => PriceWithExpiration)) public prices_cache;
    mapping(address => uint256) public precisionMapping;
    mapping(address => string) public localNetworkToOracleNetworkDenoms;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event OracleInitialized(address indexed initializer);
    event PriceUpdated(address indexed base, address indexed quote, uint256 price);

    constructor(
        uint256 _expirationOffset,
        AssetInfo[] memory _assetInfos,
        uint256 boundThreshold
    ) {
        owner = msg.sender;
        expirationOffset = _expirationOffset;
        boundThreshold = boundThreshold;
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

    function priceInBound(address base, address quote, uint256 price) public view returns (bool) {
        // Get oracle price
        GetPriceResponse memory oraclePrice = this.getPrice(base, quote);

        // Calculate acceptable lower bound using SCALE_FACTOR for precision
        uint256 acceptablePrice = (oraclePrice.price * (SCALE_FACTOR - boundThreshold)) / SCALE_FACTOR;

        // Check if provided price is acceptable
        require(price >= acceptablePrice, "Price is below the acceptable threshold");

        return true;
    }

    function getPrice(address base, address quote) external view returns (GetPriceResponse memory) {
        // This function tries to retrieve a stored price for the (base, quote) pair.
        // If it's not available directly, it checks for the inverse pair (quote, base)
        // and calculates the inverted price if found.

        PriceWithExpiration memory priceWithExpiration = prices_cache[base][quote];
        if (priceWithExpiration.exists) {
            require((block.timestamp * 1000) <= priceWithExpiration.expiration, "PriceOracle: price expired");
            return GetPriceResponse(priceWithExpiration.price, false);
        }

        priceWithExpiration = prices_cache[quote][base];
        if (priceWithExpiration.exists) {
            require((block.timestamp * 1000) <= priceWithExpiration.expiration, "PriceOracle: price expired");
            require(priceWithExpiration.price > 0, "PriceOracle: invalid price for inversion");

            // When we need the inverted price, we're essentially calculating 1 / (quote->base price).
            // However, due to internal scaling, we must adjust for both the base token's precision
            // and the contract's overall scaling factor (SCALE_FACTOR).

            // 1) Multiplying by 10**precisionMapping[base]:
            //    Each token might have a different number of decimals. By raising 10 to the power of
            //    precisionMapping[base], we ensure the "unit" for the base token is aligned correctly.
            //    This step normalizes the inverted price so that it reflects the correct scale for
            //    the base token's decimal system.

            // 2) Multiplying by SCALE_FACTOR (10^18):
            //    After adjusting for the token's precision, we also incorporate a uniform scaling factor
            //    (SCALE_FACTOR) that the contract uses to represent all prices internally at a high-precision,
            //    standardized scale. This ensures consistent calculations and comparisons across all tokens
            //    and price pairs, regardless of their native decimal representation.
            uint256 invertedPrice = (10**precisionMapping[base] * SCALE_FACTOR) / priceWithExpiration.price;

            return GetPriceResponse(invertedPrice, true);
        }

        revert("PriceOracle: price not found");
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
        uint256 adjustedPrice = _precisionAdjustedPrice(base, quote, priceWithProof.price);

        prices_cache[base][quote] = PriceWithExpiration(
            adjustedPrice,
            _getProofExpiryDate(priceWithProof.proof),
            true
        );

        emit PriceUpdated(base, quote, adjustedPrice); // Emitir el precio ajustado
    }

    function _precisionAdjustedPrice(address base, address quote, uint256 price) internal view returns (uint256) {
        uint256 basePrecision = precisionMapping[base];   // e.g., BTC might have 8 decimals
        uint256 quotePrecision = precisionMapping[quote]; // e.g., USDC might have 18 decimals

        require(basePrecision != 0 && quotePrecision != 0, "PriceOracle: precision not set for tokens");

        // The goal is to normalize the price to the SCALE_FACTOR (10^18).
        // Consider 1 BTC = 50,000 USDC.
        // - BTC (8 decimals): 1 BTC internally could be 1 * 10^8 units.
        // - USDC (18 decimals): 50,000 USDC = 50,000 * 10^18 units.
        //
        // Without scaling, mixing these different decimal systems would be complex.
        // By scaling everything to a common factor (10^18), we ensure consistent internal representation.
        // This makes calculations simpler and less error-prone, as all prices are handled at the same magnitude.

        if (quotePrecision > basePrecision) {
            // In this case, if quotePrecision is greater (e.g., 18 for USDC) than basePrecision (e.g., 8 for BTC),
            // we need to "increase" the scale of the price to match the higher precision.
            // For BTC (8 decimals) vs. USDC (18 decimals), the difference is 10.
            // We multiply the base price by 10^(18 - 8) = 10^10 to scale it appropriately.
            uint256 exponent = quotePrecision - basePrecision;
            require(price <= (type(uint256).max / SCALE_FACTOR) / (10 ** exponent), "PriceOracle: price too large");
            return (price * SCALE_FACTOR) * (10 ** exponent);
        } else if (basePrecision > quotePrecision) {
            // If the base token has more decimals than the quote token,
            // we would reduce the scale accordingly. This is not the case in our BTC/USDC example,
            // but this branch handles other scenarios.
            uint256 exponent = basePrecision - quotePrecision;
            require(exponent <= 77, "PriceOracle: exponent too large");
            return (price * SCALE_FACTOR) / (10 ** exponent);
        } else {
            // If both have the same precision, simply multiply by 10^18.
            return price * SCALE_FACTOR;
        }
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
            precisionMapping[_assetInfos[i].localNetworkName] = _assetInfos[i].localNetworkPrecision;
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

        require(
            bytes(localNetworkToOracleNetworkDenoms[quote]).length > 0,
            "PriceOracle: quote denom not found in local_network_to_oracle_network_denoms"
        );
    }
}