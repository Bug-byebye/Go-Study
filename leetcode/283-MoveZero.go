package leetcode

import(
	"fmt"
)

func moveZeroes(nums []int) {

	PlaceOfZreo := []int{-1}
	ZoneOfZero := len(nums) - 1
    firstZero := true

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			PlaceOfZreo = append(PlaceOfZreo, i)
			PlaceOfZreo = PlaceOfZreo[1:]
			break
		}
	}

	if PlaceOfZreo[0] == -1 {
		return
	}

	index := 0
	for index < len(PlaceOfZreo) {

		i := PlaceOfZreo[index]
		
        for ii := i; ii < ZoneOfZero; ii++ {
			tmp := nums[ii+1]
			nums[ii+1] = nums[ii]
			nums[ii] = tmp

			if firstZero && nums[ii] == 0 {
				PlaceOfZreo = append(PlaceOfZreo, ii)
                firstZero = false
			}
			fmt.Println(nums)
			fmt.Println(PlaceOfZreo)
		}
        firstZero = true
		ZoneOfZero--
		index++
	}
}


func moveZeroesPro(nums []int) {
    // 使用两个指针，一个指向当前处理的位置，一个指向应该放置非零元素的位置
    nonZeroIndex := 0

    // 遍历数组，将非零元素移动到前面
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[nonZeroIndex] = nums[i]
            nonZeroIndex++
        }
    }

    // 将剩余的位置填充为零
    for i := nonZeroIndex; i < len(nums); i++ {
        nums[i] = 0
    }
}


