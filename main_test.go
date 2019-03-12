package main

import (
	"reflect"
	"testing"
)

func TestGroupies(t *testing.T) {
	res := groupies([]string{"a", "b", "c", "d", "e"})
	if !reflect.DeepEqual(res, [][]string{{"a", "b", "c", "d", "e"}}) {
		t.Fatal("bad output")
	}
}
