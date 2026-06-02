type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    temp := *h
    *h = temp[:len(temp)-1]
    x := temp[len(temp)-1]
    return x
}

func minMeetingRooms(intervals []Interval) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    minHeap := MinHeap{}
    heap.Init(&minHeap)

    maxRooms := 0
    for _, interval := range intervals {
        for minHeap.Len() > 0 && minHeap[0] <= interval.start {
            heap.Pop(&minHeap)
        }

        heap.Push(&minHeap, interval.end)
        maxRooms = max(maxRooms, minHeap.Len())
    }

    return maxRooms
}