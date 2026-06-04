
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, n)

    for i := range m {
        for j := range n {
            left, top := math.MaxInt, math.MaxInt
            if i != 0 {
                top = dp[j]
            }

            if j != 0 {
                left = dp[j-1]
            }

            if i == 0 && j == 0 {
                dp[j] = grid[i][j]
            } else {
                dp[j] = grid[i][j] + min(left, top)
            }
        }
    }

    return dp[n-1]
}