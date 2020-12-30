package main

import (
	"fmt"
	"github.com/compfaculty/GoLab/misc"
	"github.com/compfaculty/GoLab/utils"
)

func main() {
	fmt.Println(utils.ReverseRunes("!oG ,olleH"))
	fmt.Println(utils.ReverseRunes("QWERTY"))
	ch := make(chan string)
	go misc.ParseIt("asdfg", ch)
	fmt.Printf("%s", <-ch)

}
