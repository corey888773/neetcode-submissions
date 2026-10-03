import heapq

class MedianFinder:

    def __init__(self):
        self.min_heap = []
        self.max_heap = []
        self.size = 0

    def addNum(self, num: int) -> None:
        self.size += 1
        if len(self.min_heap) == len(self.max_heap):
            if (self.min_heap and num < self.min_heap[0]) or not self.max_heap:
                heapq.heappush(self.max_heap, -num)
            else:
                transfer = heapq.heappop(self.min_heap)
                heapq.heappush(self.min_heap, num)
                heapq.heappush(self.max_heap, -transfer)


        else:
            if num >= -self.max_heap[0]:
                heapq.heappush(self.min_heap, num)
            else:
                transfer = -heapq.heappop(self.max_heap)
                heapq.heappush(self.min_heap, transfer)
                heapq.heappush(self.max_heap, -num)



    def findMedian(self) -> float:

        if self.size % 2 == 0: 
            return (-self.max_heap[0] + self.min_heap[0]) / 2
        
        return -self.max_heap[0]




# [1] [] 2
# [1] [2]

# [1] [] 0
# [0] [1]

# [1, 2], [4] 3
# [1, 2], [3, 4]

# [1, 2], [3] 4
# [1, 2], [3, 4]

# [1, 3], [4] 2
# [1, 2], [3, 4]
 

# [1] [3] 2 
# [1, 2], [3]

# [1] [2] 3
# [1, 2] [3]

# [1] [2] 0
# [0, 1], [2]
