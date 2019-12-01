package main

import (
  "fmt"
  "time"
)

const PROBLEM_SIZE = 134217728
var x[PROBLEM_SIZE] float64
var sum float64

func main() {

  for i:= 0; i < PROBLEM_SIZE; i++ {
    x[i] = float64(i * i)
  }

  startTime := time.Now()
  // The real work starts here
  for i := 0; i < PROBLEM_SIZE; i++ {
    sum += x[i];
  }

  elapsedTime := time.Since(startTime)

  fmt.Println("Sum =", sum)
  fmt.Println("Interval length:", elapsedTime)

}
