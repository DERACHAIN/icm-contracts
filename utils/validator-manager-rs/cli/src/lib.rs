mod config;
mod types;
mod utils;
mod admin_cmd;
mod validator_cmd;
mod delegator_cmd;
mod teleporter_cmd;

use std::error::Error;
use clap::{Parser, Subcommand};

use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger, TeleporterMessenger };
pub use config::Config;
pub use types::{NodeID, ValidationID};

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
                
        #[arg(long, default_value = "100")]
        delegation_fee_bips: u16,
        
        #[arg(long, default_value = "3600")] // 1 hour in seconds
        min_stake_duration: u64,
        
        #[arg(long, default_value = "10000")]
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


