use std::error::Error;
use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger, TeleporterMessenger};
use ethers::types::{Address, H256, Bytes, U256, U64};
use ethers::utils::parse_ether;

pub async fn handle_send_crosschain_message(messenger: &TeleporterMessenger, 
    destination_blockchain_id: &str,
    destination_address: &str,
    fee_token_address: &str,
    fee_amount: &str, // In ETH, will be converted to wei
    required_gas_limit: u64,
    message: &str,) -> Result<(), Box<dyn Error>> {
    let message_payload = message.as_bytes().to_vec();
    
    let tx_hash = messenger.send_cross_chain_message(
        destination_blockchain_id,
        destination_address,
        fee_token_address,
        fee_amount,
        required_gas_limit,
        message_payload,
    ).await?;
    
    println!("Message sent! Transaction hash: {:?}", tx_hash);

    Ok(())
}