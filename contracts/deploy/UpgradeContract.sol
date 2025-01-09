// script/UpgradeContract.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../validator-manager/DummyNativeTokenStakingManager.sol";
import {PoSValidatorManagerSettings} from "../validator-manager/interfaces/IPoSValidatorManager.sol";
import {ValidatorManagerSettings} from "../validator-manager/interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "../validator-manager/interfaces/IRewardCalculator.sol";

contract UpgradeContract is Script {
    // Add these addresses after initial deployment
    address constant PROXY_ADDRESS = address(0xD889c48750921aFFE011d1Ca30bE14Cf6DC792Cf);
    address constant PROXY_ADMIN_ADDRESS = address(0xf3b3Fe49daaB7dA2A7724091b6735095aeba7225);

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy new implementation
        DummyNativeTokenStakingManager newImplementation = new DummyNativeTokenStakingManager(ICMInitializable.Allowed);
        console.log("New implementation deployed to:", address(newImplementation));

        // 2. Upgrade proxy to new implementation
        ProxyAdmin proxyAdmin = ProxyAdmin(PROXY_ADMIN_ADDRESS);
        proxyAdmin.upgradeAndCall{value: 0}(
            ITransparentUpgradeableProxy(PROXY_ADDRESS),
            address(newImplementation),
            ""
        );
        console.log("Proxy upgraded to new implementation");

        vm.stopBroadcast();
    }
}