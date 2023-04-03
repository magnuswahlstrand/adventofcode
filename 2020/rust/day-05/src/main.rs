// use std::{fs, io};
// use std::io::BufRead;


fn main() {
    use std::time::Instant;
    let now = Instant::now();


//     let filename = "input.txt";
    println!("Part 1: {:?}", part_1().unwrap());
    // println!("Part 2: {:?}", part_2().unwrap());


    let elapsed = now.elapsed();
    println!("Elapsed: {:.2?}", elapsed);
}


fn part_1() -> Result<i32, &'static str> {
    let count =
        include_str!("../input.txt")
            .lines()
            .map(|line| line_to_int(line))
            .max().unwrap();
    Ok(count)
}

fn line_to_int(line: &str) -> i32 {
    let foo = line.chars().map(|c| {
        match c {
            'F' | 'L' => '0',
            'B' | 'R' => '1',
            _ => panic!("Invalid character"),
        }
    }).collect::<String>();
    let i = i32::from_str_radix(&foo, 2).unwrap();
    i
}

