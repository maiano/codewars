use std::collections::HashMap;

fn duplicate_encode(word: &str) -> String {
    let word = word.to_string().to_lowercase();
    word.chars()
        .map(|c| {
            if word.matches(c).count() == 1 {
                '('
            } else {
                ')'
            }
        })
        .collect()
}

fn count_duplicates(text: &str) -> u32 {
    let mut map = HashMap::new();
    for ch in text.to_lowercase().chars() {
        map.entry(ch)
            .and_modify(|counter| *counter += 1)
            .or_insert(1);
    }
    map.into_values().filter(|x| x > &1).count() as u32
}

fn create_phone_number(numbers: &[u8]) -> String {
    if let [a, b, c, d, e, f, g, h, i, j] = &numbers[..] {
        format!("({}{}{}) {}{}{}-{}{}{}{}", a, b, c, d, e, f, g, h, i, j)
    } else {
        String::new()
    }
}

fn array_diff<T: PartialEq>(a: Vec<T>, b: Vec<T>) -> Vec<T> {
    a.into_iter().filter(|x| !b.contains(x)).collect()
}

fn solution(num: i32) -> i32 {
    (3..num).fold(0, |acc, n| {
        if n % 3 == 0 || n % 5 == 0 {
            acc + n
        } else {
            acc
        }
    })
}

fn find_short(s: &str) -> u32 {
    // s.split_whitespace().map(|x| x.len()).min().unwrap_or(0) as u32
    s.split_whitespace()
        .fold(s.len() as u32, |a, b| a.min(b.len() as u32))
}

fn xo(string: &'static str) -> bool {
    string.chars().fold(0, |a, c| match c {
        'x' | 'X' => a + 1,
        'o' | 'O' => a - 1,
        _ => a,
    }) == 0
}

fn descending_order(mut x: u64) -> u64 {
    let mut digits = Vec::new();
    while x > 0 {
        digits.push(x % 10);
        x = x / 10;
    }
    digits.sort();
    digits.iter().rev().fold(0, |s, n| s * 10 + n)
}

fn high_and_low(numbers: &str) -> String {
    let arr: Vec<i32> = numbers
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();
    let min = arr.iter().min().unwrap();
    let max = arr.iter().max().unwrap();
    format!("{} {}", max, min)
}

fn disemvowel(s: &str) -> String {
    s.chars().filter(|&c| !"aeiouAEIOU".contains(c)).collect()
}

fn get_count(string: &str) -> usize {
    let vowels = vec!['a', 'e', 'i', 'o', 'u'];
    string.chars().filter(|c| vowels.contains(c)).count()
}

fn maps(values: &Vec<i32>) -> Vec<i32> {
    let mut list = Vec::new();
    for &el in values.iter() {
        list.push(el * 2);
    }
    list
}

fn greet(name: &str) -> String {
    format!("Hello, {} how are you doing today?", name)
}

fn even_or_odd(i: i32) -> &'static str {
    match i & 1 {
        0 => "Even",
        _ => "Odd",
    }
}

fn solution(phrase: &str) -> String {
    phrase.chars().rev().collect()
}

fn positive_sum(slice: &[i32]) -> i32 {
    let mut s = 0;

    for i in slice.iter() {
        if i > &0 {
            s += i;
        }
    }
    return s;
}

fn basic_op(operator: char, value1: i32, value2: i32) -> i32 {
    match operator {
        '+' => value1 + value2,
        '-' => value1 - value2,
        '*' => value1 * value2,
        '/' => value1 / value2,
        _ => 0,
    }
}

fn main() {
    println!("Hello, world!");
}
