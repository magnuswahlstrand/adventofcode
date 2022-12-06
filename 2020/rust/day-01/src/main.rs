use std::{fs, io};
use std::io::BufRead;

use itertools::Itertools;

fn main() {
    let filename = "input.txt";
    // let filename = "input_test.txt";
    println!("Part 1: {:?}", part_1(filename).unwrap());
    println!("Part 2: {:?}", part_2(filename).unwrap());
}

fn part_1(filename: &str) -> Result<u32, &str> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let input: Vec<u32> = reader
        .lines()
        .map(|line| line.unwrap().parse::<u32>().unwrap()).collect();


    for v in input.iter().combinations(2) {
        if v[0] + v[1] == 2020 {
            return Ok(v[0] * v[1]);
        }
    }

    Err("not found")
}

fn part_2(filename: &str) -> Result<usize, &str> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let input: Vec<usize> = reader
        .lines()
        .map(|line| line.unwrap().parse::<usize>().unwrap()).collect();


    for v in input.iter().combinations(3) {
        if v[0] + v[1] + v[2] == 2020 {
            return Ok(v[0] * v[1] * v[2]);
        }
    }


    Err("not found")
}
