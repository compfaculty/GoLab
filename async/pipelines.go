package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"strings"
	"sync"
)

type ChanMap chan map[string]string

func getFiles(baseDir string) chan string {
	out := make(chan string)
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(fmt.Sprintf("No directory %v %v\n", baseDir, err))
	}
	go func() {
		defer close(out)
		for _, file := range files {
			if !file.IsDir() {
				out <- path.Join(baseDir, file.Name())
			}
		}
	}()
	return out
}
func readFiles(in <-chan string) ChanMap {
	out := make(ChanMap)
	go func() {
		defer close(out)
		for file := range in {
			tmp := make(map[string]string)
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			tmp[file] = string(content)
			out <- tmp
		}
	}()
	return out
}

func producer(data []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range data {
			out <- num
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func findText(text string, ch ChanMap) chan map[string]bool {
	out := make(chan map[string]bool)
	go func() {
		defer close(out)
		for i := range ch {
			tmp := make(map[string]bool)
			for k, v := range i {
				if strings.Contains(v, text) {
					tmp[k] = true
				} else {
					tmp[k] = false
				}
				out <- tmp
			}
		}
	}()
	return out
}

func merge(chs ...<-chan int) (out chan int) {
	var wg sync.WaitGroup
	out = make(chan int)
	output := func(i int, c <-chan int) {
		fmt.Printf("goroutine num: %v\n", i)
		defer wg.Done()
		for i := range c {
			out <- i
		}
	}
	wg.Add(len(chs))
	for i, ch := range chs {
		go output(i, ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return
}
func main() {

	//p := producer(getArr(100))
	//
	//for v := range merge(sq(p), sq(p)) {
	//	fmt.Printf("Val %v\n", v)
	//}
	files := readFiles(getFiles("C:\\Users\\AlexGranovsky\\go\\src\\awesomeProject"))
	for i := range findText("getArr", files) {
		for k, v := range i {
			fmt.Printf("%-80v: %v\n", k, v)
		}
	}
}

func getArr(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}
