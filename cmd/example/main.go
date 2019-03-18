package main

import (
	"context"
	"fmt"
	"math"

	"github.com/huttotw/warm-up"
)

func main() {
	fmt.Println("starting...")

	// Define our custom function
	f := func(x float64) float64 {
		return math.Pow(x/10, 2)
	}

	// Create a new limiter that will increase in throughput according to (x/10)^2,
	// with a burst up to 0 tokens.
	l := warmup.NewLimiter(f, 0)
	defer l.Stop()

	i := 0
	for {
		ctx := context.Background()
		l.WaitN(ctx, 1)
		fmt.Println("i", i)
		i++
	}
}
