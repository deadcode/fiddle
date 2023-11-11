use std::fs::File;
use std::io::ErrorKind;

fn main() {
    let f = File::open("hello.txt");

    let f = match f {
        Ok(file) => file,
        Err(ref error) if error.kind() == ErrorKind::NotFound => {
            match File::create("hello.txt") {
                Ok(f) => {
                    println!("Created the file {:?}", f);
                    f 
                },
                Err(e) => {
                    panic!("Could not create the file {:?}", e);
                },
            }
        },
        Err(error) => {
            panic!("There was an error open the file: {:?}", error);
        },
    };     
}
