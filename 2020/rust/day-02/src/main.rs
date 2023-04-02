use std::{fs, io};
use std::io::BufRead;
use std::str::FromStr;

fn main() {
    use std::time::Instant;
    let now = Instant::now();


    let filename = "input.txt";
    // let filename = "input_test.txt";
    println!("Part 1: {:?}", part_1(filename).unwrap());
    println!("Part 2: {:?}", part_2(filename).unwrap());

    let elapsed = now.elapsed();
    println!("Elapsed: {:.2?}", elapsed);
}

struct PassportPolicy {
    min: usize,
    max: usize,
    char: char,
    password: String,
}

impl PassportPolicy {
    fn is_valid(&self) -> bool {
        let count = self.password.chars().filter(|c| *c == self.char).count();
        count >= self.min && count <= self.max
    }
    fn is_valid_2(&self) -> bool {
        let first = self.password.chars().nth(self.min - 1).unwrap();
        let second = self.password.chars().nth(self.max - 1).unwrap();

        (first == self.char) ^ (second == self.char)
    }
}

impl FromStr for PassportPolicy {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.split_whitespace();
        let mut min_max = parts.next().unwrap().split('-');
        let min = min_max.next().unwrap().parse::<usize>().unwrap();
        let max = min_max.next().unwrap().parse::<usize>().unwrap();
        let char = parts.next().unwrap().chars().next().unwrap();
        let password = parts.next().unwrap().to_string();

        Ok(PassportPolicy {
            min: min,
            max: max,
            char: char,
            password: password,
        })
    }
}

fn part_1(filename: &str) -> Result<usize, &str> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let input: usize = reader
        .lines()
        .map(|line| line.unwrap())
        .map(|line| line.parse::<PassportPolicy>().unwrap())
        .filter(|policy| policy.is_valid()).count();

    Ok(input)
}

fn part_2(filename: &str) -> Result<usize, &str> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let input: usize = reader
        .lines()
        .map(|line| line.unwrap())
        .map(|line| line.parse::<PassportPolicy>().unwrap())
        .filter(|policy| policy.is_valid_2()).count();

    Ok(input)
}
