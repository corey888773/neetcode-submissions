func longestCommonSubsequence(text1 string, text2 string) int {
    t1Len, t2Len := len(text1), len(text2)
    memo := make([][]int, t1Len)
    for i := range t1Len {
        memo[i] = make([]int, t2Len)
        for j := range t2Len {
            memo[i][j] = -1
        }
    }

    var dfs func(idx1, idx2 int) int 
    dfs = func(idx1, idx2 int) int {
        if idx1 >= t1Len || idx2 >= t2Len {
            return 0
        }

        if memo[idx1][idx2] != -1 {
            return memo[idx1][idx2]
        }

        if text1[idx1] == text2[idx2] {
            memo[idx1][idx2] = 1 + dfs(idx1+1, idx2+1)
        } else {
            memo[idx1][idx2] = max(dfs(idx1+1, idx2), dfs(idx1, idx2+1))
        }

        return memo[idx1][idx2]
    }

    return dfs(0, 0)
}

