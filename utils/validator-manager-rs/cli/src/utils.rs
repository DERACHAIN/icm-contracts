use chrono::{Duration, Utc};

pub fn get_future_timestamp(seconds_from_now: u64) -> u64 {
    let future = Utc::now() + Duration::seconds(seconds_from_now as i64);
    future.timestamp() as u64
}

// Extract bytes from revert error string
pub fn extract_revert_bytes(error_str: &str) -> Option<Vec<u8>> {
    // Look for common patterns in ethers-rs revert error strings
    if let Some(start_idx) = error_str.find("Revert(Bytes(0x") {
        let start = start_idx + "Revert(Bytes(0x".len();
        if let Some(end_idx) = error_str[start..].find("))") {
            let hex_str = &error_str[start..start + end_idx];
            if let Ok(bytes) = hex::decode(hex_str) {
                return Some(bytes);
            }
        }
    }
    None
}
