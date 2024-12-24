# Deployment guide
Step by step guide to deploy NativeTokenStakingManager contracts

## 1. ValidatorMessages
- Command
```sh
$ forge create --private-key <redacted> --rpc-url <rpc-url> \
contracts/validator-manager/ValidatorMessages.sol:ValidatorMessages
```

- Note the deployed address to use in the next step

## 2. NativeTokenStakingManager
- Command
```sh
forge create --private-key <redacted> --rpc-url <rpc-url> \
--libraries contracts/validator-manager/ValidatorMessages.sol:ValidatorMessages:<deployed-address>  \
contracts/validator-manager/NativeTokenStakingManager.sol:NativeTokenStakingManager
```

- Note the deployed address to use in the next step.

## 3. Update the implementation of TransparentProxy with deployed address

## 4. Initialize the NativeTokenStakingManager

- Deploy the RewardCalculator
```sh
$ forge create --constructor-args 100  --private-key <redacted> \
--rpc-url <rpc-url> \
contracts/validator-manager/ExampleRewardCalculator.sol:ExampleRewardCalculator
```

- Note the deployed address to use in the next step

- Call the `initialize` method on manager contract
