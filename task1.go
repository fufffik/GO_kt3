//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	results := make([]int, len(nums)) //срез

	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func(idx, val int) {
			defer wg.Done()
			results[idx] = val * val // каждая горутина пишет в свою ячейку
		}(i, n)
	}
	wg.Wait()

	for i, n := range nums {
		fmt.Printf("%d^2 = %d\n", n, results[i])
	}
}
