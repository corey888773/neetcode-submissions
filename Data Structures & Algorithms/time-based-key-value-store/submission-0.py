from collections import defaultdict

class TimeMap:
    def __init__(self):
        self.store = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.store[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.store:
            return ""

        items = self.store[key]
        
        if items[0][0] > timestamp:
            return ""

        lo, hi = 0, len(items)
        while lo < hi:
            mid = lo + (hi - lo) // 2

            if items[mid][0] > timestamp:
                hi = mid
            else:
                lo = mid + 1

        return items[lo-1][1]
        


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)