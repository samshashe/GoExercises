package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := []byte{'a', 'z', 'A', 'Z'}
	fmt.Println(str)

	fmt.Println(CountCamelCaseWords("saveChangesInTheEditor"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text to be encrypted: ")
	userAnswer, _ := reader.ReadString('\n')
	fmt.Println(EncryptString(userAnswer, 2))

	var i byte
	for i = 0; i <= 36; i++ {
		fmt.Println(i, ".", EncryptString("samson", i))
	}

}
