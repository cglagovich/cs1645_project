package main

import (
  "fmt"
  "time"
  "math"
  "sync"
  "strconv"
  "os"
)

var NUM_GO int = 1

const NROW = 1024
const NCOL = NROW

var inputArrayA  [NROW][NCOL]int
var inputArrayB  [NROW][NCOL]int
var outputArrayC [NROW][NCOL]int

func main() {
  if(len(os.Args) > 1) {
    NUM_GO, _ = strconv.Atoi(os.Args[1])
  }

  var wait_group sync.WaitGroup

  // Initialize Arrays
  for i := 0; i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      inputArrayA[i][j] = i * NCOL + j
      inputArrayB[i][j] = j * NCOL + i
      outputArrayC[i][j] = 0
    }
  }

  startTime := time.Now()


  for i := 0; i < NUM_GO; i++ {
    wait_group.Add(1)
    go func(i int) {
      defer wait_group.Done()
      mmult(i)
    }(i)
  }

  wait_group.Wait()

  elapsedTime := time.Since(startTime)

  // Verify result
  totalSum := 0.0
  for i := 0; i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      totalSum += float64(outputArrayC[i][j])
    }
  }

  //fmt.Println("Total sum =", totalSum)
  elapsed_ns := int64(elapsedTime)
  fmt.Println(elapsed_ns)
}

func mmult(rank int) {
  chunk_size := int(math.Ceil(float64(NROW) / float64(NUM_GO)))
  my_first := rank * chunk_size
  my_last := my_first + chunk_size
  for i := my_first; i < my_last && i < NROW; i++ {
    for j := 0; j < NCOL; j++ {
      temp := 0
      for k := 0; k < NROW; k++ {
        temp += inputArrayA[i][k] * inputArrayB[k][j]
      }
      // Keep temp in a register, access output matrix less
      outputArrayC[i][j] = temp
    }
  }
}
