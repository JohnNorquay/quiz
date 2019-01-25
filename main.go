package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csv1 := flag.String("csvFile", "problems.csv", "file for Questions format: question, answer")
	flag.Parse()

	csvFile, err := os.Open(*csv1)
	Check(err)
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	Check(err)

	Problems := make([]Problem, len(lines))
	for i, line := range lines {
		Problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}

	numCorrect := 0
	var answer string
	for i, line := range Problems {
		fmt.Printf("Question %d: %s  ", i+1, line.Question)
		fmt.Scan(&answer)

		if answer == line.Answer {
			numCorrect++
			fmt.Print("Correct\n")
		} else {
			fmt.Println("Wrong")
		}
	}

	fmt.Printf("\n\nYou got %d out of %d correct.\n\n", numCorrect, len(Problems))
}

func Check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
