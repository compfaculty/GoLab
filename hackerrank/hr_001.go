package main

import (
	"fmt"
)

func minimumDistances(a []int32) int32 {

	data := make(map[int32][]int32)
	for ind, item := range a {
		if _, ok := data[item]; ok {
			data[item] = append(data[item], int32(ind))
		} else {
			data[item] = []int32{int32(ind)}
		}
	}
	var ret int32 = -1
	//var min int32
	for _, val := range data {
		if len(val) >= 2 {
			localmin := val[1] - val[0]
			for i := 1; i < len(val) - 1 ; i++ {
				if val[i + 1] - val[i] < localmin {
					localmin = val[i+1] - val[i]
				}
			}
			if ret == -1 {
				ret = localmin
			} else {
				if localmin < ret {
					ret = localmin
				}
			}

		}
	}

	return ret
}
func main() {
	ret := minimumDistances([]int32{7, 1, 3, 4, 1, 7, 7})
	fmt.Printf("RET: %v", ret)
}
