use chrono::{Duration, Utc};

pub fn get_future_timestamp(seconds_from_now: u64) -> u64 {
    let future = Utc::now() + Duration::seconds(seconds_from_now as i64);
    future.timestamp() as u64
}
