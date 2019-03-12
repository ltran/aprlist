package main

import (
	"fmt"
)

func main() {
	names := []string{"a", "b", "c",
		"d", "e"}
	fmt.Println(groupies(names))
}

func groupies(names []string) [][]string {
	newNames := Shuffle(names)
	groups := [][]string{}

	// 17
	// 3
	// 4
	// 5
	size := len(names)

	group := []string{}
	for _, name := range newNames {
		group = append(group, name)
	}

	_ = groups

	// if len(names) >= 3 || len(names) <= 5 {
	// 	return [][]string{names}
	// }

	// //  {"a"}
	// //	{"a", "b", "c", "d", "e"}

	// len(names) / 3
	// len(names) / 4
	// len(names) / 5
	// // 1 R2

	// for _, n := range names {

	// }

	return [][]string{}
}

func grouper(names []string) [][]string {
	k := 5
	if len(names) >= 3 || len(names) <= 5 {
		return [][]string{names}
	}

	return append(grouper(names[k:]), names[0:k])
}

// func Shuffle(vals []string) []string {
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	ret := make([]string, len(vals))
// 	perm := r.Perm(len(vals))
// 	for i, randIndex := range perm {
// 		ret[i] = vals[randIndex]
// 	}
// 	return ret
// }
