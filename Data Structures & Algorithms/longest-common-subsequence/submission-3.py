class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        n1, n2 = len(text1), len(text2)
        dp = [[-1] * n2 for _ in range(n1)]

        def dfs(i: int, j: int) -> int:
            if i >= n1 or j >= n2:
                return 0

            if dp[i][j] != -1:
                return dp[i][j]

            if text1[i] == text2[j]:
                dp[i][j] = 1 + dfs(i+1, j+1)
            else:
                dp[i][j] = max(dfs(i+1, j),  dfs(i, j+1))

            return dp[i][j]

        return dfs(0, 0)
