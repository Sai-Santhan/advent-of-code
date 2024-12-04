fn get_input() -> &'static str {
    "forward 5
down 5
forward 8
up 3
down 8
forward 2"
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn parse_line(line: &str) -> Point {
    let (dir, amount) = line.split_once(' ').expect("must contain a white space");
    // let amount1 = amount.parse::<i32>().expect("amount1 must be a number");
    let amount2 = str::parse::<i32>(amount).expect("amount2 must be a number");

    // if dir == "forward" {
    //     return Point { x: amount1, y: 0 };
    // } else if dir == "up" {
    //     return Point { x: 0, y: -amount1 };
    // }
    // Point { x: 0, y: amount1 }

    if dir == "forward" {
        return Point { x: amount2, y: 0 };
    } else if dir == "up" {
        return Point { x: 0, y: -amount2 };
    }
    Point { x: 0, y: amount2 }
}

fn main() {
    let result = get_input()
        .lines()
        .map(parse_line)
        .fold(Point { x: 0, y: 0 }, |mut acc, point| {
            acc.x += point.x;
            acc.y += point.y;
            acc
        });

    println!("{:?}", result);
}