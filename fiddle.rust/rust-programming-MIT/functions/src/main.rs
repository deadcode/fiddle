fn main() {
    println!("Hello, world!");

    another_function(5, 6);
}

fn another_function(x: i32, y: i32) {
    println!("From another function x,y is: {},{}", x, y);
}
