package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number game!
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
You can try in with two numbers at the same time.
The greater your number is, the harder it gets.
Wanna play?`
)

func main() {
	msgWin := make([]string, 0)
	msgWin = append(msgWin, "Woohoo! You should buy a lottery ticket!", "YOU WIN!", "Great success!")
	msgLoose := make([]string, 0)
	msgLoose = append(msgLoose, "Ha-ha, looser!", "Nah, wanna try again?", "It's just not your day")
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	var guess2 int
	if len(args) == 3 {
		guess2, err = strconv.Atoi(args[len(args)-2])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}

	}
	if guess < 0 || guess2 < 0 {
		fmt.Println("Pick up a positive number")
		return
	}

	min := guess
	if guess < guess2 {
		min = guess2
	}

	var balancer int
	if guess > 10 {
		balancer = guess / 4
	}

	for turn := 1; turn < maxTurns+balancer; turn++ {
		n := rand.Intn(min) + 1

		if n == guess || n == guess2 {

			if verbose == true {
				fmt.Printf("%d ", n)
			}

			if turn == 1 {
				fmt.Println("You won it from the first try!")
				return
			} else {
				fmt.Printf(msgWin[rand.Intn(len(msgWin))])
				return
			}
		}
	}
	fmt.Printf(msgLoose[rand.Intn(len(msgLoose))])
}
