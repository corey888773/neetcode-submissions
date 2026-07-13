func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

    cordinates := func (x int) (int, int) {
        return (x / n), (x % n)
    } 

    low, high := 0, m*n
    for low < high {
        mid := low + (high - low) / 2 
        x, y := cordinates(mid)

        if matrix[x][y] == target {
            return true
        }

        if matrix[x][y] > target {
            high = mid
        } else {
            low = mid + 1
        }
    }

    return false
}