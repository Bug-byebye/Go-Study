package leetcode

func fib(n int) int {
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	fn:=fib(n-1)+fib(n-2)
	return fn
}

var fibseq []int
func fibplus(n int) int {
	fibseq = make([]int, n+1)
	fibseq[0]=0
	if n>0{
		fibseq[1]=1
	}	
	if n>=2{
		for i := 2; i <= n; i++ {
			fibseq[i]=fibseq[i-1]+fibseq[i-2]
		}
	}
	return fibseq[n]
}

var dp []int
func fibpro(n int) int {
	dp = make([]int, 2)
	dp[0]=0
	dp[1]=1
	if n==0{
		return 0
	}	
	if n>=2{
		for i := 2; i <= n; i++ {
			tmp :=dp[1]
			dp[1]=dp[0]+dp[1]
			dp[0] = tmp
		}
	}
	return dp[1]
}