mod proxy_admin;
pub use proxy_admin::ProxyAdmin;

mod warp_messenger;
pub use warp_messenger::WarpMessenger;

use ethers::{
    providers::{Http, Provider},
    types::{Address, H256, Bytes, U256},
    contract::{Contract, abigen},
    core::abi::Abi,
    middleware::SignerMiddleware,
    signers::{LocalWallet, Signer},    
};
use std::sync::Arc;
use std::error::Error;

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

pub struct ValidatorManager {
    client: SignerMiddleware<Provider<Http>, LocalWallet>,
    contract: NativeTokenStakingManager<SignerMiddleware<Provider<Http>, LocalWallet>>,
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
    pub fn new(private_key: &str, rpc_url: &str, proxy_address: &str) -> Self {
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();
        let wallet = private_key.parse::<LocalWallet>().unwrap().with_chain_id(2025021401 as u64);

        println!("Wallet address: {:?}", wallet.address());

        let client = SignerMiddleware::new(provider, wallet);
        let proxy_address: Address = proxy_address.parse().unwrap();        
        let contract = NativeTokenStakingManager::new(proxy_address, Arc::new(client.clone()));

        ValidatorManager {
            client,
            contract,
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
    pub async fn initialize_validator_registration(&self, node_id: &str, bls_public_key: &str, registration_expiry: u64,
        remaining_balance_owner_address: &str, disable_owner_address: &str, delegation_fee_bips: u16, min_stake_duration: u64, stake_amount: U256) -> Result<H256, Box<dyn Error>> {

        let node_id = Bytes::from(hex::decode(node_id).unwrap());
        let bls_public_key = Bytes::from(hex::decode(bls_public_key).unwrap());

        let validator_registration_input = ValidatorRegistrationInput {
            node_id,
            bls_public_key,
            registration_expiry,
            remaining_balance_owner: PchainOwner {
                threshold: 0,
                // addresses: vec![remaining_balance_owner_address.parse().unwrap()],
                addresses: vec![],
            },
            disable_owner: PchainOwner {
                threshold: 0,
                // addresses: vec![disable_owner_address.parse().unwrap()],
                addresses: vec![],
            },
        };
        let contract_call = self.contract.initialize_validator_registration(validator_registration_input, delegation_fee_bips, min_stake_duration);
        let call_with_value = contract_call.value(stake_amount);
        let pending_tx = call_with_value.send().await?;
        let receipt = pending_tx.await?;

        Ok(receipt.unwrap().transaction_hash)
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
}
