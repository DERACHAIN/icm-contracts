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

```sh
$ forge script contracts/validator-manager/scripts/UpgradeScript.s.sol \
--rpc-url RPC_URL --broadcast -vvvv
```

>   *Note: Should dry-run before actual deployment by omitting the `--broadcast` argument from the command*