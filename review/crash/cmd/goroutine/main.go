package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

var wg = sync.WaitGroup{} // Basically a counter that stores how many goroutines have been started and how many are still running. We use Done() when it finished and Add(x) when we add one new goroutine
var mutex = sync.Mutex{}

func main() {
	// concurrency != parallel execution
	// concurrency is when we have tasks being completed by CPU working in both of them one by one, for instance, when task1 is accessing db and waiting for this, task2 gets place to its execution
	// parallelism has more cores to work at many tasks

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from database is:", dbData[i])
	mutex.Lock()
	results = append(results, dbData[i])
	mutex.Unlock()
	wg.Done()
}
