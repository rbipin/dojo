use std::io;

fn main() {
    println!("Get User Input");

    let mut user_input = String::new();
    io::stdin()
        .read_line(&mut user_input)
        .expect("Failure to read line");
    
    println!("Your input: {}", user_input);
}
