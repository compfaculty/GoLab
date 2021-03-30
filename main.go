package main

import (
	"bufio"
	"flag"
	"fmt"

	// "github.com/compfaculty/GoLab/misc"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func start() {
	flag.Parse()
	scanner := newScanner(flag.Args())
	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", text)
}

func newScanner(args []string) *bufio.Scanner {
	if len(args) == 0 {
		return bufio.NewScanner(os.Stdin)
	}
	path := args[0]
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(fd)
}
func ReadByLines(p string, lines chan string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fd, err := os.Open(filepath.Join(dir, p))

	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	//scanner.Split()

	for scanner.Scan() {
		//log.Printf("String %v", out)
		lines <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	close(lines)
}
func main() {
	// misc.Ping("https://ukr.net")
	data := make(chan string)
	output := make(chan string)
	go ReadByLines("main.go", data)
	for s := range data {
		go func(s string) {
			output <- strings.ToUpper(s)
		}(s)
	}
	for val := range output {
		fmt.Printf("%s\n", val)
	}
	// close(data)
	close(output)
	// start()
	//	fmt.Println(utils.ReverseRunes("!oG ,olleH"))
	//	fmt.Println(utils.ReverseRunes("QWERTY"))
	//	ch := make(chan string)
	//	go misc.ParseIt("asdfg", ch)
	//	fmt.Printf("%s", <-ch)
	//	misc.RunCommand("Get-Alias")
	//	misc.ShowEnv()
	//	misc.ShowProcessInfo(11252)
}
