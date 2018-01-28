package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main12() {

	path := "problems.txt"

	if len(os.Args) > 0 {
		path = os.Args[1]
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
		line := scanner.Text()
		slice := strings.SplitN(line, ",", 2)

		fmt.Println("Question: what is the result of: ", slice[0]) // the line

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the answer: ")
		userAnswer, _ := reader.ReadString('\n')
		//fmt.Println(slice[1], userAnswer)
		if slice[1] == strings.TrimRight(userAnswer, "\n") {
			correct++
		}
		total++
	}

	fmt.Printf("You answered %d questions out of %d \n", correct, total)
}
