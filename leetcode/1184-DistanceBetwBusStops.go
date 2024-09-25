package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	dis1,dis2,mdis := 0,0,0
	if start<destination {
		for i := start; i < destination; i++ {
			dis1 += distance[i]
		}
		for i := destination; i%n != start; i++ {
			dis1 += distance[i%n]
		}
		mdis = min(dis1,dis2)
		return mdis
	}
	if start>destination {
		for i := destination; i < start; i++ {
			dis2 += distance[i]
		}
		for i := start; i%n != destination; i++ {
			dis1 += distance[i%n]
		}
		mdis = min(dis1,dis2)
		return mdis
	}

	return 0

}