package utils

import (
	"testing"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"Nirvana", "anavriN"},
		{"Abba", "abbA"},
		{"metallica", "acillatem"},
	}
	for _, c := range cases {
		got := ReverseRunes(c.s)
		if got != c.want {
			t.Errorf("Error Got %v  Should %v", got, c.want)
		}
	}

}
