use std::env;

#[derive(Debug)]
pub struct Config {
    pub rpc_url: String,
    pub proxy_admin_address: String,
    pub proxy_address: String,
}

impl Config {
    pub fn new() -> Self {
        Config {
            rpc_url: env::var("RPC_URL").unwrap(),
            proxy_admin_address: env::var("PROXY_ADMIN_ADDRESS").unwrap(),
            proxy_address: env::var("PROXY_ADDRESS").unwrap(),
        }
    }
}
