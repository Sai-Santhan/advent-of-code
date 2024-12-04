use std::fs::File;
use std::io::{BufRead, BufReader};


fn main() {
    let mut f = BufReader::new(File::open("./input.txt").expect("open failed"));

    let mut buf = Vec::<u8>::new();


    while f.read_until(b'\n', &mut buf).expect("read_until failed") != 0 {
        // this moves the ownership of the read data to s
        // there is no allocation
        let s = String::from_utf8(buf).expect("from_utf8 failed");

        let mut _floor_level : u32 = 0;

        for c in s.chars() {
            
            if c == '(' {
                _floor_level += 1;
            } else if c == ')' {
                _floor_level -= 1;
            }
        }
        // this returns the ownership of the read data to buf
        // there is no allocation
        buf = s.into_bytes();
        buf.clear();
    }
}

// fn main() {
//     if let Ok(lines) = read_lines("./input.txt") {
//         for line in lines {
//             if let Ok(

//             )
//         }
//     }

//     let path = Path::new("input.txt");
//     let display = path.display();
//     let mut s : u32 = 0;

//     let mut file = match File::open(&path) {
//         Err(why) => panic!("couldn't open {}: {}", display, why),
//         Ok(file) => file.,
//     }
// }
