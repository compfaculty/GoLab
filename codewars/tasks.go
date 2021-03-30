package codewars

func TwoSum(numbers []int, target int) [2]int {
	ret := [2]int{0, 0}
	data := make(map[int][]int)
	for index, value := range numbers {
		if cap(data[value]) < 2 {
			data[value] = append(data[value], index)
		}
	}

	for i := 0; i < len(numbers); i++ {
		if val, ok := data[target-numbers[i]]; ok {
			if cap(val) > 1 {
				ret[0], ret[1] = i, val[1]
			} else {
				ret[0], ret[1] = i, val[0]
			}
			break
		}
	}
	return ret
}

//https://www.codewars.com/kata/5544c7a5cb454edb3c000047/go
func BouncingBall(h, bounce, window float64) int {
	if (bounce <= 0 || bounce >= 1) || (window >= h) || (h <= 0) {
		return -1
	}
	var count = 1
	for {
		h *= bounce
		if h <= window {
			break
		}
		count += 2
	}
	return count
}
