package sink

import "fmt"

//untangle:track
type World struct { // want World:"tracked struct with fields: Knives, Ingredients, Bowls, Spoons"
	Knives      bool
	Ingredients bool
	Bowls       bool
	Spoons      bool
}

func (w *World) GetIngredients() bool { // want GetIngredients:"example.test/sink.World: Ingredients"
	return w.Ingredients
}

//untangle:ignore
func API() World {
	world := World{
		Knives:      true,
		Ingredients: true,
		Bowls:       false,
		Spoons:      true,
	}

	world.Bowls = true
	return world
}

func Chop() { // want Chop:"example.test/sink.World: Ingredients, Knives"
	world := API()
	fmt.Println("Chopping:", world.Knives, world.GetIngredients())
}

func Mix(world World) { // want Mix:"example.test/sink.World: Bowls, Ingredients, Spoons"
	fmt.Println("Mixing:", world.Bowls, world.Spoons, world.GetIngredients())
}
