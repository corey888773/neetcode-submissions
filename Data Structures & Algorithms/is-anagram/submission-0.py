class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        s_d, t_d = dict(), dict()

        for i in range(len(s)):
            s_d[s[i]] = s_d.get(s[i], 0) + 1
            t_d[t[i]] = t_d.get(t[i], 0) + 1

        return s_d == t_d
        