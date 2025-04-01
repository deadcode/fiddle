// Problem 1:
/*
Write a program to find the difference between the square of the sum and the sum of the squares of the first N numbers.
N will be a user-defined input that your program will take.

For example, if the user enters the number 5, you should compute the square of the sum as (1 + 2 + 3 + 4 + 5)^2 = 225.
Next, compute the sum of the squares as (1^2 + 2^2 + 3^2 + 4^2 + 5^2) = (1 + 4 + 9 + 16 + 25) = 55.
Finally, calculate the difference as 225 - 55 = 170.
*/

fn main() {
    let mut n = String::new();
    std::io::stdin()
        .read_line(&mut n)
        .expect("failed to read input.");
    let n: i32 = n.trim().parse().expect("invalid input");

    let square_of_sum = square_of_sum(n);
    println!("square of sum = {square_of_sum}");
    let sum_of_squares = {
        let mut sum: i32 = 0;
        for i in 1..=n {
            sum += i * i;
        }
        sum
    };
    println!("sum of squares = {sum_of_squares}");
    
    /* Complete the code after this line */
    println!("Diff = {}", square_of_sum - sum_of_squares); 
}

fn square_of_sum(n: i32) -> i32 {
    let mut i: i32 = 0;
    let mut sum: i32 = 0;

    while i <= n {
        sum += i;
        i += 1;
    }
    sum * sum
}