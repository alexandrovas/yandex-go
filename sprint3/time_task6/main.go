package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	iterCount := 10
	for i := 1; i <= iterCount; i++ {
		now := <-ticker.C
		fmt.Printf("[%d/%d] Count of seconds from start: %d\n", i, iterCount, int(now.Sub(start).Seconds()))
	}
	ticker.Stop()
}
