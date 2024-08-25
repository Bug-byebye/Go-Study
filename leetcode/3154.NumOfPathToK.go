package leetcode


const mx2 = 31

func binomialCoefficient(n, k int) int {
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}

func waysToReachStair(k int) (ans int) {
	for j := 0; j < 30; j++ {
		m := 1<<j - k
		if 0 <= m && m <= j+1 {
			ans += binomialCoefficient(j+1, m)
		}
	}
	return
}


