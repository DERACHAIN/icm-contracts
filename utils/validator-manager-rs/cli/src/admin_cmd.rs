use std::error::Error;

pub async fn handle_admin_proxy_info(rpc_url: &str, proxy_admin_address: &str, proxy_address: &str) -> Result<(), Box<dyn Error>> {
    let proxy_admin = l1_validator_manager::ProxyAdmin::new(rpc_url, proxy_admin_address);
    let owner = proxy_admin.owner().await?;
    let impl_address = proxy_admin
        .get_proxy_implementation(proxy_address)
        .await?;

    println!("The owner address: {:?}", owner);
    println!(
        "The implementation address of proxy {:?} is {:?}",
        proxy_address, impl_address
    );
    
    Ok(())
}

pub async fn handle_admin_warp_info(rpc_url: &str, warp_address: &str) -> Result<(), Box<dyn Error>> {
    let warp_messenger = l1_validator_manager::WarpMessenger::new(rpc_url, warp_address);
    let blockchain_id = warp_messenger.get_blockchain_id().await?;
    println!("The blockchain id is {:?}", blockchain_id);
    
    Ok(())
}