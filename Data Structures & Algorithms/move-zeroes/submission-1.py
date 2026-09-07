class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        left = 0

        for right, num in enumerate(nums):
            if left == right: continue

            while nums[left] != 0 and left < right:
                left += 1

            nums[right], nums[left] = nums[left], nums[right]



# _0,_0,1,2,0,5
# _0, 0, _1
# 1, _0, 0, _2
# 1, 2, _0, 0, _5
# 1, 2, 5, 0, 0,