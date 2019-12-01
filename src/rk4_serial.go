package main

import (
  "fmt"
  "time"
)

const PROBLEM_SIZE = 8192
var y[PROBLEM_SIZE] float64
var yt[PROBLEM_SIZE] float64
var k1[PROBLEM_SIZE] float64
var k2[PROBLEM_SIZE] float64
var k3[PROBLEM_SIZE] float64
var k4[PROBLEM_SIZE] float64
var pow[PROBLEM_SIZE] float64
var yout[PROBLEM_SIZE] float64
var c[PROBLEM_SIZE][PROBLEM_SIZE] float64

func main() {
  var h float64 = 0.3154
  var totalSum float64 = 0.0

  for i:= 0; i < PROBLEM_SIZE; i++ {
    y[i] = float64(i * i)
    pow[i] = float64(i + i)
    for j:= 0; j < PROBLEM_SIZE; j++ {
      c[i][j] = float64(i * i + j)
    }
  }

  startTime := time.Now()
  // The real work starts here
  for i := 0; i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*y[j]
    }
    k1[i] = h*(pow[i]-yt[i])
  }

  for i := 0; i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+0.5*k1[j])
    }
    k2[i] = h*(pow[i]-yt[i])
  }

  for i := 0; i < PROBLEM_SIZE; i++ {
    yt[i] = 0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+0.5*k2[j])
    }
    k3[i] = h*(pow[i]-yt[i])
  }

  for i := 0; i < PROBLEM_SIZE; i++ {
    yt[i]=0.0
    for j := 0; j < PROBLEM_SIZE; j++ {
      yt[i] += c[i][j]*(y[j]+k3[j])
    }
    k4[i] = h*(pow[i]-yt[i])

    yout[i] = y[i] + (k1[i] + 2*k2[i] + 2*k3[i] + k4[i])/6.0
    totalSum+=yout[i]
  }
  // The real work ends here
  elapsedTime := time.Since(startTime)

  fmt.Println("totalSum =", totalSum)
  fmt.Println("Interval length:", elapsedTime)

}
