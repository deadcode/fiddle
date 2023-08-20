fn main() {
    let s1 = gives_ownership();

    let s2 = String::from("World!");

    let s3 = takes_and_gives_back(s2);
}

fn gives_ownership() -> String {
    let s1 = String::from("Hello!");

    return s1;
}

fn takes_and_gives_back(str: String) -> String {
    return str;
}
