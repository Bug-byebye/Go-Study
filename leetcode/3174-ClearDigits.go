package leetcode

func clearDigits(s string) string {
	ss := []rune(s) // 将字符串转为 rune 切片，方便处理 Unicode 字符

	for i := 0; i < len(ss); i++ {
		// 判断是否是数字字符
		if ss[i] >= '0' && ss[i] <= '9' {
			if i > 0 {
				// 删除当前数字和其左边的字符
				ss = append(ss[:i-1], ss[i+1:]...)
				i -= 2 
			} else {
				// 如果数字在最前面，直接删除该数字
				ss = ss[1:]
				i-- 
			}
		}
	}

	return string(ss) 
}
