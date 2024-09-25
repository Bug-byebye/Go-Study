package leetcode

//import "fmt"

func maxScoreSightseeingPair(values []int) int {
	maxScore := -1
	for i := 0; i < len(values)-1; i++ {
		for ii := i + 1; ii < len(values); ii++ {

			score := values[i] + values[ii] + i - ii
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

/*func maxScoreSightseeingPairPRO(values []int) int {
	ml, mr := 0, len(values)-1
	maxScore := values[ml] + values[mr] + ml - mr

	for dpl, dpr := 1, len(values)-2; dpl < dpr; {
		if values[ml] < values[dpl]+1 {
			maxScore = values[dpl] + values[mr] + dpl - mr
			ml = dpl
		}

		if values[mr] < values[dpr]+1 {
			maxScore = values[dpr] + values[ml] + ml - dpr
			mr = dpr
		}
		dpr--
		dpl++
		fmt.Println(ml,mr,maxScore)
	}
	return maxScore
}*/


func maxScoreSightseeingPairPRO(values []int) int {

	maxLeft := values[0] 
	maxScore := 0

	for j := 1; j < len(values); j++ {
		
		maxScore = max(maxScore, maxLeft + values[j] - j)
		maxLeft = max(maxLeft, values[j] + j)
	}

	return maxScore
}
