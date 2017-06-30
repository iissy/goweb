package main

import (
	"fmt"
	"github"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	dir := http.Dir("public/")
	http.Handle("/", http.FileServer(dir))
	http.HandleFunc("/issue", issue)
	http.ListenAndServe(":8000", nil)
}

func issue(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("public/views/issue.html", "public/views/_list.html")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}
