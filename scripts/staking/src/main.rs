use ethers:: {
    providers::{Http, Provider},
    types::{Address},
    contract::Contract,
    core::abi::Abi,
};
use std::sync::Arc;

#[tokio::main]
async fn main() -> eyre::Result<()> {
    let provider = Provider::<Http>::try_from("http://rpc-test6.derachain.com/ext/bc/2sSg4ZrhP7vnyhQU2XF5Y4VLegcvZdZVfzpkjZdJrG7cS8gYYD/rpc")?;
    let client = Arc::new(provider);

    let contract_address: Address = "0xC0fFEE1234567890aBCdeF1234567890abcDef34".parse()?;

    let abi: Abi = serde_json::from_str(include_str!("../abis/proxy-admin.abi.json"))?;

    let contract = Contract::new(contract_address, abi, client);

    let owner: Address = contract.method("owner", ())?.call().await?;
    println!("The owner address: {:?}", owner);


    let proxy_address: Address = "0x0Feedc0de0000000000000000000000000000000".parse()?;
    let impl_address: Address = contract.method::<_, Address>("getProxyImplementation", proxy_address)?.call().await?;
    println!("The implementation address of proxy {:?} is {:?}", proxy_address, impl_address);

    Ok(())
}
