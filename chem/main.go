// Spelling with Chemistry
// see:
//   https://www.reddit.com/r/dailyprogrammer/comments/5seexn/20170206_challenge_302_easy_spelling_with/
//   https://www.reddit.com/r/programming/comments/6bgxia/using_python_to_find_the_longest_word_spellable/
//
// This demonstrates using a generated parser from trie.
//

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/egonelbre/exp/chem/element"
)

func main() {
	var matcher element.Matcher
	var src io.Reader
	src = os.Stdin
	if len(os.Args) > 1 {
		src, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewScanner(src)
	start := time.Now()
	total := 0
	for scanner.Scan() {
		matcher.Init(scanner.Text())
		matcher.Run()
		total++
	}
	stop := time.Now()
	fmt.Println(stop.Sub(start) / time.Duration(total))
}
