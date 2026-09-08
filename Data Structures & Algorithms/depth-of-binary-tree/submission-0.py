# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from collections import deque

class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        queue = deque([root])
        depth = 0

        while len(queue) > 0:
            level_size = len(queue)
            for _ in range(level_size):
                top = queue.popleft()
                
                if top.left is not None:
                    queue.append(top.left)
                
                if top.right is not None:
                    queue.append(top.right)

            depth += 1

        return depth