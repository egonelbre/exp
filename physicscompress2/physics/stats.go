package physics

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func summarize(vs []float64) {
	fmt.Printf("%d %.3f ±%.3f\n", len(vs), stats.Mean(vs), stats.StdDevS(vs))
}
