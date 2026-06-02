func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    
    queue := [][]int{} 
    for r := range rows {
        visited[r] = make([]bool, cols)
        for c := range cols {
            if grid[r][c] == 0 {
                queue = append(queue,[]int{r,c})
                visited[r][c] = true
            }
        }
    }
    
    directions := [][]int{[]int{1,0}, []int{-1,0}, []int{0,1}, []int{0,-1}}
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]; queue = queue[1:]
        
        for _, dir := range directions {
            newX, newY := x + dir[0], y + dir[1]
            
            if newX < 0 || newY < 0 || newX >= rows || newY >= cols || grid[newX][newY] == -1 || visited[newX][newY] {
                continue
            }
            
            visited[newX][newY] = true
            grid[newX][newY] = grid[x][y] + 1
            queue = append(queue, []int{newX, newY})
        } 
    }
    
    
    return
}