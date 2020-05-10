// safe-state shows an example how to do struct changes and persist them only on a non-nil error.
package main

import (
	"fmt"
)

type Ship struct {
	Vel Vector
}

type Vector struct {
	X, Y float64
}

type Event interface{}

type Input struct {
	Dir Vector
}

func (ship *Ship) guard() (*Ship, func(err *error)) {
	x := *ship
	return &x, func(err *error) {
		if *err == nil {
			*ship = x
		}
	}
}

func (ship *Ship) Handle(ev Event) (err error) {
	defer ship.guardo(&ship)(&err)

	switch ev := ev.(type) {
	case Input:
		ship.Vel.X += ev.Dir.X
		ship.Vel.Y += ev.Dir.Y
		if ship.Vel.X > 10 || ship.Vel.Y > 10 {
			return fmt.Errorf("velocity too large")
		}
	default:
		return fmt.Errorf("invalid event %T", ev)
	}

	return nil
}

func (ship *Ship) guardo(p **Ship) func(err *error) {
	x := *ship
	*p = &x
	return func(err *error) {
		if *err == nil {
			*ship = x
		}
	}
}

func (ship *Ship) Handleo(ev Event) (err error) {
	defer ship.guardo(&ship)(&err)

	switch ev := ev.(type) {
	case Input:
		ship.Vel.X += ev.Dir.X
		ship.Vel.Y += ev.Dir.Y
		if ship.Vel.X > 10 || ship.Vel.Y > 10 {
			return fmt.Errorf("velocity too large")
		}
	default:
		return fmt.Errorf("invalid event %T", ev)
	}

	return nil
}

func main() {
	{
		ship := Ship{}
		for i := 0; i < 6; i++ {
			err := ship.Handle(Input{Dir: Vector{1, 2}})
			fmt.Println(ship, err)
		}
	}
	{
		ship := Ship{}
		for i := 0; i < 6; i++ {
			err := ship.Handleo(Input{Dir: Vector{1, 2}})
			fmt.Println(ship, err)
		}
	}
}
