// Problem 1:
/* In this exercise, you will be working on creating a student management system
using Rust. The system should allow you to store and retrieve student information
based on their unique ID. For ease of work, the student structure is already
created in the code below

Next, create a StudentManager structure containing a field of student, which
will essentially be a hashmap where the key part will be an integer representing
unique ID of student and the value part will be the complete details of the
students contained in the student structure.

The StudentManager should implement the following methods:
1. new() -> Self: A constructor that initializes an empty student manager.

2. add_student(&mut self, student: Student) -> Result<(), String>:
Adds a student to the manager.
If the student's ID already exists, return an error message.
Otherwise, add the student to the manager and return Ok.

3. get_student(&self, id: i32) -> Option<&Student>: Retrieves a student
from the manager based on their ID.
If the student is found, return Some(student). Otherwise, return None.

Your task is to implement the StudentManager structure, and the mentioned methods.
Additionally, provide a sample usage of the student management system by adding
a few students and retrieving their information using the get_student() method.
*/

use std::collections::HashMap;

#[derive(Debug)]
struct Student {
    id: i32,
    name: String,
    grade: String,
}

struct StudentManager {
    map: HashMap<i32, Student>,
}

impl StudentManager {
    fn new() -> Self {
        StudentManager {
            map: HashMap::new(),
        }
    }

    fn add_student(&mut self, student: Student) -> Result<(), String> {
        if self.map.contains_key(&student.id) {
            Err(format!("Student {}:{} already exists", student.id, student.name))
        } else {
            self.map.insert(student.id, student);
            Ok(())
        }
    }

    fn get_student(&self, id: i32) -> Option<&Student> {
        self.map.get(&id)
    }
}

fn main() {
    let mut students = StudentManager::new();

    if let Err(e) = students.add_student(Student { id: 1001, name: "Batman".to_string(), grade: "12".to_string() }) {
        println!("Could not insert student: {e}");
    };
    if let Err(e) = students.add_student(Student { id: 2001, name: "Superman".to_string(), grade: "11".to_string() }) {
        println!("Could not insert student: {e}");
    }
    if let Err(e) = students.add_student(Student { id: 3001, name: "Black Adam".to_string(), grade: "5".to_string() }) {
        println!("Could not insert student: {e}");
    }

    println!("StudentManager has {} students", students.map.len());
    for (id, student) in students.map.iter() {
        println!("Found {id} at {:?}", student);
    }
    for student_id in vec![1001, 5000, 3001] {
        match students.get_student(student_id) {
            Some(student) => println!("Student id {}, found in grade {} has name {}", student.id, student.grade, student.name),
            None => println!("Student id {} not found in StudentManager", student_id),
        }
    }
}
