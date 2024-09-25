package leetcode

/*func distinctNames(ideas []string) int64 {
	var numOfName int64
	title := make(map[byte]int, len(ideas))
	for _,v := range ideas{
		title[v[0]] +=1
	}
	for i := 0; i < len(ideas); i++ {
		
	}
	return numOfName
}*/


func distinctNamesPRO(ideas []string) (ans int64) {
    group := [26]map[string]bool{}
    for i := range group {
        group[i] = map[string]bool{}
    }
    for _, s := range ideas {
        group[s[0]-'a'][s[1:]] = true
    }

    for i, a := range group { 
        for _, b := range group[:i] {
            m := 0 
            for s := range a {
                if b[s] {
                    m++
                }
            }
            ans += int64(len(a)-m) * int64(len(b)-m)
        }
    }
    return ans * 2 
}
