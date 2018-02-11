package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", basic)
	http.HandleFunc("/trailing", trailing)
	http.HandleFunc("/send-trailers-example", trailers_examples)
	http.ListenAndServe(":3030", nil)
}

func basic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server-Timing", `database; dur=1.9; desc="Database Request",render; dur=2.6; desc="Render"`)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<body><h1>Hello World</h1></body>")
}

func trailing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Trailer", "Server-Timing")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<body><h1>Hello World</h1></body>")

	w.Header().Set("Server-Timing", `database; dur=1.9; desc="Database Request",render; dur=2.6; desc="Render"`)
}

func trailers_examples(w http.ResponseWriter, req *http.Request) {
	// Before any call to WriteHeader or Write, declare
	// the trailers you will set during the HTTP
	// response. These three headers are actually sent in
	// the trailer.
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "AtEnd3")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

	w.Header().Set("AtEnd1", "value 1")
	io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
	w.Header().Set("AtEnd2", "value 2")
	w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.
}
