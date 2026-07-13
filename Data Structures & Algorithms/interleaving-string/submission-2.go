

func isInterleave(s1, s2, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    dp := make([][]int, len(s1) + 1)
    for i := range len(s1) + 1 {
        dp[i] = make([]int, len(s2)+ 1)
        for j := range len(s2) +1 { dp[i][j] = -1 }
    }

    var dfs func(i, j int) int 
    dfs = func(i, j int) int {
        if i > len(s1) || j > len(s2) {
            return 0
        }

        if i + j == len(s3) {
            return 1
        } 

        if dp[i][j] != -1 {
            return dp[i][j]
        }

        dp[i][j] = 0
        if (i < len(s1) && s3[i + j] == s1[i]) {
            dp[i][j] = max(dp[i][j], dfs(i+1, j))
        }

        if (j < len(s2) && s3[i + j] == s2[j]) {
            dp[i][j] = max(dp[i][j], dfs(i, j+1))
        }
        
        return dp[i][j]
    }

    return dfs(0, 0) == 1
}