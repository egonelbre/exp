package main

import "fmt"

type Transition struct {
	When State
	And  Event
	Then State
	Do   func()
}

type State int
type Event int

type trigger struct {
	In   State
	When Event
}

type action struct {
	To State
	Do func()
}

type Table []Transition

type Machine struct {
	State  State
	Table  Table
	lookup map[trigger]action
}

func (m *Machine) init() {
	m.lookup = make(map[trigger]action, len(m.Table))
	for _, t := range m.Table {
		m.lookup[trigger{t.When, t.And}] = action{t.Then, t.Do}
	}
}

func (m *Machine) Handle(e Event) {
	if m.lookup == nil {
		m.init()
	}
	action, ok := m.lookup[trigger{m.State, e}]
	if !ok {
		panic(fmt.Sprintf("entry missing for %v, %v", m.State, e))
	}
	action.Do()
	m.State = action.To
}

const (
	Locked = State(iota)
	Unlocked

	Coin = Event(iota)
	Pass
)

type Device interface {
	Unlock()
	Alarm()
	ThankYou()
	Lock()
}

func NewTurnstile(dev Device) *Machine {
	return &Machine{
		State: Locked,
		Table: Table{
			{Locked, Coin, Unlocked, dev.Unlock},
			{Locked, Pass, Locked, dev.Alarm},
			{Unlocked, Coin, Unlocked, dev.ThankYou},
			{Unlocked, Pass, Locked, dev.Lock},
		}}
}

type Printer struct{}

func (t Printer) Unlock()   { fmt.Println("unlock") }
func (t Printer) Alarm()    { fmt.Println("alarm") }
func (t Printer) ThankYou() { fmt.Println("thankyou") }
func (t Printer) Lock()     { fmt.Println("lock") }

func main() {
	m := NewTurnstile(Printer{})
	m.Handle(Coin)
	m.Handle(Pass)
	m.Handle(Pass)
	m.Handle(Pass)
}
