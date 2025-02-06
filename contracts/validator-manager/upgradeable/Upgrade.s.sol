// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../NativeTokenStakingManagerV2.sol";

contract UpgradeScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        address proxyAdminAddress = vm.envAddress("PROXY_ADMIN_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        NativeTokenStakingManagerV2 nativeTokenStakingManagerV2 = new NativeTokenStakingManagerV2(
                ICMInitializable.Allowed
            );

        ProxyAdmin proxyAdmin = ProxyAdmin(proxyAdminAddress);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(proxyAddress),
            address(nativeTokenStakingManagerV2),
            "" // no data
        );

        vm.stopBroadcast();

        console.log(
            "NativeTokenStakingManagerV2 implementation upgraded at: ",
            address(nativeTokenStakingManagerV2)
        );
        console.log("Proxy upgraded at: ", proxyAddress);
    }
}
