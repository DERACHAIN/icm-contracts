use cb58_rust::{NodeID, ValidationID};

fn main() {
    let nodeId = NodeID::new(&String::from("NodeID-9XHYF8YXzx4zocBPvQ3gv2K3cnteHjmAv"));
    let validationId = ValidationID::new(&String::from(
        "rtwQCxhNMYCTMDsuB5Vr252zZmQTNcoK1mbbJ48YEiXomXhEE",
    ));

    assert_eq!(
        validationId.get_byteStr(),
        "0x714b3623859a0beafd8b604dd558d1180b19b5c87fc6fbecf8437de780305c88",
    );

    cb58_rust::run().unwrap();
}
