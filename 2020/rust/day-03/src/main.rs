use std::{fs, io};
use std::io::BufRead;

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

fn part_1(filename: &str) -> Result<usize, &str> {
    common(filename, 3, 1)
}

fn part_2(filename: &str) -> Result<usize, &str> {
    Ok(
        common(filename, 1, 1).unwrap() *
            common(filename, 3, 1).unwrap() *
            common(filename, 5, 1).unwrap() *
            common(filename, 7, 1).unwrap() *
            common(filename, 1, 2).unwrap()
    )
}

fn common(filename: &str, dx: usize, dy: usize) -> Result<usize, &str> {
    let file = fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(file);
    let count = reader
        .lines()
        .step_by(dy)
        .map(|line| line.unwrap())
        .enumerate()
        .filter(|(i, row)|
            row.chars().nth((i * dx) % row.len()).unwrap() == '#'
        ).count();
    Ok(count)
}
