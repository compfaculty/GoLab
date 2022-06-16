package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func factorial(n int64) int64 {

	var ret int64 = 1
	for i := 1; n >= int64(i); i++ {
		ret *= int64(i)
		time.Sleep(1)
	}
	return ret
}

type decor func(n int64) int64

func wlogger(f decor, logger *log.Logger) decor {
	return func(n int64) (result int64) {
		defer func(t time.Time) {
			logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n, result)
		}(time.Now())
		//return f(n)
		return f(n)
	}
}

func cached(fn decor, cache *sync.Map) decor {
	return func(n int64) int64 {
		//fn := func(n int64) int64 {
		val, ok := cache.Load(n)
		if ok {
			return val.(int64)
		}
		val = fn(n)
		cache.Store(n, val)
		return val.(int64)
	}
	//return fn(n)
}

//}

func main() {
	f := cached(factorial, &sync.Map{})
	l := wlogger(f, log.New(os.Stdout, "test", 1))
	l(20)
	l(20)
	//var buf bytes.Buffer
	//defer buf.WriteTo(os.Stdout)
	//instream := make(chan int, 3)
	//var l sync.Mutex
	//var wg sync.WaitGroup
	//go func() {
	//	defer close(instream)
	//	defer fmt.Fprint(&buf, "Done")
	//	for i := 0; i < 5; i++ {
	//		wg.Add(1)
	//		go func(k int) {
	//			defer wg.Done()
	//			fmt.Fprintf(&buf, "Send %d\n", k)
	//			l.Lock()
	//			instream <- k*k
	//			l.Unlock()
	//		}(i)
	//	wg.Wait()
	//	}
	//}()
	//for val := range instream {
	//	fmt.Fprintf(&buf, "Recived %d\n", val)
	//}
}
