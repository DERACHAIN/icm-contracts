use ethers::{
    contract::{abigen, Contract},
    middleware::SignerMiddleware,
    providers::{Http, Provider},
    signers::{LocalWallet, Signer},
    types::{Address, Bytes, H256, U256},
};
use std::error::Error;
use std::sync::Arc;
use std::str::FromStr;
use ethers::utils::parse_ether;

// Generate type-safe contract bindings from ABI
abigen!(
    TeleporterMessengerContract,
    "abis/teleporter-messenger.abi.json",
);

// Implementation with signing capabilities for sending transactions
#[derive(Debug)]
pub struct TeleporterMessenger {
    contract: TeleporterMessengerContract<SignerMiddleware<Provider<Http>, LocalWallet>>,
}

impl TeleporterMessenger {
    pub fn new(private_key: &str, rpc_url: &str, teleporter_address: &str, eth_chainid: &u64) -> Self {
        // Set up the signer
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();
        let wallet = private_key.parse::<LocalWallet>().unwrap().with_chain_id(*eth_chainid);

        let client = SignerMiddleware::new(provider, wallet);
        // Create the contract instance with signer
        let teleporter_address: Address = teleporter_address.parse().unwrap();

        let contract = TeleporterMessengerContract::new(
            teleporter_address, 
            Arc::new(client.clone()),
        );

        TeleporterMessenger { contract }
    }

    // View methods (don't require signing)
    pub async fn get_blockchain_id(&self) -> Result<H256, Box<dyn Error>> {
        let blockchain_id = self.contract.blockchain_id().call().await?;        
        Ok(H256::from_slice(&blockchain_id))
    }

    pub async fn message_nonce(&self) -> Result<U256, Box<dyn Error>> {
        let nonce = self.contract.message_nonce().call().await?;
        Ok(nonce)
    }

    pub async fn get_next_message_id(&self, destination_blockchain_id: H256) -> Result<H256, Box<dyn Error>> {
        let message_id = self.contract.get_next_message_id(destination_blockchain_id.into()).call().await?;
        Ok(H256::from_slice(&message_id))
    }

    pub async fn calculate_message_id(
        &self,
        source_blockchain_id: H256,
        destination_blockchain_id: H256,
        nonce: U256,
    ) -> Result<H256, Box<dyn Error>> {
        let message_id = self.contract.calculate_message_id(
            source_blockchain_id.into(),
            destination_blockchain_id.into(),
            nonce,
        ).call().await?;
        Ok(H256::from_slice(&message_id))
    }

    pub async fn get_message_hash(&self, message_id: H256) -> Result<H256, Box<dyn Error>> {
        let message_hash = self.contract.get_message_hash(message_id.into()).call().await?;
        Ok(H256::from_slice(&message_hash))
    }

    pub async fn get_fee_info(&self, message_id: H256) -> Result<(Address, U256), Box<dyn Error>> {
        let fee_info = self.contract.get_fee_info(message_id.into()).call().await?;
        Ok(fee_info)
    }

    pub async fn message_received(&self, message_id: H256) -> Result<bool, Box<dyn Error>> {
        let received = self.contract.message_received(message_id.into()).call().await?;
        Ok(received)
    }

    pub async fn get_relayer_reward_address(&self, message_id: H256) -> Result<Address, Box<dyn Error>> {
        let address = self.contract.get_relayer_reward_address(message_id.into()).call().await?;
        Ok(address)
    }

    pub async fn check_relayer_reward_amount(
        &self,
        relayer: Address,
        fee_asset: Address,
    ) -> Result<U256, Box<dyn Error>> {
        let amount = self.contract.check_relayer_reward_amount(relayer, fee_asset).call().await?;
        Ok(amount)
    }

    pub async fn get_receipt_queue_size(&self, source_blockchain_id: H256) -> Result<U256, Box<dyn Error>> {
        let size = self.contract.get_receipt_queue_size(source_blockchain_id.into()).call().await?;
        Ok(size)
    }

    // Transaction methods that require signing
    pub async fn initialize_blockchain_id(&self) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.initialize_blockchain_id();
        let tx = contract_call.send().await?;        
        let receipt = tx.await?;
        if let Some(receipt) = receipt {
            // The return value should be in the logs of the receipt
            // For simplicity, we'll just return the transaction hash
            Ok(receipt.transaction_hash)
        } else {
            Err("Transaction failed".into())
        }
    }

    pub async fn send_cross_chain_message(
        &self,
        destination_blockchain_id: &str,
        destination_address: &str,
        fee_token_address: &str,
        fee_amount: &str, // In ETH, will be converted to wei
        required_gas_limit: u64,
        message: Vec<u8>,
    ) -> Result<H256, Box<dyn Error>> {
        // Parse destination blockchain ID (expected in hex format)
        let destination_blockchain_id = H256::from_str(destination_blockchain_id)?;
        
        // Parse addresses
        let destination_address = Address::from_str(destination_address)?;
        let fee_token_address = Address::from_str(fee_token_address)?;
        
        // Convert ETH amount to wei
        let fee_amount = parse_ether(fee_amount)?;
        
        // Create fee info struct
        let fee_info = TeleporterFeeInfo {
            fee_token_address,
            amount: fee_amount,
        };
        
        // Create message input
        let message_input = TeleporterMessageInput {
            destination_blockchain_id: destination_blockchain_id.into(),
            destination_address,
            fee_info,
            required_gas_limit: U256::from(required_gas_limit),
            allowed_relayer_addresses: vec![], // Empty for any relayer
            message: Bytes::from(message),
        };
        
        // Send the transaction
        println!("Sending cross-chain message...");
        let contract_call = self.contract.send_cross_chain_message(message_input);
        let tx = contract_call.send().await?;        
        
        // Wait for transaction to be mined
        println!("Waiting for transaction to be mined...");
        let receipt = tx.await?;
        
        // Handle the receipt
        if let Some(receipt) = receipt {
            println!("Transaction mined in block: {}", receipt.block_number.unwrap());
            
            Ok(receipt.transaction_hash)
        } else {
            Err("Transaction failed".into())
        }
    }

    pub async fn add_fee_amount(
        &self,
        message_id: H256,
        fee_token_address: Address,
        additional_fee_amount: U256,
    ) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.add_fee_amount(message_id.into(), fee_token_address, additional_fee_amount);
        let tx = contract_call.send().await?;        
        let receipt = tx.await?;
        if let Some(receipt) = receipt {
            Ok(receipt.transaction_hash)
        } else {
            Err("Transaction failed".into())
        }
    }

    pub async fn receive_cross_chain_message(
        &self,
        message_index: u32,
        relayer_reward_address: Address,
    ) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.receive_cross_chain_message(message_index, relayer_reward_address);
        let tx = contract_call.send().await?;
        let receipt = tx.await?;
        if let Some(receipt) = receipt {
            Ok(receipt.transaction_hash)
        } else {
            Err("Transaction failed".into())
        }
    }

    pub async fn redeem_relayer_rewards(
        &self,
        fee_asset: Address,
    ) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.redeem_relayer_rewards(fee_asset);
        let tx = contract_call.send().await?;
        let receipt = tx.await?;
        if let Some(receipt) = receipt {
            Ok(receipt.transaction_hash)
        } else {
            Err("Transaction failed".into())
        }
    }

    // Helper method to create a TeleporterMessageInput struct
    pub fn create_message_input(
        &self,
        destination_blockchain_id: H256,
        destination_address: Address,
        fee_token_address: Address,
        fee_amount: U256,
        required_gas_limit: U256,
        allowed_relayer_addresses: Vec<Address>,
        message: Bytes,
    ) -> TeleporterMessageInput {
        let fee_info = TeleporterFeeInfo {
            fee_token_address,
            amount: fee_amount,
        };

        TeleporterMessageInput {
            destination_blockchain_id: destination_blockchain_id.into(),
            destination_address,
            fee_info,
            required_gas_limit,
            allowed_relayer_addresses,
            message,
        }
    }
}