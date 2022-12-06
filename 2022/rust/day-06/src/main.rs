use std::{fs, io};
use std::collections::HashSet;
use std::io::BufRead;

fn main() {
    let filename = "input.txt";
    // let filename = "input_test.txt";
    println!("Part 1: {:?}", part_1_and_2(filename, 4).unwrap());
    println!("Part 2: {:?}", part_1_and_2(filename, 14).unwrap());
}

fn part_1_and_2(filename: &str, window_size: usize) -> Result<usize, &str> {
    let file = fs::File::open(filename).unwrap();
    let mut reader = io::BufReader::new(file);

    let mut s = String::new();
    reader.read_line(&mut s).unwrap();

    for (i, c) in s.trim().chars().collect::<Vec<char>>().windows(window_size).enumerate() {
        if HashSet::<char>::from_iter(c.iter().cloned()).len() == window_size {
            return Ok(i+window_size);
        }
    }

    Err("not found")
}
