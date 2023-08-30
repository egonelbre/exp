package main

import (
	"net/http"
	"net/url"
	"strings"
)

/*
This is an example how to write Go servers without routers.
It can be a useful to know how to do this, however,
your life will be probably easier with a router.

An alternative approach would be to write a switch with matchers:
https://www.youtube.com/watch?v=hmq6veCFo0Y
*/

func main() {
	http.ListenAndServe("127.0.0.1:8080", NewServer())
}

type (
	Server struct {
		Users  EntitiesServer
		Groups EntitiesServer
	}
	EntitiesServer struct {
		Name string
	}
	EntityServer struct {
		*EntitiesServer
		ID string
	}
)

func NewServer() *Server {
	server := &Server{}
	server.Users.Name = "user"
	server.Groups.Name = "group"
	return server
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch prefix, r := Route(r); prefix {
	case "":
		server.ServeDashboard(w, r)
	case "user":
		server.Users.ServeHTTP(w, r)
	case "group":
		server.Groups.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (server *Server) ServeDashboard(w http.ResponseWriter, r *http.Request) {
	ServeBody(w, `
		<a href="user/">Users</a>
		<a href="group/">Groups</a>
	`)
}

func (server *EntitiesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch prefix, r := Route(r); prefix {
	case "":
		server.ServeDashboard(w, r)
	default:
		// use an ephemeral server for serving subroute
		// you can also validate ID here
		(&EntityServer{
			EntitiesServer: server,
			ID:             prefix,
		}).ServeHTTP(w, r)
	}
}

func (server *EntityServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type R [2]string
	switch prefix, r := Route(r); (R{r.Method, prefix}) {
	case R{http.MethodGet, ""}:
		server.ServeDashboard(w, r)
	case R{http.MethodGet, "id"}:
		w.Write([]byte(server.ID))
	case R{http.MethodDelete, ""}:
		http.Error(w, "you are not allowed to do this", http.StatusUnauthorized)
	default:
		http.NotFound(w, r)
	}
}

func (server *EntitiesServer) ServeDashboard(w http.ResponseWriter, r *http.Request) {
	ServeBody(w, `
		<a href="1/">`+server.Name+` 1</a>
		<a href="2/">`+server.Name+` 2</a>
	`)
}

func (server *EntityServer) ServeDashboard(w http.ResponseWriter, r *http.Request) {
	ServeBody(w, `
		<a href="id">ID</a>
	`)
}

// ServeBody is just to demonstrate serving some content.
func ServeBody(w http.ResponseWriter, body string) {
	w.Write([]byte("<html><style>a { display: block }</style><body>" + body + "</html></body>"))
}

func Route(r *http.Request) (prefix string, r2 *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	i := strings.IndexByte(path, '/')
	if i >= 0 {
		prefix = path[:i]
		rawpath := strings.TrimLeft(r.URL.RawPath, "/")
		if k := strings.IndexByte(rawpath, '/'); k >= 0 {
			rawpath = rawpath[:k]
		}
		r2 = CloneRequest(r)
		r2.URL.Path = path[i:]
		r2.URL.RawPath = rawpath
		return prefix, r2
	}
	r2 = CloneRequest(r)
	r2.URL.Path = ""
	r2.URL.RawPath = ""
	return path, r2
}

func CloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r
	r2.URL = new(url.URL)
	*r2.URL = *r.URL
	return r2
}
