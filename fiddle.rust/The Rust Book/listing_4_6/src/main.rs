fn main() {
    let mut s1 = String::from("Hello");

    change_str(&mut s1);

    println!("s1 = {}", s1);
}

fn change_str(str: &mut String) {
    str.push_str(" World!");
}
