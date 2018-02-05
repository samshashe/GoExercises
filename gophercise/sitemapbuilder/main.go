package main

import (
	"GoExercises/gophercise/htmllinkparser/htmlparser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	input := "http://www.dw.com/am"

	GetLink(input)
	// for k, _ := range linksVisited {
	// 	fmt.Println(k)
	// }
	fmt.Printf("All links found in the domain are %d Links", len(linksVisited))
}

var linksVisited = make(map[string]bool)
var host string

func GetLink(inputUrl string) {
	fmt.Println("Parent: -------->", inputUrl)
	linksVisited[inputUrl] = true
	response, err := http.Get(inputUrl)

	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	links := htmlparser.GetDomainLinks(string(body), inputUrl)

	for _, link := range links {
		_, ok := linksVisited[link.Href]
		if ok == false {
			GetLink(link.Href)
		}
	}
}

func GetHost(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	return u.Host
}
