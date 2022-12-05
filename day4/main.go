package main

import (
	"bufio"
	"fmt"
	"log"
  "strings"
	"os"
  "strconv"
)

func main() {
	lines := ReadLines("input.txt")
	contained := 0
	overlapped := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		e1, e2 := ToInts(strings.Split(sections[0], "-")), ToInts(strings.Split(sections[1], "-"))
		if (e1[0] >= e2[0] && e1[1] <= e2[1]) || (e1[0] <= e2[0] && e1[1] >= e2[1]) {
			contained += 1
		}
		if (e1[0] >= e2[0] && e1[0] <= e2[1]) || (e1[0] <= e2[0] && e1[1] >= e2[0]) {
			overlapped += 1
		}
	}
	fmt.Println(contained)
	fmt.Println(overlapped)
}


func ToInts(a []string) []int {
	return Map(a, func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	})
}

func Map[T any, S any](a []T, f func(x T) S) []S {
	r := make([]S, 0)
	for _, e := range a {
		r = append(r, f(e))
	}
	return r
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
  if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
  if err != nil {
		log.Fatal(err)
	}
	return lines
}
