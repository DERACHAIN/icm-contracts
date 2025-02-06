// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "./NativeTokenStakingManager.sol";

contract NativeTokenStakingManagerV2 is NativeTokenStakingManager {
    constructor(ICMInitializable init) NativeTokenStakingManager(init) {}

    function foo() external pure returns (string memory) {
        return "foo";
    }
}
