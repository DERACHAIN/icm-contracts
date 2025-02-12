use dotenv::dotenv;
use std::error::Error;

use cli::Config;
use l1_validator_manager::ValidatorManager;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().ok();
    let config = Config::new();

    println!("Config: {:?}", config);
    
    let validator_manager = ValidatorManager::new(&config.private_key, &config.rpc_url, &config.proxy_address);

    let node_id = "5d7b4a79d1e63e8b54f698a7a19ebdd36dd23461";
    let validation_id = validator_manager.get_validation_id(&node_id).await?;
    println!("The validation id of node {:?} is {:?}", node_id, validation_id);

    Ok(())
}
