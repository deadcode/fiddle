// Problem 1: 
/* You are tasked with implementing a library management system using Rust. 
Your goal is to design a program that can handle books and magazines. 
To fulfill the requirements, follow the steps below:

Create a structure called Item with the following fields:
id: An integer representing the unique identifier of the item.
title: A string representing the title of the item.
year: An integer representing the publication year of the item.
type: an enumeration type. The details are coming below.

Create an enumeration called ItemType with two variants:
Book: Represents a book.
Magazine: Represents a magazine.

Implement a function called display_item_info() that takes an Item as input 
and displays its information. The function should output 
the item's ID, title, publication year, and type (book or magazine). 
*/ 
#[derive(Debug)]
enum ItemType {
    Book,
    Magazine,
}

#[derive(Debug)]
struct Item {
    id: u64,
    title: String,
    year: u32,
    kind: ItemType,
}

fn main() {
    let item1 = Item{
        id: 123456,
        title: "A Brief History of Time".to_string(),
        year: 1988,
        kind: ItemType::Book,
    };

    println!("Item1: {:?}", item1);
}