type Minh []int
 
func (h Minh) Len() int {return len(h)}
func (h Minh) Less(i, j int) bool {return h[i] < h[j]}
func (h Minh) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *Minh) Push(x any) {*h = append(*h, x.(int)) }
func (h *Minh) Pop() any {
    temp := *h
    x := temp[len(temp)-1]
    *h = temp[:len(temp)-1]
    return x 
}

func minMeetingRooms(intervals []Interval) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
    
    h := &Minh{}
    maxConcurrent := 0
    for _, inter := range intervals {
        for h.Len() > 0 && (*h)[0] <= inter.start {
            heap.Pop(h)
        }
        
        heap.Push(h, inter.end)
        maxConcurrent = max(h.Len(), maxConcurrent)
    }
 
    return maxConcurrent
}