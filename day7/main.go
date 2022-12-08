package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

  fs := map[string]int{}
  cd := ""
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var size int
		var name string

		if strings.HasPrefix(s, "$ cd") {
			cd = path.Join(cd, strings.Fields(s)[2])
		} else if _, err := fmt.Sscanf(s, "%d %s", &size, &name); err == nil {
			for d := cd; d != "/"; d = path.Dir(d) {
				fs[d] += size
			}
			fs["/"] += size
		}
	}
  fmt.Println(Part1(fs))
  fmt.Println(Part2(fs))
}

func Part1 (fs map[string]int) int{
	part1 := 0
	for _, s := range fs {
		if s <= 100000 {
			part1 += s
		}
	}
  return part1
}

func Part2 (fs map[string]int) int{
	part2 := fs["/"]
	for _, s := range fs {
		if s+70000000-fs["/"] >= 30000000 && s < part2 {
			part2 = s
		}
	}
  return part2
}
