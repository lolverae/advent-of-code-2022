package main
import (
  "fmt"
  "os"
  "bufio"
  "log"
)

// Theirs
// A is Rock      -> 1
// B is paper     -> 2
// C is scissors  -> 3

// Ours 
// X is Rock      -> 1
// Y is paper     -> 2
// Z is scissors  -> 3

// if they = me -> result = 3
// if they > me -> result = 0
// if they < me -> return = 6

// second part
// Ours 
// X is lose  -> 0
// Y is draw  -> 3
// Z is win   -> 6

func main() {
	lines := ReadLines("input.txt")
	scores1 := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	scores2 := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}

	score1 := 0
  score2 := 0
	for _, line := range lines {
		score1 += scores1[line]
    score2 += scores2[line]
	
  fmt.Println(score1,score2)
  }
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
