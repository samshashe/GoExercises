package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go Server")
		fmt.Fprintf(w, `<html>
			<body>
				Hello Gopher
			</body>
		</html>`)
	})

	http.HandleFunc("/error/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something has gone wrong", 500)
	})

	http.ListenAndServe(":3000", nil)
}
