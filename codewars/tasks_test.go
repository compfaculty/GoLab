package codewars

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		arr  []int
		sum  int
		want [2]int
	}{
		{[]int{1, 2, 3}, 4, [2]int{0, 2}},
		{[]int{1234, 5678, 9012}, 14690, [2]int{1, 2}},
		{[]int{2, 2, 3}, 4, [2]int{0, 1}},
	}
	for _, c := range cases {
		got := TwoSum(c.arr, c.sum)
		if got != c.want {
			t.Errorf("Error Got %v  Should %v", got, c.want)
		}
	}

}
func TestBouncingBall(t *testing.T) {
	cases := []struct {
		h      float64
		bounce float64
		window float64
		want   int
	}{
		{3, 0.66, 1.5, 3},
		{40, 0.4, 10, 3},
		{10, 0.6, 10, -1},
		{40, 1, 10, -1},
		{5, -1, 1.5, -1},
	}
	for _, c := range cases {
		ret := BouncingBall(c.h, c.bounce, c.window)
		if c.want != ret {
			t.Errorf("Error Got %v  Should %v", ret, c.want)
		}
	}
}
