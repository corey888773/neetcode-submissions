func maximalSquare(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i := range m { dp[i] = make([]int, n) }

    maxSquareSide := 0
    for i := range m {
        for j := range n {
            if matrix[i][j] == '0' {
                continue
            }
            
            top, left, diagonal := 0,0,0 
            if i != 0 && j != 0 {
                diagonal = dp[i-1][j-1]
            } 

            if i != 0 {
                top = dp[i-1][j]
            }

            if j != 0 {
                left = dp[i][j-1]
            }

            dp[i][j] = min(min(left, top), diagonal) + 1
            maxSquareSide = max(dp[i][j], maxSquareSide) 
        }
    }

    return maxSquareSide * maxSquareSide
}