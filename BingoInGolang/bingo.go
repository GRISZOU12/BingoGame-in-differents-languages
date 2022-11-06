// 🖋️Written by GRISZ (!!-|GRISZ|-!!#2705)
// Here is the program of the 'bingo.exe'
/*The program is divided in two parts ((before the Main function and after)before it's for the setup and after is to generate the number*/

// imports
package main

import (
	// bufio is imported for the inputs (scanners)
	"bufio"
	// is it really necessary to tell you?
	"fmt"
	// to generate pseudorandom numbers :
	"math/rand"
	// like the fmt import
	"os"
	// to convert strings into int
	"strconv"
	// to change the seed for the pseudorandom generator
	"time"
)

// ClearConsole is a function for clear the console
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

// Pause is a function that make a pause in the program and requires a user input to "unlock"
func Pause(pauseReason string) {
	pause := bufio.NewScanner(os.Stdin)
	fmt.Print(pauseReason)
	pause.Scan()
}

// ShowError is a function that display errors in red color
func ShowError(errorContent string) {
	fmt.Print("\033[31m")
	ClearConsole()
	fmt.Println(errorContent)
	Pause("Press enter to continue ...")
	fmt.Print("\033[37m")
	ClearConsole()
	MainInputCheck()
}

// MainSetup is the function to set up the bingo
func MainSetup() (input string) {
	ClearConsole()
	fmt.Println("+--------------------+")
	fmt.Println("|Welcome to the BINGO|")
	fmt.Println("+--------------------+")
	fmt.Println("You can start the bingo by enter 'start' or see the rules by enter 'info'.")
	fmt.Print(" > ")
	MainSetupInput := bufio.NewScanner(os.Stdin)
	MainSetupInput.Scan()
	input = MainSetupInput.Text()
	return
}

// MainInputCheck is a function to check and interpret the input from the function : MainSetup
func MainInputCheck() {
	UserInput := MainSetup()
	if UserInput == "start" || UserInput == "Start" || UserInput == "START" || UserInput == "sTART" {
		BingoSetup()
	} else if UserInput == "info" || UserInput == "Info" || UserInput == "INFO" || UserInput == "iNFO" {
		BingoRules()
	} else {
		ShowError("Please enter start or info not other thing.")
	}
}

// This is the main function that execute the rest of the code
func main() {
	MainInputCheck()
}

// BingoRules serves to display the rules
func BingoRules() {
	ClearConsole()
	fmt.Println("+-----+")
	fmt.Println("|RULES|")
	fmt.Println("+-----+")
	fmt.Println("\nThe bingo is a game where you can input a min number and a max number,\nThen a number will be randomly generated between your inputs range,\n then the game will begin and you will input numbers and find the previously random generated number.")
	Pause("\nPress enter to continue...")
	MainInputCheck()
}

// BingoSetup is necessary to quit the setup and enter the "bingo phase"
func BingoSetup() {
	ClearConsole()
	fmt.Println("+-------+")
	fmt.Println("|BINGO !|")
	fmt.Println("+-------+")
	GenRand()
}

// GenRand generate the random number
func GenRand() (RandNum int) {
	Min, Max := AskNumber()
	rand.Seed(time.Now().UnixNano())
	var ToGen int = (Max - Min) + Min
	RandNum = rand.Intn(ToGen)
	Pause("Press enter to close the program.")
	return
}

// AskNumber ask the user for the min and max numbers (this function call AskMinNumber and AskMaxNumber)
func AskNumber() (MinNumber int, MaxNumber int) {
	MinNumber = AskMinNumber()
	MaxNumber = AskMaxNumber()
	if MinNumber >= MaxNumber || MaxNumber > 9999 {
		ShowError("You entered a min number that is superior or equal to the max number or superior to 9999 or inferior to 0")
		os.Exit(0)
		return
	} else {
		return
	}
}

// AskMinNumber serves to ask the min number to the user this function is called by AskNumber
func AskMinNumber() (MinNumber int) {
	fmt.Print("Please enter the min number > ")
	InputMinNumber := bufio.NewScanner(os.Stdin)
	InputMinNumber.Scan()
	MinNum, err := strconv.Atoi(InputMinNumber.Text())
	if err != nil {
		ShowError("Please enter a valid number (a valid number is a number between 0 and 9999).")
		os.Exit(0)
		return
	} else {
		MinNumber = MinNum
		return MinNumber
	}
}

// AskMaxNumber function is called by AskNumber and ask the user for the max number
func AskMaxNumber() (MaxNumber int) {
	fmt.Print("Please enter the max number > ")
	InputMinNumber := bufio.NewScanner(os.Stdin)
	InputMinNumber.Scan()
	MaxNum, err := strconv.Atoi(InputMinNumber.Text())
	if err != nil {
		ShowError("Please enter a valid number.")
		return
	} else {
		MaxNumber = MaxNum
		return MaxNumber
	}
}
