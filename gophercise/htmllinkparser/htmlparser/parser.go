package htmlparser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

var links []Link

func GetLinks(htmlString string) []Link {
	links = nil
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	return getAllLinks(doc)
}

func getAllLinks(n *html.Node) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		link := Link{}

		for _, a := range n.Attr {
			if a.Key == "href" {
				link.Href = a.Val
				link.Text = strings.TrimSpace(n.FirstChild.Data)
				break
			}
		}
		links = append(links, link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAllLinks(c)
	}
	return links
}

type Link struct {
	Href string
	Text string
}
