extern crate simple_error;
extern crate sscanf;

use std::{fs, io};
use std::error::Error;
use std::io::BufRead;
use std::str::FromStr;

#[derive(Debug)]
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
struct ElfPair {
    one: Range,
    two: Range,
}

impl ElfPair {
    fn self_contained(&self) -> bool {
        self.one.contained(&self.two) | self.two.contained(&self.one)
    }
    fn overlaps(&self) -> bool {
        self.one.overlaps(&self.two) | self.two.overlaps(&self.one)
    }
}

impl FromStr for ElfPair {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        // let (elf_a, elf_b) = s.split_once(",").unwrap();

        let vals = s.split([',','-']).map(|x| x.parse::<u32>().unwrap()).collect::<Vec<u32>>();
        if vals.len() != 4 {
            return Err(Box::new(simple_error::SimpleError::new("Invalid input")));
        }

        Ok(ElfPair {
            one: Range {
                s: vals[0],
                e: vals[1],
            },
            two: Range {
                s: vals[2],
                e: vals[3],
            },
        })
    }
}
/*
fn part_1(filename: &str) -> Result<i32, io::Error> {
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

fn part_2(filename: &str) -> Result<i32, io::Error> {
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
*/

fn part_1_2(filename: &str) -> Result<i32, Box<dyn Error>> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    let sum: i32 = reader.lines()
        .filter_map(|line_result| line_result.ok()) // #1
        .filter_map(|line| line.parse().ok())
        .map(|pair: ElfPair| if pair.self_contained() { 1 } else { 0 })
        .sum();

    Ok(sum)
}

fn part_2_2(filename: &str) -> Result<i32, Box<dyn Error>> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    let sum: i32 = reader.lines()
        .filter_map(|line_result| line_result.ok()) // #1
        .filter_map(|line| line.parse().ok())
        .map(|pair: ElfPair| if pair.overlaps() { 1 } else { 0 })
        .sum();

    Ok(sum)
}

fn main() {
    let filename = "input.txt";
    println!("Part 1: {:?}", part_1_2(filename).unwrap());
    println!("Part 2: {:?}", part_2_2(filename).unwrap());
}
