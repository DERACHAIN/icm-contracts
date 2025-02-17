// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {IWarpMessenger, WarpMessage} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

contract WarpMessengerScript is Script {
    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        vm.startBroadcast(deployerPrivateKey);

        IWarpMessenger warpMessenger = IWarpMessenger(WARP_PRECOMPILE_ADDRESS);
        bytes32 blockchainId = warpMessenger.getBlockchainID();

        vm.stopBroadcast();

        console.logBytes32(blockchainId);
    }
}
