func rob(houses []int) int {
    if len(houses) == 1 {
        return houses[0]
    }

    return max(robOne(houses[:len(houses)-1]), robOne(houses[1:]))
}

func robOne(houses []int) int {
    memo := make([][]int, 2) // previous house robbed or not robbed
    for i := range 2 {
        memo[i] = make([]int, len(houses))
        for j := range len(houses) {
            memo[i][j] = -1
        }
    }

    var dfs func(houses []int, idx int, prevRobbed int) int
    dfs = func(houses []int, idx int, prevRobbed int) int {
        if idx == len(houses) {
            return 0
        }

        if memo[prevRobbed][idx] != -1 {
            return memo[prevRobbed][idx]
        }

        skip := dfs(houses, idx+1, 0)
        memo[prevRobbed][idx] = skip // we don't rob this one
        if prevRobbed == 0 {
            rob := houses[idx] + dfs(houses, idx+1, 1)
            memo[prevRobbed][idx] = max(skip, rob) // we can rob this one
        }

        return memo[prevRobbed][idx]
    }

    return dfs(houses, 0, 0)
}