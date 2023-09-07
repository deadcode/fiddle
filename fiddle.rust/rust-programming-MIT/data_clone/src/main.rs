fn main() {
    let s1 = String::from("Hello World!");
    //let s2 = s1;
    let s2 = s1.clone();

    println!("Strings are s1: {}, s2: {}", s1, s2);
}
