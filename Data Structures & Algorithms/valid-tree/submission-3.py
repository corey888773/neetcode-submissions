class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        parent = [i for i in range(n)]
        rank = [0 for i in range(n)]

        def find(x: int) -> int:
            while x != parent[x]:
                parent[x] = find(parent[x])
                x = parent[x]

            return parent[x]

        num_of_roots = n
        def union(x: int, y: int) -> bool:
            nonlocal num_of_roots
            root_x = find(x)
            root_y = find(y)

            if root_x == root_y:
                return True # has cycle

            if rank[root_x] > rank[root_y]:
                parent[root_y] = root_x
            else:
                parent[root_x] = root_y
                rank[root_y] += 1

            num_of_roots -= 1 
            return False 

        for e in edges:
            if union(e[0], e[1]):
                return False

        return num_of_roots == 1