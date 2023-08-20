fn main() {
    let s1 = String::from("Hello World!");

    let len = calc_len(&s1);

    println!("s1 = {}, len = {}", s1, len);
}

fn calc_len(s2: &String) -> usize {
    return s2.len();
}
