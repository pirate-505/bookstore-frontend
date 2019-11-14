use std::{env::args, process::exit};

struct Book {
    currency: char,
    value: u64,
}

impl Book {
    fn parse(price: &str) -> Self {
        Book {
            value: price[..price.len() - 1].parse().unwrap_or_else(|_| {
                println!("Invalid price format. Aborting...");
                exit(1)
            }),
            currency: price.chars().last().unwrap(),
        }
    }
}

fn main() {
  println!("{:?}", args().collect::<Vec<String>>());
    let prices: Vec<String> = args().skip(1).collect();

    if prices.len() == 0 {
        println!(
            "usage: {0} <cost1>[currency] [cost2 cost3 ...]\nExample: {0} 9$ 9$",
            args().nth(0).unwrap()
        );
        exit(1);
    }

    let mut books = prices.iter().map(|p| Book::parse(p));

    let total = books.clone().fold(0, |result, book| result + book.value);

    println!("Total: {}{}", total, books.nth(0).unwrap().currency)
}
