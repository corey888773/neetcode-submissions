class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        costs = [float("inf")] * n 
        costs[src] = 0
       
        for _ in range(k+1):
            temp = costs.copy()

            for (fro, to, cost) in flights:
                if costs[fro] + cost < temp[to]:
                    temp[to] = costs[fro] + cost

            costs = temp

        return costs[dst] if costs[dst] != float('inf') else -1
        
