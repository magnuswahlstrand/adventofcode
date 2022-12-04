extern crate core;

use std::fs;
use std::io;
use std::io::prelude::*;

fn part_1(filename: &str) -> io::Result<u32> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);

    let mut sum = 0;
    for line in reader.lines() {
        let line = line?;
        let item = common_item(line);
        let score = calc_score(item);
        sum += score;
    }
    Ok(sum)
}

fn calc_score(item: char) -> u32 {
    match item {
        'A'..='Z' => item as u32 - 65 + 27,
        'a'..='z' => item as u32 - 97 + 1,
        _ => panic!("Invalid item: {:?}", item)
    }
}

fn part_2(filename: &str) -> io::Result<u32> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);

    let mut sum = 0;
    for group in reader.lines().collect::<Vec<_>>().chunks_exact(3) {
        let one = group[0].as_ref().unwrap();
        let two = group[1].as_ref().unwrap();
        let three = group[2].as_ref().unwrap();
        let item = common_item_2(one, two, three);
        let score = calc_score(item);
        sum += score;
    }
    Ok(sum)
}

fn common_item(line: String) -> char {
    let first = &line[0..line.len() / 2];
    let rest = &line[line.len() / 2..];
    for c in first.chars() {
        if rest.contains(c) {
            return c;
        }
    }
    panic!("Not found");
}

fn common_item_2(one: &str, two: &str, three: &str) -> char {
    for c in one.chars() {
        if two.contains(c) && three.contains(c) {
            return c;
        }
    }
    panic!("Not found");
}

fn main() {
    let filename = "input.txt";
    // let filename = "input_test.txt";
    let part_1 = part_1(filename).unwrap();
    let part_2 = part_2(filename).unwrap();
    println!("Part 1: {:?}", part_1);
    println!("Part 2: {:?}", part_2);
}
