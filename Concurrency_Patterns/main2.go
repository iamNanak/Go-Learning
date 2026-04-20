package main

import (
	"fmt"
	"sync"
	"time"
)

type Result2 struct {
	Value string
	Err   error
}

func worker2(url string, wg *sync.WaitGroup, resultChan chan Result2) {
	defer wg.Done() // signal that this goroutine is done when the function returns

	time.Sleep(time.Millisecond * 50)

	// write by workers in chan
	resultChan <- Result2{Value: url, Err: nil}
}

func main() {

	// worker pool pattern
	jobs := []string{"image_1.png", "image_2.png", "image_3.png", "image_4.png", "image_5.png", "image_6.png", "image_7.png", "image_8.png", "image_9.png", "image_10.png", "image_11.png", "image_12.png", "image_13.png", "image_14.png", "image_15.png", "image_16.png", "image_17.png", "image_18.png", "image_19.png", "image_20.png"}

	var wg sync.WaitGroup

	resultChan := make(chan Result2, 51)

	for _, job := range jobs {
		wg.Add(1)
		go worker2(job, &wg, resultChan)
	}
	wg.Wait()
	close(resultChan)

	for res := range resultChan {
		fmt.Printf("Received: %s \n", res)
	}
}
