package leetcode

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < 2; i++ {
		if grid[i][0]==grid[i][1] && (grid[i+1][0]==grid[i][0] || grid[i+1][1]==grid[i][0]){
			return true
		}
		if grid[i][2]==grid[i][1] && (grid[i+1][2]==grid[i][1] || grid[i+1][1]==grid[i][1]){
			return true
		}
	}
	for i := 2; i > 0; i-- {
		if grid[i][2]==grid[i][1] && (grid[i-1][2]==grid[i][1] || grid[i-1][1]==grid[i][1]){
			return true
		}
		if grid[i][0]==grid[i][1] && (grid[i-1][0]==grid[i][0] || grid[i-1][1]==grid[i][0]){
			return true
		}
	}

	return false
}


func canMakeSquareProMax(grid [][]byte) bool {
	check := func(i, j int) bool {
		cnt := [2]int{}
		cnt[grid[i][j]&1]++
		cnt[grid[i][j+1]&1]++
		cnt[grid[i+1][j]&1]++
		cnt[grid[i+1][j+1]&1]++
		return cnt[0] != 2
	}
	return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1)
}
