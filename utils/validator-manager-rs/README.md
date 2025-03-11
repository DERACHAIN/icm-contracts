# ValidatorManager interaction scripts
The comprehensive scripts to interact with ValidatorManager and all relevant contracts.

## Setup

- Create `.env` file from template

```sh
$ cp .env.example .env
```

then populate all necessary credentials and secrets.

## Register validator

- TBD

## List validators

- TBD

## Remove validator

- TBD

## Send Crosschain message

- Send to C-chain

```sh
$ cargo run -p cli -- teleporter send-to-c-chain --destination-address=<dest-address> --fee-token-address=<token-address> --fee-amount=<amount> --required-gas-limit=<gas-limit> --message=<msg>
```

- Send to L1

```sh
$ $ cargo run -p cli -- teleporter send-to-l1 --destination-address=<dest-address> --fee-token-address=<token-address> --fee-amount=<amount> --required-gas-limit=<gas-limit> --message=<msg>
```

## Copyright

- TBD