class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        MAX_INT = 2**31 - 1

        costs = [MAX_INT] * n
        costs[src] = 0

        for _ in range(k+1):
            temp = costs.copy()

            for flight in flights:
                fro, to, cost = flight

                if costs[fro] + cost < temp[to]:
                    temp[to] = costs[fro] + cost

            costs = temp
            print(costs)

        if costs[dst] == MAX_INT:
            return -1

        return costs[dst]
