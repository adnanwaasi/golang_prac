package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg      = sync.WaitGroup{}
	locker  = sync.Mutex{}
	db      = []string{"waasi", "prajol", "shan10", "salih", "dark"}
	results = []string{}
)

func main() {
	t := time.Now()
	for i := range db {
		wg.Add(1)
		go readandwrite(i)
	}
	wg.Wait()
	fmt.Printf("total time for execution is %v", time.Since(t))
	fmt.Println(results)
}

func readandwrite(i int) {
	/*var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)*/
	fmt.Println("results from db are ", db[i])
	locker.Lock()
	results = append(results, db[i])
	locker.Unlock()
	wg.Done()
}
