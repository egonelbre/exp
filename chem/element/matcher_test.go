package element

import (
	"testing"
)

func BenchmarkCountMatches(t *testing.B) {
	var matcher Matcher
	for i := 0; i < t.N; i++ {
		matcher.Init("amputations")
		matcher.Run()
	}
}
