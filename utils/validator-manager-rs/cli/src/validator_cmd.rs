use crate::utils;
use crate::NodeID;

use std::error::Error;
use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger};
use ethers::types::{Address, H256, Bytes, U256, U64};
use ethers::utils::parse_ether;

pub async fn handle_validator_info(manager: &ValidatorManager, node_id: &str) -> Result<(), Box<dyn Error>> {
    let node_id = NodeID::new(node_id, "0x", "0x");
    println!("Hex NodeID: {:?}", node_id.hex_id);

    let validation_id = manager.get_validation_id(&node_id.hex_id).await?;
    println!("Validation ID: {:?}", validation_id);

    let validator = manager.get_validator(validation_id).await?;
    println!("Validator: {:?}", validator);

    let validator_info = manager.get_validator_info(validation_id).await?;
    println!("ValidatorInfo: {:?}", validator_info);

    Ok(())
}

pub async fn handle_validator_register(
    manager: &ValidatorManager, 
    node_id: &str,
    bls_public_key: &str,
    pop: &str,
    delegation_fee_bips: u16,
    min_stake_duration: u64,
    stake_amount: u64,
    expiration_seconds: u64,
) -> Result<(), Box<dyn Error>> {

    let new_validator = NodeID::new(node_id, bls_public_key, pop);
    println!("New NodeID: {:?}", new_validator.hex_id);
    
    let expiration = utils::get_future_timestamp(expiration_seconds);
    println!("Expiration timestamp: {:?}", expiration);
    
    let stake_amount_wei = parse_ether(stake_amount).unwrap();
    
    let tx_hash = manager.initialize_validator_registration(
            &new_validator.hex_id,
            &new_validator.bls_public_key,
            expiration,
            delegation_fee_bips,
            min_stake_duration,
            stake_amount_wei,
        )
        .await?;
        
    println!("InitializeValidatorRegistration TxHash {:?}", tx_hash);
    
    // Get and display the validation ID
    let validation_hexid = manager
        .get_validation_id(&new_validator.hex_id)
        .await?;
    println!("ValidationID hex: {:#x}", validation_hexid);

    let validator_info = manager.get_validator_info(validation_hexid).await?;
        println!("ValidatorInfo: {:?}", validator_info);
    
    Ok(())
}

pub async fn handle_validator_remove(
    manager: &ValidatorManager,
    validation_id: &str,
    include_uptime_proof: bool,
    message_index: u32,
) -> Result<(), Box<dyn Error>> {
    let validation_hexid = validation_id.parse::<H256>().unwrap();

    let validator = manager.get_validator(validation_hexid).await?;
    println!("Validator: {:?}", validator);
    
    let tx_hash = manager.initialize_end_validation(validation_hexid, include_uptime_proof, message_index)
        .await?;
    println!("EndValidatorRegistration TxHash {:?}", tx_hash);
    
    Ok(())
}