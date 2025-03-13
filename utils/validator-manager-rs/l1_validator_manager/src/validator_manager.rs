use crate::utils::{extract_revert_bytes};

use ethers::{
    providers::{Http, Provider},
    types::{Address, H256, Bytes, U256, U64},
    contract::{Contract, abigen},
    core::abi::Abi,
    middleware::SignerMiddleware,
    signers::{LocalWallet, Signer},    
};
use std::sync::Arc;
use std::error::Error;
use ethers::abi::AbiDecode;

// generate type-safe contract bindings
abigen!(
    NativeTokenStakingManager,
    "abis/native-token-staking-manager.abi.json",
);

#[derive(Debug)]
pub enum ValidatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Completed,
    Invalidated
}

#[derive(Debug)]
pub enum DelegatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved
}

pub struct ValidatorManager {
    client: SignerMiddleware<Provider<Http>, LocalWallet>,
    contract: NativeTokenStakingManager<SignerMiddleware<Provider<Http>, LocalWallet>>,
    wallet: LocalWallet,
}

impl ValidatorManager {
    //! ValidatorManager create
    //! 
    //! This crate is a library for interacting with the validator manager contract on the Avalanche L1 PoS blockchain.
    //! This includes but not limited to getting the validation id of a node.

    /// Create a new ValidatorManager
    /// 
    /// # Arguments
    /// * private_key - The private key of the account
    /// * rpc_url - The RPC URL of the Avalanche node
    /// * proxy_address - The address of the proxy contract
    /// * abi_str - The ABI of the proxy contract
    pub fn new(private_key: &str, rpc_url: &str, proxy_address: &str, eth_chainid: u64) -> Self {
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();
        let wallet = private_key.parse::<LocalWallet>().unwrap().with_chain_id(eth_chainid);

        println!("Wallet address: {:?}", wallet.address());

        let client = SignerMiddleware::new(provider, wallet.clone());
        let proxy_address: Address = proxy_address.parse().unwrap();
        let contract = NativeTokenStakingManager::new(proxy_address, Arc::new(client.clone()));

        ValidatorManager {
            client,
            contract,
            wallet,
        }
    }

    /// Get the validationId of a node
    /// 
    /// # Arguments    
    /// * `node_id` - The nodeId of the node without prefix `NodeID-`, e.g. `5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461`
    /// 
    /// # Returns
    /// 
    /// * `Result<H256, Box<dyn Error>>` - The validation id of the node as a H256 or an error
    /// 
    /// # Example
    /// ```
    /// use l1_validator_manager::ValidatorManager;
    /// aysnc fn main() {
    /// let validator_manager = ValidatorManager::new(&rpc_url, &proxy_address, &abi_str);
    /// let node_id = "5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461";
    /// let validation_id = validator_manager.get_validation_id(&node_id).await?;
    /// }
    /// ```
    pub async fn get_validation_id(&self, node_id: &str) -> Result<H256, Box<dyn Error>> {
        let node_id_hex = hex::decode(node_id)?;
        println!("NodeID hex: {:?}", node_id_hex);
        let validation_id = self.contract.registered_validators(Bytes::from(node_id_hex)).call().await?;
        Ok(H256::from_slice(&validation_id))
    }

    /// Initialize the validator registration
    pub async fn initialize_validator_registration(&self, node_id: &str, bls_public_key: &str, registration_expiry: u64, delegation_fee_bips: u16, min_stake_duration: u64, stake_amount: U256) -> Result<H256, Box<dyn Error>> {

        let node_id = Bytes::from(hex::decode(node_id)?);
        let bls_public_key = Bytes::from(hex::decode(bls_public_key)?);

        let validator_registration_input = ValidatorRegistrationInput {
            node_id,
            bls_public_key,
            registration_expiry,
            remaining_balance_owner: PchainOwner {
                threshold: 1,
                //addresses: vec![remaining_balance_owner_address.parse()?],
                addresses: vec![self.wallet.address()],                
            },
            disable_owner: PchainOwner {
                threshold: 1,
                //addresses: vec![disable_owner_address.parse()?],
                addresses: vec![self.wallet.address()],                
            },
        };
        let contract_call = self.contract.initialize_validator_registration(validator_registration_input, delegation_fee_bips, min_stake_duration);
        let call_with_value = contract_call.value(stake_amount);
        let pending_tx = match call_with_value.send().await {
            Ok(tx) => tx,
            Err(err) => {
                println!("Tx Error: {:?}", &err.to_string());
                if let Some(err_bytes) = extract_revert_bytes(&err.to_string()) {
                    println!("Revert bytes: {:?}", err_bytes);

                    if let Ok(err_str) = self.decode_contract_error(&err_bytes) {
                        println!("Error: {:?}", err_str);
                        return Err(Box::new(std::io::Error::new(std::io::ErrorKind::Other, err_str)))
                    } else {
                        println!("Error decoding error");
                    }
                } else {
                    println!("No revert bytes");
                }
                return Err(Box::new(err))
            }
        };

        let receipt = match pending_tx.await {
            Ok(Some(receipt)) => receipt,
            Ok(None) => {
                println!("Error: No receipt");
                return Err(Box::new(std::io::Error::new(std::io::ErrorKind::Other, "No receipt")))
            },
            Err(err) => {
                println!("Receipt Error: {:?}", err);
                return Err(Box::new(err))
            }
        };
        
        Ok(receipt.transaction_hash)
    }

    // Decode contract error using the generated error types from abigen
    fn decode_contract_error(&self, bytes: &[u8]) -> Result<String, Box<dyn Error>> {
        // The first 4 bytes are the error selector
        if bytes.len() < 4 {
            return Err("Error data too short".into());
        }
        
        // Try to decode using NativeTokenStakingManagerErrors (generated by abigen)
        let result = match NativeTokenStakingManagerErrors::decode(bytes) {
            Ok(err) => match err {
                NativeTokenStakingManagerErrors::InsufficientBalance(balance) => {
                    format!("Insufficient balance: {}", balance)
                }
                NativeTokenStakingManagerErrors::InvalidBLSKeyLength(length) => {
                    format!("Invalid BLS key length: {}", length)
                }
                NativeTokenStakingManagerErrors::InvalidNodeID(node_id) => {
                    format!("Invalid node ID: {:?}", node_id)
                }
                NativeTokenStakingManagerErrors::NodeAlreadyRegistered(node_id) => {
                    format!("Node already registered: {:?}", node_id)
                }
                NativeTokenStakingManagerErrors::InvalidDelegationFee(fee) => {
                    format!("Invalid delegation fee: {}", fee)
                }
                NativeTokenStakingManagerErrors::InvalidMinStakeDuration(duration) => {
                    format!("Invalid minimum stake duration: {}", duration)
                }
                NativeTokenStakingManagerErrors::InvalidStakeAmount(amount) => {
                    format!("Invalid stake amount: {}", amount)
                }
                NativeTokenStakingManagerErrors::InvalidRegistrationExpiry(expiry) => {
                    format!("Invalid registration expiry: {}", expiry)
                }
                NativeTokenStakingManagerErrors::InvalidPChainOwnerThreshold(threshold) => {
                    format!("Invalid PChain owner threshold: {}", threshold)
                }
                // Add other specific errors from your contract
                _ => format!("Other contract error: {:?}", err),
            },
            Err(_) => {
                // If we can't decode it with our generated types, return the raw selector
                let selector = hex::encode(&bytes[0..4]);
                format!("Unknown error with selector: 0x{}", selector)
            }
        };
        
        Ok(result)
    }

    /// Get validator information by validation ID
    /// 
    /// # Arguments
    /// * `validation_id` - The validation ID as bytes32/H256
    /// 
    /// # Returns
    /// * Result containing the Validator information
    pub async fn get_validator(&self, validation_id: H256) -> Result<Validator, Box<dyn Error>> {
        let validator = self.contract.get_validator(validation_id.into()).call().await?;
        Ok(validator)
    }

    pub async fn get_validator_info(&self, validation_id: H256) -> Result<PoSValidatorInfo, Box<dyn Error>> {
        let validator_info = self.contract.get_validator_info(validation_id.into()).call().await?;
        Ok(validator_info)
    }

    pub async fn initialize_delegator_registration(&self, validation_id: H256, stake_amount: U256) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.initialize_delegator_registration(validation_id.into());
        let call_with_value = contract_call.value(stake_amount);
        let pending_tx = call_with_value.send().await?;
        let receipt = pending_tx.await?;

        Ok(receipt.unwrap().transaction_hash)
    }

    /// Get the delegationID for a given validationID and nonce
    /// 
    /// # Arguments
    /// * `validation_id` - The validation ID as H256
    /// * `nonce` - The nonce as u64
    ///
    /// # Returns
    /// * Result containing the delegation ID as H256
    pub async fn get_delegation_id(&self, validation_id: H256, nonce: u64) -> Result<H256, Box<dyn Error>> {
        let delegation_id = self.contract.get_delegation_id(validation_id.into(), nonce).call().await?;
        Ok(H256::from_slice(&delegation_id))
    }

    /// Get delegator information by delegation ID
    /// 
    /// # Arguments
    /// * `delegation_id` - The delegation ID as H256
    ///
    /// # Returns
    /// * Result containing the Delegator information
    pub async fn get_delegator(&self, delegation_id: H256) -> Result<Delegator, Box<dyn Error>> {
        let delegator = self.contract.get_delegator(delegation_id.into()).call().await?;
        Ok(delegator)
    }

    pub async fn initialize_end_delegation(&self, delegation_id: H256, include_uptime_proof: bool, message_index: u32) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.initialize_end_delegation(delegation_id.into(), include_uptime_proof, message_index);
        let pending_tx = contract_call.send().await?;
        let receipt = pending_tx.await?;

        Ok(receipt.unwrap().transaction_hash)
    }

    pub async fn initialize_end_validation(&self, validation_id: H256, include_uptime_proof: bool, message_index: u32) -> Result<H256, Box<dyn Error>> {
        let contract_call = self.contract.initialize_end_validation(validation_id.into(), include_uptime_proof, message_index);
        let pending_tx = contract_call.send().await?;
        let receipt = pending_tx.await?;

        Ok(receipt.unwrap().transaction_hash)
    }

}