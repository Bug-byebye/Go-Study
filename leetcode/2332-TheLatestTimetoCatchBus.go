package leetcode

import (
	//"slices"
	"sort"
)

/*func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	latestTime := 0
	sort.IntSlice.Sort(buses)
	sort.IntSlice.Sort(passengers)
	slices.Sort(buses)
    slices.Sort(passengers)
	p := make(map[int]int)
	for i := 0; i < len(passengers); i++ {
		p[i] = 1
	}
	for i := len(buses) - 1; i >= 0; i++ {
		c := capacity
		for ii := buses[i]; c > 0 && ii > 0; ii-- {
			if p[ii] != 1 {
				return ii
			} else {
				c--
			}
		}
	}

	return latestTime
}
*/

func latestTimeCatchTheBus(buses, passengers []int, capacity int) (ans int) {
    
	sort.IntSlice.Sort(buses)
	sort.IntSlice.Sort(passengers)
	
    j, c := 0, 0
    for _, t := range buses {
        for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; c-- {
            j++
        }
    }

    if c > 0 {
        ans = buses[len(buses)-1]
    } else {
        ans = passengers[j-1]
    }
    for j--; j >= 0 && ans == passengers[j]; j-- { 
        ans--
    }
    return
}
