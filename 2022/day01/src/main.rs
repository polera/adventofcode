use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;
use std::collections::BTreeMap;

fn main() -> std::io::Result<()> {
    let f = File::open("puzzle_input.txt")?;
    let reader = BufReader::new(f);

    let mut entries = BTreeMap::new();
    let mut elf_number = 1;
    let mut calorie_sum = 0;

    for line  in reader.lines() {
        let entry = line.unwrap().parse::<i32>();
        if entry.is_err() {
            entries.insert(calorie_sum, elf_number);
            calorie_sum = 0;
            elf_number += 1;
            continue
        }
        calorie_sum += entry.unwrap();
    }
    // Part 1
    let top_elf : i32 = entries.keys().rev().take(1).sum();
    println!("Part 1: {top_elf}");

    // Part 2
    let top_three_sum : i32 = entries.keys().rev().take(3).sum();
    println!("Part 2: {top_three_sum}");
    Ok(())
}
