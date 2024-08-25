package leetcode

import "sort"

func countWays(nums []int) int {
	ansNum := 0
	sort.Ints(nums)
	//左右边界判断——对应选中0人和所有人的情况
	if nums[0]>0 {
		ansNum++
	}	
	if nums[len(nums)-1]<len(nums){
		ansNum++
	}
	//遍历判断是否可以只选i人
	for i := 1; i < len(nums); i++ {
		if nums[i]>i && nums[i-1]<i {
			ansNum++
		}
	}

	return ansNum
}