package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestBuf(t *testing.T) {
	b := new(Buf)
	fmt.Printf("Len %v\n", b.Length())
	n, err := b.Write([]byte("hello"))
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
	fmt.Printf("N %v\n", n)
	fmt.Printf("Len2 %v\n\n\n", b.Length())
	os.Stdout.Write(b.data)

	AnimalFromJSON("data.json")

}
