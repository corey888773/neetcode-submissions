class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        rests = {0:0}
        prefix = 0

        for i in range(1, n+1):
            prefix += nums[i-1]
            
            curr_rest = prefix % k
            if (val := rests.get(curr_rest)) is not None and i - val > 1:
                return True

            if curr_rest not in rests:
                rests[curr_rest] = i

        return False



