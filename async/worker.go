package main

import (
	"fmt"
	"sync"
)

const N = 5
func worker(index int, wg *sync.WaitGroup, cond *sync.Cond) {
	var vgr sync.WaitGroup
	wg.Done()
	vgr.Add(1)
	go func(i int) {
		vgr.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fmt.Println("Worker", i)
	}(index)
	vgr.Wait()
}
func main() {
	fmt.Print("Hello")
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		worker(i,&wg,cond)
	}
	//time.Sleep(30)
	cond.Broadcast()
	wg.Wait()

}
