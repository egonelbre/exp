// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/egonelbre/exp/bit"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var byteCount = [256]int{}

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	for i := range byteCount {
		byteCount[i] = bit.Count(uint64(i))
	}

	data, err := ioutil.ReadFile(filename)
	check(err)

	bitcount := 0
	for _, b := range data {
		bitcount += byteCount[b]
	}
	fmt.Printf("bit bias: %.4f\n", float64(bitcount)/float64(len(data)*8))

	r := bit.NewReader(bytes.NewReader(data))
	circle := byte(0)

	count2 := make(map[byte]int)
	count3 := make(map[byte]int)
	count4 := make(map[byte]int)
	count5 := make(map[byte]int)
	count6 := make(map[byte]int)
	count7 := make(map[byte]int)
	count8 := make(map[byte]int)

	var v int
	for r.Error() == nil {
		v = r.ReadBit()
		circle <<= 1
		circle |= byte(v)

		count2[circle&(1<<2-1)] += 1
		count3[circle&(1<<3-1)] += 1
		count4[circle&(1<<4-1)] += 1
		count5[circle&(1<<5-1)] += 1
		count6[circle&(1<<6-1)] += 1
		count7[circle&(1<<7-1)] += 1
		count8[circle&(1<<8-1)] += 1
	}

	PrintFrequency("2", count2, data)
	PrintFrequency("3", count3, data)
	PrintFrequency("4", count4, data)
	PrintFrequency("5", count5, data)
	PrintFrequency("6", count6, data)
	PrintFrequency("7", count7, data)
	PrintFrequency("8", count8, data)
}

type Freq struct {
	Pattern byte
	Value   float64
}

type Table []Freq

func (s Table) Len() int           { return len(s) }
func (s Table) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Table) Less(i, j int) bool { return s[i].Value > s[j].Value }

func PrintFrequency(name string, freq map[byte]int, data []byte) {
	fmt.Println("F ", name)

	table := make(Table, 0, len(freq))
	for pat, count := range freq {
		entry := Freq{pat, float64(count) / float64(len(data)*8)}
		table = append(table, entry)
	}
	sort.Sort(table)

	for _, freq := range table {
		fmt.Printf("   %8b: %.5f\n", freq.Pattern, freq.Value)
	}
}
