class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        result = []
        def backtrack(idx:int, currSum: int, arr: List[int]):
            if currSum == target:
                result.append(arr.copy())

            for i in range(idx, len(nums)):
                num = nums[i]
                if currSum+num <= target:
                    arr.append(num)
                    backtrack(i, currSum+num, arr)
                    arr.pop()


        backtrack(0, 0, [])
        return result
        