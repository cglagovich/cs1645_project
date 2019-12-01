package main

import (
  "fmt"
  "errors"
  "math"
)

type person struct {
  name string
  age int
}

func main() {
  // Code based on:
  // https://www.youtube.com/watch?v=C8LgvuEBraI
  fmt.Println("Hello World")

  // Variables. They have default value. 0 for ints
  fmt.Println("\n-------variables-------")
  var x int
  var y int = 7
  z := 5
  var sum int
  sum = x + y + z
  fmt.Println(sum)

  // Arrays
  fmt.Println("\n-------arrays-------")
  var a [5]int
  b := [5]int{5, 4, 3, 2, 1}
  a[2] = 7
  fmt.Println(a)
  fmt.Println(b)

  // Can also make a slice of ints.
  fmt.Println("\n-------slices-------")
  c := []int{1, 2, 3, 4, 5}
  // Still backed by arrays so go creates a new one behind the scenes
  d := append(c, 13)
  fmt.Println(d)

  // Maps are like dictionaries
  // map["type of keys"]"type of values"
  fmt.Println("\n-------maps-------")
  vertices := make(map[string]int)
  vertices["triangle"] = 2
  vertices["square"] = 3
  vertices["dodecagon"] = 12
  fmt.Println(vertices)
  delete(vertices, "square")
  fmt.Println(vertices["triangle"])

  // Only type of loop in go is for loop
  fmt.Println("\n-------loops-------")
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
  // Can also make it act like a while loop
  var i int
  for i < 5 {
    fmt.Println(i)
    i++
  }
  // Can also iterate over tan array
  arr := []string{"a","b","c"}
  for index, value := range arr {
    fmt.Println("index", index, "value", value)
  }
  // Or with a map
  m := make(map[string]string)
  m["a"] = "alpha"
  m["b"] = "beta"
  for key, value := range m {
    fmt.Println("key:", key, "value:", value)
  }

  // Funciton
  fmt.Println("\n-------functions-------")
  // functions can return multiple things of different types
  // Go does not have exceptions
  result, err := sqrt(16)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }

  // Structs
  fmt.Println("\n-------Structs-------")
  p := person{name: "Bob", age: 23}
  fmt.Println(p.age)

  // Pointers
  fmt.Println("\n-------Pointers-------")
  n := 7
  inc(&n)
  fmt.Println(n)

}

// Make a funcion like this
func sum(x int, y int) int {
  return x + y
}

// We can have multiple return types
func sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Undefined for negative numbers")
  }

  return math.Sqrt(x), nil
}

func inc(x *int) {
  *x++
}
