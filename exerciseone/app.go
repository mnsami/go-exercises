package exerciseone

import (
	"fmt"
	"os"
	"log"
	"bufio"
	s "strings"
)

const problemsFilename = "./problems.csv"

// check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile to read file from disk
func openFile() {
	f, err := os.Open(problemsFilename)
	if err != nil {
        log.Fatal(err)
    }
	
	defer func() {
        if err = f.Close(); err != nil {
        	log.Fatal(err)
		}
	}()
	
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		problem, answer := s.Split(scanner.Text(), ',')
        fmt.Println(problem, answer)
	}
	
    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
}

// Run method to run the exercise 1
func Run() {
	openFile()
}
