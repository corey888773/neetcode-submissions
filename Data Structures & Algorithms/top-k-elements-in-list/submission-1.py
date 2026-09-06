class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        num_counter, freq = dict(), dict()
        top_k = []

        for n in nums:
            num_counter[n] = num_counter.get(n, 0) + 1

        for key, val in num_counter.items():
            freq.setdefault(val, []).append(key)

        for i in range(len(nums), -1, -1):
            if i in freq:
                for n in freq[i]:
                    top_k.append(n)

                    if len(top_k) == k:
                        return top_k

        return []
