package leetcode

import (
	//"slices"
	"sort"
	//"sort"
)

/*func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	max := 0
	for _, val := range nums {
		sum = sum + val
		if val > max {
			max = val
		}
	}
	if sum%k == 0 || max <= sum/k {
		sort.Ints(nums)


	} else {
		return false
	}

}*/

func canPartitionKSubsets(nums []int, k int) bool {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%k != 0 {
		return false
	}
	s /= k
	cur := make([]int, k)
	n := len(nums)

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			return true
		}
		for j := 0; j < k; j++ {
			if j > 0 && cur[j] == cur[j-1] {
				continue
			}
			cur[j] += nums[i]
			if cur[j] <= s && dfs(i+1) {
				return true
			}
			cur[j] -= nums[i]
		}
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return dfs(0)
}
