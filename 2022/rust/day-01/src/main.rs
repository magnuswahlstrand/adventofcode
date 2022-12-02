use std::fs;

#[allow(dead_code)]
fn read_input(filename: &str) -> i32 {
    let data = fs::read_to_string(filename).expect("Unable to read file");
    let mut energies = Vec::new();
    for worker in data.trim().split("\n\n") {
        let a: i32 = worker.split("\n").map(|line| (line.parse::<i32>().unwrap())).sum();

        energies.push(a);
    }

    let mut max = -1;
    let mut max_index = -1;
    for i in energies.iter().enumerate() {
        if i.1 > &max {
            max = *i.1;
            max_index = i.0 as i32;
        }
    }
    println!("{}: {}", max, max_index);
    max
}

fn read_input_2(filename: &str) -> Vec<i32> {
    let data = fs::read_to_string(filename).expect("Unable to read file");
    let mut energies = Vec::new();
    for worker in data.trim().split("\n\n") {
        let a: i32 = worker.split("\n").map(|line| (line.parse::<i32>().unwrap())).sum();

        energies.push(a);
    }

    energies.sort_unstable();
    energies.reverse();
    energies
}


fn main() {
    // let filename = "test_input.txt";
    let filename = "input.txt";
    let max = read_input_2(filename);
    println!("{}", max.first().unwrap());
    println!("{}", &max[..3].iter().sum::<i32>());
}
