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

// generate type-safe contract bindings
abigen!(ProxyAdminContract, "abis/proxy-admin.abi.json",);

pub struct ProxyAdmin {
    client: Arc<Provider<Http>>,
    contract: ProxyAdminContract<Provider<Http>>,
}

impl ProxyAdmin {
    pub fn new(rpc_url: &str, proxy_admin_address: &str) -> Self {
        let provider = Provider::<Http>::try_from(rpc_url).unwrap();
        let client = Arc::new(provider);
        let proxy_admin_address: Address = proxy_admin_address.parse().unwrap();
        let contract = ProxyAdminContract::new(proxy_admin_address, client.clone());

        ProxyAdmin { client, contract }
    }

    pub async fn owner(&self) -> Result<Address, Box<dyn Error>> {
        let owner = self.contract.owner().call().await?;
        Ok(owner)
    }

    pub async fn get_proxy_implementation(&self, proxy_address: &str) -> Result<Address, Box<dyn Error>> {
        let proxy_address: Address = proxy_address.parse().unwrap();
        let impl_address = self.contract.get_proxy_implementation(proxy_address).call().await?;
        Ok(impl_address)
    }
}
