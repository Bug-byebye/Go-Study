package leetcode

func longestContinuousSubstring(s string) int {
    
    cur := s[0]
    length, lenmeom := 1, 1
    
    for i := 1; i < len(s); i++ {
        
        if s[i]-cur==1 {
            length++
        }else{
            if length>lenmeom {
                lenmeom = length
            }
            length = 1
        }
        cur = s[i]
    }	

    if length>lenmeom {
        lenmeom = length
    }
        
    return lenmeom
}