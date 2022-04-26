package main

import (
	"Jumble/ent"
	"fmt"
)

func main() {
	jumbleArray := [][]string{{"A", "B", "C"}, {"X", "Y", "Z"}}
	str := "CZ"
	x, y := FindJumble(jumbleArray, str)
	fmt.Println(x, y)
}

func FindJumble(jumbleArray [][]string, str string) (int, int) {
	j := ent.NewJumble(jumbleArray)
	return j.Find(str)
}
