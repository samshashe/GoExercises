package main

import (
	"exercise/gophercise/htmllinkparser/htmlparser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	input := "http://meskelrestaurant.com/"
	host = GetHost(input)
	GetLink(input)
	for k, _ := range linksVisited {
		fmt.Println(k)
	}
}

var linksVisited = make(map[string]int)
var host string

func GetLink(url string) {
	linksVisited[url] = linksVisited[url] + 1
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	links := htmlparser.GetLinks(string(body))

	for _, link := range links {
		//fmt.Printf("%d. Href: %s  Text: %s \n", i+1, link.Href, link.Text)
		fullpath := link.Href
		if isRelative := strings.Index(link.Href, "://"); isRelative == -1 {
			fullpath = "http://" + host + "/" + link.Href
		}

		_, ok := linksVisited[fullpath]
		if ok == false && GetHost(url) == host {
			GetLink(fullpath)
		}
	}
}

func GetHost(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	return u.Host
}
