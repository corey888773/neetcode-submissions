# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
import heapq

class Solution:    
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = []
        dummy = ListNode()

        for idx, node in enumerate(lists):
            if node:
                heapq.heappush(heap, (node.val, idx, node))

        curr = dummy
        idx = 3
        while len(heap) > 0:
            n = heapq.heappop(heap)
            next_node = n[2]
            
            if next_node.next:
                heapq.heappush(heap, (next_node.next.val, idx, next_node.next))

            curr.next = next_node
            curr = curr.next
            idx += 1

        return dummy.next
