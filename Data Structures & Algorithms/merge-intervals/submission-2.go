func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }

        return intervals[i][0] < intervals[j][0]
    })

    mergedIntervals := [][]int{}
    left, right := 0, 0
    for left < len(intervals) {
		start, end := intervals[left][0], intervals[left][1]
		for right+1 < len(intervals) && end >= intervals[right+1][0] {
			right++
			end = max(end, intervals[right][1])
		}

        mergedIntervals = append(mergedIntervals, []int{start, end})
        left = right+1 
    }

    return mergedIntervals
}

