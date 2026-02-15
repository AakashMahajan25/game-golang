package main

import (
	"fmt"
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	gameCompleted := false

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	target := r.Intn(100) + 1
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100. Can you guess it?")



	for !gameCompleted {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)

		n, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		}

		if n < target {
			fmt.Println("Too low! Try again.")
		} else if n > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the number!")
			gameCompleted = true
		}

	}
}