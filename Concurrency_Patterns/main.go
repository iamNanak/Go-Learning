package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done() // signal that this goroutine is done when the function returns

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("Downloaded / Processed: %s \n", url)

	// write by workers in chan
	resultChan <- url
}

func main() {
	// decration of wait group
	var wg sync.WaitGroup

	// create a buffered channel to hold the results
	resultChan := make(chan string, 5)

	// add the number of goroutines to wait for
	wg.Add(5)

	startTime := time.Now()

	// start the goroutines
	go worker("image_1.png", &wg, resultChan)
	go worker("image_2.png", &wg, resultChan)
	go worker("image_3.png", &wg, resultChan)
	go worker("image_4.png", &wg, resultChan)
	go worker("image_5.png", &wg, resultChan)

	// wait for all goroutines to finish
	wg.Wait()
	// close the channel after all workers are done
	close(resultChan)

	// reading the channel
	for res := range resultChan {
		fmt.Printf("Recevied: %s \n", res)
	}

	fmt.Printf("It tooks %s ms. \n", time.Since(startTime))

}
