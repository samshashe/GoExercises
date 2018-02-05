package htmlparser

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var links []Link

var excludedInvalidLinks = map[string]bool{
	"print=1":     true,
	"mailto:":     true,
	"replytocom":  true,
	"format=feed": true,
	".pdf":        true,
	".png":        true,
	".jpg":        true,
	".jpeg":       true,
	".zip":        true,
	".exe":        true,
	".gif":        true,
	".ttf":        true,
	".feed":       true,
	".bmp":        true,
	".m4a":        true,
	".mp3":        true,
	".wav":        true,
	".gz":         true,
	".rar":        true,
	".svg":        true,
	".tif":        true,
	".js":         true,
	".doc":        true,
}

func GetLinks(htmlString, inputUrl string) []Link {
	links = nil
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	return getAllLinks(doc, inputUrl)
}

func GetDomainLinks(htmlString, inputUrl string) []Link {
	links = nil
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	allLinks := getAllLinks(doc, inputUrl)
	var domainLinks []Link
	for _, link := range allLinks {
		if strings.Contains(link.Href, inputUrl) {
			fmt.Println("Found +++++ ", link.Href)
			domainLinks = append(domainLinks, link)
		}
	}
	return domainLinks
}

func GetHost(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	return u.Host
}
func GetDomain(inputUrl string) string {
	if len(inputUrl) == 0 {
		return ""
	}
	u, err := url.Parse(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	return u.Scheme + "://" + u.Host
}

func getAllLinks(n *html.Node, parentUrl string) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		link := Link{}

		for _, a := range n.Attr {
			if a.Key == "href" {
				link.Href = a.Val
				//link.Text = strings.TrimSpace(n.FirstChild.Data)
				break
			}
		}

		if len(GetHost(link.Href)) == 0 {

			link.Href = GetDomain(parentUrl) + link.Href
			links = append(links, link)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAllLinks(c, parentUrl)
	}
	return links
}

type Link struct {
	Href string
	Text string
}
