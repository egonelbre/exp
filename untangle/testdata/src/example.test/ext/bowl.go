package ext

import (
	"fmt"

	"example.test/callback"
	"example.test/sink"
)

//untangle:print
func CleanBowl(world sink.World) { // want CleanBowl:"example.test/sink.World: Bowls" "CleanBowl: World\\{Bowls\\}"
	callback.Do(func() {
		fmt.Println("Clean Bowl", world.Bowls)
	})
}
