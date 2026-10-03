from collections import deque

class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        adj = [dict() for _ in range(n)]
        for e in edges:
            fro, to = e
            adj[fro][to] = True
            adj[to][fro] = True

        visited = {0: True}
        queue = deque([(0, -1)])

        while len(queue) > 0:
            node, parent = queue.popleft()
            for neigh in adj[node]:
                if neigh not in visited:
                    queue.append((neigh, node))
                    visited[neigh] = True
                elif neigh != parent:
                    return False

        return len(visited) == n