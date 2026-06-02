
func minMeetingRooms(intervals []Interval) int {
    timeRange := int(math.Pow(10,6)) + 1
    events := make([]int, timeRange)

    for _, interval := range intervals {
        events[interval.start]++
        events[interval.end]--
    }

    maxRooms := 0
    cumulativeRooms := 0
    for _, e := range events {
        cumulativeRooms += e
        maxRooms = max(maxRooms, cumulativeRooms)
    }

    return maxRooms
}