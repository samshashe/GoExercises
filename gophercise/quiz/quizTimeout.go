package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
var (
	flagFilePath string
	flagRandom   bool
	flagTime     int
)
func init() {
	flag.StringVar(&flagFilePath, "file", "questions.csv", "path/to/csv_file")
	flag.BoolVar(&flagRandom, "random", true, "randomize order of questions")
	flag.IntVar(&flagTime, "time", 10, "test duration")
}*/

func main14() {

	path := "problems.txt"
	timeout := time.Duration(5)

	if len(os.Args) > 2 {
		path = os.Args[1]

		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		timeout = time.Duration(t)
	}

	inFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return
	} else {
		defer inFile.Close()
	}

	scanner := bufio.NewScanner(inFile)
	total, correct := 0, 0
	for scanner.Scan() {
		input := make(chan string, 1)
		line := scanner.Text()
		slice := strings.SplitN(line, ",", 2)

		fmt.Println("Question: what is the result of: ", slice[0]) // the line
		fmt.Print("Enter the answer: ")
		go readInput(input)

		select {
		case i := <-input:
			if slice[1] == strings.TrimRight(i, "\n") {
				correct++
			}
		case <-time.After(timeout * time.Second):
			fmt.Println("timed out")
		}

		total++
	}

	fmt.Printf("You answered %d questions out of %d \n", correct, total)
}

func readInput(input chan string) {
	reader := bufio.NewReader(os.Stdin)

	userAnswer, _ := reader.ReadString('\n')
	input <- userAnswer
}
