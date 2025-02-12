use avalanche_types::ids;
use hex;
use std::error::Error;
use std::str::FromStr;

const NODEID_PREFIX: &str = "NodeID-";

#[derive(Debug)]
pub struct NodeID {
    idStr: String,
    byteStr: String,
}

impl NodeID {
    pub fn new(idStr: &String) -> NodeID {
        let nodeId = ids::node::Id::from_str(idStr).unwrap();
        println!("NodeID: {:?}", nodeId.short_id());
        let shortId = nodeId.short_id();
        let hex_string = format!("0x{}", hex::encode(&shortId));
        println!("Hex: {:?}", hex_string);

        NodeID {
            idStr: idStr.clone(),
            byteStr: hex_string,
        }
    }

    pub fn get_idStr(&self) -> &str {
        &self.idStr
    }

    pub fn get_byteStr(&self) -> &str {
        &self.byteStr
    }
}

#[derive(Debug)]
pub struct ValidationID {
    idStr: String,
    byteStr: String,
}

impl ValidationID {
    pub fn new(idStr: &String) -> ValidationID {
        let validationId = ids::Id::from_str(idStr).unwrap();
        println!("ValidationID: {:?}", validationId);
        let hex_string = format!("0x{}", hex::encode(&validationId));
        println!("Hex: {:?}", hex_string);

        ValidationID {
            idStr: idStr.clone(),
            byteStr: hex_string,
        }
    }

    pub fn get_idStr(&self) -> &str {
        &self.idStr
    }

    pub fn get_byteStr(&self) -> &str {
        &self.byteStr
    }
}

pub fn run() -> Result<(), Box<dyn Error>> {
    println!("Success!");
    Ok(())
}
