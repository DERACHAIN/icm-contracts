# ValidatorManager CLI
The CLI tools to interact with ValidatorManager and Teleporter smart contracts, developed in Rust.

## Prerequisites

- [Rust v1.85+](https://www.rust-lang.org/tools/install)

## Structure

```
.
├── cli - the CLI commands
|
├── l1_validator_manager - library package
│   ├── abis - smart contract ABIs, including: ValidatorManager, ProxyAdmin, Teleporter
│   └── src - library
```

## Setup

- Create `.env` file from template

```sh
$ cp .env.example .env
```

then fill in all necessary credentials and secrets.

## Build

```sh
$ cargo build
```

## Admin commands

- Get ValidatorManager proxy information, including owner and implementation address.

```sh
$ cargo run -p cli -- admin proxy-info
```

- Get L1-specific information from WarpMessenger precompile, e.g. L1 blockchain ID

```sh
$ cargo run -p cli -- admin warp-info
```

## Validator commands

- Get validator infor by its NodeID

```sh
$ cargo run -p cli -- validator info --node-id=NODE_ID
```

- Register validator: this command will execute transaction to invoke `initializeValidatorRegistration` method on ValidatorManager smart contract. Please notice that this command only *initialize* the registration process, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- validator register --node-id=NODE_ID --bls-public-key=BLS_PUBLIC_KEY --pop=POP --delegation-fee-bips=FEE_BIPS \
--min-stake-duration=MIN_DURATION_SECONDS --stake-amount=AMOUNT_ETH --expiration=EXPIRATION_SECONDS
```

- Remove validator: this command will execute transaction to invoke `initializeEndValidation` method on ValidatorManager smart contract. Similar to register validator above, this command only *initialize* the removal process, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- validator remove --validation-id=HEX32_VALIDATION_ID --include-uptime-proof=TRUE_FALSE --message-index=0 
```

*Note: the parameters `include_uptime_proof` and `message_index` are optional. If omitted, they will use their default values*

## Delegator commands

- Get delegator info by its delegation ID

```sh
$ cargo run -p cli -- delegator info --delegation-id=HEX32_DELEGATION_ID
```

- Register delegator: this command will execute transaction to invoke `initializeDelegatorRegistration` method on ValidatorManager smart contract. Similar to validators commands, this only *initialize* the registration process, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- delegator register --validation-id=HEX32_VALIDATION_ID --stake-amount=AMOUNT_ETH
```

- Remove delegator: this command will execute transaction to invoke `initializeEndDelegation` method on ValidatorManager smart contract. Similar to commands above, this only *initialize* the removal process, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- delegator remove --delegation-id=HEX32_DELEGATION_ID --include-uptime-proof=TRUE_FALSE --message-index=0 
```

*Note: the parameters `include_uptime_proof` and `message_index` are optional. If omitted, they will use their default values*

## Teleporter commands

- Send message from DERA to C-chain, this command will execute transaction to invoke `sendCrosschainMessage` method on Teleporter smart contract on DERA. Please notice that this command only *initialize* the cross-chain messaging, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- teleporter send-to-c-chain --destination-address=DEST_ADDRESS --fee-token-address=FEE_TOKEN_ADDRESS --fee-amount=FEE_AMOUNT --required-gas-limit=GAS_LIMIT --message=BYTES_MESSAGE
```

- Send message from C-chain to DERA, this command will execute transaction to invoke `sendCrosschainMessage` method on Teleporter smart contract on C-chain. Please notice that this command only *initialize* the cross-chain messaging, which must be completed by ICM relayer.

```sh
$ cargo run -p cli -- teleporter send-to-l1 --destination-address=DEST_ADDRESS --fee-token-address=FEE_TOKEN_ADDRESS --fee-amount=FEE_AMOUNT --required-gas-limit=GAS_LIMIT --message=BYTES_MESSAGE
```

*Note: the parameters `destination_address`, `fee_token_address`, `fee_amount`, `require_gas_limit` are optional. If omitted, they will use their default values*