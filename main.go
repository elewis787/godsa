package main

import (
	"github.com/elewis787/godsa/ds/lists"
)

func main() {
	l := lists.Initialize(-1, -4, 1, 4, 6, 4, 7, 3, 1)
	lists.SortList(l, l.Len())
	lists.Visualize(l)
}
