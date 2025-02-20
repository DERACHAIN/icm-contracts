mod utils;

use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;

use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger};
use ethers::utils::parse_ether;
use ethers::types::{Address, H256, Bytes, U256};

#[derive(Debug)]
pub struct Config {
    pub private_key: String,
    pub rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
    pub warp_address: String,
    pub bootstrap_nodeid: String,
    pub bootstrap_bls_public_key: String,
    pub bootstrap_pop: String,
}

impl Config {
    pub fn new() -> Self {
        Config {
            private_key: env::var("PRIVATE_KEY").unwrap(),
            rpc_url: env::var("RPC_URL").unwrap(),
            proxy_admin_address: env::var("PROXY_ADMIN_ADDRESS").unwrap(),
            proxy_address: env::var("PROXY_ADDRESS").unwrap(),
            warp_address: env::var("WARP_ADDRESS").unwrap(),
            bootstrap_nodeid: env::var("BOOTSTRAP_NODEID").unwrap(),
            bootstrap_bls_public_key: env::var("BOOTSTRAP_BLS_PUBLIC_KEY").unwrap(),
            bootstrap_pop: env::var("BOOTSTRAP_POP").unwrap(),
        }
    }
}

pub async fn run(config: Config) -> Result<(), Box<dyn Error>> {
    // pre-checks the proxy admin and warp messenger
    let proxy_admin = ProxyAdmin::new(&config.rpc_url, &config.proxy_admin_address);
    let owner = proxy_admin.owner().await?;
    let impl_address = proxy_admin
        .get_proxy_implementation(&config.proxy_address)
        .await?;

    println!("The owner address: {:?}", owner);
    println!(
        "The implementation address of proxy {:?} is {:?}",
        config.proxy_address, impl_address
    );

    let warp_messenger = WarpMessenger::new(&config.rpc_url, &config.warp_address);
    let blockchain_id = warp_messenger.get_blockchain_id().await?;
    println!("The blockchain id is {:?}", blockchain_id);

    // pre-checks the validator manager
    let validator_manager =
        ValidatorManager::new(&config.private_key, &config.rpc_url, &config.proxy_address);

    let bootstrap_nodeid = NodeID::new(
        &config.bootstrap_nodeid,
        &config.bootstrap_bls_public_key,
        &config.bootstrap_pop,
    );
    println!("NodeID: {:?}", bootstrap_nodeid.hex_id);

    let validation_hexid = validator_manager
        .get_validation_id(&bootstrap_nodeid.hex_id)
        .await?;
    println!("ValidationID hex: {:#x}", validation_hexid);

    let validation_id = ValidationID::new(&env::var("BOOTSTRAP_VALIDATION_ID").unwrap());
    let hexid_str: String = format!("{}", hex::encode(validation_hexid.as_bytes()));
    assert_eq!(hexid_str, validation_id.hex_id);

    // initialize validator registration
    let new_validator = NodeID::new(
        "NodeID-CBRa26m4Vi974FWM6hHypGhjLF7vVGF5z",
        "0xab3e94f1eaad2a7cd13660101cbfb20b7aa1cade10ae4d6bd0085c757857bf743d77133e76f59285be154d40894c32cd",
        "0x89c581c41d128d9c2e554adbf5eb95cd96725938b597df02762c86e869c2070a44fce5faf19b98163aaf40ac4d04b6e2177b379287c3b4ea726d189db99db615eeedb7d9d1b74d5d628d7f1e0e963716d5929476912c97ee63e0c61f0b580409"
    );

    println!("New NodeID: {:?}", new_validator.hex_id);

    println!(
        "Timestamp of 1 day from now {:?}",
        utils::get_future_timestamp(24 * 3600)
    );

    let expiration = utils::get_future_timestamp(24 * 3600);
    let owner_address = "0xc0Ce63ca7605cb29aA6bcd040715D2A383a9f4aC";
    let delegation_fee_bips = 20;
    let min_stake_duration = 60*60*24*7;
    let stake_amount = parse_ether(20000).unwrap();

    // let tx_hash = validator_manager
    //     .initialize_validator_registration(
    //         &new_validator.hex_id,
    //         &new_validator.bls_public_key,
    //         expiration,
    //         &owner_address,
    //         &owner_address,
    //         delegation_fee_bips,
    //         min_stake_duration,
    //         stake_amount,
    //     )
    //     .await?;
    // println!("InitializeValidatorRegistration TxHash {:?}", tx_hash);

    let validation_hexid = validator_manager
        .get_validation_id(&new_validator.hex_id)
        .await?;
    println!("!!!NEWValidationID hex: {:#x}", validation_hexid);

    let validator = validator_manager.get_validator(validation_hexid).await?;
    println!("Validator: {:?}", validator);

    Ok(())
}

const NODEID_PREFIX: &str = "NodeID-";

#[derive(Debug, Clone)]
pub struct NodeID {
    pub node_id: String,
    pub cb58_id: String,
    pub hex_id: String,
    pub bls_public_key: String,
    pub pop: String, // Proof of possession
}

impl NodeID {
    pub fn new(node_id_str: &str, bls_public_key: &str, pop: &str) -> Self {
        let node_id = ids::node::Id::from_str(node_id_str).unwrap();
        //println!("NodeID: {:?}", node_id.short_id());
        let short_id = node_id.short_id();
        let hex_id = format!("{}", hex::encode(&short_id));
        println!("Hex: 0x{:?}", hex_id);

        NodeID {
            node_id: node_id_str.to_string(),
            cb58_id: short_id.to_string(),
            hex_id,
            bls_public_key: bls_public_key[2..].to_string(),
            pop: pop[2..].to_string(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct ValidationID {
    cb58_id: String,
    hex_id: String,
}

impl ValidationID {
    pub fn new(cb58_id: &str) -> ValidationID {
        let validation_id = ids::Id::from_str(cb58_id).unwrap();
        println!("ValidationID: {:?}", validation_id);
        let hex_id = format!("{}", hex::encode(&validation_id));
        println!("Hex: 0x{:?}", hex_id);

        ValidationID {
            cb58_id: cb58_id.to_string(),
            hex_id,
        }
    }

    pub fn newFromHex(hex_id: &H256) -> ValidationID {
        let bytes = hex::decode(&hex_id).unwrap();
        println!("Bytes: {:?}", bytes);
        let validation_id = ids::Id::from_slice(&bytes);
        println!("ValidationID: {:?}", validation_id);
        let cb58_id = validation_id.to_string();
        println!("CB58: {:?}", cb58_id);

        ValidationID {
            cb58_id,
            hex_id: hex_id.to_string(),
        }
    }
}
