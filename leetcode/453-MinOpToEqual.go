package leetcode

func minMoves(nums []int) int {
	minMove  := 0
	minofarr := nums[0]
	for i := 1; i<len(nums); i++{
		if nums[i] < minofarr{
			minofarr = nums[i]
			minMove = minMove + i*(minofarr-nums[i]) 
		}else{
			minMove = minMove + (nums[i]-minofarr)
		}
	}
	return minMove
}
