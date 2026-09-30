class Solution:
    def trap(self, height: list[int]) -> int:
        stack = []

        trapped_water = 0
        for i, hi in enumerate(height):
            while len(stack) > 0:
                if hi < height[stack[-1]]: break

                top = stack.pop()
                if len(stack) == 0: break # no left border, water overflows

                left = stack[-1]
                right = i
                water_level = min(height[left], height[right]) - height[top]
                trapped_water += (right - left - 1) * water_level

            stack.append(i)

        return trapped_water