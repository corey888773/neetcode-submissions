func longestCommonSubsequence(text1 string, text2 string) int {
    t1Len, t2Len := len(text1)+1, len(text2)+1
    memo := make([][]int, t1Len)
    for i := range t1Len { memo[i] = make([]int, t2Len) }

    for idx1 := 1; idx1 < t1Len; idx1++ {
        for idx2 := 1; idx2 < t2Len; idx2++ {
            if text1[idx1-1] == text2[idx2-1] {
                memo[idx1][idx2] = 1 + memo[idx1-1][idx2-1]
            } else {
                memo[idx1][idx2] = max(memo[idx1-1][idx2], memo[idx1][idx2-1])
            }
        }
    }

    return memo[t1Len-1][t2Len-1]
}