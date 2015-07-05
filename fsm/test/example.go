package main

import (
	"fmt"
	"time"
)

type ElectricCircuit struct {
	Description string
	Connected   bool
}

type Light struct {
	Room    string
	HasBulb bool
}

type Plumbing struct {
	Location      string
	NumberOfLeaks uint
}

type House struct {
	Address          string
	Area             uint
	ElectricCircuits []*ElectricCircuit
	Lights           []*Light
	PlumbingSystems  []*Plumbing
}

func (circuit *ElectricCircuit) NeedsRepairs() bool {
	time.Sleep(1 * time.Microsecond)
	return !circuit.Connected
}

func (light *Light) NeedsRepairs() bool {
	time.Sleep(1 * time.Microsecond)
	return !light.HasBulb
}

func (plumbing *Plumbing) NeedsRepairs() bool {
	time.Sleep(1 * time.Microsecond)
	return plumbing.NumberOfLeaks > 0
}

func (house *House) NeedsRepairs() bool {
	for _, circuit := range house.ElectricCircuits {
		if circuit.NeedsRepairs() {
			return true
		}
	}

	for _, light := range house.Lights {
		if light.NeedsRepairs() {
			return true
		}
	}

	for _, plumbing := range house.PlumbingSystems {
		if plumbing.NeedsRepairs() {
			return true
		}
	}

	return false
}

func HugeHouse() *House {
	h := &House{
		Address: "123 Fake Street",
		Area:    999,
	}
	for i := 0; i < 1e6; i += 1 {
		h.ElectricCircuits = append(h.ElectricCircuits, &ElectricCircuit{"Main", true})
	}
	for i := 0; i < 1e6/2; i += 1 {
		h.Lights = append(h.Lights, &Light{"Kitchen", true}, &Light{"Dining", true})
	}
	for i := 0; i < 1e6; i += 1 {
		h.PlumbingSystems = append(h.PlumbingSystems, &Plumbing{"Bathroom", 0})
	}
	return h
}

func main() {
	house := HugeHouse()
	start := time.Now()
	if house.NeedsRepairs() {
		fmt.Println("Fix-it! Fix-it! Fix-it!")
	} else {
		fmt.Println("All good.")
	}
	fmt.Println(time.Since(start))
}
