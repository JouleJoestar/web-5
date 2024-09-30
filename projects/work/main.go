package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			work()
		}()
	}

	wg.Wait()
	fmt.Println("Все работы завершены") //для проверки, на степике строчку убрать
}
