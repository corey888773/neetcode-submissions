# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        slow = dummy
        fast = dummy

        while fast is not None:
            for _ in range(k):
                if fast is None:
                    return dummy.next
                
                fast = fast.next

            if fast is None:
                return dummy.next

            prev = self.reverse(slow.next, fast.next, k)

            next_slow = slow.next
            slow.next = prev # here i set the dummy to point to an end of the reversed k slice
            slow = next_slow
            fast = slow
        
        return dummy.next


    def reverse(self, curr: Optional[ListNode], prev: Optional[ListNode], k: int) -> Optional[ListNode]:
        for _ in range(k):
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        return prev



# node1 -> node2 -> node3

# node2 -> node1 -> node3


# nan > 1 > 2 > 3 > 4 > 5 > 6
#  nan > C 1 > N 2 > 3 > NN 4 > 5 > 6

# N = C.N
# C.N = P
# P = C
# C = N


# 3 > 2 > 1 > 4 > 5 > 6
                

            