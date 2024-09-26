package leetcode

func differenceOfSum(nums []int) int {
	eSum,dSum := 0,0
	for _,v := range nums {
		eSum += v
		for v>0 {
			dSum += (v%10)
			v /= 10
		}
	}

	abDifference := func (a int, b int) int {
		if a>b {
			return a-b
		}else{
			return b-a
		}
	}
	return abDifference(eSum, dSum)
}


func differenceOfSumPRO(nums []int) int {
	diff := 0
	for _,v := range nums {
		diff += v
		for v>0 {
			diff -= (v%10)
			v /= 10
		}
	}
	return diff
}