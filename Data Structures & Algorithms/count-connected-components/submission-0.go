func countComponents(n int, edges [][]int) int {
    web := make([][]int, n)
    for i := range n {
        web[i] = []int{}
    } 

    for _, e := range edges {
        from := e[0]    
        to := e[1]

        web[from] = append(web[from], to)
        web[to] = append(web[to], from)
    }
    

    visited := map[int]bool{}
    counter := 0
    q := []int{}
    for i := range n {
        if _, ok := visited[i]; !ok {
            q = append(q, i)
            visited[i] = true
            counter += 1
        }

        for len(q) > 0 {
            f := q[0]
            q = q[1:]

            for _, t := range web[f]{
                if _, ok := visited[t]; !ok {
                    q = append(q, t)
                    visited[t] = true
                }
            }
        }
    }

    return counter
}
