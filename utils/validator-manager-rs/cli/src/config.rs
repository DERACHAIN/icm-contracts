use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;

#[derive(Debug)]
pub struct Config {
    pub private_key: String,
    pub rpc_url: String,
    pub c_rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
    pub warp_address: String,    
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
            teleporter_messenger: env::var("TELEPORTER_ADDRESS").unwrap(),
            blockchain_id: env::var("BLOCKCHAIN_ID").unwrap(),
            c_blockchain_id: env::var("C_BLOCKCHAIN_ID").unwrap(),
            l1_eth_chainid: env::var("L1_ETH_CHAINID").unwrap().parse::<u64>().unwrap(),
            c_eth_chainid: env::var("C_ETH_CHAINID").unwrap().parse::<u64>().unwrap(),
        }
    }
}