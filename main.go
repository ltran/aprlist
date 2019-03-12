package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names3 := []string{"a", "b", "c",
		"d", "e", "h"}
	names4 := []string{"a", "b", "c",
		"d", "e", "h"}
	names5 := []string{"a", "b", "c",
		"d", "e", "h"}
	names17 := []string{
		"a", "b", "c",
		"d", "e", "f",
		"g", "h", "i",
		"j", "k", "l",
		"m", "n", "o",
		"p", "q",
	}

	Shuffle(names3)
	Shuffle(names4)
	Shuffle(names5)
	Shuffle(names17)

	fmt.Println(grouper(names3))
	fmt.Println(grouper(names4))
	fmt.Println(grouper(names5))
	fmt.Println(grouper(names17))
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
