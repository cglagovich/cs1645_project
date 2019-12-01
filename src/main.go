// This is based on the concurrency tutorial here:
// https://www.youtube.com/watch?v=LvgVSSpwND8

package main

import (
  "fmt"
  "time"
  "sync"
)

func main() {
  // Use a wait group
  // Wait groups are one way to synchronize
  var wg sync.WaitGroup
  // Channels allow for wait and communication.
  // Odd to use with wait group
  c := make(chan string)

  go func() {
    count("sheep", c)
    wg.Done()
  }()

  msg := <- c
  fmt.Println(msg)

  wg.Wait()
}

func count(thing string, c chan string) {
  for i := 1; i <= 5; i++ {
    c <- thing
    time.Sleep(time.Millisecond * 500)
  }

  // Only sender should close the channel
  close(c)
}
