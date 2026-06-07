func maxProfit(prices []int) int {
    n := len(prices)
    memo := make([][]int, 2) // hasStock 0/1
    for i := range 2 {memo[i] = make([]int, n)}
    
    memo[0][0] = 0
    memo[1][0] = -prices[0]
    
    for i := 1; i < n; i++ {
        memo[0][i] = max(prices[i] + memo[1][i-1], memo[0][i-1])
        if i >= 2 {
            memo[1][i] = max(-prices[i] + memo[0][i-2], memo[1][i-1])
        } else {
            memo[1][i] = max(-prices[i], memo[1][i-1])
        }
    } 
    
    return memo[0][n-1]
}
