use avalanche_types::ids;
use hex;
use std::env;
use std::error::Error;
use std::str::FromStr;
use ethers::types::{Address, H256, Bytes, U256, U64};

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