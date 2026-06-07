func maxProfit(prices []int) int {
    n := len(prices)
    memo := make([][]int, 2) // hasStock 0/1
    for i := range 2 {
        memo[i] = make([]int, n)
        for j := range n {
            memo[i][j] = -1
        }
    }
    
    var dfs func(day, hasStock int) int
    dfs = func(day, hasStock int) int {
        if day >= n {
            return 0
        }
    
        if memo[hasStock][day] != -1 {
            return memo[hasStock][day]
        }
        
        cooldown := dfs(day+1, hasStock)
        
        if hasStock == 1 {
            memo[hasStock][day] = max(prices[day] + dfs(day+2, 0), cooldown)
        } else {
            memo[hasStock][day] = max(-prices[day] + dfs(day+1, 1), cooldown)
        }
        
        return memo[hasStock][day]
    }
    
    
    return dfs(0, 0)
}
