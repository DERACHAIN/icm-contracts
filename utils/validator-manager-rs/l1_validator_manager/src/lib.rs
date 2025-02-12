use ethers:: {
    providers::{Http, Provider},
    types::{Address, H256, Bytes},
    contract::Contract,
    core::abi::Abi,
};
use std::sync::Arc;
use std::error::Error;

pub struct ValidatorManager {
    abi: Abi,
    contract: Contract<Provider<Http>>,
}


impl ValidatorManager {
    //! ValidatorManager create
    //! 
    //! This crate is a library for interacting with the validator manager contract on the Avalanche L1 PoS blockchain.
    //! This includes but not limited to getting the validation id of a node.

    /// Create a new ValidatorManager
    /// 
    /// # Arguments
    /// rpc_url - The RPC URL of the Avalanche node
    /// proxy_address - The address of the proxy contract
    /// abi_str - The ABI of the proxy contract
    pub fn new(rpc_url: &str, proxy_address: &str, abi_str: &str) -> Self {
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();
        let client = Arc::new(provider);
        let proxy_address: Address = proxy_address.parse().unwrap();
        let abi: Abi = serde_json::from_str(abi_str).unwrap();
        let contract: Contract<Provider<Http>> = Contract::new(proxy_address, abi.clone(), client.clone());

        ValidatorManager {
            abi,
            contract,
        }
    }

    /// Get the validation id of a node
    /// 
    /// # Arguments    
    /// * `node_id` - The node id of the node
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
        let validation_id: H256 = self.contract.method("registeredValidators", Bytes::from(node_id_hex))?.call().await?;
        Ok(validation_id)
    }
}
