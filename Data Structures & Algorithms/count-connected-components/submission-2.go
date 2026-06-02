func countComponents(n int, edges [][]int) int {
    parents := make([]int, n)
    rootCount := n

    for i := range n { parents[i] = i }
    
    var find func(x int) int 
    find = func(x int) int {
        if parents[x] != x {
            parents[x] = find(parents[x])
        }
        
        return parents[x]
    }
    
    union := func(x, y int) {
        rootX := find(x)
        rootY := find(y)
        
        if rootX == rootY {
            return
        }
        
        rootCount--
        parents[rootX] = rootY
    }
    
    for _, edge := range edges {
        union(edge[0], edge[1])
    }
    
    return rootCount
}