import heapq

class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        MAX_INT = 2**31 - 1
        heap = []
        heapq.heappush(heap, (0, 0, src))

        connections = dict()
        for flight in flights:
            fro, to, cost = flight
            conn = connections.setdefault(fro, dict())
            conn[to] = cost

        costs = [[MAX_INT] * n for _ in range(k+2)]
        costs[0][src] = 0
            

        while len(heap) > 0:
            curr_price, stops, to = heapq.heappop(heap)

            if costs[stops][to] < curr_price:
                continue

            for conn in connections.get(to, dict()).items():
                new_dest, flight_price = conn
                if stops <= k and curr_price + flight_price < costs[stops+1][new_dest]:
                    heapq.heappush(heap, (curr_price + flight_price, stops+1, new_dest))
                    costs[stops+1][new_dest] = curr_price + flight_price


        cheapest_price = MAX_INT
        for i in range(k+2):
            cheapest_price = min(cheapest_price, costs[i][dst])

        if cheapest_price == MAX_INT:
            return -1

        return cheapest_price

        
