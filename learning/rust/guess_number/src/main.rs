use std::io;

fn main() {
    println!("Guess a Number!");

    let mut guess = String::new();
    io::stdin()
        .read_line(&mut guess)
        .expect("failure in read_lines");
    
    println!("You guessed: {}", guess);
}
