const (
    UNKNOWN = 0
    MATCHED = 1
    NOT_MATCHED = -1
)

func isMatch(s string, p string) bool {
    stringLen, patternLen := len(s), len(p)
    memo := make([][]int, stringLen+1)
    for i := range stringLen+1 {memo[i] = make([]int, patternLen+1)}

    var dfs func(sIdx, pIdx int) int
    dfs = func(sIdx, pIdx int) int {
        // end of both - MATCHED
         if pIdx == patternLen {
            if sIdx == stringLen { return MATCHED }
            return NOT_MATCHED
        }

        if memo[sIdx][pIdx] != UNKNOWN {
            return memo[sIdx][pIdx]
        }

        memo[sIdx][pIdx] = NOT_MATCHED // marking as visited
        if pIdx + 1 < patternLen && p[pIdx + 1] == '*' {
            skipPattern := dfs(sIdx, pIdx+2) // skips the pattern completely
            matchOne := NOT_MATCHED
            if sIdx < stringLen && (p[pIdx] == '.' || p[pIdx] == s[sIdx]) {
                matchOne = dfs(sIdx+1, pIdx) // matches one character of pattern and moves to the next substring char keeping the pattern
            }

            memo[sIdx][pIdx] = max(skipPattern, matchOne)
        } else if sIdx < stringLen && (p[pIdx] == '.' || p[pIdx] == s[sIdx]) {
            matchOne := dfs(sIdx+1, pIdx+1)
            memo[sIdx][pIdx] = matchOne
        }

        return memo[sIdx][pIdx]
    }

    return dfs(0, 0) == MATCHED
}