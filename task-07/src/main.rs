extern crate regex;
use std::io;
use regex::Regex;

fn main() {
    let re = Regex::new(r"\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$").unwrap();    
    let mut input = String::new();
    io::stdin().read_line(&mut input)
        .ok()
        .expect("Input error");
    println!("{:?}",&input);    
    println!("Match? {}", re.is_match(&input.trim()));
}
