import heapq

class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        m, n = len(heights), len(heights[0])
        min_efforts = [[float('inf')] * n for _ in range(m)]
        min_efforts[0][0] = 0
        min_heap = []
        heapq.heappush(min_heap, (0,0,0))

        while len(min_heap) > 0:
            curr = heapq.heappop(min_heap)
            effort, x, y = curr

            if min_efforts[x][y] < effort:
                continue

            for move in [[0, 1], [1, 0], [-1, 0], [0, -1]]:
                new_x, new_y = x + move[0], y + move[1]
                if new_x < 0 or new_x >= m or new_y < 0 or new_y >= n:
                    continue

                new_effort = max(abs(heights[x][y] - heights[new_x][new_y]), min_efforts[x][y])
                if min_efforts[new_x][new_y] > new_effort:
                    min_efforts[new_x][new_y] = new_effort
                    heapq.heappush(min_heap, (new_effort, new_x, new_y))
        

        return min_efforts[m-1][n-1]
                

            
