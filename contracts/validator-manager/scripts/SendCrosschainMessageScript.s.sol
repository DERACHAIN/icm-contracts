// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {TeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "../../teleporter/TeleporterMessenger.sol";

contract SendCrosschainMessageScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);

        bytes32 destinationChainID = keccak256("Fuji");
        address destinationAddress = 0x0000000000000000000000000000000000000001;
        bytes memory message = bytes("hello");
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo({
            feeTokenAddress: 0x0000000000000000000000000000000000000000,
            amount: 0 ether
        });

        address teleporterMessengerAddress = vm.envAddress(
            "TELEPORTER_ADDRESS"
        );
        console.log(
            "TeleporterMessenger address: %s",
            teleporterMessengerAddress
        );

        TeleporterMessenger teleporterMessenger = TeleporterMessenger(
            teleporterMessengerAddress
        );

        vm.broadcast(deployerPrivateKey);

        TeleporterMessageInput memory input = TeleporterMessageInput({
            destinationBlockchainID: destinationChainID,
            destinationAddress: destinationAddress,
            requiredGasLimit: 300_000,
            allowedRelayerAddresses: new address[](0),
            feeInfo: feeInfo,
            message: message
        });

        teleporterMessenger.sendCrossChainMessage(input);

        vm.stopBroadcast();

        console.log("Crosschain message sent");
    }
}
