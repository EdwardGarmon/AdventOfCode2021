use core::num;
use std::cmp::max;
use std::{clone, fs, vec};

fn main() {
    day4()
}

fn day2() {
    let raw = fs::read_to_string("./day2.input.txt").expect("file not found");
    let mut x: i64 = 0;
    let mut y: i64 = 0;

    for line in raw.split("\n") {
        let data: Vec<&str> = line.split(" ").collect();

        match data[0] {
            "forward" => x += data[1].parse::<i64>().unwrap(),
            "down" => y += data[1].parse::<i64>().unwrap(),
            "up" => y -= data[1].parse::<i64>().unwrap(),
            _ => {
                panic!()
            }
        }
    }

    println!("{}", x * y);
}

fn day2_2() {
    let raw = fs::read_to_string("./day2.input.txt").expect("file not found");
    let mut x: i64 = 0;
    let mut aim = 0;
    let mut y: i64 = 0;

    for line in raw.split("\n") {
        let data: Vec<&str> = line.split(" ").collect();

        match data[0] {
            "forward" => {
                y += data[1].parse::<i64>().unwrap() * aim;
                x += data[1].parse::<i64>().unwrap()
            }
            "down" => aim += data[1].parse::<i64>().unwrap(),
            "up" => aim -= data[1].parse::<i64>().unwrap(),
            _ => {
                panic!()
            }
        }
    }

    println!("{}", x * y);
}

fn day3() {
    let raw = fs::read_to_string("./day3.input.txt").expect("file not found");
    let mut counts = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    let mut l = 0;
    for line in raw.split("\n") {
        l += 1;
        let data = line.chars().map(|s| (s as i32) - 48).enumerate();

        for (i, num) in data {
            counts[i] += num;
        }
    }
    let mut g = 0;
    let mut d = 0;

    for i in 0..counts.len() {
        let n = i32::pow(2, 11 - (i as u32));
        if counts[i] > l / 2 {
            g += n;
        } else {
            d += n;
        }
    }

    println!("{}", g * d);
}

fn day3_2() {
    let raw = fs::read_to_string("./day3.input.txt").expect("file not found");

    let mut l = 0;

    let input: Vec<Vec<i32>> = raw
        .split("\n")
        .map(|s| s.chars().map(|s| (s as i32) - 48).collect())
        .collect();

    let w = input[0].len();

    let mut most: Vec<usize> = vec![];
    let mut least: Vec<usize> = vec![];

    for x in 0..input.len() {
        l += 1;
        most.push(x);
        least.push(x);
    }

    println!("{} {}", most.len(), least.len());

    let mut lastLeast: usize = 0;
    let mut lastMost: usize = 0;

    for i in 0..w {
        for i in &most {
            for m in &input[*i] {
                print!("{}", m)
            }
            println!();
        }

        let m_counts = most.iter().fold(0, |a, item| a + input[*item][i]);

        let l_counts = least.iter().fold(0, |a, item| a + input[*item][i]);

        let c = if m_counts > most.len() as i32 - m_counts {
            1
        } else if m_counts == most.len() as i32 - m_counts {
            1
        } else {
            0
        };

        let u = if l_counts > least.len() as i32 - l_counts {
            0
        } else if l_counts == least.len() as i32 - l_counts {
            0
        } else {
            1
        };

        most = most
            .into_iter()
            .filter(|&x| {
                lastMost = x;
                input[x][i] == c
            })
            .collect();

        let l = least.len();
        least = least
            .into_iter()
            .filter(|&x| {
                lastLeast = x;
                input[x][i] == u
            })
            .collect();

        println!("{} {}", most.len(), least.len());
    }

    let mut l = 0;
    let mut m = 0;

    for i in 0..input[lastMost].len() {
        l += i32::pow(2, (input[lastMost].len() - 1 - i) as u32) * input[lastMost][i];
    }

    for i in 0..input[lastLeast].len() {
        m += i32::pow(2, (input[lastLeast].len() - 1 - i) as u32) * input[lastLeast][i];
    }

    println!("{} , {}", l, m);
    println!("{}", l * m);
}

//1-location of board in board array
//2,3 location of piece in board
struct Location(usize, usize, usize);

fn day4() {}
