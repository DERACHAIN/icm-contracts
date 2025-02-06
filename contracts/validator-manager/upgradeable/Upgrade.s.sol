// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../NativeTokenStakingManager.sol";
import {ValidatorManagerSettings, PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract UpgradeScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        bytes32 l1ID = vm.envBytes32("L1_ID");
        address rewardCalculator = vm.envAddress("REWARD_CALCULATOR");
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        address proxyAdminAddress = vm.envAddress("PROXY_ADMIN_ADDRESS");

        console.log(
            "Proxy address, proxy admin address: ",
            proxyAddress,
            proxyAdminAddress
        );

        vm.startBroadcast(deployerPrivateKey);

        // Deploy new implementation
        NativeTokenStakingManager stakingManagerV2 = new NativeTokenStakingManager(
                ICMInitializable.Allowed
            );
        console.log("New implementation address: ", address(stakingManagerV2));

        // Verify ownership
        ProxyAdmin proxyAdmin = ProxyAdmin(proxyAdminAddress);
        console.log("Current owner: ", proxyAdmin.owner());
        console.log("Deployer: ", vm.addr(deployerPrivateKey));

        // Upgrade without initialization
        proxyAdmin.upgradeAndCall{value: 0}(
            ITransparentUpgradeableProxy(proxyAddress),
            address(stakingManagerV2),
            bytes("")
        );

        vm.stopBroadcast();
    }
}
