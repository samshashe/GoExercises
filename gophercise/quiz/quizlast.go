package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	path := "problems.csv"
	timeout := time.Duration(5)

	if len(os.Args) > 2 {
		path = os.Args[1]

		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		timeout = time.Duration(t)
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return
	} else {
		defer file.Close()
	}

	var total, correct int

	total, err = countLinesInFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	scanner := bufio.NewScanner(file)
	done := make(chan bool)

	timer := time.NewTicker(time.Second * time.Duration(timeout))

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			slice := strings.SplitN(line, ",", 2)

			fmt.Println("Question: what is the result of: ", slice[0]) // the line
			fmt.Print("Enter the answer: ")

			reader := bufio.NewReader(os.Stdin)
			userAnswer, _ := reader.ReadString('\n')

			if slice[1] == strings.TrimRight(userAnswer, "\n") {
				correct++
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-timer.C:
		fmt.Println("Time is up!")
	}

	fmt.Printf("You answered %d questions out of %d \n", correct, total)
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}
}
