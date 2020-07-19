package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func swap(i, j int) (x, y int) {
	return j, i
}

func main() {
	names := [6]Vector{
		Vector{
			X: 1,
		},
		Vector{
			X: 2,
		},
		Vector{
			X: 3,
		},
	}
	var lessNames []Vector = names[0:2]
	fmt.Println(lessNames)
}
