package leetcode

func findJudge(n int, trust [][]int) int {
	
	if n==1{
        return 1
    }
	
	trusted := make([]int ,n+1)
	trusting := make([]int ,n+1)
	
	for i := 0; i < len(trust); i++ {
		trusting[trust[i][0]]++
		trusted[trust[i][1]]++
	}

	for i := 0; i <= n; i++ {
		if trusted[i]==n-1 && trusting[i]==0 {
			return i
		}
	}

	return -1
}