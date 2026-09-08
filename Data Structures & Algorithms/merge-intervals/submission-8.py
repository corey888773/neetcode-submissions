class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: (x[0], x[1]))
        print(intervals)

        merged = []
        left, right = 0, 0
        while right < len(intervals):
            start = intervals[left][0]
            end = intervals[left][1]

            while right+1 < len(intervals) and intervals[right+1][0] <= end:
                right += 1
                end = max(end, intervals[right][1])

           
            merged.append([start, end])
            left = right = right+1

        return merged