use dotenv::dotenv;
use std::env;
use std::error::Error;

use ethers::{
    contract::Contract,
    core::abi::Abi,
    providers::{Http, Provider},
    types::{Address, Bytes, H256, U256, U128},
};
use std::sync::Arc;

fn convert_price(price: U256) -> Result<U256, Box<dyn Error>> {
    let base_128 = U256::from(2).pow(U256::from(128));
    let scaled_price = price.checked_mul(U256::exp10(18)).ok_or("Overflow")?;
    let value = scaled_price.checked_div(base_128).ok_or("Overflow")?;
    Ok(value)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().ok();

    let rpc_url = env::var("RPC_URL")?;
    let pair_address = env::var("PAIR_ADDRESS")?;
    let token_address = env::var("TOKEN_ADDRESS")?;

    let provider = Provider::<Http>::try_from(&rpc_url)?;
    let client = Arc::new(provider);

    let pair_address: Address = pair_address.parse()?;
    let pair_abi: Abi = serde_json::from_str(include_str!("../abis/lbpair.abi.json"))?;    
    let lb_contract: Contract<Provider<Http>> = Contract::new(pair_address, pair_abi, client.clone());

    let activeId = lb_contract.method::<_, u32>("getActiveId", ())?.call().await?;
    println!("The active id is {:?}", activeId);

    let price: U256 = lb_contract.method::<u32, U256>("getPriceFromId", activeId)?.call().await?;
    println!("The price of active id {:?} is {:?}", activeId, price);

    let converted_price = convert_price(price).unwrap();
    println!("The converted price of active id {activeId:?} is {converted_price}");

    Ok(())
}
