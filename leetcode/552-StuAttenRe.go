package leetcode

const mod = 1_000_000_007
const mx1 = 100_001

var memo [mx1][2][3]int

// 初始化 memo 数组，将所有状态标记为未计算过（-1）
func initMemo() {
    for i := range memo {
        for j := range memo[i] {
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }
}

// 递归函数 dfs(i, j, k)
// i: 剩余的天数
// j: 是否已经有一个 'A'（0 表示没有，1 表示有）
// k: 当前以多少个连续的 'L' 结尾
func dfs(i, j, k int) int {
    if i == 0 {
        return 1 // 如果没有剩余天数，这是一个有效的序列
    }
    if memo[i][j][k] != -1 {
        return memo[i][j][k] // 如果该状态已经计算过，直接返回保存的结果
    }
    
    // 计算可能的合法序列数
    res := dfs(i-1, j, 0) // 当前字符是 'P'
    if j == 0 {
        res += dfs(i-1, 1, 0) // 当前字符是 'A'，并且还没有出现过 'A'
    }
    if k < 2 {
        res += dfs(i-1, j, k+1) // 当前字符是 'L'，并且之前连续的 'L' 少于 2 个
    }
    
    // 结果取模后存储，避免溢出并进行记忆化
    memo[i][j][k] = res % mod
    return memo[i][j][k]
}

// 主函数，计算长度为 n 的合法排列数
func checkRecord(n int) int {
    initMemo() // 初始化 memo 数组
    return dfs(n, 0, 0) // 从 n 天开始，没有 'A'，没有连续 'L'
}
