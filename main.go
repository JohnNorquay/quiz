package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem struct allows us to give structure to our quiz.  By utilizing This
// approach we can have different types of input and still have the tests
// be functional
type Problem struct {
	Question string
	Answer   string
}

func main() {
	csv1 := flag.String("csvFile", "problems.csv", "file for Questions format: question, answer.")
	timeLimit := flag.Int("time", 5, "Define the time limit for the quiz.")
	flag.Parse()

	csvFile, err := os.Open(*csv1)
	Check(err)
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	Check(err)

	problems := ParseLines(lines)

	fmt.Println("Press <enter> when ready to begin quiz:")
	fmt.Scanln()
	fmt.Println("Good Luck!!\n\n")

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	numCorrect := 0
	ansChan := make(chan string)

	// problemloop is like a goto statement.  When the timer is up, the program
	// breaks back to problemloop and then finishes by printing results of the quiz.
problemloop:

	// we are running through each question that is already read into memory.
	// we are utilizing an anonymous function as a go routine.  This will run
	// in the background and keep track of how much time is left for completion
	// of the quiz.
	for i, line := range problems {
		fmt.Printf("Question %d: %s  ", i+1, line.Question) // question is asked.
		go func() {
			var answer string
			fmt.Scan(&answer) // waits for an answer to be input
			ansChan <- answer // the input answer is placed into a channel.
		}()

		// The select statement is the key to this program finishing as soon as the
		// timer is done. . .
		select {
		case <-timer.C: // if time is up, there is something in this channel.
			fmt.Println("Times Up!")
			break problemloop // if something is in the channel, we quit.
		case answer := <-ansChan: // if nothing is in the timer channel,
			// then we pull the answer waiting in the answer
			// channel.
			if answer == line.Answer {
				numCorrect++
				fmt.Print("Correct\n")
			} else {
				fmt.Println("Wrong")
			}
		}
	}

	fmt.Printf("\n\nYou got %d out of %d correct.\n\n", numCorrect, len(problems))
}

// ParseLines reads through a file already read into memory.  It takes each
// question and answer from a single line and populates the Problems struct.
// A slice of Problems are then created and returned.
func ParseLines(lines [][]string) []Problem {
	Problems := make([]Problem, len(lines))
	for i, line := range lines {
		Problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return Problems
}

// Check() checks for an error.  If present, returns the error and then panics.
func Check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
