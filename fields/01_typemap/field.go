package main

type Field interface {
	Name() string
}

type Uint struct {
	ID    string
	Value uint64
}

func (a *Uint) Name() string      { return a.ID }
func (a *Uint) Add(b *Uint) *Uint { return &Uint{Value: a.Value + b.Value} }
func (a *Uint) Sub(b *Uint) *Uint { return &Uint{Value: a.Value - b.Value} }

type Float struct {
	ID    string
	Value float64
}

func (a *Float) Name() string        { return a.ID }
func (a *Float) Add(b *Float) *Float { return &Float{Value: a.Value + b.Value} }
func (a *Float) Sub(b *Float) *Float { return &Float{Value: a.Value - b.Value} }

type Error struct{ Desc string }

func (a *Error) Name() string  { return "Error" }
func (a *Error) Error() string { return a.Desc }
