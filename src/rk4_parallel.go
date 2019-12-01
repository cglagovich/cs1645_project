package main

import (
  "fmt"
  "time"
  "sync"
  "math"
  "os"
  "strconv"
)

const PROBLEM_SIZE = 8192
// Num threads is now a command line argument, default = 1
var NUM_GO int = 1
var y[PROBLEM_SIZE] float64
var yt[PROBLEM_SIZE] float64
var k1[PROBLEM_SIZE] float64
var k2[PROBLEM_SIZE] float64
var k3[PROBLEM_SIZE] float64
var k4[PROBLEM_SIZE] float64
var pow[PROBLEM_SIZE] float64
var yout[PROBLEM_SIZE] float64
var c[PROBLEM_SIZE][PROBLEM_SIZE] float64
var h float64 = 0.3154
var totalSum float64 = 0.0
var counter1 int
var counter2 int
var counter3 int

var m1 sync.Mutex
var m2 sync.Mutex
var m3 sync.Mutex
var m4 sync.Mutex

func main() {
  if(len(os.Args) > 1) {
    NUM_GO, _ = strconv.Atoi(os.Args[1])
  }
  
  var wait_group sync.WaitGroup

  for i:= 0; i < PROBLEM_SIZE; i++ {
    y[i] = float64(i * i)
    pow[i] = float64(i + i)
    for j:= 0; j < PROBLEM_SIZE; j++ {
      c[i][j] = float64(i * i + j)
    }
  }

  startTime := time.Now()
  // The real work starts here
  for i := 0; i < NUM_GO; i++ {
    wait_group.Add(1)
    go func(rank int) {
      defer wait_group.Done()
      rk4_work(rank)
    }(i)
  }

  wait_group.Wait()
  // The real work ends here
  elapsedTime := time.Since(startTime)

  //fmt.Println("totalSum =", totalSum)
  fmt.Println(int64(elapsedTime) )
}

func rk4_work(rank int) {
  chunk_size := int(math.Ceil(float64(PROBLEM_SIZE / NUM_GO)))
  my_first := rank * chunk_size
  my_last := my_first + chunk_size

  for i := my_first; i < my_last && i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*y[j]
    }
    k1[i] = h*(pow[i]-yt[i])
  }
  m1.Lock()
  counter1++
  m1.Unlock()
  for ;counter1 < NUM_GO; {

  }

  for i := my_first; i < my_last && i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+0.5*k1[j])
    }
    k2[i] = h*(pow[i]-yt[i])
  }

  m2.Lock()
  counter2++
  m2.Unlock()
  for ;counter2 < NUM_GO; {

  }

  for i := my_first; i < my_last && i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+0.5*k2[j])
    }
    k3[i] = h*(pow[i]-yt[i])
  }

  m3.Lock()
  counter3++
  m3.Unlock()
  for ;counter3 < NUM_GO; {

  }

  for i := my_first; i < my_last && i < PROBLEM_SIZE; i++ {
    yt[i]=0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+k3[j])
    }
    k4[i] = h*(pow[i]-yt[i])

    yout[i] = y[i] + (k1[i] + 2*k2[i] + 2*k3[i] + k4[i])/6.0
    m4.Lock()
    totalSum+=yout[i]
    m4.Unlock()
  }
}
