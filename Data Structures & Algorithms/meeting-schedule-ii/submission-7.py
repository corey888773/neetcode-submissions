"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

class Solution:
    def minMeetingRooms(self, intervals: List[Interval]) -> int:
        events = []
        for i in intervals:
            events.extend([(i.start, 1), (i.end, -1)])

        events.sort(key=lambda x: (x[0], x[1]))

        num_of_rooms = 0
        max_num_of_rooms = 0
        for ev in events:
            num_of_rooms += ev[1]
            max_num_of_rooms = max(max_num_of_rooms, num_of_rooms)

        return max_num_of_rooms