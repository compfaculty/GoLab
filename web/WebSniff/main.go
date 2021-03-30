package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	fd, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	count := 0
	for scanner.Scan() {
		count++
		url := scanner.Text()
		fmt.Printf("Processing %s\n", url)
		go fetch(url, ch)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%v\n", <-ch)
	}

	fmt.Printf("Elapsed %v\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan string) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}

	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%q failed", url)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error : %q %q", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
