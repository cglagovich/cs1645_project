package main

import (
  "fmt"
  "time"
  "sync"
  "math"
)

const PROBLEM_SIZE = 134217728
const NUM_GO = 4
var x[PROBLEM_SIZE] float64
var sum[NUM_GO] float64

var c *sync.Cond
var m sync.Mutex
var counter int

func main() {
  var wait_group sync.WaitGroup
  c = sync.NewCond(&m)
  for i:= 0; i < PROBLEM_SIZE; i++ {
    x[i] = float64(i * i)
  }

  startTime := time.Now()

  for i := 0; i < NUM_GO; i++ {
    wait_group.Add(1)
    go func(i int) {
      defer wait_group.Done()
      reduce(i)
    }(i)
  }

  wait_group.Wait()

  elapsedTime := time.Since(startTime)
  fmt.Println("Sum =", sum[0])
  fmt.Println("Interval length:", elapsedTime)

}

func reduce(rank int) {
  chunk_size := int(math.Ceil((float64(PROBLEM_SIZE) / float64(NUM_GO))))
  my_first := rank * chunk_size
  my_last := my_first + chunk_size

  for i := my_first; i < my_last && i < PROBLEM_SIZE; i++ {
    sum[rank] += x[i];
  }
  half := NUM_GO/2
  steps := int(math.Floor(math.Sqrt(NUM_GO)))
  n := NUM_GO

  for i:=0; i<steps; i++ {
    c.L.Lock()
    counter++
    if counter != NUM_GO {
      c.Wait()
    } else {
      counter = 0
      c.Broadcast()
    }
    c.L.Unlock()

    if rank < half {
      sum[rank] += sum[rank+half]
    }
    if n%2 != 0 && rank == 0 {
      sum[0] += sum[n-1]
    }
    half = half/2

  }
}
