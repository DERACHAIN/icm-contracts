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

    pub async fn get_validation_id(&self, rpc_url: String, node_id: &str) -> Result<H256, Box<dyn Error>> {
        let node_id_hex = hex::decode(node_id)?;
        let validation_id: H256 = self.contract.method("registeredValidators", Bytes::from(node_id_hex))?.call().await?;
        Ok(validation_id)
    }
}
