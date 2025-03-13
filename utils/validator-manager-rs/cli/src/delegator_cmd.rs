use ethers::types::{Address, Bytes, H256, U256, U64};
use l1_validator_manager::{ProxyAdmin, ValidatorManager, WarpMessenger};
use std::error::Error;
use ethers::utils::parse_ether;

pub async fn handle_delegator_info(
    manager: &ValidatorManager,
    delegation_id: &str,
) -> Result<(), Box<dyn Error>> {
    
    let delegation_id = delegation_id.parse::<H256>().unwrap();
    
    let delegator = manager.get_delegator(delegation_id).await?;
        
    let validation_id = H256::from_slice(&delegator.validation_id);
    println!("ValidationID: {:?}", validation_id);
    
    let validator = manager.get_validator(validation_id).await?;
    println!("Validator: {:?}", validator);

    println!("Delegator Validator: {:?}", delegator.validation_id);
    println!("Delegator status: {:?}", delegator);
    
    Ok(())
}

pub async fn handle_delegator_register(
    manager: &ValidatorManager,
    validation_id: &str,
    stake_amount: u64,
) -> Result<(), Box<dyn Error>> {
    
    let validation_hexid = validation_id.parse::<H256>().unwrap();
    let stake_amount_wei = parse_ether(stake_amount).unwrap();
    
    let tx_hash = manager.initialize_delegator_registration(validation_hexid, stake_amount_wei)
        .await?;
    
    println!("InitializeDelegatorRegistration TxHash {:?}", tx_hash);
    
    // Get and display the delegation ID
    let nonce: u64 = 1; // This is a placeholder, in practice this might need to be determined
    let delegation_id = manager.get_delegation_id(validation_hexid, nonce).await?;
    println!("DelegationID: {:?}", delegation_id);
    
    Ok(())
}

pub async fn handle_delegator_remove(
    manager: &ValidatorManager,
    delegation_id: &str,
    include_uptime_proof: bool,
    message_index: u32,
) -> Result<(), Box<dyn Error>> {
    let delegation_id = delegation_id.parse::<H256>().unwrap();

    // Get current delegator state
    let delegator = manager.get_delegator(delegation_id).await?;
    println!("Delegator status: {:?}", delegator);
    println!("Delegator Validator: {:?}", delegator.validation_id);

    let validation_id = H256::from_slice(&delegator.validation_id);
    println!("ValidationID: {:?}", validation_id);

    let validator = manager.get_validator(validation_id).await?;
    println!("Validator: {:?}", validator);
    
    let tx_hash = manager.initialize_end_delegation(delegation_id, include_uptime_proof, message_index)
        .await?;
    println!("EndDelegation TxHash {:?}", tx_hash);
    
    Ok(())
}