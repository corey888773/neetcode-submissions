type item struct {
	node int
	time int
}

type PQ []item 

func (pq PQ) Len() int {return len(pq)}
func (pq PQ) Less(i, j int) bool {return pq[i].time < pq[j].time}
func (pq PQ) Swap(i, j int) {pq[i],pq[j] = pq[j], pq[i]}
func (pq *PQ) Push(x any) {*pq = append(*pq, x.(item))}
func (pq *PQ) Pop() any {
	temp := *pq
	x := temp[len(temp)-1]
	*pq = temp[:len(temp)-1]
	return x
}


func networkDelayTime(times [][]int, n int, k int) int {
    minTime := make([]int, n+1)
	for i := range n+1 { minTime[i] = math.MaxInt }
	minTime[0] = 0 // sentinel
	minTime[k] = 0

	adj := make([][]item, n+1)
	for _, t := range times {
		adj[t[0]] = append(adj[t[0]], item{t[1], t[2]})
	}

	pq := &PQ{item{k, 0}}
	for pq.Len() > 0 {
		stop := heap.Pop(pq).(item)

		if stop.time > minTime[stop.node] {
			continue
		}

		for _, neigh := range adj[stop.node] {
			nd := stop.time + neigh.time
			if nd < minTime[neigh.node] {
				minTime[neigh.node] = stop.time + neigh.time
				heap.Push(pq, item{neigh.node, nd})
			}
		}
	}

	totalMaxTime := math.MinInt
	for _, t := range minTime {
		totalMaxTime = max(totalMaxTime, t)
	}

	if totalMaxTime == math.MaxInt {
		return -1
	} else {
		return totalMaxTime
	}
}
