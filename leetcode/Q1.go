package leetcode

func getFinalState(nums []int, k int, multiplier int) []int {
	for i:=1;i<=k;i++{
		minindex := 0
		minv := nums[0] 
		for ii := 1; ii<len(nums) ; ii++{
			if nums[ii]<minv{
				minindex = ii
				minv = nums[ii]
			}
		}
		nums[minindex] = nums[minindex]*multiplier
	}
	return nums
}
