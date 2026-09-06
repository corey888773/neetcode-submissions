class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        prefixes = {0: 1}
        prefix = 0
        count = 0

        for i in range(1, len(nums)+1):
            prefix += nums[i-1]
            rest = prefix % k
            count += prefixes.get(rest, 0)

            prefixes[rest] = prefixes.get(rest, 0) + 1 

        return count
