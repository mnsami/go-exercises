package exerciseone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
)

const problemsFilename string = "./exerciseone/problems.csv"

// check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile to read file from disk
func openFile() *os.File {
	f, err := os.Open(problemsFilename)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func promptProblemToUser(problem string, problemIndex int, answer string) bool {
	fmt.Printf("Problem #%d: %s\n", problemIndex+1, problem)
	stdReader := bufio.NewReader(os.Stdin)
	userAnswer, _ := stdReader.ReadString('\n')
	userAnswer = s.Replace(userAnswer, "\n", "", -1)

	if s.Compare(userAnswer, answer) == 0 {
		return true
	}

	return false
}

// ReadProblem reads the problem
func readProblemsFromFile(file *os.File) {
	fileScanner := bufio.NewScanner(file)
	problemIndex := 0
	correctAnswers := 0

	for fileScanner.Scan() {
		line := s.Split(fileScanner.Text(), ",")
		problem := s.Replace(line[0], "\n", "", -1)
		answer := s.Replace(s.TrimSpace(line[1]), "\n", "", -1)
		result := promptProblemToUser(problem, problemIndex, answer)
		problemIndex++
		if result == true {
			correctAnswers++
		}
	}

	fmt.Printf("You scored %d correct answered out of %d.\n", correctAnswers, problemIndex)

	err := fileScanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// Run method to run the exercise One
func Run() {
	file := openFile()
	defer closeFile(file)

	readProblemsFromFile(file)
}
