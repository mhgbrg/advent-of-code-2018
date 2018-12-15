package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()

  fmt.Println(time.Since(start))
}
