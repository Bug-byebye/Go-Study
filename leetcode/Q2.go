package leetcode

import (
	"fmt"
	"strconv"
)

func gSwaps(num int) []int {
	snum := strconv.Itoa(num)
	n := len(snum)
	var swaps []int

	if n == 1 {
		swaps = append(swaps, num*10)
	}


	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			rnum := []rune(snum)

			rnum[i], rnum[j] = rnum[j], rnum[i]

			swappedStr := string(rnum)

			for swappedStr[0] == '0' {
				swappedStr = swappedStr[1:]
			}

			if swappedNum, err := strconv.Atoi(string(swappedStr)); err == nil && swappedStr!=snum {
				swaps = append(swaps, swappedNum)
			}
		}
	}

	swaps = append(swaps, num)
	return swaps
}

func ApproximatePairs(nums []int) int {

	numMap := make(map[int]int)
	count := 0

	for _, num := range nums {
		swaps := gSwaps(num)
		fmt.Print(swaps)
		for _, swap := range swaps {
			count += numMap[swap]
		}
		numMap[num]++
	}

	return count
}

// [5,12,8,5,5,1,20,3,10,10,5,5,5,5,1] 27