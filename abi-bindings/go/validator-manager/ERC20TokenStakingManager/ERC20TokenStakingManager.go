// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenstakingmanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ConversionData is an auto generated low-level Go binding around an user-defined struct.
type ConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorManagerSettings struct {
	BaseSettings             ValidatorManagerSettings
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	WeightToValueFactor      *big.Int
	RewardCalculator         common.Address
	UptimeBlockchainID       [32]byte
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status         uint8
	NodeID         []byte
	StartingWeight uint64
	MessageNonce   uint64
	Weight         uint64
	StartedAt      uint64
	EndedAt        uint64
}

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	SubnetID               [32]byte
	ChurnPeriodSeconds     uint64
	MaximumChurnPercentage uint8
}

// ValidatorMessagesValidationPeriod is an auto generated low-level Go binding around an user-defined struct.
type ValidatorMessagesValidationPeriod struct {
	SubnetID              [32]byte
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
	Weight                uint64
}

// ValidatorRegistrationInput is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistrationInput struct {
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}

// ERC20TokenStakingManagerMetaData contains all meta data concerning the ERC20TokenStakingManager contract.
var ERC20TokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InvalidTotalWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615fa0380380615fa083398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615e538061014d5f395ff3fe608060405234801561000f575f80fd5b506004361061023f575f3560e01c806380dd672f11610135578063b771b3bc116100b4578063d5f20ff611610079578063d5f20ff614610516578063df93d8de14610536578063e4a63c4014610540578063fb8b11dd14610554578063fd7ac5e714610567575f80fd5b8063b771b3bc146104ae578063ba3a4b97146104d4578063bc5fbfec146104e7578063bee0a03f146104fb578063c974d1b61461050e575f80fd5b80639e1bc4ef116100fa5780639e1bc4ef1461044e5780639e478eea14610461578063a3a65e4814610474578063a9778a7a146102ee578063afb9809614610487575f80fd5b806380dd672f146103fa5780638280a25a1461040d5780638ef34c981461041557806393e24598146104285780639ae064471461043b575f80fd5b8063467ef06f116101c1578063620658561161018657806362065856146103a757806366435abf146103ba578063732214f8146103cd57806376f78621146103d45780637d8d2f77146103e7575f80fd5b8063467ef06f146103305780634bee0040146103435780635dd6a6cb1461036457806360305d621461037757806360ad778414610394575f80fd5b806325e1c7761161020757806325e1c776146102b05780632e2194d8146102c357806335455ded146102ee57806337b9be8f1461030a5780633a1cfff61461031d575f80fd5b80630118acc4146102435780630322ed9814610258578063151d30d11461026b5780631ec447241461028a57806320d91b7a1461029d575b5f80fd5b610256610251366004614d97565b61057a565b005b610256610266366004614dd2565b61058b565b610273600a81565b60405160ff90911681526020015b60405180910390f35b610256610298366004614d97565b61081b565b6102566102ab366004614de9565b610827565b6102566102be366004614e37565b610ddd565b6102d66102d1366004614dd2565b610e51565b6040516001600160401b039091168152602001610281565b6102f761271081565b60405161ffff9091168152602001610281565b610256610318366004614e6c565b610ea5565b61025661032b366004614d97565b610eb8565b61025661033e366004614eba565b610ec4565b610356610351366004614ef8565b610fa9565b604051908152602001610281565b610256610372366004614e6c565b610fde565b61037f601481565b60405163ffffffff9091168152602001610281565b6102566103a2366004614e37565b610fea565b6103566103b5366004614f5f565b6112b2565b6102d66103c8366004614dd2565b6112d2565b6103565f81565b6102566103e2366004614d97565b6112e6565b6102566103f5366004614e6c565b6112f2565b610256610408366004614e37565b6112fe565b610273603081565b610256610423366004614f7a565b611538565b610256610436366004614dd2565b611619565b610256610449366004614e6c565b6116ad565b61035661045c366004614fa8565b6116b9565b61025661046f366004614fc8565b6116e5565b610256610482366004614eba565b6117c4565b6103567f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0081565b6104bc6005600160991b0181565b6040516001600160a01b039091168152602001610281565b6102566104e2366004614dd2565b6119ba565b6103565f80516020615dbe83398151915281565b610256610509366004614dd2565b611c13565b610273601481565b610529610524366004614dd2565b611d4f565b6040516102819190615070565b6102d66202a30081565b6103565f80516020615d9e83398151915281565b610256610562366004614f7a565b611e9e565b6103566105753660046150f0565b611f65565b6105868383835f611fc0565b505050565b5f8181525f80516020615dde8339815191526020526040808220815160e0810190925280545f80516020615dbe83398151915293929190829060ff1660058111156105d8576105d8614ffb565b60058111156105e9576105e9614ffb565b81526020016001820180546105fd9061515b565b80601f01602080910402602001604051908101604052809291908181526020018280546106299061515b565b80156106745780601f1061064b57610100808354040283529160200191610674565b820191905f5260205f20905b81548152906001019060200180831161065757829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a090910152909150815160058111156106df576106df614ffb565b1461071b575f8381526005830160205260409081902054905163170cc93360e21b81526107129160ff1690600401615193565b60405180910390fd5b606081015160405163854a893f60e01b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af4158015610792573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107b9919081019061529c565b6040518263ffffffff1660e01b81526004016107d591906152cd565b6020604051808303815f875af11580156107f1573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061081591906152df565b50505050565b6108158383835f611fec565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07545f80516020615dbe8339815191529060ff161561087957604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108bc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e091906152df565b836020013514610909576040516372b0a7e760e11b815260208401356004820152602401610712565b3061091a60608501604086016152f6565b6001600160a01b03161461095d5761093860608401604085016152f6565b604051632f88120d60e21b81526001600160a01b039091166004820152602401610712565b5f61096b6060850185615311565b905090505f805b828163ffffffff161015610bd3575f61098e6060880188615311565b8363ffffffff168181106109a4576109a4615356565b90506020028101906109b6919061536a565b6109bf906153d5565b80516040519192505f9160068801916109d791615450565b90815260200160405180910390205414610a0757805160405163a41f772f60e01b815261071291906004016152cd565b5f6002885f013584604051602001610a3692919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610a5091615450565b602060405180830381855afa158015610a6b573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610a8e91906152df565b90508086600601835f0151604051610aa69190615450565b90815260408051918290036020908101909220929092555f8381526005890190915220805460ff191660021781558251600190910190610ae690826154a5565b50604082810180515f84815260058a016020529290922060028101805492516001600160401b039485166001600160c01b031990941693909317600160801b85851602176001600160c01b0316600160c01b429590951694909402939093179092556003909101805467ffffffffffffffff19169055610b669085615574565b8251604051919550610b7791615450565b60408051918290038220908401516001600160401b031682529082907ffe3e5983f71c8253fb0b678f2bc587aa8574d8f1aab9cf82b983777f5998392c9060200160405180910390a3505080610bcc90615594565b9050610972565b5060038301805467ffffffffffffffff60401b1916600160401b6001600160401b0384168102919091179091556001840154606491610c16910460ff16836155b6565b6001600160401b03161015610c4957604051633e1a785160e01b81526001600160401b0382166004820152602401610712565b5f73__$fd0c147b4031eef6079b0498cbafa865f0$__634d847884610c6d87612319565b604001516040518263ffffffff1660e01b8152600401610c8d91906152cd565b602060405180830381865af4158015610ca8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ccc91906152df565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__6387418b8e886040518263ffffffff1660e01b8152600401610d06919061570c565b5f60405180830381865af4158015610d20573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d47919081019061529c565b90505f600282604051610d5a9190615450565b602060405180830381855afa158015610d75573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d9891906152df565b9050828114610dc45760405163baaea89d60e01b81526004810182905260248101849052604401610712565b5050506007909201805460ff1916600117905550505050565b610de68261242f565b610e06576040516330efa98b60e01b815260048101839052602401610712565b5f610e1083611d4f565b5190506002816005811115610e2757610e27614ffb565b14610e47578060405163170cc93360e21b81526004016107129190615193565b6108158383612458565b5f80610e5b6126f7565b60030154610e6990846157a0565b9050801580610e7e57506001600160401b0381115b15610e9f5760405163222d164360e21b815260048101849052602401610712565b92915050565b610eb184848484611fec565b5050505050565b6108158383835f61271b565b610ecc61293d565b5f610ed56126f7565b90505f80610ee284612974565b91509150610eef8261242f565b610efb57505050610f90565b5f828152600684016020908152604080832054600b8701909252909120546001600160a01b03918216911680610f32575080610f50565b5f848152600b86016020526040902080546001600160a01b03191690555b600483516005811115610f6557610f65614ffb565b03610f7457610f748185612d2c565b610f8a82610f8585604001516112b2565b612d56565b50505050505b610fa660015f80516020615dfe83398151915255565b50565b5f610fb261293d565b610fbe85858585612d8c565b9050610fd660015f80516020615dfe83398151915255565b949350505050565b61081584848484612ef1565b5f610ff36126f7565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561102c5761102c614ffb565b600381111561103d5761103d614ffb565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6110b382611d4f565b90506001835160038111156110ca576110ca614ffb565b146110eb578251604051633b0d540d60e21b815261071291906004016157bf565b60048151600581111561110057611100614ffb565b036111165761110e86612f1d565b505050505050565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f61113b89612319565b604001516040518263ffffffff1660e01b815260040161115b91906152cd565b606060405180830381865af4158015611176573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061119a91906157d9565b50915091508184146111c757846040015160405163089938b360e11b815260040161071291815260200190565b806001600160401b031683606001516001600160401b031610806112005750806001600160401b03168560a001516001600160401b0316115b1561122957604051632e19bc2d60e11b81526001600160401b0382166004820152602401610712565b5f888152600787016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b810267ffffffffffffffff60401b1990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f6112bb6126f7565b60030154610e9f906001600160401b038416615819565b5f6112dc82611d4f565b6080015192915050565b6105868383835f612ef1565b610eb18484848461271b565b61130661293d565b5f61130f6126f7565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561134857611348614ffb565b600381111561135957611359614ffb565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290506003815160038111156113d2576113d2614ffb565b146113f3578051604051633b0d540d60e21b815261071291906004016157bf565b60046114028260400151611d4f565b51600581111561141457611414614ffb565b14611513575f61142384612319565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f84604001516040518263ffffffff1660e01b815260040161146291906152cd565b606060405180830381865af415801561147d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114a191906157d9565b5091509150818460400151146114cd5760405163089938b360e11b815260048101839052602401610712565b806001600160401b03168460c001516001600160401b0316111561150f57604051632e19bc2d60e11b81526001600160401b0382166004820152602401610712565b5050505b61151c84612f1d565b505061153460015f80516020615dfe83398151915255565b5050565b5f6115416126f7565b90506001600160a01b0382166115755760405163caa903f960e01b81526001600160a01b0383166004820152602401610712565b5f8381526006820160205260409020546001600160a01b031633146115bb57335b604051636e2ccd7560e11b81526001600160a01b039091166004820152602401610712565b336001600160a01b038316036115ea575f928352600b01602052506040902080546001600160a01b0319169055565b5f838152600b82016020526040902080546001600160a01b0384166001600160a01b0319909116179055505050565b5f6116226126f7565b90505f61162e83611d4f565b519050600481600581111561164557611645614ffb565b14611665578060405163170cc93360e21b81526004016107129190615193565b5f8381526006830160205260409020546001600160a01b0316331461168a5733611596565b5f838152600683016020526040902054610586906001600160a01b031684612d2c565b61081584848484611fc0565b5f6116c261293d565b6116cd83338461312a565b9050610e9f60015f80516020615dfe83398151915255565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff168061172e575080546001600160401b03808416911610155b1561174c5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155611777848461336f565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b5f80516020615dbe8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f6117f786612319565b604001516040518263ffffffff1660e01b815260040161181791906152cd565b6040805180830381865af4158015611831573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118559190615830565b915091508061187b57604051632d07135360e01b81528115156004820152602401610712565b5f828152600484016020526040902080546118959061515b565b90505f036118b95760405163089938b360e11b815260048101839052602401610712565b60015f838152600580860160205260409091205460ff16908111156118e0576118e0614ffb565b14611913575f8281526005840160205260409081902054905163170cc93360e21b81526107129160ff1690600401615193565b5f828152600484016020526040812061192b91614ce6565b5f828152600584016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917f8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b910160405180910390a250505050565b5f6119c36126f7565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff1660038111156119fc576119fc614ffb565b6003811115611a0d57611a0d614ffb565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611a8657611a86614ffb565b14158015611aa75750600381516003811115611aa457611aa4614ffb565b14155b15611ac8578051604051633b0d540d60e21b815261071291906004016157bf565b5f611ad68260400151611d4f565b905080606001516001600160401b03165f03611b08576040516339b894f960e21b815260048101859052602401610712565b60408083015160608301516080840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__9363854a893f93611b7693906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611b90573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611bb7919081019061529c565b6040518263ffffffff1660e01b8152600401611bd391906152cd565b6020604051808303815f875af1158015611bef573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eb191906152df565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb046020526040902080545f80516020615dbe8339815191529190611c5a9061515b565b90505f03611c7e5760405163089938b360e11b815260048101839052602401610712565b60015f838152600580840160205260409091205460ff1690811115611ca557611ca5614ffb565b14611cd8575f8281526005820160205260409081902054905163170cc93360e21b81526107129160ff1690600401615193565b5f8281526004808301602052604091829020915163ee5b48eb60e01b81526005600160991b019263ee5b48eb92611d0f9201615853565b6020604051808303815f875af1158015611d2b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058691906152df565b611d57614d1d565b5f8281525f80516020615dde833981519152602052604090819020815160e0810190925280545f80516020615dbe833981519152929190829060ff166005811115611da457611da4614ffb565b6005811115611db557611db5614ffb565b8152602001600182018054611dc99061515b565b80601f0160208091040260200160405190810160405280929190818152602001828054611df59061515b565b8015611e405780601f10611e1757610100808354040283529160200191611e40565b820191905f5260205f20905b815481529060010190602001808311611e2357829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b6001600160a01b038116611ed05760405163caa903f960e01b81526001600160a01b0382166004820152602401610712565b5f611ed96126f7565b5f8481526007820160205260409020549091506001600160a01b03610100909104163314611f075733611596565b336001600160a01b03831603611f36575f928352600901602052506040902080546001600160a01b0319169055565b5f838152600982016020526040902080546001600160a01b0384166001600160a01b0319909116179055505050565b6040515f905f80516020615dbe833981519152907fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0690611fa890869086906158dd565b90815260200160405180910390205491505092915050565b611fcc84848484611fec565b61081557604051631036cf9160e11b815260048101859052602401610712565b5f80611ff66126f7565b5f878152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561202f5761202f614ffb565b600381111561204057612040614ffb565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6120b682611d4f565b90506002835160038111156120cd576120cd614ffb565b146120ee578251604051633b0d540d60e21b815261071291906004016157bf565b60208301516001600160a01b0316331461218a575f8281526006850160205260409020546001600160a01b031633146121275733611596565b5f82815260068501602052604090205460a082015161215691600160b01b90046001600160401b031690615574565b6001600160401b031642101561218a5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610712565b60028151600581111561219f5761219f614ffb565b036122c157600284015460808401516121c1916001600160401b031690615574565b6001600160401b03164210156121f55760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610712565b8715612207576122058288612458565b505b5f8981526007850160205260409020805460ff191660031790556060830151608082015161224091849161223b91906158ec565b613389565b505f8a8152600786016020526040812060020180546001600160401b03909316600160c01b026001600160c01b039093169290921790915561228384888c613560565b9050828a7f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed5760405160405180910390a315159450610fd69350505050565b6004815160058111156122d6576122d6614ffb565b036122fd576122e683878b613560565b506122f089612f1d565b6001945050505050610fd6565b805160405163170cc93360e21b81526107129190600401615193565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa15801561237d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526123a4919081019061590c565b91509150806123c657604051636b2f19e960e01b815260040160405180910390fd5b8151156123ec578151604051636ba589a560e01b81526004810191909152602401610712565b60208201516001600160a01b031615612428576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610712565b5092915050565b5f806124396126f7565b5f938452600601602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa1580156124a3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526124ca919081019061590c565b91509150806124ec57604051636b2f19e960e01b815260040160405180910390fd5b5f6124f56126f7565b6005810154845191925014612523578251604051636ba589a560e01b81526004810191909152602401610712565b60208301516001600160a01b03161561255f576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610712565b60208301516001600160a01b03161561259b576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610712565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246386604001516040518263ffffffff1660e01b81526004016125d891906152cd565b6040805180830381865af41580156125f2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612616919061599c565b9150915081881461263d5760405163089938b360e11b815260048101899052602401610712565b5f8881526006840160205260409020600101546001600160401b0390811690821611156126ce575f888152600684016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a26126ec565b505f8781526006830160205260409020600101546001600160401b03165b979650505050505050565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0090565b5f806127256126f7565b90505f61273187613745565b905061273c8761242f565b61274b57600192505050610fd6565b5f8781526006830160205260409020546001600160a01b031633146127705733611596565b5f87815260068301602052604090205460a082015161279f91600160b01b90046001600160401b031690615574565b6001600160401b03168160c001516001600160401b031610156127e65760c081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610712565b5f86156127fe576127f78887612458565b905061281c565b505f8781526006830160205260409020600101546001600160401b03165b600483015460408301515f916001600160a01b031690634f22429f90612841906112b2565b60a086015160c087015160405160e085901b6001600160e01b031916815260048101939093526001600160401b03918216602484018190526044840152811660648301528516608482015260a401602060405180830381865afa1580156128aa573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128ce91906152df565b90508084600a015f8b81526020019081526020015f205f8282546128f291906159bf565b90915550506001600160a01b0386161561292f575f898152600b85016020526040902080546001600160a01b0319166001600160a01b0388161790555b151598975050505050505050565b5f80516020615dfe83398151915280546001190161296e57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f61297d614d1d565b5f80516020615dbe8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f6129b088612319565b604001516040518263ffffffff1660e01b81526004016129d091906152cd565b6040805180830381865af41580156129ea573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a0e9190615830565b915091508015612a3557604051632d07135360e01b81528115156004820152602401610712565b5f82815260058085016020526040808320815160e08101909252805491929091839160ff90911690811115612a6c57612a6c614ffb565b6005811115612a7d57612a7d614ffb565b8152602001600182018054612a919061515b565b80601f0160208091040260200160405190810160405280929190818152602001828054612abd9061515b565b8015612b085780601f10612adf57610100808354040283529160200191612b08565b820191905f5260205f20905b815481529060010190602001808311612aeb57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612b7357612b73614ffb565b14158015612b945750600181516005811115612b9157612b91614ffb565b14155b15612bb557805160405163170cc93360e21b81526107129190600401615193565b600381516005811115612bca57612bca614ffb565b03612bd85760048152612bdd565b600581525b836006018160200151604051612bf39190615450565b90815260408051602092819003830190205f90819055858152600587810190935220825181548493839160ff1916906001908490811115612c3657612c36614ffb565b021790555060208201516001820190612c4f90826154a5565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff19169190921617905580516005811115612cf557612cf5614ffb565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b5f612d356126f7565b5f838152600a82016020526040812080549190559091506108158482613a2a565b5f80516020615d9e83398151915254611534906001600160a01b03168383613aa0565b60015f80516020615dfe83398151915255565b5f80612d966126f7565b600281015490915061ffff600160401b90910481169086161080612dbf575061271061ffff8616115b15612de357604051635f12e6c360e11b815261ffff86166004820152602401610712565b60028101546001600160401b039081169085161015612e1f576040516202a06d60e11b81526001600160401b0385166004820152602401610712565b8054831080612e315750806001015483115b15612e525760405163222d164360e21b815260048101849052602401610712565b5f612e5c84613aff565b90505f612e6882610e51565b90505f612e758983613b22565b5f818152600695909501602052604090942080546001600160b01b0319163317600160a01b61ffff9a909a16999099029890981767ffffffffffffffff60b01b1916600160b01b6001600160401b03989098169790970296909617875550506001909401805467ffffffffffffffff1916905550919392505050565b612efd8484848461271b565b61081557604051635bff683f60e11b815260048101859052602401610712565b5f612f266126f7565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115612f5f57612f5f614ffb565b6003811115612f7057612f70614ffb565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015281015190915061300d7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb01546001600160401b031690565b826080015161301c9190615574565b6001600160401b03164210156130505760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610712565b5f848152600784016020908152604080832080546001600160a81b031916815560018101849055600201839055600986019091529020546001600160a01b0316806130a0575060208201516130be565b5f858152600985016020526040902080546001600160a01b03191690555b5f806130cb838886614094565b915091506130e48560200151610f8587606001516112b2565b6040805183815260208101839052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f806131346126f7565b90505f6131436102d185613aff565b90505f61314f87611d4f565b905061315a8761242f565b61317a576040516330efa98b60e01b815260048101889052602401610712565b60028151600581111561318f5761318f614ffb565b146131b057805160405163170cc93360e21b81526107129190600401615193565b5f8282608001516131c19190615574565b905083600201600a9054906101000a90046001600160401b031682604001516131ea91906155b6565b6001600160401b0316816001600160401b0316111561322757604051636d51fe0560e11b81526001600160401b0382166004820152602401610712565b5f806132338a84613389565b915091505f8a8360405160200161326192919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260078b019093529120805491925060019160ff1916828002179055505f8181526007880160209081526040918290208054610100600160a81b0319166101006001600160a01b038f16908102919091178255600182018f9055600290910180546001600160401b038b81166001600160c01b03199092168217600160801b8a8316908102919091176001600160c01b031690935585519283528916938201939093529283019190915260608201849052908c9083907fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a2234269060800160405180910390a496505050505050505b9392505050565b613377614141565b6133808261418c565b61153481614206565b5f8281525f80516020615dde833981519152602052604081206002015481905f80516020615dbe83398151915290600160801b90046001600160401b03166133d1858261426e565b5f6133db876144d8565b5f888152600585016020526040808220600201805467ffffffffffffffff60801b1916600160801b6001600160401b038c811691820292909217909255915163854a893f60e01b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af4158015613484573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526134ab919081019061529c565b6040518263ffffffff1660e01b81526004016134c791906152cd565b6020604051808303815f875af11580156134e3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061350791906152df565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f8061356a6126f7565b90505f61357a8660400151611d4f565b90505f60038251600581111561359257613592614ffb565b14806135b057506004825160058111156135ae576135ae614ffb565b145b156135c0575060c08101516135fd565b6002825160058111156135d5576135d5614ffb565b036135e15750426135fd565b815160405163170cc93360e21b81526107129190600401615193565b86608001516001600160401b0316816001600160401b031611613625575f9350505050613368565b600483015460608801515f916001600160a01b031690634f22429f9061364a906112b2565b60a086015160808c01516040808e01515f90815260068b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa1580156136cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136ef91906152df565b5f878152600886016020526040902081905590506001600160a01b038716156126ec575f868152600985016020526040902080546001600160a01b0319166001600160a01b038916179055979650505050505050565b61374d614d1d565b5f8281525f80516020615dde8339815191526020526040808220815160e0810190925280545f80516020615dbe83398151915293929190829060ff16600581111561379a5761379a614ffb565b60058111156137ab576137ab614ffb565b81526020016001820180546137bf9061515b565b80601f01602080910402602001604051908101604052809291908181526020018280546137eb9061515b565b80156138365780601f1061380d57610100808354040283529160200191613836565b820191905f5260205f20905b81548152906001019060200180831161381957829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a090910152909150815160058111156138a4576138a4614ffb565b146138d7575f8481526005830160205260409081902054905163170cc93360e21b81526107129160ff1690600401615193565b60038152426001600160401b031660c08201525f84815260058381016020526040909120825181548493839160ff191690600190849081111561391c5761391c614ffb565b02179055506020820151600182019061393590826154a5565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790555f6139d38582613389565b6080840151604080516001600160401b03909216825242602083015291935083925087917ffbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42910160405180910390a3509392505050565b5f5f80516020615d9e83398151915280546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b158015613a85575f80fd5b505af1158015613a97573d5f803e3d5ffd5b50505050505050565b6040516001600160a01b0383811660248301526044820183905261058691859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061454d565b5f610e9f825f80516020615d9e833981519152546001600160a01b0316906145ae565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07545f9060ff16613b6657604051637fab81e560e01b815260040160405180910390fd5b5f80516020615dbe83398151915242613b856060860160408701614f5f565b6001600160401b0316111580613bbf5750613ba36202a300426159bf565b613bb36060860160408701614f5f565b6001600160401b031610155b15613bf957613bd46060850160408601614f5f565b604051635879da1360e11b81526001600160401b039091166004820152602401610712565b60038101546001600160401b0390613c1c90600160401b900482168583166159bf565b1115613c4657604051633e1a785160e01b81526001600160401b0384166004820152602401610712565b613c5b613c5660608601866159d2565b6145ba565b613c6b613c5660808601866159d2565b6030613c7a60208601866159e6565b905014613cac57613c8e60208501856159e6565b6040516326475b2f60e11b8152610712925060040190815260200190565b613cb684806159e6565b90505f03613ce357613cc884806159e6565b604051633e08a12560e11b8152600401610712929190615a28565b5f60068201613cf286806159e6565b604051613d009291906158dd565b90815260200160405180910390205414613d3957613d1e84806159e6565b60405163a41f772f60e01b8152600401610712929190615a28565b613d43835f61426e565b6040805160e08101909152815481525f90819073__$fd0c147b4031eef6079b0498cbafa865f0$__9063eb97ce519060208101613d808a806159e6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602090810190613dc8908b018b6159e6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602001613e1160608b0160408c01614f5f565b6001600160401b03168152602001613e2c60608b018b6159d2565b613e3590615a3b565b8152602001613e4760808b018b6159d2565b613e5090615a3b565b8152602001886001600160401b03168152506040518263ffffffff1660e01b8152600401613e7e9190615b68565b5f60405180830381865af4158015613e98573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613ebf9190810190615c1f565b5f82815260048601602052604090209193509150613edd82826154a5565b508160068401613eed88806159e6565b604051613efb9291906158dd565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90613f379085906004016152cd565b6020604051808303815f875af1158015613f53573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f7791906152df565b5f8481526005860160205260409020805460ff191660011790559050613f9d87806159e6565b5f858152600587016020526040902060010191613fbb919083615c62565b505f83815260058501602052604090206002810180546001600160c01b0319166001600160401b038916908117600160801b91909102176001600160c01b03169055600301805467ffffffffffffffff191690558061401a88806159e6565b6040516140289291906158dd565b6040518091039020847fd8a184af94a03e121609cc5f803a446236793e920c7945abc6ba355c8a30cb49898b60400160208101906140669190614f5f565b604080516001600160401b0393841681529290911660208301520160405180910390a4509095945050505050565b5f805f61409f6126f7565b5f86815260088201602052604081208054908290559192509081908015614133575f878152600685016020526040902054612710906140e990600160a01b900461ffff1683615819565b6140f391906157a0565b91508184600a015f8981526020019081526020015f205f82825461411791906159bf565b9091555061412790508282615d1b565b92506141338984613a2a565b509097909650945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661418a57604051631afcd79f60e31b815260040160405180910390fd5b565b614194614141565b61419d81614723565b6141a561473c565b610fa6606082013560808301356141c260c0850160a08601614f5f565b6141d260e0860160c08701615d2e565b6141e3610100870160e08801615d47565b6101008701356141fb61014089016101208a016152f6565b88610140013561474c565b61420e614141565b5f80516020615d9e8339815191526001600160a01b03821661424e57604051637330680360e01b81526001600160a01b0383166004820152602401610712565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f80516020615dbe8339815191525f6001600160401b0380841690851611156142a25761429b83856158ec565b90506142af565b6142ac84846158ec565b90505b60408051608081018252600284015480825260038501546001600160401b038082166020850152600160401b8204811694840194909452600160801b900490921660608201524291158061431c575060018401548151614318916001600160401b0316906159bf565b8210155b15614344576001600160401b0380841660608301528282526040820151166020820152614363565b82816060018181516143569190615574565b6001600160401b03169052505b60608101516143739060646155b6565b602082015160018601546001600160401b03929092169161439e9190600160401b900460ff166155b6565b6001600160401b031610156143d757606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610712565b85816040018181516143e99190615574565b6001600160401b03169052506040810180518691906144099083906158ec565b6001600160401b03169052506001840154604082015160649161443791600160401b90910460ff16906155b6565b6001600160401b03161015614470576040808201519051633e1a785160e01b81526001600160401b039091166004820152602401610712565b8051600285015560208101516003909401805460408301516060909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919097161717919091169390931790925550505050565b5f8181525f80516020615dde8339815191526020526040812060020180545f80516020615dbe833981519152919060089061452290600160401b90046001600160401b0316615d67565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b5f6145616001600160a01b03841683614931565b905080515f141580156145855750808060200190518101906145839190615d82565b155b1561058657604051635274afe760e01b81526001600160a01b0384166004820152602401610712565b5f61336883338461493e565b6145c76020820182614eba565b63ffffffff161580156145e757506145e26020820182615311565b151590505b1561462e576145f96020820182614eba565b6146066020830183615311565b60405163c08a0f1d60e01b815263ffffffff9093166004840152602483015250604401610712565b61463b6020820182615311565b905061464a6020830183614eba565b63ffffffff161115614663576145f96020820182614eba565b60015b6146736020830183615311565b9050811015611534576146896020830183615311565b614694600184615d1b565b8181106146a3576146a3615356565b90506020020160208101906146b891906152f6565b6001600160a01b03166146ce6020840184615311565b838181106146de576146de615356565b90506020020160208101906146f391906152f6565b6001600160a01b0316101561471b57604051630dbc8d5f60e31b815260040160405180910390fd5b600101614666565b61472b614141565b614733614aa1565b610fa681614aa9565b614744614141565b61418a614b91565b614754614141565b5f61475d6126f7565b905061ffff86161580614775575061271061ffff8716115b1561479957604051635f12e6c360e11b815261ffff87166004820152602401610712565b878911156147bd5760405163222d164360e21b8152600481018a9052602401610712565b60ff851615806147d05750600a60ff8616115b156147f35760405163170db35960e31b815260ff86166004820152602401610712565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb01546001600160401b03166001600160401b0316876001600160401b0316101561485a576040516202a06d60e11b81526001600160401b0388166004820152602401610712565b835f0361487a5760405163a733007160e01b815260040160405180910390fd5b8161489b57604051632f6bd1db60e01b815260048101839052602401610712565b97885560018801969096556002870180546001600160401b039690961669ffffffffffffffffffff1990961695909517600160401b61ffff95909516949094029390931767ffffffffffffffff60501b191660ff92909216600160501b029190911790925560038401919091556004830180546001600160a01b0319166001600160a01b03909216919091179055600590910155565b606061336883835f614b99565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015614984573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906149a891906152df565b90506149bf6001600160a01b038616853086614c28565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015614a03573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614a2791906152df565b9050818111614a8d5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610712565b614a978282615d1b565b9695505050505050565b61418a614141565b614ab1614141565b80355f80516020615dbe8339815191529081556014614ad66060840160408501615d47565b60ff161180614af55750614af06060830160408401615d47565b60ff16155b15614b2957614b0a6060830160408401615d47565b604051634a59bbff60e11b815260ff9091166004820152602401610712565b614b396060830160408401615d47565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614b6a6040830160208401614f5f565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612d79614141565b606081471015614bbe5760405163cd78605960e01b8152306004820152602401610712565b5f80856001600160a01b03168486604051614bd99190615450565b5f6040518083038185875af1925050503d805f8114614c13576040519150601f19603f3d011682016040523d82523d5f602084013e614c18565b606091505b5091509150614a97868383614c61565b6040516001600160a01b0384811660248301528381166044830152606482018390526108159186918216906323b872dd90608401613acd565b606082614c7657614c7182614cbd565b613368565b8151158015614c8d57506001600160a01b0384163b155b15614cb657604051639996b31560e01b81526001600160a01b0385166004820152602401610712565b5080613368565b805115614ccd5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b508054614cf29061515b565b5f825580601f10614d01575050565b601f0160209004905f5260205f2090810190610fa69190614d5a565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b80821115614d6e575f8155600101614d5b565b5090565b8015158114610fa6575f80fd5b803563ffffffff81168114614d92575f80fd5b919050565b5f805f60608486031215614da9575f80fd5b833592506020840135614dbb81614d72565b9150614dc960408501614d7f565b90509250925092565b5f60208284031215614de2575f80fd5b5035919050565b5f8060408385031215614dfa575f80fd5b82356001600160401b03811115614e0f575f80fd5b830160808186031215614e20575f80fd5b9150614e2e60208401614d7f565b90509250929050565b5f8060408385031215614e48575f80fd5b82359150614e2e60208401614d7f565b6001600160a01b0381168114610fa6575f80fd5b5f805f8060808587031215614e7f575f80fd5b843593506020850135614e9181614d72565b9250614e9f60408601614d7f565b91506060850135614eaf81614e58565b939692955090935050565b5f60208284031215614eca575f80fd5b61336882614d7f565b803561ffff81168114614d92575f80fd5b6001600160401b0381168114610fa6575f80fd5b5f805f8060808587031215614f0b575f80fd5b84356001600160401b03811115614f20575f80fd5b850160a08188031215614f31575f80fd5b9350614f3f60208601614ed3565b92506040850135614f4f81614ee4565b9396929550929360600135925050565b5f60208284031215614f6f575f80fd5b813561336881614ee4565b5f8060408385031215614f8b575f80fd5b823591506020830135614f9d81614e58565b809150509250929050565b5f8060408385031215614fb9575f80fd5b50508035926020909101359150565b5f80828403610180811215614fdb575f80fd5b61016080821215614fea575f80fd5b8493508301359050614f9d81614e58565b634e487b7160e01b5f52602160045260245ffd5b6006811061501f5761501f614ffb565b9052565b5f5b8381101561503d578181015183820152602001615025565b50505f910152565b5f815180845261505c816020860160208601615023565b601f01601f19169290920160200192915050565b6020815261508260208201835161500f565b5f602083015160e0604084015261509d610100840182615045565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b5f8060208385031215615101575f80fd5b82356001600160401b0380821115615117575f80fd5b818501915085601f83011261512a575f80fd5b813581811115615138575f80fd5b866020828501011115615149575f80fd5b60209290920196919550909350505050565b600181811c9082168061516f57607f821691505b60208210810361518d57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208101610e9f828461500f565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156151d7576151d76151a1565b60405290565b604080519081016001600160401b03811182821017156151d7576151d76151a1565b604051601f8201601f191681016001600160401b0381118282101715615227576152276151a1565b604052919050565b5f6001600160401b03821115615247576152476151a1565b50601f01601f191660200190565b5f82601f830112615264575f80fd5b81516152776152728261522f565b6151ff565b81815284602083860101111561528b575f80fd5b610fd6826020830160208701615023565b5f602082840312156152ac575f80fd5b81516001600160401b038111156152c1575f80fd5b610fd684828501615255565b602081525f6133686020830184615045565b5f602082840312156152ef575f80fd5b5051919050565b5f60208284031215615306575f80fd5b813561336881614e58565b5f808335601e19843603018112615326575f80fd5b8301803591506001600160401b0382111561533f575f80fd5b6020019150600581901b3603821315613559575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e1983360301811261537e575f80fd5b9190910192915050565b5f82601f830112615397575f80fd5b81356153a56152728261522f565b8181528460208386010111156153b9575f80fd5b816020850160208301375f918101602001919091529392505050565b5f606082360312156153e5575f80fd5b6153ed6151b5565b82356001600160401b0380821115615403575f80fd5b61540f36838701615388565b83526020850135915080821115615424575f80fd5b5061543136828601615388565b602083015250604083013561544581614ee4565b604082015292915050565b5f825161537e818460208701615023565b601f82111561058657805f5260205f20601f840160051c810160208510156154865750805b601f840160051c820191505b81811015610eb1575f8155600101615492565b81516001600160401b038111156154be576154be6151a1565b6154d2816154cc845461515b565b84615461565b602080601f831160018114615505575f84156154ee5750858301515b5f19600386901b1c1916600185901b17855561110e565b5f85815260208120601f198616915b8281101561553357888601518255948401946001909101908401615514565b508582101561555057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b0381811683821601908082111561242857612428615560565b5f63ffffffff8083168181036155ac576155ac615560565b6001019392505050565b6001600160401b038181168382160280821691908281146155d9576155d9615560565b505092915050565b5f808335601e198436030181126155f6575f80fd5b83016020810192503590506001600160401b03811115615614575f80fd5b803603821315613559575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b878110156156ff57848303601f19018952813536889003605e19018112615686575f80fd5b8701606061569482806155e1565b8287526156a48388018284615622565b925050506156b4868301836155e1565b868303888801526156c6838284615622565b9250505060408083013592506156db83614ee4565b6001600160401b039290921694909101939093529783019790830190600101615661565b5090979650505050505050565b6020815281356020820152602082013560408201525f604083013561573081614e58565b6001600160a01b031660608381019190915283013536849003601e19018112615757575f80fd5b83016020810190356001600160401b03811115615772575f80fd5b8060051b3603821315615783575f80fd5b60808085015261579760a08501828461564a565b95945050505050565b5f826157ba57634e487b7160e01b5f52601260045260245ffd5b500490565b60208101600483106157d3576157d3614ffb565b91905290565b5f805f606084860312156157eb575f80fd5b8351925060208401516157fd81614ee4565b604085015190925061580e81614ee4565b809150509250925092565b8082028115828204841417610e9f57610e9f615560565b5f8060408385031215615841575f80fd5b825191506020830151614f9d81614d72565b5f60208083525f84546158658161515b565b806020870152604060018084165f811461588657600181146158a2576158cf565b60ff19851660408a0152604084151560051b8a010195506158cf565b895f5260205f205f5b858110156158c65781548b82018601529083019088016158ab565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b6001600160401b0382811682821603908082111561242857612428615560565b5f806040838503121561591d575f80fd5b82516001600160401b0380821115615933575f80fd5b9084019060608287031215615946575f80fd5b61594e6151b5565b82518152602083015161596081614e58565b6020820152604083015182811115615976575f80fd5b61598288828601615255565b6040830152508094505050506020830151614f9d81614d72565b5f80604083850312156159ad575f80fd5b825191506020830151614f9d81614ee4565b80820180821115610e9f57610e9f615560565b5f8235603e1983360301811261537e575f80fd5b5f808335601e198436030181126159fb575f80fd5b8301803591506001600160401b03821115615a14575f80fd5b602001915036819003821315613559575f80fd5b602081525f610fd6602083018486615622565b5f60408236031215615a4b575f80fd5b615a536151dd565b615a5c83614d7f565b81526020808401356001600160401b0380821115615a78575f80fd5b9085019036601f830112615a8a575f80fd5b813581811115615a9c57615a9c6151a1565b8060051b9150615aad8483016151ff565b8181529183018401918481019036841115615ac6575f80fd5b938501935b83851015615af05784359250615ae083614e58565b8282529385019390850190615acb565b94860194909452509295945050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015615b5d5784516001600160a01b03168252938301936001929092019190830190615b34565b509695505050505050565b60208152815160208201525f602083015160e06040840152615b8e610100840182615045565b90506040840151601f1980858403016060860152615bac8383615045565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152615bdc8383615b02565b925060a08601519150808584030160c086015250615bfa8282615b02565b91505060c0840151615c1760e08501826001600160401b03169052565b509392505050565b5f8060408385031215615c30575f80fd5b8251915060208301516001600160401b03811115615c4c575f80fd5b615c5885828601615255565b9150509250929050565b6001600160401b03831115615c7957615c796151a1565b615c8d83615c87835461515b565b83615461565b5f601f841160018114615cbe575f8515615ca75750838201355b5f19600387901b1c1916600186901b178355610eb1565b5f83815260208120601f198716915b82811015615ced5786850135825560209485019460019092019101615ccd565b5086821015615d09575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81810381811115610e9f57610e9f615560565b5f60208284031215615d3e575f80fd5b61336882614ed3565b5f60208284031215615d57575f80fd5b813560ff81168114613368575f80fd5b5f6001600160401b038083168181036155ac576155ac615560565b5f60208284031215615d92575f80fd5b815161336881614d7256fe6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb059b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220600f26530acb9344f9685711352fcf01384e451856d9210ed4109b133d1b702f64736f6c63430008190033",
}

// ERC20TokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenStakingManagerMetaData.ABI instead.
var ERC20TokenStakingManagerABI = ERC20TokenStakingManagerMetaData.ABI

// ERC20TokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenStakingManagerMetaData.Bin instead.
var ERC20TokenStakingManagerBin = ERC20TokenStakingManagerMetaData.Bin

// DeployERC20TokenStakingManager deploys a new Ethereum contract, binding an instance of ERC20TokenStakingManager to it.
func DeployERC20TokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenStakingManager, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	ERC20TokenStakingManagerBin = strings.ReplaceAll(ERC20TokenStakingManagerBin, "__$fd0c147b4031eef6079b0498cbafa865f0$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// ERC20TokenStakingManager is an auto generated Go binding around an Ethereum contract.
type ERC20TokenStakingManager struct {
	ERC20TokenStakingManagerCaller     // Read-only binding to the contract
	ERC20TokenStakingManagerTransactor // Write-only binding to the contract
	ERC20TokenStakingManagerFilterer   // Log filterer for contract events
}

// ERC20TokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenStakingManagerSession struct {
	Contract     *ERC20TokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenStakingManagerCallerSession struct {
	Contract *ERC20TokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20TokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenStakingManagerTransactorSession struct {
	Contract     *ERC20TokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenStakingManagerRaw struct {
	Contract *ERC20TokenStakingManager // Generic contract binding to access the raw methods on
}

// ERC20TokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCallerRaw struct {
	Contract *ERC20TokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactorRaw struct {
	Contract *ERC20TokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenStakingManager creates a new instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManager(address common.Address, backend bind.ContractBackend) (*ERC20TokenStakingManager, error) {
	contract, err := bindERC20TokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// NewERC20TokenStakingManagerCaller creates a new read-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenStakingManagerCaller, error) {
	contract, err := bindERC20TokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCaller{contract: contract}, nil
}

// NewERC20TokenStakingManagerTransactor creates a new write-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenStakingManagerTransactor, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerTransactor{contract: contract}, nil
}

// NewERC20TokenStakingManagerFilterer creates a new log filterer instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenStakingManagerFilterer, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerFilterer{contract: contract}, nil
}

// bindERC20TokenStakingManager binds a generic wrapper to an already deployed contract.
func bindERC20TokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ERC20STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ERC20_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.GetWeight(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.GetWeight(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation0", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initialize", settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndValidation0", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorSet", conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ERC20TokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEndedIterator struct {
	Event *ERC20TokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegationEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerDelegationEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegationEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegationEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*ERC20TokenStakingManagerDelegationEnded, error) {
	event := new(ERC20TokenStakingManagerDelegationEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAddedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20TokenStakingManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorAddedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorAdded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*ERC20TokenStakingManagerDelegatorAdded, error) {
	event := new(ERC20TokenStakingManagerDelegatorAdded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegisteredIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*ERC20TokenStakingManagerDelegatorRegistered, error) {
	event := new(ERC20TokenStakingManagerDelegatorRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *ERC20TokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0xfe3e5983f71c8253fb0b678f2bc587aa8574d8f1aab9cf82b983777f5998392c.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*ERC20TokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitialValidatorCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0xfe3e5983f71c8253fb0b678f2bc587aa8574d8f1aab9cf82b983777f5998392c.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialValidatorCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0xfe3e5983f71c8253fb0b678f2bc587aa8574d8f1aab9cf82b983777f5998392c.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*ERC20TokenStakingManagerInitialValidatorCreated, error) {
	event := new(ERC20TokenStakingManagerInitialValidatorCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitializedIterator struct {
	Event *ERC20TokenStakingManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialized represents a Initialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialized(log types.Log) (*ERC20TokenStakingManagerInitialized, error) {
	event := new(ERC20TokenStakingManagerInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdatedIterator struct {
	Event *ERC20TokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerUptimeUpdatedIterator{contract: _ERC20TokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerUptimeUpdated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*ERC20TokenStakingManagerUptimeUpdated, error) {
	event := new(ERC20TokenStakingManagerUptimeUpdated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      uint64
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xd8a184af94a03e121609cc5f803a446236793e920c7945abc6ba355c8a30cb49.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint64 weight, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xd8a184af94a03e121609cc5f803a446236793e920c7945abc6ba355c8a30cb49.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint64 weight, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xd8a184af94a03e121609cc5f803a446236793e920c7945abc6ba355c8a30cb49.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint64 weight, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*ERC20TokenStakingManagerValidationPeriodCreated, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEndedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*ERC20TokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*ERC20TokenStakingManagerValidationPeriodEnded, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       uint64
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*ERC20TokenStakingManagerValidationPeriodRegistered, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             uint64
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*ERC20TokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *ERC20TokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	Weight             uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*ERC20TokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorWeightUpdateIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*ERC20TokenStakingManagerValidatorWeightUpdate, error) {
	event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMessagesMetaData contains all meta data concerning the ValidatorMessages contract.
var ValidatorMessagesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBLSPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"InvalidCodecID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actual\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"}],\"name\":\"packConversionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"packL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"packL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"validationPeriod\",\"type\":\"tuple\"}],\"name\":\"packRegisterL1ValidatorMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conversionID\",\"type\":\"bytes32\"}],\"name\":\"packSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"packValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackRegisterL1ValidatorMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x612160610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b1575f3560e01c8063854a893f11610079578063854a893f146101b257806387418b8e1461020f5780639b83546514610222578063a699c13514610242578063e1d68f3014610255578063eb97ce5114610268575f80fd5b8063021de88f146100b5578063088c2463146100e25780634d8478841461011257806350782b0f146101335780637f7c427a1461016b575b5f80fd5b6100c86100c3366004611904565b610289565b604080519283529015156020830152015b60405180910390f35b6100f56100f0366004611904565b61044a565b604080519283526001600160401b039091166020830152016100d9565b610125610120366004611904565b61063b565b6040519081526020016100d9565b610146610141366004611904565b6107c8565b604080519384526001600160401b0392831660208501529116908201526060016100d9565b6101a561017936600461193d565b604080515f60208201819052602282015260268082019390935281518082039093018352604601905290565b6040516100d991906119a1565b6101a56101c03660046119d5565b604080515f6020820152600360e01b602282015260268101949094526001600160c01b031960c093841b811660468601529190921b16604e830152805180830360360181526056909201905290565b6101a561021d366004611a0e565b610a1e565b610235610230366004611904565b610bff565b6040516100d99190611aaa565b6101a5610250366004611b61565b61154a565b6101a5610263366004611b93565b61158e565b61027b610276366004611ca1565b6115c4565b6040516100d9929190611d9d565b5f8082516027146102c457825160405163cc92daa160e01b815263ffffffff9091166004820152602760248201526044015b60405180910390fd5b5f805b6002811015610313576102db816001611dc9565b6102e6906008611ddc565b61ffff168582815181106102fc576102fc611df3565b016020015160f81c901b91909117906001016102c7565b5061ffff81161561033d5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561039857610354816003611dc9565b61035f906008611ddc565b63ffffffff1686610371836002611e07565b8151811061038157610381611df3565b016020015160f81c901b9190911790600101610340565b5063ffffffff81166002146103c057604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610415576103d781601f611dc9565b6103e2906008611ddc565b876103ee836006611e07565b815181106103fe576103fe611df3565b016020015160f81c901b91909117906001016103c3565b505f8660268151811061042a5761042a611df3565b016020015191976001600160f81b03199092161515965090945050505050565b5f808251602e1461048057825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044016102bb565b5f805b60028110156104cf57610497816001611dc9565b6104a2906008611ddc565b61ffff168582815181106104b8576104b8611df3565b016020015160f81c901b9190911790600101610483565b5061ffff8116156104f95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561055457610510816003611dc9565b61051b906008611ddc565b63ffffffff168661052d836002611e07565b8151811061053d5761053d611df3565b016020015160f81c901b91909117906001016104fc565b5063ffffffff81161561057a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156105cf5761059181601f611dc9565b61059c906008611ddc565b876105a8836006611e07565b815181106105b8576105b8611df3565b016020015160f81c901b919091179060010161057d565b505f805b600881101561062e576105e7816007611dc9565b6105f2906008611ddc565b6001600160401b031688610607836026611e07565b8151811061061757610617611df3565b016020015160f81c901b91909117906001016105d3565b5090969095509350505050565b5f815160261461067057815160405163cc92daa160e01b815263ffffffff9091166004820152602660248201526044016102bb565b5f805b60028110156106bf57610687816001611dc9565b610692906008611ddc565b61ffff168482815181106106a8576106a8611df3565b016020015160f81c901b9190911790600101610673565b5061ffff8116156106e95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561074457610700816003611dc9565b61070b906008611ddc565b63ffffffff168561071d836002611e07565b8151811061072d5761072d611df3565b016020015160f81c901b91909117906001016106ec565b5063ffffffff81161561076a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156107bf5761078181601f611dc9565b61078c906008611ddc565b86610798836006611e07565b815181106107a8576107a8611df3565b016020015160f81c901b919091179060010161076d565b50949350505050565b5f805f83516036146107ff57835160405163cc92daa160e01b815263ffffffff9091166004820152603660248201526044016102bb565b5f805b600281101561084e57610816816001611dc9565b610821906008611ddc565b61ffff1686828151811061083757610837611df3565b016020015160f81c901b9190911790600101610802565b5061ffff8116156108785760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b60048110156108d35761088f816003611dc9565b61089a906008611ddc565b63ffffffff16876108ac836002611e07565b815181106108bc576108bc611df3565b016020015160f81c901b919091179060010161087b565b5063ffffffff81166003146108fb57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156109505761091281601f611dc9565b61091d906008611ddc565b88610929836006611e07565b8151811061093957610939611df3565b016020015160f81c901b91909117906001016108fe565b505f805b60088110156109af57610968816007611dc9565b610973906008611ddc565b6001600160401b031689610988836026611e07565b8151811061099857610998611df3565b016020015160f81c901b9190911790600101610954565b505f805b6008811015610a0e576109c7816007611dc9565b6109d2906008611ddc565b6001600160401b03168a6109e783602e611e07565b815181106109f7576109f7611df3565b016020015160f81c901b91909117906001016109b3565b5091989097509095509350505050565b60605f80833560208501356014610a3a87870160408901611e1a565b610a476060890189611e33565b60405160f09790971b6001600160f01b0319166020880152602287019590955250604285019290925260e090811b6001600160e01b0319908116606286015260609290921b6bffffffffffffffffffffffff191660668501529190911b16607a820152607e0160405160208183030381529060405290505f5b610acd6060850185611e33565b9050811015610bf85781610ae46060860186611e33565b83818110610af457610af4611df3565b9050602002810190610b069190611e7f565b610b109080611e9d565b9050610b1f6060870187611e33565b84818110610b2f57610b2f611df3565b9050602002810190610b419190611e7f565b610b4b9080611e9d565b610b586060890189611e33565b86818110610b6857610b68611df3565b9050602002810190610b7a9190611e7f565b610b88906020810190611e9d565b610b9560608b018b611e33565b88818110610ba557610ba5611df3565b9050602002810190610bb79190611e7f565b610bc8906060810190604001611edf565b604051602001610bde9796959493929190611ef8565b60408051601f198184030181529190529150600101610ac0565b5092915050565b610c076117b1565b5f610c106117b1565b5f805b6002811015610c6e57610c27816001611dc9565b610c32906008611ddc565b61ffff1686610c4763ffffffff871684611e07565b81518110610c5757610c57611df3565b016020015160f81c901b9190911790600101610c13565b5061ffff811615610c985760405163407b587360e01b815261ffff821660048201526024016102bb565b610ca3600284611f61565b9250505f805b6004811015610d0857610cbd816003611dc9565b610cc8906008611ddc565b63ffffffff16868563ffffffff1683610ce19190611e07565b81518110610cf157610cf1611df3565b016020015160f81c901b9190911790600101610ca9565b5063ffffffff8116600114610d3057604051635b60892f60e01b815260040160405180910390fd5b610d3b600484611f61565b9250505f805b6020811015610d9857610d5581601f611dc9565b610d60906008611ddc565b86610d7163ffffffff871684611e07565b81518110610d8157610d81611df3565b016020015160f81c901b9190911790600101610d41565b50808252610da7602084611f61565b9250505f805b6004811015610e0c57610dc1816003611dc9565b610dcc906008611ddc565b63ffffffff16868563ffffffff1683610de59190611e07565b81518110610df557610df5611df3565b016020015160f81c901b9190911790600101610dad565b50610e18600484611f61565b92505f8163ffffffff166001600160401b03811115610e3957610e3961180b565b6040519080825280601f01601f191660200182016040528015610e63576020820181803683370190505b5090505f5b8263ffffffff16811015610ed25786610e8763ffffffff871683611e07565b81518110610e9757610e97611df3565b602001015160f81c60f81b828281518110610eb457610eb4611df3565b60200101906001600160f81b03191690815f1a905350600101610e68565b5060208301819052610ee48285611f61565b604080516030808252606082019092529195505f92506020820181803683370190505090505f5b6030811015610f705786610f2563ffffffff871683611e07565b81518110610f3557610f35611df3565b602001015160f81c60f81b828281518110610f5257610f52611df3565b60200101906001600160f81b03191690815f1a905350600101610f0b565b5060408301819052610f83603085611f61565b9350505f805b6008811015610fe957610f9d816007611dc9565b610fa8906008611ddc565b6001600160401b031687610fc263ffffffff881684611e07565b81518110610fd257610fd2611df3565b016020015160f81c901b9190911790600101610f89565b506001600160401b0381166060840152611004600885611f61565b9350505f805f5b600481101561106a5761101f816003611dc9565b61102a906008611ddc565b63ffffffff16888763ffffffff16836110439190611e07565b8151811061105357611053611df3565b016020015160f81c901b919091179060010161100b565b50611076600486611f61565b94505f5b60048110156110d95761108e816003611dc9565b611099906008611ddc565b63ffffffff16888763ffffffff16836110b29190611e07565b815181106110c2576110c2611df3565b016020015160f81c901b929092179160010161107a565b506110e5600486611f61565b94505f8263ffffffff166001600160401b038111156111065761110661180b565b60405190808252806020026020018201604052801561112f578160200160208202803683370190505b5090505f5b8363ffffffff16811015611217576040805160148082528183019092525f916020820181803683370190505090505f5b60148110156111c9578a61117e63ffffffff8b1683611e07565b8151811061118e5761118e611df3565b602001015160f81c60f81b8282815181106111ab576111ab611df3565b60200101906001600160f81b03191690815f1a905350600101611164565b505f60148201519050808484815181106111e5576111e5611df3565b6001600160a01b039092166020928302919091019091015261120860148a611f61565b98505050806001019050611134565b506040805180820190915263ffffffff9092168252602082015260808401525f80805b60048110156112995761124e816003611dc9565b611259906008611ddc565b63ffffffff16898863ffffffff16836112729190611e07565b8151811061128257611282611df3565b016020015160f81c901b919091179060010161123a565b506112a5600487611f61565b95505f5b6004811015611308576112bd816003611dc9565b6112c8906008611ddc565b63ffffffff16898863ffffffff16836112e19190611e07565b815181106112f1576112f1611df3565b016020015160f81c901b92909217916001016112a9565b50611314600487611f61565b95505f8263ffffffff166001600160401b038111156113355761133561180b565b60405190808252806020026020018201604052801561135e578160200160208202803683370190505b5090505f5b8363ffffffff16811015611446576040805160148082528183019092525f916020820181803683370190505090505f5b60148110156113f8578b6113ad63ffffffff8c1683611e07565b815181106113bd576113bd611df3565b602001015160f81c60f81b8282815181106113da576113da611df3565b60200101906001600160f81b03191690815f1a905350600101611393565b505f601482015190508084848151811061141457611414611df3565b6001600160a01b039092166020928302919091019091015261143760148b611f61565b99505050806001019050611363565b506040805180820190915263ffffffff9092168252602082015260a08501525f6114708284611f61565b61147b906014611f7e565b61148685607a611f61565b6114909190611f61565b90508063ffffffff168851146114cc57875160405163cc92daa160e01b815263ffffffff918216600482015290821660248201526044016102bb565b5f805b600881101561152f576114e3816007611dc9565b6114ee906008611ddc565b6001600160401b03168a61150863ffffffff8b1684611e07565b8151811061151857611518611df3565b016020015160f81c901b91909117906001016114cf565b506001600160401b031660c086015250929695505050505050565b6040515f6020820152600160e11b60228201526026810183905281151560f81b60468201526060906047015b60405160208183030381529060405290505b92915050565b6040515f602082018190526022820152602681018390526001600160c01b031960c083901b166046820152606090604e01611576565b5f60608260400151516030146115ed5760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f9861162e988a986001989297929690959094909390929101611fa6565b60405160208183030381529060405290505f5b846080015160200151518110156116a05781856080015160200151828151811061166d5761166d611df3565b6020026020010151604051602001611686929190612060565b60408051601f198184030181529190529150600101611641565b5060a08401518051602091820151516040516116c0938593929101612096565b60405160208183030381529060405290505f5b8460a00151602001515181101561173257818560a001516020015182815181106116ff576116ff611df3565b6020026020010151604051602001611718929190612060565b60408051601f1981840301815291905291506001016116d3565b5060c08401516040516117499183916020016120d1565b604051602081830303815290604052905060028160405161176a9190612102565b602060405180830381855afa158015611785573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906117a89190612113565b94909350915050565b6040805160e0810182525f808252606060208084018290528385018290528184018390528451808601865283815280820183905260808501528451808601909552918452908301529060a082019081525f60209091015290565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156118415761184161180b565b60405290565b60405160e081016001600160401b03811182821017156118415761184161180b565b604051601f8201601f191681016001600160401b03811182821017156118915761189161180b565b604052919050565b5f82601f8301126118a8575f80fd5b81356001600160401b038111156118c1576118c161180b565b6118d4601f8201601f1916602001611869565b8181528460208386010111156118e8575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611914575f80fd5b81356001600160401b03811115611929575f80fd5b61193584828501611899565b949350505050565b5f6020828403121561194d575f80fd5b5035919050565b5f5b8381101561196e578181015183820152602001611956565b50505f910152565b5f815180845261198d816020860160208601611954565b601f01601f19169290920160200192915050565b602081525f6119b36020830184611976565b9392505050565b80356001600160401b03811681146119d0575f80fd5b919050565b5f805f606084860312156119e7575f80fd5b833592506119f7602085016119ba565b9150611a05604085016119ba565b90509250925092565b5f60208284031215611a1e575f80fd5b81356001600160401b03811115611a33575f80fd5b8201608081850312156119b3575f80fd5b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611a9f5784516001600160a01b03168252938301936001929092019190830190611a76565b509695505050505050565b60208152815160208201525f602083015160e06040840152611ad0610100840182611976565b90506040840151601f1980858403016060860152611aee8383611976565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152611b1e8383611a44565b925060a08601519150808584030160c086015250611b3c8282611a44565b91505060c0840151611b5960e08501826001600160401b03169052565b509392505050565b5f8060408385031215611b72575f80fd5b8235915060208301358015158114611b88575f80fd5b809150509250929050565b5f8060408385031215611ba4575f80fd5b82359150611bb4602084016119ba565b90509250929050565b80356001600160a01b03811681146119d0575f80fd5b5f60408284031215611be3575f80fd5b611beb61181f565b9050813563ffffffff81168114611c00575f80fd5b81526020828101356001600160401b0380821115611c1c575f80fd5b818501915085601f830112611c2f575f80fd5b813581811115611c4157611c4161180b565b8060051b9150611c52848301611869565b8181529183018401918481019088841115611c6b575f80fd5b938501935b83851015611c9057611c8185611bbd565b82529385019390850190611c70565b808688015250505050505092915050565b5f60208284031215611cb1575f80fd5b81356001600160401b0380821115611cc7575f80fd5b9083019060e08286031215611cda575f80fd5b611ce2611847565b82358152602083013582811115611cf7575f80fd5b611d0387828601611899565b602083015250604083013582811115611d1a575f80fd5b611d2687828601611899565b604083015250611d38606084016119ba565b6060820152608083013582811115611d4e575f80fd5b611d5a87828601611bd3565b60808301525060a083013582811115611d71575f80fd5b611d7d87828601611bd3565b60a083015250611d8f60c084016119ba565b60c082015295945050505050565b828152604060208201525f6119356040830184611976565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561158857611588611db5565b808202811582820484141761158857611588611db5565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561158857611588611db5565b5f60208284031215611e2a575f80fd5b6119b382611bbd565b5f808335601e19843603018112611e48575f80fd5b8301803591506001600160401b03821115611e61575f80fd5b6020019150600581901b3603821315611e78575f80fd5b9250929050565b5f8235605e19833603018112611e93575f80fd5b9190910192915050565b5f808335601e19843603018112611eb2575f80fd5b8301803591506001600160401b03821115611ecb575f80fd5b602001915036819003821315611e78575f80fd5b5f60208284031215611eef575f80fd5b6119b3826119ba565b5f8851611f09818460208d01611954565b60e089901b6001600160e01b031916908301908152868860048301378681019050600481015f8152858782375060c09390931b6001600160c01b0319166004939094019283019390935250600c019695505050505050565b63ffffffff818116838216019080821115610bf857610bf8611db5565b63ffffffff818116838216028082169190828114611f9e57611f9e611db5565b505092915050565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651611fec81602a850160208b01611954565b86519083019061200381602a840160208b01611954565b60c087901b6001600160c01b031916602a9290910191820152612035603282018660e01b6001600160e01b0319169052565b61204e603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f8351612071818460208801611954565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f84516120a7818460208901611954565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f83516120e2818460208801611954565b60c09390931b6001600160c01b0319169190920190815260080192915050565b5f8251611e93818460208701611954565b5f60208284031215612123575f80fd5b505191905056fea264697066735822122017e9ed41d979fd99f0eb8ec53e0fc2d72d6822bb41f94895c2204db14e90d4f464736f6c63430008190033",
}

// ValidatorMessagesABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMessagesMetaData.ABI instead.
var ValidatorMessagesABI = ValidatorMessagesMetaData.ABI

// ValidatorMessagesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMessagesMetaData.Bin instead.
var ValidatorMessagesBin = ValidatorMessagesMetaData.Bin

// DeployValidatorMessages deploys a new Ethereum contract, binding an instance of ValidatorMessages to it.
func DeployValidatorMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorMessages, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorMessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// ValidatorMessages is an auto generated Go binding around an Ethereum contract.
type ValidatorMessages struct {
	ValidatorMessagesCaller     // Read-only binding to the contract
	ValidatorMessagesTransactor // Write-only binding to the contract
	ValidatorMessagesFilterer   // Log filterer for contract events
}

// ValidatorMessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorMessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMessagesSession struct {
	Contract     *ValidatorMessages // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorMessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMessagesCallerSession struct {
	Contract *ValidatorMessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorMessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMessagesTransactorSession struct {
	Contract     *ValidatorMessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorMessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMessagesRaw struct {
	Contract *ValidatorMessages // Generic contract binding to access the raw methods on
}

// ValidatorMessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMessagesCallerRaw struct {
	Contract *ValidatorMessagesCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactorRaw struct {
	Contract *ValidatorMessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMessages creates a new instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessages(address common.Address, backend bind.ContractBackend) (*ValidatorMessages, error) {
	contract, err := bindValidatorMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// NewValidatorMessagesCaller creates a new read-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMessagesCaller, error) {
	contract, err := bindValidatorMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesCaller{contract: contract}, nil
}

// NewValidatorMessagesTransactor creates a new write-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMessagesTransactor, error) {
	contract, err := bindValidatorMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesTransactor{contract: contract}, nil
}

// NewValidatorMessagesFilterer creates a new log filterer instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMessagesFilterer, error) {
	contract, err := bindValidatorMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesFilterer{contract: contract}, nil
}

// bindValidatorMessages binds a generic wrapper to an already deployed contract.
func bindValidatorMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.ValidatorMessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transact(opts, method, params...)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackConversionData(opts *bind.CallOpts, conversionData ConversionData) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packConversionData", conversionData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorRegistrationMessage(opts *bind.CallOpts, validationID [32]byte, valid bool) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorRegistrationMessage", validationID, valid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorWeightMessage(opts *bind.CallOpts, validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorWeightMessage", validationID, nonce, weight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackRegisterL1ValidatorMessage(opts *bind.CallOpts, validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packRegisterL1ValidatorMessage", validationPeriod)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetToL1ConversionMessage(opts *bind.CallOpts, conversionID [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetToL1ConversionMessage", conversionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackValidationUptimeMessage(opts *bind.CallOpts, validationID [32]byte, uptime uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packValidationUptimeMessage", validationID, uptime)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorRegistrationMessage(opts *bind.CallOpts, input []byte) ([32]byte, bool, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorRegistrationMessage", input)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorWeightMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorWeightMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackRegisterL1ValidatorMessage(opts *bind.CallOpts, input []byte) (ValidatorMessagesValidationPeriod, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackRegisterL1ValidatorMessage", input)

	if err != nil {
		return *new(ValidatorMessagesValidationPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorMessagesValidationPeriod)).(*ValidatorMessagesValidationPeriod)

	return out0, err

}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetToL1ConversionMessage(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetToL1ConversionMessage", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackValidationUptimeMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackValidationUptimeMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}