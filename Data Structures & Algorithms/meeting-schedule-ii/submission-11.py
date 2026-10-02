"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

import heapq

class Solution:
    def minMeetingRooms(self, intervals: List[Interval]) -> int:
        intervals.sort(key=lambda x: (x.start, x.end))
        max_num_of_rooms = 0
        heap = []

        for i in intervals:
            while len(heap) > 0 and heap[0][0] <= i.start:
                heapq.heappop(heap)    
            
            heapq.heappush(heap, (i.end, i.start))
            max_num_of_rooms = max(max_num_of_rooms, len(heap))
            
        return max_num_of_rooms