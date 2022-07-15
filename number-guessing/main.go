package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Implement a number-guessing game in which the computer computes a four digit number as a secret number and a player tries to guess that number correctly.
//Player would enter her guess and the computer would produce a feedback on the positions of the digits.
//Four-digit number can't start with 0 and have repeating digits.
//Let's say the computer computes 2658 as a secret number to be guessed by the player.
//When player enters her guess such as 1234, the computer would display -1 meaning that only one digit of 1234 exist in the secret number and its position is wrong.
//When the player enters 5678 the similarly the computer displays +2-1.
//And the game goes on until the player correctly guess the secret number and the computer displays +4.
//The game also keeps track of all guesses entered by the players so far and lists them when it displays its feedback to the player so that the player can compute her next guess correctly.

func main() {

	var secretNumber []int         //secret number of computer
	secretNumberText := []string{} //string version of

	var prediction string

	//Provide seed
	rand.Seed(time.Now().Unix())
	digitArray := rand.Perm(10) // The array will be pseudo-random permutation of the integers in the range [0,n).

	rightPlace := 0 // to count correct numbers at right place
	wrongPlace := 0 // to count correct numbers at the wrong place
	attempts := 0   // to count attempts

	if digitArray[0] == 0 {
		secretNumber = digitArray[1:5]
	} else {
		secretNumber = digitArray[0:4]
	}
	fmt.Println("Welcome to the Number Guessing Game!!! ~(˘▾˘~)")
	fmt.Println("You need to predict my 4-digits Secret Number !")
	fmt.Println("Don't worry the first element of the secret number is not 0 and,")
	fmt.Println("Each digit in my secret number does not repeat. ╮ (. ^ ᴗ ^.) ╭")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("DO NOT FORGET")
	fmt.Println(" + sign stands for correct answers in the right place")
	fmt.Println(" - sign stands for correct answers in the wrong place")

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range secretNumber {
		number := secretNumber[i]
		text := strconv.Itoa(number)
		secretNumberText = append(secretNumberText, text)
	}

	// Join our string slice.
	result := strings.Join(secretNumberText, "")

	//while rightPlace is not 4, repeat
	for rightPlace < 4 {
		rightPlace = 0 //to reset counts for each prediction try
		wrongPlace = 0
		attempts++
		fmt.Print("Enter Your Prediction: ")
		fmt.Scanln(&prediction)
		fmt.Println("You entered:", prediction)
		//Turn string into an int array
		split := strings.Split(prediction, "")
		predArray := make([]int, len(split))
		for i := range predArray {
			predArray[i], _ = strconv.Atoi(split[i])
		}
		//Compare arrays
		for m := range predArray {
			for n := range secretNumber {
				if m == n && secretNumber[n] == predArray[m] {
					rightPlace++
				} else if secretNumber[n] == predArray[m] {
					wrongPlace++
				}
			}
		}

		//Prediction end conditions
		if rightPlace == 4 {
			fmt.Println("+", rightPlace)
			fmt.Println("Congratulations! You got it on the", attempts, "th try! (^o^)/ ")
			fmt.Println("My secret number was", result, "! (ᵔᴥᵔ)")
		} else {
			fmt.Println("+", rightPlace, "-", wrongPlace)
		}
	}
}
