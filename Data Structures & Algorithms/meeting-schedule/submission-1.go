func canAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    }) 
        
    for l := 0; l < len(intervals); l++ {
        for r := l+1; r < len(intervals); r++ {
            if intervals[l].start < intervals[r].end && intervals[r].start < intervals[l].end {
                return false
            }
            
            if intervals[l].end < intervals[r].start {
                break
            }
        }
    }
    
    return true
}