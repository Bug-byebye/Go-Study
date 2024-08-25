package leetcode

//import "math"

func findPermutationDifference(s string, t string) int {
	diff := 0
	for index, val := range s {
		ii := index
		iv := val
		for index, val := range t {
			if val == iv {
				if ii>index {
					diff += (ii-index)
				}else{
					diff += (index-ii)
				}
			break
				
			}
		}
	}
	return diff
}