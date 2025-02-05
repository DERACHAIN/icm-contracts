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

use dotenv::dotenv;
use std::env;

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
    dotenv().ok();

    let rpc_url = env::var("RPC_URL")?;
    let proxy_admin_addres = env::var("PROXY_ADMIN_ADDRESS")?;
    let proxy_address = env::var("PROXY_ADDRESS")?;

    let settings = load_config()?;

    let provider = Provider::<Http>::try_from(rpc_url)?;
    let client = Arc::new(provider);

    let contract_address: Address = proxy_admin_addres.parse()?;

    let abi: Abi = serde_json::from_str(include_str!("../abis/proxy-admin.abi.json"))?;

    let contract = Contract::new(contract_address, abi, client);

    let owner: Address = contract.method("owner", ())?.call().await?;
    println!("The owner address: {:?}", owner);

    let proxy_address: Address = proxy_address.parse()?;
    let impl_address: Address = contract.method::<_, Address>("getProxyImplementation", proxy_address)?.call().await?;
    println!("The implementation address of proxy {:?} is {:?}", proxy_address, impl_address);

    Ok(())
}
