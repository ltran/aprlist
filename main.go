package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names3 := []string{"a", "b", "c",
		"d", "e", "h"}
	fmt.Println(grouper(Shuffle(names3)))

	names4 := []string{"a", "b", "c",
		"d", "e", "h"}
	fmt.Println(grouper(Shuffle(names4)))

	names5 := []string{"a", "b", "c",
		"d", "e", "h"}
	fmt.Println(grouper(Shuffle(names5)))

	names17 := []string{
		"a", "b", "c",
		"d", "e", "f",
		"g", "h", "i",
		"j", "k", "l",
		"m", "n", "o",
		"p", "q",
	}
	fmt.Println(grouper(Shuffle(names17)))
}

func grouper(names []string) [][]string {
	if len(names) >= 3 && len(names) <= 5 {
		return [][]string{names}
	}

	return append(grouper(names[3:]), names[0:3])
}

func Shuffle(vals []string) []string {
	return vals
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
