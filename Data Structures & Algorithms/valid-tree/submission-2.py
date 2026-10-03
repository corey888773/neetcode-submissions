class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        adj = [dict() for _ in range(n)]
        for fro, to in edges:
            adj[fro][to] = True
            adj[to][fro] = True


        visited = set()
        def dfs(node: int, parent: int) -> bool:
            visited.add(node)
            for neigh in adj[node]:
                if neigh not in visited:
                    if dfs(neigh, node): return True
                elif neigh != parent:
                    return True

            return False

        return not dfs(0, -1) and len(visited) == n