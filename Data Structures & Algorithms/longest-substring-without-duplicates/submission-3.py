class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = 0
        last_positions = dict()
        max_len = 0

        for right, char in enumerate(s):
            if char in last_positions and last_positions[char] >= left:
                left = last_positions[char] + 1

            last_positions[char] = right
            max_len = max(max_len, right - left + 1)

        return max_len




