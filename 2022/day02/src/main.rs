use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() -> std::io::Result<()> {
    let f = File::open("puzzle_input.txt")?;
    let reader = BufReader::new(f);

    let mut total_score_part1 = 0;
    let mut total_score_part2 = 0;
    for line in reader.lines() {
        let round = line.unwrap();
        let plays: Vec<&str> = round.split(" ").collect();
        total_score_part1 += score_round(plays[0], plays[1]);
        total_score_part2 += score_part_two(plays[0], plays[1]);
    }
    println!("Part 1 Total score: {total_score_part1}");
    println!("Part 2 Total score: {total_score_part2}");
    Ok(())
}

fn score_round(opp_selection: &str, my_selection: &str) -> i32 {
    let rps_values = HashMap::from([
        ("A", 1), // rock
        ("B", 2), // paper
        ("C", 3), // scissors
        ("X", 1), // rock
        ("Y", 2), // paper
        ("Z", 3), // scissors
    ]);

    let opp = rps_values[opp_selection];
    let me = rps_values[my_selection];

    let outcome = match (opp, me) {
        (1, 2) => 6,
        (3, 1) => 6,
        (2, 3) => 6,
        (_, _) => {
            if opp == me {
                3
            } else {
                0
            }
        }
    };

    me + outcome
}

fn score_part_two(opp_selection: &str, my_selection: &str) -> i32 {
    let rules = HashMap::from([
        ("A", "C"), // rock beats scissors
        ("B", "A"), // paper beats rock
        ("C", "B"), // scissors beats paper
    ]);

    let inverted_rules: HashMap<&str, &str> =
        rules.iter().map(|(k, v)| (v.clone(), k.clone())).collect();

    match my_selection {
        "X" => score_round(opp_selection, rules[opp_selection]), // lose
        "Y" => score_round(opp_selection, opp_selection),        // draw
        "Z" => score_round(opp_selection, inverted_rules[opp_selection]), // win
        _ => 0,
    }
}
