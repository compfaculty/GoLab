package carman

import "testing"

const testVIN = "W09000051T2123456"

func TestManufacturer(t *testing.T) {
	m := Manufacturer(testVIN)
	if m != "WO9123" {
		t.Errorf("Unexpected %v  value for %v", m, testVIN)
	}
}
