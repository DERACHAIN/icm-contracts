use ethers:: {
    providers::{Http, Provider},
    types::{Address},
    contract::Contract,
    core::abi::Abi,
};
use std::sync::Arc;

use serde::Deserialize;
use config::{Config, Environment, File};
use std::error:: Error;

#[derive(Debug, Deserialize)]
struct Settings {
    rpc_url: String,
    proxy_admin_address: String,
    proxy_address: String,
}

fn load_config() -> Result<Settings, Box<dyn Error>> {
    let config = Config::builder()
                    .add_source(File::with_name("config"))
                    .add_source(Environment::with_prefix("APP").separator("__"))
                    .build()?;

    let settings = config.try_deserialize::<Settings>()?;
    Ok(settings)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let settings = load_config()?;
    println!("RPC URL: {}", settings.rpc_url);

    let provider = Provider::<Http>::try_from(settings.rpc_url)?;
    let client = Arc::new(provider);

    let contract_address: Address = settings.proxy_admin_address.parse()?;

    let abi: Abi = serde_json::from_str(include_str!("../abis/proxy-admin.abi.json"))?;

    let contract = Contract::new(contract_address, abi, client);

    let owner: Address = contract.method("owner", ())?.call().await?;
    println!("The owner address: {:?}", owner);


    let proxy_address: Address = settings.proxy_address.parse()?;
    let impl_address: Address = contract.method::<_, Address>("getProxyImplementation", proxy_address)?.call().await?;
    println!("The implementation address of proxy {:?} is {:?}", proxy_address, impl_address);

    Ok(())
}
