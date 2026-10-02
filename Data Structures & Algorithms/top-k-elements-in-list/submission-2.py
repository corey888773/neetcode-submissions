from collections import Counter

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        num_count= [(num, count) for num, count in Counter(nums).items()]
        num_count.sort(key=lambda x: (-x[1], x[0]))

        return [item[0] for item in num_count[:k]]


        
