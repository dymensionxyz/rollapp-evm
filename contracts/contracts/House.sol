// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract House is Ownable {
    mapping(address => uint256) private balances;

    constructor(address initialOwner) Ownable(initialOwner) {}

    function depositSupply() external payable onlyOwner {}

    function withdrawSupply(uint256 amount) external onlyOwner {
        payable(msg.sender).transfer(amount);
    }

    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }

    function withdraw() external {
        uint256 amount = balances[msg.sender];
        balances[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
    }

    function addBalance(address user, uint256 amount) internal {
        balances[user] += amount;
    }

    function reduceBalance(address user, uint256 amount) internal {
        balances[user] -= amount;
    }
}
