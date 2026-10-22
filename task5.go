//go:build ignore

package main

import (
	"fmt"
	"image"
	_ "image/gif"  // регистрируют декодеры форматов
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"sync"
)

const workers = 4 // фиксированное число рабочих горутин

// processFile возвращает готовый HTML-тег или "" если файл нужно пропустить
func processFile(path string) string {
	info, err := os.Stat(path)
	if err != nil || !info.Mode().IsRegular() {
		return "" // не существует или не обычный файл
	}
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f) // читает только размеры
	if err != nil {
		return "" // не изображение
	}
	return fmt.Sprintf(`<img src="%s" width="%d" height="%d" />`,
		filepath.Base(path), cfg.Width, cfg.Height)
}

func main() {
	files := os.Args[1:]

	jobs := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range jobs {
				if tag := processFile(path); tag != "" {
					results <- tag
				}
			}
		}()
	}

	// раздаём файлы воркерам
	go func() {
		for _, p := range files {
			jobs <- p
		}
		close(jobs)
	}()

	// закрываем results, когда все воркеры закончили
	go func() {
		wg.Wait()
		close(results)
	}()

	// печатает только main, поэтому строки не перемешиваются
	for tag := range results {
		fmt.Println(tag)
	}
}
