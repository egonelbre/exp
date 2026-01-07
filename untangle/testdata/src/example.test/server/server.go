package server

import (
	"net/http"

	"example.test/ext"
	"example.test/sink"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

//untangle:print
func (s *Server) HandleChopping(w http.ResponseWriter, r *http.Request) error { // want HandleChopping:"example.test/sink.World: Ingredients, Knives" "HandleChopping: World\\{Ingredients, Knives\\}"
	sink.Chop()
	return nil
}

//untangle:print
func (s *Server) HandleMixing(w http.ResponseWriter, r *http.Request) error { // want HandleMixing:"example.test/sink.World: Bowls, Ingredients, Spoons" "HandleMixing: World\\{Bowls, Ingredients, Spoons\\}"
	world := sink.API()
	sink.Mix(world)
	return nil
}

//untangle:print
func (s *Server) HandleBowling(w http.ResponseWriter, r *http.Request) error { // want HandleBowling:"example.test/sink.World: Bowls" "HandleBowling: World\\{Bowls\\}"
	world := sink.API()
	ext.CleanBowl(world)
	return nil
}
