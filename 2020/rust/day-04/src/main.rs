// use std::{fs, io};
// use std::io::BufRead;

use std::collections::{HashMap, HashSet};

fn main() {
    use std::time::Instant;
    let now = Instant::now();


//     let filename = "input.txt";
    println!("Part 1: {:?}", part_1().unwrap());
    println!("Part 2: {:?}", part_2().unwrap());


    let elapsed = now.elapsed();
    println!("Elapsed: {:.2?}", elapsed);
}


const REQUIRED_FIELDS: [&str; 7] = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
const EYE_COLORS: [&str; 7] = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

fn part_1() -> Result<usize, &'static str> {
    let count =
        include_str!("../input_test.txt")
            .split("\n\n")
            .map(|line| line
                .split_whitespace()
                .map(|field| { field.split(':').next().unwrap() })
                .collect::<HashSet<_>>())
            .filter(|passport| REQUIRED_FIELDS.iter().all(|field| passport.contains(field)))
            .count();
    Ok(count)
}


fn part_2() -> Result<usize, &'static str> {
    let count =
        include_str!("../input.txt")
            .split("\n\n")
            .map(|line| line
                .split_whitespace()
                .map(|field| { field.split_once(':').unwrap() })
                .collect::<HashMap<_, _>>())
            .filter(|passport| REQUIRED_FIELDS.iter().all(|field| passport.contains_key(field)))
            .filter(|passport| passport.iter().all(|(f, v)| validate_field(f, v)))
            .count();
    Ok(count)
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
fn validate_field(key: &str, value: &str) -> bool {
    match key {
        "byr" => value >= "1920" && value <= "2002",
        "iyr" => value >= "2010" && value <= "2020",
        "eyr" => value.len() == 4 && value >= "2020" && value <= "2030",
        "hgt" => { // Copied from github
            if value.ends_with("cm") && value.len() == 5 {
                value[0..3].parse::<usize>().unwrap().wrapping_sub(150) <= 43
            } else if value.ends_with("in") && value.len() == 4 {
                value[0..2].parse::<usize>().unwrap().wrapping_sub(59) <= 27
            } else {
                false
            }
        }
        "hcl" => {
            if value.len() == 7 && value.starts_with('#') {
                value.chars().skip(1).all(|c| c.is_ascii_hexdigit())
            } else {
                false
            }
        }
        "ecl" => EYE_COLORS.contains(&value),
        "pid" => value.len() == 9 && value.chars().all(|c| c.is_ascii_digit()),
        "cid" => true,
        _ => false,
    }
}

