package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var crawler = Crawler{
		targets:   "urls.txt",
		requests:  make(chan *PingRequest),
		responses: make(chan *PingResponse),
	}
	go crawler.ReadLinks()
	go crawler.Start()
	for resp := range crawler.responses {
		log.Println(resp.Url, ":", resp.Status)
	}

}

func getUrls() []string {
	args := os.Args
	var path string
	if len(args) > 1 {
		path = args[1]
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(dir, "urls.txt")
	}
	log.Printf("reading from %v", path)
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fd)
	result := make([]string, 0, 100)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
