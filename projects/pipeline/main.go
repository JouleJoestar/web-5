package main

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	var prevValue string
	for value := range inputStream {
		if value != prevValue {
			outputStream <- value
			prevValue = value
		}
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go func() {
		inputStream <- "apple"
		inputStream <- "banana"
		inputStream <- "banana"
		inputStream <- "cherry"
		inputStream <- "apple"
		inputStream <- "date"
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for value := range outputStream {
		fmt.Println(value)
	}
}
