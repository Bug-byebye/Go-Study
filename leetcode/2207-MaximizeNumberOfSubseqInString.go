package leetcode

func maximumSubsequenceCount(text string, pattern string) int64 {
	a,b := pattern[0], pattern[1]
	maxn, numOfa, numOfb := 0,0,0

	for i := 0; i < len(text); i++ {
		if text[i]==b {
			maxn+=numOfa
			numOfb++
		}
		if text[i]==a {
			numOfa++
			continue
		}


	}
	m1 := maxn + numOfa
	m2 := maxn + numOfb
	maxn = max(m1,m2)
	return int64(maxn)

}