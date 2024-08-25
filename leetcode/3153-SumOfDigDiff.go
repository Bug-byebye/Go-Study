package leetcode

import (
	"fmt"
	"strconv"
	"math"
)


func SumDigitDifferences(nums []int) int64 {
	sum := 0
	digitNum := len(strconv.Itoa(nums[0]))
	strArray := make([][]int,len(nums))
	for i := range strArray{
		strArray[i] =make([]int, digitNum)
	}
	var digitOffset float64
	for i := 0; i < len(nums); i++ {
		for ii := 0; ii < digitNum; ii++ {
			if ii==0{
				digitOffset = math.Pow(10,float64(digitNum-1-ii))
			}else{
				digitOffset /= 10
			}

			strArray[i][ii] = nums[i]/int(digitOffset)
            nums[i]=nums[i]%int(digitOffset)

			fmt.Println(nums[i])
		}
	}
	fmt.Println(strArray)
	for i := 0; i < len(nums)-1; i++ {
		for ii := i+1; ii < len(nums); ii++ {
			for iii := 0; iii < digitNum; iii++ {
				if strArray[i][iii] != strArray[ii][iii] {
					sum++
				}
			}
		}
	}
	return int64(sum)
}



func SumDigitDifferencesPro(nums []int) (ans int64) {
	cnt := make([][10]int, len(strconv.Itoa(nums[0])))
	for k, x := range nums {
		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans += int64(k - cnt[i][d])
			cnt[i][d]++
			i++
		}
	}
	return
}
