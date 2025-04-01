/*
A palindrome is a word, verse, or sentence that reads the same backward or forward,
such as 'Able was I ere I saw Elba,' or a number like 1881.

Write a function named is_palindrome() that checks whether a given string is a palindrome or not.
The function should take a string as input and return a boolean value indicating whether the string is a palindrome or not.
*/

fn main() {
    let input = String::from("1211");
    println!(
        "It is {:?} that the given string is palindrome",
        palindrome(input)
    );
}

fn palindrome(input: String) -> bool {
    /* Your Code here */
    let mut first = 0;
    let mut last = input.len() - 1;

    let input_vec = input.as_bytes();

    while first < last {
        if input_vec[first] != input_vec[last] {
            return false;
        }
        first += 1;
        last -= 1;
    }
    true
}