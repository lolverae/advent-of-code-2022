package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := ReadLines("input.txt")
	sum := 0
	for _, line := range lines {
		l := len(line)
		r1, r2 := line[:l/2], line[l/2:]
		t := FindType(r1, r2)
		sum += FindPriority(t)
	}
	fmt.Println(sum)
	sum = 0
	for i := 0; i < len(lines); i += 3 {
		r1, r2, r3 := lines[i+0], lines[i+1], lines[i+2]
		b := FindBadge(r1, r2, r3)
		sum += FindPriority(b)
	}
	fmt.Println(sum)
}

func FindPriority(p byte) int {
	if p >= 97 {
		return (int(p) - 96)
	} else {
		return (int(p) - 38)
	}
}

func FindType(s1, s2 string) byte {
	for _, a1 := range []byte(s1) {
		for _, a2 := range []byte(s2) {
			if a1 == a2 {
				return a1
			}
		}
	}
	return '.'
}

func FindBadge(s1, s2, s3 string) byte {
	for _, a1 := range []byte(s1) {
		for _, a2 := range []byte(s2) {
			for _, a3 := range []byte(s3) {
				if a1 == a2 && a1 == a3 {
					return a1
				}
			}
		}
	}
	return '.'
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
