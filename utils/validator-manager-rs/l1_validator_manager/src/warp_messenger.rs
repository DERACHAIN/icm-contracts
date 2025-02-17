use ethers::{
    contract::{abigen, Contract},
    core::abi::Abi,
    middleware::SignerMiddleware,
    providers::{Http, Provider},
    signers::{LocalWallet, Signer},
    types::{Address, Bytes, H256, U256},
};
use std::error::Error;
use std::sync::Arc;

abigen!(
    WarpMessengerContract,
    r"[
    function getBlockchainID() external view returns (bytes32)
    ]"
);

pub struct WarpMessenger {
    client: Arc<Provider<Http>>,
    contract: WarpMessengerContract<Provider<Http>>,
}

impl WarpMessenger {
    pub fn new(rpc_url: &str, contract_address: &str) -> Self {
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();        
        let client = Arc::new(provider);
        let contract_address: Address = contract_address.parse().unwrap();
        let contract = WarpMessengerContract::new(contract_address, client.clone());

        WarpMessenger {
            client,
            contract,
        }
    }

    pub async fn get_blockchain_id(&self) -> Result<Bytes, Box<dyn Error>> {
        let blockchain_id = self.contract.get_blockchain_id().call().await?;
        Ok(blockchain_id.into())
    }
}
