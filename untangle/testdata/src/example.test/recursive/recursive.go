package recursive

import (
	"example.test/sink"
)

// Test case 1: Simple recursive function
//
//untangle:print
func RecursiveA(world sink.World, depth int) { // want RecursiveA:"example.test/sink.World: Knives" "RecursiveA: World\\{Knives\\}"
	if depth <= 0 {
		return
	}
	_ = world.Knives
	RecursiveA(world, depth-1)
}

// Test case 2: Mutual recursion between two functions
//
//untangle:print
func MutualA(world sink.World, depth int) { // want MutualA:"example.test/sink.World: Bowls, Ingredients" "MutualA: World\\{Bowls, Ingredients\\}"
	if depth <= 0 {
		return
	}
	_ = world.Ingredients
	MutualB(world, depth-1)
}

func MutualB(world sink.World, depth int) { // want MutualB:"example.test/sink.World: Bowls, Ingredients"
	if depth <= 0 {
		return
	}
	_ = world.Bowls
	MutualA(world, depth-1)
}

// Test case 3: Function calling another that hasn't been analyzed yet
//
//untangle:print
func CallerFirst(world sink.World) { // want CallerFirst:"example.test/sink.World: Spoons" "CallerFirst: World\\{Spoons\\}"
	CalleeSecond(world)
}

func CalleeSecond(world sink.World) { // want CalleeSecond:"example.test/sink.World: Spoons"
	_ = world.Spoons
}

// Test case 4: Complex mutual recursion (3-way cycle)
//
//untangle:print
func CycleA(world sink.World, depth int) { // want CycleA:"example.test/sink.World: Bowls, Ingredients, Knives" "CycleA: World\\{Bowls, Ingredients, Knives\\}"
	if depth <= 0 {
		return
	}
	_ = world.Knives
	CycleB(world, depth-1)
}

func CycleB(world sink.World, depth int) { // want CycleB:"example.test/sink.World: Bowls, Ingredients, Knives"
	if depth <= 0 {
		return
	}
	_ = world.Ingredients
	CycleC(world, depth-1)
}

func CycleC(world sink.World, depth int) { // want CycleC:"example.test/sink.World: Bowls, Ingredients, Knives"
	if depth <= 0 {
		return
	}
	_ = world.Bowls
	CycleA(world, depth-1)
}
