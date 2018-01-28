package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

var stores map[string]story

func init() {
	stores = make(map[string]story)

	body, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &stores)

	if err != nil {
		panic(err)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/options/", makeHandler(optionsHandler))

	http.ListenAndServe(":8080", mux)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	sto := stores["intro"]
	RenderTemplate(w, sto)
}

func optionsHandler(w http.ResponseWriter, r *http.Request, path string) {
	sto := stores[path]
	RenderTemplate(w, sto)
}

var validPath = regexp.MustCompile("^/(options)/([a-zA-Z0-9-]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
