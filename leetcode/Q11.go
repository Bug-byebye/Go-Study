package leetcode

import (
	"math"
	"sort"
)

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	c1, c2 := 0, 0
	/*r1 := coordinate1[0]
	b1 := coordinate1[1]
	r2 := coordinate2[0]
	b2 := coordinate2[1]*/
	c1 = (int(coordinate1[0]) - int(coordinate1[1])) % 2
	c2 = (int(coordinate2[0]) - int(coordinate2[1])) % 2

	if c1 != c2 {
		return false
	} else {
		return true
	}
}

func resultsArray(queries [][]int, k int) []int {
	l := len(queries)
	distance := make([]int, l)
	result := make([]int,l)
	if l<k{
		for i := 0; i < l; i++ {
			result[i]=-1
		}
		return result
	}

	for i := 0; i < l; i++ {
		dis := int(math.Abs(float64(queries[i][1])) +math.Abs(float64(queries[i][0])))
		distance[i] = dis
		
		if i>=k-1 {
			sort.Ints(distance[:i+1])
			result[i]=distance[k-1]
		}else{
			result[i] = -1
		}
		
	}
	
	return result
}


func maxScore(grid [][]int) int {
	
	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		sort.Slice(grid[i], func(a, b int) bool {
			return grid[i][a] > grid[i][b]
		})
	}
	totalScore := 0
	usedRows := make(map[int]bool)
	usedValues := make(map[int]bool)

	for col := 0; col < m; col++ {
		for row := 0; row < n; row++ {
			value := grid[row][col]
			if !usedRows[row] && !usedValues[value] {
				totalScore += value
				usedRows[row] = true
				usedValues[value] = true
				break
			}
		}
	}

	return totalScore
}
