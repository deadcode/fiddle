fn main() {
    let s1 = String::from("Hello World!");

    let (s1, len) = calc_len(s1);

    println!("s1 = {}, len = {}", s1, len);
}

fn calc_len(str: String) -> (String, usize) {
    let str_len = str.len();

    return (str, str_len);
}
