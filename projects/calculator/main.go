package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output) // Отложенная функция

		for {
			select {
			case x := <-firstChan:
				output <- x * x
			case x := <-secondChan:
				output <- x * 3
			case <-stopChan:
				return // надо завершить только после получения сигнала остановки
			}
		}
	}()

	return output
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	outputChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		time.Sleep(5 * time.Second)
		close(stopChan)
	}()

	// горутинка, которая будет отправлять данные в каналы
	go func() {
		firstChan <- 5
		secondChan <- 10
		firstChan <- 21
		secondChan <- 13
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
