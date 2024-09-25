package leetcode

import "fmt"

func maxofINTslices(nums []int, index int) int {
	maxNum := nums[0]
	for i := 0; i <= index; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	fmt.Print(maxNum,' ')
	return maxNum
}
func sumofINTslices(nums []int, index int) int {
	sumNum := 0
	for i := 0; i <= index; i++ {
		sumNum += nums[i]
	}
	fmt.Print(sumNum,' ')
	return sumNum
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	i := 0
	for i < len(chargeTimes) {
		cost := maxofINTslices(chargeTimes, i) + (i+1)*sumofINTslices(runningCosts, i)
		if cost > int(budget) {
			return i + 1
		}
		fmt.Println(cost)
		i++
	}
	return i
}

func maximumRobotsPlus(chargeTimes, runningCosts []int, budget int64) (ans int) {
    q := []int{}
    sum := int64(0)
    left := 0
    for right, t := range chargeTimes {
        // 1. 入
        for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, right)
        sum += int64(runningCosts[right])

        // 2. 出
        for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
            if q[0] == left {
                q = q[1:]
            }
            sum -= int64(runningCosts[left])
            left++
        }

        // 3. 更新答案
        ans = max(ans, right-left+1)
    }
    return
}