package leetcode

var(
	path []int
	res [][]int
	state []bool
)

func permute(nums []int) [][]int {
	le := len(nums)
	path, res = make([]int, 0, le), make([][]int, 0)
	state = make([]bool, le)
	Backtracking(nums,0)
	return res
}

func Backtracking(nums []int, cur int) {
	if cur==len(nums){
		tmp := make([]int,0,len(nums))
		copy(tmp,path)
		res = append(res, tmp)
	}
	for i:=0;i<len(nums);i++{
		if !state[i]{
			path = append(path,nums[i])
			state[i]=true
			Backtracking(nums, cur+1)
			path = path[:len(path)-1]
			state[i]=false
		}
	}
}