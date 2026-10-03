class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x: (x[1], x[0]))

        curr = 0
        intervals_removed = 0
        for i in range(1, len(intervals)):
            if intervals[i][0] < intervals[curr][1]:
                intervals_removed += 1
            else:
                curr = i 

        return intervals_removed
