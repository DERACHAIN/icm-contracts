use dotenv::dotenv;
use std::error::Error;

use cli::{Config, NodeID, ValidationID};
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

    let node_id = NodeID::new(
        "NodeID-2mhirVhzPrgDMc1nZVJwXSXg8dKr9YwGh",
        "0xb899613a28e1f55b250d587c9171fa241d11ec490f860f1b4cb8f33e7aa081956ce66c999b48d0b7712911522cc64c68",
        "0xa3b7ca2b66ffc5e878f095d9267adfc531e13d8e6e075ec544eb60594c4286b55974721640d4c5ce6f446e37dd75d2ce02670aa2a8da9c9e11342b8bf41441a7bf3084547a941d0745e5a8190c73d10966cd2f06f281e68e22b0434f0ba14f78",
    );
    println!("NodeID: {:?}", node_id);

    let validation_id = ValidationID::new("AGZiRSc8MRpkaNA5t8a5BLTafzhPxntT5HJyFrL6czD3bKNHo");
    println!("ValidationID: {:?}", validation_id);

    Ok(())
}
