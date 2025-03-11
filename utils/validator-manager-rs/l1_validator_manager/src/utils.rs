// Extract bytes from revert error string
pub fn extract_revert_bytes(error_str: &str) -> Option<Vec<u8>> {
    // Look for common patterns in ethers-rs revert error strings
    if let Some(start_idx) = error_str.find("0x") {
        let start = start_idx + "0x".len();
        let hex_str = &error_str[start..];
        if let Ok(bytes) = hex::decode(hex_str) {
            return Some(bytes);
        }
    }
    None
}