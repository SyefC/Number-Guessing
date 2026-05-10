package main

import (
	"os" 
    "math/rand/v2"          //mod
	"fmt"
)

var n int = rand.IntN(101)

func main(){
var input int //main menu input

fmt.Print("\033[H\033[2J")//clear terminal

fmt.Printf("Wanna play a game? (type 1 or 0): ")
fmt.Scan(&input)

if(input == 1){ //game
	fmt.Print("\033[H\033[2J")
	for{
	var run bool = game()
     if(run == true){
		break
	 }else if(run == false){
		continue
	 }else{
		fmt.Println("error exiting program...!")
		os.Exit(0)
	 }
	}

} else{
	fmt.Println("Exiting...")
	os.Exit(0)
}
}

func game() bool{ //game code
var guess int
fmt.Printf("Guess a number through 1-100: ")
fmt.Scan(&guess)
if(guess == n){
	fmt.Print("\033[H\033[2J")
	fmt.Printf("                                CORRECT! your input: %d == generated number: %d\n", guess, n)
	fmt.Printf("                                                  \033[32mYOU WINNNN\033[0m\n                                              Thanks For Playing!\n\n\n\n\n")
	return true
}else if(guess < n){
	fmt.Print("\033[H\033[2J")
	fmt.Printf("greater\n")
	return false
}else if(guess > n){
	fmt.Print("\033[H\033[2J")
	fmt.Printf("lesser\n")
	return false
}else{
	fmt.Println("error: invalid input or program failure. please run program back if its fail. exiting...")
	os.Exit(0)
}
return true
}