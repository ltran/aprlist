package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// read a csv line from stdin
	r := csv.NewReader(os.Stdin)
	names, err := r.Read()
	if err != nil {
		panic("failed to read")
	}

	Shuffle(names)

	fmt.Println(grouper(names))
}

func grouper(names []string) [][]string {
	if len(names) >= 3 && len(names) <= 5 {
		return [][]string{names}
	}

	return append(grouper(names[3:]), names[0:3])
}

func Shuffle(vals []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
}

func grouperIter(names []string) [][]string {
	groups := [][]string{}
	for len(names) > 5 {
		groups = append(groups, names[0:3])
		names = names[3:]
	}

	// handle leftovers
	if len(names) != 0 {
		groups = append(groups, names)
	}

	return groups
}
