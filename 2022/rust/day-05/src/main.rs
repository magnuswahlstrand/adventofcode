use std::{fs, io};
use std::error::Error;
use std::io::BufRead;

#[derive(Debug)]
struct StacksContainer {
    stacks: Vec<Vec<char>>,
}

impl StacksContainer {
    fn new(mut lines: Vec<String>) -> StacksContainer {
        lines.reverse();
        let mut stacks: Vec<Vec<char>> = Vec::new();
        let num_stacks = (&lines[0].len() + 1) / 4;
        for _ in 0..num_stacks {
            stacks.push(Vec::new());
        }
        for line in &lines {
            for (i, c) in line.chars().skip(1).step_by(4).enumerate() {
                if c != ' ' {
                    stacks[i].push(c);
                }
            }
        }

        StacksContainer {
            stacks: stacks,
        }
    }

    fn move_between(&mut self, n: usize, from: usize, to: usize) {
        for _ in 0..n {
            let c = self.stacks[from - 1].pop().unwrap();
            self.stacks[to - 1].push(c);
        }
    }
    fn move_between_v2(&mut self, n: usize, from: usize, to: usize) {
        let range = self.stacks[from - 1].len() - n..;
        let tmp: Vec<char> = self.stacks[from - 1].drain(range).collect();
        self.stacks[to - 1].extend(tmp.iter());
    }


    fn final_words(&self) {
        for c in &self.stacks {
            print!("{}", c.last().unwrap());
        }
        println!("\n")
    }
}


fn part_1(filename: &str) -> Result<&str, Box<dyn Error>> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let mut iterator = reader.lines();

    // Build initial stacks
    let initial: Vec<String> = iterator.by_ref()
        .map(|line| line.unwrap())
        .take_while(|line| line.trim().starts_with("["))
        .collect();
    let mut stacks = StacksContainer::new(initial.clone());

    // Skip empty line
    iterator.by_ref().next();

    for line in iterator.by_ref() {
        let line = line?;
        let (n, from, to) = parse_action(line);
        stacks.move_between(n, from, to)
    }

    stacks.final_words();

    Ok("")
}

fn parse_action(line: String) -> (usize, usize, usize) {
    let action = line.split(" ").collect::<Vec<&str>>();
    if action.len() < 6 {
        panic!("invalid action: {line}");
    }

    let [n, from, to] = [&action[1], &action[3], &action[5]]
        .map(|s| s.parse::<usize>().unwrap());
    (n, from, to)
}

fn part_2(filename: &str) -> Result<&str, Box<dyn Error>> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);

    let mut iterator = reader.lines();

    // Build initial stacks
    let initial: Vec<String> = iterator.by_ref()
        .map(|line| line.unwrap())
        .take_while(|line| line.trim().starts_with("["))
        .collect();
    let mut stacks = StacksContainer::new(initial.clone());

    // Skip empty line
    iterator.by_ref().next();

    for line in iterator.by_ref() {
        let line = line?;
        let (n, from, to) = parse_action(line);
        stacks.move_between_v2(n, from, to)
    }

    stacks.final_words();

    Ok("")
}

fn main() {
    let filename = "input.txt";
    // let filename = "input_test.txt";
    println!("Part 1: {:?}", part_1(filename).unwrap());
    println!("Part 2: {:?}", part_2(filename).unwrap());
}
