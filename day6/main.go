package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
  Part1()
  Part2()
}

func Part1() {
  lines := ReadLines("input.txt")
	for i := 0; i < len(lines)-4; i++ {
		if IsMarker(lines[i : i+4]) {
			fmt.Println(i+4)
			break
		}
  }
}

func Part2() {
  lines := ReadLines("input.txt")
	for i := 0; i < len(lines)-14; i++ {
		if IsMarker(lines[i : i+14]) {
			fmt.Println(i+14)
			break
		}
  }
}

func IsMarker(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true
}

func ReadLines(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  var lines string
	for scanner.Scan() {
		line := scanner.Text()
		lines = line 
	}
	if err != nil {
		log.Fatal(err)
	}
	return lines
}
