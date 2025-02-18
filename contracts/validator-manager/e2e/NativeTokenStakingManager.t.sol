// SPDX-License-Identifier: None
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import {Test} from "@forge-std/Test.sol";
import "../NativeTokenStakingManager.sol";
import {IWarpMessenger, WarpMessage} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorRegistrationInput, PChainOwner, ConversionData, InitialValidator} from "../interfaces/IValidatorManager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import "forge-std/console2.sol";

contract NativeTokenStakingManagerTest is Test {
    address public constant proxyAddress =
        0x0Feedc0de0000000000000000000000000000000;
    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;

    NativeTokenStakingManager public app;
    IWarpMessenger public warpMessenger;

    uint16 public constant DELEGATION_FEE_BIPS = 1000;
    uint64 public constant MIN_STAKE_DURATION = 24 * 3600;
    bytes public constant NODE_ID =
        bytes(hex"136c6bb8eb75447db873d66e6b6d61feb5c1a161");
    bytes public constant BLS_PUBLIC_KEY =
        bytes(
            hex"b899613a28e1f55b250d587c9171fa241d11ec490f860f1b4cb8f33e7aa081956ce66c999b48d0b7712911522cc64c68"
        );
    address public constant OWNER = 0xc0Ce63ca7605cb29aA6bcd040715D2A383a9f4aC;
    bytes32 public constant L1_ID =
        0xf63b77a038f89f4b984e931cadd3553b27194a7e2a54d7bb8ecc4cde363fc33f;
    bytes32 public constant BLOCKCHAIN_ID =
        0xf63b77a038f89f4b984e931cadd3553b27194a7e2a54d7bb8ecc4cde363fc33f;
    address public constant VALIDATOR_MANAGER_ADDRESS =
        0x0Feedc0de0000000000000000000000000000000;
    uint64 public constant DEFAULT_WEIGHT = 100;

    error InvalidInitializationStatus();

    function setUp() public {
        app = NativeTokenStakingManager(proxyAddress);
        warpMessenger = IWarpMessenger(WARP_PRECOMPILE_ADDRESS);
    }

    function xtest_initializeValidatorSet() public {
        InitialValidator[] memory initialValidators = new InitialValidator[](1);
        initialValidators[0] = InitialValidator({
            nodeID: NODE_ID,
            blsPublicKey: BLS_PUBLIC_KEY,
            weight: DEFAULT_WEIGHT
        });
        ConversionData memory conversionData = ConversionData({
            l1ID: L1_ID,
            validatorManagerBlockchainID: BLOCKCHAIN_ID,
            validatorManagerAddress: VALIDATOR_MANAGER_ADDRESS,
            initialValidators: initialValidators
        });

        vm.expectRevert(InvalidInitializationStatus.selector);

        app.initializeValidatorSet(conversionData, 0);
    }

    // Copy the struct definition
    enum ValidatorStatus {
        Unknown,
        PendingAdded,
        Active,
        PendingRemoved,
        Completed,
        Invalidated
    }

    struct Validator {
        ValidatorStatus status;
        bytes nodeID;
        uint64 startingWeight;
        uint64 messageNonce;
        uint64 weight;
        uint64 startedAt;
        uint64 endedAt;
    }

    struct ValidatorChurnPeriod {
        uint256 startedAt;
        uint64 initialWeight;
        uint64 totalWeight;
        uint64 churnAmount;
    }

    struct ValidatorManagerStorage {
        bytes32 _l1ID;
        uint64 _churnPeriodSeconds;
        uint8 _maximumChurnPercentage;
        ValidatorChurnPeriod _churnTracker;
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        mapping(bytes32 => Validator) _validationPeriods;
        mapping(bytes => bytes32) _registeredValidators;
        bool _initializedValidatorSet;
    }

    // Copy the storage location constant
    bytes32 public constant VALIDATOR_MANAGER_STORAGE_LOCATION =
        0xe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00;

    bytes32 public constant POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    // Function to get storage reference
    function _getValidatorManagerStorage(
        address validatorManagerAddress
    ) internal pure returns (ValidatorManagerStorage storage $) {
        assembly {
            // Calculate storage slot for the specific contract address
            $.slot := VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    function _getInitializedStorageSlot(
        uint64 index
    ) internal pure returns (bytes32 slot) {
        assembly {
            slot := add(VALIDATOR_MANAGER_STORAGE_LOCATION, index)
        }
    }

    // Example usage function
    function readValidatorManagerStorage(
        address validatorManagerAddress
    ) internal view returns (bytes32, uint64, uint8, bool) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage(
            validatorManagerAddress
        );
        return (
            $._l1ID,
            $._churnPeriodSeconds,
            $._maximumChurnPercentage,
            $._initializedValidatorSet
        );
    }

    function test_InitializeValidatorRegistration() public {
        address[] memory owners = new address[](1);
        owners[0] = OWNER;
        ValidatorRegistrationInput memory input = ValidatorRegistrationInput({
            nodeID: NODE_ID,
            blsPublicKey: BLS_PUBLIC_KEY,
            registrationExpiry: uint64(block.timestamp + 3600),
            remainingBalanceOwner: PChainOwner({
                threshold: 1,
                addresses: owners
            }),
            disableOwner: PChainOwner({threshold: 1, addresses: owners})
        });

        bytes32 l1Id = vm.load(
            proxyAddress,
            VALIDATOR_MANAGER_STORAGE_LOCATION
        );
        console.logBytes32(l1Id);

        bytes32 churnSlot = _getInitializedStorageSlot(1);
        bytes32 churnData = vm.load(proxyAddress, churnSlot);
        console.log("Churn data:");
        console.logBytes32(churnData);

        // uint64 churnPeriods;
        // uint8 maxChurnPct;
        // assembly {
        //     // Extract uint64 (_churnPeriodSeconds) - first 8 bytes
        //     churnPeriods := and(churnData, 0xFFFFFFFFFFFFFFFF)

        //     // Extract uint8 (_maximumChurnPercentage) - next byte after uint64
        //     // Shift right by 64 bits (8 bytes) and mask with 0xFF to get the uint8
        //     maxChurnPct := and(shr(64, churnData), 0xFF)
        // }

        // console.log("Churn periods: %s", uint256(churnPeriods));
        // console.log("Max churn percentage: %s", uint256(maxChurnPct));

        // bytes32 slotAddress = _getInitializedStorageSlot(2);
        // bytes32 dataSlot = vm.load(proxyAddress, slotAddress);
        // console.log("Started churn period %s", uint256(dataSlot));

        // bytes32 initializedSlot = bytes32(
        //     uint256(VALIDATOR_MANAGER_STORAGE_LOCATION) + 7
        // );

        // bytes32 initialized = vm.load(proxyAddress, initializedSlot);

        // console.log("Initialized:");
        // console.logBytes32(initialized);

        // bytes32 registeredValidatorsSlot = bytes32(
        //     uint256(VALIDATOR_MANAGER_STORAGE_LOCATION) + 6
        // );

        // console.log("Registered validators slot:");
        // console.logBytes32(registeredValidatorsSlot);

        bytes memory nodeId = bytes(
            hex"5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461"
        );

        // bytes32 actualSlot = keccak256(
        //     abi.encodePacked(nodeId, registeredValidatorsSlot)
        // );
        // bytes32 validationId = vm.load(proxyAddress, actualSlot);

        // console.log("ValidationId by reading slot");
        // console.logBytes32(validationId);

        bytes32 validID = app.registeredValidators(nodeId);
        console.log("ValidationId by calling contract");
        console.logBytes32(validID);

        // bytes32 minStakeAmount = vm.load(
        //     proxyAddress,
        //     POS_VALIDATOR_MANAGER_STORAGE_LOCATION
        // );
        // console.log("Min stake amount %s", uint256(minStakeAmount));

        // bytes32 maxStakeAmountSlot;
        // assembly {
        //     maxStakeAmountSlot := add(POS_VALIDATOR_MANAGER_STORAGE_LOCATION, 1)
        // }
        // bytes32 maxStakeAmount = vm.load(proxyAddress, maxStakeAmountSlot);
        // console.log("Max stake amount %s", uint256(maxStakeAmount));

        //vm.expectRevert(InvalidInitializationStatus.selector);

        app.initializeValidatorRegistration{value: 10000 ether}(
            input,
            DELEGATION_FEE_BIPS,
            MIN_STAKE_DURATION
        );
    }
}
