class Solution:
    def trap(self, height: list[int]) -> int:
        trapped_water = 0
        left, right, left_max, right_max = 0, len(height)-1, 0, 0

        while left < right:
            if height[left] < height[right]:
                left_max = max(left_max, height[left])
                trapped_water += left_max - height[left]
                left += 1
            else:
                right_max = max(right_max, height[right])
                trapped_water += right_max - height[right]
                right -= 1
        
        return trapped_water



