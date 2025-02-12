use ethers:: {
    providers::{Http, Provider},
    types::{Address, H256, Bytes},
    contract::Contract,
    core::abi::Abi,
};
use std::sync::Arc;

use serde::Deserialize;
use config::{Config, Environment, File};
use std::error::Error;

use dotenv::dotenv;
use std::env;
use hex;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().ok();

    let rpc_url = env::var("RPC_URL")?;
    let proxy_admin_addres = env::var("PROXY_ADMIN_ADDRESS")?;
    let proxy_address = env::var("PROXY_ADDRESS")?;

    let provider = Provider::<Http>::try_from(&rpc_url)?;
    let client = Arc::new(provider);

    let contract_address: Address = proxy_admin_addres.parse()?;

    let abi: Abi = serde_json::from_str(include_str!("../abis/proxy-admin.abi.json"))?;

    let contract = Contract::new(contract_address, abi, client);

    let owner: Address = contract.method("owner", ())?.call().await?;
    println!("The owner address: {:?}", owner);

    let proxy_address: Address = proxy_address.parse()?;
    let impl_address: Address = contract.method::<_, Address>("getProxyImplementation", proxy_address)?.call().await?;
    println!("The implementation address of proxy {:?} is {:?}", proxy_address, impl_address);

    let provider2 = Provider::<Http>::try_from(&rpc_url)?;
    let client2 = Arc::new(provider2);
    let abiManager: Abi = serde_json::from_str(include_str!("../abis/native-token-staking-manager.abi.json"))?;
    let managerContract = Contract::new(proxy_address, abiManager, client2);
    let nodeIdStr = "5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461";
    let nodeId = hex::decode(&nodeIdStr).unwrap();

    let validationId: H256 = managerContract.method("registeredValidators", Bytes::from(hex::decode(&nodeIdStr).unwrap()))?.call().await?;
    println!("The validation id of node {:?} is {:?}", nodeIdStr, validationId);

    Ok(())
}
