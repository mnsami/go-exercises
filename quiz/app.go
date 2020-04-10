package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	s "strings"
	"time"
)

const (
	defaultTimeLimitInSeconds int    = 30
	defaultProblemsFilename   string = "./quiz/problems.csv"
)

var (
	timeLimitInSeconds int
	problemsFilename   string
	correctAnswersCH       = make(chan int)
	correctAnswers     int = 0
)

type problem struct {
	question string
	answer   string
}

// check error
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

// ReadFile to read file from disk
func openFile() *os.File {
	f, err := os.Open(problemsFilename)
	check(err)

	return f
}

func closeFile(file *os.File) {
	err := file.Close()
	check(err)
}

func promptProblemsToUser(problems []problem) {
	for pIdx, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", pIdx+1, problem.question)
		stdReader := bufio.NewReader(os.Stdin)
		userAnswer, _ := stdReader.ReadString('\n')
		userAnswer = s.Replace(userAnswer, "\n", "", -1)
		if s.Compare(userAnswer, problem.answer) == 0 {
			correctAnswers++
		}
	}
	correctAnswersCH <- correctAnswers
	close(correctAnswersCH)
}

// ReadProblem reads the problem
func readProblemsFromFile(file *os.File) []problem {

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	check(err)

	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: s.Replace(line[0], "\n", "", -1),
			answer:   s.Replace(s.TrimSpace(line[1]), "\n", "", -1),
		}
	}

	return problems
}

// init function to set flags
func readInputFlags() {
	flag.IntVar(&timeLimitInSeconds, "timeLimit", defaultTimeLimitInSeconds, "the time limit for the quiz in seconds")
	flag.StringVar(&problemsFilename, "problemsCsv", defaultProblemsFilename, "a csv file in the format of 'question,answer'")

	flag.Parse()
}

func startQuiz() {
	fmt.Printf("Press any key to start the quiz, you have %d seconds", timeLimitInSeconds)
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadRune()
	check(err)
}

// Run method to run the exercise One
func Run() {
	readInputFlags()

	file := openFile()
	defer closeFile(file)

	problems := readProblemsFromFile(file)
	startQuiz()

	go promptProblemsToUser(problems)

	select {
	case <-time.After(time.Duration(timeLimitInSeconds) * time.Second):
		fmt.Println("You are out of time")
		finishQuiz(correctAnswers, len(problems))
	case correctAnswers := <-correctAnswersCH:
		finishQuiz(correctAnswers, len(problems))
	}
}

func finishQuiz(correctAnswers int, problemsCount int) {
	fmt.Printf("You scored %d correct answered out of %d.\n", correctAnswers, problemsCount)
	os.Exit(1)
}
