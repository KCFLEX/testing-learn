package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	quitChan := make(chan bool)

	go readUserInput(os.Stdin, quitChan)

	<-quitChan

	close(quitChan)

	fmt.Println("goodbye")

}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("----------")
	fmt.Println("Enter a whole number to check if it's prime number or not.  Enter q to quit")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(in io.Reader, quitChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			quitChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()
	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	// convert whatever user enters to int
	numCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "enter whole number", false
	}

	_, msg := isPrime(numCheck)

	return msg, false
}

func isPrime(num int) (bool, string) {
	// prime numbers do not include 0 and 1
	if num < 2 {
		return false, "number less than 2 can't be a prime number"
	}

	if num == 2 {
		return true, "2 is a prime number"
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false, fmt.Sprintf("%v is not a prime number because it is divided by %d", num, i)
		}
	}

	return true, fmt.Sprintf("%v is prime number", num)
}
