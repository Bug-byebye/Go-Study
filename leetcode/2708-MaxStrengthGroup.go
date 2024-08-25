package leetcode

import (
	//"slices"
	"sort"
)

func maxStrength(nums []int) int64 {
	strength := 1
	negativeNum := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			
			if nums[i] < 0 {
				negativeNum = append(negativeNum, nums[i])
			}
			strength *= nums[i]
		}else{
			if (len(nums)==2&&(strength<0 || nums[1]<0))||len(nums)==1 {
				return 0
			}
		}

	}
	if strength < 0 && len(nums)>1{
		sort.Ints(negativeNum)
		strength /= negativeNum[len(negativeNum)-1] 
	}
	return int64(strength)
}


func maxStrengthPro(nums []int) int64 {
    mn, mx := nums[0], nums[0]
    for _, x := range nums[1:] {
        mn, mx = min(mn, x, mn*x, mx*x),
                 max(mx, x, mn*x, mx*x)
    }
    return int64(mx)
}
