# Validator Manager contracts
The deployment guideline for Validator Manager smart contracts.

## Setup

- Determine the SUBNETID_HEX using the ICM relayer tools (command mode)

```sh
$ ./bin/app convertID --source-id=SUBNETID_CB58
```

>   *Note: the `SUBNETID_CB58` can be retrieved using `avalanche blockchain describe` command.

- Create `.env` file from template and fill in necessary credentials.

## Compile

```sh
$ forge build
```

## Upgrade

- Increase the VERSION_NUMBER in `contracts/validator-manager/NativeTokenStakingManager.sol`

```solidity
function initialize(
        PoSValidatorManagerSettings calldata settings
    ) external reinitializer(<VERSION_NUMBER>) {
        __NativeTokenStakingManager_init(settings);
    }
```

- Compile 

```sh
$ forge clean && forge build
```

- Upgrade ValidatorManager implementation

```sh
$ forge script contracts/validator-manager/scripts/UpgradeScript.s.sol \
--rpc-url RPC_URL --broadcast -vvvv
```

>   *Note: Should dry-run before actual deployment by omitting the `--broadcast` argument from the command*

- Confirm the new implementation using `utils/validator-manager-rs` command

```sh
$ cd utils/validator-manager-rs
$ cargo run -p cli -- admin proxy-info
```