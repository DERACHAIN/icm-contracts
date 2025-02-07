// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {UnsafeUpgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import "../NativeTokenStakingManager.sol";
import {ValidatorManagerSettings, PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract UpgradeScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        bytes32 l1ID = vm.envBytes32("L1_ID");
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        NativeTokenStakingManager nativeTokenStakingManagerV2 = new NativeTokenStakingManager(
                ICMInitializable.Allowed
            );

        address deployer = vm.addr(deployerPrivateKey);

        UnsafeUpgrades.upgradeProxy(
            proxyAddress,
            address(nativeTokenStakingManagerV2),
            ""
        );

        vm.stopBroadcast();

        console.log(
            "NativeTokenStakingManager new implementation deployed at: ",
            address(nativeTokenStakingManagerV2)
        );
    }
}
