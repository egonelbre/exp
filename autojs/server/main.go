package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

const ClientRootDir = "client"

var (
	addr = flag.String("listen", "127.0.0.1:8080", "port to listen to")
	dev  = flag.Bool("dev", false, "development mode")

	interval = flag.Duration("i", 300*time.Millisecond, "poll interval")
)

func main() {
	flag.Parse()

	log.Fatal(http.ListenAndServe(*addr, http.HandlerFunc(Serve)))
}

func Serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	switch r.URL.Path {
	case "/":
		ServeIndex(w, r)
	case "/~reload.js":
		ServeReloadJS(w, r)
	default:
		ServeClientFile(w, r)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ServeClientFile(w http.ResponseWriter, r *http.Request) {
	path := filepath.FromSlash(path.Join(ClientRootDir, r.URL.Path))
	// TODO: clean-path
	http.ServeFile(w, r, path)
}
