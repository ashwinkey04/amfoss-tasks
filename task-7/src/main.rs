#[macro_use]
extern crate lazy_static;
extern crate regex;

use regex::Regex;

fn main() {
    let re = Regex::new(r"^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$").unwrap();
    let text = "ashwin@amrita.edu";
    println!("Match? {}", re.is_match(text));
}
