package leetcode

var(
	tri [][]int
)
func generate(numRows int) [][]int {
	tri = make([][]int,numRows)


	for i:=0;i<numRows;i++{
       	tri[i] = make([]int, i+1)
		tri[i][0] = 1
		tri[i][i] = 1
		for n:=1;n<i;n++{
			tri[i][n] = tri[i-1][n-1]+tri[i-1][n]
		}
	}
	return tri
}