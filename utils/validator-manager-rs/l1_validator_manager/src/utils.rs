use regex::Regex;

// Extract bytes from revert error string
pub fn extract_revert_bytes(error_str: &str) -> Option<Vec<u8>> {
    let revert_re = Regex::new(r"0x([a-zA-Z0-9]+)").unwrap();
    if let Some(caps) = revert_re.captures(error_str) {
        if let Some(hex_str) = caps.get(1) {
            if let Ok(bytes) = hex::decode(hex_str.as_str()) {
                return Some(bytes);
            }
        }        
    }
    None
}