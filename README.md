# ICM Contracts
ICM smart contracts for DERA chain, including:

- ValidatorManager
- Teleporter
- ICTT (unused)

The repository is forked from [Avalanche](https://github.com/ava-labs/icm-contracts)

## Structure

- `contracts/`
  - [`governance/`](./contracts/governance/README.md) includes contracts related to L1 governance.
  - [`ictt/`](./contracts/ictt/README.md) Interchain Token Transfer contracts. Facilitates the transfer of tokens among L1s.
  - [`teleporter/`](./contracts/teleporter/README.md) includes `TeleporterMessenger`, which serves as the interface for most contracts to use ICM. 
    - [`registry/`](./contracts/teleporter/registry/README.md) includes a registry contract for managing different versions of `TeleporterMessenger`.
  - [`validator-manager/`](./contracts/validator-manager/README.md) includes contracts for managing the validator set of an L1.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- [`audits/`](./audits/README.md) includes all audits conducted on contracts in this repository.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `scripts/` includes bash scripts for interacting with TeleporterMessenger in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` performs Solidity and Golang linting.

## Prerequisites

- [Foundry](https://book.getfoundry.sh/) Use `./scripts/install_foundry.sh` to install the Ava Labs [fork](https://github.com/ava-labs/foundry) for building contracts.

## Setup

- Determine the `SUBNETID_HEX` using the ICM relayer tools (command mode)

```sh
$ ./bin/cli convertID --source-id=SUBNETID_CB58
```

prepend the result with `0x` to get the final `SUBNETID_HEX`

>   *Note: the `SUBNETID_CB58` is retrieved using `avalanche blockchain describe` command.

- Create `.env` file from template and fill in necessary information.

- Get all submodules: `git submodule update --init --recursive`

## Compile

```sh
$ forge build
```

## Test

```sh
$ forge test
```

## Deploy

- The ValidatorManager and Teleporter smart contracts is deployed during L1 deployment process, thus initial deployment is not necessary.

## Upgrade

- Increase monotonically the `VERSION_NUMBER` in `contracts/validator-manager/NativeTokenStakingManager.sol`

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

## ValidatorManager CLI tools

See the [ValidatorManager CLI documentation](./utils/validator-manager-rs/README.md) for more details.

## Docs

- [ICM Protocol Overview](./contracts/teleporter/README.md)