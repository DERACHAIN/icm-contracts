// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";

contract InitializeEndDelegationScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        bytes32 delegationID = vm.envBytes32("DELEGATION_ID");

        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        NativeTokenStakingManager nativeTokenStakingManager = NativeTokenStakingManager(
                proxyAddress
            );

        vm.broadcast(deployerPrivateKey);

        nativeTokenStakingManager.initializeEndDelegation(
            delegationID,
            false,
            0
        );

        vm.stopBroadcast();

        console.log("Initialize end delegation succeeded");
    }
}
