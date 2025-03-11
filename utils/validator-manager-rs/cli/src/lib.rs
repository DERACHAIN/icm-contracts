pub mod utils;
mod admin_cmd;
mod validator_cmd;
mod delegator_cmd;
mod teleporter_cmd;

use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;

use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger, TeleporterMessenger };
use ethers::utils::parse_ether;
use ethers::types::{Address, H256, Bytes, U256, U64};

use clap::{Parser, Subcommand};

#[derive(Debug)]
pub struct Config {
    pub private_key: String,
    pub rpc_url: String,
    pub c_rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
    pub warp_address: String,
    pub bootstrap_nodeid: String,
    pub bootstrap_bls_public_key: String,
    pub bootstrap_pop: String,
    pub teleporter_messenger: String,
    pub blockchain_id: String,
    pub c_blockchain_id: String,
    pub l1_eth_chainid: u64,
    pub c_eth_chainid: u64,
}

impl Config {
    pub fn new() -> Self {
        Config {
            private_key: env::var("PRIVATE_KEY").unwrap(),
            rpc_url: env::var("RPC_URL").unwrap(),
            c_rpc_url: env::var("C_RPC_URL").unwrap(),
            proxy_admin_address: env::var("PROXY_ADMIN_ADDRESS").unwrap(),
            proxy_address: env::var("PROXY_ADDRESS").unwrap(),
            warp_address: env::var("WARP_ADDRESS").unwrap(),
            bootstrap_nodeid: env::var("BOOTSTRAP_NODEID").unwrap(),
            bootstrap_bls_public_key: env::var("BOOTSTRAP_BLS_PUBLIC_KEY").unwrap(),
            bootstrap_pop: env::var("BOOTSTRAP_POP").unwrap(),
            teleporter_messenger: env::var("TELEPORTER_ADDRESS").unwrap(),
            blockchain_id: env::var("BLOCKCHAIN_ID").unwrap(),
            c_blockchain_id: env::var("C_BLOCKCHAIN_ID").unwrap(),
            l1_eth_chainid: env::var("L1_ETH_CHAINID").unwrap().parse::<u64>().unwrap(),
            c_eth_chainid: env::var("C_ETH_CHAINID").unwrap().parse::<u64>().unwrap(),
        }
    }
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

#[derive(Parser)]
#[command(name = "validator-manager")]
#[command(about = "Validator Manager CLI")]
#[command(version, long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Admin commands for proxy and warp messenger
    Admin {
        #[command(subcommand)]
        command: AdminCommands,
    },
    /// Validator management commands
    Validator {
        #[command(subcommand)]
        command: ValidatorCommands,
    },
    /// Delegator management commands
    Delegator {
        #[command(subcommand)]
        command: DelegatorCommands,
    },
    /// Teleporter management commands
    Teleporter {
        #[command(subcommand)]
        command: TeleporterCommands,
    },
}

#[derive(Subcommand)]
enum AdminCommands {
    /// Get proxy admin information
    ProxyInfo,
    /// Get warp messenger information
    WarpInfo,
}

#[derive(Subcommand)]
enum ValidatorCommands {
    /// Get validator information by node ID
    Info {
        #[arg(long)]
        node_id: String,
    },
    /// Register a new validator
    Register {
        #[arg(long)]
        node_id: String,
        
        #[arg(long)]
        bls_public_key: String,
        
        #[arg(long)]
        pop: String,
        
        #[arg(long)]
        owner_address: String,
        
        #[arg(long, default_value = "100")]
        delegation_fee_bips: u16,
        
        #[arg(long, default_value = "3600")] // 1 hour in seconds
        min_stake_duration: u64,
        
        #[arg(long, default_value = "20000")]
        stake_amount: u64,
        
        #[arg(long, default_value = "86400")] // 24 hours in seconds
        expiration: u64,
    },
    /// End validator registration
    Remove {
        #[arg(long)]
        validation_id: String,
        
        #[arg(long, default_value = "false")]
        include_uptime_proof: bool,
        
        #[arg(long, default_value = "0")]
        message_index: u32,
    },
}

#[derive(Subcommand)]
enum DelegatorCommands {
    /// Get delegator information
    Info {
        #[arg(long)]
        delegation_id: String,
    },
    /// Register a new delegator
    Register {
        #[arg(long)]
        validation_id: String,
        
        #[arg(long, default_value = "1000")]
        stake_amount: u64,
    },
    /// End delegator registration
    Remove {
        #[arg(long)]
        delegation_id: String,
        
        #[arg(long, default_value = "false")]
        include_uptime_proof: bool,
        
        #[arg(long, default_value = "0")]
        message_index: u32,
    },
}

#[derive(Subcommand)]
enum TeleporterCommands {
    /// Send cross chain message
    SendToCChain {
        #[arg(long)]
        destination_address: String,

        #[arg(long)]
        fee_token_address: String,

        #[arg(long)]
        fee_amount: String,

        #[arg(long)]
        required_gas_limit: u64,

        #[arg(long)]
        message: String,
    },

    SendToL1 {
        #[arg(long)]
        destination_address: String,

        #[arg(long)]
        fee_token_address: String,

        #[arg(long)]
        fee_amount: String,

        #[arg(long)]
        required_gas_limit: u64,

        #[arg(long)]
        message: String,
    },
}

pub async fn run(cfg: Config) -> Result<(), Box<dyn Error>> {
    let app = Cli::parse();
    let validator_manager = ValidatorManager::new(&cfg.private_key, &cfg.rpc_url, &cfg.proxy_address, cfg.l1_eth_chainid);
    let teleporter_messenger = TeleporterMessenger::new(&cfg.private_key, &cfg.rpc_url, &cfg.teleporter_messenger, &cfg.l1_eth_chainid);
    let c_teleporter_messenger = TeleporterMessenger::new(&cfg.private_key, &cfg.c_rpc_url, &cfg.teleporter_messenger, &cfg.c_eth_chainid);

    match app.command {
        Commands::Admin { command } => match command {
            AdminCommands::ProxyInfo => admin_cmd::handle_admin_proxy_info(&cfg.rpc_url, &cfg.proxy_admin_address, &cfg.proxy_address).await?,
            AdminCommands::WarpInfo => admin_cmd::handle_admin_warp_info(&cfg.rpc_url, &cfg.warp_address).await?,
        },
        Commands::Validator { command } => match command {
            ValidatorCommands::Info { node_id } => {
                validator_cmd::handle_validator_info(&validator_manager, &node_id).await?
            },
            ValidatorCommands::Register { 
                node_id, 
                bls_public_key, 
                pop, 
                owner_address, 
                delegation_fee_bips, 
                min_stake_duration, 
                stake_amount, 
                expiration 
            } => {
                validator_cmd::handle_validator_register(
                    &validator_manager,                    
                    &node_id, 
                    &bls_public_key, 
                    &pop, 
                    &owner_address, 
                    delegation_fee_bips, 
                    min_stake_duration, 
                    stake_amount, 
                    expiration
                ).await?
            },
            ValidatorCommands::Remove { validation_id, include_uptime_proof, message_index } => {
                validator_cmd::handle_validator_remove(&validator_manager, &validation_id, include_uptime_proof, message_index).await?
            },
        },
        Commands::Delegator { command } => match command {
            DelegatorCommands::Info { delegation_id } => {
                delegator_cmd::handle_delegator_info(&validator_manager, &delegation_id).await?
            },
            DelegatorCommands::Register { validation_id, stake_amount } => {
                delegator_cmd::handle_delegator_register(&validator_manager, &validation_id, stake_amount).await?
            },
            DelegatorCommands::Remove { delegation_id, include_uptime_proof, message_index } => {
                delegator_cmd::handle_delegator_remove(&validator_manager, &delegation_id, include_uptime_proof, message_index).await?
            },
        },
        Commands::Teleporter { command } => match command {
            TeleporterCommands::SendToCChain { destination_address, fee_token_address, fee_amount, required_gas_limit, message } => {
                teleporter_cmd::handle_send_crosschain_message(&teleporter_messenger, &cfg.c_blockchain_id, &destination_address, &fee_token_address, &fee_amount, required_gas_limit, &message).await?
            },
            TeleporterCommands::SendToL1 { destination_address, fee_token_address, fee_amount, required_gas_limit, message } => {
                teleporter_cmd::handle_send_crosschain_message(&c_teleporter_messenger, &cfg.blockchain_id, &destination_address, &fee_token_address, &fee_amount, required_gas_limit, &message).await?
            },
        },
    }
    
    Ok(())
}


