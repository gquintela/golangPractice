package main

import (
	"encoding/csv"
	"flag"
	. "fmt"
	"os"
	"strings"
	"time"
)

const red = "\033[31m"
const green = "\033[32m"
const black = "\033[0m"

func main() {

	filename := flag.String("filename", "problems.csv", "enter a filename (must be .csv)")
	totalTime := flag.Int("timer", 10, "define a time limit (seconds)")
	flag.Parse()

	//read file
	file, err := os.Open(*filename)
	if err != nil {
		exit(Sprintf("Error: can't open %s", *filename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error: check the input file content.")
	}

	problems := parseLines(lines)

	//setup
	timer := time.NewTimer(time.Duration(*totalTime) * time.Second)
	var score int
	var userAnswer string
	ansChan := make(chan string)
	for i, problem := range problems {
		Printf("Problem number %d: %s? ", i+1, problem.q)
		go func() {
			_, _ = Scanf("%s ", &userAnswer)
			ansChan <- userAnswer
		}()
		select {
		case <-timer.C:
			Printf("\nTime's up!\nYour score: %d out of %d\n", score, len(problems))
			return
		case <-ansChan:
			userAnswer = strings.TrimSpace(userAnswer)
			if userAnswer == problem.a {
				score++
				Println(string(green), "OK!")
			} else {
				Println(string(red), "Wrong!")
			}
			Print(string(black))
		}
	}
	Printf("\nYou completed all the questions!\nYour score: %d out of %d\n", score, len(problems))

}

func exit(message string) {
	Sprintln(message)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
