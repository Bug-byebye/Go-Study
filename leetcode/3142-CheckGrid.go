package leetcode

func satisfiesConditions(grid [][]int) bool {
	for l:=0;l<len(grid)-1;l++{
		for h:=0;h<len(grid[l]);h++{
			if grid[l][h]!=grid[l+1][h]{
				return  false
			}
			if h<len(grid[l])-1 && grid[l][h]==grid[l][h+1]{
				return false
			}
		}
	}
    for l,h:=len(grid)-1,0;h<len(grid[l]);h++{
			if h<len(grid[l])-1 && grid[l][h]==grid[l][h+1]{
				return false
			}
		}
	return true
}