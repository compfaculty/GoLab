package misc

import (
	"io/ioutil"
	"net/http"
)

func ParseIt(url string, ch chan string) {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ch <- string(body)
}
