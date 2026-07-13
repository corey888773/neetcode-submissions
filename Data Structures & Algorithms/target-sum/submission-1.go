const INF = -1000000

func findTargetSumWays(nums []int, target int) int {
    numsLen := len(nums)
    memo := make([][]int, 2001)
    for i := range 2001 {
        memo[i] = make([]int, numsLen)
        for j := range numsLen {
            memo[i][j] = INF
        }
    }

    var dfs func(idx, sum int) int
    dfs = func(idx, sum int) int {
        if idx == numsLen && sum == target {
            return 1
        }

        if idx >= numsLen {
            return 0
        }

        if memo[sum+1000][idx] != INF {
            return memo[sum+1000][idx]
        }

        memo[sum+1000][idx] = dfs(idx+1, sum+nums[idx]) + dfs(idx+1, sum-nums[idx])
        return memo[sum+1000][idx]
    }
    
    return dfs(0, 0)
}
