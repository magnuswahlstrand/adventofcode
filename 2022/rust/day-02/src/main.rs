use std::borrow::Borrow;
use std::fs;

// A == X Rock
// B == Y Paper
// C == Z Scissor

#[derive(Debug)]
enum Move {
    Rock,
    Paper,
    Scissor,
}

fn lookup_move(input: &str) -> Move {
    match input {
        "A" | "X" => Move::Rock,
        "B" | "Y" => Move::Paper,
        "C" | "Z" => Move::Scissor,
        _ => panic!("this is a terrible mistake!"),
    }
}

fn lookup_required_move(desired_outcome: &str, them: &Move) -> Move {
    match (desired_outcome, them) {
        // Lose
        ("X", Move::Rock) => Move::Scissor,
        ("X", Move::Paper) => Move::Rock,
        ("X", Move::Scissor) => Move::Paper,

        // Draw
        ("Y", Move::Rock) => Move::Rock,
        ("Y", Move::Paper) => Move::Paper,
        ("Y", Move::Scissor) => Move::Scissor,

        // Win
        ("Z", Move::Rock) => Move::Paper,
        ("Z", Move::Paper) => Move::Scissor,
        ("Z", Move::Scissor) => Move::Rock,

        _ => panic!("this is a terrible mistake version two!"),
    }
}

fn part_1(filename: &str) -> Option<i32> {
    let data = fs::read_to_string(filename).expect("Unable to read file");

    let mut total_score = 0;
    for mut round in data.trim().split("\n").map(|x| x.split(" ")) {
        let them = lookup_move(round.next()?);
        let you = lookup_move(round.next()?);
        let round_score = calculate_round_score(them, &you);

        let move_score = match you {
            Move::Rock => 1,
            Move::Paper => 2,
            Move::Scissor => 3,
        };

        total_score += round_score + move_score;
        // println!("Round score: {}", round_score + move_score);
    }
    return Some(total_score)
}

fn part_2(filename: &str) -> Option<i32> {
    let data = fs::read_to_string(filename).expect("Unable to read file");

    let mut total_score = 0;
    for mut round in data.trim().split("\n").map(|x| x.split(" ")) {
        let them = lookup_move(round.next()?);
        let you = lookup_required_move(round.next()?, them.borrow().clone());
        let round_score = calculate_round_score(them, &you);

        let move_score = match you {
            Move::Rock => 1,
            Move::Paper => 2,
            Move::Scissor => 3,
        };

        total_score += round_score + move_score;
        // println!("Round score: {} {}", move_score, round_score);
    }
    return Some(total_score)
}

fn calculate_round_score(them: Move, you: &Move) -> i32 {
    match (you.borrow().clone(), them.borrow().clone()) {
        // Win
        (Move::Rock, Move::Scissor) => 6,
        (Move::Paper, Move::Rock) => 6,
        (Move::Scissor, Move::Paper) => 6,

        // Lose
        (Move::Rock, Move::Paper) => 0,
        (Move::Paper, Move::Scissor) => 0,
        (Move::Scissor, Move::Rock) => 0,
        // Draw
        _ => 3,
    }
}


fn main() {
    // let filename = "test_input.txt";
    let filename = "input.txt";
    let part_1 = part_1(filename).unwrap();
    let part_2 = part_2(filename).unwrap();
    println!("Total score part 1: {}", part_1);
    println!("Total score part 2: {}", part_2);
    // println!("{}", &max[..3].iter().sum::<i32>());
}
