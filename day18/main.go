package main

import (
	"fmt"
	"math"
	// "os"
	"strings"
  "lolverae/aoc2022"
)

type Obsidian struct {
	X, Y, Z int
}

func (p Obsidian) Add(q Obsidian) Obsidian {
	return Obsidian{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

func main() {
  input := aoc2022.ReadLines("input.txt")
	lava := map[Obsidian]struct{}{}
	min := Obsidian{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Obsidian{math.MinInt, math.MinInt, math.MinInt}
  for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var p Obsidian
		fmt.Sscanf(s, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava[p] = struct{}{}

		min = Obsidian{aoc2022.GetMin(min.X, p.X), aoc2022.GetMin(min.Y, p.Y), aoc2022.GetMin(min.Z, p.Z)}
		max = Obsidian{aoc2022.GetMax(max.X, p.X), aoc2022.GetMax(max.Y, p.Y), aoc2022.GetMax(max.Z, p.Z)}
	}
	min = min.Add(Obsidian{-1, -1, -1})
	max = max.Add(Obsidian{1, 1, 1})

	delta := []Obsidian{
		{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}

	part1 := 0
  for p := range lava {
		for _, d := range delta {
			if _, ok := lava[p.Add(d)]; !ok {
				part1++
			}
		}
	}
	fmt.Println(part1)
}

