package leetcode
import(
	"slices"
)
func rotate(nums []int, k int)  {
	k = k%len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}