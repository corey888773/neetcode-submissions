func maxCoins(nums []int) int {
    n := len(nums) + 2
    vals := make([]int, n)
    vals[0], vals[n-1] = 1, 1
    copy(vals[1:], nums)

    memo := make([][]int, n)
    for i := range n {
        memo[i] = make([]int, n)
        for j := range n {
            memo[i][j] = -1
        }
    }

    var burst func(left, right int) int
    burst = func(left, right int) int {
        if right - left <=1 {
            return 0
        }

        if memo[left][right] != -1 {
            return memo[left][right]
        }

        for i := left+1; i < right; i++ {
            currentBurst := vals[left]*vals[i]*vals[right] + burst(left, i) + burst(i, right)
            memo[left][right] = max(memo[left][right], currentBurst)
        }

        return memo[left][right]
    }


    return burst(0, n-1)
}
