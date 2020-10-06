package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// flag package for command line parsing
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 2, "the time limit for the quiz in seconds")
	flag.Parse()

	// file opened using os.Open function
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFileName))
	}
	defer file.Close()

	// file reading using new reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	// fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timer.C

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n", correct, len(problems))
			return

		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
			// default:
			// 	// fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			// 	// var answer string
			// 	// fmt.Scanf("%s\n", &answer)
			// 	if answer == p.a {
			// 		correct++
			// 	}
		}
	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

// parsing each line and builds a slice of problems struct
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // trimming of the unnecessary space in the csv file, if exists
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
