type Edge struct {
    u, v, time int
}

type PQ []Edge

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq PQ) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}
func (pq *PQ) Push(x any) { *pq = append(*pq, x.(Edge)) }
func (pq *PQ) Pop() any {
    temp := *pq
    x := temp[len(temp)-1]
    *pq = temp[:len(temp)-1]
    return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    adjecency := make([][]Edge, n+1)
    for _, edge := range times {
        adjecency[edge[0]] = append(adjecency[edge[0]], Edge{edge[0], edge[1], edge[2]})
    }

    networkTimes := make([]int, n+1)
    for i := range n+1 { networkTimes[i] = math.MaxInt }
	networkTimes[0] = -1 // SENTINEL

    pq := PQ{{0, k, 0}}
    heap.Init(&pq)

    for pq.Len() > 0 {
        edge := heap.Pop(&pq).(Edge)

        if edge.time >= networkTimes[edge.v] {
            continue
        }
        networkTimes[edge.v] = edge.time
        
        for _, neigh := range adjecency[edge.v] {
            newTime := edge.time + neigh.time
            if newTime < networkTimes[neigh.v] {
                heap.Push(&pq, Edge{neigh.u, neigh.v, newTime})
            }
        } 
    }

    minTime := math.MinInt
    for _, time := range networkTimes {
        if time == math.MaxInt { return -1 }
    
        minTime = max(minTime, time)
    }

    return minTime
}
