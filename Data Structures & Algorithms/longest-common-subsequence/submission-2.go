func longestCommonSubsequence(text1 string, text2 string) int {
    t1Len, t2Len := len(text1)+1, len(text2)+1
    dp := make([]int, t2Len)

    for idx1 := 1; idx1 < t1Len; idx1++ { 
		
        diagonal := 0
        for idx2 := 1; idx2 < t2Len; idx2++ {
            temp := dp[idx2]
            if text1[idx1-1] == text2[idx2-1] {
                dp[idx2] = 1 + diagonal
            } else {
                dp[idx2] = max(dp[idx2], dp[idx2-1])
            }

            diagonal = temp
        }
    }

    return dp[t2Len-1]
}