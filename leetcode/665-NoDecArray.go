package leetcode

func checkPossibility(nums []int) bool {
	IsPossible := true

	for i := 1; i < len(nums); i++{
		if nums[i] < nums[i-1]{
			if !IsPossible {
				return false
			}
			IsPossible = false
            if i>=2&&nums[i]<nums[i-2]{
                nums[i] = nums[i-1]
            }
            
		}
	}

	return true
}