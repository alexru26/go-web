package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("test.md")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	html := markdown.ToHTML(body, nil, nil)

	fmt.Fprint(w, string(html))
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting server at :3000...")
	http.ListenAndServe(":3000", nil)
}
