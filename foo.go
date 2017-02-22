package main

import (
	"fmt"
  "math"
  "github.com/Shopify/sarama"
)

func foo(a int, b int) (int, int, int) {
  return a, b, a * b
}

func sqrt(x float64) float64 {
  z := float64(5)

  for i := 0; i < 10; i++ {
    z = z - ((z*z - x) / (2 * z))
  }

  return z
}

func main() {
  fmt.Printf("Mine %g, theirs: %g", sqrt(7.), math.Sqrt(7))
  var config = sarama.NewConfig()
  config.Producer.Return.Successes = true
}



