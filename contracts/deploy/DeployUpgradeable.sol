// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../validator-manager/DummyNativeTokenStakingManager.sol";
import {PoSValidatorManagerSettings} from "../validator-manager/interfaces/IPoSValidatorManager.sol";
import {ValidatorManagerSettings} from "../validator-manager/interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "../validator-manager/interfaces/IRewardCalculator.sol";

contract DeployUpgradeable is Script {
  bytes32 private constant l1ID = 0x0000000000000000000000000000000000000000000000000000000000000001;
  address private constant rewardCalculator = 0x71C85417237765a1779E75d2b1dd33Eb3295D208;

  function run() external {
      uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
      address deployer = vm.addr(deployerPrivateKey);
      console.log("Deployer address:", deployer);

      vm.startBroadcast(deployerPrivateKey);

      // 1. Deploy implementation
      DummyNativeTokenStakingManager implementation = new DummyNativeTokenStakingManager(ICMInitializable.Allowed);
      console.log("Implementation deployed to:", address(implementation));

      // 2. Deploy ProxyAdmin
      ProxyAdmin proxyAdmin = new ProxyAdmin(deployer);
      console.log("ProxyAdmin deployed to:", address(proxyAdmin));

      // 3. Encode initialization data
      ValidatorManagerSettings  memory validatorManagerSettings = ValidatorManagerSettings({
        l1ID: l1ID,
        churnPeriodSeconds: 3600,
        maximumChurnPercentage: 20
      });

      PoSValidatorManagerSettings memory settings = PoSValidatorManagerSettings({
        baseSettings: validatorManagerSettings,
        minimumStakeAmount: 1,
        maximumStakeAmount: 1000000,
        minimumStakeDuration: 1 days,
        minimumDelegationFeeBips: 1,
        maximumStakeMultiplier: 10,
        weightToValueFactor: 1,
        rewardCalculator: IRewardCalculator(rewardCalculator),
        uptimeBlockchainID: l1ID
      });

      bytes memory initData = abi.encodeWithSelector(
        DummyNativeTokenStakingManager.initialize.selector,
        settings
      );

      // 4. Deploy TransparentUpgradeableProxy
      TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
        address(implementation),
        address(proxyAdmin),
        initData
      );
      console.log("Proxy deployed to:", address(proxy));

      vm.stopBroadcast();
  }
}