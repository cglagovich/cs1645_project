package main

import (
  "fmt"
  "time"
)

const NROW = 1024
const NCOL = NROW

func main() {
  var inputArrayA  [NROW][NCOL]int
  var inputArrayB  [NROW][NCOL]int
  var outputArrayC [NROW][NCOL]int

  // Initialize Arrays
  for i := 0; i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      inputArrayA[i][j] = i * NCOL + j
      inputArrayB[i][j] = j * NCOL + i
      outputArrayC[i][j] = 0
    }
  }

  startTime := time.Now()

  for i := 0; i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      temp := 0
      for k := 0; k < NROW; k++ {
        temp += inputArrayA[i][k] * inputArrayB[k][j]
      }
      // Keep temp in a register, access this part of the matrix less
      outputArrayC[i][j] = temp
    }
  }

  elapsedTime := time.Since(startTime)

  // Verify result
  totalSum := 0.0
  for i := 0; i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      totalSum += float64(outputArrayC[i][j])
    }
  }

  fmt.Println("Total sum =", totalSum)
  fmt.Printf("Intervel length: %s\n", elapsedTime)
}
