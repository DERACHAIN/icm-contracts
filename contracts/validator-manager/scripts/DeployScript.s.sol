// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {UnsafeUpgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import "../NativeTokenStakingManager.sol";
import {ValidatorManagerSettings, PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        bytes32 l1ID = vm.envBytes32("L1_ID");
        address rewardCalculator = vm.envAddress("REWARD_CALCULATOR");

        vm.startBroadcast(deployerPrivateKey);

        NativeTokenStakingManager nativeTokenStakingManager = new NativeTokenStakingManager(
                ICMInitializable.Allowed
            );

        ValidatorManagerSettings
            memory validatorManagerSettings = ValidatorManagerSettings({
                l1ID: l1ID,
                churnPeriodSeconds: 60,
                maximumChurnPercentage: 20
            });

        PoSValidatorManagerSettings
            memory settings = PoSValidatorManagerSettings({
                baseSettings: validatorManagerSettings,
                minimumStakeAmount: 1000 ether,
                maximumStakeAmount: 1_000_000 ether,
                minimumStakeDuration: 60 * 60,
                minimumDelegationFeeBips: 10,
                maximumStakeMultiplier: 5,
                weightToValueFactor: 1_000,
                rewardCalculator: IRewardCalculator(rewardCalculator),
                uptimeBlockchainID: l1ID
            });

        address proxy = UnsafeUpgrades.deployTransparentProxy(
            address(nativeTokenStakingManager),
            address(deployer),
            abi.encodeCall(NativeTokenStakingManager.initialize, settings)
        );

        vm.stopBroadcast();

        console.log(
            "NativeTokenStakingManager implementation deployed at: ",
            address(nativeTokenStakingManager)
        );
        console.log("Proxy deployed at: ", address(proxy));
    }
}
