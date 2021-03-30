package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Buf struct {
	data []byte
	len  int
}

func (b *Buf) Length() int {
	return len(b.data)
}
func (b *Buf) Read(p []byte) (int, error) {
	n := copy(p, b.data)
	b.data = b.data[n:]
	if n == 0 {
		return 0, io.EOF
	}
	return n, nil
}

func (b *Buf) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *Buf) String() string {
	return string(b.data)
}

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func AnimalFromJSON(p string) []Animal {
	fd, err := os.Open(p)
	if err != nil {
		fmt.Printf("Error open file %v", err)
	}
	var beasts []Animal
	decoder := json.NewDecoder(fd)
	if err := decoder.Decode(&beasts); err != nil {
		fmt.Printf("Error open decoding %v", err)

	}
	for _, beast := range beasts {

		fmt.Printf("Animal %s %d\n", beast.Name, beast.Age)
	}
	return beasts
}
