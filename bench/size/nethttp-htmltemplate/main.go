package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	template.Must(template.New("").Parse("")).Execute(os.Stdout, nil)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
