// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {UnsafeUpgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import "../NativeTokenStakingManager.sol";
import "../RewardCalculator.sol";
import {ValidatorManagerSettings, PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";

contract UpgradeScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        bytes32 l1ID = vm.envBytes32("L1_ID");
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        NativeTokenStakingManager nativeTokenStakingManager = new NativeTokenStakingManager(
                ICMInitializable.Allowed
            );

        RewardCalculator rewardCalculator = new RewardCalculator(1);

        // reinitialize the validator manager settings
        ValidatorManagerSettings
            memory validatorManagerSettings = ValidatorManagerSettings({
                l1ID: l1ID,
                churnPeriodSeconds: 60,
                maximumChurnPercentage: 20
            });

        PoSValidatorManagerSettings
            memory settings = PoSValidatorManagerSettings({
                baseSettings: validatorManagerSettings,
                minimumStakeAmount: 10_000 ether,
                maximumStakeAmount: 100_000 ether,
                minimumStakeDuration: 60 * 60 * 24,
                minimumDelegationFeeBips: 10,
                maximumStakeMultiplier: 5,
                weightToValueFactor: 1e21,
                rewardCalculator: IRewardCalculator(rewardCalculator),
                uptimeBlockchainID: l1ID
            });

        UnsafeUpgrades.upgradeProxy(
            proxyAddress,
            address(nativeTokenStakingManager),
            abi.encodeCall(NativeTokenStakingManager.initialize, settings)
        );

        vm.stopBroadcast();

        console.log(
            "NativeTokenStakingManager proxy %s has been updated new implementation %s",
            proxyAddress,
            address(nativeTokenStakingManager)
        );
    }
}
