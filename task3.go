//go:build ignore

package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	m := make(map[int]int) // создаем мапу
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			mu.Lock() // только одна горутина за раз пишет в map
			m[k] = k * k
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	keys := make([]int, 0, len(m)) // тип, начальная длина, Ёмикость
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, m[k])
	}
}
