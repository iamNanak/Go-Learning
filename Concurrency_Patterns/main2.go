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

// func worker2(url string, wg *sync.WaitGroup, resultChan chan Result2) {
// 	defer wg.Done() // signal that this goroutine is done when the function returns

// 	time.Sleep(time.Millisecond * 50)

// 	// write by workers in chan
// 	resultChan <- Result2{Value: url, Err: nil}
// }

func worker3(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result2) {
	defer wg.Done() // signal that this goroutine is done when the function returns

	for job := range jobsChan {
		time.Sleep(time.Millisecond * 50)

		fmt.Printf("Worker processing: %s \n", job)
		resultChan <- Result2{Value: job, Err: nil}
	}

	fmt.Println("Worker exiting")
}

func main() {

	// worker pool pattern
	jobs := []string{"image_1.png", "image_2.png", "image_3.png", "image_4.png", "image_5.png", "image_6.png", "image_7.png", "image_8.png", "image_9.png", "image_10.png", "image_11.png", "image_12.png", "image_13.png", "image_14.png", "image_15.png", "image_16.png", "image_17.png", "image_18.png", "image_19.png", "image_20.png"}

	var wg sync.WaitGroup
	totalWorkers := 5

	start := time.Now()

	resultChan := make(chan Result2, 15)
	jobsChan := make(chan string, len(jobs))

	// start worker goroutines

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go worker3(jobsChan, &wg, resultChan)

	}

	// for _, job := range jobs {
	// 	wg.Add(1)
	// 	go worker2(job, &wg, resultChan)
	// }
	// wg.Wait()
	// close(resultChan)

	// send jobs to the jobs channel
	for _, job := range jobs {
		jobsChan <- job
	}

	close(jobsChan) // close the jobs channel to signal workers that no more jobs will be sent

	// we can wait for the workers to finish in a separate goroutine and close the result channel after all workers are done in case of unbuffered channel or if size of the buffered channel is less than the number of jobs
	go func() {
		wg.Wait()         // wait for all workers to finish
		close(resultChan) // close the result channel after all workers are done
	}()

	for res := range resultChan {
		fmt.Printf("Received: %s \n", res)
	}

	fmt.Printf("Total time taken: %s \n", time.Since(start))
}
