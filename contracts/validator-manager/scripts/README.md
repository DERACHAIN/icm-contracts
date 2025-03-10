# Foundry script to manage validator manager contracts

## Upgrade

```sh
$ forge script contracts/validator-manager/scripts/UpgradeScript.s.sol \
--rpc-url <rpc-endpoint> --broadcast -vvvv
```

*Note: the `ValidatorManager` contract must increase the number in `initialize` method*

```solidity
function initialize(
        PoSValidatorManagerSettings calldata settings
    ) external reinitializer(<NUMBER>) {
        __NativeTokenStakingManager_init(settings);
    }
```