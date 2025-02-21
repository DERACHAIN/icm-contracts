// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";

contract InitializeDelegatorRegistrationScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        bytes32 validationID = vm.envBytes32("VALIDATION_ID");

        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        NativeTokenStakingManager nativeTokenStakingManager = NativeTokenStakingManager(
                proxyAddress
            );

        vm.broadcast(deployerPrivateKey);

        bytes32 delegationID = nativeTokenStakingManager
            .initializeDelegatorRegistration{value: 1_000 ether}(validationID);

        vm.stopBroadcast();

        console.log("DelegationID");
        console.logBytes32(delegationID);
    }
}
