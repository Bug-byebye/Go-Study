package leetcode

//import "fmt"

var (
	num int
)

func climbStairs(n int) int {
	num = 0
	//leftS := n
	eachClimb(n)
	return num
}

func eachClimb(leftS int) {
	if leftS == 0 {
		num++
		return
	}
	leftS--
	eachClimb(leftS)
	if leftS > 0 {
		leftS--
		eachClimb(leftS)
	}
}
