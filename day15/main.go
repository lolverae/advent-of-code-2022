package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sensors := map[image.Point]int{}
	line := map[int]struct{}{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var a, b image.Point
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &a.X, &a.Y, &b.X, &b.Y)
		sensors[a] = GetAbsolute(a.X-b.X) + GetAbsolute(a.Y-b.Y)
		d := sensors[a] - GetAbsolute(2000000-a.Y)
		for x := a.X - d; x <= a.X+d; x++ {
			if !(b.X == x && b.Y == 2000000) {
				line[x] = struct{}{}
			}
		}
	}
	Part1(line)
	Part2(sensors)
}

func Part1(line map[int]struct{}) {
	fmt.Println(len(line))
}

func Part2(sensors map[image.Point]int) {
	for y := 0; y <= 4000000; y++ {
		// var sum int
	loop:
		for x := 0; x <= 4000000; x++ {
			for s, d := range sensors {
				if dx, dy := s.X-x, s.Y-y; GetAbsolute(dx)+GetAbsolute(dy) <= d {
					x += d - GetAbsolute(dy) + dx
					continue loop
				}
			}
			fmt.Println(x*4000000 + y)
			return
		}
	}
}

func GetAbsolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
