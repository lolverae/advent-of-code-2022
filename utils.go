package aoc2022

import (
	"log"
	"os"
	"strconv"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadLines(filename string) []byte {
	file, err := os.ReadFile(filename)
	Check(err)
	return file
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

func GetMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
