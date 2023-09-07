fn main() {
    let a = [10, 20, 30, 40, 50, 60];

    for elem in a.iter() {
        println!("the value is : {}", elem);
    }

    for num in (1..4).rev() {
        println!("Countdown: {}", num);
    }
}
