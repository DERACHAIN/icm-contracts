mod utils;

use dotenv::dotenv;
use std::error::Error;

use cli::{Config, NodeID, ValidationID};
use l1_validator_manager::{ValidatorManager, ProxyAdmin, WarpMessenger};
use ethers::utils::parse_ether;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    dotenv().ok();
    let config = Config::new();
    println!("Config: {:?}", config);

    cli::run(config).await?;

    Ok(())
}
