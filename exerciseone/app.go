package exerciseone

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

// check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile to read file from disk
func ReadFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
        log.Fatal(err)
    }
	
	defer func() {
        if err = f.Close(); err != nil {
        	log.Fatal(err)
		}
	}()
	
    s := bufio.NewScanner(f)
    for s.Scan() {
        fmt.Println(s.Text())
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
