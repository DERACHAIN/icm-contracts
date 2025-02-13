use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;

#[derive(Debug)]
pub struct Config {
    pub private_key: String,
    pub rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
}

impl Config {
    pub fn new() -> Self {
        Config {
            private_key: env::var("PRIVATE_KEY").unwrap(),
            rpc_url: env::var("RPC_URL").unwrap(),
            proxy_admin_address: env::var("PROXY_ADMIN_ADDRESS").unwrap(),
            proxy_address: env::var("PROXY_ADDRESS").unwrap(),
        }
    }
}

const NODEID_PREFIX: &str = "NodeID-";

#[derive(Debug, Clone)]
pub struct NodeID {
    node_id: String,
    cb58_id: String,
    hex_id: String,
    bls_public_key: String,
    pop: String, // Proof of possession
}

impl NodeID {
    pub fn new(node_id_str: &str, bls_public_key: &str, pop: &str) -> Self {
        let node_id = ids::node::Id::from_str(node_id_str).unwrap();
        println!("NodeID: {:?}", node_id.short_id());
        let short_id = node_id.short_id();
        let hex_id = format!("0x{}", hex::encode(&short_id));
        println!("Hex: {:?}", hex_id);

        NodeID {
            node_id: node_id_str.to_string(),
            cb58_id: short_id.to_string(),
            hex_id,
            bls_public_key: bls_public_key.to_string(),
            pop: pop.to_string(),
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
        let hex_id = format!("0x{}", hex::encode(&validation_id));
        println!("Hex: {:?}", hex_id);

        ValidationID {
            cb58_id: cb58_id.to_string(),
            hex_id,
        }
    }
}
