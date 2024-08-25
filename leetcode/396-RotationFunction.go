package leetcode

import (
	"fmt"
	"slices"
)

func MaxRotateFunction(nums []int) int {
	maxRF := 0 
	newnums := make([]int, len(nums))
	copy(newnums, nums)
	for i := 0; i < len(nums); i++ {

		fmt.Print(i,"   ")

		slices.Reverse(newnums)
		slices.Reverse(newnums[:i])
		slices.Reverse(newnums[i:])
		 f := 0
		for ind,val := range newnums{
			f = f+ind*val
		}
		fmt.Println(f)
		if i==0||f>maxRF{
			maxRF=f
		}
		//newnums = nums
		copy(newnums, nums)
	}
	return maxRF
}

func MaxRotateFunctionPro(nums []int) int {
	n := len(nums)
	sum := 0
	f := 0

	// 计算初始的 f(0) 值和所有元素的和
	for i, val := range nums {
		sum += val
		f += i * val
	}

	maxRF := f

	// 通过计算 f(k) 和 f(k-1) 之间的关系来优化计算
	for i := 1; i < n; i++ {
		f = f + sum - n*nums[n-i]
		if f > maxRF {
			maxRF = f
		}
		fmt.Println(i, "   ", f)
	}

	return maxRF
}