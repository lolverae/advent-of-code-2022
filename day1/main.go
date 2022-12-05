package main

import (
  "fmt"
  "os"
  "sort"
  "bufio"
  "log"
  "strconv"
)

func main() {
  file,err := os.Open("input.txt")
  if err != nil {
    log.Fatal("Whoops, you've got an error: ", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  elves := make([][]int, 0)
  elf := make([]int, 0)

  for scanner.Scan() {
    s := scanner.Text()
    if s == "" {
      elves = append(elves, elf)
      elf = make([]int, 0)
    } else {
      i, err := strconv.Atoi(s)
      if err != nil{
        log.Fatal("Whoops, you've got an error: ", err)
      }
    elf = append(elf, i)
    }
  }
  fmt.Println(caloriesPerElf(elves))
}

func caloriesPerElf(elves [][]int) (int, int) {
  totalCaloriesPerElf := make([]int, 0)
  for _, elf := range elves {
    sum := 0 
    for _, cal := range elf {
      sum += cal
    }
      totalCaloriesPerElf = append(totalCaloriesPerElf,sum)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(totalCaloriesPerElf))) 
  return totalCaloriesPerElf[0], totalCaloriesPerElf[0]+totalCaloriesPerElf[1]+totalCaloriesPerElf[2]
}

