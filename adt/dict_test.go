package adt

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	d := Dict{}
	d["a"] = 10
	d["b"] = 20
	ok, val := d.Contains("a")
	if !ok && val != 10 {
		t.Errorf("Error in contains %v  %v", ok, val)
	}
	val = d.SetDefault("a", nil)
	if val != 10 {
		t.Errorf("Should be 10 %v", val)

	}
	val = d.SetDefault("z", 100)
	{
		if val != 100 {
			t.Errorf("z should be %v", 100)
		}
	}
	for k, v := range d {
		fmt.Printf("key: %v , val %v\n", k, v)
	}
}
