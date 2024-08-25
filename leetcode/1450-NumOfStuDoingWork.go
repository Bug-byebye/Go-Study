package leetcode

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	num:=0
	for i,v := range endTime{
		if v>=queryTime && startTime[i]<=queryTime{
			num++
		}
	}
	return num
}