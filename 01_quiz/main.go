package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	filename := flag.String("filename", "problems.csv", "enter a filename (must be .csv)")
	flag.Parse()
	fmt.Println(*filename)

	//read file
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Error: can't open %s", *filename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error: check the input file content.")
	}
	problems := parseLines(lines)
	fmt.Println(problems)


}

func exit(message string) {
	fmt.Sprintln(message)
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
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
