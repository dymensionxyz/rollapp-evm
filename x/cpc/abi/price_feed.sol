// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

interface IDymPriceFeedCPC {
    /**
     * @dev Returns price for the given token by name.
     * The second return value indicating whether the operation succeeded.
     */
    function getPrice(string memory tokenName) external view returns (uint256, bool);
}