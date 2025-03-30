fn main() {
    let s1 = String::from("Hello !");

    println!("s1 = {}", s1);

    take_ownership(s1);

    let x = 5;

    println!("x = {}", x);

    make_copy(x);

    //println!("s1 = {}, x = {}", s1, x);
    println!("x = {}", x);
}

fn take_ownership(hello_str: String) {
    println!("take_ownership: hello_str = {}", hello_str);
}

fn make_copy(y: i32) {
    println!("make_copy: y = {}", y);
}
