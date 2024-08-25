package leetcode
func maxProfit(prices []int) int {
	le := len(prices)
	
	if le==0||le==1 {
		return 0
	}
	
	property := make([][]int,le)
	for i:=0;i<le;i++{
		property[i] = make([]int, 2)
	}

	property[0][0] = -prices[0]
	property[0][1] = 0

	for i:=1;i<le;i++{
		property[i][0] = max(property[i-1][0],-prices[i])
		property[i][1] = max(property[i-1][1], prices[i]+property[i-1][0])
	}
	return property[le-1][1]
} 
