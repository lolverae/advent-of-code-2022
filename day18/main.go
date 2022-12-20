package main

import (
	"fmt"
	"lolverae/aoc2022"
	"math"
	"strings"
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
	delta := []Obsidian{
		{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}
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

	surfaceArea := 0
	for p := range lava {
		for _, d := range delta {
			if _, ok := lava[p.Add(d)]; !ok {
				surfaceArea++
			}
		}
	}
	fmt.Println(surfaceArea)

	queue := []Obsidian{min}
	visited := map[Obsidian]struct{}{min: {}}
	externalSurfaceArea := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, d := range delta {
			next := current.Add(d)

			if _, isOk := lava[next]; isOk {
				externalSurfaceArea++
			} else if _, isOk := visited[next]; !isOk &&
				next.X >= min.X && next.X <= max.X &&
				next.Y >= min.Y && next.Y <= max.Y &&
				next.Z >= min.Z && next.Z <= max.Z {
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}
	fmt.Println(externalSurfaceArea)
}
