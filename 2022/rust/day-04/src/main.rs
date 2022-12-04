extern crate sscanf;

use std::{fs, io, result};
use std::io::BufRead;

#[derive(Debug)]
#[derive(sscanf::FromScanf)]
#[sscanf(format = "{s}-{e}")]
struct Range {
    s: u32,
    e: u32,
}

impl Range {
    fn contained(&self, other: &Range) -> bool {
        self.s >= other.s && self.e <= other.e
    }
    fn overlaps(&self, other: &Range) -> bool {
        (self.s >= other.s && self.s <= other.e) ||
            (self.e >= other.s && self.e <= other.e)
    }
}

#[derive(Debug)]
#[derive(sscanf::FromScanf)]
#[sscanf(format = "{one},{two}")]
struct ElfPair {
    one: Range,
    two: Range,
}

impl ElfPair {
    fn self_contained(&self) -> bool {
        self.one.contained(&self.two) | self.two.contained(&self.one)
    }
}

fn part_1(filename: &str) -> result::Result<i32, io::Error> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    let mut sum = 0;
    for line in reader.lines() {
        let line = line?;
        let pair: ElfPair = sscanf::sscanf!(line, "{ElfPair}").expect("Invalid line");
        if pair.self_contained() {
            sum += 1;
        }
    }
    Ok(sum)
}

fn part_2(filename: &str) -> result::Result<i32, io::Error> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    let mut sum = 0;
    for line in reader.lines() {
        let line = line?;
        let pair: ElfPair = sscanf::sscanf!(line, "{ElfPair}").expect("Invalid line");
        if pair.one.overlaps(&pair.two) || pair.two.overlaps(&pair.one) {
            sum += 1;
        }
    }
    Ok(sum)
}

fn main() {
    let filename = "input.txt";
    // let filename = "input_test.txt";
    println!("Part 1: {:?}", part_1(filename).unwrap());
    println!("Part 2: {:?}", part_2(filename).unwrap());
}
