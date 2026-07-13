func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := range amount + 1 {dp[i] = 1e4 + 1}
    dp[0] = 0

    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            if i - c >= 0 {
                dp[i] = min(dp[i], dp[i-c]+1)
            }
        }
    }

    if dp[amount] <= 1e4 {
        return dp[amount]
    }

    return -1
}