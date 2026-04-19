// go routines --> lightweight thread. It make function to run concurrent/parallel. If main function get terminates, go routines also get terminated
// To solve this termination problem, we can use channels to communicate between main and go routines, we can use waitgroups to wait for all go routines to finish,
//  and we can use context to cancel go routines, and we can use select to handle multiple channels, and we can use time to sleep for a certain amount of time,
// and we can use sync.Mutex to lock and unlock resources, and we can use sync.RWMutex to lock and unlock resources for read and write operations, and we can use sync.Once to execute a function only once
// and we can use sync.Pool to reuse objects, and we can use sync.Map to store key-value pairs, and we can use sync.Cond to wait for a condition to be met, and we can use sync.WaitGroup to wait for multiple goroutines to finish
// and we can use sync.Pool to reuse objects, and we can use sync.Map to store key-value pairs, and we can use sync.Cond to wait for a condition to be met, and we can use sync.WaitGroup to wait for multiple goroutines to finish


package main

import ("fmt"
	"sync")

func worker (name string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(name, "started working")
	fmt.Println(name, "finished working")
}

func main(){
	var wg sync.WaitGroup

	workers := []string{"worker1", "worker2", "worker3"}
	for _, w := range workers {
		wg.Add(1)
		go worker(w, &wg)
		fmt.Println("Started--", w, " from main")
	}
	wg.Wait()
	fmt.Println("All workers finished working")
	fmt.Println("Main function finished")
}
