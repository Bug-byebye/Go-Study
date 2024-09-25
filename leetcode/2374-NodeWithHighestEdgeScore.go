package leetcode

import (
	"fmt"
	//"sort"
)

func edgeScore(edges []int) int {
	scoreNode := make([][]int, len(edges))
	score := make([]int,len(edges))
	highestScore, maxindex := -1, -1
	for i, v := range edges {
		scoreNode[v] = append(scoreNode[v], i)
	}
	fmt.Println(scoreNode)
	for i, _ := range scoreNode {
		for _, vv := range scoreNode[i] {
			score[i] += vv
		}
	}
	fmt.Println(score)
	//sort.Ints(edges)
	for i := 0; i < len(edges); i++ {
		if highestScore<score[i] {
			highestScore = score[i]
			maxindex = i
		}
	}
	return maxindex
}