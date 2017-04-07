package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func init() {
	termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)
}

type KeyPress struct {
	Char rune
	Key  termbox.Key
}

type Repaint struct{}

var Actions = make(chan interface{})

func Listen() {
	go func() {
		tick := time.NewTicker(100 * time.Millisecond)
		defer tick.Stop()
		for range tick.C {
			select {
			case Actions <- Repaint{}:
			default:
			}
		}
	}()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				close(Actions)
				return
			default:
				Actions <- KeyPress{ev.Ch, ev.Key}
			}
		case termbox.EventError:
			//TODO: proper error handling
			panic(ev.Err)
		}
	}
}

type Process struct {
	Name     string
	progress int32
	done     int32
	mu       sync.Mutex
	err      error
}

func (process *Process) Update(value int) {
	atomic.StoreInt32(&process.progress, int32(value))
}

func (process *Process) Done() {
	atomic.StoreInt32(&process.done, 1)
}

func (process *Process) Progress() (progress int, done bool) {
	progress = int(atomic.LoadInt32(&process.progress))
	done = atomic.LoadInt32(&process.done) == 1
	return
}

func (process *Process) Fail(err error) {
	process.mu.Lock()
	process.err = err
	process.mu.Unlock()
	process.Done()
}

func (process *Process) Error() error {
	process.mu.Lock()
	err := process.err
	process.mu.Unlock()
	return err
}

type Processes struct {
	mu   sync.Mutex
	list []*Process
}

func (processes *Processes) Add(process *Process) {
	processes.mu.Lock()
	defer processes.mu.Unlock()

	processes.list = append(processes.list, process)
}

func (processes *Processes) Paint() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	processes.mu.Lock()
	defer processes.mu.Unlock()

	x, y := 0, 0
	for _, process := range processes.list {
		progress, done := process.Progress()
		line := fmt.Sprintf("%-10s %4d", process.Name, progress)
		if done {
			line += " done"
		}

		x = 0
		for _, r := range line {
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
		y++
	}

	termbox.Flush()
}

var Processing Processes

func ProcessFile(char rune) {
	process := &Process{}
	process.Name = string(char)

	go func() {
		defer process.Done()

		tick := time.NewTicker(time.Millisecond * time.Duration(rand.Intn(100)+100))
		defer tick.Stop()

		k := 0
		for range tick.C {
			process.Update(k)
			k++
			if rand.Intn(500) == 0 {
				process.Fail(errors.New("we got a zeeero"))
				return
			}
		}
	}()

	Processing.Add(process)
}

func main() {
	go Listen()
	for ev := range Actions {
		switch ev := ev.(type) {
		case Repaint:
			Processing.Paint()
		case KeyPress:
			ProcessFile(ev.Char)
			Processing.Paint()
		}
	}
}
