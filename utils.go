package aoc2022

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	Check(scanner.Err())
	return lines
}

func ToIntMust(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func MakeAbsolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
