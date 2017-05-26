package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Clock struct {
	Time      float64
	DeltaTime float64
	WallTime  float64
}

type Entity uint64
type Component interface{}
type System interface {
	Tick(*World)
}

type World struct {
	next Entity

	Clock    Clock
	Entities []Entity
	Storages map[reflect.Type]Storage
	Systems  []System
}

func NewWorld(systems ...System) *World {
	world := &World{}
	world.Systems = systems
	world.Storages = make(map[reflect.Type]Storage)
	return world
}

type Storage interface {
	For(func(e Entity, c Component))
	Add(e Entity, c Component)
	Get(e Entity) (Component, bool)
	Del(e Entity)
}

type storage map[Entity]Component

func (s storage) For(fn func(e Entity, c Component)) {
	for e, c := range s {
		fn(e, c)
	}
}
func (s storage) Add(e Entity, c Component)      { s[e] = c }
func (s storage) Get(e Entity) (Component, bool) { c, ok := s[e]; return c, ok }
func (s storage) Del(e Entity)                   { delete(s, e) }

func (w *World) GetStorage(c Component) Storage {
	s, ok := w.Storages[reflect.TypeOf(c)]
	if !ok {
		s = make(storage)
		w.SetStorage(c, s)
	}
	return s
}
func (w *World) SetStorage(c Component, s Storage) { w.Storages[reflect.TypeOf(c)] = s }

func (w *World) Spawn(cs ...Component) Entity {
	e := w.next
	w.next++
	w.Entities = append(w.Entities, e)

	for _, c := range cs {
		w.GetStorage(c).Add(e, c)
	}

	return e
}

func (w *World) Tick() {
	for _, sys := range w.Systems {
		sys.Tick(w)
	}
}

// usage

type Vector struct{ X, Y float32 }

type Body struct {
	Position Vector
	Velocity Vector
}

type BodyStorage struct {
	ID   map[Entity]int
	Body []Body
}

func NewBodyStorage() *BodyStorage {
	return &BodyStorage{
		ID:   make(map[Entity]int),
		Body: []Body{},
	}
}
func (s *BodyStorage) For(fn func(e Entity, c Component)) {
	for e, i := range s.ID {
		fn(e, s.Body[i])
	}
}

func (s *BodyStorage) Add(e Entity, c Component) {
	s.ID[e] = len(s.Body)
	body := c.(Body)
	s.Body = append(s.Body, body)
}
func (s *BodyStorage) Get(e Entity) (Component, bool) {
	i, ok := s.ID[e]
	if ok {
		return s.Body[i], true
	}
	return nil, false
}
func (s *BodyStorage) Del(e Entity) { panic("TODO") }

type Texture struct{ Name int }

type Physics struct{}

func (*Physics) Tick(world *World) {
	bodies := world.GetStorage(Body{}).(*BodyStorage)
	for i := range bodies.Body {
		body := &bodies.Body[i]
		body.Position.X += body.Velocity.X * float32(world.Clock.DeltaTime)
		body.Position.Y += body.Velocity.Y * float32(world.Clock.DeltaTime)
	}
}

type Renderer struct{}

func (*Renderer) Tick(world *World) {
	fmt.Println("=== FRAME ===")

	bodies := world.GetStorage(Body{})
	world.GetStorage(Texture{}).For(func(e Entity, c Component) {
		tex := c.(Texture)
		phy, ok := bodies.Get(e)
		if !ok {
			return
		}
		fmt.Println(tex.Name, phy.(Body).Position)
	})
}

func main() {
	world := NewWorld(
		&Physics{},
		&Renderer{},
	)

	world.SetStorage(Body{}, NewBodyStorage())

	var balls []Entity
	for i := 0; i < 10; i++ {
		body := Body{}
		body.Position = Vector{rand.Float32() - 0.5, rand.Float32() - 0.5}
		body.Velocity = Vector{rand.Float32() - 0.5, rand.Float32() - 0.5}
		balls = append(balls, world.Spawn(body, Texture{i}))
	}

	for {
		world.Tick()
		time.Sleep(3 * time.Second)
	}
}
