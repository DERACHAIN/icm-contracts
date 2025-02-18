mod utils;

use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;

use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger};
use ethers::utils::parse_ether;

#[derive(Debug)]
pub struct Config {
    pub private_key: String,
    pub rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
    pub warp_address: String,
}

impl Config {
    pub fn new() -> Self {
        Config {
            private_key: env::var("PRIVATE_KEY").unwrap(),
            rpc_url: env::var("RPC_URL").unwrap(),
            proxy_admin_address: env::var("PROXY_ADMIN_ADDRESS").unwrap(),
            proxy_address: env::var("PROXY_ADDRESS").unwrap(),
            warp_address: env::var("WARP_ADDRESS").unwrap(),
        }
    }
}

pub async fn run(config: Config) -> Result<(), Box<dyn Error>> {
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

    let validator_manager =
        ValidatorManager::new(&config.private_key, &config.rpc_url, &config.proxy_address);

    let node_id = "5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461";
    let validation_id = validator_manager.get_validation_id(&node_id).await?;
    println!(
        "The validation id of node {:?} is {:?}",
        node_id, validation_id
    );

    let node_id = NodeID::new(
        "NodeID-2mhirVhzPrgDMc1nZVJwXSXg8dKr9YwGh",
        "0xb899613a28e1f55b250d587c9171fa241d11ec490f860f1b4cb8f33e7aa081956ce66c999b48d0b7712911522cc64c68",
        "0xa3b7ca2b66ffc5e878f095d9267adfc531e13d8e6e075ec544eb60594c4286b55974721640d4c5ce6f446e37dd75d2ce02670aa2a8da9c9e11342b8bf41441a7bf3084547a941d0745e5a8190c73d10966cd2f06f281e68e22b0434f0ba14f78",
    );
    println!("NodeID: {:?}", node_id);

    let validation_id = ValidationID::new("AGZiRSc8MRpkaNA5t8a5BLTafzhPxntT5HJyFrL6czD3bKNHo");
    println!("ValidationID: {:?}", validation_id);

    println!(
        "Timestamp of 1 day from now {:?}",
        utils::get_future_timestamp(24 * 3600)
    );

    let expiration = utils::get_future_timestamp(24 * 3600);
    let owner_address = "0xc0Ce63ca7605cb29aA6bcd040715D2A383a9f4aC";
    let delegation_fee_bips = 10;
    let min_stake_duration = 3600;
    let stake_amount = parse_ether(100).unwrap();

    println!("Warp messager address {:?}", config.warp_address);
    let warp_messenger = WarpMessenger::new(&config.rpc_url, &config.warp_address);
    let blockchain_id = warp_messenger.get_blockchain_id().await?;
    println!("The blockchain id is {:?}", blockchain_id);

    let tx_hash = validator_manager
        .initialize_validator_registration(
            &node_id.hex_id,
            &node_id.bls_public_key,
            expiration,
            &owner_address,
            &owner_address,
            delegation_fee_bips,
            min_stake_duration,
            stake_amount,
        )
        .await?;
    println!("InitializeValidatorRegistration TxHash {:?}", tx_hash);

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
        println!("NodeID: {:?}", node_id.short_id());
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
}
