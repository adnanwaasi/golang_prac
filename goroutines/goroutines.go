package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg      = sync.WaitGroup{}
	locker  = sync.RWMutex{}
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
	fmt.Printf("\ntotal time for execution is %v", time.Since(t))
	fmt.Println(results)
}

func readandwrite(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("\nresults from db are ", db[i])
	write(db[i])
	Log()
	wg.Done()
}

func write(s string) {
	locker.Lock()
	results = append(results, s)
	locker.Unlock()
}

func Log() {
	locker.RLock()
	fmt.Printf("\nthe results are %v", results)
	locker.RUnlock()
}
