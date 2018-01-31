package main

import "math/rand"

const (
	Size  = 9
	Count = Size * Size
)

type Color byte

func RandomColor() Color {
	return Color(rand.Intn(2))
}

func (color Color) Rune() rune {
	switch color {
	case White:
		return '○' // '⚪'
	case Black:
		return '●' // '⚫'
	case Empty:
		return ' '
	case OOB:
		return '╳'
	}
	return '▓'
}

func (color Color) Opposite() Color {
	if color == White {
		return Black
	} else if color == Black {
		return White
	}
	return OOB
}

const (
	White = Color(iota)
	Black
	Empty
	OOB
)

type Regions struct{}

type Moves struct{}

type Board struct {
	Passes  int
	Plies   int
	Points  [2]int
	Moves   Moves
	Regions Regions
	Nodes   [Count]Node
}

func NewBoard() *Board { return &Board{} }

type Node struct {
	Color Color
}

func (board *Board) index(x, y int) int {
	return y*Size + x
}

func (board *Board) Set(x, y int, color Color) {
	if x < 0 || Size <= x || y < 0 || Size <= y {
		panic("OOB")
	}
	board.Nodes[board.index(x, y)].Color = color
}

func (board *Board) At(x, y int) Color {
	if x < 0 || Size <= x || y < 0 || Size <= y {
		return OOB
	}
	return board.Nodes[board.index(x, y)].Color
}

func (board *Board) Place(x, y int, color Color) {

}

func (board *Board) IsSuicide(x, y int, color Color) bool {
	return false
}
