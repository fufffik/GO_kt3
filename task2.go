//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squares := make(chan int, len(nums)) // буферизированный канал

	var wg sync.WaitGroup
	for _, n := range nums {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			squares <- val * val // горутина отправляет свой квадрат в канал
		}(n)
	}

	wg.Wait()
	close(squares) // закрываем, чтобы range по каналу закончился

	sum := 0
	for sq := range squares {
		sum += sq
	}
	fmt.Println("Сумма квадратов:", sum) // 220
}
