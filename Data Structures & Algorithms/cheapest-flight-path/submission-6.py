import heapq

class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        costs = [[float("inf")] * n for _ in range(k+2)]
        costs[0][src] = 0

        adj = [dict() for _ in range(n)]
        for (fro, to, price) in flights:
            adj[fro][to] = price

        heap = [(0, 0, src)]
        while len(heap) > 0:
            curr_price, stop_count, airport = heapq.heappop(heap)

            if curr_price > costs[stop_count][airport]:
                continue

            for next_airport, cost in adj[airport].items():
                new_cost = curr_price + cost
                
                if stop_count + 1 < k+2 and new_cost < costs[stop_count+1][next_airport]:
                    heapq.heappush(heap, (new_cost, stop_count+1, next_airport))
                    costs[stop_count+1][next_airport] = new_cost

        
        print(costs)
        min_cost = min(costs[i][dst] for i in range(k+2))
        return min_cost if min_cost != float('inf') else -1
        