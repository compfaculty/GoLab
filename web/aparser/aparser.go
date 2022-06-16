package aparser

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"sync"
)

// pageLinks will recursively scan a `html.Node` and will return
// a list of links found, with no duplicates
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
					//log.Printf("LINKS %v\n", links)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//log.Printf("LINK: %v\n", c)
		links = pageLinks(links, c)
	}
	return links
}

// parse given a string pointing to a URL will fetch and parse it
// returning an html.Node pointer
func parse(url string) (*html.Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}
	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse page")
	}
	return b, err
}

func analyze(url, baseurl string) (out chan string) {
	out = make(chan string)
	visited := make(map[string]string)
	var wg sync.WaitGroup
	var ana func(url, baseurl string, visited *map[string]string, out chan string)
	ana = func(url, baseurl string, visited *map[string]string, out chan string) {
		wg.Done()
		page, err := parse(url)
		if err != nil {
			fmt.Printf("Error getting page %s %s\n", url, err)
			return
		}
		//fmt.Printf("WORK ON  %v : %v\n", baseurl, url)

		title := strings.TrimSpace(pageTitle(page))
		(*visited)[url] = title
		out <- title

		//recursively find links
		links := pageLinks(nil, page)
		for _, link := range links {
			if _, ok := (*visited)[link]; !ok {
				if strings.HasPrefix(link, baseurl) {
					//fmt.Printf("RECURSIVE  %v : %v\n", i, link)
					wg.Add(1)
					go ana(link, baseurl, visited, out)
				}
			}
		}
	}
	go ana(url, baseurl, &visited, out)
	return out
}

// pageTitle given a reference to a html.Node, scans it until it
// finds the title tag, and returns its value
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// sliceContains returns true if `slice` contains `value`
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// parse given a string pointing to a URL will fetch and parse it
// returning an html.Node pointer
//func parse(url string) (*html.Node, error) {
//	r, err := http.Get(url)
//	if err != nil {
//		return nil, fmt.Errorf("Cannot get page")
//	}
//	b, err := html.Parse(r.Body)
//	if err != nil {
//		return nil, fmt.Errorf("Cannot parse page")
//	}
//	return b, err
//}
// checkDuplicates scans the visited map for pages with duplicate titles
// and writes a report
func checkDuplicates(visited *map[string]string) {
	found := false
	uniques := map[string]string{}
	fmt.Printf("\nChecking duplicates..\n")
	for link, title := range *visited {
		if uniques[title] == "" {
			uniques[title] = link
		} else {
			found = true
			fmt.Printf("Duplicate title \"%s\" in %s but already found in %s\n", title, link, uniques[title])
		}
	}

	if !found {
		fmt.Println("No duplicates were found 😇")
	}
}

func PrintResults(url string) {
	//url := "https://alterslash.org/"
	resultsCh := analyze(url, url)
	for data := range resultsCh {
		fmt.Printf("%v\n", data)
	}
}
