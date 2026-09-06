class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        n = len(nums)
        prefix = 0
        precedences = {0: 1}
        num_of_k_subs = 0
        
        for i in range(1, n+1):
            prefix += nums[i-1]
            num_of_k_subs += precedences.get(prefix - k, 0)
            precedences[prefix] = precedences.get(prefix, 0) + 1

        return num_of_k_subs