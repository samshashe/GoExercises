package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"exercise/gophercise/htmllinkparser/htmlparser"
)

func main() {
	fileNames := []string{"data/ex1.html", "data/ex2.html", "data/ex3.html", "data/ex4.html"}

	for _, fileName := range fileNames {
		fmt.Println("\nPrinting for file ", fileName)
		htmlString := readHtml(fileName)

		links := htmlparser.GetLinks(htmlString)
		for i, link := range links {
			fmt.Printf("%d. Href: %s  Text: %s \n", i+1, link.Href, link.Text)
		}
	}

}

func readHtml(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
