package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type PingResponse struct {
	Url    string
	Status string
	Body   string
}

type PingRequest struct {
	Url string
}

type Crawler struct {
	targets   string
	requests  chan *PingRequest
	responses chan *PingResponse
}

func (p *Crawler) Start() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//go readLinks("urls.txt", p.requests)
	for url := range p.requests {
		go func(uri *PingRequest) {

			resp, err := client.Get(uri.Url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			scanner.Split(bufio.ScanRunes)

			sb := strings.Builder{}
			sb.Grow(512)

			for scanner.Scan() {
				_, _ = fmt.Fprintf(&sb, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			p.responses <- &PingResponse{
				Url:    uri.Url,
				Status: resp.Status,
				Body:   sb.String(),
			}

		}(url)
		//close(p.responses)
	}

	//close(p.requests)
	//close(p.responses)
	return

}

func (p *Crawler) ReadLinks() {
	file, err := os.Open(p.targets)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.requests <- &PingRequest{scanner.Text()}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//close(urls)
}
