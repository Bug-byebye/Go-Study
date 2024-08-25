package leetcode

import "math"

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		cnt := [26]int{}
		k, maxCnt := 0, 0
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			maxCnt = max(maxCnt, cnt[b])
			if i-j+1 == k*maxCnt {
				res = min(res, dfs(j-1)+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n - 1)
}

