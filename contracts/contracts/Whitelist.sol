// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract Whitelist is Ownable {
    mapping(address => bool) private whitelist;

    error WhitelistUnauthorizedAccount(address account);

    event AddWhitelisted(address indexed account);
    event RemoveWhitelisted(address indexed account);

    constructor(address initialOwner) Ownable(initialOwner) {}

    modifier onlyWhitelisted() {
        require(isWhitelisted(msg.sender), WhitelistUnauthorizedAccount(msg.sender));
        _;
    }

    function addWhitelisted(address _address) public onlyOwner {
        whitelist[_address] = true;
        emit AddWhitelisted(_address);
    }

    function removeWhitelisted(address _address) public onlyOwner {
        whitelist[_address] = false;
        emit RemoveWhitelisted(_address);
    }

    function isWhitelisted(address _address) public view returns(bool) {
        return whitelist[_address];
    }
}
