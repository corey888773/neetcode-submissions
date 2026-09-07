

class Solution:
   


    def search(self, nums: List[int], target: int) -> int:
        def binsearch(lo: int, hi: int) -> int:
            while lo < hi:
                mid = lo + (hi - lo) // 2

                if nums[mid] >= target:
                    hi = mid
                else:
                    lo = mid + 1

            return lo if nums[lo] == target else -1 


        lo, hi = 0, len(nums)-1

        while lo < hi:
            mid = lo + (hi - lo) // 2

            if nums[mid] < nums[hi]:
                hi = mid
            else:
                lo = mid + 1

        return max(binsearch(0, lo-1), binsearch(lo, len(nums)-1))



        