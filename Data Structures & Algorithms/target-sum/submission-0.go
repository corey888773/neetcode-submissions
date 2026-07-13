func findTargetSumWays(nums []int, target int) int {
    numsLen := len(nums)
    counter := 0

    var dfs func(idx, sum int)
    dfs = func(idx, sum int) {
        if idx == numsLen && sum == target {
            counter++
        }

        if idx >= numsLen {
            return
        }

        dfs(idx+1, sum+nums[idx])
        dfs(idx+1, sum-nums[idx])
    }
    dfs(0, 0)

    return counter
}